# Terminology

This document defines the terminology used throughout the Litsxg documentation.

Terms are used consistently.
Meanings are fixed.
If a term is not defined here, it should not be assumed.

---

## Operator

An operator is any party that interacts with Litsxg.

This includes:
- running the client
- running the server
- operating both client and server

There is no distinction between “user” and “administrator” in the documentation.
All responsibility is attributed to the operator.

---

## Client

The client is the operator-facing component of Litsxg.

It provides:
- a graphical interface
- message input and display
- network connectivity over Tor

The client does not enforce security guarantees beyond those explicitly documented.

---

## Server

The server is the routing component of Litsxg.

It:
- accepts incoming connections
- authenticates operators
- routes messages
- stores limited transient state

The server is trusted by design and can observe message contents.

---

## Session

A session is an active connection between a client and the server.

Sessions:
- are established over Tor
- are not guaranteed to persist
- may terminate at any time due to network or server conditions

Session loss does not imply error recovery.

---

## Identity

Identity refers to an operator’s logical identifier within Litsxg.

Identity is:
- established via authentication
- maintained by the server
- not cryptographically protected beyond basic password handling

Identity continuity does not imply security guarantees.

---

## Authentication

Authentication is the process by which an operator claims an identity.

It exists solely to:
- distinguish operators
- route messages correctly

Authentication is not intended to provide strong security properties.

---

## Message

A message is a unit of text transmitted from one operator to another.

Messages:
- are plaintext
- may be single-line or multiline
- may be lost under certain conditions

Message delivery is best-effort.

---

## Multiline message

A multiline message is a message composed of multiple lines of text.

Multiline messages are delimited using explicit protocol markers.
They are treated as a single logical message by the server.

---

## Offline message

An offline message is a message temporarily stored by the server for later delivery.

Offline messages:
- are not guaranteed to persist
- may be lost on server restart or failure
- exist only as a convenience

---

## Tor

Tor refers to the Tor network used for connectivity.

In Litsxg, Tor is used to:
- provide reachability
- simplify deployment

Tor is not used to imply anonymity or strong privacy guarantees.

---

## Failure

Failure refers to any condition where expected behavior is interrupted.

This includes:
- network loss
- server restart
- client termination
- message loss

Failure is considered normal.

---

## Guarantee

A guarantee is a property explicitly stated in documentation.

If a property is not documented as a guarantee, it does not exist.

---

## Final note

Terminology is part of the interface.

Misinterpreting terms leads to incorrect assumptions about the tool.
