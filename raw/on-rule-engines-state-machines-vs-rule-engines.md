---
title: On Rule Engines — State Machines vs. Rule Engines
date: 2026-07-15
slug: on-rule-engines-state-machines-vs-rule-engines
summary: A deep dive into the architectural distinction between finite state machines and rule engines — when to use each, anti-patterns that arise from using the wrong one, and a decision framework for choosing. With code.
tags: llm, rules, fsm, state-machines, architecture, enterprise, design
series: on-rule-engines
---

The first essay in this series introduced the distinction between finite state machines and rule engines. This essay — the fourth in a five-part series — develops that distinction in depth. Expect code, anti-patterns, and a framework for choosing.

> Every system architect should be able to answer one question cold: is this a job for an FSM or a rule engine?

## Two questions, two tools

A **finite state machine** answers *what happens next?* It is procedural. You define states. You define transitions between them. The machine is always in exactly one state. When an event arrives, the machine consults its transition table and moves — or stays put.

A **rule engine** answers *what should we conclude?* It is declarative. You define condition-action pairs. The engine evaluates all rules against a working memory of facts. Any rule whose conditions are satisfied is eligible to fire. The engine — not the author — determines evaluation order through its conflict resolution strategy.

```go
// FSM: you define what happens next
type ClaimFSM struct {
    state string
    transitions map[string]map[string]string
}

func NewClaimFSM() *ClaimFSM {
    return &ClaimFSM{
        state: "intake",
        transitions: map[string]map[string]string{
            "intake":       {"submit": "review", "withdraw": "closed"},
            "review":       {"approve": "payment", "reject": "closed", "need_info": "intake"},
            "payment":      {"complete": "closed"},
        },
    }
}

// Rule engine: you define what should be concluded
rules := []Rule{
    {When: "claim.amount > 10000 AND claim.type = 'injury'", Then: Action("flag", "senior_review")},
    {When: "claimant.prior_claims > 3 AND claim.type = 'property'", Then: Action("flag", "fraud_check")},
    {When: "policy.active = false", Then: Action("deny", "coverage_lapsed")},
}
```

> The FSM asks: given where I am and what just happened, where do I go? The rule engine asks: given everything I know, what should I conclude?

## The distinction, in detail

| | FSM | Rule Engine |
|---|---|---|
| **Core question** | What happens next? | What should we conclude? |
| **Paradigm** | Procedural | Declarative |
| **State** | Explicit states + transition table | Working memory of facts |
| **Control flow** | Defined by transition graph | Defined by rule firing (inference + agenda) |
| **Author thinks about** | States, events, transitions | Conditions, actions, conflict resolution |
| **Best for** | Workflows, protocols, pipelines | Policies, decisions, classifications |
| **Complexity driver** | States × transitions | Rules × fact combinations |
| **Determinism** | Deterministic given state + event | Deterministic given rule set + facts |
| **Auditability** | Trace: state sequence + transitions taken | Trace: fired rules + matched facts |

The distinction is not academic. Using the wrong tool produces systems that work but are impossible to maintain.

## Anti-pattern 1: Encoding policy as an FSM

Consider a loan eligibility policy: credit score bands, debt-to-income thresholds, collateral requirements, regulatory jurisdictions. Encode this as an FSM:

```go
// ANTI-PATTERN: Policy encoded as states
// Each combination of conditions becomes a state. Explosion ensues.
type LoanPolicyFSM struct {
    creditScore int
    dti         float64
    state       string
}

func (f *LoanPolicyFSM) evaluate() string {
    // What should be 3 rules becomes a combinatorial state explosion
    switch {
    case f.creditScore < 620:
        return "deny"
    case f.creditScore >= 620 && f.creditScore <= 699 && f.dti < 0.43:
        return "approve_standard"
    case f.creditScore >= 700 && f.dti < 0.36:
        return "approve_preferred"
    // Add collateral rules? Multiply states by collateral types.
    // Add jurisdiction rules? Multiply by 50 states × federal regs.
    // The state space is the Cartesian product of all condition dimensions.
    }
    return "manual_review"
}
```

> Two conditions produce 3 outcomes. Three conditions produce 9. Five conditions with regulatory variation produce hundreds. The FSM approach scales exponentially with condition dimensions.

The fix: policy belongs in a rule engine.

```go
// CORRECT: Policy as rules — each dimension adds rules, not states
rules := []Rule{
    {When: "credit_score < 620", Then: Action("deny")},
    {When: "credit_score >= 620 AND credit_score <= 699 AND dti < 0.43", Then: Action("approve", "standard")},
    {When: "credit_score >= 700 AND dti < 0.36", Then: Action("approve", "preferred")},
    {When: "collateral.value < loan.amount * 0.8 AND loan.type = 'unsecured'",
     Then: Action("require", "additional_collateral")},
    // Adding a new condition dimension adds one rule, not a Cartesian product
    {When: "jurisdiction = 'CA' AND loan.amount > 500000",
     Then: Action("require", "california_disclosure")},
}
engine := NewRuleEngine(rules)
engine.Insert(loanFacts)
decision := engine.Run()
```

> Policy dimensions add rules linearly. States multiply transitions exponentially. That is the entire argument for using the right tool.

## Anti-pattern 2: Encoding workflow as flat rules

Now the inverse. Consider a claims processing pipeline: intake → review → investigation → payment → close. Encode this as flat rules:

```python
# ANTI-PATTERN: Workflow smuggled through working memory facts
rules = [
    Rule("step = 'intake' AND form.complete = true", actions=[
        Action("validate", "form"),
        Action("set", "step", "review"),       # process state as a fact!
    ]),
    Rule("step = 'review' AND claim.amount < 1000", actions=[
        Action("set", "step", "payment"),       # implicit transition
    ]),
    Rule("step = 'review' AND claim.amount >= 1000", actions=[
        Action("set", "step", "investigation"), # implicit transition
    ]),
    Rule("step = 'investigation' AND fraud_check.complete = true", actions=[
        Action("set", "step", "payment"),       # implicit transition
        Action("set", "fraud_flag", fraud_check.result),
    ]),
    Rule("step = 'payment' AND payment.processed = true", actions=[
        Action("set", "step", "closed"),        # implicit transition
    ]),
]
```

What is wrong with this? Everything.

1. **The process state is smuggled.** `step` is a working memory fact, not a first-class state. Nothing guarantees only one `step` fact exists. Nothing prevents contradictory transitions.
2. **Transitions are implicit.** The author intended intake → review → payment → closed. But the rules don't express this as a graph. You cannot look at the rule set and see the process.
3. **Adding a state requires discipline.** To add a "fraud_check" state between investigation and payment, you must update every rule that references `step = 'investigation'` and `step = 'payment'`. Miss one and you have a bug that only manifests when a specific combination of facts triggers the stale transition.
4. **Testing is combinatorial.** Each rule depends on `step` plus domain facts. To test the payment transition, you must set up facts that satisfy both the step condition *and* all other conditions in the rule.

> The process is there, but it is encoded indirectly through fact manipulation. You can read the rule set and not see the state machine. That is the definition of implicit.

The fix: workflow belongs in an FSM.

```python
# CORRECT: Process as explicit states and transitions
class ClaimProcess:
    states = ["intake", "review", "investigation", "payment", "closed"]
    transitions = {
        "intake":       {"validated": "review", "incomplete": "intake"},
        "review":       {"low_value": "payment", "needs_review": "investigation"},
        "investigation": {"cleared": "payment", "flagged": "closed"},
        "payment":      {"processed": "closed"},
    }

    def __init__(self):
        self.state = "intake"

    def handle(self, event):
        if event in self.transitions[self.state]:
            self.state = self.transitions[self.state][event]
            return self.state
        raise InvalidTransition(self.state, event)
```

> If you find yourself writing `step = 'something'` in every rule condition, you are encoding a state machine in a rule engine. Stop. Use an FSM.

## The decision framework

When choosing between an FSM and a rule engine, ask three questions:

**1. Is the problem primarily about sequence or about conditions?**

Sequence-driven → FSM. Condition-driven → rule engine.

A claims process is about sequence: intake, then review, then investigation, then payment. A loan eligibility policy is about conditions: credit score, DTI, collateral value.

```go
// Sequence-driven: the order of states IS the business logic
// Process: intake → review → underwriting → approval → closing
// You cannot skip review. You cannot go back from closing.
fsm := NewFSM() // correct choice

// Condition-driven: the combination of facts IS the business logic
// Policy: if credit > 700 AND DTI < 36% AND LTV < 80% → preferred rate
// The order of evaluating credit, DTI, and LTV does not matter.
engine := NewRuleEngine() // correct choice
```

**2. Does the complexity grow with states or with conditions?**

If adding a new business rule means adding a state → you are in the wrong tool. If adding a new process step means updating multiple rules → you are in the wrong tool.

```go
// Complexity smell: adding a condition dimension explodes states
// Adding "jurisdiction" to a loan FSM: 50 states × existing states = explosion
// Adding "jurisdiction" to a rule engine: one rule, maybe a decision table

// Complexity smell: adding a process step touches many rules
// Adding "fraud_check" to flat rules: update every rule with step guards
// Adding "fraud_check" to an FSM: add one state, two transitions
```

**3. Will an auditor need to trace the process or explain the decision?**

Process trace → FSM. Decision explanation → rule engine. Both → compose them (Part 5).

```go
// FSM audit: state sequence
// "Application reached underwriting via: intake → review → underwriting"
fsm.AuditTrail() // [intake, review, underwriting]

// Rule engine audit: rule trace
// "Decision: deny. Rules fired: R-17 (credit < 620), R-23 (DTI > 43%)"
engine.Explanation() // [{rule: R-17, reason: "credit_score=590 < 620"}, ...]
```

## When the problem is both

Most real business problems are both procedural and decisional. A mortgage origination system has a process (application → review → underwriting → approval → closing) and decisions at each gate (eligibility rules, pricing rules, compliance rules).

The mistake is picking one tool and forcing the other concern into it. The solution — developed in Part 5 — is composition: FSM as process skeleton, rule engine as decision muscle.

```
[Application] → [Review] → [Underwriting] → [Approval] → [Closing]
                    |            |               |
                    v            v               v
               Rule Engine   Rule Engine    Rule Engine
               (completeness (credit risk,   (final conditions,
                check)        collateral)     compliance)
```

> The FSM manages where you are. The rule engine manages what you conclude. Neither is asked to do the other's job. The composite handles both.

---

**References**

1. Charles L. Forgy. [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0). Artificial Intelligence, 19(1): 17–37, 1982.

2. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0.

> *Part 1: [On Rule Engines — From RETE to MCP](#on-rule-engines-rete-to-mcp) · Part 5: [On Rule Engines — State Machines Powered by Rule Engines](#on-rule-engines-state-machines-powered-by-rule-engines)*
