---
title: "Hacker Laws for Agentic Software Engineering: Kernighan's Law"
date: 2026-08-31
slug: hacker-laws-ase-kernighans-law
summary: "Law 8 of 12. Kernighan's Law says debugging is twice as hard as writing the code in the first place. The ASE key insight: if the agent writes clever code, the system must debug it — keep agent output boring and make the verifier the smarter half."
tags: hacker-laws, agentic-software-engineering, series, kernighans-law, debugging, verification, simplicity
series: hacker-laws-for-agentic-software-engineering
---

**Law 8 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Hofstadter's Law](https://blog.hackspree.com/#hacker-laws-ase-hofstadters-law) · Next: [Parkinson's Law](https://blog.hackspree.com/#hacker-laws-ase-parkinsons-law).

## The Law

> Debugging is twice as hard as writing the code in the first place. Therefore, if you write the code as cleverly as possible, you are, by definition, not smart enough to debug it. (Brian Kernighan, via [hacker-laws](https://github.com/dwmkerr/hacker-laws#kernighans-law))

## The Key Insight for Agentic Software Engineering

Kernighan's Law was written for human programmers; agentic software engineering makes it *worse*, because the writer and the debugger are different systems. The model writes the code; the model cannot debug it — an agent asked to evaluate its own work "tend[s] to respond by confidently praising the work" ([Anthropic's harness](https://www.anthropic.com/engineering/harness-design-long-running-apps)) — so the debugging burden falls on the system: the verifiers, the tests, the reviewers, and eventually the humans. This is the same finding Fowler's retreat made the headline of the agentic era: "code generation is no longer the bottleneck — verification is" ([Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)). Kernighan's Law names why: the generation half is now nearly free, so the debugging half — always twice as hard — is where all the cost went.

The law's prescription survives verbatim: *don't let the agent write clever code*. If the agent produces the most intricate solution it can, the system — which must debug it — is "by definition, not smart enough." The harness-level translation is the whole [boring-output discipline](https://blog.hackspree.com/#harnessing-agentic-ai-systems-schema-enforcement-self-correction): schema-enforced, conventional, simple output that the verifier can actually check; and the verifier must be built to be the smarter half — the [evaluator with hands](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble) that uses the artifact instead of reading it, the [ensembles](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble) that cross-check instead of self-praise. The 79% self-review datum from the mob post is Kernighan's Law in review form: most agent PRs were reviewed by the same developer who prompted the agent — the writer debugging its own work, which the law says is impossible ([mob programming remastered](https://blog.hackspree.com/#mob-programming-reimagined)).

The ASE reading of Kernighan's Law: **if the agent writes clever code, the system must debug it — keep agent output boring and make the verifier the smarter half.** The model generates; the harness verifies; and the division is structural, because the writer can never be trusted to debug what it wrote.

## References

- dwmkerr. [hacker-laws — Kernighan's Law](https://github.com/dwmkerr/hacker-laws#kernighans-law); Kernighan & Plauger, *The Elements of Programming Style*.
- [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) and [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc).
- [Zuill's Mob Programming, Remastered](https://blog.hackspree.com/#mob-programming-reimagined) — the self-review problem.
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [schema enforcement & self-correction](https://blog.hackspree.com/#harnessing-agentic-ai-systems-schema-enforcement-self-correction), [voting / consensual ensemble](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).
