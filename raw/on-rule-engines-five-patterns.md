---
title: On Rule Engines — Five Patterns for Composite AI
date: 2026-07-15
slug: on-rule-engines-five-patterns
summary: Five architectural patterns for blending LLMs with rule-based decision engines — NLU to rules, rules to NLG, rules orchestrating LLMs, LLM-driven rule extraction, and chatbot delegation. Each pattern answers a different question about where the LLM sits in the decision pipeline and what role it plays.
tags: llm, rules, enterprise, automation, composite-ai, decision-making, symbolic-ai, fsm
series: on-rule-engines
---

The first essay in this series traced the 40-year arc from Forgy's RETE algorithm to the Model Context Protocol and established the core distinction: rule engines produce proofs, large reasoning models produce rationales. Both are useful. Neither alone suffices for production AI.

This essay — the second in a five-part series — examines the five architectural patterns that Pierre Feillet, Allen Chan, Luigi Pichett, and Yazan Obeidi proposed in their 2023 article [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). They let an architect reason about *where* in the pipeline the LLM should sit and *what role* it should play.

> These are not theoretical constructs. They are battle-tested integration topologies from enterprise practice.

## What an enterprise decision requires

Feillet's article enumerates eight criteria. They explain why the LLM-only approach keeps hitting a wall in production.

1. **Accuracy.** A loan decision wrong 2% of the time is not 98% accurate — it is a regulatory finding and a lawsuit. Correctness means *every time*.
2. **Scalability.** Millions of claims per day cannot degrade when the rule base grows to tens of thousands of rules. RETE's performance-independence from rule count matters here.
3. **Adaptability.** Regulations and policies change. The system must accommodate new rules without a multi-month DevOps cycle.
4. **Latency.** Budgets vary — milliseconds for fraud, seconds for pre-approval, minutes for underwriting — but must be predictable. LLM latency is variable; rule engine latency is not.
5. **Auditability.** "The model's attention weights converged on that outcome" is not an acceptable answer to a regulator.
6. **Privacy.** Enterprise decisions involve PII, financial data, and health records. Sending sensitive data to a third-party LLM API is often not an option.
7. **Monitoring.** Decision volumes, rule firing frequencies, exception rates — operational requirements, not afterthoughts.
8. **Cost.** LLM inference at scale is expensive. Rule engine execution is cheap. The composite system's cost profile varies dramatically by pattern.

> Every one of these criteria pushes in the same direction: the LLM should not be the decision-maker.

## Why LLMs alone fail

LLMs "show impressive results and some reasoning capabilities" yet "fail as easily when repeating the experience." This is the architecture, not a bug. LLMs are probability distributions over token sequences. For creative tasks, variation is a feature. For mortgage decisions, it is a liability.

```python
# The fundamental problem in one line:
# An LLM called 3 times on the same loan application may give 3 different answers.
for i in range(3):
    decision = llm.generate(f"Decide on this loan: {application}")
    print(decision)  # "Approved" → "Denied" → "Approved with conditions"
```

> The rule engine guarantees correctness relative to its encoded rules. The LRM generates text that sounds like reasoning. In enterprise contexts, you need the first and want the second — which is exactly why you combine them.

## Five patterns for composite AI

### Pattern 1: NLU → Rules

<img src="images/composite-ai-fig1-nlu-rules.svg" alt="Natural Language Understanding followed by Rule Reasoning" style="width:100%;max-width:720px;">

An LLM comprehends unstructured text and extracts structured data; a rule engine reasons deterministically on that data.

```go
// Pattern 1: NLU extracts, rules decide
func intakeClaim(description string) (Decision, error) {
    // Step 1: LLM extracts structured fields from free text
    structured, err := llm.Extract(description, ClaimSchema{
        Fields: []string{"incident_type", "fault_party", "damage_types", "date"},
    })
    if err != nil {
        return Decision{}, fmt.Errorf("extraction failed: %w", err)
    }

    // Step 2: Rule engine decides on structured data
    decision, err := ruleEngine.Decide("claims-coverage", map[string]interface{}{
        "incident_type": structured.IncidentType,
        "fault_party":   structured.FaultParty,
        "damage_types":  structured.DamageTypes,
        "policy_terms":  loadPolicy(structured.PolicyID),
    })
    if err != nil {
        return Decision{}, fmt.Errorf("decision failed: %w", err)
    }

    return decision, nil
}
```

**What works.** Sequential API calls. The LLM never makes a business decision. The rule engine never parses messy language.

**What doesn't.** When expected data is missing, the system needs guardrails. Does it reject? Clarify? Escalate?

> The abstraction leaks at the schema boundary. Schema design is the hidden engineering work in Pattern 1.

### Pattern 2: Rules → NLG

<img src="images/composite-ai-fig2-rules-nlg.svg" alt="Rule Reasoning followed by Natural Language Generation with LLM" style="width:100%;max-width:720px;">

The flow reverses: a rule engine decides on structured data, then an LLM generates natural language.

```go
// Pattern 2: Rules decide, LLM communicates
func notifyCustomer(application LoanApplication) (string, error) {
    // Step 1: Rule engine produces a structured decision
    decision, err := ruleEngine.Decide("loan-underwriting", map[string]interface{}{
        "amount":        application.Amount,
        "credit_score":  application.CreditScore,
        "dti_ratio":     application.DTIRatio,
        "collateral":    application.CollateralValue,
    })
    if err != nil {
        return "", err
    }

    // Step 2: LLM generates the customer letter from the decision
    letter, err := llm.Generate(LetterTemplate, map[string]interface{}{
        "decision":     decision.Outcome,    // "approved"
        "amount":       decision.Amount,     // 350000
        "rate":         decision.Rate,       // 6.25
        "conditions":   decision.Conditions, // ["income_verification", "appraisal"]
        "tone":         application.Channel, // "email_formal"
        "language":     application.Locale,  // "es-MX"
    })

    return letter, err
}
```

> This is the safest pattern from a compliance perspective. The LLM never influences the decision — it only communicates it.

**What doesn't.** Testing NLG output is genuinely hard. The LLM may phrase the same decision in dozens of valid ways. You need semantic testing:

```go
func TestNotificationLetter(t *testing.T) {
    decision := Decision{Outcome: "denied", Reason: "DTI exceeds threshold"}
    letter := generateLetter(decision)

    // Can't assert on exact string — assert on invariants
    if !strings.Contains(letter, "denied") {
        t.Error("letter must communicate denial")
    }
    if !strings.Contains(letter, "debt-to-income") {
        t.Error("letter must mention the reason")
    }
    if strings.Contains(letter, "approved") {
        t.Error("letter must not imply approval")
    }
}
```

> Constrained generation — not free-form text — is the requirement. A Spanish-language denial letter that subtly softens rejection language creates compliance exposure.

### Pattern 3: Rules orchestrate LLM

<img src="images/composite-ai-fig3-rules-orchestrate.svg" alt="Rule Reasoning Driving Natural Language Processing with LLM" style="width:100%;max-width:720px;">

The rule engine is the master orchestrator, invoking LLMs on demand for delegated NLP tasks.

```go
// Pattern 3: Rules drive the process, LLM is a tool called on demand
func adjudicateClaim(claim Claim) (Decision, error) {
    engine := rulego.NewRuleChain("claims-adjudication",
        // Step 1: Validate coverage
        rulego.NewTransformNode("validate-coverage").
            WithScript(`msg.Metadata.covered = msg.Data.policyActive && msg.Data.claimType == "covered"`),
        // Step 2: If covered, check amount threshold
        rulego.NewSwitchNode("threshold-check").
            WithCase("below_10k", "auto-approve").
            WithCase("above_10k", "fraud-review"),
        // Step 3: Invoke LLM for fraud review — only when threshold exceeded
        rulego.NewRestNode("fraud-review").
            WithEndpoint("https://api.llm.example/v1/chat").
            WithBodyTemplate(`Analyze for anomalies:
                Claim: {{.Data.description}}
                History: {{.Data.priorClaims}}
                Amount: {{.Data.claimAmount}}
                Respond with JSON: {"risk": "low|medium|high", "findings": [...]}`),
        // Step 4: Route based on LLM result
        rulego.NewSwitchNode("fraud-route").
            WithCase("low_risk", "auto-approve").
            WithCase("medium_risk", "manual-review").
            WithCase("high_risk", "investigate"),
    )
    return engine.Run(claim)
}
```

**What works.** Costs are proportional to actual need — not every transaction invokes the LLM. The rule engine remains in control.

> The rule engine needs orchestration rules: when to call the LLM, what prompt to send, how to interpret the response, what fallback to use if the LLM returns nonsense.

**What doesn't.** The coupling is tight. The rule engine must mediate the structured-unstructured frontier — taking probabilistic LLM output and converting it into deterministic actions. These are not business rules. They are meta-rules governing the LLM interaction itself.

### Pattern 4: LLM extracts rules

<img src="images/composite-ai-fig4-llm-extract-rules.svg" alt="Extract Business Rules from Plain Text with an LLM, Run in a Logical Engine" style="width:100%;max-width:720px;">

The most ambitious pattern: LLMs at *design time* extract automation assets from plain-text policy documents.

```python
# Pattern 4: LLM reads policy, generates executable rules
policy_text = """
4.2.3 Loan Eligibility: Applicants with a credit score below 620
shall be denied. Applicants with credit score 620-699 and debt-to-income
ratio below 43% may be approved with standard rates. Applicants with
credit score 700+ and DTI below 36% qualify for preferred rates.
"""

rules = llm.extract(policy_text, format="decision_table")
# Output:
# [
#   { "when": "credit_score < 620", "then": "deny" },
#   { "when": "620 <= credit_score <= 699 AND dti < 43%", "then": "approve_standard" },
#   { "when": "credit_score >= 700 AND dti < 36%", "then": "approve_preferred" },
# ]

# Rules are validated, reviewed by humans, then deployed to the engine.
# The LLM is not in the runtime path. The rules execute deterministically.
for rule in rules:
    ruleEngine.addRule(rule, source=policy_text, paragraph="4.2.3")
```

> Pattern 4 moves the LLM from the runtime path — where its latency, cost, and non-determinism are liabilities — to the development path, where its ability to process large volumes of unstructured text is an asset.

**What doesn't.** Prompt chains or fine-tuned models, companion tools for validation and synchronization, and a human review gate. The LLM turns weeks of manual rule writing into hours of review — but cannot replace expert judgment. The lifecycle problem — keeping extracted rules in sync with evolving source documents — persists long after initial extraction. This is where the KU Leuven research program enters (Part 3).

### Pattern 5: Chatbot delegates to rules

<img src="images/composite-ai-fig5-chatbot-rules.svg" alt="Rules to Bring Reliable Reasoning in a Chatbot" style="width:100%;max-width:720px;">

An LLM drives the conversation; when a business decision is needed, the chatbot delegates to a rule engine.

```go
// Pattern 5: Chatbot handles conversation, delegates decisions
func handleMortgageChat(session ChatSession, userMessage string) (string, error) {
    // Step 1: LLM understands the user's intent and extracts parameters
    intent, params, err := llm.UnderstandIntent(userMessage, MortgageIntents{
        Intents: []string{"rate_inquiry", "apply", "check_status", "general_question"},
        Slots:   []string{"loan_amount", "income", "credit_score", "property_value"},
    })

    // Step 2: If this is a decision trigger, delegate to rule engine
    if intent == "apply" && allRequiredFieldsPresent(params) {
        decision, err := ruleEngine.Decide("mortgage-eligibility", params)
        if err != nil {
            return llm.Generate("Something went wrong. Let me connect you with a specialist.", nil)
        }

        // Step 3: LLM restitutes the deterministic decision in natural language
        return llm.Generate("mortgage-decision-response", map[string]interface{}{
            "decision":     decision,
            "missing_info": missingFields(params),
            "next_steps":   nextStepsForDecision(decision),
            "tone":         session.UserPreferences.Tone,
        })
    }

    // Not a decision trigger — LLM handles conversation freely
    return llm.Chat(session.Context, userMessage)
}
```

> The delegation boundary needs a formal contract: the rule engine exposes a decision service with a defined input schema, and the chatbot populates that schema conversationally. A missing required field is not an ambiguous conversational state — it is a slot in the schema that has not been filled.

**Two hard problems.** **Decision trigger detection** — when has the user crossed from browsing to deciding? **Incomplete context** — "my income is around 80K" when the engine needs an exact figure. The chatbot must ask, not fabricate.

## FSMs with rule engines: process skeleton, decision muscle

The five Feillet patterns describe how to combine LLMs with rule engines. But there is an orthogonal architectural dimension that predates LLMs entirely: combining finite state machines with rule engines.

```go
// The composed architecture in one structure
type MortgagePipeline struct {
    fsm    *StateMachine           // owns process state and transitions
    rules  map[string]*RuleEngine  // one rule engine per decision gate
}

func (p *MortgagePipeline) Process(app LoanApplication) error {
    for p.fsm.State != "closed" {
        state := p.fsm.State
        // Delegate to the rule engine for this state
        decision, err := p.rules[state].Decide(app.AccumulatedFacts())
        if err != nil {
            return err
        }
        // The decision determines the next transition
        p.fsm.Transition(decision.NextState)
    }
    return nil
}
```

> The FSM manages *where you are in the process*. The rule engine manages *what you know and what you should conclude*. The LLM manages *how you communicate at the boundaries*. Three concerns. Three tools. One system.

The FSM vs. rule engine distinction is explored in depth in Part 4. The composed architecture — with full code, testing strategies, and audit patterns — is developed in Part 5.

## Choosing a pattern

| Boundary | Pattern | Signal |
|---|---|---|
| Unstructured → structured | 1: NLU → Rules | LLM extracts, rules decide |
| Structured → unstructured | 2: Rules → NLG | Rules decide, LLM communicates |
| Complex NLP orchestration | 3: Rules drive LLM | Rules call LLM on demand |
| Policy → code | 4: LLM extracts rules | LLM at design time only |
| Conversational decisions | 5: Chatbot + Rules | LLM talks, rules decide |

> The patterns compose. A real system might use Pattern 4 to extract rules during development, Pattern 1 for intake, Pattern 2 for communications, and Pattern 5 for the conversational interface — with Pattern 3 orchestrating complex processes where LLM calls are needed selectively.

---

**References**

1. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

2. Pierre Feillet. [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). Medium, June 4, 2026.

3. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0. Includes [rulego-components-ai](https://github.com/rulego/rulego-components-ai) for LLM integration and MCP support.

> *Part 1: [On Rule Engines — From RETE to MCP](#on-rule-engines-rete-to-mcp) · Part 3: [On Rule Engines — Automating Decision Models](#on-rule-engines-automating-decision-models) · Part 5: [State Machines Powered by Rule Engines](#on-rule-engines-state-machines-powered-by-rule-engines)*
