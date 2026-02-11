# Audit Scope

This document defines what can and cannot be audited in Litsxg.

Auditability is limited.
Audits do not change trust assumptions.
This document exists to prevent misinterpretation.

---

## Overview

Litsxg does not position itself as an auditable security system.

Audits may verify behavior.
Audits do not add guarantees.

---

## What can be audited

The following aspects can be audited:

- source code correctness against documentation
- protocol implementation fidelity
- authentication flow implementation
- message routing logic
- failure handling behavior

Audits may confirm that the tool behaves as described.

---

## What cannot be audited meaningfully

The following cannot be meaningfully audited:

- confidentiality guarantees
- integrity guarantees
- anonymity properties
- resistance to active attackers
- resistance to targeted attacks

These properties are not provided.

---

## Server trust boundary

The server is trusted by design.

Audits cannot:
- remove server trust
- detect malicious server behavior at runtime
- prevent server abuse

Server trust is a fundamental assumption.

---

## Runtime behavior limits

Audits are limited to code and configuration.

They cannot guarantee:
- runtime environment integrity
- operator behavior
- absence of local compromise

Runtime trust is external.

---

## Tor-related auditing

Audits do not cover:

- Tor implementation correctness
- Tor network behavior
- Tor anonymity properties

Tor is an external dependency.

---

## Configuration auditing

Audits may verify:

- correct configuration usage
- absence of undocumented features

Audits do not validate operational decisions.

---

## Audit outcomes

Audit outcomes may include:

- confirmation of documented behavior
- identification of bugs
- identification of undefined behavior

Audit outcomes do not imply safety certification.

---

## No audit guarantees

Litsxg does not guarantee:

- auditability
- audit stability
- audit completeness

Audits are optional and non-authoritative.

---

## Operator responsibility

Operators are responsible for:

- deciding whether to audit
- interpreting audit results
- acting on findings

Audits do not replace responsibility.

---

## Scope enforcement

If an audit target is not described here, it is out of scope.

Audits must not infer absent guarantees.

---

## Final statement

Audits can confirm honesty, not provide safety.

Litsxg does not change under scrutiny.
