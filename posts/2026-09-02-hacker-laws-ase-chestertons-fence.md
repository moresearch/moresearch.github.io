---
title: "Hacker Laws for Agentic Software Engineering: Chesterton's Fence"
date: 2026-09-02
slug: hacker-laws-ase-chestertons-fence
summary: "Law 10 of 12. Chesterton's Fence says reforms should not be made until the reasoning behind the existing state of affairs is understood. The ASE key insight: the harness must make the agent find out why the code is there before letting it change — intent is a verification problem."
tags: hacker-laws, agentic-software-engineering, series, chestertons-fence, legacy, intent, refactoring
series: hacker-laws-for-agentic-software-engineering
---

**Law 10 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Parkinson's Law](https://blog.hackspree.com/#hacker-laws-ase-parkinsons-law) · Next: [The Bitter Lesson](https://blog.hackspree.com/#hacker-laws-ase-bitter-lesson).

## The Law

> Reforms should not be made until the reasoning behind the existing state of affairs is understood. ([hacker-laws](https://github.com/dwmkerr/hacker-laws#chestertons-fence))

## The Key Insight for Agentic Software Engineering

Chesterton's Fence is the law every agent violates on its first pass: it comes across a fence — a function that looks redundant, a workaround that looks wrong, a test that looks pointless — and removes it, because the agent has no memory of why the fence was built and no patience to find out. The unrequested feature that cost Fowler's team three days of investigation is the mild form ([Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)); the removed workaround that was load-bearing is the severe form. "Each line of a program was originally written by someone for some reason" — and in an agentic system, the someone is often an *earlier run of the same agent*, which makes the fence rule both more important and harder: the reasoning may exist only in a session log.

The law is why [legacy modernization](https://blog.hackspree.com/#factory-is-not-dead) is the clearest near-term value pool for agents, and why it is also the sharpest test of the harness: an agent told to "clean up this legacy code" is a fence-removal machine. The fix is not a better prompt ("understand before you change" is exactly the kind of instruction an agent will be told to ignore or will comply with shallowly); it is a harness rule — the verifier must check intent, not just correctness. Fowler's DSL idea is the constructive form: restrict the agent's change vocabulary until removing a fence requires explaining it first ([Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)).

The ASE reading of Chesterton's Fence: **the harness must make the agent find out why the code is there before letting it change — intent is a verification problem.** The fence rule cannot be a line in the system prompt; it has to be a gate in the pipeline: the agent must produce the fence's purpose as an artifact, and the verifier must check that artifact against the change. The mayor's answer to the man applies verbatim to the agent: "If you don't know its purpose, I certainly won't let you remove it. Go and find out the use of it, and then I may let you destroy it."

## References

- dwmkerr. [hacker-laws — Chesterton's Fence](https://github.com/dwmkerr/hacker-laws#chestertons-fence) and [Wikipedia](https://en.wikipedia.org/wiki/G._K._Chesterton#Chesterton's_fence).
- [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) — the unrequested feature; the DSL as the restricted change vocabulary.
- [The Factory Is Not Dead](https://blog.hackspree.com/#factory-is-not-dead) — legacy modernization as the agentic value pool.
- This blog's [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — [static intercepting gatekeeper](https://blog.hackspree.com/#harnessing-agentic-ai-systems-static-intercepting-gatekeeper), [verifiers are king](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc).
