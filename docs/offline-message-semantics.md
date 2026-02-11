# Offline Message Semantics

This document defines how Litsxg handles messages sent to offline operators.

Offline messaging is opportunistic.
It does not provide durability or delivery guarantees.

---

## Overview

Litsxg allows messages to be sent to operators who are not currently connected.

Such messages may be buffered temporarily.
Buffering exists only to improve convenience.

---

## Definition of offline

An operator is considered offline when:

- no active authenticated connection exists
- the server has no live session associated with the callsign

Offline status is determined solely by server state.

---

## Buffering behavior

When a message is sent to an offline operator:

- the server may buffer the message in memory
- buffering is per-callsign
- buffering occurs only while the server process remains alive

There is no persistence to disk.

---

## Buffer lifetime

Buffered messages exist:

- until the recipient authenticates
- until the server restarts
- until the server crashes
- until memory is reclaimed

Any of these events may result in loss.

---

## Delivery attempt

On successful authentication:

- the server attempts to deliver buffered messages
- delivery order is not guaranteed
- delivery may be partial or empty

Failure during delivery may result in loss.

---

## No delivery confirmation

There is no confirmation that:

- buffered messages were delivered
- messages were displayed by the client
- messages were read by the operator

Delivery is silent.

---

## Interaction with network failure

Network failure may occur during delivery.

In such cases:
- buffered messages may be lost
- delivery may stop mid-stream
- no retry occurs

There is no rollback.

---

## Multiline message handling

Offline multiline messages:

- are buffered as single logical messages
- are not streamed
- are delivered whole or not at all

Partial multiline delivery does not occur.

---

## No queue management

Litsxg does not provide:

- queue limits
- queue inspection
- queue management commands

Queues are internal and opaque.

---

## No ordering guarantees

Buffered messages:

- may be delivered out of order
- may interleave with live messages
- may be dropped selectively

Ordering must not be assumed.

---

## No backpressure

The server does not apply backpressure.

Messages sent to offline operators:
- may accumulate
- may exhaust memory
- may cause instability

This is not mitigated.

---

## Operator expectations

Operators must assume:

- offline messages may be lost
- offline delivery is best-effort
- offline messaging is a convenience feature

Relying on offline messaging for reliability is misuse.

---

## Scope enforcement

Offline message behavior outside this document is undefined.

No additional semantics exist.

---

## Final statement

Offline messaging in Litsxg is intentionally weak.

It exists to reduce friction, not to provide guarantees.
