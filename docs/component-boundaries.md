# Component Boundaries

This document defines the boundaries between components in the Litsxg system.

Boundaries are intentional.
Responsibilities do not overlap unless explicitly stated.
Behavior outside these boundaries is unsupported.

---

## Boundary principles

Litsxg enforces strict separation between components.

Each component:
- has a defined role
- owns specific state
- exposes limited interfaces

No component compensates for failures or omissions of another.

---

## Client responsibilities

The client is responsible for:

- presenting the graphical interface
- accepting operator input
- rendering received messages
- managing local interaction state
- initiating connections to the server
- following protocol rules

The client is not responsible for:
- message durability
- authentication policy enforcement
- access control beyond protocol compliance
- recovery from server-side failure

---

## Server responsibilities

The server is responsible for:

- accepting incoming connections
- authenticating operators
- maintaining operator identity mappings
- routing messages between operators
- temporarily storing offline messages
- enforcing protocol rules

The server is not responsible for:
- client-side interface behavior
- protecting message confidentiality
- masking network instability
- guaranteeing message delivery

---

## Tor responsibilities

Tor is responsible for:

- providing a transport path between client and server
- abstracting direct IP connectivity
- enabling onion service access

Tor is not responsible for:
- message integrity guarantees
- operator anonymity against the server
- application-level reliability
- protocol correctness

---

## State ownership

State ownership is strictly divided.

- Client-owned state:
  - input mode
  - connection status
  - active session state

- Server-owned state:
  - authenticated identities
  - active connections
  - offline message buffers

No state is jointly owned.

---

## Failure boundaries

Failures do not propagate responsibility across components.

Examples:
- Network failure does not shift responsibility to the client or server logic.
- Client misuse does not alter server guarantees.
- Server restart does not imply client-side recovery logic.

Each component fails independently.

---

## Interface contracts

Interfaces between components are defined by:

- the documented protocol
- documented connection behavior
- documented failure semantics

Undocumented behavior must not be relied upon.

---

## Boundary enforcement

When behavior crosses a component boundary unexpectedly, the result is undefined.

Such behavior is treated as operator error.

---

## Final statement

Component boundaries in Litsxg are strict.

Clarity is maintained by refusing to blur responsibility between components.
