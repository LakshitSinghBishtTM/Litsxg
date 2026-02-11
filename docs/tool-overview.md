# Tool Overview

This document provides a high-level operational overview of the Litsxg tool.

It describes what the tool does, how it is intended to be used, and what assumptions it makes.
It does not restate non-goals or security limits already defined elsewhere.

---

## Purpose

Litsxg exists to route plaintext messages between operators over Tor.

The tool is designed to:
- be predictable
- be stable
- expose its limits clearly
- avoid hidden behavior

It does not attempt to optimize for safety, convenience, or abstraction.

---

## Core components

Litsxg consists of two primary components:

- a client
- a server

Both components are required for operation.
Neither component is optional.

---

## Client role

The client provides the operator interface.

It is responsible for:
- accepting operator input
- displaying received messages
- managing connection state
- handling multiline message entry
- connecting to the server over Tor

The client does not attempt to enforce correctness beyond protocol rules.

---

## Server role

The server provides message routing.

It is responsible for:
- accepting incoming connections
- authenticating operators
- routing messages to recipients
- storing limited offline messages
- managing session state

The server is trusted to observe message contents and metadata.

---

## Message flow

At a high level, message flow is as follows:

1. The operator connects a client to the server over Tor.
2. The operator authenticates to establish identity.
3. Messages are sent as plaintext to the server.
4. The server routes messages to target operators.
5. Messages may be delivered immediately or stored temporarily.

This flow does not change.

---

## Failure behavior

Failure is expected.

Litsxg does not attempt to:
- mask network instability
- retry delivery silently
- preserve state beyond defined limits

Failure results in:
- message loss
- session termination
- operator-visible interruption

This behavior is intentional.

---

## Operational assumptions

Litsxg assumes that operators:

- read documentation
- understand Tor connectivity
- accept message loss
- operate their own infrastructure where required

The tool does not attempt to correct incorrect assumptions.

---

## Scope enforcement

If a behavior is not described in documentation, it must not be assumed.

Unexpected behavior is treated as operator error unless explicitly documented otherwise.

---

## Final statement

Litsxg is a narrowly scoped tool.

It does one thing: route plaintext messages over Tor.
It does not attempt to do more.
