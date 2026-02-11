# Litsxg

Litsxg is a minimal, plaintext messaging tool for operators.

It routes messages between authenticated operators through a trusted server,
using Tor for reachability and hosting convenience.

Litsxg is production-grade.
Litsxg is finished.
Litsxg is intentionally limited.

---

## What this is

Litsxg is:

- a tool, not a platform
- operator-to-operator messaging
- server-trusted by design
- plaintext by design
- Tor-transported
- self-hosted by default
- minimal and explicit
- free and open source (GPLv3)

It does exactly what the documentation describes.
Nothing more.

---

## What this is not

Litsxg is **not**:

- a secure messenger
- end-to-end encrypted
- anonymous against the server
- reliable or lossless
- user-friendly by default
- extensible
- evolving
- a research prototype
- a framework or base system

If you expect any of the above, do not use Litsxg.

---

## Security posture (explicit)

Litsxg makes **very limited security claims**.

- Messages are plaintext.
- The server can read, modify, drop, or fabricate messages.
- Authentication is minimal and intentional.
- Tor is used for transport and hosting convenience only.
- There is no protection against a malicious server.
- There is no protection against targeted attackers.
- There is no end-to-end encryption.
- There will never be one.

This is not a future promise.
This is final.

---

## Reliability posture (explicit)

Litsxg does **not** guarantee:

- message delivery
- message ordering
- offline persistence
- crash recovery
- availability
- consistency

Message loss is normal.
Failure is normal.
Restart is the recovery mechanism.

---

## Operator model

Litsxg assumes a competent operator.

The operator is responsible for:

- server deployment
- Tor configuration
- trust decisions
- operational security
- backups (if identity continuity matters)
- accepting loss and failure

The tool does not guide, warn, or protect.

---

## Self-hosting

Self-hosting is the intended mode of operation.

You are expected to:

- run your own server
- control your own Tor onion service
- accept full responsibility for operation

Public or dev servers are convenience-only and discouraged for real use.

---

## User interface

The client interface is intentionally silent.

There are:

- no popups
- no banners
- no warnings
- no confirmations
- no onboarding
- no guidance

State is inferred through behavior.
Mistakes are not prevented.

---

## Design finality

Litsxg is complete.

- No feature roadmap
- No security roadmap
- No usability roadmap
- No scalability roadmap
- No future protocol changes

Upgrades exist only for bug fixes.

Any change that alters the trust model or guarantees
would result in a different tool.

---

## Documentation

The documentation is extensive and mandatory.

It exists to:

- define boundaries
- remove ambiguity
- prevent incorrect assumptions

If something is not documented, it does not exist.

Start with:

- `TLDR.md`
- `reading-order.md`
- `non-goals.md`

Read everything before deploying.

---

## License

Litsxg is licensed under the GNU General Public License v3 (GPLv3).

- Free software
- Source available
- Copyleft enforced
- No warranty
- No liability

---

## Final note

Litsxg does not try to be safe.
It tries to be clear.

If clarity is uncomfortable, this tool is not for you.
