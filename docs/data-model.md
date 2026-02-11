# Data Model

This document defines the data model used by the Litsxg tool.

The data model is intentionally minimal.
Only data required for operation is stored.
No additional data is inferred or retained.

---

## Overview

Litsxg uses a small data model divided between:

- persistent server data
- in-memory server state
- ephemeral client state

There is no shared or synchronized data store between client and server.

---

## Persistent server data

Persistent data exists only to maintain identity continuity.

### Operator credentials

The server stores:

- operator identifier
- password hash

This data is stored in a local database.

Properties:
- plaintext passwords are not stored
- hashes are not intended to resist determined offline attacks
- no password metadata is retained

Credential storage exists solely to distinguish operators.

---

## In-memory server state

The server maintains the following in memory:

### Active sessions

- mapping of operator identifiers to active connections
- session lifetime tied to connection lifetime

Sessions are destroyed on disconnect or server restart.

---

### Offline message buffers

- temporary message queues for offline operators
- stored in memory only

Properties:
- not durable
- not ordered across restarts
- not guaranteed to be delivered

Buffers are cleared on server restart.

---

### Authentication state

- temporary state used during authentication flow
- cleared after authentication completes or fails

This state is not persisted.

---

## Client-side data

The client maintains ephemeral local data:

- current connection status
- authentication state
- input mode
- in-memory display buffer

The client does not persist:

- messages
- credentials
- session identifiers

All client data is lost on restart.

---

## Absence of message persistence

Litsxg does not persist message content to disk.

Messages exist only:
- in transit
- in temporary server buffers
- in client display memory

There is no message history store.

---

## Absence of metadata storage

Litsxg does not store:

- message timestamps for audit
- delivery receipts
- read acknowledgements
- operator activity logs

Only data required for live routing is maintained.

---

## Data lifecycle

Data lifecycle is simple:

1. Data is created when required for operation.
2. Data exists only for the duration of its usefulness.
3. Data is discarded without archival.

There is no retention policy beyond immediate operation.

---

## Data loss expectations

Data loss is expected and acceptable.

Loss may occur due to:
- server restart
- process crash
- network failure
- client termination

No recovery mechanisms are provided.

---

## Scope enforcement

If data is not described in this document, it does not exist.

Operators must not assume hidden state, backups, or recovery.

---

## Final statement

The Litsxg data model is intentionally sparse.

Data exists only to enable message routing and identity continuity.
Anything beyond that is out of scope.
