---
title: Durable Daemon Pattern — The Trap
date: 2026-07-26
slug: durable-daemon-pattern-trap
summary: Run an AI agent for a month. It learns, commits, infers, triggers. Then crash it. The chat logs survive. The commitments do not. State is not a database — it is a web. Delete one node. The web frays.
tags: daemons, durable-daemon-pattern, ai-agents, always-on-agents, memory, governance
series: durable-daemon-pattern
---

[Previously: the genealogy of daemons.](https://blog.hackspree.com/#durable-daemons)

Run an AI agent for a month. It learns that you prefer soft follow-ups with manufacturing clients. It commits on your behalf — "I'll send the deck by Tuesday." It accumulates permissions — read access to your calendar, send access to your email, API keys to your CRM. It builds inferences — Acme Corp stalls deals in Q4, the VP of Engineering responds faster on Signal than email. It develops trigger conditions — if a deal goes three days without activity, flag it. It is now more attentive than you are.

Now crash the process. Deploy new code. Rotate the API keys.

> State is not a database. It is a web. A commitment depends on a preference. That preference depends on an inference. That inference depends on a Slack message from March. Delete one node. The web frays in directions you cannot predict.

A stateless agent: nothing of the state survives — the binary is fine, the context is gone. An agent with a vector database: the chat logs persist. But not the commitments. Not the permissions. Not the triggers. Not the causal chain that connects "Acme stalls Q4" to "don't send the pricing sheet until January." The agent wakes up amnesiac. You spend two weeks re-teaching preferences. The dropped commitments become broken promises.

> The chat logs survive. The commitments do not. The agent wakes up amnesiac. This is not a bug. It is a category error. We built memory for conversations. We needed memory for obligations.

Ding, Nannapaneni, Liu, and Zhang surveyed 435 papers. [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026). Their finding: the field accumulates and retrieves well. It neglects governance. It neglects recovery. It neglects forgetting. The survey provides six diagnostic axes — Authority, Scope, Mutability, Provenance, Recoverability, Actionability. These are questions you ask *about* state: who controls it? what does it cover? how does it change? where did it come from? can it be restored? does it drive behavior? These are necessary questions. They are not sufficient answers.

> The paper gives us the diagnostic questions. It does not give us the architectural guarantees. The gap between question and guarantee is where the trap lives.

[Next: the pattern — four conditions that close the gap.](https://blog.hackspree.com/#durable-daemon-pattern-definition)

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series.*
