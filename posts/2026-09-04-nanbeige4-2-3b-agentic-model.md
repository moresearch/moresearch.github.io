---
title: "Nanbeige4.2-3B: A 3B Agentic Model, Trained From Scratch on a Looped Transformer"
date: 2026-09-04
slug: nanbeige4-2-3b-agentic-model
summary: "Nanbeige4.2-3B is a 3B-parameter agentic model pretrained from scratch on 28T tokens with a Looped Transformer that reuses the layer stack to add capacity without adding parameters, and it reportedly beats 9B-12B open models on agentic benchmarks while running locally as a personal assistant. This post walks the full recipe — architecture, environment-diverse SFT, Think/Non-Think RLHF, length-controlled reasoning RL, outcome-plus-process agentic RL — with close readings of the four sibling papers that make it legible (Looped Transformers as Programmable Computers, DeepSeek-R1, Kimi k1.5, SWE-RL), and spells out the key insights up front."
tags: nanbeige, looped-transformer, recurrent-depth, small-models, agentic-ai, rl, rlvr, rlms, local-ai, open-models, chinese-models, office-agents, tool-use, openclaw, 3b-models
---

The architecture that has the frontier's safety conversation worried just showed up where nobody was watching it: inside a 3B model that fits on a laptop. [Nanbeige4.2-3B](https://arxiv.org/abs/2607.22083), from Nanbeige Lab (submitted to arXiv on 24 July 2026), is a compact general agentic model with 3B non-embedding parameters, pretrained from scratch on 28T tokens using a **Looped Transformer** — the same architecture family, by another name, that recent reporting put at the center of the OpenAI Astra safety debate ([OpenAI Astra: Recurrent Depth and the Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety)). The looped transformer has left the rumor mill, and the paper deserves a careful read because almost every ingredient in it is a bet this blog has been tracking from the other side.

## Key insights

1. **A 3B model trained on a reused layer stack reportedly beats 9B–12B open models at agentic work.** Nanbeige4.2-3B outperforms Qwen3.5-9B and Gemma4-12B across code-agent, office-agent, and complex tool-use benchmarks while staying competitive on reasoning and alignment. Parameter count is not the binding constraint for agentic capability.
2. **"Capacity without parameters" is a computational-depth claim, not a file-compression trick.** A constant number of layers placed in a loop is provably enough for universal computation ([Looped Transformers as Programmable Computers](https://arxiv.org/abs/2301.13196)): loop iterations buy the depth that extra layers would otherwise buy. Nanbeige4.2-3B is the first from-scratch, at-scale instance of that theory being *learned* rather than hand-constructed.
3. **The recipe's real content is the RL stack, and each of its three parts has a verifiable literature behind it.** Think/Non-Think RLHF is the RLVR thinking line ([DeepSeek-R1](https://arxiv.org/abs/2501.12948)); length-controlled reasoning RL is the reasoning-efficiency line ([Kimi k1.5](https://arxiv.org/abs/2501.12599)); agentic RL with outcome and process rewards is the long-horizon agent line ([SWE-RL](https://arxiv.org/abs/2502.18449)).
4. **Think and Non-Think are trained modes, not architecture accidents.** The model is explicitly trained to reason visibly *and* to answer directly. Legibility is orthogonal to layer looping — which is the distinction the [Astra safety debate](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety) keeps blurring.
5. **Agentic capability is made in the training environment, then stabilized by reward design.** Diverse executable environments and synthesized trajectories are first-class training data; outcome-plus-process rewards exist to keep long-horizon RL from collapsing. That is harness engineering relocated into the training run.
6. **Watch the local tier.** A 3B office/code/tool agent that runs as a personal assistant is the first concrete shape of the local agent layer — [durable daemons](https://blog.hackspree.com/#durable-daemons) on user hardware rather than datacenter reservations.

## What the paper actually claims

Strip the abstract to its components and the recipe is remarkably legible for a frontier-scale training report.

**Architecture.** Pretraining from scratch on 28 trillion tokens with a Looped Transformer that *reuses the layer stack to increase capacity without adding parameters*. Instead of buying capacity with more layers, the same layers are run to greater effective depth. Capacity per parameter goes up while the parameter count stays at 3B.

**Data.** The SFT data and trajectory construction matter as much as the architecture. The lab reports expanding the diversity of executable environments, task assets, and agentic scaffolds through real-world deployment and large-scale synthesis. The model's agentic capability is grown *inside* environments — code execution, office workflows, tool use — and the scaffolds themselves are treated as training data.

**RL.** Three mechanisms, each aimed at a distinct failure mode:

1. *Mixed-mode RLHF over Think and Non-Think responses* — quality and failure-case reduction across both visibly-reasoned and direct responses. The "Think" channel is trained into existence, not assumed.
2. *Length-controlled reasoning RL* — accuracy and reasoning efficiency balanced against each other; the lab prices reasoning tokens and refuses to let the model buy accuracy solely by thinking longer.
3. *Agentic RL with outcome and process rewards* — outcome rewards grade what the agent achieved, process rewards grade how it got there, and the combination stabilizes long-horizon training, where sparse outcome-only reward makes learning collapse.

**Results, as claimed.** Nanbeige4.2-3B outperforms larger open models across agentic benchmarks and stays competitive on reasoning and alignment; performance with OpenClaw — presumably an open local-assistant agent framework — supports its use as a compact local personal assistant.

## The architecture bet, read through its founding paper

"Looped transformer" is not Nanbeige's coinage. [Looped Transformers as Programmable Computers](https://arxiv.org/abs/2301.13196) (Giannou, Rajput, Sohn, Lee, Lee, and Papailiopoulos, January 2023) proved that a *constant* number of encoder layers placed in a loop suffices for universal computation: the input sequence acts as a punchcard of instructions and memory, and a 13-layer looped transformer can emulate an instruction-set computer running iterative algorithms, function calls, conditional branches, and even in-context backpropagation. Loop iterations buy the depth that extra layers would otherwise buy — the theoretical reason a reused layer stack can increase capacity without adding parameters.

The gap between that result and Nanbeige4.2-3B is instructive. The 2023 paper *constructed* its weights by hand to demonstrate what looping permits. Nanbeige is the empirical complement: a bet that the same expressiveness can be *learned* from scratch at scale — 28T tokens, then RL — rather than engineered. And the older paper's "input as program" framing is exactly why trajectories and scaffolds show up as first-class training data: in a looped transformer, the environment-supplied trajectory is the program being run. A scaffold is not a wrapper around the model; it is input to the machine.

The safety-relevant corollary, which the Astra conversation keeps under-specifying, is that looping says nothing about legibility. Nanbeige explicitly trains Think *and* Non-Think response modes, which is direct evidence that a looped transformer can produce readable reasoning when the training reward selects for it. Recurrent depth is a parameter-efficiency claim; opaque, non-verbal reasoning is a separate training choice about where thinking happens. The model that worries safety teams would not be "a looped transformer" — it would be one that reasons only in latent space.

## The RL stack, mapped to three sibling papers

The architecture gets the headline, but the three RL ingredients are where the paper's contribution actually lives — and each one names a prior result directly.

**DeepSeek-R1: thinking can be grown, not taught.** [R1](https://arxiv.org/abs/2501.12948) (DeepSeek-AI, January 2025) showed that pure RL against verifiable rewards — with no human-labeled reasoning trajectories — incentivizes the *emergence* of advanced reasoning patterns: self-reflection, verification, dynamic strategy adaptation. Its R1-Zero variant famously produced a legible "aha moment" mid-trajectory, evidence that chain-of-thought-like behavior is an RL-grown policy, not a prompted behavior. That is the intellectual foundation of Nanbeige's whole post-pretraining program: reasoning as an emergent policy of a verifier-driven loop, with readable Think traces as an *option* the reward can select for rather than an architectural given. Everything this blog has said about [RLMs](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models) traces to this paper.

**Kimi k1.5: accuracy must be priced, or models buy it with verbosity.** [k1.5](https://arxiv.org/abs/2501.12599) (Kimi Team, January 2025) established the practical RL recipe Nanbeige's middle ingredient references: long-context RL scaling plus policy-optimization methods that explicitly control reasoning length, achieving o1-matching results (77.5 AIME, 96.2 MATH-500) *without* MCTS, value functions, or process reward models. The "length-controlled reasoning RL" in Nanbeige's abstract is k1.5's lesson applied: reasoning tokens are a budget, and a reward that ignores length trains models that think long because nothing stops them. Nanbeige states the same tradeoff — "balance accuracy and reasoning efficiency" — as an explicit pipeline stage.

**SWE-RL: agentic RL works when the environment supplies the reward.** [SWE-RL](https://arxiv.org/abs/2502.18449) (Wei et al., February 2025) was the first attempt to scale RL reasoning to real-world software engineering, training on open-source software evolution data with a lightweight rule-based reward (similarity between generated and ground-truth solutions). The result — Llama3-SWE-RL-70B at 41.0% on SWE-bench Verified, the best reported for a sub-100B model at the time — made two points Nanbeige inherits: first, that long-horizon agentic behavior (issue → patch → test) is trainable when the environment gives feedback; second, that the failure mode of such training is long-horizon instability, which is precisely why Nanbeige pairs outcome rewards with process rewards "to stabilize long-horizon training." SWE-RL is the code-agent ancestor of Nanbeige's code-agent and office-agent claims.

## Why a 3B agentic model matters for the small-model thesis

This blog has argued all year that agentic capability is a joint product of model *and* harness — that verifiers, environments, and scaffolds are where agent performance lives, and that [better harnesses let smaller models](https://blog.hackspree.com/#better-harnesses-smaller-models) recover most of the frontier's performance at a fraction of the cost ([review](https://blog.hackspree.com/#better-harnesses-smaller-models-review), [the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc)). Nanbeige4.2-3B is that thesis executed on the training side: the lab spent its effort not on pretending parameters don't matter, but on manufacturing the conditions under which 3B parameters suffice — diverse executable environments, synthesized agentic trajectories, and outcome-plus-process rewards that let long-horizon behavior actually learn. When the environment and the reward are rich enough, the model can be small.

The economics are the tell. A 3B model trained on 28T tokens is a data-rich, compute-constrained strategy: spend tokens instead of parameters, then lean on local deployment. That fits the pattern of [Chinese labs winning the local-first race](https://blog.hackspree.com/#chinese-models-local-first), and the office-capable, OpenClaw-runnable profile targets the local personal-assistant niche where [on-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem) ([specialized small language models](https://blog.hackspree.com/#specialized-small-language-models) are the shape that fits).

## What I take from this

1. **For labs copying the recipe:** the loop is necessary but not sufficient — the reusable lessons are environment-first data and reward design, not the layer reuse. Any lab can rent the architecture; the moat is the executable environments and the reward engineering, which is where R1, k1.5, and SWE-RL all point.
2. **For safety monitoring:** legibility is a dial, not a property. Nanbeige trains Think modes because its reward selects for them; a successor that drops the Think channel is the signal to watch, not the presence of looping ([Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety)).
3. **For the frontier debate:** recurrent depth now has an open, compact, verifiable instance. The conversation should move from "is looping dangerous?" to "what did the training reward select?" — reward, not trace, is where capability is actually shaped.
4. **For builders of local agents:** this is the first clear profile of the local agent generation — office/code/tool capability at 3B, RL-stabilized over long horizons, designed to sit on user hardware as a [durable daemon](https://blog.hackspree.com/#durable-daemons).

> The looped transformer was supposed to be a frontier rumor. Instead it shipped as a 3B model that beats models three to four times its size at agentic work, trained on tokens instead of parameters, and designed to run on your desk. The next frontier isn't always at the frontier — and this time it came with a readable paper trail.

---

**Related:**

- [OpenAI Astra: Recurrent Depth and the Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety) — the looped-transformer debate this paper gives an open counterexample to.
- [RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models) — Think/Non-Think RLHF and reward-driven reasoning in mechanism terms.
- [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models) and [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) — the small-model-plus-harness thesis.
- [Chinese models will win the local-first race](https://blog.hackspree.com/#chinese-models-local-first) and [On-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem) — the deployment economics of compact agents.
- [Durable Daemons](https://blog.hackspree.com/#durable-daemons) — what a local persistent agent layer looks like once it exists.

## References

- Nanbeige Lab. [Nanbeige4.2-3B: Unlocking Agentic Capabilities in a Compact Model](https://arxiv.org/abs/2607.22083), arXiv:2607.22083 [cs.AI], submitted 24 July 2026, revised 27 July 2026 (v2). The source paper for this post; all technical claims above are as reported in its abstract, and statements about OpenClaw's nature are inferences from the paper's wording rather than independent facts.
- Giannou, A., Rajput, S., Sohn, J., Lee, K., Lee, J. D., and Papailiopoulos, D. [Looped Transformers as Programmable Computers](https://arxiv.org/abs/2301.13196), arXiv:2301.13196 [cs.LG], 30 January 2023. The theoretical provenance of the reused layer stack: constant-layer looping as universal computation, with the input acting as a program of instructions and memory.
- DeepSeek-AI. [DeepSeek-R1: Incentivizing Reasoning Capability in LLMs via Reinforcement Learning](https://arxiv.org/abs/2501.12948), arXiv:2501.12948 [cs.CL], January 2025. The RLVR anchor: reasoning as an emergent, verifier-driven policy, with legible thinking as a trained option.
- Kimi Team. [Kimi k1.5: Scaling Reinforcement Learning with LLMs](https://arxiv.org/abs/2501.12599), arXiv:2501.12599 [cs.CL], January 2025. The length-control lineage behind Nanbeige's "length-controlled reasoning RL."
- Wei, Y., Duchenne, O., Copet, J., et al. [SWE-RL: Advancing LLM Reasoning via Reinforcement Learning on Open Software Evolution](https://arxiv.org/abs/2502.18449), arXiv:2502.18449 [cs.CL], February 2025. The agentic-RL ancestor: rule-based outcome rewards over long-horizon software tasks, and the instability Nanbeige's outcome-plus-process reward design addresses.
