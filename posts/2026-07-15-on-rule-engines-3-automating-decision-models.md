---
title: On Rule Engines — Automating Decision Models
date: 2026-07-15
slug: on-rule-engines-automating-decision-models
summary: The KU Leuven research program on extracting DMN decision models from text using deep learning and LLMs, generating chatbots from decision models, and building explainable assistants — the academic foundation for Pattern 4 (LLM extracts rules) and Pattern 5 (chatbot delegation).
tags: llm, rules, enterprise, automation, composite-ai, decision-making, symbolic-ai, dmn, research
series: on-rule-engines
---

The first essay in this series traced the 40-year arc from Forgy's RETE algorithm to the Model Context Protocol. The second examined the five architectural patterns for blending LLMs with rule engines. This third essay turns to the academic research that validates, extends, and operationalizes those patterns.

While Pierre Feillet was developing composite AI patterns from enterprise practice at IBM, a sustained research program at KU Leuven's LIRIS (Leuven Institute for Research on Information Systems), led by Professor Jan Vanthienen and driven primarily by Alexandre Goossens and Vedavyas Etikala, was systematically attacking the same problem from the academic side. Their central question: *can we automate the extraction of decision models from text, and if so, how?*

The six papers that span from 2021 to 2023 form a coherent research arc: survey the landscape, prove feasibility with deep learning, scale to full DMN extraction with open tools, modernize with GPT-3, add explainability, and generate conversational interfaces. Together they provide the academic backbone for Pattern 4 (LLM extracts rules) and Pattern 5 (chatbot delegation), and they connect to every other pattern in Feillet's taxonomy.

## Mapping the landscape (KSEM 2021)

The research program began with a survey. Etikala and Vanthienen's [*An Overview of Methods for Acquiring and Generating Decision Models*](https://doi.org/10.1007/978-3-030-82153-1_17) (KSEM 2021, LNCS 12817, pp. 200–208) provided a taxonomy of techniques for acquiring decision models from various knowledge sources.

The paper noted that since the introduction of the **Decision Model and Notation (DMN) standard**, there had been a surge of research interest in automated extraction — but the field lacked a systematic map. Business decisions are of significant value to organizations, but manually modeling them is costly, tedious, and time-consuming. Their survey, with 38 references, classified approaches along three dimensions:

- **Source type:** natural language text, legacy code, other models, event logs
- **Extraction target:** decision dependencies (the structural relationships between decisions), decision logic (the actual rules and conditions), full models (both)
- **Technique family:** rule-based NLP, traditional machine learning, deep learning

At the time of writing, deep learning approaches were largely unexplored for DMN extraction specifically. The survey threw down the gauntlet: can deep learning extract decision models from text?

## First extraction results (BPM 2021)

Goossens, Claessens, Parthoens, and Vanthienen took the first step in [*Extracting Decision Dependencies and Decision Logic from Text Using Deep Learning Techniques*](https://doi.org/10.1007/978-3-030-94343-1_27) (BPM 2021 Workshops, LNBIP 436, pp. 349–361).

This was the first systematic attempt to apply deep learning specifically to DMN extraction. The authors collected a large dataset of labeled and tagged sentences from real use cases and trained two architectures — **BERT** and **Bi-LSTM-CRF** — for two distinct tasks:

1. **Sentence classification:** identifying which sentences in a document describe decision logic or decision dependencies
2. **Dependency extraction:** extracting the actual structural relationships between decisions from those sentences

The results demonstrated that deep learning could achieve sufficiently high performance to support (semi)-automatic extraction from text. This was proof of concept: the problem was tractable, the architectures worked, and the "semi" qualifier — human review remains essential — was established from the start.

A preliminary version appeared at RuleML+RR 2021 as [*Deep Learning for the Identification of Decision Modelling Components from Text*](https://doi.org/10.1007/978-3-030-91167-6_11) (LNCS 12851, pp. 158–171).

## Full DMN extraction (Expert Systems with Applications, 2023)

The definitive study came with Goossens, De Smedt, and Vanthienen's [*Extracting Decision Model and Notation Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667) (Expert Systems with Applications, Vol. 211, Article 118667, 2023).

This paper represents the state of the art for deep-learning-based DMN extraction. The authors investigated and evaluated two components:

- An **automatic sentence classifier** — identifying which sentences describe decision logic or dependencies
- A **decision dependency extractor** — extracting the structural relationships between decisions

Both were trained on a large, purpose-built labeled dataset collected from real use cases, using BERT and Bi-LSTM-CRF architectures.

Five contributions mark this as the landmark paper in the field:

1. **First investigation** of deep learning specifically for extracting DMN models from text
2. Demonstrated that deep learning can **classify sentences** describing logic or dependencies with high accuracy
3. Demonstrated that deep learning can **extract decision dependencies** from sentences — the structural backbone of a DMN model
4. **First labeled and tagged dataset** made publicly available for decision model extraction research
5. **First decision tool extraction from text** made available as an open tool

The conclusion — that BERT enables (semi)-automatic extraction of decision models from text — validates Pattern 4 (LLM extracts rules) with academic rigor. It also surfaces the "semi" qualifier: extraction is not fully automatic, human review remains essential, and the question shifts from "can we extract?" to "how do we build the validation tooling around the extraction pipeline?"

## GPT-3 enters the picture (RuleML+RR 2023)

The next logical step: what happens when you replace fine-tuned BERT with a general-purpose large language model?

Goossens, Vandevelde, Vanthienen, and Vennekens addressed this in [*GPT-3 for Decision Logic Modeling*](https://ceur-ws.org/Vol-3485/paper3896.pdf) (RuleML+RR 2023 Companion, CEUR Workshop Proceedings Vol. 3485, pp. 1–14). Where the earlier work used task-specific fine-tuned models, this challenge paper explored whether prompt engineering with GPT-3 could extract decision tables and rule logic directly from text — no fine-tuning dataset needed.

This shift is significant for practitioners. Fine-tuning requires a labeled dataset, which is expensive to produce and domain-specific. Prompt engineering requires crafting effective prompts, which is cheaper to iterate but potentially less accurate on structured extraction tasks. The paper sits at the intersection of Pattern 4 (LLM extracts rules) and the broader question of whether general-purpose models can match or exceed task-specific models for structured knowledge extraction.

A companion presentation by Vanthienen and Goossens at DecisionCamp 2023, [*GPT-3 for Decision Requirements Modeling and Advice*](https://decisioncamp2023.wordpress.com/), extended this exploration to decision requirements modeling.

## Explainable assistants (BPM 2022)

Extracting a decision model is half the problem. Making it useful to humans is the other half.

Goossens, Maes, Timmermans, and Vanthienen's [*Automated Intelligent Assistance with Explainable Decision Models in Knowledge-Intensive Processes*](https://doi.org/10.1007/978-3-031-25383-6_3) (BPM 2022 Workshops, LNBIP 460, pp. 25–36) asks: once you have a DMN model — whether hand-authored or automatically extracted — how do you make it accessible to stakeholders who need to understand a decision?

The authors propose a **generic intelligent assistant** that can reason with *any* DMN model to provide explanations of decisions. This is Pattern 2 (Rules → NLG) approached from the opposite direction: not generating customer communications from a structured decision, but generating explanations of the decision *process itself* — which rules fired, what triggered them, and why the outcome was reached.

A preliminary experiment compared two explanation sources — plain text and the intelligent assistant — to evaluate the assistant's capabilities. The early findings demonstrated feasibility for equipping organizations with explainable decisions embedded in their business processes.

This connects directly to Feillet's emphasis on transparency and auditability. An extracted or authored DMN model, when paired with an explanation-capable assistant, satisfies the regulatory requirement to show *why* a decision was made — not in post-hoc rationalization but in a traceable chain from facts through rules to conclusions.

## Chatbots from decision models (RuleML+RR 2021)

The final piece of the puzzle closes the loop between decision models and conversational interfaces.

Etikala, Goossens, Van Veldhoven, and Vanthienen's [*Automatic Generation of Intelligent Chatbots from DMN Decision Models*](https://doi.org/10.1007/978-3-030-91167-6_10) (RuleML+RR 2021, LNCS 12851, pp. 142–157) addresses a limitation in decision support systems: the lack of reliable, user-friendly ways to present decision-making processes to end-users.

The solution: a framework for **automatically generating intelligent chatbots from DMN decision models**. Instead of manually building a conversational interface for each decision service (Feillet's Pattern 5), the chatbot is generated directly from the DMN model's structure:

- The model's **input data requirements** become the dialogue's information-gathering slots
- The model's **decision hierarchy** becomes the conversation flow
- The model's **output structure** determines what the chatbot communicates back

This solves the two hard problems Feillet identified for Pattern 5. The DMN model's input schema provides the formal contract for the delegation handshake — a missing required field is not an ambiguous conversational state, it is a slot in the DMN input schema that has not been filled, and the chatbot knows it needs to ask for it. The model's decision hierarchy drives the conversation structure — the chatbot knows which information to gather in which order, and when it has enough to invoke the decision service.

## The research arc

Taken together, the six papers form a coherent progression:

| Stage | Paper | Contribution |
|---|---|---|
| **Survey** | Etikala & Vanthienen (KSEM 2021) | Taxonomy of acquisition methods across sources, targets, and techniques |
| **Feasibility** | Goossens et al. (BPM 2021) | First deep learning extraction of DMN components with BERT and Bi-LSTM-CRF |
| **Scale** | Goossens, De Smedt, Vanthienen (ESWA 2023) | Full DMN extraction, open dataset, open extraction tool |
| **Modernize** | Goossens et al. (RuleML+RR 2023) | GPT-3 for decision logic modeling via prompt engineering |
| **Explain** | Goossens et al. (BPM 2022) | Generic intelligent assistant for explainable decisions from any DMN model |
| **Converse** | Etikala et al. (RuleML+RR 2021) | Automatic chatbot generation from DMN decision models |

Feillet's five patterns describe *what* to build. The KU Leuven papers describe *how* to build key components: the extraction pipeline for Pattern 4, the explanation layer for Pattern 2, and the conversational interface for Pattern 5.

## RuleGo: an open-source implementation path

The patterns and research are not confined to academic papers and enterprise platforms. **[RuleGo](https://github.com/rulego/rulego)** — an open-source rule engine in Go (Apache 2.0) — provides a concrete implementation path for each pattern:

| Pattern | RuleGo implementation |
|---|---|
| 1: NLU → Rules | RuleGo chain with LLM extraction node → rule chain processing |
| 2: Rules → NLG | RuleGo chain → LLM NLG node for customer communications |
| 3: Rules drive LLM | RuleGo DAG with conditional LLM nodes invoked on demand |
| 4: LLM extracts rules | LLM generates RuleGo chain JSON from policy documents |
| 5: Chatbot + Rules | RuleGo MCP server exposes decision services to LLM chatbot |

The **[rulego-components-ai](https://github.com/rulego/rulego-components-ai)** extension provides LLM integration and MCP server/client support, making RuleGo a practical bridge between the research literature and production systems. The same RuleGo instance that runs on a Raspberry Pi 2 with ~19 MB of memory can expose its rule chains as MCP tools, making them discoverable by an LLM-based reasoning model — a concrete implementation of the hybrid inference loop Feillet describes.

## Open questions

The Feillet articles, the KU Leuven research program, and tools like RuleGo together cover substantial ground. But several questions remain open.

**The extraction quality bar.** The KU Leuven papers demonstrate extraction works — but at what accuracy threshold does the economics flip? At 90% accuracy, a human must still review every rule. The savings come from turning a *writing* task into a *reviewing* task — faster, but the human is still in the loop for every rule. At what accuracy does the review model shift from "review every rule" to "review only low-confidence extractions" to "review only rules that conflict with existing rules" to "review a statistical sample"? This is the operational question that determines whether Pattern 4 pays off at scale.

**Rule lifecycle management.** When source documents change, extracted rules must change. Does the LLM re-extract from the amended text? Does it diff the old and new policy and propose targeted rule changes? Does it flag rules originally extracted from the amended section for human review? Governed policy evolution, versioned extraction, and conflict detection between old and new rules — this synchronization problem is where the next wave of research and engineering needs to go.

**Testing composite systems.** Neither the practitioner articles nor the academic papers develop a testing methodology for composite AI. How do you test a system where one component is deterministic and the other probabilistic? You need property-based testing: the system's output must satisfy certain invariants (decision traceability, mandatory disclosures, schema conformance) even though the surface text varies. This is an area where the industry needs more tooling and methodology.

**Vendor neutrality.** Feillet's patterns are general, but the implementations described assume IBM ODM, ADS, and watsonx.ai. A vendor-neutral reference architecture — specifying the abstract interfaces between the LLM and rule engine components — would make the patterns more broadly applicable. RuleGo demonstrates one open-source path, but the interfaces between the components are not yet standardized.

**The MCP interface standard.** MCP is still evolving. The interface between a reasoning model and a decision service — how does the LRM discover available decision services? What information does it pass? How does the rule engine communicate confidence, conditions, or partial results back? — needs standardization if composite AI is to move beyond bespoke integrations.

## The composite AI thesis

<img src="images/composite-ai-fig6-overview.svg" alt="Composite AI — Blending Neuronal and Symbolic Approaches" style="width:100%;max-width:720px;">

The central thesis running through this series is that the future of enterprise AI is composite. Not LLMs replacing rule engines. Not rule engines staying isolated. A deliberate blending where each technology does what it is good at.

LLMs handle the **perception** layer — understanding unstructured text, classifying intents, extracting entities, generating fluent natural language. They are statistical engines optimized for flexibility and fluency.

Rule engines, built on Forgy's insight from 1979, handle the **reasoning** layer — applying deterministic logic to structured data, producing auditable decisions, enforcing regulatory constraints, scaling to high transaction volumes with predictable performance. They are logical engines optimized for reliability and transparency.

The five patterns are five ways to draw the boundary between these layers. The FSM is the process skeleton that sequences decisions through states. The boundary shifts depending on the use case, the risk tolerance, and the maturity of the integration. The principle is constant: let the neuronal system handle the messiness of natural language, let the symbolic system handle the precision of business logic, and let the state machine handle the process that connects them.

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

> *Part 1: [On Rule Engines: From RETE to MCP](#on-rule-engines-rete-to-mcp) · Part 2: [On Rule Engines: Five Patterns for Composite AI](#on-rule-engines-five-patterns)*
