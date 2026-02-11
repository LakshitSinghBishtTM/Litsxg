# Operator Manual

This document describes how to operate a Litsxg server.

This manual assumes the operator controls the server.
There is no distinction between administrator and operator roles.

---

## Overview

Litsxg is operated as a single long-running server process.

Operation consists of:
- starting the server
- monitoring basic health
- restarting on failure
- managing identity data

There is no interactive control plane.

---

## Operator role

The operator is responsible for:

- server deployment
- availability
- data retention decisions
- Tor configuration
- accepting all design constraints

The tool does not enforce operational discipline.

---

## Server startup

To start the server:

- ensure required dependencies are available
- launch the server binary
- verify that it is listening on the configured port

Startup is synchronous and immediate.

---

## Runtime operation

During runtime, the server:

- accepts incoming connections
- authenticates operators
- routes messages
- buffers offline messages in memory

No manual intervention is required during normal operation.

---

## Monitoring

Litsxg provides no built-in monitoring.

Operators may monitor:
- process liveness
- resource usage
- Tor connectivity

All monitoring is external.

---

## Logging

Logging behavior is minimal.

The server does not provide:
- structured logs
- audit trails
- access logs

Any logging present is implementation-dependent.

Operators must assume logs may expose plaintext.

---

## Failure handling

On failure:

- the server process terminates
- all connections are dropped
- in-memory state is lost

The operator must restart the server manually or externally.

---

## Restart procedure

To restart the server:

1. Stop the server process.
2. Start the server process again.
3. Verify connectivity via Tor.

No state migration or cleanup is required.

---

## Identity management

Operator identities are stored persistently.

The operator may:
- delete identity data manually
- reset the database
- discard identity continuity entirely

There is no identity management interface.

---

## Data management

Persistent data is limited.

Operators are responsible for:
- securing the credential database
- deciding whether to back it up
- deleting it when appropriate

Message content is not persistently stored.

---

## Upgrade handling

Upgrades are manual.

Procedure:
- stop the server
- replace the binary
- restart the server

No compatibility guarantees exist beyond bug fixes.

---

## Multi-instance operation

Running multiple instances requires:

- separate ports
- separate Tor onion services
- separate data directories

Instances do not coordinate.

---

## Operator expectations

Operators must assume:

- message loss is possible
- failure is normal
- recovery is manual
- no safety nets exist

Operational discipline is externalized.

---

## Scope enforcement

This manual describes all supported operational behavior.

Undocumented operations are unsupported.

---

## Final statement

Operating Litsxg is straightforward but unforgiving.

The tool does exactly what is documented, and nothing more.
