---
title: Agentic-First CLI
date: 2026-08-03
slug: agentic-first-cli-design
summary: "The agent is a new category of end user — and the evidence says the interface is a performance variable, not cosmetics. SWE-agent's agent-computer interface doubled state-of-the-art on SWE-bench; frontier agents still score under 65% on terminal tasks. This post compares the design space (CLI vs function calling vs MCP vs chat) and distills the research into an agentic-first CLI checklist."
tags: cli, agentic, design, fred-brooks, the-design-of-design, conceptual-integrity, unix, structured-output, json, deterministic, agents, llm, contract, acis, swe-agent, terminal-bench
---

The agent is a new kind of end user, and the interface is the environment it lives in. That is the founding claim of [SWE-agent](https://arxiv.org/abs/2405.15793) (Yang et al., 2024), the paper that coined the term *agent-computer interface* (ACI):

> "Just as humans benefit from powerful software applications, such as integrated development environments, for complex tasks like software engineering, we posit that LM agents represent a new category of end users with their own needs and abilities, and would benefit from specially-built interfaces to the software they use."

Everything in this post follows from treating that sentence literally. The CLI has a new user, and it never blinks: it reads every byte of help text and output, remembers all of it, cannot answer a prompt, cannot see color, and pays for every token it reads. Most CLIs were designed for the old user — a human who can squint, scroll, and improvise. The agentic-first CLI is designed for the user who actually exists.

The thesis: **the interface is a first-class performance variable for agents, and the CLI is the right substrate for it — if designed with the discipline of a versioned public API. The evidence comes from the benchmark literature and the interface-design papers; the theory comes from Brooks's *The Design of Design*.**

## The interface is a performance variable

The SWE-agent result is the strongest single number in the field: with the same underlying model (GPT-4), a custom agent-computer interface achieved a **12.5% pass@1 on SWE-bench and 87.7% on HumanEvalFix — "far exceeding the previous state-of-the-art achieved with non-interactive LMs."** No new model, no new prompting trick: a better interface. The paper's conclusion is explicit: the design of the ACI changes agent behavior and performance.

Two benchmarks set the floor and the ceiling for the CLI specifically. [InterCode](https://arxiv.org/abs/2306.14898) (Yang et al., 2023) formalized interactive coding as a reinforcement-learning environment with "code as actions and execution feedback as observations" — the observation channel is the interface. [Terminal-Bench 2.0](https://arxiv.org/abs/2601.11868) (Merrill et al., 2026) built 89 hard, real-world terminal tasks and found that **frontier models and agents score under 65%** — then devoted an error analysis to *why*, because terminal interfaces, as they exist today, are bad ACIs: ambiguous output, interactive prompts, hidden state. [AgentBench](https://arxiv.org/abs/2308.03688) (Liu et al., 2023) reached the same conclusion across eight environments: how the agent observes and acts determines more of the outcome than the model's raw capability.

The practical translation: **every line your CLI emits is an observation your agent reasons over; every prompt it waits on is a stall; every hidden default is a hallucination risk.** The design of the interface is not a UX nicety. It is the agent's model of the world.

## The design space: four ways to expose a tool to an agent

| | CLI | Function calling | MCP | Chat |
|---|---|---|---|---|
| Structure | flags + `--json` | typed schema | typed schema | prose |
| Composability | pipes (Unix) | none | protocol | none |
| Observability | stdout/stderr/exit codes | app logs | protocol logs | chat log |
| Adoption cost | zero — it exists | per-tool SDK | protocol server | zero |
| Agent ergonomics | help, examples, exit codes | descriptions + schemas | tool docs | free-form |
| When it wins | everything Unix-shaped | inside one app | cross-tool discovery | humans |

Anthropic's [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents) (Dec 2024) is the most-cited engineering guidance on exactly this choice, and its conclusion favors the boring option: "the most successful implementations use simple, composable patterns rather than complex frameworks." Function calling and MCP solve real problems — typed I/O and cross-tool discovery — but each is a layer the tool must implement and maintain. The CLI already exists, is observable by construction, composes through pipes, and needs no new protocol. **The agentic-first CLI is the low-friction ACI: the discipline of a versioned API applied to the interface you already ship.**

## What the research says about CLI design

Each practice below is anchored to a source, not to taste.

**Structured output is the observation channel.** InterCode's framing — execution feedback as the observation — implies the feedback must be unambiguous. Prose output is a lossy observation: an agent that reads "Build succeeded. 12 targets, 3 warnings." will guess about the warnings, and its guesses are confident. Git solved this in 2009 with `status --porcelain`, a byte-for-byte stable machine format that ships *alongside* the human format; every state-producing command should offer the same: `--json`, data on stdout, nothing else.

**Determinism is trust.** An agent cannot satisfice — Herbert Simon's term, adopted by Brooks — a nondeterministic tool: if the same command yields different output, it must verify, and verification is the most expensive thing an agent does. So: no interactive prompts (detect non-TTY and fail fast, or provide `--yes`/`--no-input`); no hidden state (flags over config inference); sort by default; `--check` and `--dry-run` before anything destructive. An agent that trusts the tool runs once; an agent that does not runs three times.

**Help is the documentation the agent reads.** Anthropic's ACI guidance is the sharpest sentence in the field: "Carefully craft your agent-computer interface (ACI) through thorough tool documentation and testing." SWE-agent's ACI shipped documentation for its custom commands, and the paper credits it. An agent that trusts `--help` saves a full exploration cycle; a lie in help text is the most expensive bug an agentic CLI can have.

**Exit codes and stderr are the contract.** The interface has three channels — stdout (data), stderr (diagnostics), exit code (verdict) — and agents read all three. The checklists of the terminal benchmarks exist because agents misallocate effort when the channels are mixed: banners on stdout, logs where data belongs, exit 0 on failure.

**Consistency is learnability.** SWE-agent found interface design changes agent behavior; Brooks's conceptual integrity explains why: a system that feels like one mind designed it lets the agent's learned model of one subcommand transfer to the next. A CLI with five flag styles is a committee design, and the agent pays for it in tokens and mistakes.

**Quiet by default is the budget.** Brooks on budgets — design within time, memory, cost — applied to the agent's context window: verbose-by-default is a tax on every invocation, forever, at scale. `--verbose` opts in.

## Why the CLI, and why not a protocol

The counter-argument is worth taking seriously: if agents need good interfaces, build the interface from scratch — a purpose-built ACI, like SWE-agent did. The rebuttal is economics. A purpose-built ACI for your tool is what MCP servers and function schemas already are: another layer to write, document, and keep in sync with the actual tool. The CLI is the one interface that already exists, already documented, already versioned, already composable. The agentic-first discipline makes it *also* correct for agents — without inventing a protocol.

The exception is cross-tool discovery: when an agent must discover and bind tools at runtime across many systems, a protocol like [MCP](https://modelcontextprotocol.io/) earns its layer. But the interface underneath still needs the same design discipline — MCP tool descriptions and output schemas are the same contract as `--help` and `--json`, wearing a different hat. Anthropic's guidance applies at the layer you control: "reduce abstraction layers and build with basic components" in production.

## The checklist

- stdout is data only; stderr is diagnostics; logs to file
- `--json` on every state-producing command, stable documented schema
- exit codes: 0 = success, non-zero = failure, distinct code for "not run"
- no prompts: `--yes`, `--no-input`, TTY detection
- no color when piped; honor `NO_COLOR`
- deterministic: sorted output, no timestamps unless asked, no hidden config
- idempotent: `--check`, `--dry-run`, `--apply`
- quiet by default, `--verbose` opt-in
- complete, honest `--help` with examples
- one convention set across every subcommand
- versioned, additive contract

## The test

Run your CLI the way the benchmark environments run it: `--help`, one command, `--json`, another command — and read the output as a reader who never blinks, never asks, and never forgets. If any line could mean two things, the agent will choose the wrong one half the time, and it will do so confidently. The terminal benchmarks exist because that failure is measurable; the fix is the discipline above.

> The agent is a new kind of end user, and the interface is its environment. The contract is the architecture; the output is the model; determinism is respect; tokens are the budget; one mind owns the whole thing. Design for a user that never blinks — and every human at the terminal benefits too.

---

**References:**

- John Yang, Carlos E. Jimenez, Alexander Wettig, Kilian Lieret, Shunyu Yao, Karthik Narasimhan. [SWE-agent: Agent-Computer Interfaces Enable Automated Software Engineering](https://arxiv.org/abs/2405.15793). arXiv:2405.15793, 2024. — Agents as a new category of end users; custom ACI; 12.5% SWE-bench / 87.7% HumanEvalFix pass@1; interface design changes agent behavior.
- Mike A. Merrill, Alexander G. Shaw, Nicholas Carlini, Boxuan Li, Harsh Raj, Ivan Bercovich. [Terminal-Bench: Benchmarking Agents on Hard, Realistic Tasks in Command Line Interfaces](https://arxiv.org/abs/2601.11868). arXiv:2601.11868, 2026. — 89 terminal tasks; frontier agents under 65%; error analysis.
- John Yang, Akshara Prabhakar, Karthik Narasimhan, Shunyu Yao. [InterCode: Standardizing and Benchmarking Interactive Coding with Execution Feedback](https://arxiv.org/abs/2306.14898). arXiv:2306.14898, 2023. — Interactive coding as an RL environment: code as actions, execution feedback as observations.
- Xiao Liu et al. [AgentBench: Evaluating LLMs as Agents](https://arxiv.org/abs/2308.03688). arXiv:2308.03688, 2023. — Eight agent environments, including the operating system and terminal.
- [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents). Anthropic Engineering, Dec 2024. — "Carefully craft your agent-computer interface (ACI) through thorough tool documentation and testing"; simple, composable patterns over frameworks.
- [Claude Code — Best practices for agentic coding](https://docs.anthropic.com/en/docs/claude-code/best-practices). Anthropic. — Tool and interface design guidance from the team shipping the most-used agentic CLI.
- [Model Context Protocol](https://modelcontextprotocol.io/). Anthropic, 2024. — The protocol alternative for cross-tool discovery.
- [Git — Plumbing and Porcelain](https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain) and [`git status --porcelain`](https://git-scm.com/docs/git-status). — the machine-stable output contract, in production since 2009.
- Frederick P. Brooks Jr. [*The Design of Design*](https://www.cs.unc.edu/~brooks/DesignofDesign/). Addison-Wesley, 2010; and [*The Mythical Man-Month*](https://en.wikipedia.org/wiki/The_Mythical_Man-Month), 1975. — Models, the design tree, budgets, satisficing (after Herbert Simon), the separation of architecture from implementation, conceptual integrity, the one-mind rule.
- [NO_COLOR](https://no-color.org/) — the environment variable for disabling color output.
- Related: [Brooks on Software Design: Conceptual Integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) — the property the agentic contract depends on.
- Related: [Brooks on Software Design: One Mind](https://blog.hackspree.com/#brooks-design-one-mind-rule) — who owns the CLI contract, and why committees fail at it.
- Related: [The Unix Philosophy Is the Only Software Engineering Theory That Works](https://blog.hackspree.com/#unix-philosophy) — `stat` for data, `ls` for rendering; the outside-in view.
- Related: [CUPID: Properties, Not Principles](https://blog.hackspree.com/#cupid-for-joyful-coding) — composability, predictability, and the Unix property as directions, not rules.
- Related: [Task Harness Engineering](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) — the other half of the agent's toolchain: how harnesses consume CLI contracts.
