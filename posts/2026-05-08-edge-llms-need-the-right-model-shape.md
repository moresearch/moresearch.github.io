---
title: Edge LLMs live or die by model shape
date: 2026-05-08
slug: edge-llms-live-or-die-by-model-shape
summary: On-device language models are constrained less by hype and more by architecture choices that survive mobile memory, latency, and thermal limits.
tags: edge-llms, mobile, architecture
---

The promise of on-device inference is easy to say and hard to ship.

[MobileLLM: Optimizing Sub-billion Parameter Language Models for On-Device Use Cases](https://arxiv.org/abs/2402.14905) is a useful paper because it focuses on the architecture details that matter at the small end: depth, width, and parameter efficiency.

The engineering takeaway is that edge deployment is not only about compressing a big cloud model. Sometimes the winning move is to start from a model shape that was designed for the edge in the first place.

For product teams, that means evaluating model architecture and serving strategy together, not in separate meetings.
