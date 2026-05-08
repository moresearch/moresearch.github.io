---
title: Coding agent harnesses need real repositories
date: 2026-05-08
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
