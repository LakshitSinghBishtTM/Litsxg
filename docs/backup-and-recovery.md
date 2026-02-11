# Backup and Recovery

This document defines what backup and recovery mean in the context of Litsxg.

Backup is limited.
Recovery is limited.
Neither is comprehensive or guaranteed.

---

## Overview

Litsxg stores very little persistent data.

As a result:
- backup scope is narrow
- recovery scope is narrow
- most state is unrecoverable by design

This is intentional.

---

## What can be backed up

The only meaningful persistent data is:

- operator credential database
- onion service keys (if address stability is required)

No other data is durable.

---

## What cannot be backed up

The following cannot be backed up:

- message content
- offline message buffers
- session state
- connection state
- multiline input buffers

Loss of this data is permanent.

---

## Credential database backup

If identity continuity matters, the operator may back up:

- the credential database file

Considerations:
- backups contain password hashes
- backups must be protected accordingly
- restoring a backup restores identities exactly as stored

There is no partial restore.

---

## Onion service key backup

If onion address stability matters, the operator must back up:

- the onion service directory managed by Tor

Loss of these keys results in:
- loss of the onion address
- creation of a new identity for the server

This is irreversible.

---

## Backup frequency

Backup frequency is an operator decision.

Litsxg provides:
- no scheduling
- no reminders
- no automation

Backup policy is external.

---

## Recovery procedure

Recovery consists of:

1. Restoring backed-up files.
2. Restarting Tor if applicable.
3. Restarting the Litsxg server.

No validation or reconciliation occurs.

---

## Recovery limitations

Recovery does not restore:

- active sessions
- pending messages
- client state

Operators must accept that recovery is incomplete.

---

## Crash interaction

Backups do not improve crash behavior.

After a crash:
- in-memory data is lost
- only backed-up persistent data may be restored

Crash recovery remains restart-only.

---

## Backup risks

Backups introduce risk.

Examples include:
- exposure of password hashes
- exposure of onion service keys
- unintended duplication of identity data

Operators must manage these risks.

---

## Scope enforcement

If a backup or recovery behavior is not described here, it does not exist.

There are no hidden safeguards.

---

## Final statement

Backup and recovery in Litsxg are minimal by design.

Only identity continuity can be preserved.
Everything else is transient.
