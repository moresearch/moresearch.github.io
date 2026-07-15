---
title: "Composite AI: From RETE to MCP — Blending LLMs and Rule Engines for Enterprise Automation"
date: 2026-07-15
slug: composite-ai-llm-rules-automation
summary: "LLMs bring fluency but lack determinism. Rule engines bring auditability but struggle with ambiguity. Tracing a 40-year arc from Forgy's RETE algorithm to the Model Context Protocol, this post explores five composite patterns for blending neuronal and symbolic AI, the KU Leuven research program on automated decision model extraction, and why rule engines never died — they just learned to talk to your LLM."
tags: llm, rules, enterprise, automation, composite-ai, decision-making, symbolic-ai, rete, dmn, mcp
---

Enterprise decision automation sits at an uncomfortable intersection. On one side, Large Language Models promise natural interaction, contextual understanding, and flexibility. On the other, business decisions demand auditability, determinism, and compliance with regulations that do not negotiate. The question is not which technology wins. The question is how to combine them — and the answer has roots that go back nearly half a century.

Pierre Feillet, Allen Chan, Luigi Pichett, and Yazan Obeidi addressed this directly in their 2023 article [*Approaches in Using Generative AI for Business Automation*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). They proposed five architectural patterns that blend LLMs with rule-based decision engines — what they call *composite AI patterns*. In 2026, Feillet followed up with [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3), tracing a 40-year arc from Charles Forgy's RETE algorithm to the Model Context Protocol. Meanwhile, a sustained research program at KU Leuven, led by Jan Vanthienen, Alexandre Goossens, and colleagues, has been systematically attacking the problem from the academic side — building deep learning pipelines to extract DMN decision models from text, generating chatbots from decision models, and mapping the entire landscape of automated decision model acquisition.

This post weaves these threads together. It begins where the story begins: with the RETE algorithm.

## The RETE algorithm: where it all began

In 1979, Charles L. Forgy, a PhD student at Carnegie Mellon University, solved a problem that would shape the next four decades of enterprise computing. The problem was deceptively simple: given a large set of IF-THEN rules and a working memory full of facts, how do you efficiently determine which rules should fire?

The naive approach — test every rule against every fact on every cycle — is catastrophically slow. If you have 10,000 rules and 100,000 facts, that is a billion comparisons per cycle. Production systems in the 1970s ground to a halt under realistic workloads.

Forgy's insight was to observe that working memory changes slowly between cycles. Typically, only a few facts are added, removed, or modified at each step. Re-evaluating every rule from scratch means repeating vast amounts of work that has not changed. What if you could *remember* partial matches from the previous cycle and only recompute what actually changed?

The result was the **RETE algorithm** — Latin for "net" — published in Forgy's 1979 PhD thesis and in a landmark 1982 paper in *Artificial Intelligence*: [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0) (Artif. Intell. 19(1): 17–37).

### How RETE works

The algorithm has two phases.

**Compilation.** Rules are compiled into a discrimination network — a directed acyclic graph of nodes. Each node represents a condition test (e.g., `status == "active"`, `amount > 1000`). Forgy defined four node types: root nodes (entry points for facts), one-input nodes (test a single condition), two-input nodes (join facts that satisfy different patterns), and terminal nodes (all patterns satisfied — fire the rule). Conditions that appear in multiple rules are compiled once and shared. If twenty rules all check `customer.tier == "premium"`, there is one node in the network for that test, not twenty.

**Runtime execution.** Facts enter at the root and propagate through the network. Each node caches the facts (or partial matches, at join nodes) that satisfy its condition. When a new fact arrives, only the branches affected by that fact are re-evaluated. When a fact is removed, its matches are retracted from the caches. The algorithm trades memory for speed — storing partial match state at every node — and achieves performance that is, in Forgy's words, "theoretically independent of the number of rules."

This was transformative. Forgy's RETE became the core engine of **OPS5**, the production system language used to build **R1/XCON** — an expert system that configured VAX computer orders for Digital Equipment Corporation. R1/XCON was one of the first commercially successful expert systems, reportedly saving DEC millions of dollars per year by catching configuration errors that human order processors missed.

### The RETE lineage

RETE went on to become the backbone of production rule systems across the industry:

- **CLIPS** — NASA's C Language Integrated Production System, widely used in government and aerospace
- **Jess** — the Java Expert System Shell, bringing RETE to the JVM
- **Drools** — the open-source Java rule engine, now part of the Apache ecosystem, which evolved RETE into ReteOO (Rete Object-Oriented) and later PHREAK
- **IBM Operational Decision Manager (ODM)** — enterprise-grade decision automation with RETE at its core, extended with governance, versioning, and deployment pipelines
- **Soar** — the cognitive architecture that used RETE as its matching engine
- **Blaze Advisor**, **TIBCO BusinessEvents**, **BizTalk Rules Engine** — commercial rule platforms in financial services and insurance

The algorithm's core efficiency property — memory for speed, performance nearly independent of rule count — is what made it possible to run decision systems with tens of thousands of rules against millions of facts in production. Banks used it for loan origination. Insurers used it for claims adjudication. Governments used it for eligibility determination. The RETE network was invisible infrastructure, humming away inside systems that made consequential decisions about people's money, health, and legal status.

And then, around 2022, the world became captivated by a different kind of AI — one that produced fluent text rather than deterministic decisions. The question became: are rule engines obsolete?

Feillet's 2026 answer is unequivocal: no. They never died. They are running alongside your LLM right now.

It is also worth noting that the RETE algorithm, for all its historical importance, is not the only way to build a rule engine. Modern rule engines have explored alternative architectures. **[RuleGo](https://github.com/rulego/rulego)** — an open-source, Apache 2.0-licensed rule engine written in Go — uses a **Directed Acyclic Graph (DAG)** rather than a discrimination network. In RuleGo, business logic is composed of reusable, configurable component nodes wired into rule chains. A message flows along predetermined DAG paths rather than being matched against all rules in a RETE network. This trades RETE's expressive forward-chaining (any rule can fire based on any fact change) for deterministic, predictable execution paths with extremely low resource consumption — RuleGo runs on a Raspberry Pi 2 with ~19 MB of memory under 500 concurrent requests. The DAG approach represents a different point in the design space: less inferential power, more operational predictability. Both RETE-based engines and DAG-based engines have their place, and both are finding their way into composite AI architectures.

## Rule engines never died: the 40-year arc from RETE to MCP

In June 2026, Feillet published a companion piece to his 2023 article: [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). It traces a 40-year arc from Forgy's RETE to the Model Context Protocol and argues that the most important trend in enterprise AI is not replacement but convergence.

### The question that never changed

Feillet opens with a simple observation: both rule engines and Large Reasoning Models address the same fundamental problem — deriving conclusions from available knowledge. The methods differ significantly. The objective — *given what we know, what should we conclude?* — is identical.

This is more than a philosophical observation. It means that the two technologies are not addressing different problems that happen to be adjacent. They are addressing the *same* problem through different mechanisms. And that means they can — and should — be compared, contrasted, and combined.

### How rule engines reason

Rule engines separate knowledge from execution. Business logic is expressed as declarative condition-action pairs (IF/THEN structures). Facts enter working memory. The inference engine continuously evaluates which rules are satisfied. When multiple rules fire simultaneously, a conflict resolution mechanism called the *agenda* determines firing order.

The key architectural property is that the rule author specifies *what* should happen; the engine handles *how*. This separation is the source of the rule engine's transparency. Every inference step is explicit. An auditor can trace precisely which rule fired, what triggered it, and why a decision was reached. Rule engines do not generate post-hoc rationalizations. They produce *proofs* — logical derivations that are mechanically verifiable.

### How Large Reasoning Models reason

LRMs store knowledge not as explicit rules but as distributed numerical representations learned across billions of training examples. When solving problems, they navigate a high-dimensional conceptual space rather than pattern-matching against a predicate database. Their reasoning traces may look like rule chaining, but as Feillet puts it, this reasoning is "an approximation of modus ponens" executed through statistical pattern completion rather than logical evaluation.

The power is extraordinary flexibility — generalization across novel inputs, ambiguity handling, cross-domain reasoning without explicit programming.

The fragility is hallucination — producing plausible-sounding but logically invalid conclusions because "plausibility and logical validity are different things."

The core distinction Feillet draws: **rule engines guarantee correctness relative to their rules; LRMs generate probabilistically plausible conclusions.** One yields proofs. The other yields rationales. Both are useful. Confusing them is dangerous.

### Escaping pure probabilism

Feillet notes an important evolution in modern reasoning models: the ability to delegate to deterministic execution. When a calculation is needed, the model can generate and execute Python code — an exact, deterministic result fed back into the reasoning context. This is not just an LLM feature. It is a pattern: "stochastic reasoning delegating to deterministic execution" at moments where precision matters.

This pattern — the neural model recognizing its own limits and calling a deterministic subsystem — is the conceptual bridge to MCP.

### MCP and the hybrid inference loop

The Model Context Protocol formalizes the delegation pattern:

1. Reason with available context
2. Identify what is uncertain or requires precision
3. Delegate to a tool
4. Receive facts
5. Continue reasoning

Feillet draws a structural parallel to how production rule systems invoke external actions and continue with newly acquired working memory — except the orchestration is now performed by a neural model rather than a symbolic agenda. IBM ODM and Decision Intelligence expose their decision services via MCP, enabling a reasoning model to invoke a full rule engine at the precise point in its inference trace where governed, auditable, deterministic decisions are needed. The LRM handles interpretation, planning, and contextual understanding; the rule engine handles logic requiring correctness, traceability, and compliance.

### Attention and RETE: same purpose, different mechanism

Feillet is careful not to equate the two mechanisms, but he notes a shared architectural role:

- **RETE:** symbolic pattern matching with Boolean conditions, exact matches, discrete sets of eligible rules
- **Attention:** differentiable relevance computation with continuous weights, weighted sums of representations
- **Commonality:** both determine which information is most relevant for the next inference step — one operating on facts and predicates, the other on learned representations

This is not a coincidence. Both architectures face the same fundamental challenge — given a large set of knowledge and a specific context, what is relevant *right now*? RETE solves it by pre-compiling condition tests into a network and caching partial matches. Attention solves it by learning to compute relevance scores from data. The mechanisms are different. The function is the same.

### Converging trends

Feillet identifies three converging trends:

1. **Rule engines incorporating generative AI** for authoring, explaining, and maintaining policies in natural language — this is Pattern 4 (LLM extracts rules) and Pattern 2 (Rules → NLG) from the 2023 article
2. **Reasoning models incorporating deterministic subsystems** — code execution, retrieval, decision services — to ground outputs in verifiable reality
3. **Hybrid inference architectures** where stochastic reasoning orchestrates deterministic reasoning, and neural inference delegates to symbolic inference at the moments that matter

The 40-year arc from RETE to MCP is, in Feillet's telling, a story of convergence. Neither approach alone suffices for production AI. The winning architecture combines symbolic precision with neural flexibility operating in concert.

### RuleGo: a modern Go-native example of the convergence

The convergence Feillet describes is not theoretical. You can see it happening in open-source projects today. **[RuleGo](https://github.com/rulego/rulego)** is a concrete example of a modern rule engine that already embodies the hybrid architecture pattern.

RuleGo is built around a component-based DAG architecture. Each node in the graph is a pluggable component — a filter, a transformer, a router, an HTTP push, a Kafka sink, a JavaScript evaluator. Rule chains are compositions of these components. Messages enter at the root and flow along deterministic paths through the DAG to terminal nodes.

What makes RuleGo relevant to the composite AI story is its **[rulego-components-ai](https://github.com/rulego/rulego-components-ai)** extension — a dedicated library of AI-scenario components that includes LLM integration. A rule chain can call an LLM to extract user intent from a message, then branch on the structured result: if the intent is `loan_inquiry`, route to a loan eligibility rule chain; if `claim_dispute`, route to a claims adjudication chain. The LLM is a component in the DAG, invoked when the rule chain reaches its node, just like any other component. This is Pattern 3 (Rules orchestrate LLM) implemented in a lightweight, embeddable Go library rather than an enterprise platform.

RuleGo also surfaces **MCP server and client support** — the same protocol Feillet identifies as the bridge between neural and symbolic reasoning. A RuleGo instance can expose its rule chains as MCP tools, making them discoverable and invocable by an LLM-based reasoning model. Conversely, RuleGo can act as an MCP client, calling out to external decision services or LLM endpoints when its rule chain logic requires it. The hybrid inference loop Feillet describes — reason, identify uncertainty, delegate to tool, receive facts, continue — is directly implementable as a RuleGo rule chain that alternates between deterministic component nodes and MCP-mediated LLM calls.

The significance of RuleGo in this story is not that it replaces IBM ODM or the KU Leuven DMN extraction pipeline. It is that the composite AI pattern has become general enough to appear in a Go library that runs on a Raspberry Pi. The convergence is not confined to enterprise platforms. It is happening at the library level, in the open-source ecosystem, in the language (Go) that powers much of modern cloud infrastructure. The 40-year arc from RETE to MCP passes through `go get github.com/rulego/rulego`.

## What an enterprise decision actually requires

Before discussing architectures, the 2023 article enumerates eight criteria that any enterprise decision system must satisfy. These explain *why* the LLM-only approach keeps hitting a wall in production.

**1. Accuracy and Reliability.** If a loan eligibility decision is wrong 2% of the time, that is not a 98% success rate — it is a regulatory finding, a customer lawsuit, and a compliance audit. Accuracy means correctness *every time*, not average-case performance.

**2. Scalability.** A decision engine processing millions of claims per day cannot slow down when the rule base grows from hundreds to tens of thousands of rules. This is where RETE's performance-independence from rule count matters — and where LLM inference costs at scale become prohibitive.

**3. Flexibility and Adaptability.** Regulations, policies, and markets change. The system must accommodate new rules, new data sources, and new decision logic without a multi-month DevOps cycle.

**4. Real-time Decision Making.** Latency budgets vary — milliseconds for fraud detection, seconds for loan pre-approval, minutes for complex underwriting — but the system must operate within known bounds. LLM inference latency is inherently variable. Rule engine latency is predictable.

**5. Transparency and Auditability.** When a regulator asks why a claim was denied, "the model's attention weights converged on that outcome" is not an acceptable answer. You need a traceable chain of reasoning: which rules fired, which data triggered them, and what the decision path was. This is the fundamental advantage of rule engines: they produce proofs, not rationales.

**6. Security and Data Privacy.** Enterprise decisions involve PII, financial data, and health records. Sending sensitive data to a third-party LLM API is often not an option.

**7. Monitoring and Reporting.** Decision volumes, rule firing frequencies, exception rates, performance trends — these are operational requirements, not afterthoughts.

**8. Cost-effectiveness.** LLM inference at scale is expensive. Rule engine execution is cheap. The cost profile of the composite system matters and varies dramatically depending on which pattern you choose.

## Why LLMs alone fail the enterprise test

The 2023 article is direct: LLMs "show impressive results and some reasoning capabilities" yet "fail as easily when repeating the experience." They cite the phenomenon of the *stochastic parrot* — an LLM generates text that is statistically plausible based on its training distribution, not logically sound based on the facts at hand. The same prompt run twice can produce different results. A slightly reworded prompt can produce a wrong result where the original produced a correct one.

This is not a bug. It is the architecture. For creative tasks, probabilistic variation is a feature. For deciding whether someone qualifies for a mortgage, it is a liability. The authors give a concrete example: a pizza ordering bot that "depending on the runs, provides the expected outcome or a surprising one." Wrong pizza toppings are annoying. Wrong loan decisions are measured in lawsuits and regulatory actions.

The 2026 article sharpens this argument. Feillet contrasts proofs versus rationales: the rule engine guarantees correctness relative to its encoded rules; the LRM generates a text that *sounds like* reasoning. These are different products. In enterprise contexts, you often need the first and want the second — which is exactly why you combine them.

## The composite materials metaphor

The 2023 article frames the core insight through an analogy: just as materials science invented composites that surpass any single element (carbon fiber + epoxy, reinforced concrete, alloyed metals), AI systems can combine neuronal approaches (LLMs) with symbolic approaches (rule engines) to achieve what neither can do alone.

Composite AI — sometimes called neuro-symbolic AI — has been discussed in research for decades. What Feillet et al. contribute is a practitioner's taxonomy: five concrete integration patterns, each with defined pros, cons, and integration complexity. They come from experience with IBM's decision automation products — ODM and ADS — combined with LLM platforms like watsonx.ai.

The taxonomy lets an architect reason about *where* in the pipeline the LLM should sit and *what role* it should play. The answer depends on what you need the system to do.

## Pattern 1: NLU followed by rule reasoning

![Figure 1: Natural Language Understanding followed by Rule Reasoning](images/composite-ai-fig1-nlu-rules.svg)

The first pattern is conceptually the simplest: an LLM first comprehends unstructured text and extracts structured data; a causal rule engine then reasons deterministically on that data.

Think of an insurance claims intake system. A customer submits a free-text description: "I was rear-ended at the intersection of Main and Oak last Tuesday. The other driver ran the red light. My bumper is damaged and my neck hurts." The LLM extracts structured fields: `incident_type: rear_end_collision`, `fault_party: other_driver`, `damage_types: [property, bodily_injury]`, `location: Main_St_and_Oak_St`, `date: 2026-07-08`. The rule engine then processes these structured fields against the policy terms, coverage rules, and state regulations to determine coverage applicability, deductible amounts, and next steps.

**What works.** Integration is straightforward — sequential API calls passing a parameter context between engines. The LLM never makes a business decision; the rule engine never tries to parse messy natural language.

**What doesn't.** When expected data is not found in the text — the customer did not mention the date, or described the damage ambiguously — the system needs guardrails. Does it reject the claim? Ask for clarification? Escalate? These failure modes must be designed explicitly.

**Expansion: the schema design problem.** The structured schema that the LLM targets is where the abstraction leaks. If you ask the LLM to extract `damage_severity` as a free-text field, you have punted the structure problem downstream — the rule engine now has to parse "pretty bad," "minor," and "totaled" into actionable categories. If you constrain the LLM to a controlled vocabulary (`minor`, `moderate`, `severe`, `total_loss`), you need the LLM to map ambiguous descriptions onto those categories. Schema design is the hidden engineering work in Pattern 1.

## Pattern 2: Rule reasoning followed by NLG

![Figure 2: Rule Reasoning followed by Natural Language Generation with LLM](images/composite-ai-fig2-rules-nlg.svg)

The second pattern reverses the flow: a rule engine makes a decision on structured data first, then an LLM generates natural language output from that decision.

This is the most common production pattern today. A loan origination system processes an application through its rule engine and arrives at a structured decision: `approved`, `amount: $350,000`, `rate: 6.25%`, `conditions: [income_verification, property_appraisal]`. The LLM then generates the customer-facing letter.

The real win is *consistency at scale with personalization*. Every customer gets a correctly structured communication whose underlying decision is deterministic and auditable, but the surface text can vary by tone, language, channel, and customer segment. The LLM never influences the decision — it only communicates it. This is the safest pattern from a compliance perspective.

**What doesn't.** Testing NLG output is genuinely hard. The LLM may phrase the same information in dozens of valid ways. You need semantic testing: does the output contain the required decision fields? Does it omit (rather than hallucinate) restrictions or deadlines? Testing NLG is closer to testing a creative product than a function.

**Expansion: the personalization vs. compliance tension.** The more you let the LLM vary the output — adjusting tone, localizing, adapting to channels — the harder it becomes to ensure every variant faithfully represents the underlying decision. If a Spanish-language denial letter subtly softens the rejection language in a way that could be read as leaving the door open for appeal, you have created a compliance exposure. The NLG pattern requires *constrained generation* — not free-form text, but generation within guardrails that prevent legally meaningful variation.

## Pattern 3: Rule reasoning driving NLP with LLM

![Figure 3: Rule Reasoning Driving Natural Language Processing with LLM](images/composite-ai-fig3-rules-orchestrate.svg)

Pattern 3 puts the rule engine in the driver's seat as the master orchestrator, invoking LLMs on demand for delegated NLP tasks. This extends the ML model calling capability that already exists in enterprise rule engines: instead of calling a classification model for a score, the rule engine calls an LLM for text understanding or generation — but only when the reasoning path requires it.

Consider a complex claims adjudication scenario. The rule engine processes the claim: coverage is confirmed, liability is clear, but the claimed amount triggers a fraud review. The rule engine invokes an LLM: "Here is the claim description, the policy holder's history, and three years of claim notes. Summarize any inconsistencies or unusual patterns." The LLM returns a structured analysis. The rule engine incorporates that analysis into its decision path. The LLM is a tool the rule engine uses, not a peer.

**What works.** The LLM is called on demand, only when needed. This keeps inference costs proportional to actual need rather than running every transaction through an LLM.

**What doesn't.** The coupling is tight. The rule engine needs to know *when* to call the LLM, *what prompt* to send, *how to interpret* the response, and *what to do* if the response is malformed or nonsensical. The rule engine must mediate the structured-unstructured data frontier — taking the LLM's probabilistic output and converting it into something deterministic. This requires *orchestration rules*: rules about when and how to use the LLM, timeouts, fallbacks, and guardrails against calling the LLM in a tight loop.

## Pattern 4: Extract business rules from plain text with an LLM

![Figure 4: Extract Business Rules from Plain Text with an LLM, Run in a Logical Engine](images/composite-ai-fig4-llm-extract-rules.svg)

Pattern 4 is the most ambitious. Instead of using LLMs at runtime, it uses them at *design time* to extract automation assets — business rules, data models, decision tables — directly from plain-text policy documents. The extracted assets generate an automation project. The rules then execute deterministically, completely decoupled from the LLM that helped author them.

The article notes that "the success of this approach has already been prototyped with ADS." This moves the LLM from the runtime path (where its latency, cost, and non-determinism are liabilities) to the development path (where its ability to process large volumes of unstructured text is an asset).

**What works.** Policy documents are the ground truth for many business decisions. Today, human analysts manually encode them as decision tables — slow, expensive, error-prone, and creating a synchronization problem when source documents change. An LLM that can extract rules from policy text with traceability (each extracted rule links back to the source paragraph) promises dramatically decreased TCO.

**What doesn't.** This requires "an efficient prompt chain or a fine-tuned model," plus companion tools for validation and synchronization. The LLM can accelerate the authoring process — turning weeks of manual rule writing into hours of review — but cannot replace expert judgment about whether a rule correctly captures policy intent. The tools for validation and maintenance are at least as important as the extraction model itself.

**Expansion: rule lifecycle management.** When paragraph 4.2.3 of a regulation is amended, the rules extracted from it need updating. Does the LLM re-extract from the amended text? Does it diff the old and new policy and propose targeted changes? Does it flag rules from the amended section for human review? Keeping extracted rules in sync with evolving source documents is the hard problem that persists long after the initial extraction. This is precisely the territory that the KU Leuven research program has been systematically exploring.

## Pattern 5: Rules bring reliable reasoning to a chatbot

![Figure 5: Rules to Bring Reliable Reasoning in a Chatbot](images/composite-ai-fig5-chatbot-rules.svg)

Pattern 5 addresses the conversational interface. An LLM drives the conversational experience. But when a business decision is needed, the chatbot delegates to a rule-based decision engine. The chatbot recognizes the decision trigger, gathers context from the conversation, invokes the rule engine, and restitutes results through NLG.

The article references IBM Watson Orchestrate and a LangChain + ODM integration. This pattern essentially stitches Pattern 1 and Pattern 2 together within a conversational loop.

**What works.** The chatbot feels intelligent. The decisions it conveys are reliable and auditable. The user gets the best of both.

**What doesn't.** Two hard problems. First, *decision trigger detection* — in an open-ended conversation, how does the chatbot know the user has crossed from casual inquiry into decision territory? False positives are annoying; false negatives are lost business. Second, *incomplete context handling* — the user says "my income is around 80K" when the decision engine needs an exact figure. The chatbot must recognize the ambiguity and ask for clarification, not fabricate a value.

**Expansion: the delegation handshake protocol.** The delegation boundary needs a formal contract: the rule engine exposes a decision service with a well-defined input schema, and the chatbot is responsible for populating that schema from the conversation. This is essentially a *form-filling dialogue* where the form is the decision engine's input schema. The chatbot's job is to detect which decision service the user needs and fill its input schema conversationally. This is also precisely the problem that Etikala, Goossens, Van Veldhoven, and Vanthienen addressed in their 2021 work on generating chatbots from DMN models — a bridge from Feillet's practitioner patterns to the academic literature.

## The KU Leuven research program: automating decision model acquisition

While Feillet and colleagues were developing composite AI patterns from an enterprise architecture perspective, a sustained research program at KU Leuven's LIRIS (Leuven Institute for Research on Information Systems), led by Professor Jan Vanthienen, was attacking the same problem from the academic side. The central question: *can we automate the extraction of decision models from text, and if so, how?*

This body of work — spanning from 2021 to 2023, primarily driven by Alexandre Goossens, Vedavyas Etikala, and their collaborators — forms the academic backbone of Pattern 4 (LLM extracts rules). It is worth examining as a coherent research program because it shows how the problem evolves from "can we do this at all?" to "can we do it well enough to deploy?"

### Mapping the landscape (KSEM 2021)

The research program began with a survey. Etikala and Vanthienen's [*An Overview of Methods for Acquiring and Generating Decision Models*](https://doi.org/10.1007/978-3-030-82153-1_17) (KSEM 2021) provided a taxonomy of techniques for acquiring decision models from various knowledge sources: natural language text, legacy code, other models, and event logs. The paper noted that since the introduction of the DMN (Decision Model and Notation) standard, there had been a surge of interest in automated extraction — but the field lacked a systematic map. Their survey, with 38 references, provided the classification framework that the subsequent papers would populate.

The key dimensions they identified matter for practitioners: source type (text vs. code vs. logs vs. models), extraction target (decision dependencies vs. decision logic vs. full models), and technique family (rule-based NLP, traditional ML, deep learning). At the time of writing, deep learning approaches were largely unexplored. The Gauntlet had been thrown.

### First extraction results (BPM 2021)

Goossens, Claessens, Parthoens, and Vanthienen took the first step in [*Extracting Decision Dependencies and Decision Logic from Text Using Deep Learning Techniques*](https://doi.org/10.1007/978-3-030-94343-1_27) (BPM 2021 Workshops). This was the first systematic attempt to apply deep learning specifically to DMN extraction. The authors used a large dataset of labeled and tagged sentences and trained two architectures — BERT and Bi-LSTM-CRF — for two distinct tasks: classifying sentences that describe decision dependencies, and extracting the actual dependency relations between decisions.

The results demonstrated that deep learning could achieve sufficiently high performance to support (semi-)automatic extraction from text. This was proof of concept: the problem was tractable. A preliminary version appeared at RuleML+RR 2021 as [*Deep Learning for the Identification of Decision Modelling Components from Text*](https://doi.org/10.1007/978-3-030-91167-6_11) (LNCS 12851, pp. 158–171).

### Scaling up: full DMN extraction (Expert Systems with Applications, 2023)

The definitive study came with Goossens, De Smedt, and Vanthienen's [*Extracting Decision Model and Notation Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667) (Expert Systems with Applications, Vol. 211, 2023).

This paper represents the state of the art for deep-learning-based DMN extraction. The authors investigated two components: an **automatic sentence classifier** (identifying which sentences in a document describe decision logic or dependencies) and a **decision dependency extractor** (extracting the structural relationships between decisions). They trained BERT and Bi-LSTM-CRF models on a large, purpose-built labeled dataset collected from real use cases.

Five contributions are worth highlighting:

1. **First investigation** of deep learning specifically for extracting DMN models from text
2. Demonstrated that deep learning can **classify sentences** describing logic or dependencies with high accuracy
3. Demonstrated that deep learning can **extract decision dependencies** from sentences — the structural backbone of a DMN model
4. **First labeled and tagged dataset** made available for decision model extraction research
5. **First decision tool extraction from text** made available as an open tool

The conclusion — that BERT can be used for (semi)-automatic extraction of decision models from text — validates the enterprise Pattern 4 with academic rigor. It also surfaces the "semi" qualifier: extraction is not fully automatic. Human review remains essential. The question shifts from "can we extract?" to "how do we build the validation tooling around the extraction pipeline?"

### GPT-3 enters the picture (RuleML+RR 2023)

The next logical step: what happens when you replace BERT with a large language model? Goossens, Vandevelde, Vanthienen, and Vennekens addressed this in [*GPT-3 for Decision Logic Modeling*](https://ceur-ws.org/Vol-3485/paper3896.pdf) (RuleML+RR 2023 Companion, CEUR Vol. 3485).

Where the earlier work used fine-tuned BERT models for specific extraction subtasks, this paper explored whether a general-purpose LLM (GPT-3) could perform decision logic modeling — extracting decision tables and rule logic directly from text through prompt engineering rather than fine-tuning. The shift from fine-tuned extractors to prompt-based extraction with a general model is significant: it trades training cost (no fine-tuning dataset needed) for prompt engineering cost and potentially lower accuracy on structured extraction tasks. The paper sits at the intersection of Pattern 4 (LLM extracts rules) and the broader question of whether general-purpose models can match or exceed task-specific models for structured knowledge extraction.

### From extracted models to intelligent assistants (BPM 2022)

Extracting a decision model is half the problem. The other half is making it useful. Goossens, Maes, Timmermans, and Vanthienen's [*Automated Intelligent Assistance with Explainable Decision Models in Knowledge-Intensive Processes*](https://doi.org/10.1007/978-3-031-25383-6_3) (BPM 2022 Workshops, LNBIP 460, pp. 25–36) asks: once you have a DMN model — whether hand-authored or automatically extracted — how do you make it accessible to end-users who need to understand a decision?

The authors propose a generic intelligent assistant that can reason with *any* DMN model to provide explanations of decisions. This is Pattern 2 (Rules → NLG) from the opposite direction: not generating customer communications from a structured decision, but generating explanations of the decision *process itself* for stakeholders who need to understand what happened and why.

A preliminary experiment compared two explanation sources — plain text and the intelligent assistant — to evaluate the assistant's capabilities. The early results demonstrated feasibility for equipping organizations with explainable decisions embedded in their business processes. This connects directly to Feillet's emphasis on transparency and auditability: an extracted or authored DMN model, when paired with an explanation-capable assistant, satisfies the regulatory requirement to show *why* a decision was made.

### Chatbots from decision models (RuleML+RR 2021)

The final piece of the puzzle is the conversational interface. Etikala, Goossens, Van Veldhoven, and Vanthienen's [*Automatic Generation of Intelligent Chatbots from DMN Decision Models*](https://doi.org/10.1007/978-3-030-91167-6_10) (RuleML+RR 2021, LNCS 12851, pp. 142–157) closes the loop.

The paper addresses a limitation in decision support systems: the lack of reliable, user-friendly ways to present decision-making processes to end-users. The solution: a framework for **automatically generating intelligent chatbots from DMN decision models**. Instead of manually building a conversational interface for each decision service (Pattern 5 from Feillet's taxonomy), the chatbot is generated directly from the DMN model's structure — its input data requirements become the dialogue's information-gathering slots, its decision hierarchy becomes the conversation flow.

This is the academic counterpart to Feillet's Pattern 5, and it solves the two hard problems he identified: (a) the DMN model's input schema provides the formal contract for the delegation handshake, and (b) the model's required fields drive the form-filling dialogue. A missing required field is not an ambiguous conversational state — it is a slot in the DMN input schema that has not been filled, and the chatbot knows to ask for it.

### The research program as a whole

Taken together, these six papers form a coherent arc:

| Stage | Paper | What it does |
|---|---|---|
| **Survey** | Etikala & Vanthienen (KSEM 2021) | Maps the landscape of decision model acquisition |
| **Prove feasibility** | Goossens et al. (BPM 2021) | First deep learning extraction of DMN components |
| **Scale up** | Goossens, De Smedt, Vanthienen (ESWA 2023) | Full DMN extraction with BERT, released dataset and tools |
| **Modernize** | Goossens et al. (RuleML+RR 2023) | GPT-3 for decision logic modeling via prompt engineering |
| **Explain** | Goossens et al. (BPM 2022) | Intelligent assistant for explainable decisions from DMN |
| **Converse** | Etikala et al. (RuleML+RR 2021) | Automatic chatbot generation from DMN models |

The research program validates, extends, and operationalizes the composite AI vision from the academic side. Feillet's five patterns describe *what* to build. The KU Leuven papers describe *how* to build key components: the extraction pipeline (Pattern 4), the explanation layer (Pattern 2), and the conversational interface (Pattern 5).

## The decision spectrum: choosing a pattern

The five patterns are not mutually exclusive. A real enterprise system might use Pattern 4 to extract rules from policy documents during development, Pattern 1 to process unstructured claim submissions, Pattern 2 to generate customer communications, and Pattern 5 to expose the whole thing through a conversational interface — with Pattern 3 orchestrating complex multi-step processes where LLM calls are needed selectively. The patterns compose.

| Boundary | Best Pattern | Academic foundation | Open-source example |
|---|---|---|---|
| Unstructured input, structured decision | Pattern 1 (NLU → Rules) | DMN as extraction target | RuleGo + LLM component for entity extraction |
| Structured decision, unstructured output | Pattern 2 (Rules → NLG) | Goossens et al. 2022 (explainable assistant) | RuleGo chain → LLM for NLG |
| Complex multi-step process with NLP needs | Pattern 3 (Rules orchestrate LLM) | Hybrid inference via MCP | RuleGo DAG with LLM nodes + MCP |
| Policy-to-code automation | Pattern 4 (LLM extracts rules) | Goossens et al. 2021, 2023 (DMN extraction) | LLM → RuleGo chain JSON generation |
| Conversational interface to decisions | Pattern 5 (Chatbot + Rules) | Etikala et al. 2021 (chatbots from DMN) | RuleGo MCP server + LLM chatbot |

## What remains open

Feillet's two articles and the KU Leuven research program together cover substantial ground. But several questions remain open.

**Vendor neutrality.** Feillet's patterns are general, but the implementations assume IBM ODM, ADS, and watsonx.ai. The same patterns can be implemented with open-source rule engines — Drools and OpenL Tablets in the Java world, [RuleGo](https://github.com/rulego/rulego) in the Go ecosystem — as well as cloud decision services (AWS, GCP, Azure) and any LLM platform. RuleGo is a particularly clean example because its DAG-based component model maps naturally onto the composite patterns: an LLM call is just another node in the rule chain graph, with the same interface as any other component. A vendor-neutral reference architecture, specifying the abstract interfaces between the LLM and rule engine components, would make these patterns more broadly applicable across the full range of implementations.

**Testing composite systems.** Neither the practitioner articles nor the academic papers develop a testing methodology for composite AI. How do you test a system where one component is deterministic and the other is probabilistic? You need property-based testing: the system's output must satisfy certain invariants (decision traceability, mandatory disclosures, schema conformance) even though the surface text varies.

**The extraction quality bar.** The KU Leuven papers demonstrate that extraction works — but how accurate does it need to be to change the economics? If the extraction pipeline achieves 90% accuracy, a human reviewer must still check every rule. The cost savings come from turning a *writing* task into a *reviewing* task — which is faster, but the human is still in the loop for every rule. At what accuracy threshold does the review model shift from "review every rule" to "review only low-confidence extractions"? This is the operational question that determines whether Pattern 4 pays off.

**Rule lifecycle management.** When source documents change, extracted rules must change. This synchronization problem — governed policy evolution, versioned rule extraction, conflict detection between old and new rules — is the long tail of Pattern 4. Neither the enterprise patterns nor the academic papers provide a complete lifecycle model. It is where the next wave of research and engineering needs to go.

**The MCP interface standard.** Feillet's 2026 article treats MCP as the protocol that enables hybrid inference. But MCP is still evolving. The interface between an LRM and a decision service — what information does the LRM pass? How does it discover available decision services? How does the rule engine communicate confidence, conditions, or partial results back? — needs standardization if composite AI is to move beyond bespoke integrations.

## Summary: the composite AI thesis

![Figure 6: Composite AI — Blending Neuronal and Symbolic Approaches](images/composite-ai-fig6-overview.svg)

The central thesis — running through Feillet's two articles and validated by the KU Leuven research — is that the future of enterprise AI is composite. Not LLMs replacing rule engines, not rule engines staying isolated, but a deliberate blending where each technology does what it is good at.

LLMs are good at the *perception* layer: understanding unstructured text, classifying intents, extracting entities, generating fluent natural language. They are statistical engines optimized for flexibility and fluency.

Rule engines, built on Forgy's RETE insight from 1979, are good at the *reasoning* layer: applying deterministic logic to structured data, producing auditable decisions, enforcing regulatory constraints, scaling to high transaction volumes with predictable performance. They are logical engines optimized for reliability and transparency.

The five patterns are five ways to draw the boundary between these layers. The boundary shifts depending on the use case, the risk tolerance, and the maturity of the integration. But the principle is constant: let the neuronal system handle the messiness of natural language, and let the symbolic system handle the precision of business logic.

The 40-year arc from RETE to MCP is not a story of obsolescence and replacement. It is a story of infrastructure that works being augmented by capabilities that are new. The RETE network runs inside your decision engine, matching facts against rules with Boolean precision. The attention mechanism runs inside your LLM, computing relevance scores over learned representations. Both are answering the same question — *what is relevant right now?* — at different layers of the stack. The winning architecture uses both, each accountable for what it does best.

---

**References:**

1. Charles L. Forgy. [*Rete: A Fast Algorithm for the Many Pattern/Many Object Pattern Match Problem*](https://doi.org/10.1016/0004-3702(82)90020-0). Artificial Intelligence, 19(1): 17–37, 1982.

2. Pierre Feillet, Allen Chan, Luigi Pichett, Yazan Obeidi. [*Approaches in Using Generative AI for Business Automation: The Path to Comprehensive Decision Making*](https://medium.com/@pierrefeillet/approaches-in-using-generative-ai-for-business-automation-the-path-to-comprehensive-decision-3dd91c57e38f). Medium, August 4, 2023.

3. Pierre Feillet. [*Rule Engines Never Died — They're Running Alongside Your Large Reasoning Models*](https://medium.com/@pierrefeillet/rule-engines-never-died-theyre-running-alongside-your-lrm-6f39cad6e1d3). Medium, June 4, 2026.

4. Alexandre Goossens, Johannes De Smedt, Jan Vanthienen. [*Extracting Decision Model and Notation Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667). Expert Systems with Applications, 211: 118667, 2023.

5. Alexandre Goossens, Simon Vandevelde, Jan Vanthienen, Joost Vennekens. [*GPT-3 for Decision Logic Modeling*](https://ceur-ws.org/Vol-3485/paper3896.pdf). RuleML+RR Companion, CEUR Vol. 3485, 2023.

6. Alexandre Goossens, Ulysse Maes, Yves Timmermans, Jan Vanthienen. [*Automated Intelligent Assistance with Explainable Decision Models in Knowledge-Intensive Processes*](https://doi.org/10.1007/978-3-031-25383-6_3). BPM Workshops 2022, LNBIP 460, pp. 25–36.

7. Alexandre Goossens, Michelle Claessens, Charlotte Parthoens, Jan Vanthienen. [*Extracting Decision Dependencies and Decision Logic from Text Using Deep Learning Techniques*](https://doi.org/10.1007/978-3-030-94343-1_27). BPM Workshops 2021, LNBIP 436, pp. 349–361.

8. Vedavyas Etikala, Jan Vanthienen. [*An Overview of Methods for Acquiring and Generating Decision Models*](https://doi.org/10.1007/978-3-030-82153-1_17). KSEM 2021, LNCS 12817, pp. 200–208.

9. Vedavyas Etikala, Alexandre Goossens, Ziboud Van Veldhoven, Jan Vanthienen. [*Automatic Generation of Intelligent Chatbots from DMN Decision Models*](https://doi.org/10.1007/978-3-030-91167-6_10). RuleML+RR 2021, LNCS 12851, pp. 142–157.

10. Jan Vanthienen, Alexandre Goossens. [*GPT-3 for Decision Requirements Modeling and Advice*](https://decisioncamp2023.wordpress.com/). DecisionCamp 2023 (Slides and Recording, RuleML Paper).

11. [RuleGo](https://github.com/rulego/rulego) — Lightweight, high-performance, component-based rule engine for Go. Apache 2.0. Includes [rulego-components-ai](https://github.com/rulego/rulego-components-ai) for LLM integration and MCP server/client support.

> A decision system that cannot explain itself is not enterprise-grade. A decision system that cannot handle ambiguity is not useful. The composite approach accepts both constraints and designs for them. That is the engineering move — and it has been in progress since Forgy built the first discrimination network in 1979.
