# Server Architecture

This document describes the internal structure and behavior of the Litsxg server.

It defines responsibilities, state ownership, and routing behavior.
It does not assume an adversarial server model.

---

## Server role

The server is the central routing component.

Its role is to:
- accept incoming connections
- authenticate operators
- route messages between operators
- maintain limited transient state
- expose predictable behavior

The server is trusted by design.

---

## Architectural overview

The server is organized around the following functional areas:

- connection handling
- authentication management
- message routing
- offline message buffering
- state management

Each area is deliberately simple and tightly scoped.

---

## Connection handling

Connection handling is responsible for:

- accepting inbound connections
- associating connections with operators
- tracking active sessions
- cleaning up state on disconnect

Connections are long-lived but not guaranteed to persist.

---

## Authentication management

Authentication management is responsible for:

- verifying operator credentials
- establishing identity continuity
- preventing identity collisions

Authentication:
- is synchronous
- is stateful
- does not provide strong security guarantees

Authentication failure results in refusal to associate identity.

---

## Message routing

Message routing is responsible for:

- receiving plaintext messages from authenticated operators
- determining target operators
- delivering messages to active sessions
- enqueueing messages for offline operators

Routing behavior is deterministic.

---

## Offline message buffering

Offline message buffering is provided as a convenience.

Buffered messages:
- exist in server-managed memory
- may be lost on server restart or failure
- are not durable

Buffering does not imply delivery guarantees.

---

## State management

Server state includes:

- active operator sessions
- authentication mappings
- offline message buffers

State persistence is intentionally limited.
Only essential identity data is persisted.

---

## Persistence layer

The server uses a minimal persistence layer for:

- storing operator credentials
- maintaining identity continuity

Persistence is not used for:
- message storage
- delivery tracking
- audit logging

---

## Concurrency model

The server handles multiple connections concurrently.

Concurrency is managed through:
- explicit synchronization
- limited shared state
- predictable execution paths

The server does not attempt to serialize all operations globally.

---

## Failure handling

Server failure is treated as normal.

On failure:
- active sessions terminate
- buffered messages may be lost
- clients must reconnect

The server does not attempt graceful recovery beyond restart.

---

## Scope enforcement

The server enforces scope by refusing to implement features outside its role.

Requests outside protocol definition are ignored or rejected.

---

## Final statement

The Litsxg server prioritizes predictability over complexity.

Its architecture exists to route messages clearly and consistently, nothing more.
