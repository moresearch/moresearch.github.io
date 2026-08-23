---
title: "Hacker Laws for Agentic Software Engineering: Brooks' Law"
date: 2026-08-25
slug: hacker-laws-ase-brooks-law
summary: "Law 2 of 12. Brooks' Law says adding human resources to a late project makes it later. The ASE key insight: adding agents to a late task makes it later — the handoff is the ramp-up, the coordination is the overhead, and the serial fraction sets the ceiling."
tags: hacker-laws, agentic-software-engineering, series, brooks-law, parallelism, delegation, mythical-man-month
series: hacker-laws-for-agentic-software-engineering
---

**Law 2 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Conway's Law](https://blog.hackspree.com/#hacker-laws-ase-conways-law) · Next: [Amdahl's Law](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law).

## The Law

> Adding human resources to a late software development project makes it later. ([hacker-laws](https://github.com/dwmkerr/hacker-laws#brooks-law))

## The Key Insight for Agentic Software Engineering

The first temptation when a task is late is to spawn more agents — the "nine women can't make a baby in one month" intuition gets buried under the fact that an agent costs nothing to instantiate. But Brooks' reasoning survives the change of subject completely: the ramp-up time becomes the **context handoff** (the new agent must learn the task, the repo, and the state — exactly the [goldfish amnesia](https://blog.hackspree.com/#harnessing-agentic-ai-systems-state-snapshot-rollback) risk of delegation); the communication overhead becomes the **coordination** between agents (shared state, delegation contracts, merge conflicts — the [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) single point of failure); and many tasks are not divisible, because the serial reasoning and verification fraction cannot be split.

The cost structure is what changed, and it is the trap. Spawning an agent is nearly free; **integrating its output is not**. A human's ramp-up is measured in days and the org pays it once; an agent's ramp-up is measured in tokens and context, but it is paid on every delegation, and a late task invites delegating more, which multiplies the handoffs precisely when the serial path is already the bottleneck. This is the [Brooks → Amdahl](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law) pair: adding workers only helps the parallelisable fraction, and the late task is late because of the serial fraction.

The ASE reading of Brooks' Law: **nine agents can't make a feature in one day — the handoff is the ramp-up, and the serial fraction sets the ceiling.** The fix is not to spawn more workers; it is to shrink the handoff (structured artifacts, sprint contracts that define "done" before the work — the [frontier contract pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker)) and to make the serial fraction explicit and measured, so the decision to add an agent is made against the ceiling it cannot move.

## References

- dwmkerr. [hacker-laws — Brooks' Law](https://github.com/dwmkerr/hacker-laws#brooks-law) and [Wikipedia](https://en.wikipedia.org/wiki/Brooks%27s_law); Brooks, *The Mythical Man-Month*.
- [Amdahl's Law](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law) — the serial fraction is the ceiling.
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker), [state snapshot & rollback](https://blog.hackspree.com/#harnessing-agentic-ai-systems-state-snapshot-rollback) (handoff state), [asynchronous tool worker queue](https://blog.hackspree.com/#harnessing-agentic-ai-systems-async-tool-worker-queue).
