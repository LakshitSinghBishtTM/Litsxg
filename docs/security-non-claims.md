# Security Non-Claims

This document defines what Litsxg explicitly does not claim from a security perspective.

These non-claims are intentional.
They are not gaps to be filled.
They will not change.

---

## Overview

Litsxg is explicit about what it does not provide.

Operators must treat the absence of a claim as a definitive statement.
No implied or indirect security properties exist.

---

## No message confidentiality

Litsxg does not claim:

- confidentiality of message content
- protection of messages from server visibility
- protection from message inspection

Messages are plaintext.
The server can read them.

---

## No message integrity

Litsxg does not claim:

- integrity protection
- tamper detection
- resistance to message modification

Messages may be altered by a malicious server or intermediary.
This is out of scope.

---

## No authentication strength

Litsxg does not claim:

- resistance to brute-force attacks
- resistance to credential theft
- strong identity verification

Authentication exists only to distinguish operators.

---

## No anonymity guarantees

Litsxg does not claim:

- anonymity against the server
- unlinkability of operator activity
- identity hiding beyond Tor transport

Operator identities are visible to the server.

---

## No metadata protection

Litsxg does not claim protection for:

- traffic patterns
- communication frequency
- message timing
- social graph inference

Metadata protection is out of scope.

---

## No resistance to active attackers

Litsxg does not claim resistance to:

- active network attackers
- protocol manipulation
- replay attacks
- injection attacks

Such threats are not addressed.

---

## No durability guarantees

Litsxg does not claim:

- message persistence
- delivery guarantees
- recovery after failure

Message loss is expected.

---

## No availability guarantees

Litsxg does not claim:

- uptime guarantees
- resistance to denial-of-service
- high availability

Availability depends entirely on operator infrastructure.

---

## No forward security

Litsxg does not claim:

- forward secrecy
- post-compromise security
- key rotation

These concepts do not apply within the tool’s scope.

---

## No safety assurances

Litsxg does not claim suitability for:

- high-risk communication
- adversarial environments
- sensitive operational use

Use in such contexts is unsupported.

---

## Operator responsibility

Operators are responsible for:

- understanding these non-claims
- not relying on absent protections
- selecting appropriate tools for their threat model

Misuse is not mitigated.

---

## Final statement

Security non-claims define the outer boundary of Litsxg.

Anything beyond this boundary is not provided.
