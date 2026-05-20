---
draft: true
title: Self‑Improving Agent Workflows: SEP, OpenDev, and the Agent Harness Design Space
date: 2026-05-20
summary: A practical recipe that combines SEP's self‑improving skill loop with OpenDev's terminal harness architecture and an empirical survey of harness design choices.
tags: [agents, harness, workflow, research]
---

This post has been consolidated into [Harness Engineering: Best Practices for Reliable Agent Systems](#harness-engineering-best-practices-for-ai-agents). See that canonical post for the full, merged guidance.
TL;DR

Combine SEP's simple "read lessons → do work → reflect → write lessons" skill loop with OpenDev's harness-level practices (multi-model routing, context compaction, safety layers) and the empirical design-space findings from an architecture survey. Start with one skill and one loop; let usage generate the data you need to grow a robust harness.

Introduction

SEP's "The workflow that teaches itself" proposes a minimal, practical loop: make each task a stepwise skill that reads past lessons at start and writes new lessons at the end. Over time the skill improves itself from real failures and fixes. This is a low-friction way to bootstrap reliable agent behavior without designing heavyweight infrastructure first.

Why mix SEP with system papers

Two recent arXiv reports make complementary points.

- OpenDev (arXiv:2603.05344v1) argues for a compound, multi-model terminal harness with explicit context compaction, dual-agent planning/execution separation, and layered safety. Those practices solve long-running, token-budgeted CLI agent problems.
- An empirical architecture survey (arXiv:2604.18071v1) catalogs recurring harness design dimensions: subagent structure, context management, tool systems, safety, and orchestration. It shows which choices tend to co-occur across real projects.

Bringing them together

SEP gives a practical, incrementally deployable loop. OpenDev and the survey show the harness patterns that emerge as those loops scale and as teams demand reliability, safety, and multi-model efficiency.

- Start small: implement a skill with steps and a lessons file (team + personal). Let runs accumulate real failure modes.
- When recurring issues appear, promote lessons into explicit rules inside the skill or into harness-level hooks (e.g., automatic retries, sandboxing checks).
- Use the survey's design dimensions as a checklist when a skill grows into a subsystem: do we need subagent decomposition, a tool registry, stronger context persistence, or a governance layer?
- Apply OpenDev tactics for long-running terminal agents: separate planning from execution, compact context proactively, route workloads to models by role, and add staged safety vetoes.

Figures

![Knowledge‑first vs Path‑first](https://sep.com/wp-content/uploads/2026/05/visual-01-knowledge-first-vs-path-first-1024x683.jpg)
_Figure 1 — SEP: path‑first thinking (name the step, confine the scope)._ 

![Core learning loop](https://sep.com/wp-content/uploads/2026/05/visual-02-core-learning-loop-1024x683.jpg)
_Figure 2 — SEP: attach learning to the path (read lessons … write lessons)._ 

![Synthesis diagram](images/self-improving-harness.svg)
_Figure 3 — Synthesis: skill loop (left) invokes harness layers (right). The harness implements the tool registry, context compaction, safety layers, and persistence the skill depends on as it scales._

A short getting‑started checklist

1. Pick one repeatable task (dev DB reset, data import, report generation).
2. Write it as a stepwise skill with: 1) read lessons, 2) execute steps, 3) reflect, 4) write lessons.
3. Store team lessons in the repo and user lessons in a gitignored file.
4. Run the skill often; catalogue recurring failure modes.
5. When failures repeat, promote durable lessons into explicit handlers or harness hooks.
6. Use the survey's five design dimensions (subagents, context, tools, safety, orchestration) to plan harness upgrades.

Why this approach works

- Low friction: you get actionable telemetry before investing in big harness projects.
- Composability: skills remain small and learn inside their lanes; orchestrators and harness layers coordinate them at scale.
- Evidence driven: empirical architecture patterns (2604.18071v1) guide which harness investments pay off across projects.

References

- SEP: The workflow that teaches itself — https://sep.com/blog/the-workflow-that-teaches-itself-a-self-improving-agent-workflow/
- OpenDev (compound CLI agent): https://arxiv.org/html/2603.05344v1
- Agent harness design-space survey: https://arxiv.org/html/2604.18071v1

---

Notes: Figures 1–2 are from SEP's public blog (linked above). Figure 3 is a synthesized diagram included in this post to show how SEP's skill loop maps to harness layers described in the OpenDev paper and the architecture survey.
