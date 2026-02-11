# State Machines

This document defines the explicit state machines used by the Litsxg tool.

State transitions are intentional.
Implicit transitions do not exist.
If a transition is not described here, it must not be assumed.

---

## Overview

Litsxg relies on simple, explicit state machines in both the client and the server.

State machines exist to:
- make behavior predictable
- avoid hidden transitions
- define failure boundaries clearly

There is no global system state.

---

## Client state machine

The client operates within a single linear state machine.

### States

1. Disconnected  
2. Connecting  
3. Connected  
4. Authenticated  
5. Multiline Input  

---

### Disconnected

The client is not connected to the server.

Properties:
- no active network connection
- no session state
- input is accepted but not transmitted

Transitions:
- to Connecting when a connection attempt begins

---

### Connecting

The client is attempting to establish a connection to the server.

Properties:
- Tor connectivity is being established
- no protocol messages are exchanged yet

Transitions:
- to Connected on successful connection
- to Disconnected on failure

Failure during this state is expected.

---

### Connected

The client has an active network connection.

Properties:
- protocol communication is possible
- identity is not yet established

Transitions:
- to Authenticated on successful authentication
- to Disconnected on connection loss

---

### Authenticated

The client has successfully authenticated.

Properties:
- identity is bound to the session
- messages may be sent and received

Transitions:
- to Multiline Input when multiline mode is entered
- to Disconnected on connection loss

---

### Multiline Input

The client is accepting multiline input.

Properties:
- input is buffered locally
- transmission follows multiline protocol rules

Transitions:
- to Authenticated when multiline input is terminated
- to Disconnected on connection loss

Multiline mode does not persist across sessions.

---

## Server state machine

The server operates with per-connection state machines.

Each client connection is independent.

---

### Unauthenticated

The server has accepted a connection.

Properties:
- no identity is associated
- only authentication-related protocol messages are processed

Transitions:
- to Authenticated on successful authentication
- to Disconnected on connection termination

---

### Authenticated

The server has associated an identity with the connection.

Properties:
- messages are accepted for routing
- offline message delivery may occur

Transitions:
- to Disconnected on connection termination

---

### Disconnected

The connection has terminated.

Properties:
- all per-connection state is discarded
- identity association is removed

This state is terminal.

---

## Failure transitions

Failure may cause immediate transition to Disconnected from any state.

Failure sources include:
- network interruption
- Tor circuit loss
- process termination
- protocol violation

No recovery transitions exist.

---

## Absence of recovery states

There are no recovery states.

Reconnection always begins from Disconnected.
Re-authentication is always required.

State is not resumed.

---

## Scope enforcement

State machines are intentionally minimal.

Additional states will not be added.
Transitions will not be inferred.

---

## Final statement

The Litsxg state machines are simple by design.

Predictable behavior is achieved by limiting state, not expanding it.
