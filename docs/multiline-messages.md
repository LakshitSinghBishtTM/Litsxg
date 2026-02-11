# Multiline Messages

This document describes multiline message usage from the operator perspective.

Multiline messaging is explicit.
It requires discipline.
There are no safeguards.

---

## Overview

Multiline messages allow an operator to send structured or extended text as a single message.

Multiline messages:
- are plaintext
- are transmitted as a single unit
- require explicit termination

There is no automatic detection of intent.

---

## Entering multiline mode

Multiline mode is entered by sending a message with the multiline marker.

### Entry format

```
<target_callsign> $ <initial text>
```

This command:
- fixes the target callsign
- switches the client into multiline mode
- begins local buffering

---

## Writing multiline content

While in multiline mode:

- each line entered is appended to the message buffer
- lines are not transmitted immediately
- formatting is preserved exactly

Whitespace and empty lines are preserved.

---

## Terminating multiline mode

Multiline mode must be terminated explicitly.

### Termination format

```
$
```

The termination marker:
- must appear alone on a line
- immediately ends multiline mode
- triggers message transmission

Failure to terminate leaves the client in multiline mode.

---

## Transmission behavior

On termination:

- all buffered lines are joined using newline characters
- the complete message is transmitted to the server
- the server treats it as a single message

Partial transmission does not occur.

---

## Failure during multiline input

If failure occurs during multiline mode:

- all buffered content is lost
- no partial message is delivered
- the client exits multiline mode on reconnect

There is no recovery.

---

## Multiline and offline messaging

If the recipient is offline:

- the multiline message may be buffered
- buffering is volatile
- delivery is not guaranteed

Multiline messages are not treated specially when offline.

---

## No nesting or editing

Multiline mode does not support:

- nesting
- editing previously entered lines
- preview or confirmation

Once entered, content cannot be modified.

---

## No visual indicators

The interface provides no indication that multiline mode is active.

Operators must track state mentally.

---

## Operator discipline

Operators must:

- terminate multiline mode explicitly
- verify message boundaries manually
- avoid accidental multiline entry

Mistakes are not prevented.

---

## Scope enforcement

Multiline behavior outside this document is undefined.

Operators must not assume additional affordances.

---

## Final statement

Multiline messaging in Litsxg is powerful but unforgiving.

Correct usage depends entirely on operator discipline.