---
title: Agent harnesses need tasks that fight back
date: 2026-05-08
slug: agent-harnesses-need-tasks-that-fight-back
summary: Strong assistants are easier to trust when they are tested on tasks that require tools, retrieval, and real-world messiness.
tags: harness, gaia, evaluation
---

[GAIA: a benchmark for General AI Assistants](https://arxiv.org/abs/2311.12983) is useful because it resists the temptation to make evaluation too neat. It mixes reasoning, tool use, information gathering, and multimodal work in a way that feels closer to assistant reality.

That is a good model for internal harnesses too.

The harness should ask the system to do tasks that are:

- small enough to score,
- rich enough to require multiple steps,
- messy enough that shortcuts stop working.

If every test can be passed by pattern-matching the prompt, you are not measuring an assistant. You are measuring prompt luck.
