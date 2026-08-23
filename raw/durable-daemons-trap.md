---
title: Durable Daemons — State Entanglement in Long-Running Agent Systems
date: 2026-07-26
slug: durable-daemons-trap
summary: AI agent state is a dependency graph, not a key-value store. Crash it and the commitments die — not the chat logs, the obligations. The Ding et al. survey of 435 papers confirms: we accumulate well, govern poorly. This is the failure mode the durable daemons pattern must prevent.
tags: daemons, durable-daemons-pattern, state-management, ai-agents, always-on-agents, failure-modes
series: durable-daemons-pattern
---

[Previously: the architectural lineage.](https://blog.hackspree.com/#durable-daemons)

Run an AI agent for a month. It learns that you prefer soft follow-ups with manufacturing clients. It commits on your behalf — "I'll send the deck by Tuesday." It accumulates permissions — read access to your calendar, send access to your email, API keys to your CRM. It builds inferences — Acme Corp stalls deals in Q4, the VP of Engineering responds faster on Signal than email. It develops trigger conditions — if a deal goes three days without activity, flag it. It is now more attentive than you are.

Now crash the process. Deploy new code. Rotate the API keys.

> State is a dependency graph, not a key-value store. A commitment depends on a preference. That preference depends on an inference. That inference depends on a Slack message from March. Delete one node. The graph frays in directions you cannot predict.

A stateless agent: nothing of the state survives — the binary is fine, the context is gone. An agent with a vector database: the chat logs persist. But not the commitments. Not the permissions. Not the triggers. Not the causal chain that connects "Acme stalls Q4" to "don't send the pricing sheet until January." The agent wakes up amnesiac. You spend two weeks re-teaching preferences. The dropped commitments become broken promises.

In a choreography of durable daemons, this failure is catastrophic. If a sales daemon drops its commitment to send the deck, the support daemon still expects the deck was sent. The ops daemon pre-allocated capacity based on the deal progressing. Three daemons, one dropped state node, cascading inconsistency. The choreography doesn't fail loudly — it fails silently, each daemon acting on state that is no longer true.

> The chat logs survive. The obligations do not. This is not a bug. It is a category error. We built persistence for conversations. We needed persistence for commitments. In a choreography, one daemon's amnesia is every daemon's corruption.

Ding, Nannapaneni, Liu, and Zhang surveyed 435 papers. [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026). Their finding: the field accumulates and retrieves well. It neglects governance. It neglects recovery. It neglects forgetting. The survey provides six diagnostic axes — Authority, Scope, Mutability, Provenance, Recoverability, Actionability. These are questions you ask *about* state: who controls it? what does it cover? how does it change? where did it come from? can it be restored? does it drive behavior? These are necessary diagnostic questions. They are not sufficient architectural guarantees.

> The paper gives us the diagnostic axes. It does not give us the guarantees. The gap between question and guarantee is where the failure mode lives.

[Next: the pattern specification — four conditions that close the gap.](https://blog.hackspree.com/#durable-daemons-definition)

---

*Part of the [Durable Daemons](/tags/durable-daemons-pattern) series.*
