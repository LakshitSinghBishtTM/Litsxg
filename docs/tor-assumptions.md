# Tor Assumptions

This document defines the assumptions Litsxg makes about Tor.

These assumptions are narrow.
They are required for correct operation.
They do not imply stronger guarantees.

---

## Overview

Litsxg assumes Tor is available and functioning.

Tor is treated as an external dependency.
The tool does not attempt to validate or secure Tor itself.

---

## Availability assumptions

Litsxg assumes that:

- Tor may be temporarily unavailable
- Tor circuits may fail without warning
- connection attempts may take time

The tool tolerates these conditions by retrying connections.
It does not guarantee successful connection.

---

## Transport assumptions

Litsxg assumes Tor provides:

- a reliable byte stream when connected
- in-order delivery while the connection exists

If these assumptions are violated, protocol behavior is undefined.

---

## Identity assumptions

Litsxg assumes that:

- Tor does not conceal operator identity from the server
- onion service addressing does not imply anonymity

Identity is managed at the application layer.

---

## Security assumptions

Litsxg assumes that:

- Tor provides resistance to casual network observation
- Tor does not provide application-level security

Tor is not assumed to protect message content.

---

## Performance assumptions

Litsxg assumes that:

- latency may be high
- bandwidth may be limited
- performance may fluctuate

The tool does not adapt behavior based on performance.

---

## Failure assumptions

Litsxg assumes that:

- Tor failures are normal
- circuit loss can occur mid-session
- reconnection may require re-authentication

There is no session resumption.

---

## Configuration assumptions

Litsxg assumes that:

- Tor is correctly installed or embedded
- onion services are configured correctly
- local SOCKS interfaces are reachable

Misconfiguration results in failure.

---

## Operator responsibility

Operators are responsible for:

- maintaining Tor availability
- configuring onion services
- understanding Tor limitations

Tor-related issues are outside tool scope.

---

## Scope enforcement

If an assumption is not listed here, it must not be assumed.

Additional Tor properties are out of scope.

---

## Final statement

Litsxg treats Tor as infrastructure.

It relies on Tor where necessary and nowhere else.
