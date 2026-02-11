# Non-Goals

This document defines what Litsxg explicitly does not attempt to do.

These are not limitations to be addressed later.
They are design constraints.
They will not change.

---

## No end-to-end encryption

Litsxg does not provide end-to-end encryption.

Messages are not cryptographically protected between operators.
The server can read message contents.

This is by design.

---

## No cryptographic message confidentiality

Litsxg does not attempt to provide cryptographic confidentiality for messages.

There is no encryption layer intended to protect message contents from the server or from a determined adversary with server access.

Operators must not assume otherwise.

---

## No protection from a malicious server

The server is trusted by design.

Litsxg does not attempt to defend operators against a malicious, compromised, or coerced server.

Any security model that assumes an untrusted server is out of scope.

---

## No anonymity against the server

Litsxg does not attempt to anonymize operators from the server.

Operator identities are known to the server.
Message routing depends on this knowledge.

Tor is not used to conceal identity from the server.

---

## No strong authentication guarantees

Authentication exists only to maintain identity continuity.

Password handling is intentionally minimal.
It is not designed to resist offline cracking or targeted attacks.

Authentication is not a security boundary.

---

## No delivery guarantees

Litsxg provides best-effort message delivery.

There are no acknowledgements.
There are no retries.
There are no guarantees of ordering, persistence, or eventual delivery.

Message loss is possible and expected under certain conditions.

---

## No recovery promises

Litsxg does not guarantee recovery from failure.

Crashes, restarts, or network loss may result in:
- message loss
- session termination
- state reset

The tool will not attempt to recover silently or compensate for misuse.

---

## No feature extensibility

Litsxg is feature-complete.

There is no plugin system.
There is no extension mechanism.
There is no supported way to add features without modifying the code.

---

## No roadmap

There is no roadmap.

There are no planned upgrades.
There are no planned security enhancements.
There are no planned protocol changes.

Maintenance is limited to bug fixes.

---

## No safety guarantees

Litsxg does not attempt to be safe for all use cases.

It is not intended for:
- high-risk communication
- situations requiring strong confidentiality
- environments requiring fault tolerance or recovery guarantees

Using Litsxg outside its defined scope is the responsibility of the operator.

---

## Final statement

Non-goals are not shortcomings.

They are the boundaries that define Litsxg.

Assuming behavior outside these boundaries is unsupported.
