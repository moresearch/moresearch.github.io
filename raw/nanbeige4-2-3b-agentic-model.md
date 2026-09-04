---
title: "Nanbeige4.2-3B: A 3B Agentic Model, Trained From Scratch on a Looped Transformer"
date: 2026-09-04
slug: nanbeige4-2-3b-agentic-model
summary: "Nanbeige4.2-3B is a 3B-parameter agentic model pretrained from scratch on 28T tokens with a Looped Transformer that reuses its layer stack, reportedly beating 9B-12B open models at code, office and tool-use agent work — and it is close enough to run on a laptop. Key insights up front, four sibling papers as context, and an Ollama example."
tags: nanbeige, looped-transformer, recurrent-depth, small-models, agentic-ai, rl, rlvr, rlms, local-ai, open-models, chinese-models, office-agents, tool-use, openclaw, 3b-models
---

The architecture at the center of the frontier's safety conversation just shipped as a 3B model you can run on a laptop. [Nanbeige4.2-3B](https://arxiv.org/abs/2607.22083) (Nanbeige Lab, July 2026) is a compact agentic model — 3B non-embedding parameters, pretrained from scratch on 28T tokens with a **Looped Transformer** that reuses the layer stack to increase capacity without adding parameters — and it reportedly outperforms Qwen3.5-9B and Gemma4-12B across code-agent, office-agent and complex tool-use benchmarks. The same architecture family, by another name, sits at the center of the [OpenAI Astra safety debate](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety).

## Key insights

1. **For agentic work, parameters are not the binding constraint.** A 3B model on a reused layer stack beats 9B–12B open models while staying competitive on reasoning and alignment.
2. **"Capacity without parameters" is a depth claim, not a file-size trick.** A constant number of layers placed in a loop is provably enough for universal computation ([Giannou et al., 2023](https://arxiv.org/abs/2301.13196)): loop iterations buy the depth extra layers would otherwise buy. Nanbeige4.2-3B is that theory *learned* from scratch rather than hand-constructed.
3. **The recipe is the RL stack, and each part names a prior result.** Think/Non-Think RLHF is the RLVR thinking line ([DeepSeek-R1](https://arxiv.org/abs/2501.12948)); length-controlled reasoning RL is the efficiency line ([Kimi k1.5](https://arxiv.org/abs/2501.12599)); outcome-plus-process agentic RL is the long-horizon line ([SWE-RL](https://arxiv.org/abs/2502.18449)).
4. **Looping is not hidden reasoning.** Think and Non-Think are trained modes; the model reasons visibly when the reward selects for it. Legibility is orthogonal to layer reuse — the distinction the Astra conversation keeps blurring.
5. **Capability is built in the training environment, and the economics are local-first.** Scaffolds and trajectories are first-class training data; 3B plus 28T tokens is a data-rich, compute-constrained bet that ends as a durable daemon on user hardware.

## The recipe in one breath

- **Architecture.** 28T-token from-scratch pretraining through a reused layer stack: more effective depth per parameter at 3B total.
- **Data.** SFT on diverse executable environments, task assets and agentic scaffolds from real-world deployment and large-scale synthesis.
- **RL.** Mixed-mode RLHF over Think and Non-Think responses (quality, fewer failures); length-controlled reasoning RL (accuracy vs. efficiency); agentic RL with outcome *and* process rewards (stabilizes long-horizon training).
- **Results.** Beats Qwen3.5-9B and Gemma4-12B on agentic benchmarks; runs locally as a personal assistant (OpenClaw).

## Four papers that make it legible

**Looped Transformers as Programmable Computers** ([Giannou et al., 2023](https://arxiv.org/abs/2301.13196)) proved a constant number of encoder layers in a loop can emulate an instruction-set computer — the input acts as a punchcard of instructions and memory. Its weights were hand-constructed to show what looping *permits*; Nanbeige is the bet that the same expressiveness can be *learned* at scale. The "input as program" framing is also why trajectories matter: in a looped transformer, the environment-supplied trajectory is the program being run.

**DeepSeek-R1** ([2025](https://arxiv.org/abs/2501.12948)) showed pure RL against verifiable rewards grows reasoning — self-reflection, verification — with no human-labeled traces, producing legible thinking (the R1-Zero "aha moment") as an emergent, reward-selected behavior. That is the foundation of Nanbeige's whole post-pretraining program, and of why readable Think traces are an *option*, not a given.

**Kimi k1.5** ([2025](https://arxiv.org/abs/2501.12599)) is the length-control ancestor: o1-matching reasoning without MCTS, value functions or process reward models, because the RL explicitly prices reasoning length. Nanbeige's "length-controlled reasoning RL" is the same lesson — models left unpriced think long because nothing stops them.

**SWE-RL** ([Wei et al., 2025](https://arxiv.org/abs/2502.18449)) was the first RL scaled to real software engineering: rule-based outcome rewards over long-horizon issue-to-patch tasks, hitting 41% on SWE-bench Verified at 70B. It proved long-horizon agentic behavior is trainable from environment feedback — and that the failure mode is instability, which is exactly what Nanbeige's outcome-plus-process reward design addresses.

## Why it matters

Agentic capability is a joint product of model *and* harness — [verifiers, environments and scaffolds are where agent performance lives](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), and [better harnesses let smaller models](https://blog.hackspree.com/#better-harnesses-smaller-models) close most of the gap. Nanbeige4.2-3B executes that thesis on the training side: it manufactures the conditions under which 3B parameters suffice instead of adding parameters. That fits the [Chinese labs winning the local-first race](https://blog.hackspree.com/#chinese-models-local-first), and it gives the [local agent layer](https://blog.hackspree.com/#durable-daemons) its first concrete profile: office/code/tool capability at 3B, RL-stabilized over long horizons.

## Try it

The 3B-class claim is one command away on a laptop — no API key, no scaffold:

```bash
ollama run tomng/nanbeige4.1:3b "Explain systems of systems in simple terms!"
```

Run the same question through the largest model you can reach and compare: the point of the length-control and efficiency work is that the small model should stay crisp where the big one gets verbose.

> The looped transformer was supposed to be a frontier rumor. Instead it shipped as a 3B model that beats models three to four times its size at agentic work — and it runs on your desk. The next frontier isn't always at the frontier.

---

**Related:**

- [OpenAI Astra: Recurrent Depth and the Limits of Chain-of-Thought Safety](https://blog.hackspree.com/#openai-astra-recurrent-depth-safety)
- [RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models)
- [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models) and [the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc)
- [Durable Daemons](https://blog.hackspree.com/#durable-daemons)

## References

- Nanbeige Lab. [Nanbeige4.2-3B: Unlocking Agentic Capabilities in a Compact Model](https://arxiv.org/abs/2607.22083), arXiv:2607.22083 [cs.AI], July 2026. Source paper; claims above are as reported in its abstract.
- Giannou, A., et al. [Looped Transformers as Programmable Computers](https://arxiv.org/abs/2301.13196), arXiv:2301.13196 [cs.LG], January 2023.
- DeepSeek-AI. [DeepSeek-R1: Incentivizing Reasoning Capability in LLMs via Reinforcement Learning](https://arxiv.org/abs/2501.12948), arXiv:2501.12948 [cs.CL], January 2025.
- Kimi Team. [Kimi k1.5: Scaling Reinforcement Learning with LLMs](https://arxiv.org/abs/2501.12599), arXiv:2501.12599 [cs.CL], January 2025.
- Wei, Y., et al. [SWE-RL: Advancing LLM Reasoning via Reinforcement Learning on Open Software Evolution](https://arxiv.org/abs/2502.18449), arXiv:2502.18449 [cs.CL], February 2025.
