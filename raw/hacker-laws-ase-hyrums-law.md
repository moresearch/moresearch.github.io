---
title: "Hacker Laws for Agentic Software Engineering: Hyrum's Law"
date: 2026-08-29
slug: hacker-laws-ase-hyrums-law
summary: "Law 6 of 12. Hyrum's Law says with enough users, all observable behaviours of a system will be depended on by somebody. The ASE key insight: agents are the most thorough users of your interfaces — they depend on every observable behaviour you didn't promise, so the implicit interface IS the contract."
tags: hacker-laws, agentic-software-engineering, series, hyrums-law, implicit-interfaces, cli, contracts
series: hacker-laws-for-agentic-software-engineering
---

**Law 6 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Goodhart's Law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law) · Next: [Hofstadter's Law](https://blog.hackspree.com/#hacker-laws-ase-hofstadters-law).

## The Law

> With a sufficient number of users of an API, it does not matter what you promise in the contract: all observable behaviours of your system will be depended on by somebody. (Hyrum Wright, via [hacker-laws](https://github.com/dwmkerr/hacker-laws#hyrums-law-the-law-of-implicit-interfaces))

## The Key Insight for Agentic Software Engineering

Hyrum's Law assumed human users who squint, scroll, and improvise. The agent is a user who never blinks: it reads every byte of help text, every line of output, every exit code, every timestamp, every ordering — and it depends on all of it, because it has no tolerance for ambiguity and no memory of what "should have worked." The [agentic-first CLI](https://blog.hackspree.com/#agentic-first-cli-design) discipline is Hyrum's Law taken as a design contract: "a lie in help text is the most expensive bug an agentic CLI can have," and "no timestamps unless asked" is not a preference but a dependency hazard — the agent will start parsing the timestamp and break when it changes format.

The law sharpens in two directions. First, the *model-facing* interfaces: the tool schemas, the `--json` outputs, the structured contracts the agent reads — every observable behaviour, including the ones you did not promise, becomes part of the de facto API. The [type-graph mirror](https://blog.hackspree.com/#deepseek-harness) exists precisely because spec drift is a Hyrum's Law failure: the tool catalogue that drifts from the implementation is an implicit interface that somebody — some agent — will depend on. Second, the *model itself* becomes an interface others depend on: once agents are built on a model's observable behaviours (its tool-calling format, its refusal patterns, its output ordering), those behaviours are frozen by the ecosystem the way an API is frozen by its users — which is why compatibility-breaking changes in a harness's session format are so expensive ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

The ASE reading of Hyrum's Law: **the agent will depend on every observable behaviour you didn't promise — for agents, the implicit interface IS the contract.** The defense is to make the promised contract exhaustive: stable, versioned, additive `--json` schemas; deterministic ordering; explicit exit-code semantics; and honest `--help`. You cannot stop agents from depending on your behaviour, but you can decide which behaviour they depend on. Hyrum's Law does not say the contract is meaningless — it says the contract must cover everything observable.

## References

- dwmkerr. [hacker-laws — Hyrum's Law](https://github.com/dwmkerr/hacker-laws#hyrums-law-the-law-of-implicit-interfaces) and [Hyrum's Law](https://www.hyrumslaw.com/).
- [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) — the agent as the user who never blinks; the contract as the architecture.
- [DeepSeek Harness: Everything Is a Plugin](https://blog.hackspree.com/#deepseek-harness) — the type-graph mirror; session-format compatibility as a Hyrum constraint.
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [schema enforcement & self-correction](https://blog.hackspree.com/#harnessing-agentic-ai-systems-schema-enforcement-self-correction).
