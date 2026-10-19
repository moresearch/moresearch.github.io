---
title: "AI-driven CEP"
date: 2026-10-19
slug: ai-driven-cep
summary: "Traditional complex event processing compiles hand-written patterns over time windows. AI-driven CEP separates the pipeline into a historian that keeps the timeline and a decision model that judges it, returning typed verdicts instead of text. The result is semantic aggregation without regex wrappers, and one inference pass for many operational questions."
tags: complex-event-processing, cep, stream-processing, semantic-streaming, decision-models, system-one, small-language-models, event-driven-architecture, ai
---

A payment fails three times in ninety seconds. A deploy lands two minutes earlier. A customer writes, "I've been charged twice and nobody is answering." The raw events are simple. The meaning is not.

Traditional complex event processing (CEP) is one of the most successful ideas in event-driven systems. It compiles patterns over time windows, matches them incrementally, and fires when a sequence completes. It is deterministic, fast, and explainable. It is also closed-world: it can only detect the patterns someone anticipated.

Modern streams violate that assumption. The most valuable signals are wrapped in language — log lines, support tickets, deploy notes, chat messages — and the patterns are semantic. A rule can match `ECONNRESET`. It cannot tell you whether the reset is a gateway outage, a retry bug, or a customer who pressed "pay" twice.

The fix is not to put a chat model in the hot path. It is to separate the pipeline into a **historian** that owns state and a **judge** that owns meaning. The historian does what stream processors are good at: ingest, order, window, replay. The judge is a decision model, not a text generator: it takes a case file and a batch of declarative analytical questions and returns typed verdicts with probabilities. Ordinary code decides what to do. That is AI-driven CEP.

## The Dawn of Semantic Streaming

CEP has a precise lineage. David Luckham's *The Power of Events* (2002) framed it as the analysis of event clouds to find complex events; Cugola and Margara's survey defines it as processing event streams to identify meaningful events and react to them. Engines such as [Esper](https://www.espertech.com/esper/), [Flink CEP](https://nightlies.apache.org/flink/flink-docs-release-1.20/docs/libs/cep/), [Drools Fusion](https://docs.drools.org/), and [SASE](https://doi.org/10.1145/1142473.1142520) compile patterns into automata or match networks and evaluate them incrementally as events arrive.

A rule looks like this, in pseudocode:

```text
SEQ(payment_failed, payment_failed, payment_failed)
  WITHIN 10 minutes
  PARTITION BY customer_id
  -> emit duplicate_charge_risk
```

The engine keeps partial matches, applies temporal constraints, and fires when the pattern completes. [Rete](https://doi.org/10.1016/0004-3702%2882%2990020-0) is the classic algorithm for the many-pattern/many-object problem; streaming engines add windows, watermarks, and state. [Bucchi, Grez, and Riveros](https://arxiv.org/abs/1709.05369) give the formal semantics, and [Stonebraker's eight requirements](https://doi.org/10.1145/1107499.1107504) remain a good acceptance test for the streaming layer. The result is exactly what an operator wants from a rule: deterministic, cheap, and legible.

The trouble is that the rule encodes an ontology. It knows the event types its author anticipated, the field names the schema exposes, and the string variants the regex covers. The long tail is everything else:

- `connection reset by peer`, `ECONNRESET`, `socket hang up`, and `upstream prematurely closed connection` are the same failure in four strings.
- A customer writes "I was billed twice"; the payment events show one capture and one void. The rule sees a failed payment, not a duplicate charge.
- A deploy causes a burst of retries. The rule sees a spike; the operator needs to know whether the spike is benign.
- A new error code appears. The rule is silent — not because nothing happened, but because nobody wrote the branch.

You can fix each case with another rule. That is the trap. Rule engines scale by accretion: every missed variant becomes a regex, every new context becomes a condition, every exception becomes a flag. The rule set becomes a second-order program that nobody can test exhaustively. The failure mode is not a crash; it is a slow drift into false positives and alert fatigue.

The temptation to put a generative large language model (LLM) in the loop is understandable. Hand it the logs and ask what happened. But a chat model in the hot path converts a fixed cost into a variable cost, adds seconds of latency, and returns prose that must be parsed. It can hallucinate an event, ignore a constraint, or follow an instruction embedded in a log line. The output is unbounded, so the application needs a parser, a validator, and a fallback path. You have replaced rule sprawl with prompt sprawl.

The missing primitive is narrower: a **decision** over a bounded set of answers, with a probability attached, produced by a zero-text-generation decision model.

That primitive now exists as a product category. TypeSafe's [Jev](https://typesafe.ai/blog/introducing-system-one-models-and-jev) is a System One model: unstructured state in, typed probabilistic decisions out. It gives up string generation entirely. Together AI's open-weight [Tev1](https://github.com/togethercomputer/tev1) is a 0.8B/4B decision model fine-tuned from Qwen3.5 that picks from a list, answers true or false, or places an item on a rubric, and can answer up to 64 questions about the same text in one call. Tev1 is not zero-text-generation in the same sense — it emits a constrained answer token — but its output space is still finite, and the application never parses prose. Convai's [Laya](https://huggingface.co/convaiinnovations/laya) is the Apache-2.0, non-autoregressive counterpart: typed answers in a single forward pass, no text generation, and nothing to parse. Microsoft's [Decision-1](https://commandline.microsoft.com/microsoft-decision-1-model-foundry/) is a decision-scoring model aimed at routing, classification, prioritization, verification, and workflow control. The model category is new, and the benchmarks are still mostly vendor-run. The architectural idea — a learned classifier inside a stateful pipeline — is not.

The word "semantic" has a history in CEP, and it is worth keeping the two meanings apart. Semantic CEP and stream reasoning already applied ontologies, RDF streams, and logical inference to event processing; [Dell'Aglio, Della Valle, and van Harmelen](https://doi.org/10.3233/DS-170006) survey that line of work, and [DSCEP](https://arxiv.org/abs/2002.05869) is a distributed infrastructure for it. That is formal entailment over structured knowledge. The semantic streaming described here is narrower and learned: natural-language understanding over unstructured context, with a typed decision as the output. Both are useful. They answer different questions.

## Separation of State and Judgment

The architecture that works splits the pipeline into two decoupled layers, with a third component coordinating them.

```text
raw events ──▶ historian ──▶ case file ──▶ judge ──▶ verdicts
                 │                                      │
                 └──────────── event log ◀──────────────┘
```

### The historian

The historian is the high-throughput data infrastructure. It filters and routes the stream, ingests messages, validates schemas, orders events, handles late arrivals, maintains watermarks, and caches a rolling chronological timeline. It is event-sourced: the raw stream is append-only, immutable, and replayable. It also owns entity state and retention policy. Kafka Streams, Flink, or a dedicated event store can host it; the point is that the stateful, replayable part of the pipeline stays in infrastructure built for state and replay.

The historian does not decide what the events mean. It preserves the evidence and makes it addressable. That is a narrower job than "understand the stream," and it is the job that stream processors already do well. The [Dataflow model](https://doi.org/10.14778/2824032.2824076) gives the vocabulary for correctness over unbounded, out-of-order data; [event sourcing](https://martinfowler.com/eaaDev/EventSourcing.html) gives the vocabulary for treating the log as the source of truth. The historian is where those two ideas meet.

When the application needs a judgment, the historian produces a **case file**: a bounded, ordered window of events for one entity, plus the structured state and policy text the decision needs. The case file is the unit of context. It has a token budget, a serializer version, and an evidence order. It is not the whole stream, and it is not a random sample. It is the smallest timeline that can support the question.

### The judge

The judge is a decision model. It receives the case file and a set of declarative questions. It returns a typed answer for each question — a yes/no probability, a choice among options, or a score on a rubric — with probabilities, and a confidence where the primitive provides one. It does not generate text, choose its own actions, or write to the stream. It is a pure function of the case file and the question set. The case file is the engineering timeline; the verdict batch is where it becomes operational clarity.

The [System One documentation](https://docs.typesafe.ai/concepts/system-one) describes the contract precisely: one state, one or more questions, all questions evaluated independently, answers returned with probabilities. The primitives are `Choice`, `Score`, and `Noul` (a yes/no probability). A request can mix them. A `Choice` question enumerates its options; a `Score` question enumerates its levels; a `Noul` question has two sides. The model's output space is fixed before inference starts.

That fixed output space is the difference between a judge and a chatbot. There is no prose to parse, no JSON to repair, no "answer yes or no" prompt followed by a parser that accepts "Yes.", "yes", "Y", and "Absolutely." The application validates the distribution against a schema and thresholds it. If the top choice is below a confidence threshold, or the distribution is too flat, the application can abstain and route the case to a human or a slower reasoning path.

### The application agent

Between the historian and the judge sits the localized application agent — not an autonomous LLM, but ordinary software running next to the historian. It decides when to judge, how to pack the case file, which questions to ask, how to threshold the answers, and what side effects to perform. This is where control flow lives.

The split gives each layer a single responsibility:

- The historian is the source of truth. It can replay a day of events against a new judge offline.
- The judge is replaceable and testable. It has no hidden state; the same case file, question set, and model version should produce the same distribution, modulo floating-point and serving nondeterminism.
- The application agent is ordinary code. It can enforce hard constraints, rate limits, human gates, and audit logging.

A sketch in Go:

```go
type Event struct {
	ID      string
	Entity  string
	Time    time.Time
	Kind    string
	Text    string
	Payload map[string]any
}

type Question struct {
	Name         string
	Type         string // "noul", "choice", or "score"
	Instructions string
	Criteria     any // options, levels, or yes/no descriptions
}

type Verdict struct {
	Name         string
	Choice       string
	Probability  float64
	Distribution map[string]float64
}

func (p *Pipeline) OnEvent(ctx context.Context, e Event) error {
	// The historian appends the raw event and maintains the window.
	if err := p.history.Append(ctx, e); err != nil {
		return err
	}
	window, err := p.history.Window(ctx, e.Entity, 15*time.Minute)
	if err != nil {
		return err
	}

	// Ordinary code decides whether this event is worth judging.
	if !p.trigger.ShouldJudge(e, window) {
		return nil
	}

	// The case file is ordered, bounded, and versioned.
	caseFile := p.pack(window)

	// The judge answers all questions in one request.
	verdicts, err := p.judge.Decide(ctx, caseFile, p.questions)
	if err != nil {
		return err
	}

	// Verdicts become derived events, so they compose with the rest of the stream.
	return p.emit(ctx, e.Entity, verdicts)
}
```

The important line is `p.emit`. A verdict is not a dashboard annotation; it is a new event. It can trigger a rule, feed another window, open a ticket, or be replayed. The loop is closed.

### The runtime objection

There is a serious objection to putting any model in the event loop, and it is the subject of an [earlier post on this site](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern): a per-event model call turns a fixed cost into a variable cost and gives up replayability, auditability, versioning, and determinism. That objection still applies here. It sets the bar for this architecture rather than disqualifying it.

The first answer is that the historian is a cost governor. The judge should not run on every event. It should run when a deterministic trigger fires: a state transition, a threshold crossing, a new case, a clock deadline, or a rule that returns "uncertain." The historian's job is to reduce millions of events into thousands of cases. The inference bill scales with cases, not with the raw stream.

The second answer is that the judge should be a frozen artifact, not a vendor endpoint that changes under you. Fine-tune it, evaluate it against a golden set, hash it, and pin it. Record the model hash, the question-set hash, the serializer version, and the threshold table with every emitted verdict. That is what makes replay possible. Where the decision can be compiled into a deterministic rule or a small tree model, do that instead; the decision model is for the residual — cases whose meaning depends on unstructured language and cannot be captured by features cheaply. The judge is a waypoint, not a free lunch. The existing post's warning is the right default: convert variable costs into fixed ones wherever you can, and be honest about the cases where you cannot.

The third answer is locality. A 0.8B model can run on the same machine as the stream processor; a 4B model can run on a modest GPU. That removes the network round trip, keeps sensitive data inside the boundary, and makes the latency budget predictable. It does not make inference free, but it makes the cost legible.

## Multi-Dimensional Verdicts in a Single Pass

The superpower of this design is not that the model is smart. It is that the model's output is shaped like software.

### Semantic aggregation over strings

The historian stores the raw text; the judge reads the window as a whole. That lets it do something a regex cannot: aggregate semantically across events. It can recognize that four different log lines describe the same failure even when their formatting and wording differ, that three customer messages express the same intent, or that a sequence of retries is one incident rather than four independent errors. It can weigh repetition, recency, and context instead of matching one line at a time.

This is where the log-parsing literature becomes relevant. LLM-based parsers such as [LLMParser](https://arxiv.org/abs/2404.18001), [LILAC](https://arxiv.org/abs/2310.01796), and the systems surveyed by [Beck et al.](https://arxiv.org/abs/2504.04877) show that language models can normalize heterogeneous log messages into templates without hand-written regexes. [Loghub](https://arxiv.org/abs/2008.06448) is the dataset collection behind much of that work. But parsing is only the first step. A template tells you what kind of event occurred; it does not tell you whether the case is a duplicate charge, a policy violation, or a benign retry storm. The judge operates on the parsed *and* raw evidence and returns the operational meaning.

The same shift shows up in the research on LLM-augmented CEP pipelines. [Zeeshan, Kumar, and Pirttikangas](https://arxiv.org/abs/2501.00906) use a multi-agent LLM layer to augment CEP for the Internet of Multimedia Things; [Chandra, Agarwal, and Singh](https://arxiv.org/abs/2510.09646) combine ontology-driven CEP with LLM reasoning for real-time health analytics. Those systems use language models as reasoning components. The decision-model architecture here is narrower: the model is a judge with a fixed questionnaire, not a general reasoning agent.

### Deterministic output controls

A decision model's output space is declared in advance. `Choice` questions enumerate the options; `Score` questions enumerate the levels; `Noul` questions have two sides. The model returns a distribution over that space. The application validates the distribution against the schema, thresholds it, and acts.

This eliminates a whole class of glue code:

- No regex wrappers to extract a label from prose.
- No JSON repair for malformed model output.
- No silent failure when the model invents an option that does not exist.
- No prompt-injection path through *generated* text, because there is no generated text. Adversarial text in the input can still bias a judgment, which is why the judge is not a security boundary.

It also gives the application a principled way to abstain. If the top choice's probability is below a threshold, or the distribution is too flat, route to a human or a slower reasoning path. Calibration research is clear that neural confidence is not automatically trustworthy; [Guo et al.](https://arxiv.org/abs/1706.04599) showed that modern neural networks are often overconfident, and the problem has not disappeared. Conformal prediction offers a distribution-free way to set abstention thresholds with coverage guarantees ([Angelopoulos and Bates](https://arxiv.org/abs/2107.07511)), at the cost of exchangeability assumptions that streaming data often violate. The engineering answer is to measure calibration per question, per slice, and per time period, and to treat confidence as a signal to be validated rather than a fact to be trusted.

### Constant-factor performance

The performance argument is simple: the judge evaluates the entire question batch against one shared state in one request. Adding a question does not add a round trip or a separate request; the model emits the verdicts together. TypeSafe's documentation describes questions as evaluated independently and in parallel and reports most queries around 100 ms; the Jev announcement claims 70–500 ms end to end and a parallel sampler that produces all outputs in a single query. Tev1 accepts up to 64 questions about the same text in one call. Microsoft's Decision-1 announcement claims 2.5× the speed of the runner-up decision model and 35× the speed of GPT-6 Sol on its benchmark. These are vendor claims, not independent measurements, but they are consistent with the same architectural claim: the cost is dominated by encoding the shared state, not by the number of questions.

"Constant-factor" is the right intuition, not a law. Attention over a longer state still costs more; a 64-question batch still has more output heads than a three-question batch; a 4B model is slower than a 0.8B model. The win is that the number of *decisions* is decoupled from the number of *model invocations*. You pay for the case file, not for each question. Whether the batch contains three self-checks or twenty detection rules, the number of model invocations is one. And the decision count still scales with the number of cases. That is why the historian's trigger policy matters as much as the judge's latency.

| | Rule-based CEP | Generative LLM in the loop | Decision-model judge |
|---|---|---|---|
| Pattern specification | Hand-written rules and schemas | Natural-language prompt | Declarative question set |
| Input | Structured fields | Text | Text + structured state |
| Output | Pattern match | Free text | Typed verdicts + distributions |
| Determinism | High | Low | Bounded output space; versioned artifact |
| Latency | Microseconds to milliseconds | Seconds | Tens to hundreds of milliseconds |
| Cost | Low | High and variable | Low and predictable per case |
| Failure mode | Missed variants, rule sprawl | Hallucination, parsing, prompt injection | Misclassification, calibration drift |
| Audit trail | Rule trace | Opaque | Evidence + probabilities + model hash |

## A Worked Case: Duplicate Charges

Consider a payments pipeline. The raw stream contains `payment_attempt`, `gateway_response`, `charge`, `refund`, `support_ticket`, `deploy`, and `metric_alert` events. A legacy rule detects three failed attempts in ten minutes and opens a "payment failure" incident.

The rule is right about the failures and wrong about the case. The customer's ticket says "charged twice"; the payment events show one capture and one void; the gateway log says `upstream prematurely closed connection`; the deploy event landed two minutes before the spike. The rule fires on the failures and misses the duplicate charge. A second rule for duplicate charges needs a definition of "duplicate" across authorization, capture, void, and refund events, plus the customer's intent. That definition is exactly the kind of thing that is easy to state in English and hard to encode in a pattern language.

With the historian/judge split, the application packs a fifteen-minute case file for the customer: the ordered events, the ticket text, the deploy record, and the relevant policy. The judge answers a batch:

- `duplicate_charge` (boolean): does the evidence indicate the customer was charged more than once for the same order?
- `customer_intent` (choice): refund, cancellation, billing question, fraud report, other.
- `root_cause` (choice): gateway, retry bug, deploy, issuer, unknown.
- `severity` (score 0–4).
- `requires_human` (boolean).

The application thresholds the answers: a high-probability duplicate charge with high severity opens a refund case and pages the payments on-call; a flat distribution routes to a human; a clear gateway signature links the incident to the deploy. The verdict is emitted as a `case_triaged` event. The next window can reason over both the raw events and the verdict.

The deterministic rule still has a job. It handles the obvious case cheaply and provides a hard constraint: if the rule fires, the case must be reviewed, regardless of what the judge says. The judge handles the ambiguous case. This is a hybrid, not a replacement.

## Designing the Judge

The question set is the new rule set, and it deserves the same engineering discipline.

Not every judgment needs a decision model. If the task is pure classification with plenty of labels, a fine-tuned encoder such as [ModernBERT](https://arxiv.org/abs/2412.13663) or a few-shot method such as [SetFit](https://arxiv.org/abs/2209.11055) can be cheaper and easier to operate. The decision model earns its place when the input is unstructured, the question set is heterogeneous, and the output needs calibrated probabilities over a small option space.

**Ask narrow questions.** "What happened?" is not a decision. "Does the evidence indicate a duplicate charge?" is. Each question should have a clear answer space and explicit criteria for each option. If a question needs a paragraph of reasoning, split it.

**Keep questions independent.** The System One documentation is explicit: all questions see the same state and are evaluated independently. That is a feature — no question's answer becomes hidden context for another — but it means you cannot ask a follow-up that depends on a previous verdict in the same batch. If you need a cascade, make two batches and let the application agent route between them.

**Include the null option.** A choice question should include "none of the above" or "not enough information." Without it, the model is forced to choose among bad options. The `Noul` primitive gives you a probability for yes; for questions where "unknown" is a real answer, use a `Choice` with a "not enough information" option and calibrate the abstention threshold.

**Keep the option set small.** Decision models are trained on a bounded label space. Tev1's training range is 2–24 options, and its documentation recommends staying in that range. If you need hundreds of labels, decompose the decision into a hierarchy of smaller questions rather than one giant choice.

**Version everything.** The question set, the criteria text, the state serializer, the window policy, the model hash, and the threshold table are all part of the decision. A change to any of them can change the verdict. Store the version with every emitted event so you can replay and explain.

**Evaluate per question, not just end to end.** Build a golden set of case files with labels. Measure precision, recall, and calibration for each question; slice by event type, customer segment, language, and time. A model that is 95% accurate overall can be 60% accurate on the slice that matters. Report abstention rate and the cost of human review, not just accuracy.

**Monitor drift.** The input distribution changes when a new service ships, a log format changes, or customer behavior shifts. The label distribution changes when the business changes. Monitor the distribution of verdicts, the confidence distribution, and the agreement between the judge and human reviewers. Re-run the golden set on every model or question-set change. Treat the judge like any other production model: it needs a model card, an owner, and a rollback plan.

## What Can Go Wrong

The design is not magic. It moves the failure modes; it does not eliminate them.

- **The judge is confidently wrong.** A constrained output space prevents prose hallucinations but not misclassification. Calibration and abstention are not optional.
- **The case file is incomplete.** If the historian truncates the window or drops late events, the judge reasons over a false timeline. Watermarks and retention policies are part of decision quality.
- **The question is ambiguous.** Two engineers can read the same criteria differently. The model will too. Write criteria as if they were tests, and measure inter-rater agreement on the labels.
- **The model is stale.** A fine-tuned decision model is a snapshot of the world at training time. New failure modes require new data and a new model. The historian's replay log is the raw material for that retraining.
- **The judge becomes an oracle.** The most dangerous failure is organizational: treating a probability as a verdict. The output is evidence for a decision, not the decision itself. Keep hard rules, human gates, and audit trails where the consequences are high.
- **The economics invert.** If every event triggers a judgment, the inference bill scales with the stream. Use deterministic triggers, sampling, and caching. The historian should earn its name by reducing the number of cases the judge sees.

## The Boundary

AI-driven CEP is not a claim that rules are obsolete. It is a claim about where to put the boundary. Rules are excellent at exact constraints, safety invariants, and cheap filters. Decision models are excellent at ambiguity, paraphrase, intent, and context. The historian is excellent at state, order, and replay. Put each where it belongs.

The shift is already visible in the tooling. Decision models are emerging as a product category: TypeSafe's Jev, Together AI's open-weight Tev1, Microsoft's Decision-1. Small language models make local inference practical; constrained decoding makes typed outputs reliable; log-parsing research shows that the unstructured part of the stream is tractable. The remaining work is architectural, not just modeling: build the historian, define the questions, calibrate the thresholds, and keep the control flow in code.

Traditional CEP asked: *what pattern matches this stream?* AI-driven CEP asks: *what is the meaning of this case, and what decision does it support?* The first question is still necessary. The second is the one that operations has always been trying to answer.

---

**References and related reading:**

**CEP foundations**

- David Luckham, [*The Power of Events: An Introduction to Complex Event Processing in Distributed Enterprise Systems*](https://www.informit.com/store/power-of-events-an-introduction-to-complex-events-and-9780201727890) (Addison-Wesley, 2002).
- Gianpaolo Cugola and Alessandro Margara, ["Processing flows of information: From data stream to complex event processing,"](https://doi.org/10.1145/2187671.2187677) *ACM Computing Surveys* 44(3), 2012.
- Eugene Wu, Yanlei Diao, and Shariq Rizvi, ["High-performance complex event processing over streams,"](https://doi.org/10.1145/1142473.1142520) SIGMOD 2006.
- Charles L. Forgy, ["Rete: A fast algorithm for the many pattern/many object pattern match problem,"](https://doi.org/10.1016/0004-3702%2882%2990020-0) *Artificial Intelligence* 19(1), 1982.
- Michael Stonebraker, Uğur Çetintemel, and Stan Zdonik, ["The 8 requirements of real-time stream processing,"](https://doi.org/10.1145/1107499.1107504) *SIGMOD Record* 34(4), 2005.
- Marco Bucchi, Alejandro Grez, and Cristian Riveros, ["Foundations of Complex Event Processing,"](https://arxiv.org/abs/1709.05369) arXiv:1709.05369.
- Apache Flink, ["CEP: Complex Event Processing."](https://nightlies.apache.org/flink/flink-docs-release-1.20/docs/libs/cep/)
- EsperTech, ["Esper: Event Processing."](https://www.espertech.com/esper/)
- Drools, ["Complex Event Processing."](https://docs.drools.org/)

**Streaming and event sourcing**

- Tyler Akidau et al., ["The Dataflow Model,"](https://doi.org/10.14778/2824032.2824076) *PVLDB* 8(12), 2015.
- Tyler Akidau, ["The World Beyond Batch: Streaming 101."](https://www.oreilly.com/radar/the-world-beyond-batch-streaming-101/)
- Jay Kreps, ["The Log: What every software engineer should know about real-time data's unifying abstraction."](https://www.linkedin.com/blog/engineering/distributed-systems/log-what-every-software-engineer-should-know-about-real-time-datas-unifying)
- Martin Fowler, ["Event Sourcing."](https://martinfowler.com/eaaDev/EventSourcing.html)
- Martin Fowler, ["What do you mean by event-driven?"](https://martinfowler.com/articles/201701-event-driven.html)
- Apache Kafka, ["Kafka Streams."](https://kafka.apache.org/documentation/streams/)

**Semantic CEP and stream reasoning**

- Daniele Dell'Aglio, Emanuele Della Valle, and Frank van Harmelen, ["Stream reasoning: A survey and outlook,"](https://doi.org/10.3233/DS-170006) *Data Science* 1(1–2), 2017.
- Vitor Pinheiro de Almeida, Sukanya Bhowmik, and Markus Endler, ["DSCEP: An Infrastructure for Distributed Semantic Complex Event Processing,"](https://arxiv.org/abs/2002.05869) arXiv:2002.05869.
- Gulay Unel, ["Stream Reasoning on Expressive Logics,"](https://arxiv.org/abs/1808.04738) arXiv:1808.04738.
- Artur d'Avila Garcez and Luis C. Lamb, ["Neurosymbolic AI: The 3rd Wave,"](https://arxiv.org/abs/2012.05876) arXiv:2012.05876.

**Decision models and small models**

- TypeSafe AI, ["Introducing System One Models & Jev,"](https://typesafe.ai/blog/introducing-system-one-models-and-jev) September 15, 2026.
- TypeSafe AI, ["System One" documentation.](https://docs.typesafe.ai/concepts/system-one)
- TypeSafe AI, ["How to build with TypeSafe."](https://docs.typesafe.ai/concepts/how-to-build-with-system-one)
- Together AI, ["tev1" open-weight decision model.](https://github.com/togethercomputer/tev1) See also the [Ollama model page](https://ollama.com/library/tev1) and the [4B weights](https://huggingface.co/togethercomputer/Tev1-4B-experimental).
- Together AI, ["How to train your own Jev."](https://www.together.ai/blog/how-to-train-your-own-jev)
- Convai Innovations, [Laya](https://huggingface.co/convaiinnovations/laya) — an Apache-2.0, non-autoregressive System One decision model.
- Microsoft, ["Introducing Microsoft-Decision-1, our model for fast decision-making,"](https://commandline.microsoft.com/microsoft-decision-1-model-foundry/) October 9, 2026.
- Peter Belcak et al., ["Small Language Models are the Future of Agentic AI,"](https://arxiv.org/abs/2506.02153) arXiv:2506.02153.
- Zhenyan Lu et al., ["Small Language Models: Survey, Measurements, and Insights,"](https://arxiv.org/abs/2409.15790) arXiv:2409.15790.
- Lewis Tunstall et al., ["Efficient Few-Shot Learning Without Prompts" (SetFit),](https://arxiv.org/abs/2209.11055) arXiv:2209.11055.
- Benjamin Warner et al., ["Smarter, Better, Faster, Longer: A Modern Bidirectional Encoder for Fast, Memory Efficient, and Long Context Finetuning and Inference,"](https://arxiv.org/abs/2412.13663) arXiv:2412.13663.

**Structured outputs and evaluation**

- Brandon T. Willard and Rémi Louf, ["Efficient Guided Generation for Large Language Models,"](https://arxiv.org/abs/2307.09702) arXiv:2307.09702.
- Saibo Geng et al., ["Grammar-Constrained Decoding for Structured NLP Tasks without Finetuning,"](https://arxiv.org/abs/2305.13971) arXiv:2305.13971.
- Chuan Guo et al., ["On Calibration of Modern Neural Networks,"](https://arxiv.org/abs/1706.04599) arXiv:1706.04599.
- Anastasios N. Angelopoulos and Stephen Bates, ["A Gentle Introduction to Conformal Prediction and Distribution-Free Uncertainty Quantification,"](https://arxiv.org/abs/2107.07511) arXiv:2107.07511.
- Nelson F. Liu et al., ["Lost in the Middle: How Language Models Use Long Contexts,"](https://arxiv.org/abs/2307.03172) arXiv:2307.03172.

**Logs and operational context**

- Jieming Zhu et al., ["Loghub: A Large Collection of System Log Datasets for AI-driven Log Analytics,"](https://arxiv.org/abs/2008.06448) arXiv:2008.06448.
- Viktor Beck et al., ["System Log Parsing with Large Language Models: A Review,"](https://arxiv.org/abs/2504.04877) arXiv:2504.04877.
- Van-Hoang Le and Hongyu Zhang, ["Log Parsing: How Far Can ChatGPT Go?"](https://arxiv.org/abs/2306.01590) arXiv:2306.01590.
- Zeyang Ma et al., ["LLMParser: An Exploratory Study on Using Large Language Models for Log Parsing,"](https://arxiv.org/abs/2404.18001) arXiv:2404.18001.
- Zhihan Jiang et al., ["LILAC: Log Parsing using LLMs with Adaptive Parsing Cache,"](https://arxiv.org/abs/2310.01796) arXiv:2310.01796.

**AI and CEP**

- Talha Zeeshan, Abhishek Kumar, and Susanna Pirttikangas, ["Large Language Model Based Multi-Agent System Augmented Complex Event Processing Pipeline for Internet of Multimedia Things,"](https://arxiv.org/abs/2501.00906) arXiv:2501.00906.
- Ritesh Chandra, Sonali Agarwal, and Navjot Singh, ["Real-Time Health Analytics Using Ontology-Driven Complex Event Processing and LLM Reasoning: A Tuberculosis Case Study,"](https://arxiv.org/abs/2510.09646) arXiv:2510.09646.
- [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) — the compile-time/runtime wall that this architecture has to respect.
