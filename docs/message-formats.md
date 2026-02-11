# Message Formats

This document defines the concrete message formats used by the Litsxg protocol.

All formats are plaintext.
All formats are line-based.
No implicit fields exist.

If a format is not defined here, it must not be assumed.

---

## Overview

Messages exchanged between client and server fall into the following categories:

- authentication commands
- routed messages
- multiline message blocks
- control commands

Each category has a fixed format.

---

## Line termination

All protocol lines are terminated by a newline character.

- Lines are processed sequentially.
- Trailing newline characters are not part of message content.
- Empty lines have no semantic meaning unless explicitly stated.

---

## Authentication command format

Authentication commands are issued before message routing is allowed.

### Format

```
$ <callsign> <password>
```

### Properties

- `$` indicates an authentication-related command
- `<callsign>` is the operator identifier
- `<password>` is plaintext and transmitted as-is
- fields are separated by single spaces

Authentication is session-scoped.

---

## Routed message format

Routed messages are sent after authentication completes.

### Format

```
<target_callsign> <message_content>
```

### Properties

- `<target_callsign>` identifies the recipient
- `<message_content>` is plaintext
- message content may not be empty
- the first space separates address and content

The server prepends the sender identity before delivery.

---

## Delivered message format

Messages delivered to recipients are formatted by the server.

### Format

```
<sender_callsign>: <message_content>
```

### Properties

- sender identity is added by the server
- content is delivered verbatim
- no metadata is attached

Delivery order is not guaranteed.

---

## Multiline message initiation

Multiline messages are explicitly initiated.

### Format

```
<target_callsign> $ <initial_content>
```

### Properties

- `$` indicates entry into multiline mode
- `<initial_content>` may be empty
- subsequent lines are treated as message content

Multiline mode is per-session and explicit.

---

## Multiline message continuation

After initiation, subsequent lines are treated as message content.

### Format

```
<content_line>
```

### Properties

- lines are appended in order
- no transformation is applied
- content is buffered until termination

---

## Multiline message termination

Multiline messages are explicitly terminated.

### Format

```
$
```

### Properties

- termination marker must appear alone on a line
- buffered content is transmitted as a single message
- termination exits multiline mode

Failure to terminate leaves the client in multiline mode.

---

## Control command format

Certain commands affect protocol state rather than routing messages.

### Format

```
$ <command>
```

### Properties

- commands are prefixed with `$`
- commands are interpreted by the server
- unsupported commands are ignored

Control command behavior is limited.

---

## Whitespace rules

Whitespace handling is strict.

- leading whitespace is significant
- trailing whitespace is preserved in message content
- multiple spaces are not collapsed

Operators must format input carefully.

---

## Invalid formats

Invalid message formats result in:

- silent ignore
- session termination
- undefined behavior

Clients must not rely on error feedback.

---

## Scope enforcement

No additional message formats exist.

Any format not described in this document is unsupported.

---

## Final statement

Message formats in Litsxg are intentionally minimal.

They exist to be read, understood, and implemented without abstraction.