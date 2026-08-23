---
title: "Agents Aren't Magic. They're Distillation at Scale."
date: 2026-07-19
slug: agents-are-distillation-at-scale
summary: "A 350M parameter model, fine-tuned for a single epoch, crushes ChatGPT on tool calling by 51 points. The future of agents is not bigger models. It's smaller ones that know exactly what to do and nothing else."
tags: ai, agents, distillation, small-models, tool-calling, efficiency
---

"Agents aren't magic. They're distillation at scale." — **Andrej Karpathy**

As [@0xMortyx put it](https://x.com/0xMortyx/status/2078468804276019504): 99.99% of an LLM's capacity is wasted on data it never needed. **Small model + right tools + closed loop = terrifying capability.**

In March 2026, a team at AWS proved the formula with numbers. They took `facebook/opt-350m` — 350 million parameters, from 2022, smaller than what runs on a laptop — and fine-tuned it on ToolBench for a single epoch.

| Model | ToolBench Pass Rate |
|---|---|
| Fine-tuned OPT-350M (this work) | **77.55%** |
| ToolLLaMA-DFS | 30.18% |
| ChatGPT-CoT | 26.00% |
| ToolLLaMA-CoT | 16.27% |

A 350M model, one epoch, tripled the closest competitor's score. It didn't just beat ChatGPT. It rendered the large model a rounding error.

## The 99.99% problem

A large language model knows the capital of Burkina Faso, the plot of *Anna Karenina*, and seventeen ways to cook an egg. When you ask it to call an API, every one of those facts is dead weight. The model is carrying an encyclopedia through a door that only needs a key.

Tool calling requires one thing: understanding the instruction, selecting the right API, formatting the call. The model doesn't need Dostoevsky. It needs to know the difference between GET and POST. For that, 350 million parameters isn't just sufficient — it's optimal. The large model is distracted by its own knowledge. The small model is focused because it has no capacity for anything else.

> A generalist knows everything and can do anything, badly. A specialist knows one thing and does it perfectly. Stop hiring generalists for narrow tasks.

## Distillation is the mechanism

The paper calls it Supervised Fine-Tuning. What's actually happening: the model is being taught the *pattern*, not the reasoning.

187,542 instruction-solution pairs from ToolBench. Each one a Thought-Action-Action Input triplet: think, select the tool, format the call. Feed enough triplets through a small model and it internalizes the rhythm. Not the underlying logic. Not the world knowledge. Just the surface form of effective tool use — the cadence of think-act-observe, the template of the API call, the recovery strategy when the call fails.

> You don't need the teacher's brain. You need the teacher's answers to enough questions that the student internalizes the shape of correct responses. Once the student has the shape, the teacher is overhead.

## The closed loop

Here's where it gets terrifying.

A model at 77.55% pass rate can generate its own training data — its successful trajectories become instruction-solution pairs for the next round. Round 1: 187K human examples. Round 2: human examples + N self-generated successes. Round 3: even more. Each round compresses the failure modes into the distribution. The model improves without growing.

> The closed loop is the moat, not the model. Switching models is expensive not because the model is irreplaceable but because the loop has accumulated months of self-generated training data the replacement hasn't seen.

## Small model + right tools + closed loop = terrifying capability

The AWS paper is not isolated. AgentSymbiotic: 8B LLaMA reaches 48.5% on WebArena, approaching Claude-3.5 at 52.1%. SCoRe: a 7B student matches a 72B teacher across 12 benchmarks — student generates trajectories, teacher corrects only the earliest error, RL closes the gap. Every result has the same shape: small model, targeted training, closed-loop refinement. The large model is bootstrapping, not destination.

The economics are brutal. If 350M params handles 77% of tool calls and costs 1/1000th of a frontier model — why is the frontier model handling routine API calls?

> The future is not one omniscient model. It's a swarm of tiny specialists, each distilled to do one thing perfectly, coordinated by a router that costs nothing. The large model trains them. The small models do the work. The loop is the product.

## How to build

Stop fine-tuning large models for narrow tasks. You are paying for parameters you don't need.

Start with the smallest model that can absorb the pattern. Fine-tune on high-quality task-specific examples. Deploy. Log successes and failures. Feed successes back into training. The loop tightens. If you need a large model, use it only as a judge — to evaluate, correct, and generate the next round of training data.

The AWS team proved 350 million parameters is enough. Your tool-calling task is not harder than ToolBench. Your model is too large.

---

*Paper: P. Jhandi, O. Kazi, S. Subramanian, N. Sendas — [Small Language Models for Efficient Agentic Tool Calling](https://arxiv.org/abs/2512.15943) (AAAI 2026). Quote via [@0xMortyx](https://x.com/0xMortyx/status/2078468804276019504), citing Andrej Karpathy.*
