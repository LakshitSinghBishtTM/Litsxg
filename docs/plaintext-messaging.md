# Plaintext Messaging

This document defines the plaintext messaging model used by Litsxg.

Plaintext messaging is a deliberate design choice.
It is not a transitional state.
It will not change.

---

## Overview

All messages in Litsxg are plaintext.

Plaintext means:
- message content is readable by the server
- message content is readable in transit by any party with access
- no cryptographic transformation is applied to message bodies

This applies to all messages without exception.

---

## Rationale

Plaintext messaging exists to:

- keep the protocol simple
- make behavior transparent
- avoid false security assumptions
- reduce implementation complexity

Plaintext is treated as an explicit constraint, not a compromise.

---

## Message visibility

Plaintext messages are visible to:

- the server process
- the server operator
- any party with access to server memory or logs
- any party with access to the client runtime

Operators must assume full visibility at all times.

---

## Server handling of plaintext

The server:

- receives messages in plaintext
- processes routing decisions based on plaintext
- forwards messages in plaintext
- may buffer plaintext messages temporarily

The server does not attempt to obfuscate or transform content.

---

## Client handling of plaintext

The client:

- transmits plaintext messages
- displays plaintext messages
- stores plaintext messages only in memory

No client-side encryption or obfuscation exists.

---

## Multiline plaintext handling

Multiline messages:

- are concatenated into a single plaintext block
- preserve operator formatting
- are treated as opaque text by the server

Multiline messages do not alter plaintext properties.

---

## Logging implications

Plaintext messaging implies that:

- accidental logging may expose message content
- debugging output may reveal messages
- system-level logs may capture plaintext

Operators are responsible for managing logging exposure.

---

## Interaction with Tor

Tor does not change plaintext properties.

Tor:
- transports plaintext data
- does not encrypt message content end-to-end
- does not hide plaintext from the server

Transport security does not imply content confidentiality.

---

## Security implications

Plaintext messaging implies:

- no confidentiality guarantees
- no integrity guarantees
- no deniability
- no resistance to inspection

These implications are intentional and accepted.

---

## Operator responsibility

Operators are responsible for:

- deciding whether plaintext messaging is acceptable
- avoiding sensitive content where inappropriate
- selecting tools aligned with their threat model

Using plaintext where confidentiality is required is unsupported.

---

## Scope enforcement

Plaintext messaging is fundamental to Litsxg.

There is no configuration option to disable it.
There is no future plan to change it.

---

## Final statement

Plaintext messaging defines the core character of Litsxg.

It makes the tool explicit, predictable, and honest about its limits.
