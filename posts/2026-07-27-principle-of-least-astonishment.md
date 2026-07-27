---
title: Designing Software for AI Coding Agents
date: 2026-07-27
slug: principle-of-least-astonishment
summary: AI coding agents are now your most demanding users. They read your APIs literally, follow your conventions religiously, and fail silently when surprised. POLA, convention over configuration, and conceptual integrity are not aesthetic preferences. They are the difference between an agent that builds on your system and one that cannot.
tags: software-engineering, ai-agents, ai-coding, pola, convention-over-configuration, conceptual-integrity, api-design
---

In 1972, an anonymous language designer wrote: "Every construct in the system should behave exactly as its syntax suggests." This is the Principle of Least Astonishment. For fifty years it was advice about human users. Today it is a requirement for AI coding agents — and they are less forgiving.

> A human user is astonished, confused, then adapts. An AI agent is astonished, hallucinates, then fails. The failure is silent. The output looks correct. It is wrong in a way you won't notice until it breaks.

AI coding agents — Claude Code, Cursor, Copilot, Codex — consume software differently than humans. They read your documentation literally. They infer behavior from names, types, and conventions. They do not develop superstitions. They do not ask clarifying questions when confused. They do not "get a feel" for your API after using it for a week. They read your interface and assume it is honest. When it isn't, they produce code that compiles, passes tests, and does the wrong thing. The cost of astonishment has gone up.

## The new user

An AI coding agent approaches your software like a new team member who has read every page of your documentation, memorized every convention in your ecosystem, and will never ask you a question. It builds a mental model from the surface of your system — function signatures, type names, directory structure, config files, error messages. If the surface is honest, the model is accurate. If the surface lies, the model is wrong. The agent does not know it is wrong. It writes code. The code runs. The bug is architectural.

> A human developer learns your system. An AI coding agent reads your system. The difference is that the agent assumes everything it reads is true.

This is why POLA matters more for agents than for humans. A human sees `getUser()` create a user and thinks "that's wrong, I'll check the docs." An agent sees `getUser()`, reads the name, assumes it gets a user, and writes code that calls it in a loop — creating a thousand users. The agent was not wrong. The function name was a lie. The agent believed it.

## Convention over configuration: the agent's map

AI coding agents are convention-first. They know Rails puts models in `app/models/`, that `Sales` maps to `sales`, that `has_many` expects a foreign key named `table_id`. They know this because conventions are machine-readable patterns. An agent can navigate a Rails codebase faster than a human because the conventions form a predictable topology. Every directory has a known purpose. Every naming pattern maps to a known behavior. The agent does not need to read configuration. The convention is the configuration.

> Conventions are an API for agents. Every convention you follow is a decision the agent doesn't have to reverse-engineer. Every convention you violate is a trap.

When a framework has strong conventions, agents thrive. Rails, Django, Next.js — agents produce idiomatic code because the conventions are legible. When a codebase has no conventions, or worse, violates the conventions of its own ecosystem, agents produce garbage. The agent assumes the ecosystem's defaults. The codebase overrode them silently. The agent didn't notice. The PR looks fine. The bug is in production.

Convention over configuration was designed to reduce developer toil. Its second-order effect — entirely unforeseen by DHH in 2004 — is that it makes software legible to AI. A convention is a promise: "this name means this behavior, always." Agents trust promises. When you keep the promise, the agent is productive. When you break it, the agent is dangerous.

## Conceptual integrity: the agent's trust model

Fred Brooks argued that conceptual integrity is the most important quality of a system. One design voice. Every part consistent with every other. In *The Design of Design* (2010), he decomposed it into three principles: **propriety** (no unnecessary features), **orthogonality** (no unintended coupling), **generality** (no artificial limitations). These principles were written for human teams. They are existential for agent-consumed software.

> An agent builds a model of your system from its surface. If the surface has conceptual integrity, the model is accurate. If it doesn't, the model is a hallucination waiting to execute.

**Propriety** means every function earns its place. An agent sees a public method and assumes it is safe to call. If half your public surface is internal machinery you never meant to expose, the agent will call it. The agent is not wrong. You exposed it. Propriety for agents means: if it's public, it's intended. If it's not intended, it's not public. The surface is honest.

**Orthogonality** means changes don't cascade. An agent modifies one file and assumes it hasn't broken three others. If your system has hidden coupling — a config flag that changes behavior in an unrelated module, a global state that flips between test and production — the agent cannot see it. The agent makes a local change. The system fails globally. The agent is not wrong. The coupling was invisible. Orthogonality for agents means: the dependency graph is the source graph. If `a.go` doesn't import `b.go`, changing `a.go` doesn't break `b.go`. The structure is honest.

**Generality** means no artificial limits. An agent uses your API within its documented parameters and hits a wall you didn't document — a rate limit, a size cap, a timeout that applies only on Tuesdays. The agent is not wrong. The limit was hidden. Generality for agents means: if there's a constraint, it's in the type signature, the error type, or the documentation. The constraint is honest.

> Propriety is an honest surface. Orthogonality is an honest dependency graph. Generality is an honest contract. Together they mean: the agent can trust what it reads.

## The agent design checklist

Building software for AI coding agents is not about optimizing for LLMs. It is about making your system's surface match its behavior. The same principles that make a system usable by humans make it usable by agents — but agents cannot compensate for violations the way humans can.

1. **Names are contracts.** If your function is called `getUser`, it must get a user. If it creates, updates, deletes, or sends email, rename it. An agent believes the name.

2. **Conventions are infrastructure.** Follow your ecosystem's conventions. Put files where the framework expects them. Name things what the framework expects them to be named. Every deviation is a decision the agent must reverse-engineer. Agents are bad at reverse-engineering.

3. **The public surface is the API.** If a function is public, it is documented. If it is not documented, it is not public. Agents cannot distinguish "public for internal use" from "public for external use." Neither can most humans, but agents pay a higher price.

4. **Errors are part of the interface.** An agent reads your error messages. If your error says "permission denied," the agent assumes the fix is to add permissions. If the real problem is a missing config flag and the error is wrong, the agent will spiral. Error messages are the agent's debugger. Make them honest.

5. **Dependencies are visible or they don't exist.** An agent sees what's in the import graph. If module A affects module B through a channel not in the import graph — global state, a database trigger, a cron job, an env var — the agent cannot see it. The agent will break it. Make dependencies visible or eliminate them.

> The agent believes the name. The agent follows the convention. The agent trusts the public surface. The agent reads the error message. The agent sees the import graph. If any of these are lies, the agent will produce code that is wrong in ways you will not detect until it matters.

## Open questions

**Do we need agent-aware type systems?** Types are machine-readable contracts. An agent that can query the type system — not just read docs — is an agent that can verify its own assumptions. Should we design type systems that agents can interrogate directly?

**Should frameworks publish agent-specific conventions?** A framework has implicit conventions that humans learn through experience. An agent needs them explicit. Should every framework ship a `CONVENTIONS.md` written for LLM consumption — structured, exhaustive, machine-parseable?

**How do you test an agent's understanding of your API?** You can't ask the agent "do you understand this?" — it will say yes either way. Do we need agent-specific integration tests: give the agent a task that requires understanding a particular convention, and verify it doesn't hallucinate?

**Does conceptual integrity become measurable?** Brooks argued for it qualitatively. But if an agent's success rate on your API correlates with your system's conceptual integrity, do we finally have a metric for it? Can we measure "agent success rate" and call it the integrity score?

**What happens when agents train on agent-generated code?** An agent writes code that works but violates conventions. Another agent reads that code as a training example. The violation becomes a pattern. The pattern becomes a de facto convention. Who governs the convention when the convention emerges from agent output rather than human intent?

**Is there a POLA for agents that differs from POLA for humans?** An agent is never astonished by verbosity. It is astonished by inconsistency. A human is the opposite. Do we need different design principles for agent-first APIs than for human-first APIs?

> The best-designed system for an AI coding agent is the best-designed system for a human. The difference is that the agent will not forgive you for cutting corners.

---

**References:**

- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- David Heinemeier Hansson. [Convention over Configuration](https://en.wikipedia.org/wiki/Convention_over_configuration). Ruby on Rails.
- Geoffrey James. (1987). *The Tao of Programming*.
- Related: [Durable Daemons — Pattern Specification](https://blog.hackspree.com/#durable-daemons-definition) — POLA for agent architecture.
