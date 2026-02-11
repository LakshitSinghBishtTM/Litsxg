# OPSEC

This document describes operational security considerations for Litsxg.

OPSEC is external.
Litsxg does not enforce it.
Failure to apply OPSEC is an operator decision.

---

## Overview

Litsxg assumes a competent operator.

The tool does not:
- protect against operator mistakes
- enforce safe defaults
- prevent unsafe deployments

Operational security exists outside the tool.

---

## Trust assumptions

Litsxg assumes:

- the operator trusts the server
- the operator controls the server
- the operator accepts plaintext visibility

Violating these assumptions breaks the model.

---

## Server-side OPSEC

Operators running a server must consider:

- who has system access
- who can read logs
- who can access memory
- who can access backups

Anyone with server access can read messages.

---

## Credential handling OPSEC

Operators must assume:

- password hashes may be exposed via backups
- plaintext passwords transit the server
- credential reuse increases risk

Password hygiene is external.

---

## Logging OPSEC

Operators must treat logs as sensitive.

Logs may contain:
- message content
- operator identifiers
- connection metadata

Logging must be controlled externally.

---

## Host system OPSEC

The host system is trusted.

Operators must secure:
- operating system access
- user accounts
- process permissions
- filesystem permissions

A compromised host compromises Litsxg completely.

---

## Network OPSEC

Tor hides direct IP exposure.

It does not:
- hide message content
- prevent timing correlation
- prevent traffic analysis

Network OPSEC beyond Tor is external.

---

## Client-side OPSEC

Client systems must be trusted.

Operators must assume:
- local compromise exposes messages
- screen content can be observed
- memory can be inspected

Client security is not enforced.

---

## Identity OPSEC

Callsigns are persistent identities.

Operators should consider:
- reuse across contexts
- correlation risk
- accidental disclosure

Identity management is manual.

---

## Physical OPSEC

Physical access overrides all software controls.

Operators must secure:
- machines
- storage media
- backup devices

Physical compromise is total compromise.

---

## What OPSEC will not do

OPSEC will not:

- make plaintext safe
- prevent server abuse
- correct misuse
- add security properties not designed into the tool

OPSEC cannot change the model.

---

## Operator responsibility

Operators are responsible for:

- deciding whether OPSEC is sufficient
- applying external controls
- accepting residual risk

Litsxg does not validate OPSEC.

---

## Scope enforcement

If an OPSEC measure is not documented here, it is external.

The tool provides no additional safeguards.

---

## Final statement

Litsxg assumes disciplined operators.

OPSEC failures are not tool failures.
