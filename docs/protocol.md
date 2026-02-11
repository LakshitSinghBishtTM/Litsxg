# Protocol

This document defines the Litsxg protocol.

The protocol is intentionally simple.
It is line-based, plaintext, and synchronous.
There are no implicit behaviors.

If a behavior is not documented here, it must not be assumed.

---

## Overview

The Litsxg protocol defines communication between client and server.

Properties:
- plaintext
- line-oriented
- stateful per connection
- synchronous request/response behavior

There is no version negotiation.
There is no extensibility mechanism.

---

## Transport assumptions

The protocol operates over a reliable byte stream.

In practice:
- the transport is a TCP connection over Tor
- ordering is provided by the transport
- connection loss terminates the protocol session

The protocol does not attempt to recover from transport failure.

---

## Line-based framing

Protocol messages are framed as lines.

- Each message is terminated by a newline character.
- Lines are processed sequentially.
- Empty lines are ignored unless explicitly documented.

There is no binary framing.

---

## Session scope

A protocol session exists only for the lifetime of a connection.

Properties:
- authentication state is bound to the session
- session state is not transferable
- session state is lost on disconnect

Reconnection always starts a new session.

---

## Authentication requirement

No message routing occurs before authentication.

Until authentication completes:
- only authentication-related commands are processed
- all other input is ignored

Authentication rules are defined separately.

---

## Command and message distinction

The protocol distinguishes between:

- commands
- messages

Commands affect protocol state.
Messages are routed to other operators.

Commands are identified by explicit syntax.

---

## Operator identifiers

Operators are identified by a callsign.

Properties:
- callsigns are plaintext
- callsigns are meaningful only within the server
- callsigns are not cryptographically bound to a transport

Callsign validity rules are enforced by the server.

---

## Message addressing

Messages are addressed explicitly.

Each outbound message includes:
- a target callsign
- message content

The server uses the target callsign to route the message.

---

## Multiline message handling

The protocol supports multiline messages.

Multiline messages:
- use explicit delimiters
- are transmitted as a single logical message
- are not streamed incrementally

Multiline handling rules are defined separately.

---

## Error handling

The protocol has minimal error handling.

Properties:
- most protocol violations result in silent ignore or disconnect
- explicit error messages are not guaranteed
- clients must not rely on error feedback

Error handling behavior is deterministic but minimal.

---

## Absence of acknowledgements

The protocol does not include:

- delivery acknowledgements
- read receipts
- retry signals
- sequencing metadata

Delivery is best-effort.

---

## Absence of versioning

The protocol does not support version negotiation.

Client and server are expected to implement the same protocol definition.
Incompatible versions result in undefined behavior.

---

## Protocol invariants

The following invariants hold:

- plaintext is always visible to the server
- protocol state is per-connection
- no hidden negotiation occurs
- no recovery is attempted after failure

---

## Scope enforcement

Protocol behavior outside this document is undefined.

Operators must not assume undocumented commands, extensions, or guarantees.

---

## Final statement

The Litsxg protocol exists to be readable and predictable.

Its simplicity is intentional.
