# System Overview

This document provides a high-level view of the Litsxg system.

It describes the components involved, how they interact, and the assumptions that govern their interaction.
Detailed behavior is covered in component-specific documents.

---

## System composition

Litsxg consists of three primary elements:

- client
- Tor network
- server

All communication between client and server occurs over Tor.

---

## Client

The client is the operator-facing component.

Its responsibilities include:
- presenting a graphical interface
- collecting operator input
- displaying received messages
- managing connection and session state
- communicating with the server over Tor

The client maintains minimal local state.
It does not persist messages or session data across restarts.

---

## Server

The server is the routing and coordination component.

Its responsibilities include:
- accepting incoming connections
- authenticating operators
- routing messages to recipients
- temporarily storing offline messages
- maintaining in-memory session state

The server is trusted by design.
It has visibility into message contents and routing metadata.

---

## Tor network

Tor is used as the transport layer between client and server.

In Litsxg, Tor provides:
- reachability without direct IP exposure
- simplified deployment via onion services

Tor is not used to claim anonymity or strong privacy guarantees beyond transport-level properties.

---

## Communication model

The communication model is client-server.

- Clients initiate connections.
- The server does not initiate contact with clients.
- All messages pass through the server.

There is no peer-to-peer communication.

---

## Message flow overview

At a high level:

1. The client establishes a Tor connection to the server.
2. The operator authenticates to claim an identity.
3. Messages are sent as plaintext to the server.
4. The server routes messages to target operators.
5. Messages may be delivered immediately or queued temporarily.

This flow is fixed.

---

## State management

State is divided as follows:

- Client state:
  - current session
  - input mode
  - connection status

- Server state:
  - active sessions
  - operator identities
  - temporary offline messages

State persistence is intentionally limited.

---

## Failure model

Failure is expected and treated as normal.

Examples include:
- Tor circuit failure
- network interruption
- server restart
- client termination

The system does not attempt automatic recovery beyond reconnecting when possible.

---

## Scope boundaries

Behavior outside this overview is undefined unless documented elsewhere.

Operators must not assume additional layers, services, or safeguards.

---

## Final statement

The Litsxg system is intentionally simple.

Its architecture favors clarity and predictability over abstraction or resilience.
