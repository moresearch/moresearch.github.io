---
draft: true
title: Harnesses Should Observe the OS
date: 2026-05-03
slug: good-harnesses-watch-the-whole-operating-system
summary: Desktop and multimodal agents need execution harnesses that see the same operating system complexity the user sees.
tags: harness, osworld, multimodal
---

This post has been consolidated into [Harness Engineering: Best Practices for Reliable Agent Systems](#harness-engineering-best-practices-for-ai-agents). See that canonical post for the full, merged guidance.
The right environment can make an average agent look impressive or make a strong agent fail honestly.

[OSWorld: Benchmarking Multimodal Agents for Open-Ended Tasks in Real Computer Environments](https://arxiv.org/abs/2404.07972) is valuable because it evaluates agents in real desktop settings rather than collapsing everything into a synthetic task wrapper.

That is the bar harness engineers should aim for. If your agent acts inside an operating system, your evaluation should capture:

- window state,
- clipboard and file effects,
- long action sequences,
- recovery after mistakes.

The more your harness reflects the real environment, the less likely you are to ship a demo that falls apart the moment a user clicks somewhere unexpected.

## The operating system is part of the task

Desktop and multimodal agents do not act in a vacuum. They act in an environment with windows, focus changes, files, delays, and all the messy state that ordinary users barely notice until something goes wrong.

A harness that trims that away may be easier to manage, but it also stops measuring the thing the user will actually experience.

## State is where many failures hide

Operating system tasks are difficult partly because the agent has to maintain awareness across a longer chain of effects. A single wrong click can change focus. A clipboard action can overwrite an assumption. A file operation can succeed but still leave the system in the wrong place for the next step.

That is why it helps to capture the broader environment, not just the final answer.

## Honest environments create honest confidence

OSWorld is useful because it pushes evaluation toward the real surface area of desktop interaction. The lesson is not that every harness must copy the benchmark exactly. The lesson is that environment fidelity matters.

If the product is supposed to work across the operating system, the harness should watch the operating system too. Otherwise the team risks learning the wrong lesson from a passing score.


Harness engineering is infrastructure engineering. The harness is the environment in which the agent operates. Designing a harness means deciding what the agent can see, what it can do, and how its actions are evaluated. The same design problem appears in any sandbox: the browser sandbox for JavaScript, the container sandbox for microservices, the test sandbox for CI/CD. The harness is the interface between the agent and the world. The interface determines what the agent can learn. The design of the interface is an engineering decision with consequences for everything the agent does downstream.


> A task that does not resist teaches the agent that resistance does not exist. The real world resists. The agent must learn to overcome resistance. The resistance is the teacher.
