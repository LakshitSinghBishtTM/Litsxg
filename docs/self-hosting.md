# Self-Hosting

This document defines the self-hosting model for Litsxg.

Self-hosting is the preferred and expected mode of operation.
Managed hosting is not provided.

---

## Overview

Litsxg is designed to be self-hosted by the operator.

Self-hosting means:
- the operator controls the server
- the operator controls the infrastructure
- the operator accepts all operational responsibility

There is no external service dependency beyond Tor.

---

## Why self-hosting

Self-hosting exists to:

- avoid third-party trust dependencies
- maintain explicit control over data
- reduce operational ambiguity

Self-hosting aligns with the tool’s design constraints.

---

## Hosting environments

Litsxg may be self-hosted in environments such as:

- personal servers
- virtual private servers
- on-premises machines
- isolated network hosts

The environment choice does not alter tool behavior.

---

## Hardware requirements

Litsxg has minimal hardware requirements.

Typical requirements include:
- modest CPU capacity
- sufficient memory for in-memory buffers
- stable storage for identity data

There is no hardware acceleration.

---

## Operating system considerations

Litsxg does not depend on a specific operating system.

Operators are responsible for:
- system security
- process isolation
- resource management

OS hardening is external.

---

## Network configuration

When self-hosting with Tor:

- no public inbound ports are required
- the server should bind locally
- Tor handles external reachability

Network configuration errors result in unreachability.

---

## Data ownership

Self-hosting ensures:

- credential data remains under operator control
- no third-party storage is involved

Message content is transient and not retained.

---

## Maintenance responsibilities

The operator is responsible for:

- software updates
- Tor updates
- system updates
- data management

There is no automatic maintenance.

---

## Failure acceptance

Self-hosting requires acceptance of failure.

Operators must accept:
- message loss
- downtime
- manual recovery

No service-level guarantees exist.

---

## Backup considerations

Self-hosting allows backups, but they are optional.

If backups are used:
- only credential data is relevant
- backups must be secured
- backups may expose hashes

Message backups are not possible.

---

## Migration

Migrating a self-hosted instance requires:

- copying identity data if continuity is desired
- preserving onion service keys if address stability is required

Migration is manual.

---

## Scope enforcement

Self-hosting behavior outside this document is unsupported.

The tool does not adapt to hosting environments.

---

## Final statement

Self-hosting Litsxg is simple and direct.

Control is explicit, and responsibility is absolute.
