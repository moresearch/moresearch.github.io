---
title: "Hacker Laws for Agentic Software Engineering: The Law of Leaky Abstractions"
date: 2026-08-28
slug: hacker-laws-ase-leaky-abstractions
summary: "Law 12 of 12. Spolsky's Law of Leaky Abstractions says all non-trivial abstractions, to some degree, are leaky. The ASE key insight: every abstraction the agent lives on leaks — and the leak layer is where the harness must put the verifier, because the model cannot see the leak."
tags: hacker-laws, agentic-software-engineering, series, leaky-abstractions, spolsky, abstractions, verification
series: hacker-laws-for-agentic-software-engineering
---

**Law 12 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [The Bitter Lesson](https://blog.hackspree.com/#hacker-laws-ase-bitter-lesson).

## The Law

> All non-trivial abstractions, to some degree, are leaky. (Joel Spolsky, via [hacker-laws](https://github.com/dwmkerr/hacker-laws#the-law-of-leaky-abstractions))

## The Key Insight for Agentic Software Engineering

The agent lives on a stack of abstractions, and every one of them leaks: tool calling is an abstraction over the model's output format, and it leaks when the format changes; structured output is an abstraction over the model's ability to follow a schema, and it leaks when the schema is wrong; the context window is an abstraction over the model's attention, and it leaks in the middle ([Lost in the Middle](https://arxiv.org/abs/2307.03172)); "the model understood the task" is the deepest abstraction of all, and it leaks exactly when the model misunderstood. Spolsky's warning applies with a new twist: the *user of the abstraction is also the thing that leaks*. The agent cannot see the leak — it cannot tell the difference between a tool that failed and a tool that succeeded, which is precisely the [silent crash](https://blog.hackspree.com/#harnessing-agentic-ai-systems-schema-enforcement-self-correction) anti-pattern: the abstraction returns `""` or `"ok"` and the agent confidently proceeds on a false premise.

The consequence is that the harness is the leak-detection layer, and the verifier is the leak detector. Every abstraction boundary in the agent's stack needs a check on the far side of it: the schema boundary needs [enforcement with feedback](https://blog.hackspree.com/#harnessing-agentic-ai-systems-schema-enforcement-self-correction) (the model is still in the loop to fix what leaked); the tool boundary needs the [gatekeeper](https://blog.hackspree.com/#harnessing-agentic-ai-systems-static-intercepting-gatekeeper) and legible exit codes (the three channels exist so failures are visible); the outcome boundary needs the [evaluator with hands](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble) that uses the artifact instead of trusting the abstraction. The [agentic-first CLI](https://blog.hackspree.com/#agentic-first-cli-design) discipline is Leaky Abstractions in contract form: make the leak visible in the interface, because the agent will not see it otherwise.

The ASE reading of the Law of Leaky Abstractions: **agent abstractions leak — and the leak layer is where the harness must put the verifier, because the model cannot see the leak.** Each abstraction saves the agent from the underlying complexity until the day the complexity breaks through, and on that day the agent — unlike a human — has no prior experience of the underlying system to draw on. The harness is the one component that can be told about all the layers at once, so it is the one that must watch the seams.

## References

- dwmkerr. [hacker-laws — The Law of Leaky Abstractions](https://github.com/dwmkerr/hacker-laws#the-law-of-leaky-abstractions); Spolsky, *[The Law of Leaky Abstractions](https://www.joelonsoftware.com/2002/11/11/the-law-of-leaky-abstractions/)* (2002).
- Liu et al., [Lost in the Middle](https://arxiv.org/abs/2307.03172) — the context abstraction leaking in the middle.
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [schema enforcement & self-correction](https://blog.hackspree.com/#harnessing-agentic-ai-systems-schema-enforcement-self-correction), [static intercepting gatekeeper](https://blog.hackspree.com/#harnessing-agentic-ai-systems-static-intercepting-gatekeeper), [voting / consensual ensemble](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).
- [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) — the interface as the visible contract.
