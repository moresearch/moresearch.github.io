---
title: On Rule Engines — From RETE to MCP
date: 2026-07-15
slug: on-rule-engines-rete-to-mcp
summary: Tracing a 40-year arc from Charles Forgy's RETE algorithm to the Model Context Protocol, this essay explores how rule engines and large reasoning models answer the same fundamental question through different mechanisms — and why the most important trend is not replacement but convergence.
tags: llm, rules, enterprise, automation, composite-ai, symbolic-ai, rete, mcp, fsm
series: on-rule-engines
---

Enterprise decision automation sits at an uncomfortable intersection. Large Language Models promise natural interaction and flexibility. Business decisions demand auditability, determinism, and compliance with regulations that do not negotiate. The question is not which technology wins. The question is how to combine them — and the answer has roots going back nearly half a century.

> The question is not which technology wins. The question is how to combine them.

Pierre Feillet addressed this in his 2023 article [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f) and his 2026 follow-up [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). The second article traces a 40-year arc from Charles Forgy's RETE algorithm to the Model Context Protocol and argues that the most important trend in enterprise AI is not replacement but convergence.

This essay — the first in a five-part series — begins where the story begins: with the RETE algorithm, the architecture of rule engines, the rise of large reasoning models, and the fundamental distinction between finite state machines and rule engines that every system architect should understand.

## The RETE algorithm (1979)

In 1979, Charles L. Forgy, a PhD student at Carnegie Mellon University, solved a problem that would shape the next four decades of enterprise computing. Given a large set of IF-THEN rules and a working memory full of facts, how do you efficiently determine which rules should fire?

The naive approach — test every rule against every fact on every cycle — is catastrophically slow. Ten thousand rules against a hundred thousand facts is a billion comparisons per cycle. Production systems in the 1970s ground to a halt under realistic workloads.

Forgy's insight was that working memory changes slowly between cycles. Typically only a few facts are added or removed at each step. Re-evaluating every rule from scratch means repeating vast amounts of work that has not changed. What if you could *remember* partial matches and only recompute what actually changed?

> The RETE algorithm trades memory for speed — and achieves performance theoretically independent of the number of rules.

The result was the RETE algorithm — Latin for "net" — published in Forgy's 1979 PhD thesis and in a landmark 1982 paper in *Artificial Intelligence*: [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0).

### How it works

RETE has two phases. **Compilation:** rules are compiled into a discrimination network — a directed acyclic graph where each node tests a condition. Conditions appearing in multiple rules are compiled once and shared. If twenty rules check `customer.tier == "premium"`, there is one node, not twenty.

Consider these three loan underwriting rules:

```python
# Rule 1: High-value loan requires collateral
if loan.amount > 500_000 and loan.type == "mortgage":
    loan.requires_collateral = True

# Rule 2: Premium customers get rate discount
if customer.tier == "premium" and loan.type == "mortgage":
    loan.rate_discount = 0.25

# Rule 3: High-value premium mortgage gets priority review
if loan.amount > 500_000 and customer.tier == "premium" and loan.type == "mortgage":
    loan.review_priority = "high"
```

A RETE network compiles these into shared condition nodes. The test `loan.type == "mortgage"` appears in all three rules but exists once in the network. The test `loan.amount > 500_000` appears in two rules but exists once. When a fact changes — say the loan amount updates — only the branches downstream of the `loan.amount` node re-evaluate. Rules 1 and 3 might be affected; Rule 2 is untouched.

**Runtime:** facts enter at the root and propagate through the network. Each node caches the facts — or partial matches, at join nodes — that satisfy its condition. When a fact changes, only affected branches re-evaluate. When a fact is removed, its matches are retracted from the caches. The algorithm trades memory for speed — storing partial match state at every node — and achieves performance that is, in Forgy's words, theoretically independent of the number of rules.

### The lineage

RETE became the core of **OPS5**, which powered **R1/XCON** — an expert system that configured VAX computer orders for Digital Equipment Corporation, one of the first commercially successful AI systems, reportedly saving DEC millions of dollars per year by catching configuration errors that human order processors missed.

It went on to become the backbone of production rule systems across the industry: **CLIPS** (NASA's C Language Integrated Production System), **Jess** (the Java Expert System Shell), **Drools** (open-source, now Apache, which evolved RETE into ReteOO and later PHREAK), **IBM Operational Decision Manager** (enterprise-grade decision automation with governance, versioning, and deployment pipelines), **Soar** (the cognitive architecture), **Blaze Advisor**, and **TIBCO BusinessEvents**.

> The RETE network was invisible infrastructure, humming inside systems that made consequential decisions about people's money, health, and legal status.

Banks used it for loan origination. Insurers for claims adjudication. Governments for eligibility determination. For four decades, RETE was the silent backbone of automated decision-making.

And then, around 2022, the world became captivated by a different kind of AI — one that produced fluent text rather than deterministic decisions. The question became: are rule engines obsolete?

Feillet's 2026 answer is unequivocal: no. They never died. They are running alongside your LLM right now.

### Beyond RETE

Modern rule engines have evolved beyond the discrimination network. **[RuleGo](https://github.com/rulego/rulego)** — an open-source rule engine in Go (Apache 2.0) — uses a Directed Acyclic Graph. Business logic is composed of component nodes wired into rule chains; messages flow along predetermined DAG paths rather than being matched against all rules. A RuleGo rule chain looks like this:

```go
// A RuleGo rule chain for loan application intake
ruleChain := rulego.NewRuleChain("loan-intake",
    // Node 1: Validate required fields
    rulego.NewTransformNode("validate-fields").
        WithScript(`msg.Metadata.valid = msg.Data.amount > 0 && msg.Data.applicantId != ""`),
    // Node 2: Branch on validation result
    rulego.NewSwitchNode("route").
        WithCase("valid", "extract-intent").
        WithCase("invalid", "return-error"),
    // Node 3: LLM intent extraction
    rulego.NewRestNode("extract-intent").
        WithEndpoint("https://api.llm.example/v1/chat").
        WithBodyTemplate(`Extract intent from: {{.Data.description}}`),
    // Node 4: Route to decision engine
    rulego.NewRestNode("decision-engine").
        WithEndpoint("http://odm.example/decisions/loan-eligibility"),
)
```

This trades RETE's expressive forward-chaining for deterministic execution with extremely low resource consumption — ~19 MB of memory under 500 concurrent requests on a Raspberry Pi 2. Both approaches have their place.

## Two ways to answer the same question

Feillet's 2026 article opens with a deceptively simple observation: both rule engines and Large Reasoning Models address the identical problem — *given what we know, what should we conclude?*

> The methods differ. The objective is the same. This means they can — and should — be compared, contrasted, and combined.

### How rule engines reason

Rule engines separate knowledge from execution. Business logic is declarative. Facts enter working memory. The inference engine evaluates which rules are satisfied. When multiple rules fire, a conflict resolution mechanism — the *agenda* — determines firing order.

```python
# Declarative: you say WHAT, not HOW
rules = [
    Rule(
        when=[Fact("loan.amount", ">", 500_000), Fact("loan.type", "==", "mortgage")],
        then=[Action("set", "loan.requires_collateral", True)]
    ),
    Rule(
        when=[Fact("customer.tier", "==", "premium"), Fact("loan.type", "==", "mortgage")],
        then=[Action("set", "loan.rate_discount", 0.25)]
    ),
]

# The engine handles: which rules to evaluate, in what order,
# what happens when multiple rules fire, how to resolve conflicts.
engine = RuleEngine(rules)
engine.insert(loan_facts)
engine.run()  # produces: {requires_collateral: True, rate_discount: 0.25}
```

Every inference step is explicit. An auditor can trace precisely which rule fired, what triggered it, and why a decision was reached. Rule engines do not generate post-hoc rationalizations. **They produce proofs.**

### How Large Reasoning Models reason

LRMs store knowledge not as explicit rules but as distributed numerical representations learned across billions of examples. Their reasoning is what Feillet calls "an approximation of modus ponens" executed through statistical pattern completion.

```
Prompt: "Should this $600K mortgage for a premium-tier customer require collateral?"

LLM: "Based on the loan amount of $600K exceeding the typical threshold of
$500K for mortgages, and considering this is a premium-tier customer..."
```

The power is flexibility — generalization across novel inputs, ambiguity handling, cross-domain reasoning without explicit programming. The fragility is hallucination — producing plausible-sounding but logically invalid conclusions.

> Plausibility and logical validity are different things. One yields proofs. The other yields rationales. Both are useful. Confusing them is dangerous.

### The core distinction

Feillet gives a concrete example from the 2023 article: a pizza ordering bot that "depending on the runs, provides the expected outcome or a surprising one." Wrong pizza toppings are annoying. Wrong loan decisions are lawsuits.

```python
# Rule engine: same input → same output, every time
result1 = engine.decide(loan_application)  # deny, DTI 47%
result2 = engine.decide(loan_application)  # deny, DTI 47%
assert result1 == result2  # always passes

# LLM: same input → statistically similar output, maybe not the same
response1 = llm.generate("Review this loan: " + application)  # "I recommend denial..."
response2 = llm.generate("Review this loan: " + application)  # "This application should be..."
# response1 and response2 may differ in wording, tone, or even conclusion
```

> The rule engine guarantees correctness relative to its encoded rules. The LRM generates text that sounds like reasoning. In enterprise contexts, you need the first and want the second — which is exactly why you combine them.

## The hybrid inference loop

Modern reasoning models have evolved beyond raw token prediction. When a calculation is needed, the model generates and executes code — an exact result fed back into context. This is a pattern: stochastic reasoning delegating to deterministic execution at moments where precision matters.

The Model Context Protocol formalizes this delegation. An MCP tool definition for a decision service looks like:

```json
{
  "name": "check_loan_eligibility",
  "description": "Evaluate a mortgage application against underwriting rules",
  "inputSchema": {
    "type": "object",
    "properties": {
      "loan_amount": {"type": "number"},
      "applicant_income": {"type": "number"},
      "credit_score": {"type": "integer"},
      "property_value": {"type": "number"},
      "loan_type": {"type": "string", "enum": ["mortgage", "heloc", "refinance"]}
    },
    "required": ["loan_amount", "applicant_income", "credit_score", "property_value"]
  }
}
```

The loop is:

```
1. Reason with available context
2. Identify what is uncertain or requires precision
3. Delegate to a tool (MCP call to the rule engine)
4. Receive deterministic facts
5. Continue reasoning with grounded information
```

Feillet draws a structural parallel to production rule systems invoking external actions — except orchestration is now performed by a neural model rather than a symbolic agenda. IBM ODM and Decision Intelligence expose their decision services via MCP, enabling a reasoning model to invoke a full rule engine at the precise point where governed, deterministic decisions are needed.

> The LRM handles interpretation and context. The rule engine handles logic requiring correctness and traceability. Neither is asked to do the other's job.

### Attention and RETE: same purpose, different mechanism

Feillet is careful not to equate the two, but he notes a shared architectural role. Both answer the question *what is relevant right now?*

RETE pre-compiles condition tests into a network and caches partial matches. Attention learns to compute relevance scores from data. The mechanisms differ. The function is identical.

### RuleGo: the convergence in a Go library

The convergence is not confined to enterprise platforms. **[RuleGo](https://github.com/rulego/rulego)** embodies the hybrid architecture in a Go library. An LLM call is just another node in the rule chain:

```go
// RuleGo rule chain: deterministic nodes + LLM nodes, same interface
chain := rulego.NewRuleChain("claims-fraud-check",
    // Deterministic: validate claim data
    rulego.NewTransformNode("validate").
        WithScript(`msg.Metadata.skip = msg.Data.claimAmount < 10000`),
    // Deterministic: branch
    rulego.NewSwitchNode("threshold-check").
        WithCase("skip", "approve").
        WithCase("review", "llm-analysis"),
    // LLM call: anomaly detection in claim notes
    rulego.NewRestNode("llm-analysis").
        WithEndpoint("https://api.llm.example/v1/chat").
        WithBodyTemplate(`Review for anomalies:
            Claim: {{.Data.description}}
            History: {{.Data.priorClaims}}
            Amount: {{.Data.claimAmount}}`),
    // Deterministic: route based on LLM result
    rulego.NewSwitchNode("fraud-route").
        WithCase("clean", "approve").
        WithCase("suspicious", "investigate"),
)
```

The **[rulego-components-ai](https://github.com/rulego/rulego-components-ai)** extension provides LLM integration and MCP server/client support. A RuleGo instance exposes its rule chains as MCP tools, discoverable by any reasoning model. The hybrid inference loop Feillet describes — reason, identify uncertainty, delegate, receive, continue — is directly implementable.

> The 40-year arc from RETE to MCP passes through `go get github.com/rulego/rulego`.

## Finite state machines vs. rule engines

Before moving to composite AI patterns — the subject of the next essay — there is an architectural distinction worth making explicit.

A **finite state machine** answers *what happens next?* It encodes explicit states and transitions. Its logic is procedural.

A **rule engine** answers *what should we conclude?* It encodes condition-action pairs against a working memory. Its logic is declarative.

```go
// FSM: procedural — you define the sequence
type LoanFSM struct { state string }
func (f *LoanFSM) Transition(event string) error {
    switch f.state {
    case "application":
        if event == "submit" { f.state = "review" }
    case "review":
        if event == "approve" { f.state = "underwriting" }
        if event == "reject" { f.state = "closed" }
    }
}

// Rule Engine: declarative — you define the conditions
rules := []Rule{
    {When: "loan.amount > 500K AND loan.type = mortgage", Then: "require_collateral"},
    {When: "customer.tier = premium AND loan.amount > 250K", Then: "priority_review"},
}
```

> If the problem is procedural — steps, stages, sequences — reach for an FSM. If the problem is decisional — eligibility, pricing, risk, compliance — reach for a rule engine. Most real business processes are both. That is the subject of the fourth and fifth essays.

The full comparison and anti-patterns are explored in Part 4 of this series. The composed architecture — FSM as process skeleton, rule engine as decision muscle — is developed in Part 5.

## The 40-year arc

<img src="images/composite-ai-fig6-overview.svg" alt="Composite AI — Blending Neuronal and Symbolic Approaches" style="width:100%;max-width:720px;">

The story from RETE to MCP is not obsolescence and replacement. It is infrastructure that works being augmented by capabilities that are new. The RETE network matches facts against rules with Boolean precision. The attention mechanism computes relevance over learned representations. Both answer the same question — *what is relevant right now?* — at different layers of the stack.

> A decision system that cannot explain itself is not enterprise-grade. A decision system that cannot handle ambiguity is not useful. The composite approach accepts both constraints and designs for them.

The next essay examines the five composite AI patterns that Feillet and his co-authors proposed in 2023: concrete architectures for combining neuronal and symbolic AI in production systems.

---

**References**

1. Charles L. Forgy. [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0). Artificial Intelligence, 19(1): 17–37, 1982.

2. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

3. Pierre Feillet. [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). Medium, June 4, 2026.

4. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0. Includes [rulego-components-ai](https://github.com/rulego/rulego-components-ai) for LLM integration and MCP support.

> *Part 2: [On Rule Engines — Five Patterns for Composite AI](#on-rule-engines-five-patterns) · Part 4: [On Rule Engines — State Machines vs. Rule Engines](#on-rule-engines-state-machines-vs-rule-engines)*
