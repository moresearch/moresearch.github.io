---
title: "The OpenHands Software Agent SDK: Event-Sourced Foundations for Production Agents"
date: 2026-08-26
slug: openhands-software-agent-sdk
summary: "A close reading of the MLSys 2026 paper (arXiv:2511.03690) that redesigned OpenHands into a four-package SDK: optional sandboxing, immutable config with a single event-sourced source of truth, the Action-Execution-Observation tool contract, security in the loop, and a local-to-remote workspace factory — validated by 61% fewer production failures and SOTA on 3 of 5 benchmarks."
tags: [openhands, agent-sdk, event-sourcing, harness-engineering, agent-architecture, mlsys, mcp, sandboxing, security-analyzer, multi-llm-routing, software-agents, production-agents]
---

[The OpenHands Software Agent SDK: A Composable and Extensible Foundation for Production Agents](https://arxiv.org/abs/2511.03690) (Xingyao Wang, Simon Rosenberg, Juan Michelini, Calvin Smith, Hoang Tran, Engel Nyst, Rohit Malhotra, Xuhui Zhou, Valerie Chen, Robert Brennan, Graham Neubig; accepted at MLSys 2026) is a rare thing in the agent literature: an architecture paper that ships. It documents the complete redesign of the most popular open-source software-engineering agent — OpenHands, 64k+ GitHub stars in 18 months — into a modular SDK, and it defends the redesign with production telemetry rather than vibes. Repos: [software-agent-sdk](https://github.com/OpenHands/software-agent-sdk), [benchmarks](https://github.com/OpenHands/benchmarks).

This is also the substrate this blog wrote about yesterday: the [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models) paper built its harness-optimizer search space on the SDK's API surface — system prompts, tools, hooks, context management, sub-agents. Reading the SDK paper explains *why* that search space worked: it is the parts catalog of a production agent harness, deliberately factored so components can be swapped, serialized, and recomposed.

The paper's own framing is a design problem: *"no consensus exists on how to architect these foundations."* OpenAI's SDK, Claude's Agent SDK, Google ADK, and LangGraph all differ wildly in isolation models, state management, and deployment assumptions. The paper's contribution is a well-validated reference design. The SDK works in six lines:

```python
from openhands.sdk import LLM, Conversation
from openhands.tools.preset.default import get_default_agent
llm = LLM(model="openhands/claude-sonnet-4-5-20250929", api_key="...")
agent = get_default_agent(llm=llm)
conversation = Conversation(agent=agent, workspace="/path/to/project")
conversation.send_message("Write 3 facts about this project into FACTS.txt.")
conversation.run()
```

## The redesign is the story: four principles from V0's debt

The SDK is not a greenfield design; it is the answer to architectural debt accumulated in the original monolithic OpenHands (V0). Four pain points, four principles:

| V0 pain | V1 design principle |
|---|---|
| **Universal sandboxing** — every tool call ran in a Docker container; two processes (agent + sandbox) with divergent states; local workflows needed duplicated MCP/tool implementations; one tenant's heavy actions could crash shared containers | **Optional isolation** — agent and tools run in one process by default (aligning with MCP's local-first assumptions); sandboxing is opt-in for production |
| **Mutable config sprawl** — 140+ fields, 15 classes, 2.8K lines of configuration across parallel hierarchies (CLI, web UI, GitHub App, SaaS), each with its own precedence rules; identical parameters could diverge | **Stateless by default, one source of truth** — every component is an immutable, validated Pydantic model; the only mutable entity is `ConversationState`, which records all state in an append-only event log |
| **Monorepo coupling** — agent core, evaluation suite, frontend, CLI, and GitHub integrations in one repository; benchmark dependencies and version conflicts leaked into production | **Strict separation of concerns** — the agent core is a shared library; applications (CLI, GUI, GitHub App) consume it via APIs |
| **Monolith logic** — new behaviors required editing core logic or branching per entry point | **Two-layer composability** — four independent deployment packages (sdk, tools, workspace, server) plus a typed component model for safe extension |

Each principle is a decision with a rejected alternative, and the paper is honest about the tradeoffs: event sourcing beat a database-backed model because a DB couples the SDK to a storage backend and breaks offline replay; optional isolation beat both mandatory containerization (V0's fragility) and fully-local-only execution (production safety needs).

## The spine: event-sourced state as single source of truth

The architectural centerpiece is the event model. All interactions are immutable events appended to a log, organized in a two-tier hierarchy:

- **LLM-convertible events** — `MessageEvent`, `ActionEvent` (tool calls), `SystemPromptEvent`, `CondensationSummaryEvent`, and `ObservationBaseEvent` subclasses (tool results, user rejections, agent errors) — these are what the model sees.
- **Internal events** — `ConversationStateUpdateEvent`, `CondensationRequest`, `Condensation`, `PauseEvent` — bookkeeping that never reaches the LLM.

Components (Agent, Tool, LLM) are immutable and serializable; *all* mutable variables live in `ConversationState`, which combines metadata fields with the append-only `EventLog` behind a FIFO lock. Persistence is dual-path: metadata serializes to a single `state.json`, events write as individual JSON files — so resuming a conversation means loading `state.json` and replaying the event directory, with automatic detection of incomplete conversations.

The design's payoff is measurable and the paper measures it: replaying 39,870 events from 433 real SWE-Bench conversations, per-event persistence is a 0.20 ms median (0.40 ms per action+observation cycle), full state replay is 4.1 ms, and crash recovery completes in 7.4 ms (32.1 ms even for the longest observed conversation at 358 events). Against LLM round-trips of 1–30 s, event sourcing is free. This is the same family of design as [DeepSeek Harness's enforced append-only session log](https://blog.hackspree.com/#deepseek-harness) and the [spatiotemporal-composability](https://blog.hackspree.com/#spatiotemporal-composability) calculus — the durable-facts-on-a-log spine is becoming the consensus architecture for agent harnesses, and this paper is the largest-scale validation of it yet.

## Stateless components, typed contracts

The immutable-component principle shapes everything else.

**The LLM abstraction** is deliberately vendor-agnostic. Through LiteLLM it reaches 100+ providers over both the Chat Completions and the newer OpenAI Responses APIs; it natively consumes extended-thinking fields (`ThinkingBlock` for Anthropic, `ReasoningItemModel` for OpenAI); and it supports models *without* function calling at all via `NonNativeToolCallingMixin`, which converts tool schemas into text prompts and parses tool calls with regex extraction — dramatically expanding the usable model pool. On top sits `RouterLLM`, which routes different requests to different models while keeping a unified `LLM` interface:

```python
class MultimodalRouter(RouterLLM):
    def select_llm(self, messages: list[Message]) -> str:
        has_images = any(m.contains_image for m in messages)
        return "primary" if has_images else "secondary"

# route text to cheaper model, images to multimodal model
router = MultimodalRouter(llms_for_routing={
    "primary": LLM(model="claude-sonnet-4-5"),
    "secondary": LLM(model="devstral-small")})
agent = Agent(llm=router, tools=tools)
```

This is the harness as a cost-engineering surface — [every token has a price tag](https://blog.hackspree.com/#every-token-has-a-price-tag), and the SDK lets you express the routing policy in a method.

**The tool system** follows an Action–Execution–Observation contract: the LLM proposes JSON tool calls, which are validated into typed `Action` schemas before execution; `ToolExecutor` runs the logic; `Observation` structures the result for the LLM. MCP tools are first-class citizens — their JSON schemas are auto-translated into `Action` models, so external MCP servers behave identically to native tools (validated, type-safe, serialized). The tool *registry* decouples specifications from implementations: Python executors aren't serializable, so tools travel across process or network boundaries as lightweight JSON specs, reconstructed lazily at runtime. That single mechanism enables distributed agent architectures — and it is exactly what made the [harness optimizer](https://blog.hackspree.com/#better-harnesses-smaller-models) search space editable: harnesses are just compositions of serializable specs.

## Security inside the loop

The SDK treats safety as part of the control loop rather than an afterthought, through two abstractions: `SecurityAnalyzer`, which rates each tool call low/medium/high/unknown risk, and `ConfirmationPolicy`, which decides whether approval is required. The built-in pair — `LLMSecurityAnalyzer` plus `ConfirmRisky` — blocks actions above a configurable threshold (default: high), pausing the agent in a `WAITING FOR CONFIRMATION` state until a human acts. Because risk assessment is separated from enforcement, policies can be adapted mid-session (e.g., relaxing restrictions for read-only `grep`), and custom analyzers slot in without touching tool executors. `SecretRegistry` complements it with late-bound secrets: tools fetch credentials only at execution time, and any secret that appears in output is masked as `<secret-hidden>`, with support for live rotation mid-conversation. These are precisely the features that are *absent* from the feature-comparison table's OpenAI/Claude/Google/LangChain columns — the security analyzer, confirmation policies, secrets with auto-masking, and agent stuck detection are OpenHands-unique.

## Local-first, deploy-anywhere

The paper's most immediately useful trick is the `Conversation` factory. Instantiate with a path or `LocalWorkspace` and you get a `LocalConversation` — the full agent loop (LLM calls, tool invocation, event callbacks, state updates) running in-process, perfect for notebooks and debugging, with pause/resume. Swap in a `RemoteWorkspace` and the *same code* transparently constructs a `RemoteConversation` that serializes the agent config and delegates execution to an agent server over HTTP/WebSocket. Going from prototype to production is a two-line diff:

```python
+from openhands.workspace import DockerWorkspace
+with DockerWorkspace(...) as workspace:
+    conversation = Conversation(agent=agent, workspace=workspace)
+    conversation.send_message("Create hello.py")
+    conversation.run()
```

The server itself is a first-class artifact: REST endpoints for conversation control, WebSocket for event streaming, session-based authentication, and official Docker images bundling the full stack (API server, VSCode Web, VNC desktop, Chromium). Each agent instance runs in an independent container with dedicated filesystem and resource limits — the containerized model that makes SaaS-style multi-tenancy tractable, with per-conversation event logs providing tenant isolation by construction.

## The numbers that validate the architecture

**Production reliability.** A 15-day parallel rollout ran V0 and V1 against live users. System-attributable errors per 1k conversations fell from 78.0 to 30.0 — a 61% reduction. The entire V0 infrastructure-error class (69.8/1k) vanished: it was dominated by inter-pod HTTP failures between the conversation manager and execution runtime (`HTTPStatusError` 401 at 43.0/1k, `AgentRuntimeNotReadyError` at 18.8/1k, connection timeouts at 3.1/1k), all eliminated by V1's co-located execution model. The residual V1 SDK errors (29.7/1k) were dominated by one condensation bug exposed during the extended-thinking rollout — caught in production, fixed, and instructive: even a good architecture ships bugs, but co-location removed an entire failure class structurally rather than by patching symptoms.

**Capability preservation.** On SWE-Bench Verified with matched models, Claude Sonnet 4 scores an identical 68.0% on V0 and V1 — the redesign neither helped nor hurt. With Sonnet 4.5, V1 gains +8.2 points (64.6 → 72.8), attributed to extended-thinking support that the event-sourced architecture integrated naturally but would have been "significant engineering effort" to retrofit into V0.

**Benchmark SOTA.** Across 14 models (7 closed, 7 open) and five task categories — issue resolution (SWE-Bench Verified), greenfield development (Commit0), frontend development (SWE-Bench Multimodal), testing (SWT-Bench), and information gathering (GAIA) — the SDK hits state-of-the-art on **3 of 5**: Commit0 at 56.2% (published SOTA: 12.5%), SWE-Bench Multimodal at 44.1% (no published SOTA yet), and GAIA at 80.0% (SOTA: 74.6%). It remains within 2.6 points (SWE-Bench Verified, 76.6 vs 79.2) and 5.2 points (SWT-Bench, 78.8 vs 84.0) of SOTA on the other two — using a *single model per evaluation*, while several SOTA systems use multi-model orchestration. The uniform harness also surfaces model specialization that a model-agnostic substrate makes visible: Claude models dominate issue resolution and testing, while GPT-5.4 leads long-horizon greenfield development.

**Three-tier QA.** The engineering discipline extends to testing: programmatic tests with mocked LLM calls run on every commit (seconds, no API cost); LLM-based integration and example tests run daily and on-demand for PRs ($0.5–$3, under 5 minutes, real models); benchmark evaluations run on-demand at $100–$1000 for comprehensive capability checks. The integration-test framework — subclass, implement `setup()`, `tools()`, `verify_result()` — is itself a template for [harness QA](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems).

## What this means

Read as a reference architecture, the SDK codifies a specific set of bets, and the blog's prior coverage suggests they are the right ones:

**The harness is the product.** The minimal example is six lines because everything that matters — state, tools, security, context, deployment — lives in the foundation, not in the agent loop. This is the same conclusion as [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models): the agent's value is the software around the model, and now that software is a maintained, versioned, open-source artifact instead of bespoke glue.

**Event-sourced state is the consensus spine.** DeepSeek Harness's append-only session log, Cordis's composability calculus, and now OpenHands V1's `ConversationState` — three independent designs converging on immutable-event logs with derived state. The 0.2 ms persistence latency and 7.4 ms recovery numbers make the objection ("event sourcing is too slow") an empirical dead end.

**Local-first beats sandbox-first.** The mandatory-sandbox assumption was the source of most of V0's production failures. Opt-in isolation — one process by default, containers when you need them — aligns with MCP's local-first reality and removes an entire infrastructure failure class. The lesson generalizes beyond OpenHands: [terminal agents](https://blog.hackspree.com/#terminal-agents-survey) live in the local substrate, and harnesses that treat local execution as the default rather than the exception are structurally simpler.

The honest limits are stated: the SDK currently targets single-agent conversations (multi-agent coordination "requires further design"); LLM-based security analysis remains subject to adversarial prompts and inconsistent classification; and a comprehensive multi-tenant security audit is future work. The security architecture, in particular, is defense-in-depth with a probabilistic core — the `LLMSecurityAnalyzer` is itself a model, so the guardrail inherits the model's failure modes, which is exactly the kind of leak the [leaky-abstractions law](https://blog.hackspree.com/#hacker-laws-ase-leaky-abstractions) predicts.

For anyone building software-engineering agents, this paper is the current best answer to "what should the foundation look like": immutable validated components, one event-sourced source of truth, a typed tool contract with MCP as a first-class citizen, security inside the loop, and a workspace abstraction that makes the notebook-to-production journey a two-line diff. The architecture has now been validated twice over — by 61% fewer production failures on one axis, and by being the substrate the harness-optimization literature [builds on](https://blog.hackspree.com/#better-harnesses-smaller-models) on the other.
