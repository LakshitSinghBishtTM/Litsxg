# Lifecycle

This document defines the lifecycle of the Litsxg tool.

Lifecycle stages are fixed.
There is no dynamic reconfiguration or adaptive behavior.
Each stage transitions deterministically to the next.

---

## Overview

Litsxg has simple, linear lifecycles for both client and server components.

Lifecycle management prioritizes:
- predictability
- explicit transitions
- absence of hidden states

There is no global lifecycle controller.

---

## Client lifecycle

### Startup

On startup, the client:

- initializes the graphical interface
- prepares local state
- begins connection attempts to the server

No state is restored from previous runs.

---

### Connection establishment

During connection establishment, the client:

- attempts to connect using system Tor if available
- falls back to embedded Tor if configured
- retries indefinitely with fixed behavior

Connection attempts do not require operator confirmation.

---

### Authentication

After establishing a connection, the client:

- initiates authentication
- waits for server acceptance or rejection

Authentication must succeed before message exchange.

---

### Active operation

During active operation, the client:

- accepts operator input
- transmits messages to the server
- displays received messages
- manages multiline input mode

The client remains in this stage until failure or termination.

---

### Failure handling

On failure, including:

- network interruption
- Tor circuit loss
- server restart
- protocol error

the client:

- terminates the active session
- clears connection state
- returns to connection establishment

No state is preserved across failure.

---

### Shutdown

On shutdown, the client:

- terminates network connections
- releases resources
- exits without cleanup guarantees

There is no shutdown handshake.

---

## Server lifecycle

### Startup

On startup, the server:

- initializes persistent storage
- prepares in-memory state
- begins listening for connections

No prior session state is restored.

---

### Accepting connections

During operation, the server:

- accepts incoming connections
- assigns per-connection execution contexts
- processes authentication and messages

Connections are handled independently.

---

### Normal operation

During normal operation, the server:

- routes messages
- buffers offline messages temporarily
- manages active session mappings

There is no background maintenance process.

---

### Failure and restart

On failure or restart:

- all active connections terminate
- in-memory state is lost
- offline messages are discarded

After restart, the server resumes from a clean state.

---

### Shutdown

On shutdown, the server:

- stops accepting new connections
- terminates existing connections
- exits without draining buffers

Graceful shutdown is not guaranteed.

---

## Absence of upgrade lifecycle

Litsxg does not support live upgrades.

Upgrades require:
- stopping the process
- replacing the binary
- restarting the tool

No state migration is performed.

---

## Lifecycle invariants

Across all lifecycles:

- state is ephemeral
- transitions are explicit
- recovery is manual

These invariants do not change.

---

## Final statement

The Litsxg lifecycle is intentionally simple.

Operational clarity is achieved by limiting transitions and preserving explicit control.
