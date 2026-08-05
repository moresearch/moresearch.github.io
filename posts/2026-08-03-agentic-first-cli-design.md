---
title: Agentic-First CLI
date: 2026-08-03
slug: agentic-first-cli-design
summary: "The CLI has a new user that reads everything, asks nothing, remembers everything, and pays per token. Agentic-first design is not a new discipline — it is Fred Brooks's The Design of Design applied to a user who never blinks. The contract is the architecture, the output is the model, determinism is respect, tokens are the budget."
tags: cli, agentic, design, fred-brooks, the-design-of-design, conceptual-integrity, unix, structured-output, json, deterministic, agents, llm, contract, best-practices
---

The CLI has a new user, and it never blinks. It reads every byte of help text, every line of output, every error message — and remembers all of it. It cannot answer a prompt, because the tool will wait forever. It cannot see color. It pays for every token it reads. It is an AI agent, and it is now the most important reader of your command-line interface.

Most CLIs were designed for a different user: a human who can squint, scroll, and improvise. The agentic-first CLI is designed for the user who actually exists. This post is that design, using Fred Brooks's *The Design of Design* (2010) as its main reference — a design manual written fifteen years before its best audience existed, drawing on ideas fifty years old.

The thesis: **Agentic-first CLI design is not a new discipline. It is The Design of Design applied to a user who never blinks.** The contract is the architecture. The output is the model. Determinism is respect. Tokens are the budget. Conceptual integrity is the difference between a tool an agent learns once and a tool it must re-learn every invocation.

## The user is the judge

Brooks's first move is to take the user seriously: a design succeeds or fails in the user's hands. The agent's properties as a user, stated plainly: it reads everything verbatim; it never asks questions; it remembers everything within a context window; it pays per token; it cannot be surprised into learning — it generalizes from what it reads. Every practice below falls out of those properties: reads everything → structured output. Never asks → no prompts. Remembers everything → stability is a promise. Pays per token → quiet by default.

## The contract is the architecture

Brooks's most powerful doctrine is the separation of architectural design from implementation: the architecture is the contract visible to users; the implementation beneath it is free to change. For a CLI, the architecture is not the code — it is the command names, the flags, the output schema, the exit codes, the help text. That is the surface agents compile against.

Git is the canonical example. Its plumbing commands are the implementation; its porcelain commands are the architecture. And `status --porcelain` exists precisely so a machine can depend on a byte-for-byte stable format while the rest of the tool evolves. Porcelain mode is Brooks's separation made into a flag. The consequence: treat the contract as a versioned public API. A breaking change breaks every agent that ever learned the tool; additive change is free. Version the schema *before* you need to break it.

## The output is the model

Brooks on models: designers work through models, and the model determines what questions you can ask. For an agent, the CLI's output *is* its model of the system — the only evidence it has about what the command did. Prose output is a lossy model: an agent that reads "Build succeeded. 12 targets, 3 warnings." will guess about the warnings, and its guesses are confident. Structured output — `--json`, a stable schema, data on stdout and nothing else — gives the agent a model with no ambiguity.

Three rules: **stdout is data** (no banners, no "Done!", no logs); **stderr is diagnostics** (logs go to a file or behind `--log`); **`--json` is not optional** (if a command produces structured state, it must emit it as structured data). The Unix tradition knew this: `ls` does not know anything about files — `stat` provides the data, `ls` only renders it. The agent is the user who wants `stat` and does not want `ls`'s column widths.

## Determinism is respect

Herbert Simon's satisficing, which Brooks adopts, says designers accept the first acceptable solution rather than optimize. An agent cannot satisfice a nondeterministic tool: if the same command yields different output on successive runs, it must verify — and verification is the most expensive thing an agent does. So: no interactive prompts (detect non-TTY and fail fast, or provide `--yes` / `--no-input`); no hidden state (flags over config-file inference, because the agent cannot see what the config implies); no implicit ordering (sort by default); idempotent by default, with `--check` and `--dry-run` before anything destructive. An agent that trusts the tool runs once. An agent that does not runs three times.

## One mind, one contract

Conceptual integrity is the center of Brooks's argument: the system feels like one mind designed it. For a CLI, that means one convention set across every subcommand — the same flag shapes, the same error format, the same exit-code semantics, the same output style. The agent's learned model of one subcommand transfers to the next; a CLI with five flag styles is a committee design, and the agent pays for it in tokens and mistakes.

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — *The Mythical Man-Month*, 1975

The corollary Brooks sharpens in *The Design of Design*: the design must proceed from one mind. Every added flag is a subtraction from every existing flag, because the agent must learn them all. The essential skill of the designer is saying no — to smart people with good arguments — and having the authority to make it stick.

## Tokens are the budget

Brooks on budgets: design within budgets — time, memory, cost. The agentic-first budget is the context window; verbose-by-default is a tax on every invocation, forever, at scale. Quiet by default; `--verbose` opts in. Help text is written for reading — complete, accurate, with examples: an agent that trusts `--help` saves a full exploration cycle, and a lie in help text is the most expensive bug an agentic CLI can have. And the design tree — Brooks's model of exploration: try a branch, read the output, backtrack — is exactly what an agent does when it runs your tool. Complete help and stable output make the tree shallow; the cheapest investment is making the first branch always correct.

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

*The Design of Design* ends where all good design books do: with the user. Run your CLI the way the agent would — `--help`, one command, `--json`, another command — and read the output as a reader who never blinks, never asks, and never forgets. If any line could mean two things, the agent will choose the wrong one half the time, and it will do so confidently.

> The agent is the user, and the user is the judge. The contract is the architecture; the output is the model; determinism is respect; tokens are the budget; one mind owns the whole thing. Design for a user that never blinks — and every human at the terminal benefits too.

---

**References:**

- Frederick P. Brooks Jr. [*The Design of Design: Essays from a Computer Scientist*](https://www.cs.unc.edu/~brooks/DesignofDesign/). Addison-Wesley, 2010. — models, the design tree, budgets and constraints, satisficing, the separation of architectural design from implementation, conceptual integrity, the one-mind rule.
- Frederick P. Brooks Jr. [*The Mythical Man-Month*](https://en.wikipedia.org/wiki/The_Mythical_Man-Month). Addison-Wesley, 1975; 20th Anniversary Edition 1995. — conceptual integrity ("the most important consideration in system design"); the separation of architecture from implementation.
- Herbert A. Simon. *The Sciences of the Artificial*. MIT Press, 1969. — satisficing: designers accept the first acceptable solution.
- [Git — Plumbing and Porcelain](https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain) and [`git status --porcelain`](https://git-scm.com/docs/git-status). — the machine-stable output contract.
- [NO_COLOR](https://no-color.org/) — the environment variable for disabling color output.
- Related: [Brooks on Software Design: Conceptual Integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) — the property the agentic contract depends on.
- Related: [Brooks on Software Design: One Mind](https://blog.hackspree.com/#brooks-design-one-mind-rule) — who owns the CLI contract, and why committees fail at it.
- Related: [The Unix Philosophy Is the Only Software Engineering Theory That Works](https://blog.hackspree.com/#unix-philosophy) — `stat` for data, `ls` for rendering; the outside-in view.
- Related: [CUPID: Properties, Not Principles](https://blog.hackspree.com/#cupid-for-joyful-coding) — composability, predictability, and the Unix property as directions, not rules.
- Related: [Task Harness Engineering](https://blog.hackspree.com/#task-harness-engineering) — the other half of the agent's toolchain: how harnesses consume CLI contracts.
