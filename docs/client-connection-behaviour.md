# Client Connection Behaviour

This document defines how the Litsxg client establishes and maintains connections to the server.

Connection behaviour is deterministic.
There is no adaptive logic or operator prompting.
All retries follow fixed rules.

---

## Overview

The client is responsible for establishing a connection to the server over Tor.

Connection logic is simple:
- attempt connection
- on failure, retry
- on success, authenticate

There is no manual connection control.

---

## Initial connection attempt

On startup, the client immediately attempts to connect to the configured server address.

Properties:
- no confirmation prompt is shown
- no delay is introduced intentionally
- connection attempts begin automatically

The operator does not initiate the connection manually.

---

## Tor selection logic

The client follows a fixed Tor selection order:

1. Attempt connection using an existing system Tor instance.
2. If unavailable, attempt connection using an embedded Tor instance.
3. If both fail, retry according to retry rules.

This order does not change during runtime.

---

## System Tor usage

When system Tor is available:

- the client connects through the configured SOCKS interface
- Tor configuration is not modified
- Tor lifecycle is managed externally

Failure of system Tor triggers fallback behavior.

---

## Embedded Tor usage

If system Tor is unavailable:

- the client may start an embedded Tor instance
- the embedded Tor runs locally
- the client connects through the embedded SOCKS interface

Embedded Tor exists to reduce setup friction.
It does not change security properties.

---

## Retry behavior

On connection failure:

- the client waits a fixed interval
- the client retries connection
- retries continue indefinitely

There is no retry limit.
There is no exponential backoff.

---

## Successful connection

On successful connection:

- the client transitions to a connected state
- authentication begins immediately
- message exchange remains disabled until authentication completes

Connection success does not imply authentication success.

---

## Connection loss during operation

If the connection is lost during operation:

- the active session is terminated
- local connection state is cleared
- authentication state is discarded
- unsent input is lost

The client immediately resumes retry behavior.

---

## No operator intervention

The client does not:

- ask the operator to confirm retries
- provide connection diagnostics
- offer alternative connection paths

Connection management is automatic and opaque.

---

## Absence of connection indicators

The client does not provide explicit connection indicators.

Connection state is inferred through:
- message delivery
- input behavior
- session resets

There is no status display.

---

## Failure expectations

Connection failure is normal.

Operators must expect:
- intermittent connectivity
- repeated reconnects
- session loss

The client does not attempt to stabilize connections.

---

## Scope enforcement

Connection behavior outside this document is undefined.

Operators must not assume configurability or override mechanisms.

---

## Final statement

Client connection behaviour in Litsxg is fixed and predictable.

It exists to connect when possible and retry when not.