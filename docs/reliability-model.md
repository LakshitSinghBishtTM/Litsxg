# Reliability Model

This document defines the reliability model of Litsxg.

Reliability is intentionally limited.
The tool prioritizes simplicity and predictability over guarantees.

---

## Overview

Litsxg operates under a best-effort reliability model.

This means:
- messages may be lost
- sessions may terminate
- state may be discarded

Reliability behavior is deterministic but minimal.

---

## Best-effort delivery

Litsxg attempts to deliver messages when possible.

Properties:
- no acknowledgements
- no retries
- no delivery confirmation

If delivery fails, the message is lost.

---

## Connection reliability

Connections are not assumed to be stable.

Properties:
- Tor circuits may fail
- connections may reset
- reconnection requires re-authentication

Connection loss is expected.

---

## Server availability

Server availability is not guaranteed.

Properties:
- no high-availability support
- no replication
- no failover

Availability depends entirely on operator infrastructure.

---

## Offline message reliability

Offline messages are buffered opportunistically.

Properties:
- buffering is in memory
- buffers are volatile
- delivery is not guaranteed

Offline buffering does not imply reliability.

---

## Ordering reliability

Message ordering is not guaranteed.

Properties:
- messages may arrive out of order
- ordering may differ across recipients
- ordering may differ across sessions

Ordering must not be relied upon.

---

## Persistence reliability

Persistence is minimal.

Properties:
- no durable message storage
- limited identity persistence
- volatile runtime state

Persistence exists only where necessary.

---

## Failure propagation

Failure propagates immediately.

Examples:
- server failure terminates all sessions
- client failure loses all local state

No attempt is made to contain failure.

---

## Absence of recovery mechanisms

Litsxg does not provide:

- replay
- reconciliation
- delayed retry
- state repair

Recovery is manual.

---

## Operator expectations

Operators must assume:

- message loss is possible
- delivery is not guaranteed
- failure may occur at any time

Designing workflows that assume reliability is misuse.

---

## Scope enforcement

If a reliability property is not documented here, it does not exist.

Operators must not assume hidden guarantees.

---

## Final statement

Litsxg is reliable only in its predictability.

It fails clearly and without concealment.
