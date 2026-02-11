# Threat Model Table

This document presents the Litsxg threat model in tabular form.

The table is descriptive, not evaluative.
“Not mitigated” means intentionally ignored.

---

## Threat coverage table

| Threat category                         | Considered | Mitigated | Notes |
|----------------------------------------|------------|-----------|-------|
| Passive network observation            | Yes        | Partial   | Tor provides resistance at the transport layer |
| Active network manipulation            | No         | No        | Out of scope |
| Server compromise                      | Yes        | No        | Server is trusted by design |
| Malicious server operator              | Yes        | No        | Accepted risk |
| Message content disclosure             | Yes        | No        | Messages are plaintext |
| Metadata analysis                      | No         | No        | Out of scope |
| Traffic correlation                    | No         | No        | Out of scope |
| Authentication brute force             | Yes        | No        | No rate limiting or hardening |
| Credential theft                       | Yes        | No        | No recovery or protection mechanisms |
| Denial of service                      | Yes        | No        | No mitigation |
| Replay attacks                         | No         | No        | Out of scope |
| Message tampering                      | Yes        | No        | No integrity protection |
| Identity impersonation (server-side)   | Yes        | No        | Server can impersonate |
| Client compromise                      | Yes        | No        | Local environment trusted |
| Legal or coercive threats              | Yes        | No        | No resistance |
| Accidental data retention              | Partial    | Partial   | Minimal persistence only |
| Infrastructure failure                 | Yes        | Partial   | Failure accepted, not prevented |

---

## Interpretation rules

- “Considered” means explicitly acknowledged.
- “Mitigated” means a deliberate mechanism exists.
- Absence of mitigation is intentional.

No row implies future work.

---

## Scope reminder

This table reflects the final threat model.

Operators must not:
- infer mitigations from consideration
- assume partial mitigation implies robustness
- treat ignored threats as oversights

---

## Final statement

The threat model is narrow by design.

The table exists to prevent misinterpretation, not to inspire confidence.
