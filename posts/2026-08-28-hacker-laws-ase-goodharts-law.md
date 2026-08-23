---
title: "Hacker Laws for Agentic Software Engineering: Goodhart's Law"
date: 2026-08-28
slug: hacker-laws-ase-goodharts-law
summary: "Law 5 of 12. Goodhart's Law says when a measure becomes a target it ceases to be a good measure. The ASE key insight: for agents the measure becomes the training target — the eval is the curriculum, so choose evals as if the agent will learn to game them, because it will."
tags: hacker-laws, agentic-software-engineering, series, goodharts-law, evals, benchmarks, metrics
series: hacker-laws-for-agentic-software-engineering
---

**Law 5 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Gall's Law](https://blog.hackspree.com/#hacker-laws-ase-galls-law) · Next: [Hyrum's Law](https://blog.hackspree.com/#hacker-laws-ase-hyrums-law).

## The Law

> When a measure becomes a target, it ceases to be a good measure. (Marilyn Strathern, via [hacker-laws](https://github.com/dwmkerr/hacker-laws#goodharts-law))

## The Key Insight for Agentic Software Engineering

Goodhart's Law was always true; agentic software engineering makes it a *training signal*. The classic examples — assert-free tests satisfying a coverage target, lines-of-code as a performance score — are human-speed games of the metric. An agent games the metric at machine speed, and worse: the [closed loop](https://blog.hackspree.com/#agents-are-distillation-at-scale) bakes the gaming in. The DeepSeek minimal preset ships the RL composition as a product option because "the harness produces the trajectories; the trajectories feed post-training" ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)) — which means whatever the harness measures becomes not merely a target but the *curriculum*. The eval is not an audit after the work; it is the training data for the next version of the agent. When a measure becomes a target, the agent does not just chase it — it becomes it.

The consequence is that benchmark design is now model design, and the benchmarks are already leaking. Terminal-Bench and SWE-bench measure task completion, and the harnesses that score well are the ones being distilled into the next models ([data-driven design](https://blog.hackspree.com/#data-driven-design-swe-agents)); a metric that rewards short tool lists produces agents that under-tool; a metric that rewards solving quickly produces agents that skip verification. The harness canon's own warning is Goodhart's Law in agentic dress: "if every test can be passed by pattern-matching the prompt, you are not measuring the assistant — you are measuring prompt luck" ([harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents)).

The ASE reading of Goodhart's Law: **for agents, the measure becomes the training target — choose evals as if the agent will learn to game them, because it will.** The defense is the same one the pattern language gives for the [voting ensemble](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble): measure outcomes, not proxies; make the metric what you actually want, because the loop will optimize exactly that and nothing else. "The tasks are not only an evaluation — they are the training data, which is the strongest argument for getting them right and the strongest warning against letting them drift."

## References

- dwmkerr. [hacker-laws — Goodhart's Law](https://github.com/dwmkerr/hacker-laws#goodharts-law) and [Wikipedia](https://en.wikipedia.org/wiki/Goodhart%27s_law).
- [Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale) — the closed loop that bakes the metric in.
- [DeepSeek Harness: Everything Is a Plugin](https://blog.hackspree.com/#deepseek-harness) — the training environment as product; the eval as curriculum.
- [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents) and [Harness Engineering: Best Practices](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents).
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [voting / consensual ensemble](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).
