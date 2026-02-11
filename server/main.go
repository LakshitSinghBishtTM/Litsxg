package main
import (
    "bufio"
    "crypto/sha256"
    "database/sql"
    "encoding/hex"
    "fmt"
    "net"
    "strings"
    "sync"
    _ "github.com/mattn/go-sqlite3"
)
type Client struct {
    callsign string
    conn     net.Conn
    writer   *bufio.Writer
}
type Server struct {
    clients      map[string]*Client
    pendingMsgs  map[string][]string
    pendingAuth  map[string]string
    mu           sync.RWMutex
    db           *sql.DB
}
func initDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./litsxg.db")
    if err != nil {
        return nil, err
    }
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            callsign TEXT PRIMARY KEY,
            password_hash TEXT NOT NULL
        )
    `)
    if err != nil {
        return nil, err
    }
    return db, nil
}
func hashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}
func (s *Server) userExists(callsign string) (bool, error) {
    var exists bool
    err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE callsign=?)", callsign).Scan(&exists)
    return exists, err
}
func (s *Server) verifyPassword(callsign, password string) (bool, error) {
    var storedHash string
    err := s.db.QueryRow("SELECT password_hash FROM users WHERE callsign=?", callsign).Scan(&storedHash)
    if err != nil {
        return false, err
    }
    return storedHash == hashPassword(password), nil
}
func (s *Server) createUser(callsign, password string) error {
    _, err := s.db.Exec("INSERT INTO users (callsign, password_hash) VALUES (?, ?)", callsign, hashPassword(password))
    return err
}
func (s *Server) deleteUser(callsign string) error {
    _, err := s.db.Exec("DELETE FROM users WHERE callsign=?", callsign)
    return err
}
func (s *Server) sendToClient(callsign, message string) {
    s.mu.RLock()
    client, online := s.clients[callsign]
    s.mu.RUnlock()
    if online {
        client.writer.WriteString(message + "\n")
        client.writer.Flush()
    } else {
        s.mu.Lock()
        s.pendingMsgs[callsign] = append(s.pendingMsgs[callsign], message)
        s.mu.Unlock()
    }
}
func (s *Server) deliverPendingMessages(callsign string) {
    s.mu.Lock()
    messages := s.pendingMsgs[callsign]
    delete(s.pendingMsgs, callsign)
    s.mu.Unlock()
    for _, msg := range messages {
        s.sendToClient(callsign, msg)
    }
}
func (s *Server) handleClient(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    writer := bufio.NewWriter(conn)
    var currentCallsign string
    var authenticated bool
    connAddr := conn.RemoteAddr().String()
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        line = strings.TrimRight(line, "\r\n")
        if strings.HasPrefix(line, "$ ") {
            parts := strings.SplitN(line[2:], " ", 2)
            if len(parts) == 1 && parts[0] == "die" {
                if !authenticated {
                    writer.Flush()
                    continue
                }
                s.deleteUser(currentCallsign)
                s.mu.Lock()
                delete(s.clients, currentCallsign)
                s.mu.Unlock()
                writer.Flush()
                conn.Close()
                return
            }
            if len(parts) == 2 {
                callsign := parts[0]
                password := parts[1]
                exists, _ := s.userExists(callsign)
                if exists {
                    valid, _ := s.verifyPassword(callsign, password)
                    if valid {
                        currentCallsign = callsign
                        authenticated = true
                        s.mu.Lock()
                        s.clients[callsign] = &Client{callsign: callsign, conn: conn, writer: writer}
                        delete(s.pendingAuth, connAddr)
                        s.mu.Unlock()
                        writer.Flush()
                        s.deliverPendingMessages(callsign)
                    } else {
                        writer.Flush()
                    }
                } else {
                    s.mu.RLock()
                    lastAttempt := s.pendingAuth[connAddr]
                    s.mu.RUnlock()
                    if lastAttempt == callsign {
                        err := s.createUser(callsign, password)
                        if err != nil {
                            writer.Flush()
                        } else {
                            currentCallsign = callsign
                            authenticated = true
                            s.mu.Lock()
                            s.clients[callsign] = &Client{callsign: callsign, conn: conn, writer: writer}
                            delete(s.pendingAuth, connAddr)
                            s.mu.Unlock()
                            writer.Flush()
                        }
                    } else {
                        s.mu.Lock()
                        s.pendingAuth[connAddr] = callsign
                        s.mu.Unlock()
                        writer.Flush()
                    }
                }
            }
            continue
        }
        if !authenticated {
            writer.Flush()
            continue
        }
        parts := strings.SplitN(line, " ", 2)
        if len(parts) < 2 {
            continue
        }
        target := parts[0]
        content := parts[1]
        if strings.HasPrefix(content, "$ ") {
            multilineMsg := content[2:]
            for {
                nextLine, err := reader.ReadString('\n')
                if err != nil {
                    break
                }
                nextLine = strings.TrimRight(nextLine, "\r\n")
                if nextLine == "$" {
                    break
                }
                if multilineMsg == "" {
                    multilineMsg = nextLine
                } else {
                    multilineMsg += "\n" + nextLine
                }
            }
            content = multilineMsg
        }
        message := fmt.Sprintf("%s: %s", currentCallsign, content)
        s.sendToClient(target, message)
    }
    if authenticated {
        s.mu.Lock()
        delete(s.clients, currentCallsign)
        s.mu.Unlock()
    }
    s.mu.Lock()
    delete(s.pendingAuth, connAddr)
    s.mu.Unlock()
}
func main() {
    db, err := initDB()
    if err != nil {
        return
    }
    defer db.Close()
    server := &Server{clients: make(map[string]*Client), pendingMsgs: make(map[string][]string), pendingAuth: make(map[string]string), db: db}
    listener, err := net.Listen("tcp", ":9696")
    if err != nil {
        return
    }
    defer listener.Close()
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go server.handleClient(conn)
    }
}