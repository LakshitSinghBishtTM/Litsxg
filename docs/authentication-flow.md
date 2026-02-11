# Authentication Flow

This document defines the authentication flow used by the Litsxg protocol.

Authentication exists solely to establish identity continuity.
It is not a security boundary.
It does not provide strong protection.

---

## Overview

Authentication is required before message routing is permitted.

The authentication flow:
- is synchronous
- is session-scoped
- completes entirely before normal operation

Authentication state does not persist across connections.

---

## Authentication goals

The authentication mechanism exists to:

- associate a callsign with a connection
- prevent accidental identity collisions
- enable message routing by identity

It does not exist to provide confidentiality, integrity, or resistance to attack.

---

## Initial authentication attempt

Upon connection, the client submits authentication credentials.

### Format

```
$ <callsign> <password>
```

This command must be the first meaningful input sent by the client.

---

## Existing identity authentication

If the callsign already exists:

- the server compares the provided password hash with stored data
- on match, authentication succeeds
- on mismatch, authentication fails

On failure:
- the connection remains unauthenticated
- message routing is not enabled
- the server does not provide detailed error feedback

---

## New identity registration

If the callsign does not exist:

- the server records the authentication attempt
- the connection remains unauthenticated
- the client must repeat the authentication command

On repeated submission with the same callsign:
- the server creates the new identity
- authentication succeeds
- identity continuity begins

This mechanism avoids accidental identity creation.

---

## Successful authentication

On successful authentication:

- the connection is associated with the callsign
- pending offline messages may be delivered
- normal message routing becomes available

Authentication state is bound to the active connection only.

---

## Authentication failure behavior

Authentication failure results in:

- refusal to associate identity
- continued unauthenticated session state
- possible silent ignore of subsequent messages

The protocol does not guarantee explicit failure messages.

---

## Session scope

Authentication state:

- exists only for the lifetime of the connection
- is discarded on disconnect
- must be re-established on reconnect

There is no session resumption.

---

## Identity uniqueness

Callsigns are unique within the server.

The server enforces uniqueness.
There is no namespace separation.

Identity ownership is determined solely by authentication success.

---

## Absence of authentication recovery

Litsxg does not support:

- password recovery
- identity recovery
- account migration

Loss of credentials results in loss of identity access.

---

## Absence of authentication hardening

The authentication mechanism does not include:

- rate limiting
- brute-force protection
- challenge-response protocols
- multi-factor authentication

Such features are out of scope.

---

## Scope enforcement

Authentication behavior outside this document is undefined.

Operators must not assume additional safeguards.

---

## Final statement

Authentication in Litsxg is intentionally minimal.

It exists to support routing, not to provide security guarantees.