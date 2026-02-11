# Resource Usage

This document defines the expected resource usage characteristics of Litsxg.

Resource usage is intentionally modest.
There is no dynamic scaling or optimization layer.
The tool assumes sufficient resources are provisioned by the operator.

---

## Overview

Litsxg is a lightweight tool.

It is designed to:
- run continuously
- consume predictable resources
- avoid background tasks and bloat

Resource usage depends primarily on connection count and message volume.

---

## CPU usage

CPU usage is generally low.

CPU is used for:
- handling network I/O
- routing messages
- basic authentication checks

There is no cryptographic workload beyond basic hashing.
There are no background computations.

CPU spikes may occur during:
- connection churn
- Tor circuit establishment
- bursts of message activity

---

## Memory usage

Memory usage is modest but unbounded.

Memory is used for:
- active connection state
- offline message buffers
- in-memory protocol handling

Offline message buffers are stored in memory.
There are no explicit limits.

Memory exhaustion is possible under misuse.

---

## Disk usage

Disk usage is minimal.

Disk is used for:
- credential database storage
- Tor onion service keys
- optional logs if enabled externally

Message content is not written to disk by design.

---

## Network usage

Network usage depends on:
- number of active connections
- message volume
- Tor circuit behavior

All network traffic flows through Tor.

Bandwidth usage is generally low but variable.

---

## Tor-related overhead

Tor introduces additional overhead.

This includes:
- increased latency
- circuit establishment costs
- background Tor maintenance traffic

These costs are external to Litsxg but unavoidable.

---

## File descriptor usage

Each active connection consumes:
- one network socket
- associated Tor resources

There is no connection pooling.

Operators must ensure sufficient file descriptor limits.

---

## Resource limits

Litsxg does not enforce:

- memory limits
- connection limits
- message size limits
- buffer limits

All limits are external.

---

## Failure under resource pressure

Under resource pressure:

- performance may degrade
- messages may be dropped
- the process may terminate

There is no graceful degradation mechanism.

---

## Monitoring resource usage

Litsxg does not provide resource metrics.

Operators must rely on:
- operating system tools
- process monitors
- Tor diagnostics

Resource visibility is external.

---

## Operator responsibility

Operators are responsible for:

- provisioning sufficient resources
- setting system-level limits
- monitoring usage
- mitigating exhaustion

Resource exhaustion is not mitigated by the tool.

---

## Scope enforcement

If a resource behavior is not described here, it does not exist.

Operators must not assume hidden safeguards.

---

## Final statement

Litsxg consumes only what it needs.

It does not protect itself from excess or misuse.
