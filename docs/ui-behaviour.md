# UI Behaviour

This document defines the behavior of the Litsxg client interface.

The interface is intentionally minimal.
There are no prompts, banners, or guidance layers.
Behavior is direct and literal.

---

## Overview

The Litsxg client provides a single interactive interface.

Its purpose is to:
- accept operator input
- display received messages
- reflect internal state implicitly

The interface does not explain itself.

---

## Input handling

Operator input is accepted as plaintext.

Properties:
- input is sent exactly as entered
- no preprocessing or sanitization occurs
- leading and trailing whitespace is preserved

Input semantics are defined by the protocol, not the interface.

---

## Output handling

Received messages are displayed verbatim.

Properties:
- messages appear in the order received by the client
- no formatting is applied beyond prefixing sender identity
- multiline messages are displayed as contiguous blocks

The interface does not annotate messages.

---

## Prompt semantics

The interface uses no explicit prompts.

There is:
- no mode indicator
- no status indicator
- no connection indicator

Operator context must be inferred.

---

## Multiline mode behavior

Multiline mode is implicit.

Behavior:
- entering multiline mode changes how input is interpreted
- no visual distinction is provided
- termination must be explicit

Failure to terminate multiline mode results in continued buffering.

---

## Connection state visibility

Connection state is not explicitly displayed.

Indicators of connection state include:
- message delivery
- sudden session reset
- silent input failure

There is no explicit feedback.

---

## Error visibility

Errors are not displayed.

Examples include:
- authentication failure
- connection loss
- protocol violation

Errors manifest only through behavior.

---

## No dialogs or alerts

The interface does not display:

- warnings
- confirmations
- popups
- banners
- notifications

All interaction is text-based and immediate.

---

## No configuration interface

The client does not expose:

- settings panels
- toggles
- preferences

All behavior is fixed.

---

## No accessibility guarantees

The interface does not guarantee:

- accessibility compliance
- assistive technology support
- localization

Interface behavior is uniform.

---

## Operator responsibility

Operators are responsible for:

- understanding input semantics
- recognizing interface state
- correcting mistakes manually

The interface does not prevent misuse.

---

## Scope enforcement

If a UI behavior is not documented here, it does not exist.

Operators must not assume hidden affordances.

---

## Final statement

The Litsxg interface is deliberately silent.

It reflects the system without interpretation or assistance.
