# Port and Binding Model

This document defines how network ports and bindings are used by Litsxg.

Port usage is fixed.
Binding behavior is explicit.
There is no dynamic port negotiation.

---

## Overview

Litsxg uses a simple port and binding model.

- the server listens on a fixed local port
- Tor maps an onion service port to the local server port
- clients connect only via Tor

There is no direct clearnet exposure by default.

---

## Server binding

The server binds to a local network interface.

Properties:
- binding is typically to 127.0.0.1
- binding does not expose the server publicly
- binding is static and configured at startup

The server does not rebind during runtime.

---

## Server listening port

The server listens on a single TCP port.

Properties:
- one port per server instance
- no port multiplexing
- no dynamic port allocation

The port number is fixed by configuration.

---

## Onion service port mapping

Tor maps an onion service port to the server’s local listening port.

Properties:
- onion service port and local port may differ
- mapping is static
- mapping is defined in Tor configuration

Clients are unaware of the local port.

---

## Client connection target

Clients connect to:

- the onion service address
- the onion service port

Clients never connect directly to the server’s local interface.

---

## Embedded Tor binding

When embedded Tor is used:

- a local SOCKS interface is bound to a local address
- a local ephemeral port may be selected
- the client connects through this interface

Embedded Tor binding is internal to the client.

---

## Port conflicts

Port conflicts may occur if:

- the configured server port is already in use
- the embedded Tor SOCKS port cannot be allocated

On conflict:
- startup fails
- the affected component does not retry alternative ports
- operator intervention is required

---

## No dynamic rebinding

Litsxg does not support:

- automatic port reassignment
- fallback port selection
- runtime rebinding

All bindings are static for the lifetime of the process.

---

## Firewall considerations

When using onion services:

- no inbound firewall rules are required for clearnet access
- only local bindings are used

Firewall configuration is the operator’s responsibility.

---

## Scope enforcement

If a port or binding behavior is not described here, it does not exist.

Operators must not assume automatic discovery or negotiation.

---

## Final statement

The port and binding model in Litsxg is intentionally static.

Predictability is preferred over flexibility.
