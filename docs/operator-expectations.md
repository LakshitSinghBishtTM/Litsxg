# Operator Expectations

This document defines what is expected of operators using Litsxg.

These expectations are mandatory.
Failure to meet them results in unsupported operation.

---

## Responsibility

Operators are fully responsible for:

- understanding documented behavior
- deploying and operating infrastructure
- managing their own risk
- interpreting failures correctly

Litsxg does not assume responsibility for misuse or incorrect assumptions.

---

## Documentation literacy

Operators are expected to:

- read relevant documentation before operation
- understand non-goals and limits
- follow defined operating procedures

Skipping documentation does not change tool behavior.

---

## Security assumptions

Operators must understand and accept that:

- messages are plaintext
- the server can read message contents
- authentication does not provide strong security
- no cryptographic confidentiality is provided

Assuming additional security properties is operator error.

---

## Failure acceptance

Operators must accept that:

- network failures occur
- message loss is possible
- sessions may terminate without recovery
- offline messages are not guaranteed

Litsxg does not attempt to compensate for failure.

---

## Infrastructure competence

Operators are expected to:

- understand Tor connectivity
- manage onion services where applicable
- handle server restarts and updates
- maintain system availability if required

The tool does not abstract infrastructure complexity.

---

## Self-hosting preference

Operators are expected to self-host where possible.

Shared or public servers are provided for convenience only.
Dependence on shared infrastructure is discouraged.

---

## No entitlement to support

Operators must not expect:

- guaranteed support
- feature requests to be accepted
- behavior outside documentation

Support policy is defined separately.

---

## Interface expectations

Operators must not expect:

- guidance from the interface
- warnings for misuse
- safeguards against incorrect input

The interface reflects tool state, not operator intent.

---

## Compliance with scope

Operators must operate within defined scope.

Use cases outside scope are unsupported and unaddressed.

---

## Final statement

Litsxg assumes capable operators.

The tool does not protect operators from their own assumptions.
