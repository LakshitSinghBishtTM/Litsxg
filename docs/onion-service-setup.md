# Onion Service Setup

This document describes how to expose the Litsxg server as a Tor onion service.

This is the preferred deployment model.
No alternative exposure methods are required or documented here.

---

## Overview

Litsxg servers are intended to be reachable via a Tor onion service.

An onion service:
- removes the need for a public IP address
- avoids direct network exposure
- provides a stable service identifier

This document assumes familiarity with basic Tor operation.

---

## Prerequisites

Before configuring an onion service, the operator must have:

- Tor installed and functioning
- permission to modify Tor configuration
- the Litsxg server binary available

Misconfigured Tor environments are out of scope.

---

## Onion service model

Litsxg uses a single onion service.

Properties:
- one onion address per server instance
- TCP service exposed through Tor
- no additional ports or services

The onion address uniquely identifies the server.

---

## Basic Tor configuration

To expose the Litsxg server, configure Tor with an onion service.

Example configuration:

```
HiddenServiceDir /var/lib/tor/litsxg/
HiddenServicePort 9696 127.0.0.1:9696
```

Properties:
- `HiddenServiceDir` stores onion service keys
- `HiddenServicePort` maps the onion service port to the local server port

Paths and ports may be adjusted as needed.

---

## Onion address generation

After starting Tor with the configured service:

- Tor generates a new onion service key
- the onion address is written to a file in `HiddenServiceDir`

Example:

```
/var/lib/tor/litsxg/hostname
```

This file contains the onion address used by clients.

---

## Server binding requirements

The Litsxg server must:

- listen on the local address specified in Tor configuration
- accept connections on the mapped port
- not bind to public network interfaces unless explicitly intended

Local-only binding is preferred.

---

## Permissions and security

The onion service directory must:

- be readable only by the Tor process
- not be exposed to other system users
- not be backed up insecurely

Loss of the onion service key results in a new address.

---

## Restart behavior

On restart:

- the onion address remains stable if keys are preserved
- active client connections terminate
- clients must reconnect

There is no graceful reconnection.

---

## Multiple instances

Running multiple server instances requires:

- separate onion service directories
- distinct local ports
- distinct Tor configuration entries

Instances are isolated.

---

## Failure modes

Common failure modes include:

- Tor not running
- incorrect port mapping
- permission errors on the service directory

These failures result in the server being unreachable.

---

## Scope enforcement

This document describes onion service exposure only.

- no clearnet hosting is documented
- no reverse proxies are described
- no load balancing is supported

---

## Final statement

Onion services are the expected deployment mechanism for Litsxg.

They provide sufficient reachability without increasing system complexity.