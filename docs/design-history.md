# Design History

This document records how Litsxg arrived at its current design.

The history is descriptive.
It does not justify changes.
It does not imply evolution.

---

## Overview

Litsxg did not evolve through incremental feature growth.

Its design converged toward reduction.
Features were removed until only necessary behavior remained.

The current design is the result of deliberate constraint.

---

## Initial motivations

Litsxg originated from a need for:

- direct operator-to-operator messaging
- minimal infrastructure requirements
- explicit behavior without abstraction
- self-hosted control

Early goals did not include safety guarantees or user comfort.

---

## Early design choices

From early iterations, the following choices were fixed:

- centralized server model
- plaintext protocol
- explicit authentication
- line-based messaging

These choices were never treated as temporary.

---

## Rejected design directions

Several directions were considered and rejected early:

- end-to-end encryption
- client-side key management
- protocol extensibility
- rich user interfaces
- message history and persistence

These were rejected to preserve clarity and scope.

---

## Tor integration decision

Tor was introduced as:

- a deployment convenience
- a reachability mechanism
- a way to avoid clearnet exposure

Tor was never intended to redefine the security model.

---

## GUI inclusion

A graphical client was included to:

- reduce friction for operators
- provide direct interaction
- avoid command-line exclusivity

The GUI was intentionally kept silent and minimal.

---

## Authentication simplification

Authentication was simplified to:

- identify operators
- prevent accidental identity collision

No attempt was made to harden authentication beyond this.

---

## Removal of safeguards

Several potential safeguards were explicitly removed:

- warnings
- confirmations
- retries
- recovery mechanisms

Safeguards were seen as obscuring behavior.

---

## Finalization point

The design was considered finished when:

- all documented behavior matched implementation
- no feature depended on future work
- no security claims exceeded reality

At this point, development shifted to maintenance only.

---

## No historical branching

There are no alternative design branches.

Abandoned ideas remain abandoned.
They are not scheduled for reconsideration.

---

## Historical permanence

Design history informs understanding only.

It does not create obligations to:
- revisit decisions
- accommodate feedback
- evolve the tool

History explains why the tool is fixed.

---

## Operator interpretation

Operators should read this history to:

- understand intent
- avoid proposing rejected ideas
- align expectations

Disagreement does not change the design.

---

## Final statement

Litsxg is the result of deliberate reduction.

Its history ends where its design stabilizes.
