---
title: Good harnesses watch the whole operating system
date: 2026-05-08
slug: good-harnesses-watch-the-whole-operating-system
summary: Desktop and multimodal agents need execution harnesses that see the same operating system complexity the user sees.
tags: harness, osworld, multimodal
---

The right environment can make an average agent look impressive or make a strong agent fail honestly.

[OSWorld: Benchmarking Multimodal Agents for Open-Ended Tasks in Real Computer Environments](https://arxiv.org/abs/2404.07972) is valuable because it evaluates agents in real desktop settings rather than collapsing everything into a synthetic task wrapper.

That is the bar harness engineers should aim for. If your agent acts inside an operating system, your evaluation should capture:

- window state,
- clipboard and file effects,
- long action sequences,
- recovery after mistakes.

The more your harness reflects the real environment, the less likely you are to ship a demo that falls apart the moment a user clicks somewhere unexpected.
