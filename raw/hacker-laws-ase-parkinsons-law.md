---
title: "Hacker Laws for Agentic Software Engineering: Parkinson's Law"
date: 2026-09-01
slug: hacker-laws-ase-parkinsons-law
summary: "Law 9 of 12. Parkinson's Law says work expands so as to fill the time available for its completion. The ASE key insight: agent work expands to fill the budget — the context window and the token budget — so the budget is the discipline and the scope is a contract agreed before the work."
tags: hacker-laws, agentic-software-engineering, series, parkinsons-law, scope, budgets, context
series: hacker-laws-for-agentic-software-engineering
---

**Law 9 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Kernighan's Law](https://blog.hackspree.com/#hacker-laws-ase-kernighans-law) · Next: [Chesterton's Fence](https://blog.hackspree.com/#hacker-laws-ase-chestertons-fence).

## The Law

> Work expands so as to fill the time available for its completion. ([hacker-laws](https://github.com/dwmkerr/hacker-laws#parkinsons-law))

## The Key Insight for Agentic Software Engineering

Parkinson's Law described bureaucracies; agentic software engineering gives it a *meter*. An agent does not have a deadline, it has a context window and a token budget — and its work expands to fill both. Give an agent a vague task and a large window and it will produce a large result: more features it wasn't asked for, more refactoring it wasn't asked to do, more files touched than the task required. The unrequested feature that cost Fowler's team three days of investigation is Parkinson's Law with teeth ([Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)): the work expanded to fill the scope the agent inferred, and nobody had bounded the scope.

The law has a second, economic face. Work expands to fill the *budget* — and the budget is metered. This is the [token economics](https://blog.hackspree.com/#every-token-has-a-price-tag) argument in one sentence: an agent given an uncapped budget will spend it, because "the employee's rational strategy is to maximize usage — to tokenmaxx" — and the employee is now a loop that never gets tired. Parkinson's Law is why the [budget throttler](https://blog.hackspree.com/#harnessing-agentic-ai-systems-budget-throttler) and the [infinite execution vortex](https://blog.hackspree.com/#harnessing-agentic-ai-systems-budget-throttler) are the same pattern's two sides: without a ceiling, the work expands without bound; with a ceiling, the work expands to the ceiling and stops.

The fix is the discipline the pattern language keeps naming: **the scope is a contract agreed before the work** — the [sprint contract](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) that defines "done" before the agent starts, and the [tasks-that-fight-back](https://blog.hackspree.com/#harnessing-agentic-ai-systems-state-snapshot-rollback) principle that keeps tasks small enough to score. Parkinson's Law is not defeated by asking the agent to be brief; it is defeated by making the container small.

The ASE reading of Parkinson's Law: **agent work expands to fill the budget — the budget is the discipline, and the scope is a contract agreed before the work.** The deadline is not a time; it is a context window, a token ceiling, and a definition of done. Make the container small and the work will fit it.

## References

- dwmkerr. [hacker-laws — Parkinson's Law](https://github.com/dwmkerr/hacker-laws#parkinsons-law) and [Wikipedia](https://en.wikipedia.org/wiki/Parkinson%27s_law).
- [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag) — the meter; tokenmaxxing as the rational strategy.
- [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) — the three-day unrequested feature.
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [budget throttler](https://blog.hackspree.com/#harnessing-agentic-ai-systems-budget-throttler), [orchestrator-worker](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) (sprint contracts), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents).
