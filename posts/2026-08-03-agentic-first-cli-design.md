---
title: "Agentic-First CLI: The Agent Is the User"
date: 2026-08-03
slug: agentic-first-cli-design
summary: "The CLI has a new user that reads everything, asks nothing, remembers everything, and pays per token. Agentic-first design is not a new discipline — it is Fred Brooks's The Design of Design applied to a user who never blinks. The contract is the architecture, the output is the model, determinism is respect, tokens are the budget."
tags: cli, agentic, design, fred-brooks, the-design-of-design, conceptual-integrity, unix, structured-output, json, deterministic, agents, llm, contract, best-practices
---

The CLI has a new user, and it never blinks. It reads every byte of help text, every line of output, every error message — and remembers all of it. It cannot ask a question when the tool prompts it, because the tool will wait forever. It cannot see color. It pays for every token the tool makes it read. It is an AI agent, and it is now the most important reader of your command-line interface.

Most CLIs were designed for a different user: a human at a terminal who can squint, scroll, re-run, and improvise. The agentic-first CLI is designed for the user who actually exists. This post is about what that design looks like, using Fred Brooks's *The Design of Design* (2010) as its main reference — a design manual written fifteen years before its best audience existed, drawing on ideas that are fifty years old.

The thesis: **Agentic-first CLI design is not a new discipline. It is The Design of Design applied to a user who never blinks.** The contract is the architecture. The output is the model. Determinism is respect. Tokens are the budget. And conceptual integrity is the difference between a tool an agent can learn once and a tool an agent must re-learn on every invocation.

## The user is the judge

Brooks's first move is to take the user seriously: a design succeeds or fails in the user's hands, and the user's experience *is* the design. The classic error is designing for a model of the user instead of the user.

The agent's properties as a user, stated plainly:

- reads everything, verbatim
- never asks questions — cannot
- remembers everything, within a context window
- pays per token read and written
- runs a command hundreds of times without boredom — but each run costs
- cannot be surprised into learning; it generalizes from what it reads

Every best practice below falls out of those properties. Reads everything → output must be structured. Never asks → no prompts. Remembers everything → stability is a promise. Pays per token → quiet by default.

## The contract is the architecture

Brooks's most powerful doctrine, from *The Design of Design* and *The Mythical Man-Month* before it, is the separation of architectural design from implementation: the architecture is the contract visible to users; the implementation beneath it is free to change. For a CLI, the architecture is not the code. It is the command names, the flags, the output schema, the exit codes, the help text. That is the surface agents compile against.

Git is the canonical example. Its plumbing commands are the implementation; its porcelain commands are the architecture. And `status --porcelain` exists precisely so a machine — a script, an agent — can depend on a byte-for-byte stable format while the rest of the tool evolves beneath it. Porcelain mode is Brooks's separation made into a flag.

The consequence: an agentic-first CLI treats its contract as a public API with the seriousness of a versioned service. A breaking change breaks every agent that ever learned the tool; additive change is free. Version the schema *before* you need to break it.

## The output is the model

Brooks on models: designers work through models, and the model determines what questions you can ask of the design. For an agent, the CLI's output *is* its model of the system — the only evidence it has about what the command did.

Prose output is a lossy model. An agent that reads "Build succeeded. 12 targets, 3 warnings." holds a fuzzy model; it will guess about the warnings, the structure, the next step — and its guesses are confident. Structured output — `--json`, a stable schema, data on stdout and nothing else — gives the agent a model with no ambiguity.

Three rules:

- **stdout is data.** No banners, no "Done!", no logs.
- **stderr is diagnostics.** Logs go to a file or behind `--log`.
- **`--json` is not optional.** If a command produces structured state, it must be able to emit it as structured data.

The Unix tradition already knew this. `ls` does not know anything about files; `stat` provides the data, `ls` only renders it. The agent is the user who wants `stat` and does not want `ls`'s column widths.

## Determinism is respect

Herbert Simon's satisficing, which Brooks adopts, says designers accept the first acceptable solution rather than optimize. An agent cannot satisfice a nondeterministic tool. If the same command yields different output on successive runs, the agent must verify — and verification is the most expensive thing an agent does.

- No interactive prompts. Detect non-TTY and fail fast, or provide `--yes` / `--no-input`.
- No hidden state. Flags over config-file inference, because the agent cannot see what the config implies.
- No implicit ordering. Sort output by default.
- Idempotent by default, with `--check` and `--dry-run` before anything destructive.

An agent that trusts the tool runs once. An agent that does not runs three times.

## One mind, one contract

Conceptual integrity is the center of Brooks's argument: the system feels like one mind designed it. For a CLI, integrity means one convention set across every subcommand — the same flag shapes, the same error format, the same exit-code semantics, the same output style. The agent's learned model of one subcommand transfers to the next. A CLI with five flag styles is a committee design, and the agent pays for it in tokens and mistakes.

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — *The Mythical Man-Month*, 1975

And the corollary Brooks sharpens in *The Design of Design*: the design must proceed from one mind. Every added flag is a subtraction from every existing flag, because the agent must learn them all. The essential skill of the designer is saying no — to smart people with good arguments — and having the authority to make it stick.

## Tokens are the budget

Brooks on budgets: design within budgets — time, memory, cost. The agentic-first budget is the context window. Verbose-by-default is a tax on every invocation, forever, at scale.

- **Quiet by default.** `--verbose` opts in.
- **Help text is written for reading** — complete, accurate, with examples. An agent that trusts `--help` saves a full exploration cycle; an agent that catches `--help` lying wastes one. A lie in help text is the most expensive bug an agentic CLI can have.
- **The design tree is what the agent does.** Brooks's model of design exploration — try a branch, read the output, backtrack — is exactly an agent running your tool. Complete help and stable output make the tree shallow. The cheapest design investment is making the first branch always correct.

## The checklist

- stdout is data only; stderr is diagnostics; logs to file
- `--json` on every state-producing command, with a stable, documented schema
- exit codes: 0 = success, non-zero = failure, a distinct code for "not run" (help, dry-run, invalid input)
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

**Related on this site:**

- [Brooks on Software Design: Conceptual Integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) — the property the agentic contract depends on.
- [Brooks on Software Design: One Mind](https://blog.hackspree.com/#brooks-design-one-mind-rule) — who owns the CLI contract, and why committees fail at it.
- [The Unix Philosophy Is the Only Software Engineering Theory That Works](https://blog.hackspree.com/#unix-philosophy) — `stat` for data, `ls` for rendering; the outside-in view.
- [CUPID Is Not a Set of Principles. That Is the Point.](https://blog.hackspree.com/#cupid-for-joyful-coding) — composability, predictability, and the Unix property as directions, not rules.
- [Task Harness Engineering](https://blog.hackspree.com/#task-harness-engineering) — the other half of the agent's toolchain: how harnesses consume CLI contracts.
