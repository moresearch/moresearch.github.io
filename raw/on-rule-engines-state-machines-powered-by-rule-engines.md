---
title: On Rule Engines — State Machines Powered by Rule Engines
date: 2026-07-15
slug: on-rule-engines-state-machines-powered-by-rule-engines
summary: The composed architecture — FSM as process skeleton, rule engine as decision muscle. A complete mortgage origination pipeline with Go code, testing strategies, layered audit trails, and LLM composition at the boundaries.
tags: llm, rules, fsm, state-machines, architecture, enterprise, golang, testing, audit
series: on-rule-engines
---

Part 4 established the distinction: FSMs handle sequence, rule engines handle decisions. This essay — the final in a five-part series — develops the composed architecture in full. An FSM manages process state and transitions. At each state where a decision is required, the FSM delegates to a rule engine. The LLM handles natural language at the boundaries. Three concerns. Three tools. One system.

> The FSM knows where you are. The rule engine knows what to conclude. The LLM knows how to say it. Each does one thing. Together they do everything.

## The mortgage origination pipeline

A mortgage application moves through five states. At each state gate, a rule engine makes a decision. The decision determines the transition.

```
[Intake] ──→ [Review] ──→ [Underwriting] ──→ [Approval] ──→ [Closing]
   │              │              │                  │              │
   v              v              v                  v              v
Rule Engine   Rule Engine   Rule Engine        Rule Engine    Rule Engine
(completeness (document     (credit risk,       (final         (funding
 check)        verification) collateral, DTI)    conditions)    verification)
```

### The domain types

```go
type LoanApplication struct {
    ID            string
    ApplicantID   string
    Amount        float64
    PropertyValue float64
    CreditScore   int
    AnnualIncome  float64
    MonthlyDebt   float64
    LoanType      string // "conventional", "fha", "va", "jumbo"
    Documents     []Document
    State         string // mirrors FSM state for persistence
}

type Decision struct {
    Outcome   string            // "approve", "deny", "need_info", "escalate"
    NextState string            // the FSM transition target
    Data      map[string]any    // structured decision payload
    Rules     []RuleTrace       // which rules fired and why
    Timestamp time.Time
}

type RuleTrace struct {
    RuleID      string
    Condition   string
    MatchedFacts map[string]any
}
```

### The FSM

```go
type MortgageFSM struct {
    state       string
    transitions map[string]map[string]string
    onEnter     map[string]func(*LoanApplication) error
    onExit      map[string]func(*LoanApplication) error
}

func NewMortgageFSM() *MortgageFSM {
    return &MortgageFSM{
        state: "intake",
        transitions: map[string]map[string]string{
            "intake": {
                "complete":   "review",
                "incomplete": "intake",
                "withdraw":   "closed",
            },
            "review": {
                "verified":     "underwriting",
                "missing_docs": "review",
                "deny":         "closed",
            },
            "underwriting": {
                "approve":      "approval",
                "deny":         "closed",
                "need_info":    "review",
            },
            "approval": {
                "conditions_met": "closing",
                "conditions_failed": "underwriting",
                "withdraw": "closed",
            },
        },
        onEnter: map[string]func(*LoanApplication) error{
            "intake":       validateApplication,
            "underwriting": lockRate,
            "closing":      generateClosingDocs,
        },
    }
}

func (f *MortgageFSM) Transition(app *LoanApplication, event string) error {
    target, ok := f.transitions[f.state][event]
    if !ok {
        return fmt.Errorf("invalid transition: %s --%s--> ?", f.state, event)
    }

    if fn, ok := f.onExit[f.state]; ok {
        if err := fn(app); err != nil {
            return fmt.Errorf("onExit %s: %w", f.state, err)
        }
    }

    f.state = target
    app.State = target

    if fn, ok := f.onEnter[target]; ok {
        if err := fn(app); err != nil {
            return fmt.Errorf("onEnter %s: %w", target, err)
        }
    }

    return nil
}
```

> The FSM is a library concern — small, testable, and independent of both rule logic and LLM integration.

### The rule engines

Each decision gate has its own rule engine, scoped to the decision at that state:

```go
type MortgagePipeline struct {
    fsm   *MortgageFSM
    gates map[string]*RuleEngine // one engine per decision state
}

func NewMortgagePipeline() *MortgagePipeline {
    return &MortgagePipeline{
        fsm: NewMortgageFSM(),
        gates: map[string]*RuleEngine{
            "intake":       NewIntakeRules(),
            "review":       NewReviewRules(),
            "underwriting": NewUnderwritingRules(),
            "approval":     NewApprovalRules(),
            "closing":      NewClosingRules(),
        },
    }
}

func NewUnderwritingRules() *RuleEngine {
    return NewRuleEngine([]Rule{
        {
            ID:      "UW-01",
            When:    "credit_score < 620",
            Then:    Action{Outcome: "deny", NextState: "deny", Reason: "credit_score_minimum"},
        },
        {
            ID:      "UW-02",
            When:    "dti_ratio > 0.43 AND loan_type != 'va'",
            Then:    Action{Outcome: "deny", NextState: "deny", Reason: "dti_exceeds_threshold"},
        },
        {
            ID:      "UW-03",
            When:    "ltv_ratio > 0.95 AND loan_type = 'conventional'",
            Then:    Action{Outcome: "deny", NextState: "deny", Reason: "insufficient_equity"},
        },
        {
            ID:      "UW-04",
            When:    "credit_score >= 700 AND dti_ratio < 0.36 AND ltv_ratio < 0.80",
            Then:    Action{Outcome: "approve", NextState: "approve", Rate: "preferred", Adjustment: -0.25},
        },
        {
            ID:      "UW-05",
            When:    "credit_score >= 620 AND dti_ratio < 0.43 AND ltv_ratio < 0.95",
            Then:    Action{Outcome: "approve", NextState: "approve", Rate: "standard"},
        },
        {
            ID:      "UW-06",
            When:    "documents MISSING 'tax_returns' OR documents MISSING 'pay_stubs'",
            Then:    Action{Outcome: "need_info", NextState: "need_info", MissingDocs: []string{"tax_returns", "pay_stubs"}},
        },
    })
}
```

> Each gate's rule engine knows only its own domain. The intake engine checks completeness. The underwriting engine evaluates credit risk. The approval engine applies final conditions. Changes to one gate's rules never affect another gate.

### The pipeline

```go
func (p *MortgagePipeline) Process(app *LoanApplication) (*Decision, []AuditEntry, error) {
    var audit []AuditEntry

    for p.fsm.State() != "closed" && p.fsm.State() != "deny" {
        state := p.fsm.State()
        engine, ok := p.gates[state]
        if !ok {
            return nil, audit, fmt.Errorf("no rule engine for state: %s", state)
        }

        // Gather facts accumulated so far
        facts := app.AccumulatedFacts()

        // Delegate to the rule engine for this state gate
        decision, err := engine.Decide(facts)
        if err != nil {
            return nil, audit, fmt.Errorf("decision at %s: %w", state, err)
        }

        // Record audit entry
        audit = append(audit, AuditEntry{
            State:      state,
            Decision:   decision.Outcome,
            RulesFired: decision.Rules,
            Transition: decision.NextState,
            Timestamp:  decision.Timestamp,
        })

        // The rule engine's decision drives the FSM transition
        if err := p.fsm.Transition(app, decision.NextState); err != nil {
            return nil, audit, fmt.Errorf("transition from %s: %w", state, err)
        }

        // If this was a terminal decision, return it
        if decision.Outcome == "deny" || decision.Outcome == "approve" {
            return &decision, audit, nil
        }
    }

    return nil, audit, fmt.Errorf("pipeline exited without terminal decision")
}
```

> The FSM calls the rule engine. The rule engine returns a decision with a NextState. The FSM transitions. The loop continues until a terminal state. This is the entire architecture.

## Testing the composite

Each component tests independently. The composition tests with mocks.

### Testing the FSM in isolation

```go
func TestMortgageFSM_Transitions(t *testing.T) {
    fsm := NewMortgageFSM()

    // Happy path: intake → review → underwriting → approval → closing
    app := &LoanApplication{ID: "LOAN-001"}

    if err := fsm.Transition(app, "complete"); err != nil {
        t.Fatal(err)
    }
    if fsm.State() != "review" {
        t.Errorf("expected review, got %s", fsm.State())
    }

    // Invalid transition: can't jump from review to closing
    if err := fsm.Transition(app, "conditions_met"); err == nil {
        t.Error("expected error for invalid transition review → closing")
    }
}

func TestMortgageFSM_InvalidTransitions(t *testing.T) {
    fsm := NewMortgageFSM()
    app := &LoanApplication{ID: "LOAN-002"}

    // Cannot approve from intake
    if err := fsm.Transition(app, "approve"); err == nil {
        t.Error("intake --approve--> ? should be invalid")
    }
}
```

### Testing a rule engine in isolation

```go
func TestUnderwritingRules_CreditDenial(t *testing.T) {
    engine := NewUnderwritingRules()

    decision, err := engine.Decide(map[string]any{
        "credit_score": 590,
        "dti_ratio":    0.30,
        "ltv_ratio":    0.75,
        "loan_type":    "conventional",
    })

    if err != nil {
        t.Fatal(err)
    }
    if decision.Outcome != "deny" {
        t.Errorf("expected deny for credit_score=590, got %s", decision.Outcome)
    }
    if decision.NextState != "deny" {
        t.Errorf("expected next_state=deny, got %s", decision.NextState)
    }
}

func TestUnderwritingRules_PreferredRate(t *testing.T) {
    engine := NewUnderwritingRules()

    decision, err := engine.Decide(map[string]any{
        "credit_score": 720,
        "dti_ratio":    0.32,
        "ltv_ratio":    0.75,
        "loan_type":    "conventional",
    })

    if err != nil {
        t.Fatal(err)
    }
    if decision.Outcome != "approve" {
        t.Errorf("expected approve, got %s", decision.Outcome)
    }
    if decision.Rate != "preferred" {
        t.Errorf("expected preferred rate for 720 credit, got %s", decision.Rate)
    }
}
```

### Testing the pipeline end-to-end

```go
func TestPipeline_HappyPath(t *testing.T) {
    pipeline := NewMortgagePipeline()
    app := &LoanApplication{
        ID:            "LOAN-003",
        Amount:        350000,
        PropertyValue: 450000,
        CreditScore:   720,
        AnnualIncome:  120000,
        MonthlyDebt:   3200,
        LoanType:      "conventional",
        Documents:     []Document{
            {Type: "tax_returns", Status: "verified"},
            {Type: "pay_stubs", Status: "verified"},
        },
    }

    decision, audit, err := pipeline.Process(app)
    if err != nil {
        t.Fatal(err)
    }
    if decision.Outcome != "approve" {
        t.Errorf("expected approve, got %s", decision.Outcome)
    }

    // Verify the process trace
    expectedStates := []string{"intake", "review", "underwriting", "approval", "closing"}
    for i, entry := range audit {
        if entry.State != expectedStates[i] {
            t.Errorf("step %d: expected state %s, got %s", i, expectedStates[i], entry.State)
        }
    }
}

func TestPipeline_LowCredit(t *testing.T) {
    pipeline := NewMortgagePipeline()
    app := &LoanApplication{
        CreditScore: 580,
        // ... other fields ...
    }

    decision, audit, err := pipeline.Process(app)
    if err != nil {
        t.Fatal(err)
    }
    if decision.Outcome != "deny" {
        t.Errorf("expected deny for credit 580, got %s", decision.Outcome)
    }

    // Verify the audit trail shows which rule fired
    lastEntry := audit[len(audit)-1]
    if lastEntry.State != "underwriting" {
        t.Errorf("expected denial at underwriting gate, got %s", lastEntry.State)
    }
}
```

> Three layers of testing. FSM in isolation: does it transition correctly? Rule engine in isolation: does it decide correctly? Pipeline end-to-end: does the composition work? Each layer can be tested independently. The combinatorial explosion of testing all state × rule combinations is avoided by testing the FSM with mocked decisions and the rule engine with canned facts.

## The layered audit trail

A regulator asks: "Why was loan LOAN-004 denied?" The composite system produces a layered answer.

```go
type FullAuditReport struct {
    ApplicationID string
    ProcessTrace  []ProcessStep    // FSM layer: where were we and when?
    DecisionTrace []DecisionRecord // Rule engine layer: what did we conclude and why?
    LLMTrace      []LLMRecord      // LLM layer: what did the model generate?
}

func (p *MortgagePipeline) FullAudit(app *LoanApplication) (*FullAuditReport, error) {
    decision, audit, err := p.Process(app)
    if err != nil {
        return nil, err
    }

    report := &FullAuditReport{ApplicationID: app.ID}

    for _, entry := range audit {
        // Process layer: FSM trace
        report.ProcessTrace = append(report.ProcessTrace, ProcessStep{
            State:     entry.State,
            EnteredAt: entry.Timestamp,
            Event:     entry.Transition,
        })

        // Decision layer: rule trace
        for _, rule := range entry.RulesFired {
            report.DecisionTrace = append(report.DecisionTrace, DecisionRecord{
                State:        entry.State,
                RuleID:       rule.RuleID,
                Condition:    rule.Condition,
                MatchedFacts: rule.MatchedFacts,
            })
        }
    }

    return report, nil
}
```

The output is a traceable chain: process trace (which states, in which order), decision trace (which rules fired, against which facts), and — when LLMs are composed at the boundaries — generation trace (what was communicated, in which words).

```
LOAN-004 Audit Report
=====================
Process Trace:
  intake(09:15) --complete--> review
  review(09:22) --verified--> underwriting
  underwriting(09:23) --deny--> closed

Decision Trace:
  underwriting: UW-01 fired — credit_score=590 < threshold=620
  underwriting: UW-03 fired — ltv_ratio=0.97 > maximum=0.95

Conclusion: deny. Reasons: [credit_score_minimum, insufficient_equity]
```

> The process says where. The rules say why. The audit combines both without conflating them.

## Composing with LLMs at the boundaries

The FSM + rule engine architecture composes naturally with Feillet's five patterns from Part 2. The LLM handles natural language at intake and notification. The FSM and rule engines handle the process and decisions between them.

```go
type MortgageSystem struct {
    pipeline *MortgagePipeline
    llm      *LLMClient
}

func (s *MortgageSystem) HandleInquiry(session ChatSession, message string) (string, error) {
    // Pattern 1: NLU at intake — LLM extracts structured fields
    intent, params, err := s.llm.UnderstandIntent(message, MortgageSchema)
    if err != nil {
        return s.llm.Generate("clarify", nil)
    }

    if intent == "apply" {
        app := params.ToLoanApplication()

        // FSM + Rule Engine: deterministic core
        decision, audit, err := s.pipeline.Process(app)
        if err != nil {
            return s.llm.Generate("error", map[string]any{"error": err})
        }

        // Pattern 2: NLG at notification — LLM communicates the decision
        return s.llm.Generate("mortgage_decision", map[string]any{
            "decision": decision,
            "audit":    audit,
            "tone":     session.UserPreferences.Tone,
            "language": session.UserPreferences.Locale,
        })
    }

    // Non-decision conversation: LLM handles freely
    return s.llm.Chat(session.Context, message)
}
```

> The LLM handles the messiness of human language. The FSM handles the discipline of process state. The rule engine handles the precision of business logic. Each does one thing. Each is testable independently. The composition handles everything.

## The design principles

**1. Separate process from policy.** The FSM encodes the sequence of states. The rule engine encodes the conditions for decisions. Never smuggle process state through working memory facts. Never encode branching policy logic in transition guards.

**2. One rule engine per decision gate.** Each gate's rules are scoped to that decision. Changing underwriting rules never affects intake rules. This is the single-responsibility principle applied to decision automation.

**3. Test in layers.** FSM with mocked decisions. Rule engine with canned facts. Pipeline end-to-end for integration. The combinatorial explosion of testing all state × rule combinations is avoided by layer isolation.

**4. Audit in layers.** Process trace (FSM) + decision trace (rule engine) + communication trace (LLM) = complete auditability. Each layer is independently queryable. Combined, they tell the full story.

**5. The LLM is a boundary component.** It handles natural language at intake and notification. It never makes a business decision. It never manages process state. The deterministic core — FSM + rule engines — is LLM-free at runtime.

> The composite architecture is not a compromise. It is the recognition that no single tool handles all three concerns well. Use the FSM for process. Use the rule engine for policy. Use the LLM for language. The system that does all three is the system that survives.

---

**References**

1. Charles L. Forgy. [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0). Artificial Intelligence, 19(1): 17–37, 1982.

2. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

3. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0.

> *Part 1: [On Rule Engines — From RETE to MCP](#on-rule-engines-rete-to-mcp) · Part 4: [On Rule Engines — State Machines vs. Rule Engines](#on-rule-engines-state-machines-vs-rule-engines)*
