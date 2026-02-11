# Network Failure Semantics

This document defines how Litsxg behaves under network failure.

Network failure is expected.
It is not treated as exceptional.
The tool does not attempt to mask or compensate for it.

---

## Overview

Network failure includes any condition where communication between client and server is interrupted.

Examples include:
- Tor circuit loss
- temporary network unavailability
- server restart
- client process termination

Failure semantics are deterministic.

---

## Failure detection

Network failure is detected implicitly.

Indicators include:
- failed reads
- failed writes
- closed connections
- timeouts at the transport layer

There is no explicit failure signaling protocol.

---

## Client-side behavior on failure

When a network failure occurs, the client:

- terminates the active connection
- discards session state
- clears authentication state
- discards buffered input
- resumes connection retry behavior

No attempt is made to preserve state.

---

## Server-side behavior on failure

When a network failure occurs, the server:

- terminates the affected connection
- discards per-connection state
- removes identity association for that connection

Server operation for other connections continues unaffected.

---

## Message loss semantics

Messages may be lost due to network failure.

Loss may occur:
- before reaching the server
- while buffered in memory
- before delivery to a recipient

There is no detection or reporting of message loss.

---

## Offline message interaction

Offline message buffering does not alter failure semantics.

Buffered messages:
- may be lost on server restart
- are not acknowledged
- are not retried

Network failure may prevent offline delivery entirely.

---

## Authentication impact

Network failure invalidates authentication state.

On reconnect:
- authentication must be repeated
- identity continuity depends on successful re-authentication

There is no session resumption.

---

## Multiline message failure

Failure during multiline input results in:

- loss of all buffered multiline content
- exit from multiline mode
- no partial transmission

Operators must re-enter content manually.

---

## Timing considerations

There are no timing guarantees.

Network recovery time:
- is variable
- is not estimated
- is not communicated to the operator

The client retries indefinitely.

---

## Absence of recovery guarantees

Litsxg does not provide:

- automatic recovery
- state reconciliation
- partial message repair

Recovery is manual and external.

---

## Operator expectations

Operators must assume that:

- network failure can occur at any time
- message loss is possible
- sessions can terminate without warning

Operating under different assumptions is unsupported.

---

## Scope enforcement

Failure semantics outside this document are undefined.

Operators must not infer additional behavior.

---

## Final statement

Litsxg treats network failure as a normal condition.

Correct operation requires accepting loss and retry as fundamental properties.