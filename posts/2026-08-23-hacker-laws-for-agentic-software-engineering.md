---
title: "Hacker Laws for Agentic Software Engineering"
date: 2026-08-23
slug: hacker-laws-for-agentic-software-engineering
summary: "The index of a series that takes twelve laws from dwmkerr's hacker-laws catalog and asks what each means when the engineer is an agentic AI system. The law does not change; the subject does — Brooks was about people, and in agentic software engineering 'human resources' can be spawned in milliseconds, but the reasoning survives the change of subject. Each law gets its own post, focused on one key insight for agentic software engineering."
tags: hacker-laws, agentic-software-engineering, series, patterns, agents, laws, principles
series: hacker-laws-for-agentic-software-engineering
---

[dwmkerr/hacker-laws](https://github.com/dwmkerr/hacker-laws) is the best-known catalog of the laws, theorems, and principles that software engineers cite to explain why things keep going wrong. This series takes twelve of them and asks a question the catalog was never written to answer: **what does each law mean when the engineer is an agentic AI system?**

The law does not change; the subject does. Brooks' Law was about people — and in agentic software engineering, "human resources" can be spawned in milliseconds. Amdahl's Law was about processors — and now the processors are agents. Goodhart's Law was about KPI-gaming by employees — and now the optimizer is a model that will game the metric at machine speed. The reasoning survives the change of subject; that is exactly why the laws are worth re-reading for the agentic era, and why this blog's own catalog — the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) pattern language — keeps landing on the same conclusions the laws reached decades ago.

Each post in this series is one law, one commit, one push, and one key insight for agentic software engineering — the insight is the whole point, and the law is the evidence.

| # | Law | The ASE key insight |
|---|---|---|
| 1 | [Conway's Law](https://blog.hackspree.com/#hacker-laws-ase-conways-law) | The software will mirror the topology of the agents that made it — design the topology (the harness), not just the prompt. |
| 2 | [Brooks' Law](https://blog.hackspree.com/#hacker-laws-ase-brooks-law) | Nine agents can't make a feature in one day — the handoff is the ramp-up, and the serial fraction sets the ceiling. |
| 3 | [Amdahl's Law](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law) | Measure the serial fraction of your agent pipeline — the verifier is usually it, and no number of workers beats the ceiling it sets. |
| 4 | [Gall's Law](https://blog.hackspree.com/#hacker-laws-ase-galls-law) | Grow the agent from a working single loop; don't design the multi-agent system from scratch. |
| 5 | [Goodhart's Law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law) | For agents, the measure becomes the training target — choose evals as if the agent will learn to game them, because it will. |
| 6 | [Hyrum's Law](https://blog.hackspree.com/#hacker-laws-ase-hyrums-law) | The agent will depend on every observable behaviour you didn't promise — for agents, the implicit interface IS the contract. |
| 7 | [Hofstadter's Law](https://blog.hackspree.com/#hacker-laws-ase-hofstadters-law) | An agent task always takes longer than you expect, recursively — so the ceiling is a system property, not an estimate. |
| 8 | [Kernighan's Law](https://blog.hackspree.com/#hacker-laws-ase-kernighans-law) | If the agent writes clever code, the system must debug it — keep agent output boring and make the verifier the smarter half. |
| 9 | [Parkinson's Law](https://blog.hackspree.com/#hacker-laws-ase-parkinsons-law) | Agent work expands to fill the budget — the budget is the discipline, and the scope is a contract agreed before the work. |
| 10 | [Chesterton's Fence](https://blog.hackspree.com/#hacker-laws-ase-chestertons-fence) | The harness must make the agent find out why the code is there before letting it change — intent is a verification problem. |
| 11 | [The Bitter Lesson](https://blog.hackspree.com/#hacker-laws-ase-bitter-lesson) | The loop that leverages computation beats the hand-crafted prompt — and the agent will apply the same lesson to your harness. |
| 12 | [The Law of Leaky Abstractions](https://blog.hackspree.com/#hacker-laws-ase-leaky-abstractions) | Agent abstractions leak — and the leak layer is where the harness must put the verifier, because the model cannot see the leak. |

## The series

1. [Conway's Law](https://blog.hackspree.com/#hacker-laws-ase-conways-law) — the org chart is now humans, agents, and the harness that wires them.
2. [Brooks' Law](https://blog.hackspree.com/#hacker-laws-ase-brooks-law) — adding agents to a late task makes it later.
3. [Amdahl's Law](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law) — agent parallelism is bounded by the serial fraction.
4. [Gall's Law](https://blog.hackspree.com/#hacker-laws-ase-galls-law) — complex agent systems evolve from simple loops.
5. [Goodhart's Law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law) — the eval is the curriculum; the metric is the target.
6. [Hyrum's Law](https://blog.hackspree.com/#hacker-laws-ase-hyrums-law) — agents depend on every observable behaviour.
7. [Hofstadter's Law](https://blog.hackspree.com/#hacker-laws-ase-hofstadters-law) — agent timelines are recursive estimates.
8. [Kernighan's Law](https://blog.hackspree.com/#hacker-laws-ase-kernighans-law) — debugging agent output is the bottleneck.
9. [Parkinson's Law](https://blog.hackspree.com/#hacker-laws-ase-parkinsons-law) — agent work expands to fill the budget.
10. [Chesterton's Fence](https://blog.hackspree.com/#hacker-laws-ase-chestertons-fence) — understand before the agent changes anything.
11. [The Bitter Lesson](https://blog.hackspree.com/#hacker-laws-ase-bitter-lesson) — compute and the loop beat the hand-crafted prompt.
12. [The Law of Leaky Abstractions](https://blog.hackspree.com/#hacker-laws-ase-leaky-abstractions) — every agent abstraction leaks; the verifier lives at the leak.

## References

- dwmkerr. [hacker-laws](https://github.com/dwmkerr/hacker-laws) — the source catalog; each law post links to its own entry and the primary source.
- This blog's [Harnessing Agentic AI Systems: A Pattern Language](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) — the pattern language the laws keep landing on.
- [Loop Engineering is what the NATO conference asked for in 1968](https://blog.hackspree.com/#loop-engineering) — the loop as the unit of design, which most of these laws end up constraining.
