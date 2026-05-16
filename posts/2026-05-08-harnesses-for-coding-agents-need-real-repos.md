---
title: Coding Harnesses Need Real Repos
date: 2026-04-11
slug: coding-agent-harnesses-need-real-repositories
summary: Coding-agent quality becomes measurable when the harness uses actual repos, issues, and test outcomes instead of idealized toy prompts.
tags: harness, coding-agents, swe-bench
---

[SWE-bench: Can Language Models Resolve Real-World GitHub Issues?](https://arxiv.org/abs/2310.06770) changed the conversation because it moved coding evaluation closer to the work people actually do: read a real issue, edit a real repo, and satisfy a real test suite.

That is also the right pattern for harness design inside product teams.

A credible coding harness should prefer:

- messy repositories over perfect examples,
- failing tests over vague grading,
- issue-driven tasks over isolated snippets.

If the agent only looks good in a clean-room benchmark, the harness is flattering the system instead of measuring it.

## Real repositories force real tradeoffs

Toy prompts are attractive because they are easy to score and easy to explain. They are also easy to overfit. A coding agent can look competent when the entire task fits neatly inside a paragraph and a single file.

Real repositories are different. They introduce layout confusion, hidden assumptions, partial context, and tests that fail for reasons the agent has to discover. That is much closer to the work engineers actually delegate.

## Why issue-driven evaluation matters

The issue format is valuable because it mirrors a normal entry point for engineering work. A task begins with an imperfect description, not with a fully specified solution path. The agent has to read, inspect, and decide what matters.

That is the sort of friction a harness should preserve. It reveals whether the system can navigate ambiguity without immediately collapsing into guesswork.

## Tests are better than vibes

A strong coding harness needs explicit outcomes, and failing tests are one of the cleanest ways to get them. They are not perfect, but they are far better than a human observer deciding whether an output felt plausible.

The broader lesson from SWE-bench is not only about benchmarking. It is about honesty. If you want to know whether a coding agent can help on real work, give it real repositories and real consequences. Everything else is too easy to stage-manage.
