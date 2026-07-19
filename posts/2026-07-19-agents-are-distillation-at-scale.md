---
title: "Agents Aren't Magic. They're Distillation at Scale."
date: 2026-07-19
slug: agents-are-distillation-at-scale
summary: "A 350M parameter model, fine-tuned for a single epoch, crushes ChatGPT on tool calling by 51 percentage points. The future of agents is not bigger models. It's smaller ones that know exactly what to do and nothing else."
tags: ai, agents, distillation, small-models, tool-calling, efficiency
---

In March 2026, a team at AWS published a result that should reset how you think about AI agents.

They took `facebook/opt-350m` — a 350-million-parameter model from 2022, smaller than what you can run on a laptop — and fine-tuned it on ToolBench for a single epoch. One pass through the data. Then they ran it against the standard agentic tool-calling benchmark.

The results:

| Model | ToolBench Pass Rate |
|---|---|
| Fine-tuned OPT-350M (this work) | **77.55%** |
| ToolLLaMA-DFS | 30.18% |
| ChatGPT-CoT | 26.00% |
| ToolLLaMA-CoT | 16.27% |

A 350M parameter model, trained for one epoch, beat ChatGPT by 51 percentage points. It didn't just edge past. It tripled the score of the closest competitor. The small model didn't compete with the large model. It rendered the large model irrelevant.

> The paper's title is understated: *Small Language Models for Efficient Agentic Tool Calling: Outperforming Large Models with Targeted Fine-tuning.* The title could have been: *Everything You Believe About Model Size Is Wrong.*

## The 99.99% problem

Why does a 350M model crush a model with, presumably, hundreds of billions of parameters on a specific task?

Because **99.99% of an LLM's capacity is wasted on data it never needed for the task at hand.**

A large language model knows the capital of Burkina Faso. It knows the plot of *Anna Karenina*. It knows the chemical formula for polyethylene, the history of the Ming dynasty, and seventeen ways to cook an egg. When you ask it to call an API, all of that knowledge is dead weight. The model is carrying an encyclopedia through a door that only requires a key.

Tool calling is a narrow skill. It requires understanding the instruction, selecting the right API, formatting the call correctly, and interpreting the result. That's it. The model doesn't need to know who wrote *The Brothers Karamazov*. It needs to know the difference between a GET and a POST. For that, 350 million parameters is not just sufficient — it's optimal. The larger model is distracted by its own knowledge. The small model is focused because it has no choice but to be.

> A generalist knows everything and can do anything, badly. A specialist knows one thing and does it perfectly. Tool calling is a specialist skill. Stop hiring generalists for it.

## Distillation is the mechanism

The paper calls it Supervised Fine-Tuning. But what's actually happening is **distillation at scale.**

The training data — 187,542 instruction-solution pairs from ToolBench — encodes the tool-calling knowledge of every model and human that contributed to the benchmark. Each example is a Thought-Action-Action Input triplet: think about the problem, select the tool, format the call. Feed enough of these triplets through a small model and it internalizes the pattern. Not the underlying reasoning. Not the world knowledge. *The pattern.*

That's what distillation means in this context. You're not teaching the model to reason about tools. You're teaching it to reproduce the surface form of effective tool use — the rhythm of think-act-observe, the template of the API call, the mapping from error message to recovery strategy. The pattern is learnable. The pattern is also transferable. And the pattern does not require 175 billion parameters to encode.

> Distillation says: you don't need the teacher's brain. You need the teacher's answers to enough questions that the student internalizes the shape of correct responses. Once the student has the shape, the teacher is overhead.

## The closed loop

This is where it gets terrifying.

A small model fine-tuned on 187,542 examples achieves a 77.55% pass rate. That means 22.45% of its tool calls still fail. But the model that scored 77.55% can now generate new instruction-solution pairs — its own successful trajectories. Those trajectories become training data for the next fine-tuning round. The model teaches itself.

Round 1: 187K human-curated examples → 77.55% pass rate.
Round 2: 187K human examples + N successful self-generated trajectories → ?
Round 3: human examples + even more self-generated trajectories → ?

Each round compresses the failure modes into the training distribution. The model learns not just what works but what fails and how to recover. The closed loop doesn't need more parameters. It needs more data — and the model is generating the data.

This is the dynamic that turns a 350M parameter curiosity into a production system that handles 90%+ of routine tool calls at a fraction of the cost. Each successful call produces a training example. Each training example makes the next call more likely to succeed. The loop tightens. The model improves without growing.

> The closed loop is the moat, not the model. Once the loop is running, switching to a different model is expensive not because the model is irreplaceable but because the loop has accumulated months of self-generated training data that the replacement model hasn't seen.

## Small model + right tools = terrifying capability

The AWS paper is not an isolated result. It's part of a pattern that has been building for two years.

AgentSymbiotic gets an 8B LLaMA to 48.5% on WebArena, approaching Claude-3.5's 52.1%. SCoRe uses a 7B student that matches a 72B teacher across 12 benchmarks — by having the student generate trajectories, the teacher correct only the earliest error, and RL close the gap. Every one of these results has the same shape: small model, targeted training, closed-loop refinement. The large model is the bootstrapping mechanism, not the destination.

The practical implication is uncomfortable for anyone who has bet on scale. If a 350M model can handle 77% of tool-calling tasks, and a 7B model can handle 90%, and those models cost 1/1000th as much to run as a frontier model — why is the frontier model handling routine API calls? The economics don't just favor the small model. They make the large model look like a rounding error.

> The future of agents is not a single omniscient model that does everything. It's a swarm of tiny, specialized models, each distilled to do one thing perfectly, coordinated by a router that costs almost nothing. The large model is the trainer. The small models are the workforce.

## What this means for how you build

The paper has a clear engineering implication: **stop fine-tuning large models for narrow tasks.** You are paying for parameters you don't need, training on knowledge the model already has, to produce capabilities that a smaller model could acquire faster and execute more reliably.

Start with the smallest model that can absorb the pattern. Fine-tune it on high-quality task-specific examples. Deploy it. Log its successes and failures. Feed the successes back into the training set. Let the closed loop do the work. If you need a larger model, use it only as a judge — to evaluate the small model's outputs, to correct its failures, to generate the next round of training data. The large model is the teacher. The small model is the worker. The loop is the product.

The AWS team proved that 350 million parameters is enough. Your tool-calling task is not harder than ToolBench. Your model is almost certainly too large.

---

*Paper: P. Jhandi, O. Kazi, S. Subramanian, N. Sendas — [Small Language Models for Efficient Agentic Tool Calling: Outperforming Large Models with Targeted Fine-tuning](https://arxiv.org/abs/2512.15943) (AAAI 2026 Workshop on Agentic AI Benchmarks and Applications for Enterprise Tasks).*
