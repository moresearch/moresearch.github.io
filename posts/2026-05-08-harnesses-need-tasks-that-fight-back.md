---
title: Agent harnesses need tasks that fight back
date: 2026-04-27
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

## Good tasks create resistance

A useful harness case should push back a little. It should require the system to gather information, make a choice, and survive the possibility that the first path is wrong. Without that resistance, it becomes too easy for a model to look capable by sounding capable.

This is what I like about GAIA as a reference point. It treats assistant work as something broader than isolated reasoning. The assistant has to operate, not just answer.

## Multi-step work is the real surface area

Most assistants become interesting only when they need to combine several behaviors:

- retrieve something,
- use a tool,
- inspect intermediate state,
- finish with a verifiable answer.

A harness that removes those transitions removes a big portion of the actual risk. Many failures live in the handoffs between steps, not in any single step on its own.

## Score what matters

The trick is to keep tasks hard enough to be honest without making them impossible to grade. That usually means designing cases with clear evidence of success while preserving enough messiness that shortcuts stop working.

That balance is worth the effort. If the tasks fight back a little, the results are much more believable. And believable results are what let teams trust an assistant enough to give it more room to work.
