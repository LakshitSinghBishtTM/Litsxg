# Litsxg

Litsxg is a lightweight messaging tool composed of a small server and a simple client.
It provides direct user-to-user messaging with minimal runtime overhead and explicit,
predictable behavior.

The project is intentionally simple in implementation and explicit in operation.

---

## Status

- Server: working
- Client / GUI: working
- Network: TCP (Tor onion service supported)

---

## Features

- Centralized messaging server
- Callsign-based user identities
- Password-based authentication
- Direct messaging between users
- Queued delivery for offline users
- Multiline message support
- Tor onion service exposure
- Minimal persistent state
- Small and readable codebase

---

## Project Layout

```
Litsxg/
├── server/              # Server source code
├── gui/                 # Client / GUI source code
├── docs/                # Detailed design and protocol documentation
├── README.md            # This file
└── LICENSE              # Open source
```

---

## Requirements

- Go (1.20 or newer recommended)
- SQLite3
- Tor

---

## Building

### Build the Server

```
cd server
go build -o litsxg-server
```

### Build the Client / GUI

```
cd gui
go build -o litsxg-gui
```

---

## Running the Server

```
./litsxg-server
```

By default, the server listens on:

```
TCP port 9696
```

On first startup, the server creates a SQLite database file named:

```
litsxg.db
```

in the current working directory. This database stores user credentials.

---

## Running the Client / GUI

```
./litsxg-gui
```

The client connects to a running Litsxg server using its address and port.
The server may be accessed over a local network or via a Tor onion address.

---

## Connection Model

* Clients connect to the server over TCP
* Each connection represents a single session
* A client must authenticate before sending messages
* Connections are closed on network failure or client exit

---

## Authentication

Users authenticate using a **callsign** and **password**.

### Authentication Command

```
$ <callsign> <password>
```

### Authentication Behavior

* If the callsign exists, the password is verified
* If the callsign does not exist:

  * The first attempt records intent
  * Repeating the same callsign confirms registration and creates the account

This two-step behavior prevents accidental account creation.

---

## Messaging

### Sending a Message

```
<target_callsign> <message>
```

Example:

```
ajay hello
```

* If the target user is online, the message is delivered immediately
* If the target user is offline, the message is queued in memory

Messages are delivered in the order they were sent.

---

## Multiline Messages

Multiline messages are supported using a block format.

Example:

```
ajay $ 
this is line one
this is line two
this is line three
$
```

Everything between the `$` markers is sent as a single message.

---

## Offline Message Handling

* Messages sent to offline users are stored in memory
* Pending messages are delivered when the user successfully authenticates
* Pending messages are cleared after delivery
* Messages are not persisted across server restarts

---

## Account Deletion

An authenticated user may delete their account using:

```
$ die
```

Behavior:

* The user entry is removed from the database
* The client session is terminated
* The callsign becomes available for reuse

---

## Server State

The server maintains the following state:

* Connected clients
* Pending messages for offline users
* Pending authentication attempts
* User credential database

All in-memory state is lost on server shutdown.

---

## Tor / Onion Service

The server can be exposed through Tor using a standard onion service configuration.

Example Tor configuration:

```
HiddenServiceDir /var/lib/tor/litsxg/
HiddenServicePort 9696 127.0.0.1:9696
```

After restarting Tor, the onion address is available at:

```
/var/lib/tor/litsxg/hostname
```

Clients can connect using this onion address and TCP port `9696`.

---

## Network Behavior

* The server does not retry failed deliveries
* No acknowledgements are implemented at the protocol level

---

## Data Handling

* User passwords are stored as SHA-256 hashes
* Messages are transmitted in plaintext
* Message contents are not stored on disk
* No encryption is performed at the application layer

---

## Documentation

The `docs/` directory contains detailed documentation covering:

* Protocol specification
* Message formats
* Authentication flow
* Multiline message semantics
* Failure and error handling
* Operational behavior
* Design decisions and scope

A reading order is provided for navigating the documentation.

---

## Development Notes

* The codebase is intentionally minimal
* Most complexity is documented explicitly
* Behavior is preferred over abstraction
* Changes should preserve existing protocol semantics

---

## License

Copyright (C) 2026 Lakshit Singh Bisht

This project is licensed under the GNU General Public License v3.0.
See LICENSE for details.

---

