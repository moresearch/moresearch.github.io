---
title: "Emerging Harness Patterns and Anti-Patterns for Agentic AI Systems"
date: 2026-08-23
slug: harness-patterns-for-agentic-ai-systems
summary: "A catalog of sixteen emerging harness patterns and twelve anti-patterns for agentic AI systems, organized into four families — Safety, Security & Isolation; State, Memory & Context; Tool Binding & Interfacing; Orchestration & Topology — each pattern and anti-pattern explicitly subtitled with its description, its reference, its connection to this blog's own archive, and its tradeoffs. It closes with patterns on the 2026 frontier: the generator-evaluator loop, sprint contracts, context resets with handoff artifacts, the interop layer (MCP, ACP, AGENTS.md), and live-environment evaluators."
tags: harness, harness-engineering, agent-harness, patterns, anti-patterns, agents, agentic-ai, security, sandboxing, memory, orchestration, tool-binding, verification, token-economics, emerging-patterns
---

A harness is the environment an agent lives in: everything between the model and the world. It decides what the agent can see, what it can do, and how the outcome is judged. This blog has been arguing for months that the harness — not the model — is where the leverage concentrates, and the numbers have arrived: the same model finished between 47% and 67% of real tasks across eight different harnesses while cost per finished task varied sevenfold, and the same weights score 3.4 points apart on the public Terminal-Bench board under two major harnesses. Nothing about the intelligence changed. The wrapper changed.

This post is the second edition of this blog's pattern catalog. The first edition, published earlier this year, distilled thirteen patterns from the blog's own archive — the append-only session log, the loop as a config row, policy waterfalls with monotonic guards, capability seams, durable execution, the sandbox stack, the verifier layer, and the rest. This edition reorganizes the catalog around a different axis, and it reflects the field's movement in the intervening months. The agent ecosystem has converged on four families of harness concern — **Safety, Security & Isolation**; **State, Memory & Context**; **Tool Binding & Interfacing**; **Orchestration & Topology** — and inside each family a set of named patterns and anti-patterns has emerged from the frameworks and the literature: ephemeral sandboxes, intercepting gatekeepers, breakpoints, budget throttlers, rolling-window compression, memory routers, state snapshots, schema enforcement, tool registries, orchestrator-workers, blackboards, pipelines, and ensembles — and the twelve ways to get each of them wrong.

Every pattern below is subtitled, described, referenced to the primary source, connected to the relevant posts in this blog's archive, and weighed: what it buys, what it costs, and where its guarantee stops. The anti-patterns are the same shape with the sign flipped — each one is a recurring failure with a reference that documents it and a way out that the pattern half of this catalog already supplies. The frontier section at the end covers what is emerging as of mid-2026 and is not yet in any framework's stable documentation: the generator-evaluator loop, sprint contracts, context resets with handoff artifacts, the interop layer, and live-environment evaluators.

| Family | Patterns | Anti-patterns |
|---|---|---|
| Safety, Security & Isolation | Ephemeral Sandbox Wrapper · Static Intercepting Gatekeeper · HITL Breakpoint · Token & Time Budget Throttler | The Naked Prompt · The Infinite Execution Vortex · Prompt-Driven Authorization |
| State, Memory & Context | Rolling Window Compression · Semantic Memory Router · State Snapshot & Rollback · Tiered Hierarchical Memory | The Context Avalanche · Goldfish Amnesia · The RAG Firehose |
| Tool Binding & Interfacing | Schema Enforcement & Self-Correction · Asynchronous Tool Worker Queue · Mock Tool Virtualization · Dynamic Tool Discovery / Registry | The Silent Crash · Bloated Utility Belt · The Schema Free-for-All |
| Orchestration & Topology | Orchestrator-Worker · Blackboard · Sequential Pipeline Routing · Voting / Consensual Ensemble | The Committee Paradox · The God Agent · State Race Conditions |

# Part I — Emerging Harness Patterns

## Safety, Security & Isolation Patterns

The safety family is where the harness earns its name. The pattern set here is the operational translation of the blog's earlier arguments: defense in depth from [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents), fail-closed policy from the [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), and subtraction from [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface). What has emerged since is a concrete vocabulary: the wrapper, the gatekeeper, the breakpoint, and the throttle — four distinct seams where the harness inserts itself between the model and the world.

### Pattern 1 — Ephemeral Sandbox Wrapper

**Description.** Spawns isolated, short-lived virtual environments (e.g., Docker, WASM) for safe, untrusted agent code execution.

**Reference.** [LangChain Sandbox Runtime Documentation](https://python.langchain.com/docs/integrations/) — sandboxing is now a first-class integration category in the LangChain ecosystem, alongside chat models and vector stores, with the explicit contract that agent-authored code runs in disposable environments rather than on the host.

**Connection.** This is the microVM end of the spectrum this blog mapped in [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents): a container is not a sandbox, a microVM is not a sandbox — each is a layer. The ephemeral wrapper adds the *lifetime* dimension: the environment is created per task, mutated freely, and destroyed. It composes with the DeepSeek harness's sandbox vocabulary ([file-effects-only, bash-equivalent trust](https://blog.hackspree.com/#deepseek-harness)) and with the [system boundary](https://blog.hackspree.com/#spatiotemporal-composability) concept — the wrapper is where the harness reifies the outside world so that recovery can be promised. It is also the training-loop pattern from the [minimal preset](https://blog.hackspree.com/#deepseek-harness): the RL environment is an ephemeral shell, which is why the trajectories it produces are clean.

**Tradeoff.** Ephemerality costs startup latency and state loss — an agent that needs a persistent filesystem across steps fights the wrapper, which is why the pattern is usually paired with a persistent-but-isolated volume. Real isolation needs a hypervisor or OS-level sandbox, which is heavy; WASM containers are lighter and faster but constrain what the agent can do. And the wrapper is only as good as its teardown: if the environment survives, the "ephemeral" claim is marketing. The DeepSeek lesson applies — [teardown derived from the load, not remembered](https://blog.hackspree.com/#spatiotemporal-composability).

### Pattern 2 — Static Intercepting Gatekeeper

**Description.** Intercepts model-generated tool calls against a strict blocklist before passing them to external APIs.

**Reference.** [Llama Guard: LLM-based Input-Output Safeguard for Human-AI Conversations](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/) (Meta AI, 2023) — a classification model, tuned to be the gate between model output and the world, that scores both prompt and response against a safety taxonomy before anything external happens.

**Connection.** The gatekeeper is the *static* answer to the same problem the [tool pipeline](https://blog.hackspree.com/#deepseek-harness) solves dynamically: pre-execute waterfalls, monotonic guards, and a fail-closed approval seam that can only deny or abstain, never force-allow. The DeepSeek doctrine applies directly — the gatekeeper must run *before* policy observers ever see the call, because "[pre-execute listeners, approval, and guards must never observe — or worse, approve — a call that can only fail](https://blog.hackspree.com/#deepseek-harness)". The gatekeeper is the harness's version of [`pledge(2)`](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering): a restricted interface where the wrong thing is unexpressible.

**Tradeoff.** The honest caveat is in the name: Llama Guard is *not* static — it is itself an LLM, which means it can be fooled, bypassed, or prompted around, and it adds a model call to every tool invocation. A true static blocklist (deny specific tool names, argument patterns, URL prefixes) is fast, deterministic, and auditable, but it can only catch what it enumerates. The pattern works when the blocklist is the floor and the model-based classifier is the next layer up — the [monotonic guard ordering](https://blog.hackspree.com/#deepseek-harness): deny by default, allow by exception, never let policy be argued around.

### Pattern 3 — Human-in-the-Loop (HITL) Breakpoint

**Description.** Freezes the harness execution loop to demand manual approval for high-risk mutations like financial transactions.

**Reference.** [LangGraph Interrupts (State Management & Breakpoints)](https://docs.langchain.com/oss/python/langgraph/interrupts) — `interrupt()` as a first-class graph primitive that pauses execution, persists the graph state, and resumes from the exact checkpoint after a human approves, edits, or rejects.

**Connection.** The breakpoint is the governance seam from the [always-on agents survey](https://blog.hackspree.com/#always-on-agents): authority and scope are not properties of the model, they are properties of the harness, and the breakpoint is where the harness asserts them. It is also the [approval seam](https://blog.hackspree.com/#deepseek-harness) — `allowed-once` as the only granting outcome, with a missing answerer resolving to `unavailable` — and the [OpenWorker approval gate](https://blog.hackspree.com/#openworker-outcome-layer) before consequential actions. The pattern's power is that it is a *state machine* primitive, not a prompt: the loop does not merely ask permission, it persists where it paused.

**Tradeoff.** Breakpoints trade autonomy for safety and latency for trust — an agent that must stop for every risky action cannot run unattended, which is why the pattern must be paired with an explicit automation path (permission presets, scoped allowlists) or the harness drowns in approval requests and humans rubber-stamp everything, which is worse than no approval at all. The LangGraph guidance to not wrap interrupts in try/except is the deep lesson: a breakpoint that can be swallowed is a breakpoint that does not exist. And the trust boundary must be drawn honestly — a breakpoint protects the mutation, not the model's intent behind it.

### Pattern 4 — Token & Time Budget Throttler

**Description.** Monitors continuous tool loops to forcefully terminate agents exceeding maximum token costs or time boundaries.

**Reference.** [AutoGPT — Maximum Loop & Cost Control](https://docs.agpt.co/classic/configuration/) — the first widely deployed agent included iteration limits and cost controls as first-class settings, precisely because the autonomous loop it popularized demonstrated the failure they prevent.

**Connection.** This is the enforcement half of the [Bill as Assertion](https://blog.hackspree.com/#every-token-has-a-price-tag) argument. The token economics post computed the shape: the unit got cheaper but the task did not (an autocomplete action eats a few hundred tokens; an agentic task eats millions), bills tripled even as prices fell 98%, and the only response to unbounded consumption is a hard ceiling. The OWASP Top 10 for LLM applications now names this as a first-class risk — [LLM10:2025 Unbounded Consumption](https://genai.owasp.org/llm-top-10/). The DeepSeek harness's [runaway-loop guard](https://blog.hackspree.com/#deepseek-harness) shows the honest version of the pattern's limit: it "only sends reminders and eventually goes quiet" — which is why the throttler must be enforced in the harness, not requested of the model.

**Tradeoff.** A budget that is too tight kills legitimate long-horizon work; a budget that is too loose is theater. Time and token ceilings interact with [prefix caching](https://blog.hackspree.com/#deepseek-harness): a throttler that resets the request mid-session can invalidate the warm prefix and *increase* the bill it was meant to cap. And the throttler must decide what "exceeded" means — total spend, per-step spend, wall-clock time, or loop iterations — because each catches a different failure mode, and the monetary ceiling is the one the firm actually cares about ([the bill is 2-14% of labor cost](https://blog.hackspree.com/#every-token-has-a-price-tag), and the caps were about the shape of the bill, not the money).

## State, Memory & Context Patterns

The state family is where the field moved fastest. The [always-on agents survey](https://blog.hackspree.com/#always-on-agents) showed the gap: the literature is good at accumulating and retrieving state and almost silent on governing it. The emerging patterns below are the *retrieval and management* half — the governance half (validation, forgetting, rollback) is still under-built, which is exactly why the anti-patterns in this family are the most common in production.

### Pattern 5 — Rolling Window Compression

**Description.** Automatically summarizes older conversation histories in background threads to keep the active context window lean.

**Reference.** [MemGPT: Towards LLMs as Operating Systems](https://arxiv.org/abs/2310.08560) (Packer et al., UC Berkeley) — virtual context management: data moves between fast and slow memory tiers under an OS-style paging discipline, with interrupts to manage control flow.

**Connection.** This blog's prior treatment is the [append-only log](https://blog.hackspree.com/#deepseek-harness), and the two patterns are in direct tension — the append-only log *cannot* rewrite the past, and compression is precisely a rewrite. The DeepSeek answer is the disciplined form: the summarizer must not touch the log; it produces a new prefix-extension message in the Claude-shaped eight-section checkpoint format, and the [compaction bug](https://blog.hackspree.com/#deepseek-harness) — a fresh system prompt at the front invalidating the entire cache prefix — is the pattern's canonical failure. The pattern is a paging discipline: the context window is RAM, the log is disk, and the summary is the page table.

**Tradeoff.** Summaries lose detail, and the loss is permanent unless the source log is preserved — which is why compression and the append-only log must be separate layers. Compaction preserves continuity but not a clean slate: Anthropic's 2026 harness work found models exhibit "context anxiety" and begin wrapping up prematurely as the context fills, and compaction alone did not fix it ([Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps)). And every compaction is a billing event — if the summarization call is not a genuine prefix-extension of the warm request, it pays full price for the whole replayed history.

### Pattern 6 — Semantic Memory Router

**Description.** Intercepts ongoing tasks, queries vector stores, and injects context fragments just-in-time into the agent's prompt.

**Reference.** [Pinecone — Retrieval-Augmented Generation](https://www.pinecone.io/learn/retrieval-augmented-generation/) — RAG as the standard architecture for grounding model output in external knowledge, with retrieval interposed between the task and the prompt.

**Connection.** The router is the *decision* layer on top of RAG: not every chunk belongs in every prompt, and the harness — not the model — decides what to inject. This connects to the [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) principle that every line the agent reads is an observation it reasons over; injected context is the highest-leverage observation there is. It also inherits the [always-on governance](https://blog.hackspree.com/#always-on-agents) questions: provenance of the retrieved state, authority over what can be injected, and recoverability when the injection is wrong.

**Tradeoff.** The router is a point of failure that is also a point of attack: retrieved content is untrusted input, and prompt injection through a vector store is the vector of choice ([LLM08:2025 Vector and Embedding Weaknesses](https://genai.owasp.org/llm-top-10/)). Just-in-time injection competes with [prefix-cache discipline](https://blog.hackspree.com/#deepseek-harness) — context that changes every turn destroys the cache bookmark unless the injected region sits outside the prefix. And the router must fight the [RAG firehose](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) anti-pattern below: retrieval quality is decided by chunking, metadata filtering, and reranking, not by top-K volume.

### Pattern 7 — State Snapshot & Rollback

**Description.** Saves complete system state snapshots at checkpoint N, allowing recovery if an agent hits an error loop at step N+3.

**Reference.** [Temporal — Workflow State & Replay](https://docs.temporal.io/) — durable execution as a platform primitive: every workflow step is persisted, completed steps never re-execute, and a workflow can be deterministically replayed from its event history after a crash.

**Connection.** This is condition 4 of the [durable daemons](https://blog.hackspree.com/#durable-daemons-definition) pattern — crash-proof execution — made operational. The [execution post](https://blog.hackspree.com/#durable-daemons-execution) made the stakes concrete: a trading daemon that checkpoints before placing an order and after the confirmation resumes on recovery with no reconciliation; "a duplicate order loses capital. A duplicate email is embarrassing." The snapshot is also the answer to the always-on survey's rollback stage: the agentic equivalent of a database transaction, which is exactly what DBOS (1-2 ms step latency, a single Postgres write) and Temporal provide.

**Tradeoff.** Snapshots are only as good as their granularity: a snapshot of the *conversation* does not capture the *world* — an email sent, a payment moved, a file written outside the boundary. The [system boundary](https://blog.hackspree.com/#spatiotemporal-composability) argument applies with double force: recovery is promised exactly as wide as the reification. External side effects still require [idempotency keys](https://blog.hackspree.com/#durable-daemons-execution), because replay will re-fire them and only the external system can deduplicate. And the rollback itself can be the failure: "recovery guarantees clean removal, but removal still has to be invoked."

### Pattern 8 — Tiered Hierarchical Memory

**Description.** Divides storage into immediate short-term context, scratchpad workspace, and long-term historical database storage.

**Reference.** [Lilian Weng — LLM Powered Autonomous Agents](https://lilianweng.github.io/posts/2023-06-23-agent/) (Lil'Log) — the canonical agent architecture essay: short-term memory is in-context learning, long-term memory is an external vector store with fast retrieval, and the agent is planning + memory + tool use.

**Connection.** The tiered memory is the substrate the [always-on agents survey](https://blog.hackspree.com/#always-on-agents) asked for: task ledgers, permissions, commitments, provenance, and trigger conditions are state types with different lifetimes, and the tiered pattern is how a harness gives each a home. It is also where the governance gap shows: the survey's six axes — authority, scope, mutability, provenance, recoverability, actionability — are exactly the questions a tiered design must answer per tier, and the answer is usually "not yet." Modern frameworks formalize the tiers as thread state (short-term) plus a store (long-term), with semantic, episodic, and procedural memory as distinct stores ([LangGraph Memory overview](https://docs.langchain.com/oss/python/langgraph/memory)).

**Tradeoff.** More tiers means more consistency work: what lives in the scratchpad that should have been promoted to long-term, what was deleted from short-term that the long-term still references, and — the hard one — what must be *forgotten* from all tiers including cached contexts and fine-tuned weights. Tiering also fights the [append-only discipline](https://blog.hackspree.com/#deepseek-harness): a tier that rewrites itself is a tier where the past is negotiable. And the memory is only as good as its index: recall failures are silent, which is why the [goldfish amnesia](https://blog.hackspree.com/#always-on-agents) anti-pattern is the tiered pattern's most common degeneration.

## Tool Binding & Interfacing Patterns

The tool family is where the harness meets the agent's hands. The blog's prior argument in [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) was that the interface is a performance variable — SWE-agent's custom ACI achieved 12.5% pass@1 on SWE-bench with the same GPT-4 model, and Terminal-Bench 2.0 found frontier agents under 65% on terminal tasks with the error analysis blaming the interfaces. The emerging patterns below are the binding layer: how a tool call becomes a typed, scheduled, testable, discoverable thing.

### Pattern 9 — Schema Enforcement & Self-Correction

**Description.** Forces raw LLM text into JSON Schema, catching parsing failures and feeding structural fixes back internally.

**Reference.** [Instructor — Python Library (Pydantic-based LLM Tool Interfacing)](https://python.useinstructor.com/) — the model's output is validated against a Pydantic schema, and on failure the validation error is fed back to the model as part of the retry loop, with `max_retries` and `token_budget` bounding the correction cost.

**Connection.** This is the [agentic-first contract](https://blog.hackspree.com/#agentic-first-cli-design) applied on the way back in: the model's tool calls and responses are the same contract as `--json` and exit codes, and a malformed call is the same bug class as a lie in `--help`. The pattern is also the structured-output companion to the [frozen request](https://blog.hackspree.com/#deepseek-harness): if the bill is the assertion, the schema is the grammar — validation failures become part of the log, which is exactly what the append-only model wants, because the correction is an append, not a rewrite.

**Tradeoff.** Schema enforcement costs tokens: every failed parse is a retry, and retries are where the bill grows (which is why Instructor's `token_budget` — a cumulative cap on validation retries — exists). Strict schemas over-constrain open-ended output, and loose schemas under-catch. And the retry loop is only as good as the error message: a validation error that does not tell the model *which field* failed produces the same failure again — the [silent crash](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) anti-pattern with extra steps.

### Pattern 10 — Asynchronous Tool Worker Queue

**Description.** Offloads long-running processes to background task workers and hands a tracking ID to the looping agent.

**Reference.** [Celery — Distributed Task Queue](https://docs.celeryq.dev/en/stable/getting-started/introduction.html) — the production standard for background task execution: the caller gets a task ID immediately and polls or receives a callback when the work is done.

**Connection.** This is how a harness stops an agent from blocking its own loop on a tool that takes minutes. It composes with the [tool pipeline](https://blog.hackspree.com/#deepseek-harness) — long-running tools are the natural candidates for the around-dispatch timeout/retry/metrics wrappers — and with the [durable daemons](https://blog.hackspree.com/#durable-daemons-execution) choreography: a worker queue is a daemon that satisfies all four conditions, and the tracking ID is the shared state another daemon observes. The pattern is also the practical answer to context budgeting: the agent holds a tracking ID, not the worker's output, so the output does not occupy the context window until it is ready.

**Tradeoff.** Asynchrony introduces the state-consistency problems of the [race condition](https://blog.hackspree.com/#durable-daemons-execution) anti-pattern: the agent may finish before the worker does, may poll the wrong ID, or may be torn down while the worker is still running — which is why cancellation must be a contract (the DeepSeek distinction between `ABORTED_BEFORE_DISPATCH` and `ABORTED`, so "cancellation never abandons the body"). The queue also needs exactly-once or idempotent workers, because a retried task can double-execute. And every queue is infrastructure: brokers, workers, visibility timeouts, and dead-letter handling that the harness team owns forever.

### Pattern 11 — Mock Tool Virtualization

**Description.** Swaps production APIs out for lightweight mock responses inside development environments during multi-agent unit testing routines.

**Reference.** [VCR.py — HTTP Request Mocking Library](https://vcrpy.readthedocs.io/) — record-and-replay HTTP interactions: the first real call is recorded to cassette, and subsequent runs replay the recorded response without touching the network.

**Connection.** This is the harness-engineering principle of [isolating external dependencies behind controllable fakes](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) made concrete, and it is the enabling condition for [deterministic replay](https://blog.hackspree.com/#data-driven-design-swe-agents): a harness that cannot reproduce a run cannot measure a change. Mock virtualization is also how the [tasks that fight back](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) requirement is squared with CI budgets — real repos, real browsers, and real APIs are expensive at the sample sizes data-driven design demands, so the harness runs against high-fidelity recordings in the fast loop and escalates to real infrastructure for the slow loop.

**Tradeoff.** A mock that drifts from production teaches the agent the wrong lessons — "if every test can be passed by pattern-matching the prompt, you are not measuring the assistant; you are measuring prompt luck" — which is why the cassette must be re-recorded on a schedule and why the [agent harnesses need real environments](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) for the final gate. Mocks hide latency, nondeterminism, and rate limits, and an agent trained only against mocks fails the first time it meets a 429. The pattern's honesty rule: mock for the unit test, record for the integration test, real for the eval.

### Pattern 12 — Dynamic Tool Discovery / Registry

**Description.** Stores tool specs in databases, matching and surfacing capabilities dynamically to the agent based on semantic text queries.

**Reference.** [Toolformer — Language Models Can Teach Themselves to Use Tools](https://arxiv.org/abs/2302.04761) (Meta AI) — models learn which tool to invoke and how, via self-supervised insertion of API calls; the registry framing generalizes the paper's core claim that tool selection can be taught and scored rather than hard-coded.

**Connection.** The registry is the [capability seam](https://blog.hackspree.com/#deepseek-harness) made discoverable: DeepSeek's three roles — Service Definition, Providers, Consumer — behind a stable context key, with a provider advertising its start-time capabilities statically so a request needing one it lacks is rejected before any run exists. The modern protocol form is MCP: [Model Context Protocol](https://modelcontextprotocol.io/docs/learn/architecture) standardizes how tools declare themselves, and the harness serves the catalog to the agent mid-session — DeepSeek's `cordis_inspect` tool is exactly a registry query. The blog's [type-graph mirror](https://blog.hackspree.com/#deepseek-harness) argument applies: the registry is only trustworthy if the tool specs cannot drift from the tool implementations.

**Tradeoff.** A registry solves discovery and creates a new attack and failure surface: every registered tool is a target for [prompt injection through tool descriptions](https://genai.owasp.org/llm-top-10/), and a registry that surfaces too many tools produces the [bloated utility belt](https://blog.hackspree.com/#agentic-first-cli-design) anti-pattern. Dynamic discovery also fights prefix stability — tool lists that reorder by relevance invalidate the cache prefix (the DeepSeek cache trio includes "sort tool descriptions in one fixed order"). The Toolformer lesson is the boundary: the model should learn *which* tool, not the harness should trust it with *all* tools. Discovery and authorization are separate seams.

## Orchestration & Topology Patterns

The orchestration family is where agents multiply. The blog's prior treatments are [Loop Engineering](https://blog.hackspree.com/#loop-engineering) (the loop as the unit of design), [mob programming remastered](https://blog.hackspree.com/#mob-programming-reimagined) (many minds, one surface), and the [durable daemons choreography](https://blog.hackspree.com/#durable-daemons-execution) (no orchestrator, shared state). The emerging patterns below are the four shapes the industry has actually shipped: boss-worker, blackboard, pipeline, and ensemble.

### Pattern 13 — Orchestrator-Worker (Boss-Worker)

**Description.** Directs traffic using a highly capable central agent that delegates atomic sub-tasks to smaller, faster, specialized agents.

**Reference.** [Microsoft AutoGen — Multi-Agent Conversation Framework](https://arxiv.org/abs/2308.08155) — agents as conversable, customizable components that can be composed into flexible conversation patterns, with the orchestrator delegating work to specialist agents that may pair LLMs with human inputs and tools.

**Connection.** The orchestrator-worker shape is the [harness building the harness](https://blog.hackspree.com/#deepseek-harness) at runtime: one agent that curates a delegation tree instead of drowning in context. It connects to DeepSeek's [workflow seam](https://blog.hackspree.com/#deepseek-harness) — the model writes an orchestration script with `agent()` calls that spawn subagents under caps — and to the [subagent registry](https://blog.hackspree.com/#deepseek-harness) with providers mounted by name. It is also the distillation pattern's topology: a swarm of [tiny specialists coordinated by a router](https://blog.hackspree.com/#agents-are-distillation-at-scale). The Anthropic long-running harness is the strongest recent instance: a planner expands the prompt into a spec, a generator works one feature at a time, and the two negotiate what "done" means before code is written ([Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Tradeoff.** The orchestrator is a single point of failure and a bottleneck: if it mis-delegates, the error cascades downstream (Anthropic deliberately kept the planner's spec high-level because "if the planner tried to specify granular technical details upfront and got something wrong, the errors in the spec would cascade"). Every delegation is a context handoff that can lose state — the [goldfish amnesia](https://blog.hackspree.com/#always-on-agents) risk — and every worker is a new surface for the [committee paradox](https://blog.hackspree.com/#buzz-block-agents) if the delegation loop has no termination condition. The boss-worker pattern needs the exit condition to be part of the harness, not the conversation.

### Pattern 14 — Blackboard (Shared Workspace)

**Description.** Connects disjointed agents to a unified state schema where the harness coordinates simultaneous data writes and events.

**Reference.** [CrewAI — Multi-Agent Coordination Framework](https://docs.crewai.com/en/concepts/crews) — crews of role-specialized agents coordinate through shared state and task dependencies, and its Flows layer provides event-driven, state-managed workflow execution.

**Connection.** The blackboard is one of the oldest ideas in AI — the Hearsay-II speech-understanding architecture (1970s) — and the agent era reinvented it. This blog's prior treatment is the [durable daemons choreography](https://blog.hackspree.com/#durable-daemons-execution): "Each daemon is a process. Each daemon is durable. The choreography is the composition. No orchestrator. Just state." The blackboard is the shared-state composition made explicit, and the [Buzz/GitButler layer](https://blog.hackspree.com/#buzz-block-agents) is its infrastructure face: virtual branches let multiple agents write the same repository simultaneously, with conflicts detected at write time rather than merge time. The harness's job in the blackboard pattern is exactly what CrewAI's Flows formalize: coordinating which agent may write what, and routing events to listeners.

**Tradeoff.** Shared state is shared risk. The blackboard concentrates the [race condition](https://blog.hackspree.com/#durable-daemons-execution) anti-pattern: simultaneous writes without transactional locks corrupt the state, and the durability of the board decides whether the corruption is recoverable. The blackboard also fights [per-agent scope](https://blog.hackspree.com/#deepseek-harness) — the tension this blog flagged between Pattern 7 (isolation) and choreography (shared state) — and the resolution is explicit scope discipline: which state is shared, which is scoped, who owns the boundary. A blackboard without an owner is a tragedy of the commons for state.

### Pattern 15 — Sequential Pipeline Routing

**Description.** Passes analytical payloads through rigid linear stages, using LLMs solely for classification or transformation tasks at each hop.

**Reference.** [LangChain Chains — Architectural Pattern](https://python.langchain.com/docs/versions/migrating_chains/overview/) — chains compose LLM calls and deterministic logic into fixed sequences, and the migration-era documentation is explicit about when the rigid linear form is the right call and when it is not.

**Connection.** The pipeline is the [loop engineering](https://blog.hackspree.com/#loop-engineering) answer to "the most successful implementations use simple, composable patterns rather than complex frameworks" (Anthropic's *Building Effective Agents*): a rigid linear stage where each hop has one job and a deterministic contract is the easiest harness to verify, replay, and bill. It is also the [rule-engine](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) lineage — LLMs used for classification or transformation at defined seams, with the deterministic machinery doing the routing. This blog's [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) discipline applies per hop: each stage is a tool with a `--json` contract, exit codes, and honest `--help`.

**Tradeoff.** Rigid pipelines cannot route around failure: a stage that misclassifies poisons every downstream stage, and there is no re-planning. Pipelines also serialize — the whole flow runs at the speed of its slowest LLM hop, which is why the pattern's cost multiplies with the [token economics](https://blog.hackspree.com/#every-token-has-a-price-tag) (each hop is a full request; prefix caching helps only if the prefix is stable). The pattern's own docs warn about its limits — chains are right when the flow is genuinely linear, wrong when the flow needs to branch, loop, or recover, which is when the orchestrator-worker or graph shapes win.

### Pattern 16 — Voting / Consensual Ensemble

**Description.** Queries multiple independent model setups with identical prompts, using harness code to calculate majority agreement on outputs.

**Reference.** [LMSYS Chatbot Arena — Consensus Data & Research](https://lmarena.ai/leaderboard) — the reference for agreement-based evaluation of model output: millions of pairwise human votes form the consensus signal that ranks models, establishing that *independent judgments aggregated* is a workable quality measure.

**Connection.** The ensemble is the verification pattern made statistical: where [verifiers are king](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), the ensemble is the king's court — multiple independent judgments aggregated by harness code instead of one verdict from one model. It connects to the [data-driven design](https://blog.hackspree.com/#data-driven-design-swe-agents) principle that a single run is a data point and a thousand runs is a distribution: majority agreement is a distribution over independent samples. And it is the natural judge for the [self-evaluation problem](https://blog.hackspree.com/#mob-programming-reimagined) — 79% of 25,264 agent-generated PRs were reviewed by the same developer who prompted the agent; the ensemble replaces self-review with cross-review.

**Tradeoff.** Ensembles multiply cost linearly (n independent calls for n votes) and can multiply latency; the [cost per finished task varies 7x across harnesses](https://blog.hackspree.com/#deepseek-harness) before you add an ensemble. Majority agreement is a strong signal only when the members are *independent* — correlated members (same model family, same prompt shape, same failure mode) vote as one and add nothing. And the arena reference is honest about the limit: human-vote consensus measures preference, not correctness, and an LLM-judge ensemble inherits every bias of its members — "algorithmic verification misses context. Agentic verification hallucinates." The ensemble is a signal, not a ground truth.

# Part II — Anti-Patterns

## Safety, Security & Isolation Anti-Patterns

### Anti-Pattern 1 — The Naked Prompt (Implicit Trust)

**Description.** Giving an LLM direct API keys or terminal access without a sanitization, parsing, or proxy layer in the harness.

**Reference.** [OWASP Top 10 for LLM Applications](https://genai.owasp.org/llm-top-10/) — the relevant entries are LLM01 Prompt Injection, LLM02 Sensitive Information Disclosure, and LLM07 System Prompt Leakage in the 2025 numbering (LLM06 Sensitive Information Disclosure in the 2023/24 numbering, which is where the classic citation comes from).

**The failure.** The model is not an application boundary; it is an untrusted input processor. A naked prompt hands the model the keys and asks it to be careful, which is the [Prompt-Driven Authorization](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) failure at its most literal. The fix is the pattern half of this family: the [gatekeeper](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface), the [sandbox wrapper](https://blog.hackspree.com/#sandboxing-ai-agents), and credentials handled by a [transparent secrets proxy](https://blog.hackspree.com/#sandboxing-ai-agents) so the agent never sees them. The xz lesson from [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) applies: the tool that executes arbitrary code with your credentials is the highest-value target in your supply chain, and "an agent with a thousand dependencies is a thousand doors."

### Anti-Pattern 2 — The Infinite Execution Vortex (Uncapped Loops)

**Description.** Failing to program hard timeout, token, or monetary ceilings in the harness, letting an agent retry a broken tool loop indefinitely.

**Reference.** [AutoGPT Issue Tracker — Infinite Loop & Token Drain Reports](https://github.com/Significant-Gravitas/AutoGPT/issues) — the first autonomous agent's open issue tracker is a museum of the failure: agents that loop on the same failing step, burning tokens and wall-clock until a human intervenes.

**The failure.** The vortex is the [Token & Time Budget Throttler](https://blog.hackspree.com/#deepseek-harness) pattern's absence made visible, and it is an economic failure before it is a technical one: an uncapped loop is the [tragedy of the commons](https://blog.hackspree.com/#every-token-has-a-price-tag) staged inside one agent run, and the meter is what makes any mechanism possible. The fix is structural, not prompt-based: hard ceilings enforced by the harness ([LLM10:2025 Unbounded Consumption](https://genai.owasp.org/llm-top-10/)), loop-iteration limits, and the [runaway-loop guard](https://blog.hackspree.com/#deepseek-harness) with its honest caveat — reminders are not enforcement, so the ceiling must be enforced in code.

### Anti-Pattern 3 — Prompt-Driven Authorization

**Description.** Relying on instructions in the system prompt (e.g., "Do not delete user data") for security, instead of hard-coding permission checks into the tool logic.

**Reference.** [Simon Willison — Prompt Injection series](https://simonwillison.net/series/prompt-injection/) — the running documentation of the threat model: instructions are data, prompt injection is the mechanism, and "you can't solve AI security problems with more AI."

**The failure.** A system prompt is not an access control list; it is a document the model may be instructed to ignore. The pattern's doctrine is the opposite of everything in the safety family: authorization must be [monotonic, structural, and fail-closed](https://blog.hackspree.com/#deepseek-harness) — the tool denies by default, and no amount of persuasion in the prompt can force-allow. The DeepSeek formulation is the standard: monotonic guards "deny or abstain and can never force-allow, so owner policy that must not be reordered cannot be argued around." The always-on survey's authority axis says the same thing: who can modify the agent's state is a property of the harness, not of the model's reading comprehension.

## State, Memory & Context Anti-Patterns

### Anti-Pattern 4 — The Context Avalanche (Memory Dumping)

**Description.** Shoveling every raw log, conversation history, and tool output directly into the chat history until the LLM context fills up and degrades reasoning.

**Reference.** [Lost in the Middle: How Language Models Use Long Contexts](https://arxiv.org/abs/2307.03172) (Liu et al., Stanford) — model performance degrades significantly when relevant information sits in the middle of a long input context; performance is highest at the beginning and end.

**The failure.** The avalanche is the [Rolling Window Compression](https://blog.hackspree.com/#deepseek-harness) pattern's absence: raw logs are not context, and a filled context is not a well-informed agent — it is a degraded one that reads the middle worst exactly when the middle holds the answer. The fix is the [append-only log with a logged-surface invariant](https://blog.hackspree.com/#deepseek-harness): the model sees a *derived* view — 44 event types in DeepSeek's log, exactly three visible to the model — never the raw stream. "The system never stores the conversation it sends; it recomputes the model's context from the log."

### Anti-Pattern 5 — Goldfish Amnesia (State Blindness)

**Description.** Running a multi-turn agent loop completely statelessly, causing the agent to forget its high-level goal and repeat identical tool calls.

**Reference.** [LangGraph — Persistent State Architecture](https://docs.langchain.com/oss/python/langgraph/memory) — thread-scoped persistence as the primitive that makes multi-turn agents remember; statelessness is documented as the thing persistence exists to prevent.

**The failure.** A stateless loop is not an agent; it is a request-response function with a longer prompt. The always-on survey's finding is the diagnosis: agents "accumulate state aggressively and have almost no machinery for unwinding it," and the opposite failure — no state at all — repeats the same tool call because nothing records that it already ran. The fix is the [state lifecycle](https://blog.hackspree.com/#always-on-agents) made real: write, validate, retrieve, update, forget — with the ledger of what was tried as the minimum viable memory. The durable daemons' condition 2 (stateful memory) and condition 3 (autonomous action) are the specification: "A chatbot fails all three. An always-on agent satisfies them."

### Anti-Pattern 6 — The RAG Firehose

**Description.** Injecting top-K vector database chunks into the prompt blindly based on raw keyword match, drowning out the actual user instruction.

**Reference.** [Pinecone — Advanced RAG: Chunking & Metadata Filtering](https://www.pinecone.io/learn/advanced-rag/) — retrieval quality as a function of chunking strategy, metadata filtering, and reranking, not raw top-K volume.

**The failure.** The firehose is the [Semantic Memory Router](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) pattern's absence: retrieval without judgment. The agent's actual instruction competes with fifty chunks of loosely-related text, and [Lost in the Middle](https://blog.hackspree.com/#agentic-first-cli-design) predicts the outcome — the instruction is in the middle, the chunks drown it, performance degrades. The fix is routing discipline: retrieved context is a first-class observation, selected by the harness with chunking, metadata filters, and reranking, and injected where it cannot bury the instruction. And every injected chunk is untrusted input — the firehose is also the injection vector for [LLM08:2025 Vector and Embedding Weaknesses](https://genai.owasp.org/llm-top-10/).

## Tool Binding & Interfacing Anti-Patterns

### Anti-Pattern 7 — The Silent Crash (Exception Swallowing)

**Description.** Catching API errors in the harness background but returning blank or generic strings to the agent, forcing it to hallucinate success.

**Reference.** [Instructor — Error Handling & Retry Mechanics](https://python.useinstructor.com/concepts/retrying/) — the counter-pattern's fix is documented as the retry loop: validation failures and API errors are surfaced to the model with the details needed to correct, bounded by `max_retries` and `token_budget`.

**The failure.** The silent crash converts a real error into a hallucinated success: the agent receives `""` or `"ok"`, and because it cannot see the error, it confidently proceeds on a false premise. This is the [interface as contract](https://blog.hackspree.com/#agentic-first-cli-design) argument inverted — the three channels (stdout data, stderr diagnostics, exit code verdict) exist precisely so failures are legible, and a harness that swallows them destroys the contract. The fix is the layered-check discipline from [harness engineering](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents): failures must be legible at the tool-selection layer, the state-transition layer, and the outcome layer — "the planner chose the wrong tool" instead of "the agent failed."

### Anti-Pattern 8 — Bloated Utility Belt (Tool Over-Provisioning)

**Description.** Giving a single agent dozens of complex tools at once, which confuses the model's reasoning and triggers incorrect tool selection.

**Reference.** [Toolformer — Scaling Down Tool Cardinality](https://arxiv.org/abs/2302.04761) (Meta AI) — the research line shows tool use degrades as the tool space grows unmanageable; the interface-design literature (SWE-agent's ACI) reached the same conclusion independently: fewer, better-shaped tools beat more tools.

**The failure.** Every tool in the prompt is a decision the model must make, and a bloated registry turns tool selection into a [needle-in-a-haystack retrieval problem](https://blog.hackspree.com/#agentic-first-cli-design) that models lose. The fix is the [capability seam](https://blog.hackspree.com/#deepseek-harness): split only when roles evolve independently, present only what the task needs (DeepSeek's presets are per-session tool compositions, and code mode replaces the tool list with a generated SDK entirely), and let the [registry](https://blog.hackspree.com/#deepseek-harness) discover on demand instead of dumping everything into the prompt. The SWE-agent result is the anchor: with the same model, a minimal well-designed ACI more than doubled state-of-the-art.

### Anti-Pattern 9 — The Schema Free-for-All

**Description.** Accepting unstructured text strings back from the LLM for complex arguments, leading to downstream software failures when values are malformed.

**Reference.** [Pydantic — Core Validation Reference](https://docs.pydantic.dev/) — the reference for typed, validated data structures in Python; the entire library exists because unvalidated strings become runtime failures one layer down.

**The failure.** The free-for-all moves the parsing problem downstream where there is no model to correct it: a malformed date string fails in a database insert, a truncated JSON fails in a deserializer, and the error surfaces far from the agent that produced it — the [silent crash](https://blog.hackspree.com/#deepseek-harness) with more steps. The fix is [Schema Enforcement & Self-Correction](https://blog.hackspree.com/#agentic-first-cli-design): force the output into a schema at the harness boundary, feed validation errors back while the model is still in the loop, and let the append-only log record the corrections. The [agentic-first contract](https://blog.hackspree.com/#deepseek-harness) is the specification: stable, versioned, additive schemas, and honest failures when the schema cannot be satisfied.

## Orchestration & Topology Anti-Patterns

### Anti-Pattern 10 — The Committee Paradox (Infinite Multi-Agent Debates)

**Description.** Designing two or more open-ended autonomous agents that review each other's work without a rigid harness-driven exit condition, causing eternal arguments.

**Reference.** [Microsoft AutoGen — Multi-Agent Conversation Framework](https://arxiv.org/abs/2308.08155) — the framework's paper and documentation treat termination as a first-class design concern of multi-agent conversation; the community's infinite-loop reports are the empirical record of what happens when it is not.

**The failure.** Two agents reviewing each other with no exit condition is an [infinite loop with a nicer name](https://blog.hackspree.com/#loop-engineering). The committee needs the harness to define *when it is done* — a termination condition, a verdict threshold, a budget, or a human breakpoint — and the [Voting / Consensual Ensemble](https://blog.hackspree.com/#buzz-block-agents) pattern shows the shape: disagreement is resolved by an aggregation rule, not by talking longer. The self-review data from the mob post ([79% of agent PRs reviewed by the prompting developer](https://blog.hackspree.com/#mob-programming-reimagined)) is the failure's milder cousin: review without independence is theater, and review without termination is waste.

### Anti-Pattern 11 — The God Agent (Monolithic Orchestration)

**Description.** Building one massive agent with an immense prompt to manage every phase of a business workflow, instead of decoupling it into specialized workers.

**Reference.** [CrewAI — Architecture Guides: Decomposing Monoliths into Crews](https://docs.crewai.com/en/concepts/crews) — the framework's core argument for crews is that role-specialized agents with focused prompts beat one agent doing everything.

**The failure.** The god agent is the [Bloated Utility Belt](https://blog.hackspree.com/#agentic-first-cli-design) at the orchestration scale: one context window holding every phase of the workflow, every tool, and every rule, with the [context avalanche](https://blog.hackspree.com/#deepseek-harness) and [lost-in-the-middle](https://blog.hackspree.com/#agentic-first-cli-design) degradation as the inevitable consequences. The fix is [decomposition with handoff](https://blog.hackspree.com/#durable-daemons-execution): the [orchestrator-worker](https://blog.hackspree.com/#deepseek-harness) shape, one feature at a time, with structured artifacts carrying state between sessions — the Anthropic harness's three-agent architecture is the modern template (planner, generator, evaluator), and its planner deliberately stays high-level because "errors in the spec would cascade into the downstream implementation."

### Anti-Pattern 12 — State Race Conditions

**Description.** Allowing asynchronous multi-agent setups to write to a shared memory space without harness-level transactional locks, corrupting the system state.

**Reference.** [Temporal — Distributed Design & Durable Execution](https://docs.temporal.io/) — the platform's documentation on event-sourced, deterministic execution is the standard answer to the race: state changes go through the workflow runtime, not through unsynchronized writes.

**The failure.** The race is the [Blackboard](https://blog.hackspree.com/#durable-daemons-execution) pattern without its harness: two agents writing the same ledger, one overwriting the other's checkpoint, the state corrupting silently. The fix is the durability discipline from the [durable daemons execution post](https://blog.hackspree.com/#durable-daemons-execution): writes through a single ordered log, exactly-once within the control boundary, idempotency keys for external effects, and choreography through shared state with clear ownership — "no RPC. No message bus. No central orchestrator" is only safe when the shared state itself is the coordination mechanism, and the coordination mechanism must be transactional. The Buzz/GitButler layer shows the version-control face: virtual branches detect conflicts at write time, not merge time.

# Patterns on the 2026 Frontier

The four families above are the settled vocabulary — the patterns frameworks document and teams adopt. The frontier is what is emerging in 2026 and is not yet in any stable documentation: the shapes that harness teams at the frontier are converging on, and that this blog's archive anticipated. Five are worth naming.

### Frontier 1 — The Generator–Evaluator Loop

The most consequential harness result of early 2026 is Anthropic's [Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps) (March 2026): a GAN-inspired architecture that separates the agent doing the work from the agent judging it. The founding observation is the self-evaluation failure: "When asked to evaluate work they've produced, agents tend to respond by confidently praising the work — even when, to a human observer, the quality is obviously mediocre." The fix is structural: a generator agent creates, a standalone evaluator grades against explicit criteria, and the critique flows back as the next iteration's input. The frontend experiment ran 5-15 iterations per generation, with the evaluator navigating the live page via Playwright MCP before scoring; the full-stack version ran a planner, a generator, and an evaluator for six hours at $200 — against a solo run of 20 minutes at $9 — and produced an application the solo run could not approach. The tradeoff is the cost: the harness was "over 20x more expensive," and the pattern is only worth it when the output quality difference justifies the bill. This is the [verifiers are king](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) argument with the verifier given agency: not a gate that blocks, but a critic that iterates.

### Frontier 2 — Sprint Contracts

The same harness introduced a second pattern worth stealing: before each sprint, the generator and evaluator "negotiated a sprint contract: agreeing on what 'done' looked like for that chunk of work before any code was written." The generator proposed what it would build and how success would be verified; the evaluator reviewed the proposal; the two iterated until they agreed; only then did the generator build against the contract. The pattern converts an under-specified spec into testable acceptance criteria at the moment of maximum leverage — before the work exists — and it is the [tasks that fight back](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) principle applied to the contract itself: the agent defines the verifier before it defines the artifact. The tradeoff is the negotiation overhead — two LLM agents spending tokens agreeing on "done" before any code is written — which is why the contract is bounded to a sprint-sized chunk, not the whole project.

### Frontier 3 — Context Resets with Handoff Artifacts

Anthropic's harness work distinguished two answers to the filling context window: *compaction*, which summarizes in place so the same agent continues on a shortened history, and *context resets*, which clear the context entirely and hand state to a fresh agent through a structured artifact. The reset is the stronger move, and it exists because of "context anxiety": "models tend to lose coherence on lengthy tasks as the context window fills," and some exhibit premature wrap-up ("beginning to wrap up work prematurely as they approach what they believe is their context limit") — Sonnet 4.5 exhibited it strongly enough that compaction alone was insufficient. The reset provides "a clean slate, at the cost of the handoff artifact having enough state for the next agent to pick up the work cleanly" — plus orchestration complexity, token overhead, and latency. This is the [append-only log](https://blog.hackspree.com/#deepseek-harness) argument at the session boundary: the log never rewrites, and the handoff artifact is a new prefix, not an edit. The frontier tradeoff is now visible in model generations — Opus 4.5 "largely removed that behavior on its own," so the harness could drop resets entirely; the pattern is a function of the model, not just the harness.

### Frontier 4 — The Interop Layer: MCP, ACP, AGENTS.md

The tool-seam and instruction-seam standardization is the quietest frontier and the one with the most leverage. [MCP](https://modelcontextprotocol.io/docs/learn/architecture) standardized how tools declare and execute; [ACP](https://agentclientprotocol.com/) (Agent Client Protocol) standardizes how editors and agents connect, so "each editor must build custom integrations for every agent" stops being the default; and [AGENTS.md](https://agents.md/) — "a README for agents," used by over 60,000 open-source projects — standardizes where agent instructions live. The blog's DeepSeek teardown called the strategy [interop-as-product-strategy](https://blog.hackspree.com/#deepseek-harness): "when a vendor adopts its rivals' formats, your files become portable, whichever harness you run." The frontier pattern is the reverse: when the *community* standardizes the seam, the harness that implements the standard inherits the ecosystem — DeepSeek runs unmodified Claude Code hooks, and the [subagent seam](https://blog.hackspree.com/#deepseek-harness) mounts rival harnesses as providers rather than reimplementing them. The tradeoff is the same one the seam pattern always carries: every protocol you adopt is a contract you do not control, and every standard is a boundary where a substitution can hide.

### Frontier 5 — Live-Environment Evaluators

The evaluator in Anthropic's harness was not a static grader: it was given the Playwright MCP, "which let it interact with the live page directly before scoring each criterion and writing a detailed critique" — navigating the app, clicking through UI features, testing API endpoints and database states "the way a user would." This is the frontier version of the blog's [harnesses should observe the whole system](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) principle: the verifier does not read the artifact, it *uses* it. The tradeoff is wall-clock time — "because the evaluator was actively navigating the page rather than scoring a static screenshot, each cycle took real wall-clock time. Full runs stretched up to four hours" — which is why the pattern belongs in the slow loop of the [eval → task → agent harness taxonomy](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents), reserved for the final gate. The frontier insight is that live evaluation catches what static verification cannot: the solo run's application "looked impressive but still had real bugs when you actually tried to use them," and only a verifier with hands found them.

# How the patterns compose — and where they conflict

The catalog is not a menu; the patterns interact, and some of the interactions are contradictions that have to be managed explicitly.

**The memory family contains its own war.** The [append-only log](https://blog.hackspree.com/#deepseek-harness) (Pattern 1 of the first edition), [rolling window compression](https://blog.hackspree.com/#deepseek-harness), and [tiered hierarchical memory](https://blog.hackspree.com/#always-on-agents) disagree about whether the past can be rewritten. The resolution is layering: the log is immutable, the derived view is compressed, the tiers are projections — and the [compaction prefix bug](https://blog.hackspree.com/#deepseek-harness) is what happens when the layers touch.

**Prefix-cache discipline constrains every family.** The [Bill as Assertion](https://blog.hackspree.com/#every-token-has-a-price-tag) requirement — no volatile data in the prompt prefix, fixed tool ordering, append-only history — constrains the [Semantic Memory Router](https://blog.hackspree.com/#deepseek-harness) (inject context where it cannot disturb the prefix), the [Dynamic Tool Registry](https://blog.hackspree.com/#deepseek-harness) (discovery must not reorder the tool list), and the [Rolling Window Compression](https://blog.hackspree.com/#deepseek-harness) (summaries must be prefix-extensions). The interface discipline of the [agentic-first CLI](https://blog.hackspree.com/#agentic-first-cli-design) — "no timestamps unless asked" — is a billing requirement before it is a UX nicety.

**Isolation and orchestration pull against each other.** The [Ephemeral Sandbox Wrapper](https://blog.hackspree.com/#sandboxing-ai-agents) and [per-agent scope](https://blog.hackspree.com/#deepseek-harness) want isolation; the [Blackboard](https://blog.hackspree.com/#durable-daemons-execution) and the [durable daemons choreography](https://blog.hackspree.com/#durable-daemons-definition) want shared state as the coordination mechanism. The resolution is explicit scope discipline — which state is shared, which is scoped, who owns the boundary — the least-solved governance question in the [always-on survey](https://blog.hackspree.com/#always-on-agents).

**Safety and speed are the same budget.** The [HITL Breakpoint](https://blog.hackspree.com/#deepseek-harness), the [Static Gatekeeper](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface), and the [Token & Time Budget Throttler](https://blog.hackspree.com/#every-token-has-a-price-tag) each add latency and tokens to every step; the [verification layer](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) adds more. The token post's arithmetic — the bill is 2-14% of labor cost — is the budget all of them share, and the teams that win make the checks *be* the measurement data: layered checks that fail legibly are both the gate and the instrument.

**Defense in depth disagrees with itself.** The [sandbox stack](https://blog.hackspree.com/#sandboxing-ai-agents) says add layers; [zero overhead](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) says subtract them. Both are right, and the resolution is honesty about what each layer is for: the sandbox constrains the compromised agent, the small dependency graph shrinks what can be compromised, the gatekeeper mediates capabilities — none of them is a wall, and the pattern is to know which one you are building.

The meta-pattern across all of it is the one the DeepSeek teardown's two independent analyses converged on: **the harness is the product.** The same model swinging from 47% to 67% across eight harnesses, harness-level improvements lifting scores without any training, the same weights 3.4 points apart under two harnesses — the wrapper is where the leverage is, which is why the wrapper deserves a pattern catalog, and an anti-pattern catalog, at all.

# What to steal regardless

The catalog above is drawn from this blog's archive and from primary sources verified for this edition. The portable items, independent of any particular stack:

**From the pattern half:** the ephemeral sandbox wrapper with honest teardown (P1); the intercepting gatekeeper with monotonic, fail-closed denial (P2); the breakpoint as a persisted state primitive, not a prompt (P3); hard token/time/monetary ceilings enforced in code, with the bill as a CI assertion (P4); compression as a projection over an immutable log (P5); retrieval with routing judgment, injected where it cannot bury the instruction (P6); checkpoints with idempotency keys at the boundary (P7); tiers with explicit authority, scope, and forgetting per tier (P8); schema enforcement with bounded self-correction (P9); async workers with cancellation as a contract (P10); mocks for the fast loop and real environments for the final gate (P11); registries whose specs cannot drift from implementations (P12); delegation with harness-defined termination (P13); shared state with transactional ownership (P14); linear stages with deterministic contracts (P15); and ensembles of independent judgments as a signal, never a ground truth (P16).

**From the anti-pattern half:** the twelve names above are a diagnostic checklist. When an agent behaves inexplicably, ask which anti-pattern the harness is committing: Is the model trusted with keys and told to be careful (A1)? Is there no ceiling on the loop (A2)? Is authorization a sentence in the prompt instead of a check in the tool (A3)? Is the context a dump instead of a derivation (A4)? Is the loop stateless (A5)? Is retrieval drowning the instruction (A6)? Are errors being swallowed into hallucinated success (A7)? Is the tool belt bloated (A8)? Are complex arguments untyped (A9)? Are agents debating without an exit condition (A10)? Is one agent doing everything (A11)? Is shared state being written without locks (A12)?

**From the frontier:** separate the generator from the evaluator, because self-evaluation is reliably lenient (F1); negotiate the contract before the code, because "done" agreed in advance beats "done" discovered after (F2); reset the context and hand off structured artifacts instead of compacting in place when the model exhibits context anxiety (F3); implement the interop standards, because the seam is the ecosystem (F4); and give the verifier hands, because the artifact that looks impressive is not the artifact that works (F5).

The tradeoffs are not defects in the patterns; they are the prices. Every harness pattern in this catalog buys a guarantee and charges a cost, and the honest harness engineer is the one who can state both for each layer of the stack. The guarantee stops where the declaration stops — file effects only, bash-equivalent trust, observational equivalence, within the control boundary, majority agreement is not ground truth. Name the boundary, design within it, and verify the design works. That is the engineering method, and it is the whole catalog in one sentence.

# References

## Primary sources for this edition's patterns and anti-patterns

- [LangChain — Sandbox integrations](https://python.langchain.com/docs/integrations/) — sandboxing as a first-class integration component (Pattern 1).
- Meta AI. [Llama Guard: LLM-based Input-Output Safeguard for Human-AI Conversations](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/) (December 2023) — the intercepting gatekeeper (Pattern 2).
- [LangGraph — Interrupts (state management & breakpoints)](https://docs.langchain.com/oss/python/langgraph/interrupts) — HITL breakpoints (Pattern 3); [LangGraph — Memory overview](https://docs.langchain.com/oss/python/langgraph/memory) — persistent state (Anti-Pattern 5).
- [AutoGPT — Classic configuration: maximum loop & cost control](https://docs.agpt.co/classic/configuration/) and [AutoGPT issue tracker](https://github.com/Significant-Gravitas/AutoGPT/issues) — budget throttling (Pattern 4) and the infinite-loop/token-drain record (Anti-Pattern 2).
- Packer, Wooders, Lin, Fang, Patil, Stoica, Gonzalez. [MemGPT: Towards LLMs as Operating Systems](https://arxiv.org/abs/2310.08560) (UC Berkeley, 2023) — virtual context management (Pattern 5).
- [Pinecone — Retrieval-Augmented Generation](https://www.pinecone.io/learn/retrieval-augmented-generation/) — the RAG router (Pattern 6); [Pinecone — Advanced RAG: chunking & metadata filtering](https://www.pinecone.io/learn/advanced-rag/) (Anti-Pattern 6).
- [Temporal — durable execution, workflow state & replay](https://docs.temporal.io/) — snapshots and rollback (Pattern 7), distributed-design discipline (Anti-Pattern 12).
- Weng, L. [LLM Powered Autonomous Agents](https://lilianweng.github.io/posts/2023-06-23-agent/) (Lil'Log, 2023) — tiered hierarchical memory (Pattern 8).
- [Instructor — Python library for structured LLM output](https://python.useinstructor.com/) and [retry logic with validation & token budgets](https://python.useinstructor.com/concepts/retrying/) — schema enforcement and self-correction (Pattern 9), error-handling mechanics (Anti-Pattern 7).
- [Celery — Distributed Task Queue](https://docs.celeryq.dev/en/stable/getting-started/introduction.html) — the asynchronous worker queue (Pattern 10).
- [VCR.py — HTTP request mocking / record-and-replay](https://vcrpy.readthedocs.io/) — mock tool virtualization (Pattern 11).
- Schick et al. [Toolformer: Language Models Can Teach Themselves to Use Tools](https://arxiv.org/abs/2302.04761) (Meta AI, 2023) — tool learning and cardinality (Pattern 12; Anti-Pattern 8).
- Wu, Bansal, Zhang, Wu, Li, Zhu, Jiang, Zhang, Zhang, Liu, White, Burger, Wang. [AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation](https://arxiv.org/abs/2308.08155) (Microsoft, 2023) — orchestrator-worker (Pattern 13), termination as a design concern (Anti-Pattern 10).
- [CrewAI — Crews (multi-agent coordination)](https://docs.crewai.com/en/concepts/crews) and [CrewAI — Flows (event-driven, state-managed workflows)](https://docs.crewai.com/en/concepts/flows) — the blackboard (Pattern 14), decomposing monoliths (Anti-Pattern 11).
- [LangChain — Chains (architectural pattern)](https://python.langchain.com/docs/versions/migrating_chains/overview/) — sequential pipeline routing (Pattern 15).
- [LMSYS Chatbot Arena / LMArena — consensus data & leaderboard](https://lmarena.ai/leaderboard) — voting and consensus evaluation (Pattern 16).
- [OWASP Top 10 for LLM Applications](https://genai.owasp.org/llm-top-10/) (2025 edition: LLM01 Prompt Injection, LLM02 Sensitive Information Disclosure, LLM06 Excessive Agency, LLM07 System Prompt Leakage, LLM08 Vector and Embedding Weaknesses, LLM10 Unbounded Consumption; 2023/24 edition: LLM06 Sensitive Information Disclosure) — the naked prompt (Anti-Pattern 1) and the uncapped loop (Anti-Pattern 2).
- [Simon Willison — Prompt Injection series](https://simonwillison.net/series/prompt-injection/) — prompt-driven authorization (Anti-Pattern 3).
- Liu, Lin, Hewitt, Paranjape, Bevilacqua, Petroni, Liang. [Lost in the Middle: How Language Models Use Long Contexts](https://arxiv.org/abs/2307.03172) (Stanford, 2023) — the context avalanche (Anti-Pattern 4).
- [Pydantic — core validation reference](https://docs.pydantic.dev/) — the schema free-for-all (Anti-Pattern 9).

## Frontier sources (2026)

- Rajasekaran, P. (Anthropic Labs). [Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps) (March 2026) — the generator-evaluator loop, sprint contracts, context resets vs. compaction, context anxiety, live-environment evaluators with Playwright MCP; the solo-run vs. harness-run cost comparison ($9/20 min vs. $200/6 hr).
- [Model Context Protocol — architecture](https://modelcontextprotocol.io/docs/learn/architecture) — the tool interop standard.
- [Agent Client Protocol (ACP)](https://agentclientprotocol.com/) — the editor/agent interop standard (Zed Industries, JetBrains).
- [AGENTS.md](https://agents.md/) — the agent instruction-file standard, "a README for agents," used by 60k+ open-source projects.
- Anthropic. [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents) (December 2024) — "simple, composable patterns rather than complex frameworks."
- OpenAI. [Harness Engineering: Leveraging Codex in an Agent-First World](https://openai.com/index/harness-engineering/) — the harness as the product surface for agents.

## This blog's archive (each pattern's connection is linked inline)

- [DeepSeek Harness: Everything Is a Plugin](https://blog.hackspree.com/#deepseek-harness) — append-only session log, golden rule and ~120x prefix-cache discount, tool pipeline (waterfalls, monotonic guards, fail-closed approval), capability seams, code mode, per-agent scope, four presets, subagent registry, interop-as-strategy, same-model-across-eight-harnesses benchmark, known limitations.
- [Spatiotemporal Composability: The Missing Calculus for Self-Evolving Agents](https://blog.hackspree.com/#spatiotemporal-composability) — temporal/spatial composability, revertible effects, the one-door context, the system boundary, composition vs. verification.
- [Harness Engineering: Best Practices for Reliable Agent Systems](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) — the consolidated harness canon: layered checks, pass-fail contracts, replay, incidents as baselines, tasks that fight back, real repos and browsers, mock isolation.
- [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) — Sonar AC/DC: Guide/Verify/Solve, algorithmic vs. agentic verification, the three nested loops.
- [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) — Fowler's retreat findings, agents don't learn, DSLs as the pledge(2) bridge.
- [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents) — data-driven design, the harness as measurement instrument, distributions over data points.
- [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) — SWE-agent's ACI, Terminal-Bench 2.0, structured output, determinism, exit codes as contract.
- [Durable Daemons](https://blog.hackspree.com/#durable-daemons) series — [pattern specification](https://blog.hackspree.com/#durable-daemons-definition), [runtime and implementation](https://blog.hackspree.com/#durable-daemons-execution) (Temporal, DBOS, exactly-once, idempotency keys, choreography), [limits](https://blog.hackspree.com/#durable-daemons-limits).
- [Always-On Agents: state, memory, and the governance gap](https://blog.hackspree.com/#always-on-agents) — the 435-paper survey, six diagnostic axes, the state lifecycle.
- [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents) — defense in depth, the thirteen layers, compliance is not security.
- [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) — subtraction as the only compounding supply-chain defense.
- [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag) — the Jevons paradox arithmetic, the meter, the measurement asymmetry.
- [Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale) — small specialists, the closed loop.
- [Buzz and the Identity Problem](https://blog.hackspree.com/#buzz-block-agents) — portable agent identity, GitButler virtual branches, the "who did what, why, and was it correct?" test.
- [Zuill's Mob Programming, Remastered](https://blog.hackspree.com/#mob-programming-reimagined) — the self-review problem (79% of 25,264 agent PRs), driver/navigators as review architecture.
- [Loop Engineering is what the NATO conference asked for in 1968](https://blog.hackspree.com/#loop-engineering) — loops over prompts, "the simulation becomes the system."
- [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer) — deliverables over chat, approval gates, personas as configuration.
- [The Factory Is Not Dead](https://blog.hackspree.com/#factory-is-not-dead) — Bemer's 1968 "software factory," the factory framing the harness catalog inherits.
