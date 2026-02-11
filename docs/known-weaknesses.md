# Known Weaknesses

This document enumerates known weaknesses in Litsxg.

These weaknesses are understood.
They are accepted.
They are not scheduled to be fixed.

---

## Overview

Litsxg prioritizes clarity over protection.

As a result, several weaknesses are inherent to its design.
Operators must be aware of them before deployment.

---

## Server visibility of messages

The server can read all message content.

This includes:
- live messages
- buffered offline messages
- multiline message content

This is a fundamental property of the tool.

---

## Server control over routing

The server controls message routing.

The server can:
- delay messages
- drop messages
- modify messages
- impersonate operators

There is no mitigation.

---

## Weak authentication model

Authentication is intentionally minimal.

Weaknesses include:
- plaintext password transmission
- lack of rate limiting
- lack of brute-force protection
- lack of credential recovery

Authentication does not protect against targeted attacks.

---

## Lack of message integrity

Messages have no integrity protection.

Weaknesses include:
- undetectable modification
- undetectable injection
- replay without detection

Message authenticity is not guaranteed.

---

## Lack of delivery guarantees

Messages may be lost.

Weaknesses include:
- loss during network failure
- loss during server restart
- loss during client failure

Loss is silent.

---

## Lack of persistence guarantees

Offline message buffering is volatile.

Weaknesses include:
- loss on server crash
- loss on server restart
- loss due to memory pressure

There is no durability.

---

## Tor instability dependence

Litsxg depends on Tor.

Weaknesses include:
- unpredictable latency
- circuit failure
- connection churn

The tool does not compensate for these conditions.

---

## No protection against traffic analysis

Litsxg does not protect against:

- timing analysis
- volume analysis
- communication pattern inference

Tor transport alone is insufficient.

---

## No resistance to denial-of-service

The server has no DoS protections.

Weaknesses include:
- unbounded connection attempts
- resource exhaustion
- forced restart

Availability depends entirely on operator environment.

---

## No client-side safeguards

The client does not protect operators from misuse.

Weaknesses include:
- no warnings
- no confirmations
- no sanity checks

Operator mistakes propagate directly.

---

## No audit or monitoring guarantees

The tool does not provide:

- audit logs
- tamper evidence
- monitoring hooks

Visibility is external.

---

## Weakness permanence

These weaknesses are not temporary.

They are design consequences.
They will remain.

---

## Operator responsibility

Operators are responsible for:

- understanding these weaknesses
- deciding whether they are acceptable
- deploying accordingly

Using Litsxg while ignoring these weaknesses is unsupported.

---

## Final statement

Known weaknesses define the limits of Litsxg.

Ignoring them does not make the tool stronger.
