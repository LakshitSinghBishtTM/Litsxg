# Deployment Guide

This document describes how to deploy Litsxg in a production setting.

Deployment is simple by design.
There is no orchestration layer, installer, or automation framework.
All deployment steps are explicit and manual.

---

## Overview

Litsxg is deployed as a standalone server process exposed via a Tor onion service.

The canonical deployment model is:
- single server instance
- single onion service
- self-hosted environment
- operator-controlled lifecycle

Other deployment models are out of scope.

---

## Deployment requirements

Before deployment, the operator must have:

- a supported operating system
- Tor installed and functioning
- the Litsxg server binary
- filesystem access for persistent data

No container runtime is required.
No package manager integration is assumed.

---

## Environment preparation

The operator must:

- choose a dedicated directory for the server
- ensure correct filesystem permissions
- decide where persistent data will reside

The server does not create or manage directories beyond its own runtime needs.

---

## Server binary placement

Place the server binary in a fixed location.

Example:
- `/opt/litsxg/server`
- `/usr/local/bin/litsxg-server`

The binary location does not affect runtime behavior.

---

## Persistent data location

The server stores minimal persistent data.

The operator must decide:
- where the credential database is stored
- how it is secured
- whether it is backed up

Message content is not stored persistently.

---

## Tor installation

Tor must be installed and operational.

The operator is responsible for:
- installing Tor via system means
- enabling the Tor service
- ensuring Tor starts on boot if required

Tor installation is not managed by Litsxg.

---

## Onion service configuration

The operator must configure a Tor onion service.

This includes:
- selecting a local server port
- mapping the onion service port
- securing the onion service directory

Details are defined in `onion-service-setup.md`.

---

## Server startup order

The recommended startup order is:

1. Start Tor.
2. Verify Tor is running.
3. Start the Litsxg server.
4. Verify the server is listening locally.
5. Verify the onion service is reachable.

Deviating from this order may result in temporary unavailability.

---

## Process supervision

Litsxg does not include a supervisor.

The operator may optionally use:
- system service managers
- external watchdogs
- manual supervision

Supervision behavior is external.

---

## Firewall considerations

When using onion services:

- no inbound clearnet ports need to be opened
- only local loopback traffic is required

Firewall configuration is the operator’s responsibility.

---

## Resource provisioning

Litsxg has low resource requirements.

The operator must ensure:
- sufficient memory for in-memory buffers
- sufficient CPU for concurrent connections

There is no dynamic scaling.

---

## Deployment validation

After deployment, the operator should verify:

- the server accepts connections
- authentication works as expected
- messages are routed correctly

There is no built-in validation tool.

---

## Failure handling

Deployment must assume failure.

The operator must be prepared to:
- restart the server
- restart Tor
- accept message loss

Failure handling is manual.

---

## Scope enforcement

This guide defines the supported deployment process.

Automated deployment, clustering, and managed hosting are out of scope.

---

## Final statement

Deploying Litsxg is intentionally uncomplicated.

Operational responsibility remains entirely with the operator.
