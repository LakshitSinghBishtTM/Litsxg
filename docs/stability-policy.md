# Stability Policy

This document defines what stability means in the context of Litsxg.

Stability refers to behavior, not comfort.
Stable does not mean forgiving.

---

## Overview

Litsxg is considered stable.

Stability means:
- behavior matches documentation
- interfaces do not change arbitrarily
- failures occur in documented ways

Stability does not imply safety or convenience.

---

## Behavioral stability

The following behaviors are stable and fixed:

- plaintext messaging
- protocol semantics
- authentication flow
- failure handling
- absence of guarantees

These behaviors will not change.

---

## Interface stability

The client interface is stable.

This means:
- no new UI elements will be added
- no guidance layers will be introduced
- no usability abstractions will appear

Interface behavior is final.

---

## Protocol stability

The protocol is stable.

This means:
- no new commands
- no version negotiation
- no backward compatibility layers

Protocol changes occur only to fix defects in documented behavior.

---

## Stability across upgrades

Upgrades may occur to fix bugs.

Operators must assume:
- client and server upgrades may need coordination
- mismatched versions may fail
- stability does not imply compatibility

Compatibility is not guaranteed.

---

## Failure stability

Failure behavior is stable.

Failures will continue to:
- terminate sessions
- lose state
- require manual recovery

Failure will not be softened.

---

## Data stability

Only identity data is stable across restarts.

All other data:
- is ephemeral
- may be lost
- should not be relied upon

This behavior is permanent.

---

## No deprecation policy

Litsxg does not deprecate features.

Features are either:
- present and final
- absent and permanent

There is no transition period.

---

## Operator expectations

Operators must expect:

- no feature evolution
- no gradual improvement
- no accommodation of new use cases

Stability means fixed constraints.

---

## Scope enforcement

If a stability property is not described here, it does not exist.

Operators must not infer additional assurances.

---

## Final statement

Litsxg is stable because it does not change.

Predictability comes from immobility, not polish.
