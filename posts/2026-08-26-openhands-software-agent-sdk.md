---
title: "The OpenHands Software Agent SDK: Event-Sourced Foundations for Production Agents"
date: 2026-08-26
slug: openhands-software-agent-sdk
summary: "Six insights from the MLSys 2026 paper (arXiv:2511.03690): architectural debt drove the redesign; event-sourced state is the consensus spine; local-first beats sandbox-first; serializability is the composability engine; security belongs inside the loop; and the harness is the product."
tags: [openhands, agent-sdk, event-sourcing, harness-engineering, agent-architecture, mlsys, mcp, sandboxing]
---

[The OpenHands Software Agent SDK](https://arxiv.org/abs/2511.03690) (Wang et al., MLSys 2026) is a rare thing in the agent literature: an architecture paper that ships. It documents the redesign of the most popular open-source software-engineering agent — OpenHands, 64k+ GitHub stars in 18 months — into a modular SDK, and defends it with production telemetry: **61% fewer system-attributable failures** in a 15-day live comparison, and state-of-the-art results on 3 of 5 benchmarks. Repos: [software-agent-sdk](https://github.com/OpenHands/software-agent-sdk), [benchmarks](https://github.com/OpenHands/benchmarks).

It is also the substrate this blog wrote about yesterday: the [harness optimizer](https://blog.hackspree.com/#better-harnesses-smaller-models) in "Better Harnesses, Smaller Models" searches this SDK's design space — system prompts, tools, hooks, context management, sub-agents. The SDK paper explains why that space was searchable at all. A complete agent runs in six lines:

```python
from openhands.sdk import LLM, Conversation
from openhands.tools.preset.default import get_default_agent
llm = LLM(model="openhands/claude-sonnet-4-5-20250929", api_key="...")
agent = get_default_agent(llm=llm)
conversation = Conversation(agent=agent, workspace="/path/to/project")
conversation.send_message("Write 3 facts about this project into FACTS.txt.")
conversation.run()
```

## Insight 1 — The redesign is the story: four principles from V0's debt

The SDK is not a greenfield design; it is the answer to four failures of the original monolithic OpenHands (V0):

| V0 pain | V1 principle |
|---|---|
| Universal sandboxing: two processes with divergent state; local workflows needed duplicated tool/MCP code | Optional isolation: one process by default, containers opt-in for production |
| Config sprawl: 140+ fields, 15 classes, 2.8K lines across parallel hierarchies; identical parameters diverging | Stateless by default: immutable validated components; one source of truth for state |
| Monorepo coupling: agent core, eval suite, apps, and benchmarks in one repo; version conflicts leaking into production | Strict separation: the core is a shared library consumed via APIs |
| Monolith logic: new behaviors required editing core or branching per entry point | Two-layer composability: four deployable packages plus a typed component model |

Each choice has a rejected alternative, stated honestly: event sourcing beat a database-backed model because a DB couples the SDK to a storage backend and breaks offline replay; optional isolation beat both mandatory containerization (V0's fragility) and fully-local-only execution (production needs safety).

## Insight 2 — Event-sourced state is the consensus spine

The centerpiece is the state model. Every component (Agent, Tool, LLM) is an immutable, validated, serializable Pydantic model; *all* mutable variables live in `ConversationState`, which records interactions in an append-only `EventLog`. Events are two-tiered: LLM-convertible ones (messages, tool calls, observations, system prompts) are what the model sees; internal ones (`CondensationRequest`, `PauseEvent`, state updates) are pure bookkeeping. Persistence is dual-path — metadata in `state.json`, events as individual JSON files — so resuming a conversation means loading the file and replaying the log, with automatic detection of incomplete conversations.

The overhead objection is empirically dead: replaying 39,870 events from 433 real SWE-Bench conversations, persistence costs 0.20 ms median per event, full state replay 4.1 ms, crash recovery 7.4 ms (32.1 ms worst case). Against LLM round-trips of 1–30 s, event sourcing is free.

This is the third independent convergence this blog has covered on the same spine: [DeepSeek Harness's enforced append-only session log](https://blog.hackspree.com/#deepseek-harness), the [spatiotemporal-composability](https://blog.hackspree.com/#spatiotemporal-composability) calculus, and now OpenHands V1. Durable facts on a log, state as a pure derivation — that is the consensus architecture of the agent harness.

## Insight 3 — Local-first beats sandbox-first

V0's mandatory sandboxing wasn't just awkward; in production it was the dominant source of failure. The conversation manager and execution runtime talked over inter-pod HTTP, and the 15-day V0/V1 rollout shows what that bought: `HTTPStatusError` 401s at 43.0 per 1k conversations, runtime-not-ready races at 18.8 per 1k, connection timeouts at 3.1 per 1k. V1 runs the agent and tools co-located by default, which **eliminated the entire infrastructure-error class (69.8 → 0.0 per 1k) and cut system-attributable failures 61% (78.0 → 30.0 per 1k)**. The architecture removed a failure class structurally instead of patching its symptoms.

The same principle drives the deployment story. `Conversation` is a factory: give it a path or `LocalWorkspace` and you get an in-process loop with pause/resume (perfect for notebooks and debugging); give it a `RemoteWorkspace` and the *same code* serializes the agent config and delegates execution to an agent server over REST/WebSocket. Prototype to production is a two-line diff:

```python
+from openhands.workspace import DockerWorkspace
+with DockerWorkspace(...) as workspace:
+    conversation = Conversation(agent=agent, workspace=workspace)
+    conversation.send_message("Create hello.py")
+    conversation.run()
```

## Insight 4 — Serializability is the composability engine

The tool system runs on an Action–Execution–Observation contract: LLM tool calls validate into typed `Action` schemas before execution; `ToolExecutor` runs them; `Observation` structures results for the model. MCP tools are first-class — their schemas translate automatically into the same contract, so external servers behave identically to native tools.

The key move is the **tool registry**: Python executors aren't serializable, so tools travel across process and network boundaries as lightweight JSON specs, reconstructed lazily at runtime. That single mechanism is what makes distributed architectures and editable harnesses possible — specs can be composed, swapped, and searched. It is exactly why the [harness optimizer](https://blog.hackspree.com/#better-harnesses-smaller-models) could treat harnesses as a searchable space: a harness is just a composition of serializable specs.

## Insight 5 — Security belongs inside the loop

The SDK treats safety as a control-loop component, not an afterthought, with two abstractions: `SecurityAnalyzer`, which rates each tool call low/medium/high/unknown risk, and `ConfirmationPolicy`, which decides when approval is required — pausing the agent in a `WAITING FOR CONFIRMATION` state until a human responds. Because assessment is separated from enforcement, trust can adapt mid-session (relax restrictions for read-only `grep`) and custom analyzers slot in without touching tool executors. `SecretRegistry` completes the picture: late-bound secrets fetched only at execution time, masked as `<secret-hidden>` in any output, rotatable mid-conversation.

These are precisely the features absent from the paper's comparison of the OpenAI, Claude, Google, and LangChain SDKs: the security analyzer, confirmation policies, secret auto-masking, and agent stuck detection are OpenHands-unique.

## Insight 6 — The harness is the product (and the numbers back it)

Everything else follows the same logic. The LLM layer reaches 100+ providers, consumes native extended-thinking fields, and supports non-function-calling models via a text-prompt fallback — widening the usable model pool; `RouterLLM` lets routing policy (images to a multimodal model, text to a cheap one) live in a method, making the harness a [cost-engineering surface](https://blog.hackspree.com/#every-token-has-a-price-tag). The `Condenser` halves API cost with no degradation by summarizing overflow.

The redesign's capability claims hold up: matched models score an identical 68.0% on SWE-Bench Verified across V0 and V1 (parity — the redesign didn't hurt), while Sonnet 4.5 gains +8.2 points on V1 thanks to extended-thinking support the event-sourced architecture absorbed naturally. Across 14 models and five task categories, the SDK hits **SOTA on 3 of 5** — Commit0 at 56.2% vs. published SOTA 12.5%, GAIA at 80.0% vs. 74.6%, SWE-Bench Multimodal at 44.1% — using a single model per evaluation where several SOTA systems use multi-model orchestration. A three-tier QA pipeline (mocked-LLM tests on every commit, $0.5–3 LLM tests daily, $100–1000 benchmark runs on demand) keeps the whole thing honest.

The stated limits are worth repeating: single-agent focus (multi-agent coordination is future work); and a security core that is probabilistic — `LLMSecurityAnalyzer` is itself a model, so the guardrail inherits the model's failure modes, exactly the [leaky abstraction](https://blog.hackspree.com/#hacker-laws-ase-leaky-abstractions) this blog keeps returning to.

## What this means

Read as a reference architecture, the SDK codifies three bets that now look convergent across the field: **event-sourced state is the spine**, **local-first execution is the default**, and **the harness is the product** — everything that matters (state, tools, security, context, deployment) lives in the foundation, which is why the minimal agent is six lines and why the [harness-optimization literature](https://blog.hackspree.com/#better-harnesses-smaller-models) builds directly on it. The architecture has been validated twice over: by 61% fewer production failures on one axis, and by being the substrate the next wave of agent research searches on the other.
