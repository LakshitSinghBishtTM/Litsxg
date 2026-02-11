# Error Handling

This document defines how errors are handled in Litsxg.

Error handling is intentionally minimal.
Errors are not abstracted, hidden, or softened.
Most errors result in termination of the affected operation.

---

## Overview

Litsxg treats errors as normal operational events.

The tool does not attempt to:
- classify errors for operators
- recover silently
- provide diagnostic guidance through the interface

Error handling exists to preserve predictable behavior, not continuity.

---

## Client-side error handling

### Connection errors

Connection errors include:
- inability to establish a Tor connection
- connection loss during operation
- failure to reach the server

Handling behavior:
- the active session is terminated
- local connection state is cleared
- reconnection attempts begin automatically

No warning or explanation is presented.

---

### Protocol errors

Protocol errors include:
- invalid message formats
- unexpected command sequences
- protocol violations

Handling behavior:
- input may be ignored
- the session may be terminated
- behavior may be undefined

The client does not attempt to correct protocol misuse.

---

### Write failures

Write failures include:
- inability to transmit data
- partial writes
- broken connections

Handling behavior:
- the connection is treated as failed
- local state is reset
- buffered input is discarded

There are no retries.

---

### Interface errors

Interface-level errors are not surfaced.

The client assumes:
- the interface reflects internal state
- failures are reflected indirectly through behavior

No dialogs, alerts, or prompts are shown.

---

## Server-side error handling

### Authentication errors

Authentication errors include:
- invalid credentials
- identity conflicts
- incomplete authentication flows

Handling behavior:
- identity association is refused
- unauthenticated state is maintained
- detailed error messages are not guaranteed

Authentication errors do not terminate the server.

---

### Routing errors

Routing errors include:
- unknown target callsigns
- disconnected recipients

Handling behavior:
- messages may be buffered
- messages may be dropped
- no delivery feedback is guaranteed

Routing errors are not surfaced to senders.

---

### Internal server errors

Internal errors include:
- database access failure
- in-memory state inconsistency
- unexpected runtime conditions

Handling behavior:
- the affected connection may terminate
- the server process may terminate
- in-memory state may be lost

The server does not attempt state repair.

---

## Absence of error reporting

Litsxg does not provide:

- structured error codes
- human-readable error explanations
- machine-readable diagnostics

Error context is not transmitted over the protocol.

---

## Absence of recovery mechanisms

Litsxg does not implement:

- automatic retries
- rollback mechanisms
- state reconciliation
- deferred error handling

Recovery is manual and external.

---

## Error visibility

Errors are visible only through their effects.

Examples:
- message loss
- session termination
- connection reset

Operators must infer cause through observation and logs.

---

## Scope enforcement

If an error behavior is not described here, it is undefined.

Operators must not rely on undocumented error handling.

---

## Final statement

Error handling in Litsxg favors clarity over comfort.

Failure is surfaced by termination, not explanation.
