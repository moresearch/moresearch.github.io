---
title: Durable Daemon Pattern — The Problem
date: 2026-07-26
slug: durable-daemon-pattern-problem
summary: AI agent state is a web, not a database. Crash it and the commitments die. The literature knows we're good at accumulating, bad at governing. This is the gap the durable daemon pattern exists to close.
tags: daemons, durable-daemon-pattern, ai-agents, always-on-agents, memory, governance
series: durable-daemon-pattern
---

[Previously: where daemons come from.](https://blog.hackspree.com/#durable-daemon-pattern-genealogy)

> What happens when an AI agent runs for a month?

> State is not a database. It is a web. Delete one node. The web frays.

It learns your preferences — which prospects need soft follow-ups, which Slack channels are noise, which meetings you always skip. It makes commitments on your behalf. It accumulates permissions. It builds inferences. It develops trigger conditions — if a deal goes three days without activity, flag it. It is now more attentive than you are. A commitment depends on a preference. That preference depends on an inference. That inference depends on a Slack message from March. You cannot wipe and rebuild this. It is entangled.

> What survives a crash?

> The chat logs survive. The commitments do not. The agent wakes up amnesiac.

A stateless agent: nothing of the state survives — the binary is fine, the context is gone. An agent with a vector database: the chat logs persist. But not the commitments. Not the permissions. Not the triggers. Not the causal chain. You spend two weeks re-teaching preferences. The dropped commitments become broken promises.

> What does the literature say?

> The field accumulates and retrieves well. It neglects governance, recovery, and forgetting. That gap is what happens when state outruns the runtime.

Ding, Nannapaneni, Liu, and Zhang surveyed 435 papers. [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026). The survey provides six diagnostic axes — Authority, Scope, Mutability, Provenance, Recoverability, Actionability. These are questions you ask *about* state: who controls it? what does it cover? how does it change? where did it come from? can it be restored? does it drive behavior?

These are diagnostic questions. They tell you whether your state is well-governed. They do not give you the architectural guarantees that make it well-governed. The paper gives us the questions. It does not give us the pattern that closes the gap.

> State without durability is a liability. Every piece of knowledge is a piece of the system a crash can destroy.

[Next: the four conditions of the pattern.](https://blog.hackspree.com/#durable-daemon-pattern-definition)

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series.*
