# Consistency Model

This document defines the consistency model of Litsxg.

Consistency is weak by design.
No strong consistency guarantees are provided.
Predictability is preferred over correctness guarantees.

---

## Overview

Litsxg operates under a minimal consistency model.

The system does not attempt to:
- synchronize global state
- enforce ordering guarantees
- reconcile divergent views

Each component operates independently.

---

## Session consistency

Within a single active session:

- messages are processed sequentially
- state transitions follow documented rules

Beyond this, no consistency is guaranteed.

---

## Cross-session consistency

Across sessions:

- no state is shared
- no history is preserved
- no continuity exists beyond identity

Each session is independent.

---

## Client-server consistency

Between client and server:

- no state reconciliation occurs
- divergence is resolved by disconnect
- server state is authoritative

Clients adapt to server behavior.

---

## Multi-client consistency

When multiple clients interact:

- message delivery order is not guaranteed
- views may differ between clients
- delivery may be asymmetric

Consistency between clients is not enforced.

---

## Offline message consistency

Offline messages:

- may be delivered partially
- may be delivered out of order
- may be lost entirely

No attempt is made to ensure completeness.

---

## Failure-induced inconsistency

Failures may introduce inconsistency.

Examples include:
- lost messages
- duplicate delivery
- mismatched client views

Such inconsistency is accepted.

---

## Absence of convergence guarantees

Litsxg does not guarantee:

- eventual consistency
- strong consistency
- causal consistency

There is no convergence model.

---

## Identity consistency

Identity consistency is limited to:

- callsign uniqueness
- authentication-based association

Identity does not imply message continuity.

---

## Operator expectations

Operators must assume:

- inconsistent views are possible
- state divergence is normal
- no automatic reconciliation exists

Designing workflows that assume consistency is misuse.

---

## Scope enforcement

If a consistency property is not described here, it does not exist.

Operators must not infer stronger guarantees.

---

## Final statement

Litsxg accepts inconsistency as a consequence of simplicity.

The tool prioritizes explicit behavior over correctness models.
