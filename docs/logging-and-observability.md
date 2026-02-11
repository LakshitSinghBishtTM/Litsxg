# Logging and Observability

This document defines logging and observability behavior in Litsxg.

Logging is minimal.
Observability is external.
The tool does not attempt to explain itself.

---

## Overview

Litsxg provides no structured observability layer.

There is:
- no metrics system
- no tracing
- no event stream
- no audit framework

Visibility into operation is intentionally limited.

---

## Server logging

The server does not provide a defined logging interface.

Any logging that exists is:
- implementation-dependent
- intended for debugging only
- not part of the protocol or guarantees

Operators must not rely on log format or content.

---

## Client logging

The client does not expose logs to the operator.

Client behavior is observable only through:
- interface behavior
- message flow
- connection success or failure

There is no debug console.

---

## Message logging

Litsxg does not intentionally log message content.

However:
- plaintext messages may appear in logs if enabled
- debugging output may expose message content
- system-level logging may capture traffic

Operators must assume logs may contain plaintext.

---

## Authentication logging

Authentication events are not logged explicitly.

There is no guarantee of:
- login attempt logging
- failure logging
- identity access logging

Authentication visibility is minimal.

---

## Connection logging

Connection events are not explicitly logged.

There is no guarantee of:
- connection open logs
- connection close logs
- failure cause logging

Connection state is inferred from behavior.

---

## Observability via external tools

Operators may use external tools to observe behavior, such as:

- process supervisors
- system logs
- resource monitors
- Tor logs

These tools are outside Litsxg scope.

---

## Tor logging interaction

Tor maintains its own logs.

These logs may reveal:
- connection attempts
- service availability
- circuit behavior

Tor logs are separate from Litsxg logs.

---

## No telemetry

Litsxg does not emit:

- telemetry
- analytics
- usage statistics
- error reports

There is no outbound reporting.

---

## No health endpoints

Litsxg does not expose:

- health checks
- status endpoints
- administrative interfaces

Process liveness is the only indicator.

---

## Operator responsibility

Operators are responsible for:

- deciding whether to enable logging
- managing log retention
- securing log storage
- interpreting failure from limited signals

Observability is a deployment concern.

---

## Scope enforcement

If observability behavior is not documented here, it does not exist.

Operators must not assume hidden diagnostics.

---

## Final statement

Litsxg does not explain itself.

It runs, fails, and restarts without commentary.
