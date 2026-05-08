---
title: On-device LLMs are a systems design problem
date: 2026-05-08
slug: on-device-llms-are-a-systems-design-problem
summary: The on-device future depends on more than one model choice; it depends on compression, acceleration, fallback policy, and deployment design.
tags: edge-llms, review, systems
---

There is a reason review papers are useful in fast-moving fields: they show how many moving parts a clean demo hides.

[On-Device Language Models: A Comprehensive Review](https://arxiv.org/abs/2409.00088) is valuable because it frames edge LLM deployment as a systems problem spanning compression, hardware acceleration, runtime strategy, and hybrid edge-cloud design.

That framing is worth stealing.

If a team says it is "doing on-device AI," the real question is whether it has a clear answer for:

- what runs locally,
- what falls back remotely,
- how quality and latency trade off,
- how the deployment will be debugged in the field.
