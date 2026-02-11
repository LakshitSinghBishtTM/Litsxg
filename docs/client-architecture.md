# Client Architecture

This document describes the internal structure and behavior of the Litsxg client.

It focuses on responsibility boundaries, state handling, and interaction with the server.
Implementation details are described only where they affect observable behavior.

---

## Client role

The client is an operator-facing component.

Its role is limited to:
- providing a graphical interface
- accepting operator input
- displaying received messages
- managing local interaction state
- communicating with the server over Tor

The client does not attempt to provide safety guarantees or hide system behavior.

---

## Architectural overview

The client is organized around the following functional areas:

- interface layer
- connection management
- protocol handling
- local state tracking

Each area has a defined purpose and minimal overlap with others.

---

## Interface layer

The interface layer is responsible for:

- rendering output text
- collecting operator input
- displaying connection state implicitly through behavior
- handling multiline input mode

The interface:
- does not validate semantic correctness
- does not guide the operator
- does not warn about misuse

Visual output reflects server responses and local input only.

---

## Connection management

Connection management is responsible for:

- establishing connections to the server over Tor
- handling reconnect attempts
- tracking connection availability

Connection logic:
- prefers existing system Tor where available
- falls back to embedded Tor where configured
- retries indefinitely with fixed behavior

Connection failure is treated as normal.

---

## Protocol handling

Protocol handling is responsible for:

- formatting outbound messages
- parsing inbound messages
- enforcing message framing rules
- managing multiline message delimiters

Protocol handling follows documented protocol definitions strictly.
Undocumented protocol behavior is ignored.

---

## Local state

The client maintains minimal local state, including:

- current connection status
- authentication state
- input mode (single-line or multiline)
- in-memory message history for display

Local state is not persisted across restarts.

---

## Multiline input mode

The client supports a multiline input mode.

This mode:
- is explicitly entered by the operator
- remains active until explicitly terminated
- sends content according to protocol rules

The client does not infer operator intent.

---

## Concurrency model

The client uses concurrent routines to:

- read from the network
- write to the interface
- process operator input

Concurrency is limited and controlled.
Shared state access is synchronized explicitly.

---

## Error handling

The client handles errors by:

- terminating the affected connection
- resetting local connection state
- continuing operation where possible

Errors are not surfaced as warnings or dialogs.

---

## Persistence policy

The client does not persist:

- messages
- credentials
- session state

All client-side data is ephemeral.

---

## Scope enforcement

The client enforces scope by omission.

If a feature is not implemented, it is unavailable.
No compatibility layers or fallbacks are provided.

---

## Final statement

The Litsxg client is intentionally simple.

Its architecture exists to expose system behavior directly, not to abstract it away.
