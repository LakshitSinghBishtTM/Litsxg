# Security Claims

This document defines the complete and exclusive set of security claims made by Litsxg.

Only the claims listed here are made.
No implicit or derived claims exist.
Anything not listed here is not provided.

---

## Overview

Litsxg makes a small number of narrow security claims.

These claims are factual descriptions of behavior.
They are not guarantees of safety beyond their stated scope.

---

## Transport-level obscurity

Litsxg claims that:

- client-server communication is transported over Tor
- direct IP addresses are not exchanged between client and server

This provides resistance to casual network-level observation.

This does not provide anonymity against the server.

---

## Limited credential exposure

Litsxg claims that:

- plaintext passwords are not stored by the server
- stored credentials are represented as hashes

This limits exposure from accidental inspection of stored data.

This does not protect against determined attacks or credential theft.

---

## Identity continuity

Litsxg claims that:

- operators can maintain a stable identity across sessions
- identity collisions are prevented by authentication

This enables consistent message routing.

This does not imply strong authentication security.

---

## Explicit protocol behavior

Litsxg claims that:

- protocol behavior is deterministic
- protocol rules are explicitly documented
- no hidden negotiation or behavior exists

This allows operators to reason about system behavior.

---

## Minimal data retention

Litsxg claims that:

- message content is not durably stored
- server-side state is minimal and transient
- client-side state is ephemeral

This reduces long-term data accumulation.

This does not guarantee data absence at all times.

---

## Predictable failure behavior

Litsxg claims that:

- failure results in connection termination
- failure semantics are documented
- no silent recovery occurs

This makes failure visible through behavior.

---

## Scope of claims

All security claims are limited to:

- what is explicitly stated
- what is directly observable
- what is documented

Claims do not extend to third-party components or operator environments.

---

## Absence of implied claims

Litsxg explicitly does not claim:

- confidentiality of message content
- integrity protection
- authentication hardening
- anonymity against the server
- resistance to active attackers
- resistance to targeted attacks

Absence of a claim is intentional.

---

## Operator responsibility

Operators are responsible for:

- evaluating whether these claims are sufficient
- deploying the tool within this scope
- not inferring additional protections

Using Litsxg outside these claims is unsupported.

---

## Final statement

The security claims of Litsxg are intentionally narrow.

They describe exactly what the tool provides, and nothing more.
