---
title: "Emerging Harness Patterns and Anti-Patterns for Agentic AI Systems"
date: 2026-08-23
slug: harness-patterns-for-agentic-ai-systems
summary: "A problem-first pattern language for agentic AI systems: fifteen problems, each solved by one or more harness patterns, failed by one or more anti-patterns, and advanced by frontier patterns — with references grouped per problem. The intro argues the central technical insight with evidence: the unit of design is the agentic AI system, not the agent — we harness the whole system, including the agent(s), because the agent is the one component you cannot design. Every pattern and anti-pattern is stated in a fixed, consistent form (Forces, Solution, Consequences, Tradeoffs, Evidence, Related / Smell, Anti-solution, Failure, Refactoring, Evidence, Related)."
tags: harness, harness-engineering, agent-harness, pattern-language, patterns, anti-patterns, agents, agentic-ai, agentic-systems, security, sandboxing, memory, orchestration, tool-binding, verification, token-economics, emerging-patterns
---

A harness is the environment an agentic AI system lives in: everything between the models and the world. It decides what the system can see, what it can do, and how the outcome is judged. This blog has been arguing for months that the harness — not the model — is where the leverage concentrates, and the numbers have arrived: the same model finished between 47% and 67% of real tasks across eight different harnesses while cost per finished task varied sevenfold, and the same weights score 3.4 points apart on the public Terminal-Bench board under two major harnesses. Nothing about the intelligence changed. The wrapper changed.

## The unit of design is the agentic AI system, not the agent

This post is deliberately titled for **agentic AI systems**, and the distinction is the whole argument. A *harness for an AI agent* is an instrument around one model — a wrapper for a single component. A *harness for an agentic AI system* is the entire runtime around one or many agents: the loop, the tools, the memory and state, the interfaces the agents read, the policies that constrain them, the verifiers that judge them, the orchestrators that compose them, the humans in the loop, and the evaluation and billing apparatus that closes the loop. **We harness the whole system, including the agent(s). We do not care about the agent as such** — not because agents do not matter, but because the agent is the one component we cannot design, and everything that can be designed is the system.

The definitional case is made by the people who sell the components. Google's Agents whitepaper defines the agent as "a program that extends beyond the standalone capabilities of a Generative AI model" — an application of model, tools, and an orchestration layer; the thing you ship is the application, and the model is a part of it. MemGPT's title makes the same claim in one phrase: *LLMs as operating systems*. An operating system is not a CPU; it is everything between the CPU and the user — process management, memory paging, scheduling, permissions, the file system, the shell. That is exactly what a harness is for an agentic AI system, and the analogy runs deeper than rhetoric: the OS's core discipline is that hardware is replaceable and the interface is stable. The agent is the hardware. The harness is the OS.

The empirical case is stronger. In every controlled comparison, the agent was fixed and the system changed, and the system decided the outcome:

- **Same model, eight harnesses:** the same model ran thirty real work tasks through eight different harnesses; success swung from 46.7% to 66.7%, and cost per finished task varied by a factor of seven ([DeepSeek Harness teardown](https://blog.hackspree.com/#deepseek-harness)).
- **Same weights, public board:** the same weights score 83.8% under one major harness and 80.4% under another — 3.4 points decided entirely by software with no weights in it.
- **Same model, different interface:** SWE-agent achieved 12.5% pass@1 on SWE-bench and 87.7% on HumanEvalFix with GPT-4 — "far exceeding the previous state-of-the-art achieved with non-interactive LMs" — by changing only the agent-computer interface ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)). AgentBench reached the same conclusion across eight environments: how the agent observes and acts determines more of the outcome than the model's raw capability.
- **No training, harness-level only:** terminal-bench-rl reports that harness-level improvements — native tool calling, multimodal support, execution optimization — lifted scores *without training*, and Meta-Harness reached 76.4% on Terminal-Bench 2.0 and "was itself discovered through automated harness evolution."
- **Same model, different harness, dramatic quality gap:** Anthropic's long-running harness turned a one-sentence prompt into a 16-feature application that a solo agent run could not approach — at 20x the cost, same model ([Harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps)).
- **Smaller model, better system:** a 350M-parameter specialist, distilled and harnessed, beats ChatGPT on tool calling by 51 points ([Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale)).

Why it is a mistake to focus on the agent is a catalog of failure modes, each documented in this blog's archive:

**It misattributes failures.** When an agentic system fails, the agent-focus blames the model; the system-focus asks which tool it selected, what context it was given, what interface it read, which verifier approved it. Fowler's team spent three days investigating a feature an agent inserted that nobody asked for — that is a systems failure (no verifier caught an unrequested change), not a model failure ([Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)). The self-review datum is the same error at review time: 79% of 25,264 agent-generated PRs were reviewed by the same developer who prompted the agent ([mob programming remastered](https://blog.hackspree.com/#mob-programming-reimagined)). And "[the model decided]" is not an audit answer — regulators and boards accept checkpointed state, provenance chains, and deterministic replay ([durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution)).

**It optimizes the least controllable part.** The agent is the only stochastic component — "given the same input, you get different outputs" — and the only reliable method for stochastic components is measurement, which is a system function ([Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents)). Every lever you actually pull is a system lever: the interface, the memory, the tools, the policies, the verifiers. "The transition from folk expertise to data-driven design is the transition from craftsmanship to engineering."

**It ignores where improvement lives.** "Agents don't learn. Every mistake an agent makes, it will make again unless the harness explicitly prevents it" ([Fowler](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)). The loop that improves — read lessons, do work, reflect, write lessons — is a system loop ([harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents)). The training environment is a system component: DeepSeek ships the RL composition as a product preset because "the harness produces the trajectories; the trajectories feed post-training" ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). The closed loop is a system: distillation's moat is "the loop has accumulated months of self-generated training data" ([distillation](https://blog.hackspree.com/#agents-are-distillation-at-scale)).

**It confuses a component with its properties.** Every governance-relevant property is a system property, not an agent property: identity lives in the keypair, auditability in the log, authority in the guards, cost in the throttle, safety in the sandbox, correctness in the verifier ([Buzz](https://blog.hackspree.com/#buzz-block-agents), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), [always-on agents](https://blog.hackspree.com/#always-on-agents)). The test for agent infrastructure is "can you answer 'who did what, why, and was it correct?' six months after the fact, after rotating credentials, switching models, and changing platforms?" — the agent cannot answer that question; only the system can. The systems view also explains the deepest boundary from the composability calculus: recovery and correctness are promised *as wide as the system reifies* — "composition makes change safe; verification makes change right," and both are properties of the system, not of the agent ([spatiotemporal composability](https://blog.hackspree.com/#spatiotemporal-composability)).

The pattern language that follows treats the whole agentic AI system as the subject. Every pattern below names a decision about the system — the loop, the memory, the interface, the policy, the verifier, the orchestration — never about the model. The model is a socket you plug into the system. The system is what you build.

## How to read this pattern language

The pattern language is organized **by problem, not by component**. Each unit of the language is one recurring engineering problem that every agentic AI system eventually faces; inside each problem unit live the *pattern* that solves it, the *anti-pattern* that fails it, and — where the frontier has moved — the *frontier pattern* that is evolving it. This is the structure of Christopher Alexander's *A Pattern Language* (1977): the pattern is "a set of problems and documented solutions," each one "our current best guess as to what arrangement... will work to solve the problem presented," cross-referenced into a network. The linear order matters: read top to bottom, and the problems build — first make the system safe (problems 1–4), then make it remember (5–7), then make it act (8–11), then make it scale (12–14), then make it trustworthy (15).

Each **pattern** is stated in exactly this form, in this order:

- **Forces** — the competing pressures that make the problem hard.
- **Solution** — what to build.
- **Consequences** — what the pattern buys.
- **Tradeoffs** — what it costs.
- **Evidence** — the primary sources and where it shows up.
- **Related** — the other entries it composes with or contradicts.

Each **anti-pattern** is stated in exactly this form, in this order:

- **Smell** — how to recognize it in a system.
- **Anti-solution** — what teams reach for.
- **Failure** — why it breaks, with evidence.
- **Refactoring** — the pattern in this language that fixes it.
- **Evidence** — the sources that document the failure.
- **Related** — the entries it leads to.

Each problem unit ends with **References for this problem** — the sources for that unit alone, so the language can be read as a sequence of self-contained decisions. The catalog as a whole:

| # | Problem | Patterns | Anti-patterns |
|---|---|---|---|
| 1 | Containing untrusted execution | P1 Ephemeral Sandbox Wrapper | A1 The Naked Prompt |
| 2 | Refusing a tool call before it happens | P2 Static Intercepting Gatekeeper | A3 Prompt-Driven Authorization |
| 3 | Stopping the loop for a human | P3 HITL Breakpoint | — |
| 4 | Bounding the loop | P4 Token & Time Budget Throttler | A2 The Infinite Execution Vortex |
| 5 | Keeping a long session in a lean window | P5 Rolling Window Compression · F3 Context Resets · F6 Context Engineering | A4 The Context Avalanche |
| 6 | Choosing what context to inject | P6 Semantic Memory Router | A6 The RAG Firehose |
| 7 | Making state survive and giving it homes | P7 State Snapshot & Rollback · P8 Tiered Hierarchical Memory | A5 Goldfish Amnesia |
| 8 | Typing tool output and keeping errors legible | P9 Schema Enforcement & Self-Correction | A7 The Silent Crash · A9 The Schema Free-for-All |
| 9 | Not blocking the loop on long tools | P10 Asynchronous Tool Worker Queue | — |
| 10 | Making the fast loop repeatable | P11 Mock Tool Virtualization | — |
| 11 | Discovering capabilities without bloating the prompt | P12 Dynamic Tool Discovery / Registry · F4 The Interop Layer | A8 The Bloated Utility Belt |
| 12 | Dividing a workflow across agents | P13 Orchestrator-Worker · F2 Sprint Contracts | A11 The God Agent |
| 13 | Coordinating through shared state without corruption | P14 Blackboard (Shared Workspace) | A12 State Race Conditions |
| 14 | Keeping linear flows linear | P15 Sequential Pipeline Routing | — |
| 15 | Producing a verdict no single model can fake | P16 Voting / Consensual Ensemble · F1 Generator–Evaluator Loop · F5 Live-Environment Evaluators | A10 The Committee Paradox |

# Part I — Fifteen problems, their patterns, and their anti-patterns

## Problem 1 — Containing untrusted execution

An agentic system executes code and commands that are untrusted by construction — written by a model, possibly steered by injected instructions. A mistake or an attack must die with the task that produced it. This problem is where the harness earns its name: it is the operational translation of defense in depth from [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents), fail-closed policy from the [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), and subtraction from [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface).

### Pattern 1 — Ephemeral Sandbox Wrapper

**Forces.** Isolation wants a real boundary; performance wants none. Ephemerality wants nothing to survive; long-running work wants persistence. Teardown wants completeness; cleanup wants to be free.

**Solution.** Spawn isolated, short-lived virtual environments (e.g., Docker, WASM) for safe, untrusted agent code execution. Create per task, mutate freely, destroy on completion.

**Consequences.** Blast radius is bounded by lifetime, not by trust. Training trajectories come out clean because the RL environment is an ephemeral shell. The wrapper is where the system reifies the outside world, which is what makes recovery promises possible.

**Tradeoffs.** Ephemerality costs startup latency and state loss — an agent that needs a persistent filesystem across steps fights the wrapper, which is why the pattern is usually paired with a persistent-but-isolated volume. Real isolation (hypervisor, OS-level) is heavy; WASM is light and fast but constrains what the agent can do. The wrapper is only as good as its teardown: if the environment survives, the "ephemeral" claim is marketing — teardown must be [derived from the load, not remembered](https://blog.hackspree.com/#spatiotemporal-composability).

**Evidence.** LangChain now lists sandboxing as a first-class integration component ([Sandbox integrations](https://python.langchain.com/docs/integrations/)); Replit's thirteen-layer stack is the production-grade instance ([Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents)); the DeepSeek harness ships a file-effects-only sandbox vocabulary with enforcement reported `full` or `partial` ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Composes with P2 (the gatekeeper runs outside the wrapper); the sandbox-stack pattern from the first edition ([Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents)) is this pattern's defense-in-depth cousin; it is the refactoring for A1.

### Anti-Pattern 1 — The Naked Prompt (Implicit Trust)

**Smell.** API keys in the prompt or environment; the model told to "be careful"; no sanitization, parsing, or proxy layer between the model and the credentials.

**Anti-solution.** Treat the model as an application boundary and the prompt as an access control list.

**Failure.** The model is an untrusted input processor; prompt injection converts instructions into actions; system prompts leak; credentials exfiltrate. OWASP's relevant entries are LLM01 Prompt Injection, LLM02 Sensitive Information Disclosure, and LLM07 System Prompt Leakage in the 2025 numbering (LLM06 Sensitive Information Disclosure in the 2023/24 numbering, where the classic citation lives) ([OWASP Top 10](https://genai.owasp.org/llm-top-10/)).

**Refactoring.** P1 (the sandbox wrapper), P2 (the gatekeeper), and a transparent secrets proxy so the agent never sees the credentials — the Replit pattern ([Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents)). The xz lesson applies: the tool that executes arbitrary code with your credentials is the highest-value target in your supply chain, and "an agent with a thousand dependencies is a thousand doors" ([Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface)).

**Evidence.** OWASP ([LLM Top 10](https://genai.owasp.org/llm-top-10/)); Simon Willison's running documentation of the threat model ([prompt injection series](https://simonwillison.net/series/prompt-injection/)).

**Related.** Leads to A3 (prompt-driven authorization is the same error one level down).

**References for this problem.** LangChain sandbox integrations ([docs](https://python.langchain.com/docs/integrations/)); Meta's Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); OWASP Top 10 for LLM Applications ([2025](https://genai.owasp.org/llm-top-10/)); Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); archive: [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents), [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), [spatiotemporal composability](https://blog.hackspree.com/#spatiotemporal-composability).

## Problem 2 — Refusing a tool call before it happens

The system's hands must be able to say no before anything reaches an external API — and the denial must be final: structural, not prose. This is the "who may act" problem: interception and authorization are the same seam.

### Pattern 2 — Static Intercepting Gatekeeper

**Forces.** Security wants denial to be final; usability wants appeals. Determinism wants a fixed rule; the threat surface wants adaptivity. Auditing wants every decision recorded; latency wants none.

**Solution.** Intercept model-generated tool calls against a strict blocklist before passing them to external APIs.

**Consequences.** A deterministic floor that is fast, auditable, and cannot be argued around. The gatekeeper is the system's version of [`pledge(2)`](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering): a restricted interface where the wrong thing is unexpressible.

**Tradeoffs.** The honest caveat is in the name: [Llama Guard](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/) is *not* static — it is itself an LLM, which means it can be fooled, bypassed, or prompted around, and it adds a model call to every tool invocation. A true static blocklist catches only what it enumerates. The pattern works when the blocklist is the floor and the model-based classifier is the next layer up — the [monotonic guard ordering](https://blog.hackspree.com/#deepseek-harness): deny by default, allow by exception, never let policy be argued around.

**Evidence.** Llama Guard, the LLM-based input-output safeguard (Meta AI, 2023), is the reference classifier ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); the DeepSeek tool pipeline encodes the ordering — pre-execute waterfalls, monotonic guards, fail-closed approval with `allowed-once` as the only granting outcome ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Composes with P1 (the wrapper contains what the gatekeeper misses) and P3 (the breakpoint is the gatekeeper with a human in it); it is the refactoring for A3.

### Anti-Pattern 3 — Prompt-Driven Authorization

**Smell.** "Do not delete user data" in the system prompt; permission checks that are sentences, not code.

**Anti-solution.** Policy as prose — the belief that the model will read and obey the instructions.

**Failure.** Instructions are data; prompt injection is the mechanism; "you can't solve AI security problems with more AI" ([Willison](https://simonwillison.net/series/prompt-injection/)). A system prompt is a document the model may be instructed to ignore.

**Refactoring.** Authorization must be [monotonic, structural, and fail-closed](https://blog.hackspree.com/#deepseek-harness): monotonic guards "deny or abstain and can never force-allow, so owner policy that must not be reordered cannot be argued around." Who can modify the agent's state is a property of the system, not of the model's reading comprehension ([always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Evidence.** Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); the DeepSeek tool pipeline's monotonic-guard doctrine ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Is the deeper form of A1; is fixed by P2 (the gatekeeper).

**References for this problem.** Meta's Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); OWASP Top 10 ([2025](https://genai.owasp.org/llm-top-10/)); Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (monotonic guards, approval seam), [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) (DSLs as the pledge(2) bridge), [always-on agents](https://blog.hackspree.com/#always-on-agents) (authority axis).

## Problem 3 — Stopping the loop for a human

Some mutations — financial transactions, deletions, releases — must not happen without human authority. The loop must stop, persist exactly where it paused, and resume only after a person decides. There is no named anti-pattern for this problem; its failure modes are the rubber-stamp (approval without attention) and the swallowed breakpoint, both covered in the tradeoffs.

### Pattern 3 — Human-in-the-Loop (HITL) Breakpoint

**Forces.** Autonomy wants the loop to run; safety wants it to stop. Latency wants no pauses; accountability wants a record of every pause. Approval wants to be meaningful; convenience wants it to be fast.

**Solution.** Freeze the harness execution loop to demand manual approval for high-risk mutations like financial transactions. Persist the exact state at the pause, then resume from that checkpoint after a human approves, edits, or rejects.

**Consequences.** Authority becomes a property of the system — a state-machine primitive, not a prompt: the loop does not merely ask permission, it persists where it paused. Every human decision is recorded in the state, which makes the breakpoint an audit seam ([always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Tradeoffs.** An agent that must stop for every risky action cannot run unattended — the pattern must be paired with an explicit automation path (permission presets, scoped allowlists) or the harness drowns in approvals and humans rubber-stamp everything, which is worse than no approval at all. A breakpoint that can be swallowed is a breakpoint that does not exist: the LangGraph guidance is explicit that interrupts must not be wrapped in try/except ([LangGraph interrupts](https://docs.langchain.com/oss/python/langgraph/interrupts)). And the breakpoint protects the mutation, not the model's intent behind it.

**Evidence.** LangGraph's `interrupt()` primitive makes HITL a first-class graph construct ([state management & breakpoints](https://docs.langchain.com/oss/python/langgraph/interrupts)); OpenWorker ships approval gates before every consequential action ([OpenWorker outcome layer](https://blog.hackspree.com/#openworker-outcome-layer)); the DeepSeek approval seam is fail-closed — a missing answerer resolves to `unavailable` ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Composes with P2 (the gatekeeper decides what is routine, the breakpoint what is consequential); the automation path it needs is P4's discipline; it is the governance seam the [always-on survey](https://blog.hackspree.com/#always-on-agents) names as authority and scope.

**References for this problem.** LangGraph interrupts ([docs](https://docs.langchain.com/oss/python/langgraph/interrupts)); archive: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer) (approval gates), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (approval seam, `allowed-once`), [always-on agents](https://blog.hackspree.com/#always-on-agents) (authority and scope axes).

## Problem 4 — Bounding the loop

Loops can burn unbounded tokens, wall-clock, and money retrying a broken step. The ceiling must be enforced by the system, never requested of the model — this is the enforcement half of the [Bill as Assertion](https://blog.hackspree.com/#every-token-has-a-price-tag).

### Pattern 4 — Token & Time Budget Throttler

**Forces.** Long-horizon work wants room; unbounded consumption wants none. Enforcement wants to be structural; convenience wants to be prompt-based. The bill wants a shape; the work wants a budget.

**Solution.** Monitor continuous tool loops to forcefully terminate agents exceeding maximum token costs or time boundaries. Enforce in the harness, never in the prompt.

**Consequences.** The vortex cannot happen: the bill is bounded by construction, and the system — not the model — owns the ceiling. "Make the cache hit a CI gate, make the bill a test."

**Tradeoffs.** A budget that is too tight kills legitimate long-horizon work; too loose is theater. The throttler interacts with [prefix caching](https://blog.hackspree.com/#deepseek-harness): a reset mid-session can invalidate the warm prefix and *increase* the bill it was meant to cap. And the throttler must decide what "exceeded" means — total spend, per-step spend, wall-clock, or loop iterations — because each catches a different failure mode. The honest limit is documented: DeepSeek's runaway-loop guard "only sends reminders and eventually goes quiet" — reminders are not enforcement ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Evidence.** AutoGPT shipped iteration and cost limits as first-class settings in 2023 precisely because its loop demonstrated the failure ([AutoGPT — maximum loop & cost control](https://docs.agpt.co/classic/configuration/)); OWASP names it a first-class risk — [LLM10:2025 Unbounded Consumption](https://genai.owasp.org/llm-top-10/); the token economics post computed the shape — prices down 98%, consumption up ~150x, bills tripled, and "the caps were about the shape of the bill" ([Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag)).

**Related.** Is the refactoring for A2; composes with P3 (the breakpoint stops the loop, the throttler ends it); connects to the bill-as-assertion pattern from the first edition ([Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag)).

### Anti-Pattern 2 — The Infinite Execution Vortex (Uncapped Loops)

**Smell.** No ceiling on iterations, tokens, time, or money; the same failing step retried; the loop "working on it."

**Anti-solution.** "It'll converge" — trust the model to stop.

**Failure.** Unbounded token drain and wall-clock burn; an economic failure before a technical one — the [tragedy of the commons](https://blog.hackspree.com/#every-token-has-a-price-tag) staged inside one run.

**Refactoring.** P4 (the token & time budget throttler), enforced in code — OWASP [LLM10:2025 Unbounded Consumption](https://genai.owasp.org/llm-top-10/), loop-iteration limits, and the honest caveat that reminders are not enforcement ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Evidence.** AutoGPT's open issue tracker is a museum of the failure ([issue tracker](https://github.com/Significant-Gravitas/AutoGPT/issues)); the token economics post computed the shape of the bill ([Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag)).

**Related.** Is the absence of P4; composes with A1 (naked prompts make vortices more dangerous); is the single-loop form of A10 (the committee paradox).

**References for this problem.** AutoGPT configuration ([docs](https://docs.agpt.co/classic/configuration/)) and issue tracker ([issues](https://github.com/Significant-Gravitas/AutoGPT/issues)); OWASP Top 10 — LLM10 Unbounded Consumption ([2025](https://genai.owasp.org/llm-top-10/)); archive: [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (runaway-loop guard, prefix caching).

## Problem 5 — Keeping a long session in a lean window

Context is finite and conversation is not; raw dumps degrade reasoning. The system must keep the session inside the window — and the frontier has split the answer into three: compression (P5), resets (F3), and the umbrella discipline of context engineering (F6).

### Pattern 5 — Rolling Window Compression

**Forces.** Fidelity wants the full history; the window wants a summary. Continuity wants in-place compaction; clarity wants a clean slate. The cache wants a stable prefix; the summarizer wants to rewrite it.

**Solution.** Automatically summarize older conversation histories in background threads to keep the active context window lean.

**Consequences.** Long sessions become possible at bounded cost. The pattern is a paging discipline: the context window is RAM, the log is disk, the summary is the page table. It is the practical implementation of MemGPT's virtual context management ([MemGPT](https://arxiv.org/abs/2310.08560)).

**Tradeoffs.** Summaries lose detail, and the loss is permanent unless the source log is preserved — which is why compression and the [append-only log](https://blog.hackspree.com/#deepseek-harness) must be separate layers. Compaction preserves continuity but not a clean slate: Anthropic's 2026 harness work found models exhibit "context anxiety" — "beginning to wrap up work prematurely as they approach what they believe is their context limit" — and compaction alone did not fix it ([harness design for long-running application development](https://www.anthropic.com/engineering/harness-design-long-running-apps)). And every compaction is a billing event: if the summarization call is not a genuine prefix-extension of the warm request, it pays full price for the whole replayed history — the exact bug DeepSeek's compaction fix corrected ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Evidence.** MemGPT formalized virtual context management with OS-style memory tiers ([paper](https://arxiv.org/abs/2310.08560)); DeepSeek's compaction engine produces the Claude-shaped eight-section checkpoint as a trailing prefix-extension message ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)); Anthropic documents the context-anxiety failure mode and the reset alternative ([frontier](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Related.** Is the refactoring for A4; conflicts with and composes with the append-only log (first edition); pairs with F3 (context resets) as the two answers to the filling window.

### Anti-Pattern 4 — The Context Avalanche (Memory Dumping)

**Smell.** Raw logs, every tool output, and full transcripts in the history; context perpetually near capacity; performance degrading as the session grows.

**Anti-solution.** "Long context solves it" — dump everything and let the model sort it out.

**Failure.** Model performance degrades significantly when relevant information sits in the middle of a long input context; performance is highest at the beginning and end ([Lost in the Middle](https://arxiv.org/abs/2307.03172)). A filled context is not a well-informed agent — it is a degraded one that reads the middle worst exactly when the middle holds the answer.

**Refactoring.** P5 (rolling window compression) over a [derived view](https://blog.hackspree.com/#deepseek-harness): 44 event types in the DeepSeek log, exactly three visible to the model — "the system never stores the conversation it sends; it recomputes the model's context from the log."

**Evidence.** Liu et al., *Lost in the Middle* (Stanford, 2023) ([paper](https://arxiv.org/abs/2307.03172)); the DeepSeek logged-surface invariant ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Is the absence of P5 and P8; is the substrate of A11 (the god agent) and the retrieval form of A6.

### Frontier 3 — Context Resets with Handoff Artifacts

**Problem.** Models lose coherence as the context window fills — "context anxiety": "beginning to wrap up work prematurely as they approach what they believe is their context limit." Compaction preserves continuity but not a clean slate, and for Sonnet 4.5 compaction alone was insufficient.

**Forces.** Continuity wants compaction; clarity wants a reset. State wants to survive; the window wants to be emptied. Cost wants fewer calls; quality wants the clean slate.

**Solution.** Distinguish *compaction* (summarize in place, same agent continues) from *context resets* (clear the context entirely, hand state to a fresh agent through a structured artifact). Use the reset when the model exhibits context anxiety.

**Consequences.** "A reset provides a clean slate, at the cost of the handoff artifact having enough state for the next agent to pick up the work cleanly." This is the [append-only log](https://blog.hackspree.com/#deepseek-harness) argument at the session boundary: the log never rewrites, and the handoff artifact is a new prefix, not an edit.

**Tradeoffs.** Resets add "orchestration complexity, token overhead, and latency to each harness run." And the frontier twist: Opus 4.5 "largely removed that behavior on its own," so the harness could drop resets entirely — the pattern is a function of the model, not just the system.

**Evidence.** Anthropic's harness work ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Related.** Is the companion of P5 (compression); composes with P13 (the reset is the delegation handoff); its umbrella is F6.

### Frontier 6 — Context Engineering

**Problem.** The system's behavior is decided by the configuration of tokens at inference time — the entire context state, not the prompt. Prompt engineering optimizes the wrong unit.

**Forces.** Prompting wants the right words; context wants the right state. The window wants curation; the loop wants accumulation.

**Solution.** Treat context engineering as the discipline of answering "what configuration of context is most likely to generate our model's desired behavior?" — curating the entire context state (system instructions, tools, MCP servers, external data, message history) as it evolves over turns ([Anthropic — Effective context engineering for AI agents](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)).

**Consequences.** The unit of design becomes the system's context state — which is the systems argument of this post in one sentence. "In contrast to the discrete task of writing a prompt, context engineering is iterative" and spans the whole runtime: this is P5, P6, and F3 unified under one discipline.

**Tradeoffs.** The context is a moving target that must be "cyclically refined" — and every refinement is a potential cache-prefix violation and a potential governance gap ([token economics](https://blog.hackspree.com/#every-token-has-a-price-tag), [always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Evidence.** Anthropic, *Effective context engineering for AI agents* (September 2025) ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)).

**Related.** Is the umbrella over P5, P6, F3; is the memory family's frontier expression of "the system, not the agent."

**References for this problem.** MemGPT ([arXiv:2310.08560](https://arxiv.org/abs/2310.08560)); Liu et al., Lost in the Middle ([arXiv:2307.03172](https://arxiv.org/abs/2307.03172)); Anthropic, Effective context engineering for AI agents ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (compaction, logged-surface invariant), [always-on agents](https://blog.hackspree.com/#always-on-agents).

## Problem 6 — Choosing what context to inject

Not every retrieved fact belongs in every prompt. Retrieval without judgment drowns the instruction — and every injected chunk is untrusted input. The system must decide, and the decision cannot be the model's.

### Pattern 6 — Semantic Memory Router

**Forces.** Grounding wants relevant facts; focus wants them selected. Freshness wants live retrieval; the cache wants a stable prefix. Utility wants injected context; security wants retrieved content treated as untrusted input.

**Solution.** Intercept ongoing tasks, query vector stores, and inject context fragments just-in-time into the agent's prompt.

**Consequences.** Answers get grounded; the prompt stays lean; attention is curated by the system rather than dumped by default. Injected context is the highest-leverage observation there is — "every line your CLI emits is an observation your agent reasons over" ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)).

**Tradeoffs.** The router is a point of failure and a point of attack: retrieved content is untrusted input, and prompt injection through a vector store is the vector of choice ([LLM08:2025 Vector and Embedding Weaknesses](https://genai.owasp.org/llm-top-10/)). Just-in-time injection competes with [prefix-cache discipline](https://blog.hackspree.com/#deepseek-harness) — context that changes every turn destroys the cache bookmark unless the injected region sits outside the prefix. And retrieval quality is decided by chunking, metadata filtering, and reranking, not by top-K volume — which is the difference between this pattern and its anti-pattern A6.

**Evidence.** Pinecone's RAG guides are the reference architecture ([Retrieval-Augmented Generation](https://www.pinecone.io/learn/retrieval-augmented-generation/)); the always-on survey's provenance and authority axes define what the router must track about the state it injects ([always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Related.** Is the refactoring for A6; composes with P8 (the tiers are the router's stores); the governance it inherits is the [always-on](https://blog.hackspree.com/#always-on-agents) authority question.

### Anti-Pattern 6 — The RAG Firehose

**Smell.** Top-K chunks injected by raw keyword match; the user instruction buried; retrieval results dominating the prompt.

**Anti-solution.** "More chunks equals better grounding."

**Failure.** Retrieval without judgment drowns the instruction — and the instruction sits in the middle, which is exactly where [Lost in the Middle](https://arxiv.org/abs/2307.03172) predicts degradation. Every injected chunk is untrusted input, making the firehose an injection vector ([LLM08:2025 Vector and Embedding Weaknesses](https://genai.owasp.org/llm-top-10/)).

**Refactoring.** P6 (the semantic memory router) with chunking, metadata filtering, and reranking deciding what is injected ([Pinecone advanced RAG](https://www.pinecone.io/learn/advanced-rag/)).

**Evidence.** Pinecone's advanced RAG guide is the counter-documentation ([chunking & metadata filtering](https://www.pinecone.io/learn/advanced-rag/)); OWASP's vector-and-embedding entry documents the attack ([LLM Top 10](https://genai.owasp.org/llm-top-10/)).

**Related.** Is the absence of P6; composes with A4 (the firehose is the avalanche's retrieval form).

**References for this problem.** Pinecone RAG guide ([learn](https://www.pinecone.io/learn/retrieval-augmented-generation/)) and advanced RAG ([learn](https://www.pinecone.io/learn/advanced-rag/)); OWASP Top 10 — LLM08 Vector and Embedding Weaknesses ([2025](https://genai.owasp.org/llm-top-10/)); Liu et al., Lost in the Middle ([arXiv:2307.03172](https://arxiv.org/abs/2307.03172)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design), [always-on agents](https://blog.hackspree.com/#always-on-agents), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (prefix caching).

## Problem 7 — Making state survive and giving it homes

State must survive crashes and error loops, and different state types need different lifetimes. The system must remember across turns, crashes, and sessions — the survival half (P7) and the organization half (P8) are two patterns for one problem.

### Pattern 7 — State Snapshot & Rollback

**Forces.** Durability wants every step persisted; latency wants none. Exactly-once wants no re-execution; replay wants determinism. The conversation wants to be saved; the world wants to be left alone.

**Solution.** Save complete system state snapshots at checkpoint N, allowing recovery if an agent hits an error loop at step N+3. Completed steps never re-execute; in-flight steps resume from the last checkpoint.

**Consequences.** Crash-proof execution, audit by construction, and the agentic equivalent of a database transaction. This is condition 4 of the durable-daemons pattern made operational ([durable daemons](https://blog.hackspree.com/#durable-daemons-definition)).

**Tradeoffs.** Snapshots are only as good as their granularity: a snapshot of the *conversation* does not capture the *world* — an email sent, a payment moved, a file written outside the boundary. The [system boundary](https://blog.hackspree.com/#spatiotemporal-composability) argument applies with double force: recovery is promised exactly as wide as the reification. External side effects still require [idempotency keys](https://blog.hackspree.com/#durable-daemons-execution), because replay will re-fire them and only the external system can deduplicate — "a duplicate order loses capital. A duplicate email is embarrassing." And the rollback itself can be the failure: "recovery guarantees clean removal, but removal still has to be invoked."

**Evidence.** Temporal is the platform-grade instance — durable execution, workflow state, and deterministic replay ([Temporal](https://docs.temporal.io/)); DBOS turns a single Postgres write into a 1-2 ms checkpoint ([durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution)); the DeepSeek session log makes replay the everyday debugging tool — crash at step 80, recompute and continue from the same log ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Is the refactoring for A12 (state race conditions); composes with P14 (the blackboard needs the snapshot to be safe); its boundary is the [spatiotemporal](https://blog.hackspree.com/#spatiotemporal-composability) system boundary.

### Pattern 8 — Tiered Hierarchical Memory

**Forces.** Speed wants RAM; capacity wants disk. Currency wants the hot path; provenance wants the archive. Governance wants per-tier authority; simplicity wants one store.

**Solution.** Divide storage into immediate short-term context, scratchpad workspace, and long-term historical database storage.

**Consequences.** State has homes with different lifetimes — task ledgers, permissions, commitments, provenance, and triggers each fit a tier — and recall happens at the right latency. Modern frameworks formalize the tiers as thread state (short-term) plus a store (long-term), with semantic, episodic, and procedural memory as distinct stores ([LangGraph memory overview](https://docs.langchain.com/oss/python/langgraph/memory)).

**Tradeoffs.** More tiers mean more consistency work: what lives in the scratchpad that should have been promoted, what was deleted from short-term that the long-term still references, and — the hard one — what must be *forgotten* from all tiers including cached contexts and fine-tuned weights. Tiering fights the [append-only discipline](https://blog.hackspree.com/#deepseek-harness): a tier that rewrites itself is a tier where the past is negotiable. And the memory is only as good as its index: recall failures are silent, which is why the tiered pattern's most common degeneration is A5.

**Evidence.** Lilian Weng's canonical essay — short-term memory is in-context learning, long-term memory is an external vector store with fast retrieval ([LLM Powered Autonomous Agents](https://lilianweng.github.io/posts/2023-06-23-agent/)); the always-on survey's six axes — authority, scope, mutability, provenance, recoverability, actionability — are exactly the questions a tiered design must answer per tier ([always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Related.** Is the refactoring for A5; supplies the stores for P6 (the router); the governance gap it exposes is the [always-on](https://blog.hackspree.com/#always-on-agents) lifecycle.

### Anti-Pattern 5 — Goldfish Amnesia (State Blindness)

**Smell.** A multi-turn loop with no persisted state; identical tool calls repeated; the high-level goal forgotten mid-task.

**Anti-solution.** "The prompt has the goal" — statelessness as simplicity.

**Failure.** A stateless loop is not an agent; it is a request-response function with a longer prompt. Nothing records what was tried, so everything is tried again. The always-on survey's finding is the diagnosis: agents "accumulate state aggressively and have almost no machinery for unwinding it" — and the opposite failure, no state at all, repeats the same tool call ([always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Refactoring.** P8 (tiered memory) with the state lifecycle made real — write, validate, retrieve, update, forget — and the ledger of what was tried as the minimum viable memory. The durable daemons' conditions 2 and 3 are the specification: "A chatbot fails all three. An always-on agent satisfies them" ([durable daemons definition](https://blog.hackspree.com/#durable-daemons-definition)).

**Evidence.** LangGraph's persistent-state architecture is documented as the thing statelessness is the absence of ([LangGraph memory](https://docs.langchain.com/oss/python/langgraph/memory)); the always-on survey codes 435 papers and finds the governance half missing ([always-on agents](https://blog.hackspree.com/#always-on-agents)).

**Related.** Is the absence of P8; is the delegation risk of P13 (handoffs lose state).

**References for this problem.** Temporal ([docs](https://docs.temporal.io/)); Lilian Weng, LLM Powered Autonomous Agents ([post](https://lilianweng.github.io/posts/2023-06-23-agent/)); LangGraph memory overview ([docs](https://docs.langchain.com/oss/python/langgraph/memory)); archive: [durable daemons series](https://blog.hackspree.com/#durable-daemons) ([definition](https://blog.hackspree.com/#durable-daemons-definition), [execution](https://blog.hackspree.com/#durable-daemons-execution)), [always-on agents](https://blog.hackspree.com/#always-on-agents), [spatiotemporal composability](https://blog.hackspree.com/#spatiotemporal-composability) (system boundary).

## Problem 8 — Typing tool output and keeping errors legible

Model output is text; tools and downstream systems need types. Malformed values fail far from their cause, and swallowed errors become hallucinated successes. The boundary must be typed, and failures must stay legible.

### Pattern 9 — Schema Enforcement & Self-Correction

**Forces.** Flexibility wants free text; correctness wants types. Retry wants bounded cost; passing bad data wants none. The log wants append-only corrections; the loop wants no re-requests.

**Solution.** Force raw LLM text into JSON Schema, catching parsing failures and feeding structural fixes back internally.

**Consequences.** A typed contract at the harness boundary — the same contract as `--json` and exit codes on the way out ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)). Validation failures become part of the log: the correction is an append, not a rewrite, which is exactly what the [append-only model](https://blog.hackspree.com/#deepseek-harness) wants.

**Tradeoffs.** Schema enforcement costs tokens: every failed parse is a retry, and retries are where the bill grows — which is why the pattern must be paired with a cumulative retry budget. Strict schemas over-constrain open-ended output; loose schemas under-catch. And the retry loop is only as good as the error message: a validation error that does not tell the model *which field* failed produces the same failure again — A7 with extra steps.

**Evidence.** Instructor is the reference implementation — Pydantic validation with `max_retries` and `token_budget` bounding the correction cost ([Instructor](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)).

**Related.** Is the refactoring for A9 and A7; composes with the frozen-request pattern from the first edition ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

### Anti-Pattern 7 — The Silent Crash (Exception Swallowing)

**Smell.** API errors caught in the background; blank or generic strings returned to the agent; no stderr, no exit-code verdict.

**Anti-solution.** Catch-and-continue: "the agent doesn't need to know about the error."

**Failure.** A real error becomes a hallucinated success: the agent receives `""` or `"ok"`, cannot see the error, and confidently proceeds on a false premise. The three channels — stdout data, stderr diagnostics, exit code verdict — exist precisely so failures are legible, and swallowing them destroys the contract ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)).

**Refactoring.** P9 (schema enforcement & self-correction) with errors surfaced to the model with the details needed to correct, bounded by retry and token budgets ([Instructor error handling](https://python.useinstructor.com/concepts/retrying/)); layered checks so failures stay legible — "the planner chose the wrong tool" instead of "the agent failed" ([harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents)).

**Evidence.** Instructor's retry mechanics are the documented counter-pattern ([retrying](https://python.useinstructor.com/concepts/retrying/)).

**Related.** Is the absence of P9; feeds A9 (swallowed errors become malformed downstream values).

### Anti-Pattern 9 — The Schema Free-for-All

**Smell.** Complex arguments passed as raw strings; parsing deferred to the consumer; "the model formats it."

**Anti-solution.** Trust the model's output format.

**Failure.** The parsing problem moves downstream where there is no model to correct it: a malformed date fails in a database insert, a truncated JSON fails in a deserializer, and the error surfaces far from the agent that produced it — A7 with more steps.

**Refactoring.** P9 (schema enforcement at the harness boundary) with validation errors fed back while the model is still in the loop, and the corrections recorded in the append-only log.

**Evidence.** Pydantic exists because unvalidated strings become runtime failures one layer down ([Pydantic — core validation](https://docs.pydantic.dev/)).

**Related.** Is the absence of P9; feeds A7's downstream.

**References for this problem.** Instructor ([docs](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)); Pydantic core validation ([docs](https://docs.pydantic.dev/)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) (stdout/stderr/exit-code contract), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) (layered checks), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (frozen request).

## Problem 9 — Not blocking the loop on long tools

Some tools take minutes. The loop must stay live while the work happens elsewhere, and cancellation must be a contract. There is no named anti-pattern for this problem; its failure modes — the blocking loop and the abandoned worker — are covered in the tradeoffs.

### Pattern 10 — Asynchronous Tool Worker Queue

**Forces.** Responsiveness wants the loop unblocked; correctness wants the result. Asynchrony wants a tracking ID; consistency wants the work done. Cancellation wants a contract; the worker wants to finish.

**Solution.** Offload long-running processes to background task workers and hand a tracking ID to the looping agent.

**Consequences.** The loop stays live; the context holds a tracking ID instead of a worker's output, so the output does not occupy the window until it is ready. A worker queue is a daemon that satisfies all four durable-daemon conditions, and the tracking ID is the shared state another daemon observes ([durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution)).

**Tradeoffs.** Asynchrony introduces the consistency problems of A12: the agent may finish before the worker does, may poll the wrong ID, or may be torn down while the worker is still running — which is why cancellation must be a contract (the DeepSeek distinction between `ABORTED_BEFORE_DISPATCH` and `ABORTED`, so "cancellation never abandons the body"). The queue needs exactly-once or idempotent workers, because a retried task can double-execute. And every queue is infrastructure: brokers, workers, visibility timeouts, and dead-letter handling that the system team owns forever.

**Evidence.** Celery is the production standard for background task execution ([Celery — distributed task queue](https://docs.celeryq.dev/en/stable/getting-started/introduction.html)); DeepSeek's cancellation contract is the harness-side discipline ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Composes with P7 (the queue's workers need durability) and P14 (the blackboard's slow writers); its consistency risk is A12.

**References for this problem.** Celery — distributed task queue ([docs](https://docs.celeryq.dev/en/stable/getting-started/introduction.html)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (cancellation contract, around-dispatch wrappers), [durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution) (exactly-once, choreography).

## Problem 10 — Making the fast loop repeatable

Real APIs are slow, flaky, and costly; the system needs deterministic replay at scale. There is no named anti-pattern for this problem; its failure mode — the mock that drifts from production — is covered in the tradeoffs.

### Pattern 11 — Mock Tool Virtualization

**Forces.** Fidelity wants the real API; repeatability wants it frozen. Speed wants mocks; honesty wants drift detection. Determinism wants replay; realism wants live.

**Solution.** Swap production APIs out for lightweight mock responses inside development environments during multi-agent unit testing routines.

**Consequences.** Deterministic replay becomes possible — "a harness that cannot reproduce a run cannot measure a change" ([data-driven design](https://blog.hackspree.com/#data-driven-design-swe-agents)) — and the sample sizes data-driven design demands become affordable. This is the enabling condition for the [tasks-that-fight-back](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) fast loop.

**Tradeoffs.** A mock that drifts from production teaches the agent the wrong lessons — "if every test can be passed by pattern-matching the prompt, you are not measuring the assistant; you are measuring prompt luck." Mocks hide latency, nondeterminism, and rate limits; an agent trained only against mocks fails the first time it meets a 429. The pattern's honesty rule: mock for the unit test, record for the integration test, real for the eval.

**Evidence.** VCR.py is the record-and-replay reference — the first real call is recorded to cassette, subsequent runs replay it ([VCR.py](https://vcrpy.readthedocs.io/)); the harness canon codifies the layering — eval harnesses for fast filtering, task harnesses for realism, agent harnesses for integration ([harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents)).

**Related.** Composes with P10 (mocks are the workers' test doubles); is the fast-loop complement to F5 (live-environment evaluators, the slow loop).

**References for this problem.** VCR.py ([docs](https://vcrpy.readthedocs.io/)); archive: [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) (tasks that fight back, real repos and browsers), [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents) (deterministic replay, sample size).

## Problem 11 — Discovering capabilities without bloating the prompt

Tool spaces grow, and static lists bloat the prompt and confuse selection. Capabilities must be discoverable — and the frontier has standardized the discovery seam itself (F4).

### Pattern 12 — Dynamic Tool Discovery / Registry

**Forces.** Discoverability wants a catalog; the prompt wants it small. Dynamism wants reordering; the cache wants it fixed. Discovery wants open access; authorization wants a separate gate.

**Solution.** Store tool specs in databases, matching and surfacing capabilities dynamically to the agent based on semantic text queries.

**Consequences.** Many tools can exist behind small prompts; the registry becomes the single source of truth for what the system can do — and the type-graph mirror from the DeepSeek design keeps specs from drifting from implementations ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). The registry is also the interop seam: MCP standardizes how tools declare themselves, and the system serves the catalog to the agent mid-session ([Model Context Protocol](https://modelcontextprotocol.io/docs/learn/architecture)).

**Tradeoffs.** The registry is a new attack surface: every registered tool is a target for [prompt injection through tool descriptions](https://genai.owasp.org/llm-top-10/), and a registry that surfaces too many tools produces A8. Dynamic discovery fights prefix stability — tool lists that reorder by relevance invalidate the cache prefix (the DeepSeek cache trio includes "sort tool descriptions in one fixed order"). And the Toolformer lesson is the boundary: the model should learn *which* tool, but the system should not trust it with *all* tools — discovery and authorization are separate seams.

**Evidence.** Toolformer showed models can learn tool selection ([Toolformer](https://arxiv.org/abs/2302.04761)); the DeepSeek capability seam — Service Definition, Providers, Consumer behind a stable context key, with static start-time capability advertisement ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)) — is the pattern's canonical form; MCP is the protocol face ([architecture](https://modelcontextprotocol.io/docs/learn/architecture)).

**Related.** Is the refactoring for A8; composes with P2 (the gatekeeper authorizes what the registry discovered); is the runtime face of the capability-seam pattern from the first edition.

### Anti-Pattern 8 — Bloated Utility Belt (Tool Over-Provisioning)

**Smell.** Dozens of complex tools per agent; incorrect tool selection; the model "forgetting" tools exist.

**Anti-solution.** "More tools equals more capable."

**Failure.** Every tool in the prompt is a decision the model must make; a bloated registry turns tool selection into a needle-in-a-haystack retrieval problem that models lose. Toolformer's research line and SWE-agent's ACI work both found the same thing: fewer, better-shaped tools beat more tools.

**Refactoring.** P12 (the registry) with per-session tool compositions — DeepSeek's presets are per-session compositions, and code mode replaces the tool list with a generated SDK entirely ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). The SWE-agent result is the anchor: with the same model, a minimal well-designed ACI more than doubled state-of-the-art ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)).

**Evidence.** Toolformer ([paper](https://arxiv.org/abs/2302.04761)); SWE-agent's ACI ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)).

**Related.** Is the absence of P12; is the tool-scale form of A11 (the god agent).

### Frontier 4 — The Interop Layer: MCP, ACP, AGENTS.md

**Problem.** Every harness re-implements the same seams: how tools declare themselves, how editors and agents connect, where agent instructions live. The system needs the seams to be standard.

**Forces.** Interoperability wants open protocols; vendors want moats. Portability wants standards; integration wants control.

**Solution.** Adopt the emerging protocol layer: [MCP](https://modelcontextprotocol.io/docs/learn/architecture) for tool declaration and execution, [ACP](https://agentclientprotocol.com/) for editor-agent connection ("each editor must build custom integrations for every agent" stops being the default), and [AGENTS.md](https://agents.md/) — "a README for agents," used by over 60,000 open-source projects — for where agent instructions live.

**Consequences.** The seam becomes the ecosystem: DeepSeek runs unmodified Claude Code hooks and mounts rival harnesses as subagent providers rather than reimplementing them ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). This blog's DeepSeek teardown named the strategy [interop-as-product-strategy](https://blog.hackspree.com/#deepseek-harness): "when a vendor adopts its rivals' formats, your files become portable, whichever harness you run."

**Tradeoffs.** Every protocol you adopt is a contract you do not control, and every standard is a boundary where a substitution can hide — the supply-chain argument from [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface).

**Evidence.** MCP architecture ([docs](https://modelcontextprotocol.io/docs/learn/architecture)); ACP ([agentclientprotocol.com](https://agentclientprotocol.com/)); AGENTS.md ([agents.md](https://agents.md/)).

**Related.** Is the protocol face of P12 (the registry); is the interop layer for P13 (providers mount by standard).

**References for this problem.** Toolformer ([arXiv:2302.04761](https://arxiv.org/abs/2302.04761)); Model Context Protocol ([architecture](https://modelcontextprotocol.io/docs/learn/architecture)); Agent Client Protocol ([agentclientprotocol.com](https://agentclientprotocol.com/)); AGENTS.md ([agents.md](https://agents.md/)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (capability seams, Typert, interop-as-strategy), [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) (tool-selection research), [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface).

## Problem 12 — Dividing a workflow across agents

No single context can hold a whole workflow. The work must be divided, delegated, and handed off — and the frontier has added the contract that makes delegation safe (F2).

### Pattern 13 — Orchestrator-Worker (Boss-Worker)

**Forces.** Context limits want decomposition; coherence wants a single plan. Delegation wants specialists; control wants oversight. Termination wants a rule; conversation wants to continue.

**Solution.** Direct traffic using a highly capable central agent that delegates atomic sub-tasks to smaller, faster, specialized agents.

**Consequences.** Bounded contexts, parallelizable work, and a plan-then-execute shape. The orchestrator is the system's delegation tree curator — the workflow seam where the model writes orchestration scripts with `agent()` calls under caps ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)) — and the topology of the distillation pattern: a swarm of [tiny specialists coordinated by a router](https://blog.hackspree.com/#agents-are-distillation-at-scale).

**Tradeoffs.** The orchestrator is a single point of failure: if it mis-delegates, the error cascades downstream — Anthropic deliberately kept its planner's spec high-level because "if the planner tried to specify granular technical details upfront and got something wrong, the errors in the spec would cascade" ([harness design](https://www.anthropic.com/engineering/harness-design-long-running-apps)). Every delegation is a context handoff that can lose state (the A5 risk), and every worker is a new surface for A10 if the delegation loop has no termination condition. The exit condition must be part of the system, not the conversation.

**Evidence.** AutoGen is the reference framework — conversable, customizable agents composed into conversation patterns ([AutoGen](https://arxiv.org/abs/2308.08155)); Anthropic's three-agent architecture — planner, generator, evaluator — is the strongest recent instance ([harness design](https://www.anthropic.com/engineering/harness-design-long-running-apps)); DeepSeek's subagent registry mounts providers by name, including rival harnesses' own binaries ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)).

**Related.** Is the refactoring for A11; composes with P16 (the ensemble as the delegator's judge) and F2 (sprint contracts as the delegation's contract).

### Anti-Pattern 11 — The God Agent (Monolithic Orchestration)

**Smell.** One massive agent with an immense prompt managing every phase of a business workflow; every tool, every rule, every stage in one context window.

**Anti-solution.** "One agent, all context, total control."

**Failure.** The context avalanche and lost-in-the-middle degradation are guaranteed by construction: one window holding every phase, every tool, and every rule degrades exactly as the middle fills. The bloated utility belt is the tool-scale version of the same error.

**Refactoring.** P13 (orchestrator-worker) with decomposition and handoff: one feature at a time, structured artifacts carrying state between sessions. Anthropic's three-agent architecture — planner, generator, evaluator — is the modern template, and its planner deliberately stays high-level because "errors in the spec would cascade into the downstream implementation" ([harness design](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Evidence.** CrewAI's core argument for crews is that role-specialized agents with focused prompts beat one agent doing everything ([CrewAI — crews](https://docs.crewai.com/en/concepts/crews)).

**Related.** Is the absence of P13; composes A4 and A8.

### Frontier 2 — Sprint Contracts

**Problem.** Under-specified specs produce over- and under-built artifacts; "done" is discovered after the work, not before.

**Forces.** Specification wants precision; agility wants latitude. Verification wants criteria; creativity wants freedom.

**Solution.** Before each sprint, the generator and evaluator "negotiated a sprint contract: agreeing on what 'done' looked like for that chunk of work before any code was written." The generator proposes what it will build and how success will be verified; the evaluator reviews the proposal; the two iterate until they agree; only then does the generator build against the contract.

**Consequences.** The spec's ambiguity is resolved at the moment of maximum leverage — before the work exists — and the agent defines the verifier before it defines the artifact, which is the [tasks-that-fight-back](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) principle applied to the contract itself.

**Tradeoffs.** Negotiation overhead — two agents spending tokens agreeing on "done" before any code is written — which is why the contract is bounded to a sprint-sized chunk, not the whole project.

**Evidence.** Anthropic's harness ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Related.** Composes with P13 (the delegation's contract) and F1 (the evaluator grades against the contract).

**References for this problem.** AutoGen ([arXiv:2308.08155](https://arxiv.org/abs/2308.08155)); CrewAI crews ([docs](https://docs.crewai.com/en/concepts/crews)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (workflow seam, subagent registry), [Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents).

## Problem 13 — Coordinating through shared state without corruption

Multiple agents need a shared working memory, and unsynchronized writes corrupt it. The board must be shared and safe at once.

### Pattern 14 — Blackboard (Shared Workspace)

**Forces.** Sharing wants a unified schema; isolation wants per-agent scopes. Concurrency wants parallel writes; consistency wants ordered ones. Ownership wants a coordinator; choreography wants none.

**Solution.** Connect disjointed agents to a unified state schema where the harness coordinates simultaneous data writes and events.

**Consequences.** Choreography without an orchestrator — "each daemon is a process. Each daemon is durable. The choreography is the composition. No orchestrator. Just state" ([durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution)). Conflicts can be detected at write time rather than merge time — GitButler's virtual branches are the version-control face ([Buzz](https://blog.hackspree.com/#buzz-block-agents)). The blackboard is one of the oldest ideas in AI — the Hearsay-II speech architecture (1970s) — and the agent era reinvented it.

**Tradeoffs.** Shared state is shared risk. The blackboard concentrates A12: simultaneous writes without transactional locks corrupt the state, and the durability of the board decides whether the corruption is recoverable. The blackboard also fights [per-agent scope](https://blog.hackspree.com/#deepseek-harness) — the tension between isolation and choreography — and the resolution is explicit scope discipline: which state is shared, which is scoped, who owns the boundary. A blackboard without an owner is a tragedy of the commons for state.

**Evidence.** CrewAI's Flows layer formalizes the event-driven, state-managed workflow ([CrewAI — Flows](https://docs.crewai.com/en/concepts/flows)); the durable-daemons series specifies the choreography through shared state ([pattern specification](https://blog.hackspree.com/#durable-daemons-definition)); GitButler documents seven multi-agent collaboration patterns on virtual branches ([Buzz](https://blog.hackspree.com/#buzz-block-agents)).

**Related.** Composes with P7 (the board must be snapshotable) and P10 (the board's slow writers); its ownership question is the [always-on](https://blog.hackspree.com/#always-on-agents) authority axis.

### Anti-Pattern 12 — State Race Conditions

**Smell.** Asynchronous multi-agent writes to shared memory with no transactional locks; state corruption that appears "random."

**Anti-solution.** "It usually works" — hope as a concurrency strategy.

**Failure.** Two agents writing the same ledger, one overwriting the other's checkpoint, the state corrupting silently. The blackboard without its harness.

**Refactoring.** P7 (snapshots) with the durability discipline from the durable daemons: writes through a single ordered log, exactly-once within the control boundary, idempotency keys for external effects, and choreography through shared state with clear ownership — "no RPC. No message bus. No central orchestrator" is only safe when the shared state itself is the coordination mechanism, and the coordination mechanism must be transactional ([durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution)). GitButler's virtual branches detect conflicts at write time, not merge time ([Buzz](https://blog.hackspree.com/#buzz-block-agents)).

**Evidence.** Temporal's event-sourced, deterministic execution is the standard answer to the race ([Temporal](https://docs.temporal.io/)).

**Related.** Is the absence of P7 and P14's ownership discipline; is the async risk of P10.

**References for this problem.** CrewAI Flows ([docs](https://docs.crewai.com/en/concepts/flows)); Temporal ([docs](https://docs.temporal.io/)); archive: [durable daemons series](https://blog.hackspree.com/#durable-daemons) ([definition](https://blog.hackspree.com/#durable-daemons-definition), [execution](https://blog.hackspree.com/#durable-daemons-execution)), [Buzz and the Identity Problem](https://blog.hackspree.com/#buzz-block-agents) (GitButler virtual branches), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (per-agent scope), [always-on agents](https://blog.hackspree.com/#always-on-agents).

## Problem 14 — Keeping linear flows linear

Some flows are genuinely linear — classify, transform, emit — and the simplest verifiable shape is the right one. There is no named anti-pattern for this problem; its failure mode — forcing non-linear flows into pipelines — is covered in the tradeoffs.

### Pattern 15 — Sequential Pipeline Routing

**Forces.** Simplicity wants rigid stages; adaptivity wants branching. Determinism wants fixed routing; robustness wants recovery. Billing wants few calls; quality wants specialized hops.

**Solution.** Pass analytical payloads through rigid linear stages, using LLMs solely for classification or transformation tasks at each hop.

**Consequences.** The easiest harness to verify, replay, and bill: a linear stage where each hop has one job and a deterministic contract, in the spirit of "the most successful implementations use simple, composable patterns rather than complex frameworks" ([Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents)). Each hop is a tool with a `--json` contract, exit codes, and honest `--help` ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)).

**Tradeoffs.** Rigid pipelines cannot route around failure: a stage that misclassifies poisons every downstream stage, and there is no re-planning. Pipelines serialize — the whole flow runs at the speed of its slowest LLM hop, and each hop is a full request ([token economics](https://blog.hackspree.com/#every-token-has-a-price-tag)). The pattern's own lineage is explicit about the boundary: chains are right when the flow is genuinely linear, wrong when it needs to branch, loop, or recover — which is when P13 or a graph shape wins.

**Evidence.** LangChain's chains are the canonical instance, and the migration-era docs are explicit about when the rigid linear form is right and when it is not ([LangChain Chains](https://python.langchain.com/docs/versions/migrating_chains/overview/)).

**Related.** Is the disciplined baseline; is the opposite of A11 at orchestration scale; composes with P9 (each hop's output is schema-enforced).

**References for this problem.** LangChain chains ([docs](https://python.langchain.com/docs/versions/migrating_chains/overview/)); Anthropic, Building Effective Agents ([post](https://www.anthropic.com/engineering/building-effective-agents)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design), [Loop Engineering](https://blog.hackspree.com/#loop-engineering), [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag).

## Problem 15 — Producing a verdict no single model can fake

One judgment is unreliable, and self-evaluation is reliably lenient. The system must decide what is good without depending on a single model's opinion — and the frontier has evolved the gate into an iterative critic (F1) with hands (F5).

### Pattern 16 — Voting / Consensual Ensemble

**Forces.** Reliability wants multiple judgments; cost wants one call. Independence wants diverse members; convenience wants one family. Agreement wants a signal; correctness wants ground truth.

**Solution.** Query multiple independent model setups with identical prompts, using harness code to calculate majority agreement on outputs.

**Consequences.** A statistical signal where [verifiers are king](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc): multiple independent judgments aggregated by system code instead of one verdict from one model. Majority agreement is a distribution over independent samples — "a single run is a data point; a thousand runs is a distribution" ([data-driven design](https://blog.hackspree.com/#data-driven-design-swe-agents)) — and it replaces self-review with cross-review.

**Tradeoffs.** Ensembles multiply cost linearly and can multiply latency; the [cost per finished task varies 7x across harnesses](https://blog.hackspree.com/#deepseek-harness) before you add an ensemble. Majority agreement is only a strong signal when the members are *independent* — correlated members (same family, same prompt shape, same failure mode) vote as one and add nothing. And the reference is honest about the limit: human-vote consensus measures preference, not correctness, and an LLM-judge ensemble inherits every bias of its members — "algorithmic verification misses context. Agentic verification hallucinates." The ensemble is a signal, not a ground truth.

**Evidence.** LMSYS Chatbot Arena / LMArena is the reference for agreement-based evaluation at scale — millions of pairwise votes forming the consensus signal ([LMArena leaderboard](https://lmarena.ai/leaderboard)); the mob post's self-review datum motivates it — 79% of 25,264 agent PRs reviewed by the prompting developer ([mob programming remastered](https://blog.hackspree.com/#mob-programming-reimagined)).

**Related.** Composes with P13 (the ensemble as the delegator's judge); is the aggregation answer to A10; is the statistical cousin of F1.

### Anti-Pattern 10 — The Committee Paradox (Infinite Multi-Agent Debates)

**Smell.** Two or more autonomous agents reviewing each other's work; no exit condition; the conversation "making progress" without converging.

**Anti-solution.** "More debate equals better decisions."

**Failure.** Two agents reviewing each other with no exit condition is an infinite loop with a nicer name ([Loop Engineering](https://blog.hackspree.com/#loop-engineering)). Tokens burn, no verdict arrives, and the system never terminates.

**Refactoring.** Termination as a system property: a termination condition, a verdict threshold, a budget, or a human breakpoint — and P16's aggregation rule as the shape of disagreement resolution. Review without independence is theater, and review without termination is waste ([mob programming remastered](https://blog.hackspree.com/#mob-programming-reimagined)).

**Evidence.** AutoGen's framework treats termination as a first-class design concern of multi-agent conversation; the community's infinite-loop reports are the empirical record ([AutoGen](https://arxiv.org/abs/2308.08155)).

**Related.** Is the absence of P13 and P16's termination discipline; is the orchestration form of A2 (the vortex).

### Frontier 1 — The Generator–Evaluator Loop

**Problem.** Self-evaluation is reliably lenient: "when asked to evaluate work they've produced, agents tend to respond by confidently praising the work — even when, to a human observer, the quality is obviously mediocre" ([Anthropic](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Forces.** Feedback wants independence; convenience wants self-review. Iteration wants speed; quality wants depth. Cost wants few cycles; taste wants many.

**Solution.** Separate the agent doing the work from the agent judging it — a GAN-inspired generator-evaluator loop where the evaluator grades against explicit criteria and the critique flows back as the next iteration's input.

**Consequences.** "The separation doesn't immediately eliminate that leniency on its own; the evaluator is still an LLM that is inclined to be generous towards LLM-generated outputs. But tuning a standalone evaluator to be skeptical turns out to be far more tractable than making a generator critical of its own work." The frontend experiment ran 5-15 iterations per generation; the full-stack version ran a planner, a generator, and an evaluator for six hours and produced an application the solo run could not approach.

**Tradeoffs.** The harness was "over 20x more expensive" — $200 for six hours against $9 for twenty minutes, same model — and the pattern is only worth it when the output-quality difference justifies the bill. This is the [verifiers are king](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) argument with the verifier given agency: not a gate that blocks, but a critic that iterates.

**Evidence.** Anthropic, *Harness design for long-running application development* (March 2026) ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Related.** Composes with P16 (the ensemble as the evaluator); pairs with F2 (the contract the evaluator grades against).

### Frontier 5 — Live-Environment Evaluators

**Problem.** Static verification misses what only use reveals: the artifact "looks impressive but still has real bugs when you actually try to use them" ([Anthropic](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Forces.** Depth wants live interaction; cost wants static scoring. Reality wants the real system; CI wants speed.

**Solution.** Give the evaluator hands: in Anthropic's harness the evaluator was given the Playwright MCP, "which let it interact with the live page directly before scoring each criterion and writing a detailed critique" — navigating the app, clicking through UI features, testing API endpoints and database states "the way a user would."

**Consequences.** The verifier does not read the artifact, it *uses* it — the frontier version of this blog's [harnesses should observe the whole system](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) principle. Only a verifier with hands finds the solo run's broken entity wiring.

**Tradeoffs.** Wall-clock: "because the evaluator was actively navigating the page rather than scoring a static screenshot, each cycle took real wall-clock time. Full runs stretched up to four hours." The pattern belongs in the slow loop of the [eval → task → agent harness taxonomy](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents), reserved for the final gate.

**Evidence.** Anthropic's harness ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)).

**Related.** Is the slow-loop complement of P11 (mocks); composes with F1 (the evaluator's tools).

**References for this problem.** LMArena leaderboard ([lmarena.ai](https://lmarena.ai/leaderboard)); AutoGen ([arXiv:2308.08155](https://arxiv.org/abs/2308.08155)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), [Zuill's Mob Programming, Remastered](https://blog.hackspree.com/#mob-programming-reimagined) (self-review), [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents), [Loop Engineering](https://blog.hackspree.com/#loop-engineering).

# Part II — How the patterns compose — and where they conflict

The catalog is not a menu; the patterns interact across problems, and some of the interactions are contradictions that have to be managed explicitly.

**The memory problem contains its own war.** The [append-only log](https://blog.hackspree.com/#deepseek-harness) (first edition), P5 (rolling window compression), and P8 (tiered memory) disagree about whether the past can be rewritten. The resolution is layering: the log is immutable, the derived view is compressed, the tiers are projections — and the [compaction prefix bug](https://blog.hackspree.com/#deepseek-harness) is what happens when the layers touch.

**Prefix-cache discipline constrains every problem.** The [Bill as Assertion](https://blog.hackspree.com/#every-token-has-a-price-tag) requirement — no volatile data in the prompt prefix, fixed tool ordering, append-only history — constrains P6 (inject context where it cannot disturb the prefix), P12 (discovery must not reorder the tool list), and P5 (summaries must be prefix-extensions). The interface discipline of the [agentic-first CLI](https://blog.hackspree.com/#agentic-first-cli-design) — "no timestamps unless asked" — is a billing requirement before it is a UX nicety.

**Isolation and coordination pull against each other.** P1 and [per-agent scope](https://blog.hackspree.com/#deepseek-harness) want isolation; P14 and the [durable daemons choreography](https://blog.hackspree.com/#durable-daemons-definition) want shared state as the coordination mechanism. The resolution is explicit scope discipline — which state is shared, which is scoped, who owns the boundary — the least-solved governance question in the [always-on survey](https://blog.hackspree.com/#always-on-agents).

**Safety and speed are the same budget.** P2, P3, and P4 each add latency and tokens to every step; the [verification layer](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) adds more. The token post's arithmetic — the bill is 2-14% of labor cost — is the budget all of them share, and the systems that win make the checks *be* the measurement data: layered checks that fail legibly are both the gate and the instrument.

**Defense in depth disagrees with itself.** The [sandbox stack](https://blog.hackspree.com/#sandboxing-ai-agents) says add layers; [zero overhead](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) says subtract them. Both are right, and the resolution is honesty about what each layer is for: the sandbox constrains the compromised agent, the small dependency graph shrinks what can be compromised, the gatekeeper mediates capabilities — none of them is a wall, and the pattern is to know which one you are building.

**Problems interact across the linear order.** The problems are ordered so that later ones assume earlier ones are solved: bounded loops (P4) make schema retries (P9) affordable; snapshots (P7) make the blackboard (P14) safe; decomposition (P13) makes verification (P15) tractable; containment (P1) makes everything else survivable. A system that skips a problem does not skip its cost — it pays the cost as one of the anti-patterns.

The meta-pattern across all of it is the one the DeepSeek teardown's two independent analyses converged on: **the system is the product.** The same model swinging from 47% to 67% across eight harnesses, harness-level improvements lifting scores without any training, the same weights 3.4 points apart under two harnesses — the wrapper is where the leverage is, which is why the wrapper deserves a pattern language, and an anti-pattern catalog, at all.

# Part III — What to steal regardless

The catalog above is drawn from this blog's archive and from primary sources verified for this edition. The portable items, independent of any particular stack:

**From the pattern half:** the ephemeral sandbox wrapper with honest teardown (P1); the intercepting gatekeeper with monotonic, fail-closed denial (P2); the breakpoint as a persisted state primitive, not a prompt (P3); hard token/time/monetary ceilings enforced in code, with the bill as a CI assertion (P4); compression as a projection over an immutable log (P5); retrieval with routing judgment, injected where it cannot bury the instruction (P6); checkpoints with idempotency keys at the boundary (P7); tiers with explicit authority, scope, and forgetting per tier (P8); schema enforcement with bounded self-correction (P9); async workers with cancellation as a contract (P10); mocks for the fast loop and real environments for the final gate (P11); registries whose specs cannot drift from implementations (P12); delegation with system-defined termination (P13); shared state with transactional ownership (P14); linear stages with deterministic contracts (P15); and ensembles of independent judgments as a signal, never a ground truth (P16).

**From the anti-pattern half:** the twelve names above are a diagnostic checklist. When an agentic system behaves inexplicably, ask which problem the system skipped: Is the model trusted with keys and told to be careful (A1, problem 1)? Is there no ceiling on the loop (A2, problem 4)? Is authorization a sentence in the prompt instead of a check in the tool (A3, problem 2)? Is the context a dump instead of a derivation (A4, problem 5)? Is the loop stateless (A5, problem 7)? Is retrieval drowning the instruction (A6, problem 6)? Are errors being swallowed into hallucinated success (A7, problem 8)? Is the tool belt bloated (A8, problem 11)? Are complex arguments untyped (A9, problem 8)? Are agents debating without an exit condition (A10, problem 15)? Is one agent doing everything (A11, problem 12)? Is shared state being written without locks (A12, problem 13)?

**From the frontier:** separate the generator from the evaluator, because self-evaluation is reliably lenient (F1, problem 15); negotiate the contract before the code, because "done" agreed in advance beats "done" discovered after (F2, problem 12); reset the context and hand off structured artifacts instead of compacting in place when the model exhibits context anxiety (F3, problem 5); implement the interop standards, because the seam is the ecosystem (F4, problem 11); give the verifier hands, because the artifact that looks impressive is not the artifact that works (F5, problem 15); and engineer the context, not the prompt, because the behavior is a function of the whole context state (F6, problem 5).

The tradeoffs are not defects in the patterns; they are the prices. Every pattern in this catalog buys a guarantee and charges a cost, and the honest engineer is the one who can state both for each layer of the system. The guarantee stops where the declaration stops — file effects only, bash-equivalent trust, observational equivalence, within the control boundary, majority agreement is not ground truth. Name the boundary, design within it, and verify the design works. That is the engineering method, and it is the whole pattern language in one sentence.

# References

## Pattern-language foundations

- Alexander, C., Ishikawa, S., Silverstein, M. [A Pattern Language: Towns, Buildings, Construction](https://en.wikipedia.org/wiki/A_Pattern_Language) (Oxford University Press, 1977) — 253 patterns "written as a set of problems and documented solutions," cross-referenced into a network; the origin of the problem-first pattern-language form used here.
- Gamma, Helm, Johnson, Vlissides. [Design Patterns: Elements of Reusable Object-Oriented Software](https://en.wikipedia.org/wiki/Design_Patterns) (Addison-Wesley, 1994) — the GoF catalog that made named patterns with forces and consequences the standard form in software.

## Primary sources (grouped by problem; each problem unit also carries its own references)

**Problems 1–4 (safety).** LangChain sandbox integrations ([docs](https://python.langchain.com/docs/integrations/)); Meta AI, Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); LangGraph interrupts ([docs](https://docs.langchain.com/oss/python/langgraph/interrupts)); AutoGPT configuration ([docs](https://docs.agpt.co/classic/configuration/)) and issue tracker ([issues](https://github.com/Significant-Gravitas/AutoGPT/issues)); OWASP Top 10 for LLM Applications ([2025](https://genai.owasp.org/llm-top-10/)); Willison, prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); archive: [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents), [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface), [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), [OpenWorker outcome layer](https://blog.hackspree.com/#openworker-outcome-layer).

**Problems 5–7 (memory & state).** Packer et al., MemGPT ([arXiv:2310.08560](https://arxiv.org/abs/2310.08560)); Liu et al., Lost in the Middle ([arXiv:2307.03172](https://arxiv.org/abs/2307.03172)); Pinecone RAG ([learn](https://www.pinecone.io/learn/retrieval-augmented-generation/)) and advanced RAG ([learn](https://www.pinecone.io/learn/advanced-rag/)); Temporal ([docs](https://docs.temporal.io/)); LangGraph memory ([docs](https://docs.langchain.com/oss/python/langgraph/memory)); Weng, LLM Powered Autonomous Agents ([post](https://lilianweng.github.io/posts/2023-06-23-agent/)); Anthropic, Effective context engineering for AI agents ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)); archive: [durable daemons series](https://blog.hackspree.com/#durable-daemons), [always-on agents](https://blog.hackspree.com/#always-on-agents), [spatiotemporal composability](https://blog.hackspree.com/#spatiotemporal-composability), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness).

**Problems 8–11 (tools & interfaces).** Instructor ([docs](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)); Pydantic ([docs](https://docs.pydantic.dev/)); Celery ([docs](https://docs.celeryq.dev/en/stable/getting-started/introduction.html)); VCR.py ([docs](https://vcrpy.readthedocs.io/)); Schick et al., Toolformer ([arXiv:2302.04761](https://arxiv.org/abs/2302.04761)); Model Context Protocol ([architecture](https://modelcontextprotocol.io/docs/learn/architecture)); Agent Client Protocol ([agentclientprotocol.com](https://agentclientprotocol.com/)); AGENTS.md ([agents.md](https://agents.md/)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents), [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness).

**Problems 12–15 (orchestration & verification).** Wu et al., AutoGen ([arXiv:2308.08155](https://arxiv.org/abs/2308.08155)); CrewAI crews ([docs](https://docs.crewai.com/en/concepts/crews)) and flows ([docs](https://docs.crewai.com/en/concepts/flows)); LangChain chains ([docs](https://python.langchain.com/docs/versions/migrating_chains/overview/)); LMArena leaderboard ([lmarena.ai](https://lmarena.ai/leaderboard)); Anthropic, Building Effective Agents ([post](https://www.anthropic.com/engineering/building-effective-agents)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); OpenAI, Harness Engineering ([post](https://openai.com/index/harness-engineering/)); archive: [verifiers-are-king](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), [mob programming remastered](https://blog.hackspree.com/#mob-programming-reimagined), [Loop Engineering](https://blog.hackspree.com/#loop-engineering), [durable daemons execution](https://blog.hackspree.com/#durable-daemons-execution), [Buzz and the Identity Problem](https://blog.hackspree.com/#buzz-block-agents), [agents-are-distillation-at-scale](https://blog.hackspree.com/#agents-are-distillation-at-scale).

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
- [Zuill's Mob Programming, Remastered](https://blog.hackspree.com/#mob-programming-reimagined) — the self-review problem (79% of 25,264 agent PRs), driver/navigators as review architecture, Minsky's Society of Mind.
- [Loop Engineering is what the NATO conference asked for in 1968](https://blog.hackspree.com/#loop-engineering) — loops over prompts, "the simulation becomes the system."
- [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer) — deliverables over chat, approval gates, personas as configuration.
- [The Factory Is Not Dead](https://blog.hackspree.com/#factory-is-not-dead) — Bemer's 1968 "software factory," the factory framing the harness catalog inherits.
