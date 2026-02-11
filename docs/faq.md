# FAQ

This document answers common questions about Litsxg.

The answers are direct.
They do not soften constraints.
They do not provide reassurance.

---

## What is Litsxg?

Litsxg is a messaging tool.

It routes plaintext messages between operators via a server over Tor.
It is production-grade and intentionally minimal.

---

## Who is Litsxg for?

Litsxg is for operators who:

- control their own server
- accept plaintext messaging
- accept message loss
- value explicit behavior over safeguards

It assumes technical competence and operational discipline.

---

## Is Litsxg secure?

Litsxg is secure only within its stated claims.

It does not provide:
- end-to-end encryption
- message confidentiality
- integrity protection
- anonymity against the server

If you require these properties, Litsxg is not suitable.

---

## Does Tor make messages private?

No.

Tor provides transport reachability and resistance to casual network observation.
It does not hide message content from the server.

Messages remain plaintext.

---

## Can the server read messages?

Yes.

The server can read, modify, drop, or fabricate messages.
This is a design property, not a bug.

---

## Are messages stored?

Messages are not durably stored.

They may exist temporarily:
- in memory
- in offline buffers
- in client display buffers

They are lost on restart or failure.

---

## Are offline messages reliable?

No.

Offline messages are buffered opportunistically.
They may be lost without notice.

There are no delivery guarantees.

---

## Is there a mobile client?

No.

Only the documented client exists.
No additional platforms are planned.

---

## Will end-to-end encryption be added later?

No.

The design is finished.
Cryptographic upgrades are explicitly out of scope.

---

## Can I use Litsxg for sensitive communication?

No.

Using Litsxg for high-risk or sensitive communication is misuse.
The tool does not defend against such threats.

---

## Can I run a public server?

You can, but it is discouraged.

Public servers:
- require trust from operators
- provide no guarantees
- increase misuse risk

Self-hosting is preferred.

---

## Is Litsxg anonymous?

No.

Operators authenticate with persistent identities.
The server knows who is connected and what they send.

---

## Does Litsxg log activity?

Litsxg does not intentionally log activity.

However:
- plaintext may appear in logs if enabled externally
- Tor may produce logs
- system logs may capture data

Logging control is external.

---

## Is Litsxg actively maintained?

Litsxg is considered complete.

Maintenance consists only of:
- bug fixes
- correctness fixes

There is no feature roadmap.

---

## Where do I get help?

Documentation is the primary reference.

There is no guaranteed support.
There is no SLA.

---

## Final statement

If you are uncomfortable with any answer above, do not use Litsxg.

The tool will not adapt to expectations.
