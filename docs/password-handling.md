# Password Handling

This document defines how passwords are handled in Litsxg.

Password handling is intentionally minimal.
It exists only to support identity continuity.
It is not designed to resist determined attacks.

---

## Overview

Litsxg uses passwords to associate an operator identity with a connection.

Passwords are:
- provided by the operator
- transmitted in plaintext
- stored only in hashed form

No additional password protections are implemented.

---

## Password transmission

Passwords are transmitted:

- as plaintext
- within the authentication command
- over a Tor transport

There is no additional encryption layer at the protocol level.

---

## Password storage

The server stores:

- a hash of the password
- associated with an operator callsign

Plaintext passwords are not stored.

The storage format is fixed and minimal.

---

## Hashing purpose

Password hashing exists to:

- prevent casual inspection of stored credentials
- avoid storing plaintext passwords

Hashing does not exist to provide strong security guarantees.

---

## Absence of salting and hardening guarantees

Litsxg does not guarantee:

- resistance to offline brute-force attacks
- use of memory-hard password hashing
- resistance to credential reuse attacks

Password hardening is out of scope.

---

## Password creation and strength

Litsxg does not enforce:

- minimum password length
- complexity requirements
- password rotation

Operators are responsible for choosing passwords.

---

## Password recovery

Litsxg does not support:

- password recovery
- password reset
- account recovery

Loss of a password results in loss of identity access.

---

## Password reuse

Litsxg does not detect or prevent:

- password reuse
- shared passwords across identities

Password hygiene is the operator’s responsibility.

---

## Password lifecycle

Password lifecycle is simple:

- created implicitly on first authentication
- used for identity verification
- stored until explicitly removed or server data is deleted

There is no expiration mechanism.

---

## Password exposure risk

Operators must assume that:

- server compromise exposes password hashes
- plaintext passwords are visible during authentication
- embedded Tor does not change password visibility

These risks are accepted.

---

## Scope enforcement

If a password behavior is not documented here, it does not exist.

Operators must not assume additional protections.

---

## Final statement

Password handling in Litsxg is intentionally basic.

It supports identity continuity and nothing more.
