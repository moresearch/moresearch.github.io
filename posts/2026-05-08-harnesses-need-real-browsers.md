---
draft: true
title: Harnesses Need Real Browsers
date: 2026-04-19
slug: harnesses-need-real-browsers-not-polite-demos
summary: Web agent evaluation gets more useful when the harness runs against real pages, real interaction flows, and explicit pass-fail criteria.
tags: harness, agents, browser
---

This post has been consolidated into [Harness Engineering: Best Practices for Reliable Agent Systems](#harness-engineering-best-practices-for-ai-agents). See that canonical post for the full, merged guidance.
If an agent claims it can use the web, the harness should make it prove it on the web.

That is why [WebVoyager: Building an End-to-End Web Agent with Large Multimodal Models](https://arxiv.org/abs/2401.13919) matters. The paper pushes evaluation into realistic browser tasks instead of stopping at tidy prompt-response examples.

The harness lesson is straightforward:

- use real interfaces,
- preserve the messy interaction sequence,
- score outcomes with concrete checks.

A fake DOM or hand-picked screenshot can still be useful for unit tests, but it should not be the final word on whether a browser agent actually works.

## Browser work is interaction work

Web tasks are not only about extracting the right answer from a page. They involve navigation, state, timing, and the strange little contingencies that happen when an interface is built for humans rather than benchmarks.

That is why polite demos can be misleading. A perfectly staged page or a frozen snapshot removes most of the difficulty that makes browser use interesting in the first place.

## Real pages expose real weaknesses

Once the harness runs against real interfaces, all the inconvenient details return:

- buttons move,
- forms span multiple steps,
- state has to persist across actions,
- success depends on the whole sequence, not one isolated click.

Those details are not noise. They are the task.

## Keep simpler fixtures, but do not stop there

There is still room for lightweight browser fixtures. They help with unit tests and targeted debugging. But a serious claim about browser competence should survive a more honest environment.

That is why work like WebVoyager is useful. It reminds teams that if a product pitch includes autonomous web interaction, the evaluation should include actual web interaction. Otherwise the harness is grading presentation quality, not browser ability.


Harness engineering is infrastructure engineering. The harness is the environment in which the agent operates. Designing a harness means deciding what the agent can see, what it can do, and how its actions are evaluated. The same design problem appears in any sandbox: the browser sandbox for JavaScript, the container sandbox for microservices, the test sandbox for CI/CD. The harness is the interface between the agent and the world. The interface determines what the agent can learn. The design of the interface is an engineering decision with consequences for everything the agent does downstream.
