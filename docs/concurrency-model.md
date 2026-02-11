# Concurrency Model

This document defines how concurrency is handled in Litsxg.

Concurrency is limited by design.
Parallelism exists only where required for correct operation.
There is no attempt to maximize throughput or utilization.

---

## Overview

Litsxg uses concurrency to separate:

- network I/O
- operator interaction
- internal state updates

Concurrency exists to prevent blocking, not to increase performance.

---

## Client concurrency model

The client uses a small number of concurrent execution paths.

### Responsibilities

Concurrency in the client is used to:

- read incoming network data
- process operator input
- update the interface without blocking

These activities operate independently.

---

### Network read path

Incoming data from the server is processed in a dedicated execution path.

Properties:
- blocking reads are isolated from the interface
- received messages are forwarded to the display layer
- failure results in connection termination

There is no buffering beyond what is required for protocol parsing.

---

### Network write path

Outgoing messages are written synchronously.

Properties:
- writes occur only when the operator submits input
- failed writes result in connection teardown
- no retry logic exists

Write failure is treated as terminal for the session.

---

### Interface updates

Interface updates occur independently of network I/O.

Properties:
- updates reflect current client state
- no attempt is made to batch or coalesce updates
- visual state mirrors internal state directly

Interface responsiveness is prioritized over ordering guarantees.

---

### Client state synchronization

Shared client state is limited and synchronized explicitly.

Properties:
- minimal shared state
- explicit locking where required
- no lock-free or speculative behavior

State corruption is avoided by limiting scope, not complexity.

---

## Server concurrency model

The server handles multiple connections concurrently.

Each connection is processed independently.

---

### Connection handling

Each inbound connection is assigned an independent execution context.

Properties:
- no shared execution path between connections
- per-connection failures do not affect others

Connection lifetimes are independent.

---

### Shared server state

The server maintains limited shared state:

- active operator mappings
- offline message buffers

Access to shared state is synchronized explicitly.

---

### Synchronization strategy

The server uses simple synchronization primitives.

Properties:
- coarse-grained locking
- no lock-free data structures
- no optimistic concurrency control

This favors correctness and clarity over throughput.

---

### Message routing concurrency

Message routing may involve:

- reading shared state
- writing to a target connection
- appending to an offline buffer

Routing operations are serialized only where required to maintain consistency.

---

## Absence of advanced concurrency features

Litsxg intentionally avoids:

- worker pools
- asynchronous task queues
- background schedulers
- speculative execution

Concurrency complexity is treated as a risk, not an optimization.

---

## Failure behavior under concurrency

Concurrency does not change failure semantics.

Under failure:
- connections terminate
- shared state is cleaned up
- no recovery attempts are made

Race conditions are prevented by limiting concurrent interactions.

---

## Scope enforcement

If a concurrent behavior is not documented here, it does not exist.

Operators must not assume implicit parallelism or performance guarantees.

---

## Final statement

The Litsxg concurrency model is intentionally conservative.

Correctness and predictability are preferred over performance or scale.
