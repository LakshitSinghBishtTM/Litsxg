package main
import(
	"bufio"
	_ "embed"
	"image/color"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/net/proxy"
)
//go:embed tor-bundle/tor/tor
var torBinary []byte
type darkTheme struct{}
func(d darkTheme)Color(n fyne.ThemeColorName,v fyne.ThemeVariant)color.Color{
	switch n{
	case theme.ColorNameBackground,theme.ColorNameInputBackground,theme.ColorNameFocus:
		return color.Black
	case theme.ColorNameForeground,theme.ColorNameDisabled,theme.ColorNamePrimary:
		return color.White
	}
	return theme.DefaultTheme().Color(n,v)
}
func(d darkTheme)Font(s fyne.TextStyle)fyne.Resource{
	return theme.DefaultTheme().Font(fyne.TextStyle{Monospace:true})
}
func(d darkTheme)Icon(n fyne.ThemeIconName)fyne.Resource{
	return theme.DefaultTheme().Icon(n)
}
func(d darkTheme)Size(n fyne.ThemeSizeName)float32{
	return theme.DefaultTheme().Size(n)
}
type GuiClient struct{
	out *widget.Entry
	in *widget.Entry
	lines []string
	conn net.Conn
	writer *bufio.Writer
	torCmd *exec.Cmd
	socks string
	mu sync.Mutex
	multiline bool
}
func(c *GuiClient)append(line string){
	c.mu.Lock()
	c.lines=append(c.lines,line)
	text:=strings.Join(c.lines,"\n")
	c.mu.Unlock()
	fyne.Do(func(){c.out.SetText(text)})
}
func freePort()(int,error){
	l,err:=net.Listen("tcp","127.0.0.1:0")
	if err!=nil{return 0,err}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port,nil
}
func extractTor()(string,error){
	cache,_:=os.UserCacheDir()
	dir:=filepath.Join(cache,"litsxg","tor")
	os.MkdirAll(dir,0700)
	bin:=filepath.Join(dir,"tor")
	if runtime.GOOS=="windows"{bin+=".exe"}
	if _,err:=os.Stat(bin);err==nil{return bin,nil}
	if err:=os.WriteFile(bin,torBinary,0700);err!=nil{return "",err}
	return bin,nil
}
func(c *GuiClient)startEmbeddedTor()error{
	p,err:=freePort()
	if err!=nil{return err}
	bin,err:=extractTor()
	if err!=nil{return err}
	c.socks="127.0.0.1:"+strconv.Itoa(p)
	c.torCmd=exec.Command(bin,"--SocksPort",c.socks)
	c.torCmd.Stdout=io.Discard
	c.torCmd.Stderr=io.Discard
	if err:=c.torCmd.Start();err!=nil{return err}
	for i:=0;i<60;i++{
		conn,err:=net.DialTimeout("tcp",c.socks,time.Second)
		if err==nil{conn.Close();return nil}
		time.Sleep(500*time.Millisecond)
	}
	return err
}
func tryConnect(addr,socks string)(net.Conn,error){
	d,err:=proxy.SOCKS5("tcp",socks,nil,proxy.Direct)
	if err!=nil{return nil,err}
	return d.Dial("tcp",addr)
}
func(c *GuiClient)attachConn(conn net.Conn){
	c.mu.Lock()
	c.conn=conn
	c.writer=bufio.NewWriter(conn)
	c.mu.Unlock()
	sc:=bufio.NewScanner(conn)
	sc.Buffer(make([]byte,1024),1024*1024)
	for sc.Scan(){
		c.append("> "+sc.Text())
	}
	c.mu.Lock()
	c.writer=nil
	c.conn=nil
	c.multiline=false
	c.mu.Unlock()
}
func(c *GuiClient)send(s string){
	c.mu.Lock()
	w:=c.writer
	c.mu.Unlock()
	if w==nil{return}
	if _,err:=w.WriteString(s+"\n");err!=nil{
		c.mu.Lock()
		c.writer=nil
		c.mu.Unlock()
		return
	}
	w.Flush()
}
func(c *GuiClient)connectLoop(addr string){
	for{
		var conn net.Conn
		var err error
		conn,err=tryConnect(addr,"127.0.0.1:9050")
		if err!=nil{
			if c.socks==""{
				if err:=c.startEmbeddedTor();err!=nil{
					time.Sleep(time.Second)
					continue
				}
			}
			conn,err=tryConnect(addr,c.socks)
		}
		if err!=nil{
			c.socks=""
			time.Sleep(2*time.Second)
			continue
		}
		c.attachConn(conn)
		time.Sleep(2*time.Second)
	}
}
func main(){
	a:=app.New()
	a.Settings().SetTheme(&darkTheme{})
	w:=a.NewWindow("Litsxg")
	client:=&GuiClient{lines:[]string{}}
	out:=widget.NewMultiLineEntry()
	out.Disable()
	out.Wrapping=fyne.TextWrapOff
	out.TextStyle=fyne.TextStyle{Monospace:true}
	in:=widget.NewMultiLineEntry()
	in.TextStyle=fyne.TextStyle{Monospace:true}
	in.Wrapping=fyne.TextWrapOff
	in.OnSubmitted=func(s string){
		if s==""{return}
		if s=="$"&&client.multiline{
			client.send("$")
			client.multiline=false
			in.SetText("")
			return
		}
		if strings.HasPrefix(s,"$ ")&&!client.multiline{
			client.multiline=true
			client.append("> "+s)
			client.send(s)
			in.SetText("")
			return
		}
		if client.multiline{
			client.append("> "+s)
			client.send(s)
			in.SetText("")
			return
		}
		client.append("> "+s)
		client.send(s)
		in.SetText("")
	}
	client.out=out
	client.in=in
	bg:=canvas.NewRectangle(color.Black)
	w.SetContent(container.NewBorder(nil,in,nil,nil,container.NewMax(bg,out)))
	go client.connectLoop("<write_your_onion_address_here>:9696")
	w.Resize(fyne.NewSize(720,520))
	w.ShowAndRun()
}
