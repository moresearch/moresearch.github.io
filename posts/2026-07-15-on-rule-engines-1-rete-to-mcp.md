---
title: On Rule Engines — From RETE to MCP
date: 2026-07-15
slug: on-rule-engines-rete-to-mcp
summary: Tracing a 40-year arc from Charles Forgy's RETE algorithm to the Model Context Protocol, this essay explores how rule engines and large reasoning models answer the same fundamental question through different mechanisms — and why the most important trend is not replacement but convergence.
tags: llm, rules, enterprise, automation, composite-ai, symbolic-ai, rete, mcp, fsm
series: on-rule-engines
---

Enterprise decision automation sits at an uncomfortable intersection. Large Language Models promise natural interaction and flexibility. Business decisions demand auditability, determinism, and compliance with regulations that do not negotiate. The question is not which technology wins. The question is how to combine them — and the answer has roots going back nearly half a century.

Pierre Feillet addressed this in his 2023 article [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f) and his 2026 follow-up [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). The second article traces a 40-year arc from Charles Forgy's RETE algorithm to the Model Context Protocol and argues that the most important trend in enterprise AI is not replacement but convergence.

This essay — the first in a three-part series — begins where the story begins: with the RETE algorithm, the architecture of rule engines, the rise of large reasoning models, and the fundamental distinction between finite state machines and rule engines that every system architect should understand.

## The RETE algorithm (1979)

In 1979, Charles L. Forgy, a PhD student at Carnegie Mellon University, solved a problem that would shape the next four decades of enterprise computing. Given a large set of IF-THEN rules and a working memory full of facts, how do you efficiently determine which rules should fire?

The naive approach — test every rule against every fact on every cycle — is catastrophically slow. Ten thousand rules against a hundred thousand facts is a billion comparisons per cycle. Production systems in the 1970s ground to a halt under realistic workloads.

Forgy's insight was that working memory changes slowly between cycles. Typically only a few facts are added or removed at each step. Re-evaluating every rule from scratch means repeating vast amounts of work that has not changed. What if you could *remember* partial matches and only recompute what actually changed?

The result was the RETE algorithm — Latin for "net" — published in Forgy's 1979 PhD thesis and in a landmark 1982 paper in *Artificial Intelligence*: [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0).

### How it works

RETE has two phases. **Compilation:** rules are compiled into a discrimination network — a directed acyclic graph where each node tests a condition. Conditions appearing in multiple rules are compiled once and shared. If twenty rules check `customer.tier == "premium"`, there is one node, not twenty. Forgy defined four node types: root nodes (entry points for facts), one-input nodes (test a single condition), two-input nodes (join facts that satisfy different patterns), and terminal nodes (all patterns satisfied — fire the rule).

**Runtime:** facts enter at the root and propagate through the network. Each node caches the facts — or partial matches, at join nodes — that satisfy its condition. When a new fact arrives, only the branches affected by that fact are re-evaluated. When a fact is removed, its matches are retracted from the caches. The algorithm trades memory for speed — storing partial match state at every node — and achieves performance that is, in Forgy's words, theoretically independent of the number of rules.

### The lineage

RETE became the core of **OPS5**, which powered **R1/XCON** — an expert system that configured VAX computer orders for Digital Equipment Corporation, one of the first commercially successful AI systems, reportedly saving DEC millions of dollars per year by catching configuration errors that human order processors missed.

It went on to become the backbone of production rule systems across the industry: **CLIPS** (NASA's C Language Integrated Production System), **Jess** (the Java Expert System Shell), **Drools** (open-source, now Apache, which evolved RETE into ReteOO and later PHREAK), **IBM Operational Decision Manager** (enterprise-grade decision automation with governance, versioning, and deployment pipelines), **Soar** (the cognitive architecture), **Blaze Advisor**, and **TIBCO BusinessEvents**.

Banks used it for loan origination. Insurers for claims adjudication. Governments for eligibility determination. The RETE network was invisible infrastructure, humming inside systems that made consequential decisions about people's money, health, and legal status.

And then, around 2022, the world became captivated by a different kind of AI — one that produced fluent text rather than deterministic decisions. The question became: are rule engines obsolete?

Feillet's 2026 answer is unequivocal: no. They never died. They are running alongside your LLM right now.

### Beyond RETE

Modern rule engines have also evolved beyond the discrimination network. **[RuleGo](https://github.com/rulego/rulego)** — an open-source rule engine in Go (Apache 2.0) — uses a Directed Acyclic Graph rather than a RETE network. Business logic is composed of component nodes wired into rule chains; messages flow along predetermined DAG paths rather than being matched against all rules. This trades RETE's expressive forward-chaining for deterministic execution with extremely low resource consumption — ~19 MB of memory under 500 concurrent requests on a Raspberry Pi 2. Both approaches have their place, and both are finding their way into composite AI architectures.

## Two ways to answer the same question

Feillet's 2026 article opens with a simple observation: both rule engines and Large Reasoning Models address the identical problem — *given what we know, what should we conclude?* The methods differ significantly. The objective is the same. This is more than a philosophical point. It means these technologies can and should be compared, contrasted, and combined.

### How rule engines reason

Rule engines separate knowledge from execution. Business logic is expressed as declarative condition-action pairs (IF/THEN structures). Facts enter working memory. The inference engine continuously evaluates which rules are satisfied. When multiple rules fire simultaneously, a conflict resolution mechanism called the *agenda* determines firing order.

The key architectural property is that the rule author specifies *what* should happen; the engine handles *how*. This separation is the source of the rule engine's transparency. Every inference step is explicit. An auditor can trace precisely which rule fired, what triggered it, and why a decision was reached. Rule engines do not generate post-hoc rationalizations. They produce *proofs* — logical derivations that are mechanically verifiable.

### How Large Reasoning Models reason

LRMs store knowledge not as explicit rules but as distributed numerical representations learned across billions of training examples. When solving problems, they navigate a high-dimensional conceptual space rather than pattern-matching against a predicate database. Their reasoning traces may look like rule chaining, but as Feillet puts it, this reasoning is "an approximation of modus ponens" executed through statistical pattern completion rather than logical evaluation.

The power is extraordinary flexibility — generalization across novel inputs, ambiguity handling, cross-domain reasoning without explicit programming.

The fragility is hallucination — producing plausible-sounding but logically invalid conclusions because, as Feillet puts it, "plausibility and logical validity are different things."

### The core distinction

**Rule engines guarantee correctness relative to their rules; LRMs generate probabilistically plausible conclusions.** One yields proofs, the other rationales. Both are useful. Confusing them is dangerous.

Feillet gives a concrete example from the 2023 article: a pizza ordering bot from DeepLearning.ai's tutorial that "depending on the runs, provides the expected outcome or a surprising one." Wrong pizza toppings are annoying. Wrong loan decisions are measured in lawsuits and regulatory actions. The stochastic parrot is not a bug — it is the architecture. For creative tasks, probabilistic variation is a feature. For mortgage decisions, it is a liability.

## The hybrid inference loop

Modern reasoning models have evolved beyond raw token prediction. When a calculation is needed, the model can generate and execute Python code — an exact, deterministic result fed back into the reasoning context. This is a pattern: stochastic reasoning delegating to deterministic execution at moments where precision matters.

The Model Context Protocol formalizes this delegation:

1. Reason with available context
2. Identify what is uncertain or requires precision
3. Delegate to a tool
4. Receive facts
5. Continue reasoning

Feillet draws a structural parallel to how production rule systems invoke external actions and continue with newly acquired working memory — except the orchestration is now performed by a neural model rather than a symbolic agenda. IBM ODM and Decision Intelligence expose their decision services via MCP, enabling a reasoning model to invoke a full rule engine at the precise point in its inference trace where governed, auditable, deterministic decisions are needed. The LRM handles interpretation, planning, and contextual understanding; the rule engine handles logic requiring correctness, traceability, and compliance.

### Attention and RETE: same purpose, different mechanism

Feillet is careful not to equate the two, but he notes a shared architectural role:

- **RETE:** symbolic pattern matching with Boolean conditions, exact matches, discrete sets of eligible rules
- **Attention:** differentiable relevance computation with continuous weights, weighted sums of representations
- **Commonality:** both determine which information is most relevant for the next inference step — one operating on facts and predicates, the other on learned representations

This is not a coincidence. Both architectures face the same fundamental challenge: given a large set of knowledge and a specific context, what is relevant *right now*? RETE solves it by pre-compiling condition tests into a network and caching partial matches. Attention solves it by learning to compute relevance scores from data. The mechanisms are different. The function is the same.

### Converging trends

Feillet identifies three:

1. **Rule engines incorporating generative AI** for authoring, explaining, and maintaining policies in natural language
2. **Reasoning models incorporating deterministic subsystems** — code execution, retrieval, decision services — to ground outputs in verifiable reality
3. **Hybrid inference architectures** where stochastic reasoning orchestrates deterministic reasoning, and neural inference delegates to symbolic inference at the moments that matter

### RuleGo: the convergence in a Go library

The convergence is not confined to enterprise platforms. **[RuleGo](https://github.com/rulego/rulego)** is a modern rule engine that already embodies the hybrid architecture. Its DAG-based component model treats an LLM call as just another node in the rule chain graph — a filter, a transformer, an HTTP push, an LLM intent extractor, all share the same interface. The **[rulego-components-ai](https://github.com/rulego/rulego-components-ai)** extension provides LLM integration and MCP server/client support.

A RuleGo instance can expose its rule chains as MCP tools, making them discoverable and invocable by an LLM-based reasoning model. Conversely, RuleGo can act as an MCP client, calling out to LLM endpoints when its rule chain logic requires it. The hybrid inference loop Feillet describes — reason, identify uncertainty, delegate, receive, continue — is directly implementable as a RuleGo rule chain alternating between deterministic component nodes and MCP-mediated LLM calls.

The significance is not that RuleGo replaces IBM ODM. It is that the composite AI pattern has become general enough to appear in a Go library running on a Raspberry Pi. The 40-year arc from RETE to MCP passes through `go get github.com/rulego/rulego`.

## Finite state machines vs. rule engines

Before moving to composite AI patterns — the subject of the second essay in this series — there is an architectural distinction worth making explicit: the difference between finite state machines and rule engines, and why confusing them leads to brittle systems.

A **finite state machine** answers the question *what happens next?* It encodes a fixed set of states and explicit transitions between them. You are in state S. Event E arrives. If a transition from S on E is defined, you move to the target state. The FSM knows where you are, what you can do, and where you go next. Its logic is procedural: the sequence matters, the state matters, the transition guard matters.

A **rule engine** answers the question *what should we conclude?* It encodes condition-action pairs evaluated against a working memory of facts. Any rule whose conditions are satisfied is eligible to fire. The rule engine does not care about sequence — it cares about satisfaction. Its logic is declarative: the facts matter, the conditions matter, the conclusions matter. Order of rule evaluation is the engine's concern, not the author's.

The distinction is sharp:

| Dimension | FSM | Rule Engine |
|---|---|---|
| Core question | What happens next? | What should we conclude? |
| Logic style | Procedural | Declarative |
| State model | Explicit states + transitions | Working memory of facts |
| Control flow | Defined by transition graph | Defined by rule firing (agenda) |
| Best for | Workflows, protocols, pipelines | Policies, decisions, classifications |
| Determinism | Deterministic given state + event | Deterministic given rule set + facts |
| Complexity grows with | Number of states × transitions | Number of rules × fact combinations |

The mistake is using one where the other belongs. Encoding a loan eligibility policy as an FSM produces an explosion of states — one for every combination of credit score band, debt ratio, collateral type, and regulatory jurisdiction. Encoding a multi-step claims workflow as a flat rule set produces fragile priority chains (`rule-1: if step=pending_review then...`) where the implicit process state is smuggled through working memory facts. Each approach can express the other's domain, but the expression is awkward, verbose, and hard to maintain.

The rule of thumb: if the problem is *procedural* — steps, stages, sequences, approvals, handoffs — reach for an FSM or a workflow engine. If the problem is *decisional* — eligibility, pricing, risk, compliance, classification — reach for a rule engine. Most real business processes are both, and that is the subject of the third essay.

## The 40-year arc

<img src="images/composite-ai-fig6-overview.svg" alt="Composite AI — Blending Neuronal and Symbolic Approaches" style="width:100%;max-width:720px;">

The story from RETE to MCP is not obsolescence and replacement. It is infrastructure that works being augmented by capabilities that are new. The RETE network matches facts against rules with Boolean precision. The attention mechanism computes relevance over learned representations. Both answer the same question — *what is relevant right now?* — at different layers of the stack.

The next essay examines the five composite AI patterns that Feillet and his co-authors proposed in 2023: concrete architectures for combining neuronal and symbolic AI in production systems.

---

**References**

1. Charles L. Forgy. [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0). Artificial Intelligence, 19(1): 17–37, 1982.

2. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

3. Pierre Feillet. [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). Medium, June 4, 2026.

4. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0. Includes [rulego-components-ai](https://github.com/rulego/rulego-components-ai) for LLM integration and MCP support.

> *Continue to Part 2: [On Rule Engines: Five Patterns for Composite AI](#on-rule-engines-five-patterns)*
