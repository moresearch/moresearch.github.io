---
title: Correctness First
date: 2026-07-20
slug: correctness-first
summary: Linux won the world because a lawsuit froze BSD at the wrong moment. OpenBSD is what the world lost — and what agent builders most need to understand.
tags: openbsd, linux, correctness, security, agents, software-engineering
---

In 1991, Linus Torvalds wanted a Unix-like system for his 386. BSD Net/2 had been released, but he didn't know about it. He wrote his own kernel. The next year, AT&T sued BSDi, alleging BSD contained proprietary Unix code. BSD development froze for two years. Linux, written from scratch with no legal baggage, absorbed the energy. By the time the lawsuit settled in 1994, Linux had won.

Linus later said: "If 386BSD had been available when I started on Linux, Linux would probably never have happened."

The world runs on Linux not because it was better designed, but because it was available. The cathedral lost to the bazaar by accident.

![The BSD Daemon — mascot of BSD since 1976, drawn by John Lasseter before Pixar](/images/beastie.png)

## Two architectures, two theories of quality

BSD is an operating system. One source tree. One team. The kernel, libc, core utilities, daemons, manual pages — written, reviewed, and shipped by the same people. When a kernel interface changes, every userland tool that depends on it is updated in the same commit. Configuration syntax is consistent because the same hands that wrote the daemon wrote its parser.

Linux is a kernel assembled into an operating system by distributions. The kernel comes from Linus. The C library from GNU. Userspace from a hundred independent projects. Each has its own maintainers, release schedule, coding style, and idea of what "good" means. The distribution's job is integration, not design. The result works. It is not clean. Everyone who has debugged a Linux system at 3am knows the feeling of crossing a component boundary and discovering the assumptions changed.

> Linux is dirty in the way a city is dirty — it works, it's full of life, but nobody designed it from scratch. BSD is clean in the way a well-designed building is clean — the structure is visible, the materials are consistent, the wiring is labeled.

## OpenBSD: correctness as the only feature

![Puffy, the OpenBSD mascot — a pufferfish. Defensive, uncompromising, hard to swallow.](/images/puffy.gif)

If BSD represents the cathedral, OpenBSD is the cathedral with the strictest building code. Theo de Raadt forked it from NetBSD in 1995 with an uncompromising thesis: correctness over features, every time.

> "We are non-stop trying to find ways across our entire source tree that small little programmer errors result in problems. At some point, we have to start asking ourselves whether features are the thing, or whether quality is the issue." — Theo de Raadt

The project has audited its entire source tree — millions of lines — line by line. Multiple times. Not primarily with tools. With humans reading code, asking "what happens when this fails?" The result: an operating system with only two remote holes in the default install across nearly three decades.

This is not a security achievement. It is a correctness achievement that makes security possible.

OpenBSD's innovations read like a list of things every other OS later adopted: W^X memory (2003, now universal), `pledge(2)` (process-level system call restrictions), `unveil(2)` (filesystem sandboxing), LibreSSL (forked OpenSSL, removed half the code, broke nothing). Each was born from asking: *how do we make the wrong thing impossible, not just harder?*

> Most projects respond to security by adding layers — a firewall, a sandbox, a scanner. OpenBSD responds by removing the bug. The difference is the difference between a house with a reinforced door and a house with no structural flaws.

## The collision with the agent era

We are now deploying software that writes software — agents that generate, edit, and ship code at scale. These agents do not understand correctness. They produce tokens statistically likely to satisfy a prompt. They don't know that this return value needs checking, that this buffer needs bounds, that this error path leaks a descriptor. They pattern-match from training data that, statistically, also doesn't know these things.

> Generated code is code nobody has read. Tests cover what you thought to test. Vulnerabilities live in what you didn't think of.

The agent era promises speed. OpenBSD proves that correctness at speed is an oxymoron. Shipping code faster than anyone can review it is not automation — it's an attack surface factory.

The resolution is not to reject agents. It is to demand of agent output what OpenBSD demands of human output: adversarial review, minimal diffs, restricted capabilities, and a culture that refuses to ship code that hasn't been read.

## What agent builders should steal

**Audit the diff, not the demo.** OpenBSD doesn't audit the running kernel. It audits the source. Evaluate what the agent changed, not just whether tests pass. The diff is where bugs live.

**Restrict by default.** OpenBSD ships with almost nothing running. Agent platforms should ship with almost no capabilities. Every new tool is attack surface — for bugs, prompt injection, unintended behavior. An agent that can rewrite any file has maximum blast radius. An agent restricted to specific files and specific side effects can fail without destroying things.

**The harness is your code review.** The agent has no concept of correctness. The harness has to. If your evaluation scores "did it compile?" and "did the test pass?" and "did the reviewer LLM nod?" — you are not asking enough. Write criteria as specific as an OpenBSD code review: "this error must propagate," "this resource must be freed."

**Remove generated code aggressively.** A 500-line generated function that works is worse than a 30-line one that works. Every line the agent wrote is a line someone needs to read. Prefer agents that produce minimal diffs.

**Culture scales. Checklists don't.** OpenBSD's real innovation is cultural: correctness is everyone's job, speed is never an excuse, shipped bugs are taken personally. Agent teams need the same culture. If your team treats agent output as "good enough if the tests pass," you will ship vulnerabilities at agent speed — faster than anyone can review.

---

OpenBSD has been building secure software for over two decades with a team that fits in a single room. They do it by being slower, more careful, and more uncompromising than almost any other project.

> Reliability was always the prerequisite for security. OpenBSD proved it with decades of results.

The agent era will prove it again — either by adopting that lesson, or by suffering what happens when you don't.
