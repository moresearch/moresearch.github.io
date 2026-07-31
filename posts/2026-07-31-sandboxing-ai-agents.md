---
title: Sandboxes Are Hard
date: 2026-07-31
slug: sandboxing-ai-agents
summary: Amjad Masad, CEO of Replit, argues that the real problem with AI agent security is not the agent. It is the sandbox. Sandboxing is extraordinarily difficult, and most implementations — including those from dedicated vendors — make fundamental errors. The first rule of security is humility: assume every layer will fail. Build the next one to catch it.
tags: ai-agents, security, sandboxing, defense-in-depth, replit, infrastructure
---

The headlines say AI agents are escaping their sandboxes. Amjad Masad, CEO of Replit, says the headlines have it backwards. The problem is not that agents are clever. The problem is that sandboxes are hard — and most people build them wrong.

> "The first rule of security is humility. Assume zero-days exist — because they do. Assume your isolation will eventually fail."

Masad would know. Replit has run arbitrary, untrusted code since 2016 — surviving attacks from hobbyists, researchers, and state actors. The company's experience is the argument. A sandbox is not a product you buy. It is an architecture you build — layer by layer, each layer assuming the one below it will break.

> Having a sandbox is not the same as having a security architecture. The first is a feature. The second is a stack.

Replit's security stack is thirteen layers deep. Zero-trust service-to-service auth with short-lived tokens. Linux containers hardened with seccomp-bpf. Per-customer GCP Projects — even free-tier users get their own project for maximum tenant isolation. MicroVM migration in progress to eliminate the shared kernel entirely. An append-only Git sidecar so that even if the agent deletes its `.git` directory, history is recoverable. A transparent proxy for secrets — the agent's code never receives passwords. MCP tool calls proxied through a scanning layer that blocks prompt injection attacks. Continuous penetration testing via HackerOne and Trail of Bits. An internal AI red-teaming harness that continuously scans, prioritizes, and validates findings before engaging engineers.

> Thirteen layers. Each layer assumes the one below it will fail. When one does — and one always does — the rest hold. This is not paranoia. This is engineering.

The design principle is defense in depth. It is the opposite of trusting a single boundary. A container is not a sandbox. A microVM is not a sandbox. A hypervisor is not a sandbox. Each is a layer of a sandbox. The sandbox is the stack — the accumulated constraints that make the agent's behavior predictable even when one constraint fails.

This connects to a deeper argument this blog has been making. The stack is the design abstraction. Each layer eliminates a region of the failure space. A container eliminates host compromise. A microVM eliminates kernel sharing. A transparent proxy eliminates secret leakage. An append-only sidecar eliminates history deletion. Each layer is a constraint on what can go wrong. The stack is the accumulated set of constraints that make deployment safe.

> The agent is non-deterministic. The sandbox is the constraint that bounds the non-determinism. When the agent does something unexpected — and it will — the sandbox ensures the unexpected stays contained.

The hardest part of sandboxing is not the technology. It is the humility. You must accept that every layer you build will eventually fail. You must build the next layer anyway. You must resist the temptation to declare the system secure because it passed a penetration test or earned a compliance certification. Compliance is not security. A clean pen test is a snapshot, not a guarantee. The only honest posture is: assume compromise, contain the blast radius, monitor continuously, respond fast, harden the root cause, and start again.

> The companies that get sandboxing wrong are the ones that believe their own marketing. The companies that get it right are the ones that have been attacked for a decade and learned what humility costs.

---

**References:**

- Amjad Masad. [Sandboxes Are Hard](https://www.linkedin.com/feed/update/urn:li:activity:7488800810252890112/). LinkedIn, July 2026.
- Replit. [Defense in Depth: How Replit Secures Every Layer of the Vibe Coding Stack](https://replit.com/blog/defense-in-depth-how-replit-secures-every-layer-of-the-vibe-coding-stack). Replit Blog, 2026.
- Related: [The Stack as a Design Abstraction](https://blog.hackspree.com/#the-stack-is-the-abstraction) — Each layer is a constraint on the failure space.
- Related: [Durable Daemons — Pattern Specification](https://blog.hackspree.com/#durable-daemons-definition) — The four conditions for trustworthy agents.
