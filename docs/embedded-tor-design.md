# Embedded Tor Design

This document describes the design and purpose of the embedded Tor functionality in the Litsxg client.

Embedded Tor exists for operational convenience.
It does not alter the security model or guarantees of the tool.

---

## Purpose

The embedded Tor component allows the client to operate without requiring a pre-existing system Tor installation.

Its purpose is to:
- reduce initial setup friction
- improve portability
- allow operation in restricted environments

Embedded Tor is optional.

---

## Scope of embedded Tor

Embedded Tor is limited to:

- starting a local Tor process
- exposing a local SOCKS interface
- terminating with the client process

It does not expose configuration controls to the operator.

---

## Lifecycle

The embedded Tor lifecycle is bound to the client lifecycle.

- started when required
- used for the duration of client operation
- terminated when the client exits

Embedded Tor does not persist state across runs.

---

## Selection behavior

The client prefers system Tor.

Embedded Tor is used only when:
- system Tor is unavailable
- the client is configured to allow embedded Tor

The selection logic is fixed and not configurable at runtime.

---

## Configuration constraints

Embedded Tor operates with:

- default Tor behavior
- minimal configuration
- no operator customization

Advanced Tor configuration is out of scope.

---

## Security implications

Embedded Tor does not change:

- message visibility to the server
- authentication behavior
- protocol guarantees
- failure semantics

Operators must not assume improved security from embedded Tor usage.

---

## Performance considerations

Embedded Tor may have different performance characteristics than system Tor.

Examples include:
- increased startup latency
- variable circuit establishment time

The client does not adapt behavior based on performance.

---

## Failure behavior

If embedded Tor fails:

- connection attempts fail
- the client retries according to connection rules
- no diagnostic information is provided

Failure handling is identical to system Tor failure.

---

## Operator responsibility

Operators remain responsible for:

- understanding Tor behavior
- accepting Tor instability
- choosing whether embedded Tor is appropriate

Embedded Tor is a convenience, not an abstraction.

---

## Scope enforcement

Embedded Tor functionality is limited to what is described here.

No additional Tor features are exposed through the client.

---

## Final statement

Embedded Tor exists to reduce setup burden.

It does not change the fundamental properties of the Litsxg tool.
