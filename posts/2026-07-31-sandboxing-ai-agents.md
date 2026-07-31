---
title: Sandboxes Are Hard
date: 2026-07-31
slug: sandboxing-ai-agents
summary: Amjad Masad, CEO of Replit, argues that the real problem with AI agent security is not the agent. It is the sandbox. Sandboxing is extraordinarily difficult, and most implementations — including those from dedicated vendors — make fundamental errors. The first rule of security is humility: assume every layer will fail. Build the next one to catch it.
tags: ai-agents, security, sandboxing, defense-in-depth, replit, infrastructure
---

AI agents escape sandboxes. The headlines blame the agents. Amjad Masad, CEO of Replit, blames the sandboxes. His argument: sandboxing is an extraordinarily difficult infrastructure problem, and most implementations — including those from dedicated vendors — make fundamental errors. The agent is not the threat. The single layer of isolation pretending to be a security architecture is the threat.

> "The first rule of security is humility. Assume zero-days exist — because they do. Assume your isolation will eventually fail."

Masad would know. Replit has run arbitrary, untrusted code since 2016 — surviving attacks from hobbyists, researchers, and state actors. The company's experience is the argument. A sandbox is not a product you buy. It is an architecture you build — layer by layer, each layer assuming the one below it will break.

> Having a sandbox is not the same as having a security architecture. The first is a feature. The second is a stack.

Replit's security stack is thirteen layers deep. Zero-trust service-to-service auth. Linux containers hardened with seccomp-bpf. Per-customer GCP Projects for tenant isolation. MicroVM migration underway to eliminate the shared kernel. An append-only Git sidecar — history survives even if the agent deletes `.git`. A transparent secrets proxy — application code never sees credentials. MCP tool calls scanned for prompt injection. Built-in auth via Clerk — the agent never implements authentication from scratch. Forkable databases so development never touches production. Determinate Nix for supply chain integrity. HackerOne and Trail of Bits for continuous external assessment. An internal AI red-teaming harness that scans, prioritizes, and validates findings before engineers are engaged.

> Thirteen layers. Each layer assumes the one below it will fail. When one does — and one always does — the rest hold. This is not paranoia. This is engineering.

The design principle is defense in depth. It is the opposite of trusting a single boundary. A container is not a sandbox. A microVM is not a sandbox. A hypervisor is not a sandbox. Each is a layer of a sandbox. The sandbox is the stack — the accumulated constraints that make the agent's behavior predictable even when one constraint fails.

The stack is the design abstraction. Each layer constrains a region of the failure space. A container constrains host compromise. A microVM constrains kernel sharing. A transparent proxy constrains secret leakage. An append-only sidecar constrains history deletion. No single layer is sufficient. The stack is the accumulated set of constraints.

> The agent is non-deterministic. The sandbox is the constraint that bounds the non-determinism. When the agent does something unexpected, the sandbox ensures the unexpected stays contained.

The hardest part of sandboxing is not the technology. It is the posture. You must accept that every layer will eventually fail and build the next one anyway. You must resist the temptation to declare the system secure because it passed a penetration test. Compliance is not security. A clean pen test is a snapshot. The only honest posture: assume compromise, contain the blast radius, monitor continuously, respond fast, harden the root cause, repeat.

> The companies that get sandboxing wrong are the ones that believe their own marketing. The companies that get it right are the ones that have been attacked for a decade and learned what humility costs.

---

**References:**

- Amjad Masad. [Sandboxes Are Hard](https://www.linkedin.com/feed/update/urn:li:activity:7488800810252890112/). LinkedIn, July 2026.
- Replit. [Defense in Depth: How Replit Secures Every Layer of the Vibe Coding Stack](https://replit.com/blog/defense-in-depth-how-replit-secures-every-layer-of-the-vibe-coding-stack). Replit Blog, 2026.
- George Fahmy. [The Agent Sandbox Taxonomy](https://github.com/kajogo777/the-agent-sandbox-taxonomy). GitHub, 2026. — Open-source taxonomy of sandboxing approaches.
- [gVisor](https://gvisor.dev/) — Google's application kernel for container sandboxing.
- [Firecracker](https://firecracker-microvm.github.io/) — AWS's microVM for multi-tenant isolation. Powers Lambda and Fargate.
- [Trail of Bits](https://www.trailofbits.com/) — Security assessment of Replit's infrastructure.
- [Clerk](https://clerk.com/) — Authentication-as-a-service. Eliminates entire classes of auth bugs.
- [Semgrep](https://semgrep.dev/) — Static analysis for vulnerability detection in the scanning pipeline.
- [HoundDog](https://hounddog.ai/) — Privacy issue detection in pre-publish security scanning.
- Related: [The Stack as a Design Abstraction](https://blog.hackspree.com/#the-stack-is-the-abstraction) — Each layer constrains the failure space.
- Related: [Durable Daemons — Pattern Specification](https://blog.hackspree.com/#durable-daemons-definition) — Trustworthy agents require defense in depth.
