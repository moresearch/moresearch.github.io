---
title: "Architecting Principles for Systems-of-Systems (Applied to Agentic AI)"
date: 2026-08-30
slug: systems-of-systems-architecting-principles
summary: "Mark Maier's 1998 Systems Engineering paper defined a system-of-systems: elements that are operationally and managerially independent, geographically distributed, and producing behavior that resides in no single element — architected by evolution, not top-down design. Every multi-agent deployment is a system-of-systems, whether its builders admit it or not. The five principles and what they mean when the elements are agents."
tags: [systems-of-systems, architecture, maier, multi-agent-systems, emergent-behavior, interfaces, standards, orchestration, harness-engineering, agentic-ai]
---

## The definition that keeps coming back

In 1998, Mark Maier published a paper in *Systems Engineering* that defined a category and then told architects how to fail at it less often: [*Architecting Principles for Systems-of-Systems*](https://doi.org/10.1002/(sici)1520-6858(1998)1:4%3C267::aid-sys3%3E3.0.co;2-d). A system-of-systems (SoS), he argued, is an assemblage of components that individually qualify as systems and that collectively satisfy five characteristics:

1. **Operational independence** — the elements operate usefully on their own; disassemble the SoS and each part still does its job.
2. **Managerial independence** — the elements are separately acquired and managed, each with a continuing existence of its own.
3. **Geographic distribution** — the elements live apart and integration happens over networks.
4. **Emergent behavior** — the SoS does things that no element does; the capability resides in the interactions.
5. **Evolutionary development** — the SoS never appears fully formed; functions, purposes, and elements are added, removed, and modified with experience.

Read that list against any serious multi-agent deployment and the match is exact. Each agent can be invoked alone and is built by a team that owns it (1 and 2). The agents run on different services and machines, integrated over protocols (3). The orchestrated outcome — a report assembled from six specialized agents, a negotiation closed by an ensemble vote — belongs to no single agent (4). And the platform accretes new agents and capabilities continuously, never by one grand design (5). Maier was not anticipating LLM agents in 1998. He was describing a structural category, and agentic systems fell into it the moment they became multi-agent. The paper's value is that it tells you what kind of architecting works when you are in that category — and the honest answer is: not the kind most teams attempt.

## Why classical architecture fails the SoS test

Maier's central claim is that a SoS cannot be architected the way a single system is. Classical systems architecture assumes an architect with authority: someone who can decide the internal design of every component and the sequence in which they are built. A system-of-systems has no such authority — the elements are operationally and managerially independent, so **no one controls the whole**. The SoS architect therefore cannot design the elements; they can only design the conditions under which independently evolving elements interoperate.

That is why the paper is titled "architecting principles" rather than "design principles." The architect's instruments are interfaces, standards, and evolution — not control. For agentic systems this is the difference between teams that try to *design* their multi-agent platform and teams that *cultivate* it. The teams that fail are the ones who specify every agent's internal prompt, tool list, and memory layout from a central committee and expect the ensemble to behave. The teams that succeed spend their architecture budget on the harness: the contracts, the routing, the shared state, the evaluation — the things that cross the boundaries. This blog's [harness patterns series](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) is, read back through this lens, a catalog of SoS architecting: every pattern is a decision about the interfaces between agents, not a recipe for any single agent's internals.

The principles that follow are Maier's, each translated into the terms of agentic architecture.

## Principle 1 — The architecture must have stable intermediate forms

Maier's sharpest principle: a SoS must be fieldable and useful at *every* stage of its evolution, not merely in the imagined end state. You cannot architect for the final integrated system and defer usefulness until everything is connected, because the end state never arrives — the SoS evolves forever. Every intermediate form must be a working system.

For agents this is a design constraint with teeth: **every agent must be useful alone before it is useful together.** An agent that cannot run standalone — that depends on the fleet to make any sense — is not a stable intermediate form; it is a liability you can only ship in a big bang. The systems that last are the ones built as [durable daemons](https://blog.hackspree.com/#durable-daemons): each element is a long-lived, independently operable process that earns its keep on its own and becomes more valuable when networked. The same logic applies to the harness: a tool, a skill, or an integration should be shippable in isolation. If the architecture only works in its final, fully-connected state, the architecture is not a system-of-systems — it is a distributed monolith with an evolution problem.

## Principle 2 — Minimize commonality

Maier's recommendation is counterintuitive on its face: **minimize commonality, don't maximize it.** The natural instinct of the integrating architect is to force a shared schema, a shared framework, a shared ontology across every element — one context, one vocabulary, one platform for all. The SoS principle says the opposite: commonality should be forced only where the interface standards demand it, and everywhere else the elements should retain maximum autonomy.

The reason is managerial and operational independence. The moment you force element internals to conform to a common standard, you have assumed ownership of elements you do not own — and you have frozen the SoS at the pace of its slowest-moving component. For agent systems this shows up as the universal-schema trap: the team that insists every agent emit events in one canonical format, store memory in one vector schema, and expose one tool-calling convention. The blackboard pattern is the instructive exception — [shared state is shared risk](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems), and the shared schema is only worth its cost where the interface genuinely requires it. Everywhere else, let each agent be internally whatever its owning team needs it to be, and pay for the integration at the boundary.

## Principle 3 — Design for the evolution of standards

If the interface standards are the backbone of a SoS, then **standards that cannot evolve become the constraint that freezes the SoS.** Maier's principle is to design the standards with their own evolution path in mind — versioned, replaceable, with a defined migration — or they will outlive their usefulness and hold the whole system hostage to them.

For agentic platforms this is the protocol lesson, and the evidence is all around: tool schemas, [function-calling contracts](https://en.wikipedia.org/wiki/System_of_systems), context protocols, and memory formats are the interface standards of the multi-agent age, and they are young. The platforms that design their protocols to evolve — versioned schemas, additive changes, legacy support windows, explicit deprecation — will survive the ones that bake today's tool format in stone. The failure mode is familiar: a single version of the "official" agent protocol, a shared codebase that only one team may touch, a breaking change that forces every element to upgrade simultaneously. That is not evolution; that is a deadline. A SoS is only as evolvable as its least evolvable standard.

## Principle 4 — The leverage is at the interfaces

Maier's most practical directive: in a system-of-systems, **the architect's leverage is concentrated in the interfaces between elements** — the network structure, the exchange formats, the protocols — and not in the internals of the elements. The architect who spends effort redesigning an element they don't own is wasting effort; the architect who decides *what crosses the boundary* is shaping the whole system.

For agentic architecture this is where the modern practice converges with the 1998 principle. The [harness is the product](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems): the routing rules, the tool contracts, the context budget, the evaluation gates — these are interface decisions, and they are where multi-agent platforms win or lose. Whether an ensemble behaves like a [directed orchestra or a blackboard of collaborators](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard) is decided at the boundaries: what each agent is allowed to see, write, and call. The internals of any individual agent — its prompt, its fine-tune, its tools — are mostly someone else's problem, sometimes literally. The boundary is your problem, and it is the whole ballgame. This is also the economic lens: interfaces are where the [mechanisms and incentives](https://blog.hackspree.com/#mechanism-design-and-network-economics-for-agentic-markets) live, and where the value of the ensemble is created or destroyed.

## Principle 5 — Assess by emergent behavior

If the capability of a SoS resides in the interactions, then **you cannot evaluate a SoS by summing or averaging its elements' performance.** Maier's assessment principle follows directly from the definition: the metric that matters is the emergent behavior of the whole. Each element may be flawless by its own tests and the ensemble still fails end-to-end — and vice versa.

This is the principle that most multi-agent teams violate. The agent has a 95% pass rate on its unit suite; therefore the pipeline must be fine. But the pipeline is a system-of-systems, and its behavior is emergent: the failure is in the handoff, the context that got truncated, the tool that the orchestrator called with the wrong argument, the ordering that made two agents overwrite shared state. This blog's [empirical-SE discipline](https://blog.hackspree.com/#empirical-se-what-studies-say) is the practical expression of the same rule: evaluate the assembled system with end-to-end, outcome-level tests, not per-element proxies. And the reason to test the whole rather than the parts is the same reason the [blind men and the elephant](https://blog.hackspree.com/#blind-men-elephant-business-process-automation) story is a SoS story: the whole does not exist inside any single element, and no element's view of it is the truth.

## The four species of system-of-systems

Maier also gave the field its taxonomy, later extended by the DoD's [Systems Engineering Guide for Systems of Systems](https://en.wikipedia.org/wiki/System_of_systems) — and the taxonomy is the first diagnostic an architect should run, because **the wrong topology choice is a category error.** Four types:

- **Directed** — built and managed to fulfill specific purposes; elements integrated and operated under central management. Translation: the orchestrator-worker topology, where one coordinator owns the run and the workers are extensions of it.
- **Acknowledged** — recognized objectives and a designated integrator with resources, but the elements retain independent ownership, funding, and development. Translation: the platform model — a company's central platform coordinates agents that individual teams own and evolve at their own pace. Most serious enterprise deployments are acknowledged whether they know it or not.
- **Collaborative** — elements voluntarily cooperate toward agreed purposes, with weak or absent central management. Translation: blackboard ensembles and voting patterns, where agents choose to share state or consensus without a director.
- **Virtual** — no central management and no agreed purpose; large-scale behavior emerges from the interactions. Translation: [agentic markets](https://blog.hackspree.com/#mechanism-design-and-network-economics-for-agentic-markets), where no coordinator exists, agents act on their own incentives, and the system-level outcome is emergent.

The category errors are the expensive ones. Directing a virtual SoS — mandating behavior in elements no one owns — produces resistance and gaming, the thing [mechanism design](https://blog.hackspree.com/#mechanism-design-and-network-economics-for-agentic-markets) exists to avoid. Treating an acknowledged SoS as directed — assuming central ownership of elements that other teams control — produces integration that breaks the moment the other teams ship. And treating a directed SoS as collaborative produces drift where the coordinator should simply decide. The taxonomy is not academic: it tells you how much control you actually have, which tells you which principles apply.

## What this means

Five checks for any multi-agent architecture, straight from Maier:

1. **Name your species.** Directed, acknowledged, collaborative, or virtual — and design the topology, the interfaces, and the evaluation to match the control you actually have, not the control you wish you had.
2. **Every element must stand alone.** If an agent is not useful in isolation, it is not a stable intermediate form; fix that before networking it.
3. **Resist forced commonality.** The shared schema is a cost, not a virtue. Pay it only where the interface genuinely requires it.
4. **Budget your architecture effort at the boundaries.** The harness, the contracts, the routing, the shared state — not the internals of agents you don't own.
5. **Evaluate the emergent whole.** End-to-end outcomes over per-agent metrics, always.

Maier's 1998 insight survives because it is structural: whenever a large system is made of independently operating, independently owned, geographically distributed parts whose value emerges from their interaction, the architect who tries to design it like a monolith loses, and the architect who cultivates it wins. The multi-agent age did not invent the system-of-systems. It just made the category impossible to ignore.
