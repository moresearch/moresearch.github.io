---
title: Durable Daemons — An Event-Driven Choreography Pattern for Persistent AI Agents
date: 2026-07-25
slug: durable-daemons
summary: Durable daemons are an event-driven choreography pattern for AI agents. Each daemon persists, remembers, and acts autonomously. Together they coordinate through shared state — no central orchestrator, no message bus. This is the Unix inheritance: Maxwell's demon to Beastie to Go.
tags: daemons, durable-daemons-pattern, event-driven, choreography, bsd, unix, go, systems
series: durable-daemons-pattern
---

Imagine a being that sits between two chambers. It observes molecules. When a fast one approaches the gate, it opens. When a slow one approaches, it closes. It uses no energy. Only information. It sorts order from chaos without touching either. This is Maxwell's demon. This is a daemon.

> A daemon observes. A daemon decides. A daemon acts. It uses information, not energy. It runs forever.

In 1976, Stallman's ITS at MIT had DAEMON — a background process that watched for new files, woke up, and acted. John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to. Three eras. Three kinds of daemon. One architectural pattern.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988. The trident is fork(2). The tennis shoes are unexplained.](/images/bsd-daemon-medium.gif)

Imagine a service. It receives a request. It produces a response. It forgets. Now imagine a daemon. It maintains a queue. A schedule. A watch on a directory. It acts when its internal state crosses a threshold — not when it is asked, but when conditions are met. A service is reactive. A daemon is autonomous. And daemons compose. One daemon writes state. Another daemon observes it. A third reacts. No orchestrator. No message bus. Just shared state and independent triggers. This is choreography. `cron` has watched the clock since 1975: check time, decide, fork and exec, repeat. No one asked `cron` to do anything. It asked itself. Now imagine a hundred `cron`s, each watching different conditions, each acting on shared state. That is the durable daemons pattern.

> A service answers and forgets. A daemon maintains and acts. Durable daemons coordinate without a coordinator.

This is the Unix inheritance. *Daemon* is Greek δαίμων — guardian spirit, intermediary between mortals and gods. The pitchfork is `fork(2)`, the system call that spawns a child process to do the daemon's work. BSD gave daemons conceptual integrity — one source tree, one team, one design voice. Daemons are first-class citizens of that design. Beastie, drawn by a comic artist paid for cracking a wall safe, redrawn by a Pixar founder, became the mascot. The tennis shoes are unexplained.

Go is the natural language for this inheritance. Compile to a single static binary. No runtime to install. No `libc` dance. Goroutines give you lightweight concurrency — one per workflow, cheap as a function call. The standard library handles signals, file descriptors, and process management. Deploy a daemon by copying a file. Compose daemons by pointing them at the same Postgres instance. Each daemon is a process. The composition is the shared state. Go was built for daemons. It just didn't know it yet.

> The genealogy is not an analogy. It is the architecture. AI agents that persist, remember, and act are daemons. The pattern starts here.

[Next: the state persistence problem — what happens when an AI agent's state outruns its runtime.](https://blog.hackspree.com/#durable-daemons-trap)

---

*Part of the [Durable Daemons](/tags/durable-daemons-pattern) series.*
