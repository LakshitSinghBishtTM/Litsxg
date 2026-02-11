# Multiline Protocol

This document defines the multiline message protocol used by Litsxg.

Multiline handling is explicit.
There is no implicit aggregation or automatic termination.
Operator intent must be stated clearly.

---

## Overview

The multiline protocol allows an operator to send a message composed of multiple lines as a single logical unit.

Multiline messages:
- are plaintext
- are transmitted as a single message
- require explicit entry and exit

There is no streaming mode.

---

## Entry into multiline mode

Multiline mode is entered explicitly by the operator.

### Entry format

```
<target_callsign> $ <initial_content>
```

Properties:
- `$` indicates entry into multiline mode
- `<initial_content>` may not be empty
- the target callsign is fixed for the duration of the multiline message

After this command, the client enters multiline input mode.

---

## Multiline input phase

While in multiline mode:

- each subsequent line is treated as message content
- lines are buffered locally by the client
- content is not transmitted immediately

Whitespace and formatting are preserved exactly as entered.

---

## Termination of multiline mode

Multiline mode is terminated explicitly.

### Termination format

```
$
```

Properties:
- the termination marker must appear alone on a line
- no surrounding whitespace is allowed
- termination triggers transmission of buffered content

Upon termination, the client exits multiline mode.

---

## Transmission behavior

On termination:

- buffered lines are concatenated using newline characters
- the resulting content is transmitted as a single message
- the server routes the message normally

There is no partial transmission.

---

## Failure during multiline mode

Failure during multiline mode may occur due to:

- connection loss
- client termination
- server restart

In such cases:
- buffered content is lost
- no partial message is delivered
- multiline mode is not resumed

There is no recovery.

---

## Server handling

The server treats multiline messages as opaque content.

The server:
- does not inspect internal structure
- does not validate formatting
- does not attempt to split content

From the server perspective, a multiline message is equivalent to a single-line message.

---

## Absence of nesting

Multiline mode cannot be nested.

Attempting to enter multiline mode while already in multiline mode results in undefined behavior.

---

## Absence of timeouts

There is no timeout for multiline mode.

The client remains in multiline mode until:
- termination marker is received
- connection is lost

Operators must terminate multiline mode explicitly.

---

## Scope enforcement

Multiline behavior outside this document is undefined.

Operators must not assume automatic termination or recovery.

---

## Final statement

The multiline protocol in Litsxg is explicit and unforgiving.

It exists to allow structured input without adding protocol complexity.