# Crash Recovery

This document defines crash recovery behavior in Litsxg.

Crash recovery is minimal.
The tool does not attempt state reconstruction or continuity.
Restart is the only recovery mechanism.

---

## Overview

A crash is any unexpected termination of the client or server process.

Examples include:
- process crash
- forced termination
- system restart
- power loss

Crash handling behavior is deterministic and simple.

---

## Client crash behavior

When the client crashes:

- the network connection is terminated
- authentication state is lost
- buffered input is lost
- multiline input state is lost

No client state is preserved.

---

## Client restart behavior

On restart, the client:

- starts with a clean state
- initiates connection attempts
- requires re-authentication

There is no crash recovery logic.

---

## Server crash behavior

When the server crashes:

- all client connections terminate
- all in-memory state is lost
- offline message buffers are discarded

Persistent identity data remains intact.

---

## Server restart behavior

On restart, the server:

- initializes fresh in-memory state
- reloads persistent identity data
- accepts new connections

No attempt is made to restore sessions or messages.

---

## Message loss on crash

Messages may be lost due to crash:

- messages in transit are lost
- buffered offline messages are lost
- partially delivered messages are lost

There is no detection or reporting of loss.

---

## Authentication impact

Crash invalidates authentication state.

After restart:
- all operators must re-authenticate
- identity continuity depends on credential correctness

There is no session carryover.

---

## No crash consistency guarantees

Litsxg does not provide:

- atomic operations across crash
- transaction semantics
- write-ahead logging

Crash consistency is not a goal.

---

## No automatic restart handling

Litsxg does not include:

- watchdog processes
- automatic restart logic
- health monitoring

Restart is external.

---

## Operator responsibility

Operators are responsible for:

- running the server under appropriate supervision
- restarting processes after crashes
- accepting message loss

Crash tolerance is externalized.

---

## Scope enforcement

If crash recovery behavior is not described here, it does not exist.

Operators must not assume hidden recovery logic.

---

## Final statement

Litsxg does not recover from crashes.

It restarts cleanly and continues operation from that point.
