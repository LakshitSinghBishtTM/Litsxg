# Project Scope

This document defines the scope of the Litsxg tool.

Scope is fixed.
Anything not explicitly included here is out of scope by default.

---

## In-scope functionality

Litsxg provides the following functionality:

- Routing plaintext messages between authenticated operators
- Operating over Tor for network reachability
- Maintaining identity continuity via simple authentication
- Supporting single-line and multiline message input
- Providing a minimal graphical interface for interaction
- Allowing offline message buffering on the server
- Supporting self-hosted deployment

These functions define the complete operational scope of the tool.

---

## Out-of-scope functionality

The following are explicitly out of scope:

- End-to-end encryption
- Cryptographic message confidentiality
- Metadata minimization beyond Tor transport
- Protection from a malicious or compromised server
- Strong authentication or password security
- Message durability guarantees
- Delivery acknowledgements or retries
- Multi-device synchronization
- Account recovery mechanisms
- Automatic updates
- Extension or plugin systems

Out-of-scope items will not be added.

---

## Scope stability

The scope of Litsxg is stable.

No new functionality is planned.
No expansion of scope is planned.
The scope will not evolve.

Changes are limited to correcting defects that violate documented behavior.

---

## Deployment scope

Litsxg is designed to be deployed by operators.

Self-hosting is preferred.
Shared or public servers may exist for convenience, but are not part of the core operating model.

Operating Litsxg in hosted or managed environments does not change its scope.

---

## Documentation scope

Documentation describes:

- What the tool does
- What the tool does not do
- How the tool behaves under normal and failure conditions

Documentation does not describe hypothetical features or future behavior.

---

## Scope enforcement

Behavior outside documented scope is unsupported.

Expectations beyond scope are operator errors.

---

## Final statement

The scope of Litsxg is narrow by design.

Stability is achieved by refusing to expand it.
