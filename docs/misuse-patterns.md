# Misuse Patterns

This document describes common misuse patterns for Litsxg.

These patterns are not edge cases.
They are predictable outcomes of incorrect assumptions.
The tool does not attempt to prevent them.

---

## Overview

Misuse occurs when Litsxg is operated outside its defined scope.

Most misuse patterns result from assuming properties that the tool does not provide.
Documentation exists to prevent this, not to mitigate consequences.

---

## Treating Litsxg as a secure messenger

Using Litsxg as if it provides strong security guarantees is misuse.

Examples include assuming:
- message confidentiality
- integrity protection
- server untrustworthiness
- anonymity beyond Tor transport

Litsxg does not support these assumptions.

---

## Using Litsxg for high-risk communication

Using Litsxg for high-risk or sensitive communication is misuse.

Examples include:
- activism under threat
- whistleblowing
- coordination under adversarial surveillance
- legal or safety-critical communication

The tool is not designed for these contexts.

---

## Relying on message delivery guarantees

Assuming messages will be delivered reliably is misuse.

Examples include:
- assuming offline messages will persist
- assuming messages will arrive in order
- assuming retries will occur

Message loss is expected.

---

## Assuming recovery after failure

Assuming the tool will recover state after failure is misuse.

Examples include:
- expecting session resumption
- expecting message replay
- expecting buffered input preservation

All state is ephemeral.

---

## Trusting the server incorrectly

Assuming the server is neutral, honest, or benevolent is misuse.

The server:
- can read messages
- can modify messages
- can impersonate operators

Trust is required, not enforced.

---

## Treating Tor as a security boundary

Assuming Tor provides application-level security is misuse.

Tor:
- provides transport reachability
- does not protect message content
- does not prevent server visibility

Tor does not compensate for protocol limits.

---

## Ignoring documentation boundaries

Skipping documentation and inferring behavior is misuse.

Examples include:
- assuming undocumented features
- assuming implied guarantees
- assuming “reasonable behavior”

Only documented behavior exists.

---

## Depending on public or dev servers

Treating public or development servers as reliable infrastructure is misuse.

Such servers:
- may disappear
- may change behavior
- may lose data

Self-hosting is preferred.

---

## Automating against the protocol

Building automation that relies on undocumented protocol behavior is misuse.

The protocol:
- is not designed for automation
- has no stability guarantees for such use
- may change under bugfixes

Automation is unsupported.

---

## Treating Litsxg as a platform

Attempting to extend Litsxg into a platform is misuse.

Examples include:
- adding plugins
- layering services
- building dependency chains

The tool is closed by design.

---

## Operator responsibility

Operators are responsible for avoiding misuse.

The tool does not:
- warn
- block
- compensate

Misuse results in undefined behavior.

---

## Final statement

Misuse patterns are predictable.

Litsxg does not attempt to protect operators from them.
