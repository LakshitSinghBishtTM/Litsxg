# Design Philosophy

This document describes the principles that guide the design of Litsxg.

These principles are fixed.
They explain why the tool behaves as it does.
They are not subject to revision.

---

## Tool-first design

Litsxg is designed as a tool.

It is not a platform.
It is not a framework.
It is not a service ecosystem.

The tool exists to perform a narrow function reliably and predictably.
Anything outside that function is out of scope.

---

## Deliberate minimalism

Litsxg is intentionally minimal.

- Fewer features reduce ambiguity.
- Fewer abstractions reduce failure modes.
- Fewer assumptions reduce misuse.

Minimalism here is not aesthetic.
It is a method for maintaining control over behavior.

---

## Explicit boundaries over implicit guarantees

Litsxg favors explicit boundaries.

If a property is not stated, it does not exist.
If a guarantee is not written, it must not be assumed.

The tool does not attempt to imply safety, reliability, or confidentiality through design complexity.

---

## Closed design

Litsxg has a closed design.

The feature set is complete.
There is no extension mechanism.
There is no roadmap.

Changes are limited to correcting defects that violate documented behavior.

---

## Simplicity over coverage

Litsxg does not attempt to cover all use cases.

It prefers:
- simple message routing
- predictable failure
- understandable behavior

Coverage of edge cases is secondary to clarity of core behavior.

---

## Server-visible state is acceptable

Litsxg accepts server-visible state.

The server:
- knows operator identities
- sees message contents
- manages routing state

This is not treated as a flaw to be hidden.
It is a design decision.

---

## Tor as infrastructure, not a promise

Tor is used as infrastructure.

It provides:
- reachability
- deployment convenience
- resistance to casual network observation

Tor is not used to claim anonymity or strong privacy guarantees.

---

## Failure as a normal condition

Litsxg treats failure as normal.

Network loss, crashes, and partial delivery are expected.
The tool does not attempt to mask these conditions.

Predictable failure is preferred over complex recovery logic.

---

## Operator responsibility

Litsxg assumes operator responsibility.

Operators are expected to:
- read documentation
- understand boundaries
- deploy and operate infrastructure correctly

The tool does not compensate for incorrect assumptions.

---

## Final principle

Litsxg is designed to do a small number of things clearly.

Its philosophy is to state limits explicitly and remain within them.
