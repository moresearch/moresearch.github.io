---
title: Harnesses need real browsers, not polite demos
date: 2026-05-08
slug: harnesses-need-real-browsers-not-polite-demos
summary: Web agent evaluation gets more useful when the harness runs against real pages, real interaction flows, and explicit pass-fail criteria.
tags: harness, agents, browser
---

If an agent claims it can use the web, the harness should make it prove it on the web.

That is why [WebVoyager: Building an End-to-End Web Agent with Large Multimodal Models](https://arxiv.org/abs/2401.13919) matters. The paper pushes evaluation into realistic browser tasks instead of stopping at tidy prompt-response examples.

The harness lesson is straightforward:

- use real interfaces,
- preserve the messy interaction sequence,
- score outcomes with concrete checks.

A fake DOM or hand-picked screenshot can still be useful for unit tests, but it should not be the final word on whether a browser agent actually works.
