# Reading Order

This document defines the canonical reading order for the Litsxg documentation.

The order is intentional.
Each file assumes context established by the files before it.
Skipping files is allowed, but misunderstanding is the responsibility of the operator.

---

## Entry and orientation

These files define what the tool is and whether you should proceed at all.

1. README.md  
2. TLDR.md  
3. reading-order.md  
4. non-goals.md  
5. design-philosophy.md  
6. project-scope.md  
7. terminology.md  

Stop here if any of these conflict with your expectations.

---

## Operator framing and mental model

These files establish how Litsxg is intended to be operated and understood.

8. tool-overview.md  
9. operator-expectations.md  

Proceed only if you accept the responsibilities described.

---

## System architecture

These files describe the structure of the tool at a system level.

10. system-overview.md  
11. component-boundaries.md  
12. client-architecture.md  
13. server-architecture.md  
14. data-model.md  
15. state-machines.md  
16. concurrency-model.md  
17. lifecycle.md  

---

## Protocol definition

These files define the on-wire and logical behavior of the system.

18. protocol.md  
19. message-formats.md  
20. authentication-flow.md  
21. multiline-protocol.md  
22. error-handling.md  
23. protocol-non-guarantees.md  

Do not infer protocol properties beyond what is explicitly stated.

---

## Tor and networking behavior

These files define how Litsxg uses Tor and how network failures are handled.

24. why-tor.md  
25. tor-assumptions.md  
26. onion-service-setup.md  
27. client-connection-behaviour.md  
28. network-failure-semantics.md  
29. embedded-tor-design.md  
30. port-and-binding-model.md  

Tor is used for reachability.
No anonymity guarantees beyond those stated should be assumed.

---

## Security model and limits

These files define the security boundaries of the tool.

31. threat-model.md  
32. threatmodeltable.md  
33. security-claims.md  
34. security-non-claims.md  
35. password-handling.md  
36. plaintext-messaging.md  
37. known-weaknesses.md  
38. misuse-patterns.md  
39. attacker-models.md  

Security properties are narrow and deliberate.

---

## Reliability and failure semantics

These files describe behavior under failure and loss.

40. reliability-model.md  
41. offline-message-semantics.md  
42. crash-recovery.md  
43. consistency-model.md  

Assume failure is possible at all times.

---

## Operation and deployment

These files are required for operating Litsxg in production.

44. operator-manual.md  
45. deployment-guide.md  
46. self-hosting.md  
47. dev-server-policy.md  
48. upgrade-policy.md  
49. backup-and-recovery.md  
50. logging-and-observability.md  
51. resource-usage.md  
52. support-policy.md  

Operation without reading these files is unsupported.

---

## Interface behavior

These files describe operator interaction with the interface.

53. ui-behaviour.md  
54. multiline-messages.md  

The interface does not compensate for misuse.

---

## Practical reference

These files address common questions and mistakes.

55. faq.md  
56. common-mistakes.md  
57. opsec.md  
58. no-bloat-policy.md  

These files do not override earlier documents.

---

## Project governance and closure

These files define stability, governance, and final context.

59. stability-policy.md  
60. audit-scope.md  
61. contributing.md  
62. license.md  
63. philosophy.md  
64. comparisons.md  
65. design-history.md  
66. non-roadmap.md  
67. conclusion.md  

---

## Final note

This documentation set is not optimized for speed or convenience.

It is optimized to make assumptions explicit and limits unavoidable.
Reading less does not make the tool safer.
