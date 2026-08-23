---
title: "Hacker Laws for Agentic Software Engineering: Gall's Law"
date: 2026-08-27
slug: hacker-laws-ase-galls-law
summary: "Law 4 of 12. Gall's Law says a complex system that works has evolved from a simple system that worked. The ASE key insight: grow the agent from a working single loop — the god agent is the complex system designed from scratch, and it never works."
tags: hacker-laws, agentic-software-engineering, series, galls-law, complexity, evolution, god-agent
series: hacker-laws-for-agentic-software-engineering
---

**Law 4 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Amdahl's Law](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law) · Next: [Goodhart's Law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law).

## The Law

> A complex system that works is invariably found to have evolved from a simple system that worked. A complex system designed from scratch never works and cannot be patched up to make it work. You have to start over with a working simple system. (John Gall, via [hacker-laws](https://github.com/dwmkerr/hacker-laws#galls-law))

## The Key Insight for Agentic Software Engineering

Gall's Law is the anti-pattern catalogue in one paragraph. The [god agent](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) — the one massive agent with an immense prompt managing every phase of a workflow — is precisely the "complex system designed from scratch": every tool, every rule, every stage in one context window, assembled in one go and expected to work. It never does, and "cannot be patched up to make it work. You have to start over with a working simple system." The working simple system is the single loop — read, act, verify, reflect — and the evolution is the [harness pattern language](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems): from the [sequential pipeline](https://blog.hackspree.com/#harnessing-agentic-ai-systems-sequential-pipeline-routing) (linear, boring, working) you grow the [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) (delegation added when the context overflows), and from that the [ensembles and evaluators](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble) (verification added when the failures demand it). Each step is a small change to a system that already works.

The empirical record agrees. Anthropic's own long-running harness grew exactly this way: the earlier harness was an initializer plus a coding agent working one feature at a time; the frontier version added a planner and an evaluator, one at a time, each addressing a specific observed gap ([Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps)). The most successful guidance in the field is Gall's Law in product form: "the most successful implementations use simple, composable patterns rather than complex frameworks" ([Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents)). And the blog's own [self-improving loop](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) — read lessons, do work, reflect, write lessons — is the simple system from which everything else evolves.

The ASE reading of Gall's Law: **grow the agent from a working single loop; don't design the multi-agent system from scratch.** When a complex agentic architecture fails, the fix is not to add more scaffolding to it — it is to start over with the smallest loop that works and let the harness evolve. The system that works was found, not designed.

## References

- dwmkerr. [hacker-laws — Gall's Law](https://github.com/dwmkerr/hacker-laws#galls-law) and [Wikipedia](https://en.wikipedia.org/wiki/Gall%27s_law).
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) (the god agent), [sequential pipeline routing](https://blog.hackspree.com/#harnessing-agentic-ai-systems-sequential-pipeline-routing), [voting / consensual ensemble](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).
- [Harness Engineering: Best Practices for Reliable Agent Systems](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) — the simple self-improving loop.
- Anthropic. [Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps) and [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents).
