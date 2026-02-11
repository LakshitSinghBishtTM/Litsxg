# Why Tor

This document explains why Litsxg uses Tor.

Tor is used intentionally and narrowly.
Its inclusion does not imply broad privacy or anonymity guarantees.

---

## Purpose of Tor in Litsxg

Tor is used as a transport mechanism.

Its purpose in Litsxg is to:
- provide reachability without direct IP exposure
- simplify deployment through onion services
- allow operators to host servers without public network configuration

Tor is infrastructure, not a security claim.

---

## What Tor provides

In the context of Litsxg, Tor provides:

- a routable network path between client and server
- resistance to casual network-level observation
- stable addressing via onion services

These properties are sufficient for the intended scope of the tool.

---

## What Tor does not provide

Tor does not provide, within Litsxg:

- anonymity against the server
- message confidentiality
- protection from a malicious server
- application-level security guarantees
- delivery reliability

Operators must not assume Tor compensates for protocol or design limits.

---

## Why Tor over alternatives

Tor was chosen over alternatives because:

- onion services remove the need for public IPs
- deployment is simpler for self-hosting
- connectivity works across restrictive networks

The choice is practical, not ideological.

---

## Tor and identity visibility

Even when using Tor:

- the server knows operator identities
- authentication occurs in plaintext
- message contents are visible to the server

Tor does not hide identity from the server.

---

## Tor instability as a design factor

Tor connections may be unstable.

Litsxg treats this as normal.
The tool does not attempt to mask or correct Tor behavior.

Failure over Tor results in:
- connection loss
- session termination
- message loss

This is expected.

---

## Embedded Tor usage

The client may use an embedded Tor instance.

This exists to:
- reduce operator setup complexity
- improve portability

Embedded Tor does not change security properties.

---

## Operator responsibility

Operators are responsible for:

- understanding Tor behavior
- accepting Tor instability
- managing onion service configuration

Tor usage does not absolve operators of operational responsibility.

---

## Scope enforcement

If Tor behavior is not documented here, it must not be assumed.

Tor is not a substitute for application-level guarantees.

---

## Final statement

Litsxg uses Tor because it is useful.

It does not use Tor to make promises it cannot keep.
