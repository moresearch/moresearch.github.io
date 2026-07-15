---
title: On Rule Engines — Five Patterns for Composite AI
date: 2026-07-15
slug: on-rule-engines-five-patterns
summary: Five architectural patterns for blending LLMs with rule-based decision engines — NLU to rules, rules to NLG, rules orchestrating LLMs, LLM-driven rule extraction, and chatbot delegation. Each pattern answers a different question about where the LLM sits in the decision pipeline and what role it plays.
tags: llm, rules, enterprise, automation, composite-ai, decision-making, symbolic-ai, fsm
series: on-rule-engines
---

The first essay in this series traced the 40-year arc from Forgy's RETE algorithm to the Model Context Protocol and established the core distinction: rule engines produce proofs, large reasoning models produce rationales. Both are useful. Neither alone suffices for production AI.

This essay — the second in the series — examines the five architectural patterns that Pierre Feillet, Allen Chan, Luigi Pichett, and Yazan Obeidi proposed in their 2023 article [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). They let an architect reason about *where* in the pipeline the LLM should sit and *what role* it should play. Each pattern is explained with its tradeoffs, failure modes, and the hidden engineering work that makes it production-ready.

But first: what does an enterprise decision actually require?

## What an enterprise decision requires

Feillet's article enumerates eight criteria. They explain why the LLM-only approach keeps hitting a wall in production.

1. **Accuracy.** A loan decision wrong 2% of the time is not 98% accurate — it is a regulatory finding and a lawsuit. Correctness means *every time*.
2. **Scalability.** Millions of claims per day cannot degrade when the rule base grows to tens of thousands of rules. RETE's performance-independence from rule count matters here.
3. **Adaptability.** Regulations and policies change. The system must accommodate new rules without a multi-month DevOps cycle.
4. **Latency.** Budgets vary — milliseconds for fraud, seconds for pre-approval, minutes for underwriting — but must be predictable. LLM latency is variable; rule engine latency is not.
5. **Auditability.** "The model's attention weights converged on that outcome" is not an acceptable answer to a regulator. Rule engines produce proofs, not rationales.
6. **Privacy.** Enterprise decisions involve PII, financial data, and health records. Sending sensitive data to a third-party LLM API is often not an option.
7. **Monitoring.** Decision volumes, rule firing frequencies, exception rates, and performance trends are operational requirements, not afterthoughts.
8. **Cost.** LLM inference at scale is expensive. Rule engine execution is cheap. The composite system's cost profile varies dramatically by pattern.

## Why LLMs alone fail

LLMs "show impressive results and some reasoning capabilities" yet "fail as easily when repeating the experience." This is the architecture, not a bug. LLMs are probability distributions over token sequences. For creative tasks, variation is a feature. For mortgage decisions, it is a liability. The same prompt run twice can produce different results. A slightly reworded prompt can produce a wrong result where the original produced a correct one.

The rule engine guarantees correctness relative to its encoded rules. The LRM generates text that *sounds like* reasoning. In enterprise contexts, you need the first and want the second — which is exactly why you combine them.

## Five patterns for composite AI

The 2023 article's core contribution is a practitioner's taxonomy: five integration patterns, each with defined tradeoffs. They come from experience with IBM ODM and ADS, but the patterns are general. They are not theoretical constructs. They are battle-tested integration topologies.

### Pattern 1: NLU → Rules

<img src="images/composite-ai-fig1-nlu-rules.svg" alt="Natural Language Understanding followed by Rule Reasoning" style="width:100%;max-width:720px;">

An LLM comprehends unstructured text and extracts structured data; a rule engine reasons deterministically on that data. An insurance claim — "I was rear-ended at Main and Oak last Tuesday. The other driver ran the red light. My bumper is damaged and my neck hurts." — becomes structured fields (`incident_type: rear_end_collision`, `fault_party: other_driver`, `damage_types: [property, bodily_injury]`), which the rule engine processes against policy terms, coverage rules, and state regulations.

**What works.** Integration is straightforward — sequential API calls passing a parameter context between engines. The LLM never makes a business decision; the rule engine never parses messy natural language. The separation of concerns is clean.

**What doesn't.** When expected data is missing from the text, the system needs guardrails. Does it reject the claim? Ask for clarification? Escalate to a human? These failure modes must be designed explicitly.

**The hidden work.** Schema design is where the abstraction leaks. If you ask the LLM to extract `damage_severity` as free text, the rule engine must parse "pretty bad," "minor," and "totaled" into actionable categories. If you constrain the LLM to a controlled vocabulary, the LLM must map ambiguous descriptions onto discrete labels — itself a classification task with error potential. Schema design at the structured/unstructured boundary is the engineering problem that determines whether Pattern 1 works in production.

### Pattern 2: Rules → NLG

<img src="images/composite-ai-fig2-rules-nlg.svg" alt="Rule Reasoning followed by Natural Language Generation with LLM" style="width:100%;max-width:720px;">

The flow reverses: a rule engine decides on structured data, then an LLM generates natural language output. A loan decision (`approved, $350,000, 6.25%, conditions: [income_verification, property_appraisal]`) becomes a customer letter.

This is the most common production pattern and the safest from a compliance perspective — the LLM never influences the decision, only communicates it. The real win is *consistency at scale with personalization*: every customer gets a correctly structured, deterministic communication whose surface text varies by tone, language, channel, and customer segment.

**What doesn't.** Testing NLG output is genuinely hard. The LLM may phrase the same information in dozens of valid ways. You cannot assert on output strings. You need semantic testing: does the output contain required decision fields? Does it omit restrictions, conditions, or deadlines? Does the tone match the channel?

A Spanish-language denial letter that subtly softens rejection language in a way that could be read as leaving the door open for appeal creates compliance exposure. Constrained generation — not free-form text from a decision payload, but generation within guardrails that prevent legally meaningful variation — is the requirement.

### Pattern 3: Rules orchestrate LLM

<img src="images/composite-ai-fig3-rules-orchestrate.svg" alt="Rule Reasoning Driving Natural Language Processing with LLM" style="width:100%;max-width:720px;">

The rule engine is the master orchestrator, invoking LLMs on demand for delegated NLP tasks. This extends the ML model calling capability that already exists in enterprise rule engines like IBM ODM and ADS — instead of calling a classification model for a score, the rule engine calls an LLM for text understanding or generation, but only when the reasoning path requires it.

Consider a complex claims adjudication: coverage is confirmed, liability is clear, but the claimed amount triggers a fraud review. The rule engine invokes an LLM: "Here is the claim description, the policy holder's history, and three years of claim notes. Summarize any inconsistencies or unusual patterns." The LLM returns structured analysis. The rule engine incorporates it into its decision path. The LLM is a tool the rule engine uses, not a peer.

**What works.** Costs are proportional to actual need — not every transaction invokes the LLM.

**What doesn't.** The rule engine needs orchestration rules: *when* to call the LLM, *what prompt* to send, *how to interpret* the response, *what fallback* to use if the response is malformed or nonsensical. These are not business rules but meta-rules governing the LLM interaction itself. A production implementation requires an orchestration framework that abstracts LLM invocation behind a clean interface — timeouts, retries, fallbacks, and guardrails against calling the LLM in a tight loop.

### Pattern 4: LLM extracts rules

<img src="images/composite-ai-fig4-llm-extract-rules.svg" alt="Extract Business Rules from Plain Text with an LLM, Run in a Logical Engine" style="width:100%;max-width:720px;">

The most ambitious pattern: LLMs at *design time* extract automation assets — business rules, data models, decision tables, rule signatures — from plain-text policy documents. The extracted assets generate an automation project in a decision platform. Rules then execute deterministically, completely decoupled from the LLM that helped author them. The article notes this was prototyped with IBM ADS.

Policy documents, regulations, and procedure manuals are ground truth. Today, human analysts manually encode them as decision tables — slow, expensive, error-prone, and creating a synchronization problem when source documents change. An LLM that extracts rules with traceability (each extracted rule links back to its source paragraph) promises dramatically decreased total cost of ownership.

**What doesn't.** This requires prompt chains or fine-tuned models, companion tools for validation and synchronization, and a human review gate. The LLM can turn weeks of manual rule writing into hours of review — but cannot replace expert judgment about whether a rule correctly captures policy intent. The lifecycle problem — keeping extracted rules in sync with evolving source documents — is the hard part that persists long after initial extraction.

### Pattern 5: Chatbot delegates to rules

<img src="images/composite-ai-fig5-chatbot-rules.svg" alt="Rules to Bring Reliable Reasoning in a Chatbot" style="width:100%;max-width:720px;">

An LLM drives the full conversational experience — understanding user intents, managing dialogue state, gathering context across turns. When a business decision is needed, the chatbot delegates to a rule-based decision engine. It recognizes the decision trigger, gathers required parameters from the conversation, invokes the rule engine, and restitutes results through NLG. This is Pattern 1 and Pattern 2 stitched together in a conversational loop.

The article references IBM Watson Orchestrate and a LangChain + ODM integration as examples.

**What works.** The chatbot benefits from conversational UX while delegating corporate reasoning to deterministic engines. This is the pattern most end-users actually interact with.

**What doesn't.** Two hard problems. **Decision trigger detection:** in open-ended conversation, when has the user crossed from "what are your rates?" to "I'd like to apply"? False positives annoy; false negatives lose business. **Incomplete context:** "my income is around 80K" when the engine needs an exact figure. The chatbot must recognize the ambiguity and ask, not fabricate.

The delegation boundary needs a formal contract: the rule engine exposes a decision service with a defined input schema, and the chatbot populates that schema conversationally. This is a form-filling dialogue where the form is the DMN model's input requirements.

## FSMs with rule engines: process skeleton, decision muscle

The five Feillet patterns describe how to combine LLMs with rule engines. But there is an orthogonal architectural dimension that predates LLMs entirely: combining finite state machines with rule engines. This pattern — FSM as process skeleton, rule engine as decision muscle — remains one of the most underappreciated choices in enterprise systems.

### The pattern

An FSM manages the process lifecycle: which stage the transaction is in, which transitions are valid, what happens on entry and exit to each state. At each state where a decision is required, the FSM delegates to a rule engine. The rule engine receives the accumulated facts and produces a structured decision that determines the next transition.

Consider a mortgage origination pipeline:

```
[Application] → [Review] → [Underwriting] → [Approval] → [Closing]
                    |            |               |
                    v            v               v
               Rule Engine   Rule Engine    Rule Engine
               (completeness (credit risk,   (final conditions,
                check)        collateral)     compliance)
```

The FSM owns the pipeline. It knows the application moves from Review to Underwriting only after completeness checks pass. It knows Underwriting can transition to Approval (decision: approve), Rejection (decision: deny), or back to Review (decision: need more information). The rule engines own the decisions at each gate.

### Why this works

**Process changes do not touch policy.** Adding a new "Fraud Check" state between Review and Underwriting requires modifying the FSM transition graph. None of the rule sets change.

**Policy changes do not touch process.** Changing the debt-to-income threshold from 43% to 45% is a one-line rule change. The FSM never knows about it.

**Testing is independently tractable.** The FSM can be tested with mocked rule engine responses. The rule engine can be tested with canned fact sets. The combinatorial explosion of testing both together is avoided.

**Auditability is layered.** A regulator asks: "Why was this loan denied?" The audit trail shows process trace (Application → Review → Underwriting, valid transitions) and decision trace (rules R-17, R-23, R-41 fired; output: `deny, DTI exceeds threshold, insufficient collateral`) as separate, composable artifacts.

### Where LLMs fit in

The FSM + rule engine pattern composes naturally with Feillet's five patterns. At any state in the FSM:

- **Pattern 1 (NLU → Rules):** The FSM is in the Intake state. Unstructured text arrives. An LLM extracts structured fields. The rule engine decides on completeness. The FSM transitions.
- **Pattern 2 (Rules → NLG):** The FSM reaches the Notification state. The rule engine has produced a decision. An LLM generates the customer communication. The FSM transitions to Closed.
- **Pattern 3 (Rules orchestrate LLM):** The FSM is in the Investigation state. The rule engine determines additional analysis is needed, invokes an LLM for anomaly detection, incorporates the result, and drives the next transition.
- **Pattern 5 (Chatbot + Rules):** The entire FSM is wrapped in a conversational interface. The chatbot tracks which FSM state the conversation is in, gathers parameters, and communicates decisions back to the user.

### The architectural principle

The FSM manages *where you are in the process*. The rule engine manages *what you know and what you should conclude*. The LLM manages *how you communicate at the boundaries*. Three concerns. Three tools. One system.

This principle is old — it predates LLMs, predates RETE, goes back to the early days of business process management and expert systems. But it is worth restating because the current AI discourse tends to collapse everything into "ask the LLM." A mortgage origination system is not one decision. It is a sequence of decisions embedded in a process. The FSM provides the sequence. The rule engines provide the decisions. The LLM provides the natural language interface at the boundaries.

## Choosing a pattern

The five patterns compose. A real system might use Pattern 4 to extract rules during development, Pattern 1 for intake, Pattern 2 for communications, and Pattern 5 for the conversational interface — with Pattern 3 orchestrating complex processes where LLM calls are needed selectively.

| Boundary | Pattern | Open-source path |
|---|---|---|
| Unstructured → structured | 1: NLU → Rules | RuleGo + LLM extraction node |
| Structured → unstructured | 2: Rules → NLG | RuleGo chain → LLM NLG |
| Complex NLP orchestration | 3: Rules drive LLM | RuleGo DAG with LLM + MCP nodes |
| Policy → code | 4: LLM extracts rules | LLM → RuleGo chain JSON |
| Conversational decisions | 5: Chatbot + Rules | RuleGo MCP server + LLM chatbot |

The next essay examines the KU Leuven research program that provides the academic foundation for these patterns — deep learning pipelines that extract DMN decision models from text, intelligent assistants that explain decisions, and frameworks that generate chatbots directly from decision models.

---

**References**

1. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

2. Pierre Feillet. [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). Medium, June 4, 2026.

3. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0. Includes [rulego-components-ai](https://github.com/rulego/rulego-components-ai) for LLM integration and MCP support.

> *Part 1: [On Rule Engines: From RETE to MCP](#on-rule-engines-rete-to-mcp) · Part 3: [On Rule Engines: Automating Decision Models](#on-rule-engines-automating-decision-models)*
