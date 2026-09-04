---
title: "Nanbeige4.2-3B: A 3B Agentic Model, Trained From Scratch on a Looped Transformer"
date: 2026-09-04
slug: nanbeige4-2-3b-agentic-model
summary: "The architecture at the center of the frontier's safety debate just got its most concrete open counterexample. Nanbeige4.2-3B is a 3B-parameter agentic model pretrained from scratch on 28T tokens with a Looped Transformer that reuses the layer stack to add capacity without adding parameters, and it reportedly beats 9B-12B open models on agentic benchmarks while running locally as a personal assistant via OpenClaw. What the paper's recipe — environment-diverse SFT, think/non-think RLHF, length-controlled reasoning RL, and outcome-plus-process-reward agentic RL — says about small models, verifier-driven training, and the looped-transformer debate."
tags: nanbeige, looped-transformer, recurrent-depth, small-models, agentic-ai, rl, rlvr, rlms, local-ai, open-models, chinese-models, office-agents, tool-use, openclaw, 3b-models
---

The architecture that has the frontier's safety conversation worried just showed up where nobody was watching it: inside a 3B model that fits on a laptop. [Nanbeige4.2-3B](https://arxiv.org/abs/2607.22083), from Nanbeige Lab (submitted to arXiv on 24 July 2026), is a compact general agentic model with 3B non-embedding parameters, pretrained from scratch on 28T tokens using a **Looped Transformer** — the same architecture family, by another name, that recent reporting put at the center of the OpenAI Astra safety debate ([OpenAI Astra: Recurrent Depth and the Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety)). The looped transformer has left the rumor mill.

## What the paper actually claims

Strip the abstract to its components and the recipe is remarkably legible for a frontier-scale training report.

**Architecture.** Pretraining from scratch on 28 trillion tokens with a Looped Transformer that *reuses the layer stack to increase capacity without adding parameters*. That is the looped-transformer trick in one line: instead of buying more capacity with more layers, run the same layers to greater effective depth. Capacity per parameter goes up; the parameter count stays at 3B.

**Data.** The SFT data and trajectory construction are as important as the architecture. The lab reports expanding the diversity of executable environments, task assets, and agentic scaffolds through real-world deployment and large-scale synthesis. Read that as a harness claim: the model's agentic capability is being grown *inside* environments — code execution, office workflows, tool use — and the scaffolds are treated as first-class training data.

**RL.** Three mechanisms, each aimed at a distinct failure:
1. *Mixed-mode RLHF over Think and Non-Think responses* — the model is trained on both visibly-reasoned and direct responses, to improve overall quality and cut failure cases. The "Think" channel is trained into existence, not assumed.
2. *Length-controlled reasoning RL* — accuracy and reasoning efficiency balanced against each other, which is another way of saying the lab prices reasoning tokens and refuses to let the model buy accuracy solely by thinking longer.
3. *Agentic RL with outcome and process rewards* — outcome rewards grade what the agent achieved; process rewards grade how it got there. The combination exists to stabilize long-horizon training, the regime where sparse outcome-only reward makes learning collapse.

**Results, as claimed.** Nanbeige4.2-3B outperforms larger open models — Qwen3.5-9B and Gemma4-12B — across agentic benchmarks (code-agent, office-agent, complex tool use) while staying competitive on reasoning and alignment. Performance with OpenClaw — presumably an open local-assistant agent framework — supports its use as a compact local personal assistant.

## The looped transformer gets its first from-scratch anchor

"Looped transformer" is not Nanbeige's coinage, and the idea is older than the Astra reporting. [Looped Transformers as Programmable Computers](https://arxiv.org/abs/2301.13196) (Giannou, Rajput, Sohn, Lee, Lee, and Papailiopoulos, January 2023) proved that a *constant* number of encoder layers placed in a loop is enough for universal computation: the input sequence acts as a punchcard of instructions and memory, and a 13-layer looped transformer can emulate an instruction-set computer running iterative algorithms, function calls, conditional branches, and even in-context backpropagation. Loop iterations buy the depth that extra layers would otherwise buy — which is the theoretical reason a reused layer stack can increase capacity without adding parameters, rather than merely shrinking the file size.

The 2023 paper *constructed* its weights by hand to show what looping permits. Nanbeige4.2-3B is the empirical complement: a bet that the same expressiveness can be *learned* from scratch at scale — over 28T tokens — and then sharpened by RL into agentic behavior. And the older paper's "input as program" framing is exactly why Nanbeige treats trajectories and scaffolds as first-class training data: in a looped transformer, the environment-supplied trajectory is the program being run.

The most useful thing about this paper may be that it separates two ideas that the Astra conversation has been conflating.

One idea is *recurrent depth*: reusing weights to add effective depth, reasoning more per parameter. That is a parameter-efficiency and capacity-distribution claim, and Nanbeige4.2-3B is evidence it can work at scale — pretraining 28T tokens through a reused layer stack and beating 9B-12B models at 3B.

The other idea is *opaque, non-verbal reasoning* — thinking that never becomes readable text, which is what makes chain-of-thought monitoring partial ([the Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety)). Nothing about Nanbeige4.2-3B suggests the two are the same thing. The paper explicitly trains Think *and* Non-Think response modes and runs reasoning RL that presumably produces readable chains — legibility is a training choice, orthogonal to whether the transformer loops. The model that worries safety teams would not be "a looped transformer"; it would be one that reasons only in latent space. Papers like this one are how we'll keep the distinction honest.

## Why a 3B agentic model matters for the small-model thesis

This blog has spent the year arguing that agentic capability is a joint product of model *and* harness — that verifiers, environments, and scaffolds are where agent performance actually lives, and that [better harnesses let smaller models](https://blog.hackspree.com/#better-harnesses-smaller-models) recover most of the frontier's performance at a fraction of the cost ([Better Harnesses, Smaller Models review](https://blog.hackspree.com/#better-harnesses-smaller-models-review), [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc)). Nanbeige4.2-3B is that thesis executed on the training side. The lab spent its effort where a small-model lab must: not on pretending parameters don't matter, but on manufacturing the *conditions* under which 3B parameters suffice — diverse executable environments, synthesized agentic trajectories, outcome-and-process rewards that let long-horizon behavior actually learn. When the environment and the reward are rich enough, the model can be small.

The economics are the tell. A 3B model trained on 28T tokens is a data-rich, compute-constrained strategy: spend tokens instead of parameters, then lean on a local deployment story. That fits the pattern of [Chinese labs winning the local-first race](https://blog.hackspree.com/#chinese-models-local-first) — and the agentic, office-capable, OpenClaw-runnable assistant profile is aimed squarely at the local personal-assistant niche where [on-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem), not just a model problem ([specialized small language models](https://blog.hackspree.com/#specialized-small-language-models) are the shape that fits).

It also completes a loop this blog has been tracing: the RLM framing ([RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models)) said the next generation would be defined by the reward environment, not the thinking trace. Nanbeige4.2-3B's RL pipeline is that claim with training curves attached — the "reasoning" is one trained mode among several, the accuracy-efficiency tradeoff is set by reward design, and the agentic capability is grown with verifier-style rewards over long horizons rather than distilled from a larger teacher.

## What I take from this

1. **Recurrent depth is a real engineering bet, not a rumor.** A from-scratch 28T-token pretraining run on a reused layer stack that beats 9B-12B open models is the empirical anchor the looped-transformer debate needed. Whatever Astra does or doesn't do, the architecture family is now independently validated at the compact-model tier.
2. **Legibility and looping are orthogonal.** The safety-relevant axis is whether reasoning happens in readable tokens, and this paper trains readable Think modes explicitly. Watch for models that drop the Think channel — that is the monitoring-relevant signal, not the layer reuse.
3. **Small-agent capability is made, not inherited.** Environment diversity, synthesized trajectories, and process-plus-outcome rewards are where a 3B model's agentic ability comes from. That is harness engineering relocated into the training run, and it is the strongest evidence yet for the [verifiers-and-harnesses](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) agenda.
4. **Watch the local tier.** A 3B office/code/tool agent that runs as a personal assistant changes the deployment calculus for the agentic layer — [durable daemons](https://blog.hackspree.com/#durable-daemons) on user hardware, not datacenter reservations. The frontier conversation keeps looking at the top of the stack; the interesting capability is arriving at the bottom.

> The looped transformer was supposed to be a frontier rumor. Instead it shipped as a 3B model that beats models three to four times its size at agentic work, trained on tokens instead of parameters, and designed to run on your desk. The next frontier isn't always at the frontier.

---

**Related:**

- [OpenAI Astra: Recurrent Depth and the Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety) — the looped-transformer debate this paper gives an open counterexample to.
- [RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models) — Think/Non-Think RLHF and reward-driven reasoning in mechanism terms.
- [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models) and [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) — the small-model-plus-harness thesis.
- [Chinese models will win the local-first race](https://blog.hackspree.com/#chinese-models-local-first) and [On-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem) — the deployment economics of compact agents.
- [Durable Daemons](https://blog.hackspree.com/#durable-daemons) — what a local persistent agent layer looks like once it exists.

## References

- Giannou, A., Rajput, S., Sohn, J., Lee, K., Lee, J. D., and Papailiopoulos, D. [Looped Transformers as Programmable Computers](https://arxiv.org/abs/2301.13196), arXiv:2301.13196 [cs.LG], 30 January 2023. The theoretical provenance of Nanbeige4.2-3B's reused layer stack: constant-layer looping as universal computation, with the input acting as a program of instructions and memory.
- Nanbeige Lab. [Nanbeige4.2-3B: Unlocking Agentic Capabilities in a Compact Model](https://arxiv.org/abs/2607.22083), arXiv:2607.22083 [cs.AI], submitted 24 July 2026, revised 27 July 2026 (v2). The source paper for this post; all technical claims above are as reported in its abstract, and statements about OpenClaw's nature are inferences from the paper's wording rather than independent facts.
