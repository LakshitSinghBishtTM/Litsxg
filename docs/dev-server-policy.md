# Dev Server Policy

This document defines the policy for development and convenience servers in Litsxg.

Development servers exist for convenience only.
They are not authoritative.
They are not reliable.
They may disappear without notice.

---

## Overview

A development server is any Litsxg server not intended for long-term operation.

Examples include:
- test servers
- demonstration servers
- temporary convenience servers

These servers are optional.

---

## Purpose of dev servers

Dev servers exist to:

- allow quick experimentation
- test connectivity and behavior
- provide short-lived access

They are not intended for sustained use.

---

## Data expectations

On dev servers:

- identity data may be wiped at any time
- offline messages may be discarded without notice
- behavior may change without warning

No data persistence is implied.

---

## Stability expectations

Dev servers provide no stability guarantees.

Operators must expect:
- frequent restarts
- inconsistent behavior
- protocol mismatches

Stability is not a goal.

---

## Security expectations

Dev servers provide no additional security.

Operators must assume:
- full message visibility
- no operational discipline
- no access control guarantees

Dev servers should not be trusted.

---

## Operator responsibility

Operators who run dev servers must:

- clearly label them as such
- avoid implying reliability
- avoid implying security guarantees

Misrepresentation is misuse.

---

## Client behavior

Clients do not distinguish between dev and production servers.

It is the operator’s responsibility to:
- communicate server intent
- manage expectations

The tool does not enforce server classification.

---

## Migration from dev to production

Dev servers should not be upgraded into production servers.

Recommended practice:
- deploy a fresh production server
- create a new onion service
- treat dev data as disposable

Identity continuity is optional.

---

## Scope enforcement

Dev server behavior is intentionally unrestricted.

There is no policy enforcement mechanism.

---

## Final statement

Dev servers exist to reduce friction, not to provide service.

They should never be treated as authoritative infrastructure.
