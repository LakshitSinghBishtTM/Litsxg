# Protocol Non-Guarantees

This document defines what the Litsxg protocol explicitly does not guarantee.

Non-guarantees are part of the protocol definition.
They are not deficiencies.
They will not change.

---

## Overview

The Litsxg protocol is intentionally minimal.

As a result, it does not provide many properties commonly assumed in modern messaging systems.
Operators must not infer such properties from protocol simplicity.

---

## No delivery guarantees

The protocol does not guarantee:

- message delivery
- eventual delivery
- delivery within any time bound

Messages may be lost due to:
- network failure
- server restart
- client disconnect
- routing conditions

Loss is acceptable.

---

## No ordering guarantees

The protocol does not guarantee:

- message ordering
- preservation of send order
- consistent ordering across recipients

Messages may arrive out of order or not at all.

---

## No acknowledgement semantics

The protocol does not include:

- delivery acknowledgements
- read receipts
- confirmation messages
- negative acknowledgements

Senders receive no feedback regarding message outcome.

---

## No retry semantics

The protocol does not implement:

- automatic retries
- retransmission logic
- backoff strategies

If a message is lost, it is lost.

---

## No persistence guarantees

The protocol does not guarantee:

- message persistence
- storage durability
- survival across server restarts

Any persistence that exists is temporary and opportunistic.

---

## No session continuity

The protocol does not guarantee:

- session resumption
- authentication continuity across connections
- state recovery after disconnect

Each connection represents a new protocol session.

---

## No version compatibility guarantees

The protocol does not guarantee:

- backward compatibility
- forward compatibility
- negotiated protocol versions

Client and server implementations must match.

---

## No error signaling guarantees

The protocol does not guarantee:

- explicit error messages
- structured error reporting
- consistency of error behavior

Errors may be silent.

---

## No extensibility guarantees

The protocol does not guarantee:

- future extension points
- compatibility with added commands
- safe evolution

The protocol is closed.

---

## Assumption discipline

Operators must not assume:

- implied guarantees
- best-effort behavior beyond what is stated
- recovery semantics

Absence of a guarantee must be treated as explicit.

---

## Final statement

The Litsxg protocol defines what it does.

Everything else is intentionally undefined.
