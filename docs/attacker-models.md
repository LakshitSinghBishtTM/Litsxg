# Attacker Models

This document enumerates attacker models relevant to Litsxg.

These models describe possible attackers and their capabilities.
They do not imply defenses.
Only explicitly stated mitigations exist.

---

## Overview

Attacker models are used to clarify assumptions.

Litsxg does not attempt to be robust against most attackers.
Understanding these models is required to avoid misuse.

---

## Passive network observer

### Description

A passive network observer can:
- observe network traffic
- record traffic metadata
- analyze traffic patterns

Examples include:
- local network observers
- passive ISP-level observers

### Impact

- plaintext messages may be observed in transit without Tor
- metadata may be collected

### Mitigation

- Tor transport provides resistance to casual observation

No further mitigation exists.

---

## Active network attacker

### Description

An active network attacker can:
- inject traffic
- modify traffic
- drop traffic
- perform man-in-the-middle attacks

### Impact

- message tampering
- message loss
- session termination

### Mitigation

None.

Active network attackers are out of scope.

---

## Malicious server operator

### Description

A malicious server operator controls the server.

Capabilities include:
- reading all messages
- modifying messages
- dropping messages
- impersonating operators
- fabricating messages

### Impact

- total compromise of confidentiality and integrity
- complete loss of trust

### Mitigation

None.

The server is trusted by design.

---

## Compromised server

### Description

A compromised server is under third-party control.

Capabilities include:
- memory inspection
- database access
- runtime modification

### Impact

- exposure of credentials and messages
- undefined behavior

### Mitigation

None.

Server compromise is catastrophic.

---

## Malicious client operator

### Description

A malicious client operator controls their own client.

Capabilities include:
- protocol misuse
- message flooding
- malformed input

### Impact

- message spam
- resource exhaustion
- undefined protocol behavior

### Mitigation

Minimal.

Protocol violations may result in disconnect.

---

## Curious administrator

### Description

A curious administrator has legitimate access to server systems.

Capabilities include:
- inspecting logs
- reading memory
- accessing stored data

### Impact

- message disclosure
- credential exposure

### Mitigation

Limited.

Password hashing prevents casual plaintext inspection only.

---

## Local system attacker

### Description

A local system attacker controls the client or server host.

Capabilities include:
- keylogging
- memory inspection
- process manipulation

### Impact

- total compromise of local component

### Mitigation

None.

Local system security is assumed.

---

## Automated attacker

### Description

An automated attacker can:
- script protocol interactions
- brute-force credentials
- flood connections

### Impact

- denial of service
- credential compromise

### Mitigation

None.

Rate limiting and hardening are out of scope.

---

## Legal or coercive attacker

### Description

A legal or coercive attacker can:
- compel disclosure
- seize infrastructure
- enforce compliance

### Impact

- forced access to data
- forced behavior changes

### Mitigation

None.

Litsxg does not attempt resistance.

---

## Operator responsibility

Operators are responsible for:

- understanding attacker capabilities
- determining acceptability
- choosing appropriate tools

Deploying Litsxg against unsupported attackers is misuse.

---

## Final statement

Attacker models define who Litsxg does not protect against.

Assuming protection where none exists is incorrect.
