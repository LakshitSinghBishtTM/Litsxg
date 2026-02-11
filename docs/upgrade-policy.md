# Upgrade Policy

This document defines the upgrade policy for Litsxg.

Upgrades are intentionally rare.
The design is finished.
No feature evolution is planned.

---

## Overview

Litsxg is considered complete.

The tool is production-grade in its current form.
There is no roadmap for new features, protocol changes, or security upgrades.

Upgrades exist only to correct defects.

---

## Design finality

The following aspects are final and will not change:

- plaintext protocol
- authentication model
- threat model
- Tor-based transport
- absence of end-to-end encryption
- absence of protocol extensibility

Any change to these would constitute a different tool.

---

## Permitted upgrade scope

The only permitted upgrades are:

- bug fixes
- crash fixes
- correctness fixes for documented behavior
- build or portability fixes

Upgrades will not introduce new guarantees.

---

## Non-permitted upgrades

The following will not be added:

- end-to-end encryption
- cryptographic redesigns
- protocol extensions
- feature additions
- usability layers
- automation or orchestration

Requests for these are out of scope.

---

## Compatibility expectations

Compatibility is narrow.

Operators must assume:
- upgrades may require client and server alignment
- backward compatibility is not guaranteed
- protocol versioning will not be introduced

Upgrade mismatches may result in failure.

---

## Upgrade procedure

To upgrade Litsxg:

1. Stop the server.
2. Replace the server binary.
3. Restart the server.
4. Restart clients as needed.

No state migration is performed.

---

## Data handling during upgrades

During upgrade:

- in-memory state is lost
- offline messages are lost
- persistent identity data remains intact

Operators must accept this loss.

---

## Rollback policy

Rollback is manual.

Operators must:
- retain old binaries if rollback is required
- accept possible data inconsistency
- restart processes manually

No automated rollback exists.

---

## Operator responsibility

Operators are responsible for:

- deciding when to upgrade
- coordinating client restarts
- accepting upgrade-related disruption

Upgrades are operational decisions.

---

## Scope enforcement

If an upgrade behavior is not described here, it does not exist.

There will be no silent evolution.

---

## Final statement

Litsxg is not an evolving platform.

Upgrades exist only to keep the tool functioning as designed.
