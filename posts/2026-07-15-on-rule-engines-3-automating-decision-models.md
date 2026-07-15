---
title: On Rule Engines — Automating Decision Models
date: 2026-07-15
slug: on-rule-engines-automating-decision-models
summary: The KU Leuven research program on extracting DMN decision models from text using deep learning and LLMs, generating chatbots from decision models, and building explainable assistants — the academic foundation for Pattern 4 (LLM extracts rules) and Pattern 5 (chatbot delegation).
tags: llm, rules, enterprise, automation, composite-ai, decision-making, symbolic-ai, dmn, research
series: on-rule-engines
---

The first essay traced the 40-year arc from Forgy's RETE algorithm to MCP. The second examined the five architectural patterns for blending LLMs with rule engines. This third essay turns to the academic research that validates, extends, and operationalizes those patterns.

While Pierre Feillet was developing composite AI patterns from enterprise practice at IBM, a sustained research program at KU Leuven's LIRIS (Leuven Institute for Research on Information Systems), led by Professor Jan Vanthienen and driven primarily by Alexandre Goossens and Vedavyas Etikala, was systematically attacking the same problem from the academic side.

> Feillet's five patterns describe *what* to build. The KU Leuven papers describe *how* to build it.

## Mapping the landscape (KSEM 2021)

The research program began with a survey. Etikala and Vanthienen's [*An Overview of Methods for Acquiring and Generating Decision Models*](https://doi.org/10.1007/978-3-030-82153-1_17) (KSEM 2021) provided a taxonomy of techniques for acquiring decision models from various knowledge sources. Business decisions are of significant value — but manually modeling them is costly, tedious, and time-consuming.

The survey classified approaches along three dimensions: **source type** (text, legacy code, models, event logs), **extraction target** (dependencies, logic, full models), and **technique family** (rule-based NLP, traditional ML, deep learning). Deep learning approaches were largely unexplored for DMN extraction. The Gauntlet had been thrown.

## First extraction results (BPM 2021)

Goossens, Claessens, Parthoens, and Vanthienen took the first step in [*Extracting Decision Dependencies and Decision Logic from Text Using Deep Learning Techniques*](https://doi.org/10.1007/978-3-030-94343-1_27) (BPM 2021 Workshops). This was the first systematic attempt to apply deep learning specifically to DMN extraction.

The approach: collect a labeled dataset of sentences from real use cases, train two architectures — **BERT** and **Bi-LSTM-CRF** — for two tasks:

```python
# Task 1: Sentence classification — does this sentence describe decision logic?
sentences = [
    ("Applicants with credit score below 620 shall be denied.", "decision_logic"),
    ("The loan officer reviews the application package.", "process_description"),
    ("DTI ratio is calculated as total monthly debt / gross monthly income.", "definition"),
]
classifier = FineTunedBERT(sentences, labels=["decision_logic", "process_description", "definition"])

# Task 2: Dependency extraction — which decisions depend on which?
# Input: "The eligibility decision depends on the credit assessment and the income verification."
# Output: eligibility -> [credit_assessment, income_verification]
extractor = BiLSTMCRF(dependency_sentences)
```

The results demonstrated sufficiently high performance to support (semi)-automatic extraction. A preliminary version appeared at RuleML+RR 2021 as [*Deep Learning for the Identification of Decision Modelling Components from Text*](https://doi.org/10.1007/978-3-030-91167-6_11).

> The "semi" in semi-automatic was established from the start: extraction works, but human review remains essential.

## Full DMN extraction (Expert Systems with Applications, 2023)

The definitive study came with Goossens, De Smedt, and Vanthienen's [*Extracting Decision Model and Notation Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667) (Expert Systems with Applications, Vol. 211, 2023). Five contributions:

1. **First investigation** of deep learning specifically for extracting DMN models from text
2. Sentence **classification** for logic/dependency detection with high accuracy
3. **Dependency extraction** from sentences — the structural backbone of a DMN model
4. **First labeled dataset** made publicly available for decision model extraction research
5. **First extraction tool** made available as open source

The extracted model looks like:

```yaml
# DMN model extracted from policy text by BERT-based pipeline
decisions:
  - id: eligibility
    label: "Determine Loan Eligibility"
    dependencies: [credit_assessment, income_verification, collateral_check]
    logic:
      - when: "credit_score < 620"
        then: "deny"
      - when: "620 <= credit_score <= 699 AND dti_ratio < 43%"
        then: "approve_standard"
      - when: "credit_score >= 700 AND dti_ratio < 36%"
        then: "approve_preferred"

  - id: pricing
    label: "Determine Interest Rate"
    dependencies: [eligibility, market_conditions]
    logic:
      - when: "eligibility = approve_preferred AND ltv_ratio < 80%"
        then: "rate = base_rate - 0.5%"
      - when: "eligibility = approve_standard"
        then: "rate = base_rate + 0.25%"
```

> The leap from "can we extract?" to "here is the tool and the dataset" is what makes this paper the landmark in the field.

## GPT-3 enters the picture (RuleML+RR 2023)

Goossens, Vandevelde, Vanthienen, and Vennekens explored the next logical step in [*GPT-3 for Decision Logic Modeling*](https://ceur-ws.org/Vol-3485/paper3896.pdf) (RuleML+RR 2023 Companion). Replace fine-tuned BERT with prompt-engineered GPT-3:

```python
# Fine-tuned approach (BPM 2021, ESWA 2023):
# Requires labeled dataset, domain-specific training, high accuracy on known formats
model = FineTunedBERT.train(labeled_sentences, labels)
rules = model.extract(policy_text)

# Prompt-engineered approach (RuleML+RR 2023):
# No training data needed, general model, potentially lower structured accuracy
rules = llm.extract(policy_text, prompt="""
    Extract decision rules from the following policy text.
    Output as a decision table in JSON format.
    Each rule must include: when (conditions), then (conclusion).
    Link each rule to its source paragraph.
""")
```

> The shift trades training cost for prompt engineering cost. No fine-tuning dataset needed — but structured extraction accuracy may be lower.

A companion presentation by Vanthienen and Goossens at DecisionCamp 2023, [*GPT-3 for Decision Requirements Modeling and Advice*](https://decisioncamp2023.wordpress.com/), extended this to decision requirements modeling.

## Explainable assistants (BPM 2022)

Extracting a model is half the problem. Goossens, Maes, Timmermans, and Vanthienen's [*Automated Intelligent Assistance with Explainable Decision Models in Knowledge-Intensive Processes*](https://doi.org/10.1007/978-3-031-25383-6_3) (BPM 2022 Workshops) asks: once you have a DMN model, how do you make it accessible?

They propose a **generic intelligent assistant** that can reason with *any* DMN model to provide explanations:

```python
class DecisionAssistant:
    """Generic assistant: works with any DMN model."""
    def __init__(self, dmn_model: DMNModel):
        self.model = dmn_model

    def explain(self, decision_id: str, inputs: dict) -> Explanation:
        """Explain why a decision reached its conclusion."""
        trace = self.model.execute(decision_id, inputs)
        return Explanation(
            decision=trace.outcome,
            fired_rules=[step.rule for step in trace.steps],
            input_facts=trace.facts_used,
            reasoning_chain=[
                f"Rule {step.rule.id} fired because {step.rule.condition} matched {step.matched_facts}"
                for step in trace.steps
            ],
        )
```

> An extracted DMN model paired with an explanation-capable assistant satisfies the regulatory requirement to show *why* a decision was made — not in post-hoc rationalization but in a traceable chain from facts through rules to conclusions.

## Chatbots from decision models (RuleML+RR 2021)

Etikala, Goossens, Van Veldhoven, and Vanthienen close the loop in [*Automatic Generation of Intelligent Chatbots from DMN Decision Models*](https://doi.org/10.1007/978-3-030-91167-6_10) (RuleML+RR 2021). Their framework generates a chatbot directly from a DMN model's structure:

```python
# A DMN model becomes a conversational interface automatically
dmn = DMNModel.load("mortgage-eligibility.dmn")

chatbot = ChatbotGenerator(dmn).generate()
# Generated chatbot behavior:
#   Slot 1: "What is the loan amount?"        → dmn.inputs.loan_amount
#   Slot 2: "What is your annual income?"     → dmn.inputs.applicant_income
#   Slot 3: "What is your credit score?"       → dmn.inputs.credit_score
#   Slot 4: "What is the property value?"      → dmn.inputs.property_value
#   --- all required inputs gathered ---
#   Invoke: dmn.decide("eligibility", inputs)
#   Response: "Based on your credit score of 720 and DTI of 32%,
#              you qualify for preferred rates at 6.0%."

# The DMN schema IS the conversation contract
```

> A missing required field is not an ambiguous conversational state. It is a slot in the DMN input schema that has not been filled, and the chatbot knows it needs to ask for it.

This solves the two hard problems Feillet identified for Pattern 5: the DMN input schema provides the formal delegation contract, and missing required fields are unambiguously identifiable slots to ask about.

## The research arc

| Stage | Paper | Contribution |
|---|---|---|
| Survey | Etikala & Vanthienen (KSEM 2021) | Taxonomy of acquisition methods |
| Feasibility | Goossens et al. (BPM 2021) | First deep learning DMN extraction |
| Scale | Goossens, De Smedt, Vanthienen (ESWA 2023) | Full extraction, open dataset and tools |
| Modernize | Goossens et al. (RuleML+RR 2023) | GPT-3 for decision logic modeling |
| Explain | Goossens et al. (BPM 2022) | Explainable assistant from any DMN model |
| Converse | Etikala et al. (RuleML+RR 2021) | Chatbots from DMN models |

## RuleGo: an open-source implementation path

The patterns and research are not confined to academic papers and enterprise platforms. **[RuleGo](https://github.com/rulego/rulego)** provides a concrete implementation path:

```go
// Pattern 4 implemented with RuleGo: LLM generates RuleGo chain JSON from policy
policyText := readPolicy("underwriting-policy-2026.txt")
chainJSON := llm.Extract(policyText, PromptConfig{
    Format: "rulego_chain",
    Schema: rulego.ChainSchema,
})

// The generated chain runs deterministically, LLM-free at runtime
chain := rulego.LoadChain(chainJSON)
decision := chain.Run(loanApplication)

// Generated chain structure:
// {
//   "ruleChain": {
//     "nodes": [
//       {"id": "credit-check", "type": "switch",
//        "cases": [
//          {"when": "credit_score < 620", "then": "deny"},
//          {"when": "credit_score >= 620 && credit_score <= 699", "then": "dti-check"},
//          {"when": "credit_score >= 700", "then": "preferred-check"}
//        ]},
//       {"id": "dti-check", "type": "switch", ...},
//       {"id": "preferred-check", "type": "switch", ...}
//     ]
//   }
// }
```

## Open questions

**The extraction quality bar.** At what accuracy threshold does the economics flip? At 90%, a human reviews every rule. At 95%? At 99%? The savings come from turning a *writing* task into a *reviewing* task — but the threshold where you stop reviewing every rule is where the operational gains live.

**Rule lifecycle management.** When source documents change, extracted rules must change. Governed policy evolution, versioned extraction, conflict detection between old and new rules — this synchronization problem is where the next wave of research needs to go.

**Testing composite systems.** How do you test a system where one component is deterministic and the other probabilistic? Property-based testing: invariants that must hold regardless of surface variation.

**Vendor neutrality.** The patterns are general but the implementations assume IBM products. RuleGo demonstrates one open-source path, but the interfaces between components are not yet standardized.

**The MCP interface standard.** How does an LRM discover available decision services? What information passes between them? How are partial results and confidence signals communicated?

## The composite AI thesis

<img src="images/composite-ai-fig6-overview.svg" alt="Composite AI — Blending Neuronal and Symbolic Approaches" style="width:100%;max-width:720px;">

The central thesis running through this series: the future of enterprise AI is composite. LLMs handle the **perception** layer — unstructured text, intents, entities, fluency. Rule engines, built on Forgy's insight from 1979, handle the **reasoning** layer — deterministic logic, auditable decisions, regulatory compliance. The FSM handles the **process** layer — sequencing decisions through states and transitions.

> Let the neuronal system handle the messiness of natural language. Let the symbolic system handle the precision of business logic. Let the state machine handle the process that connects them.

A decision system that cannot explain itself is not enterprise-grade. A decision system that cannot handle ambiguity is not useful. The composite approach accepts both constraints and designs for them. That has been the engineering move since Forgy built the first discrimination network in 1979.

---

**References**

1. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

2. Pierre Feillet. [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). Medium, June 4, 2026.

3. Alexandre Goossens, Johannes De Smedt, Jan Vanthienen. [*Extracting Decision Model and Notation Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667). Expert Systems with Applications, 211: 118667, 2023.

4. Alexandre Goossens, Simon Vandevelde, Jan Vanthienen, Joost Vennekens. [*GPT-3 for Decision Logic Modeling*](https://ceur-ws.org/Vol-3485/paper3896.pdf). RuleML+RR Companion, CEUR Vol. 3485, 2023.

5. Alexandre Goossens, Ulysse Maes, Yves Timmermans, Jan Vanthienen. [*Automated Intelligent Assistance with Explainable Decision Models in Knowledge-Intensive Processes*](https://doi.org/10.1007/978-3-031-25383-6_3). BPM Workshops 2022, LNBIP 460, pp. 25–36.

6. Alexandre Goossens, Michelle Claessens, Charlotte Parthoens, Jan Vanthienen. [*Extracting Decision Dependencies and Decision Logic from Text Using Deep Learning Techniques*](https://doi.org/10.1007/978-3-030-94343-1_27). BPM Workshops 2021, LNBIP 436, pp. 349–361.

7. Vedavyas Etikala, Jan Vanthienen. [*An Overview of Methods for Acquiring and Generating Decision Models*](https://doi.org/10.1007/978-3-030-82153-1_17). KSEM 2021, LNCS 12817, pp. 200–208.

8. Vedavyas Etikala, Alexandre Goossens, Ziboud Van Veldhoven, Jan Vanthienen. [*Automatic Generation of Intelligent Chatbots from DMN Decision Models*](https://doi.org/10.1007/978-3-030-91167-6_10). RuleML+RR 2021, LNCS 12851, pp. 142–157.

9. Jan Vanthienen, Alexandre Goossens. [*GPT-3 for Decision Requirements Modeling and Advice*](https://decisioncamp2023.wordpress.com/). DecisionCamp 2023.

10. [RuleGo](https://github.com/rulego/rulego) — Lightweight, component-based rule engine for Go. Apache 2.0. Includes [rulego-components-ai](https://github.com/rulego/rulego-components-ai) for LLM integration and MCP support.

> *Part 1: [On Rule Engines — From RETE to MCP](#on-rule-engines-rete-to-mcp) · Part 2: [On Rule Engines — Five Patterns for Composite AI](#on-rule-engines-five-patterns) · Part 4: [On Rule Engines — State Machines vs. Rule Engines](#on-rule-engines-state-machines-vs-rule-engines)*
