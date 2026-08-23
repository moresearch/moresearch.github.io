---
title: "Hacker Laws for Agentic Software Engineering: Conway's Law"
date: 2026-08-24
slug: hacker-laws-ase-conways-law
summary: "Law 1 of 12. Conway's Law says the technical boundaries of a system will reflect the structure of the organisation. The ASE key insight: the organisation is now humans, agents, and the harness that wires them — and the software will mirror the topology of the agents that made it, so you change the software by changing the topology."
tags: hacker-laws, agentic-software-engineering, series, conways-law, topology, harness
series: hacker-laws-for-agentic-software-engineering
---

**Law 1 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Next: [Brooks' Law](https://blog.hackspree.com/#hacker-laws-ase-brooks-law).

## The Law

> The technical boundaries of a system will reflect the structure of the organisation. ([hacker-laws](https://github.com/dwmkerr/hacker-laws#conways-law))

## The Key Insight for Agentic Software Engineering

Conway's Law was about the org chart, and the org chart just changed: it now includes the agents, and — more importantly — the **harness that wires them**. In an agentic software engineering shop, the structure that shapes the software is not only who reports to whom; it is how the agents communicate, what shared state they read, which delegations are allowed, and where the verifiers sit. The software will mirror that topology, because the topology decides what each agent can see and therefore what it can build. Give two agents the same repository but different shared-state boundaries and they will produce systems with seams in different places — the seam lands where the communication seam lands.

The harness is the org chart. The [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) topology produces central-plan, delegated-feature systems; the [blackboard](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard) topology produces choreographed, shared-state systems; the [sequential pipeline](https://blog.hackspree.com/#harnessing-agentic-ai-systems-sequential-pipeline-routing) topology produces layered, linear systems. The model is the same in all three — the topology is what changes, and the topology is a harness decision, not a model decision ([the system, not the agent](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems)).

The flip side is sharper: agents also mirror each other. The [distillation loop](https://blog.hackspree.com/#agents-are-distillation-at-scale) — small models trained on a large model's outputs — makes the "organisation" a single voice, and the software it produces is correspondingly single-voiced. When the agentic org has one mind, the software has one architecture, for better and for worse; when it has many independent minds with narrow channels between them, the software is modular in exactly the places the channels are narrow.

The ASE reading of Conway's Law: **the software will mirror the topology of the agents that made it, so design the topology — the harness — to get the seams you want.** If you want modular systems, make the communication channels between agents narrow and the shared state explicit; if you want integrated systems, wire the agents tightly. The law was never an excuse for the org chart; it was an instruction to take the org chart seriously. In agentic software engineering, the org chart is a file you can edit.

## References

- dwmkerr. [hacker-laws — Conway's Law](https://github.com/dwmkerr/hacker-laws#conways-law) and [Conway's Law on Wikipedia](https://en.wikipedia.org/wiki/Conway%27s_law).
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker), [blackboard](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard), [sequential pipeline routing](https://blog.hackspree.com/#harnessing-agentic-ai-systems-sequential-pipeline-routing).
- [Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale) — the single-voice organisation.
