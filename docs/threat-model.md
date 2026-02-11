# Threat Model

This document defines the threat model for Litsxg.

The threat model is intentionally narrow.
Only explicitly listed threats are considered.
Anything outside this model is not defended against.

---

## Purpose

The purpose of this document is to state, unambiguously:

- which threats Litsxg considers
- which threats Litsxg ignores
- where responsibility lies

This is not a risk assessment.
It is a boundary definition.

---

## Assets

The following assets exist within Litsxg:

- operator identities (callsigns)
- plaintext message content
- message routing metadata
- server availability

No other assets are assumed.

---

## Trusted components

The following components are trusted by design:

- the server process
- the server operator
- the local execution environment

Trust in these components is required for operation.

---

## Adversaries considered

Litsxg considers the following adversaries:

### Casual network observers

Examples:
- local network observers
- passive ISP-level observers

Mitigation:
- Tor transport provides resistance to casual observation

---

### Curious but non-malicious parties

Examples:
- accidental log readers
- administrators without active attack intent

Mitigation:
- basic password hashing
- minimal data retention

These mitigations are limited.

---

### Unreliable infrastructure

Examples:
- unstable networks
- Tor circuit failure
- process crashes

Mitigation:
- explicit failure handling
- acceptance of loss

---

## Adversaries not considered

Litsxg does not defend against the following:

### Malicious server operators

The server can:
- read messages
- modify routing
- impersonate operators

This is accepted.

---

### Compromised servers

If the server is compromised:
- all data is exposed
- behavior is undefined

There is no mitigation.

---

### Active network attackers

Examples:
- traffic correlation
- active manipulation
- protocol injection

These threats are out of scope.

---

### Targeted attacks

Examples:
- brute-force authentication attacks
- credential theft
- denial-of-service attacks

Litsxg does not implement defenses against these.

---

### Legal or coercive threats

Examples:
- subpoenas
- coercion
- compelled disclosure

Litsxg does not attempt resistance.

---

## Threats explicitly ignored

The following are explicitly ignored:

- metadata analysis resistance
- forward secrecy
- post-compromise security
- message deniability
- plausible deniability
- censorship resistance beyond Tor reachability

---

## Operator responsibility

Operators are responsible for:

- evaluating whether this threat model is acceptable
- deploying the tool only within this scope
- not relying on unclaimed protections

Operating outside this model is unsupported.

---

## No threat model evolution

The threat model is fixed.

There are no plans to:
- expand defended threats
- introduce new mitigations
- revise assumptions

---

## Final statement

Litsxg defends against very little.

This is intentional.

Understanding this threat model is mandatory for correct operation.
