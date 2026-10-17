---
title: Software Factory Engineering
date: 2026-10-17
slug: software-factory-engineering
summary: "Dark factories are production functions, not rooms full of agents. Patrick Debois's account of building five or six factories at Tessl shows why a factory has to be specialized around a class of software artifact products, owned like a product, and verified like a production line. It also introduces the role of software factory engineer, a factory development lifecycle, and the case for engineering the factory as the environment in which coding agents operate."
tags: software-factory, dark-factory, agents, environment-engineering, software-engineering, factory-engineering, software-factory-engineer, software-factory-development-lifecycle, specialization, verification, patrick-debois
---

Ask whether software engineering is becoming a factory and the answer is probably yes. The more useful question is: *which* factory?

Patrick Debois — one of the people who gave the industry DevOps — has been building software factories at Tessl. In a recent interview, [*Is the Future of Software Engineering a Factory?*](https://www.youtube.com/watch?v=FnCnv2tGsK0), he describes what happened when the whole organization was told to build one: "we ended up with five or six different factories." Not because the team wanted duplication, but because the work was not the same. As he puts it, "not all projects are similar." The factories differed by repository shape, tooling, context, and workflow.

[![Patrick Debois interviewed about software factories: five or six factories, ownership, and what it takes to survive production](images/software-factory-engineering-video.jpg)](https://www.youtube.com/watch?v=FnCnv2tGsK0)

Then came the message that captures the whole problem. One morning, someone asked on Slack: the factory had created fifteen tickets about improving itself. *Who owns this?*

That question is the subject of this post, and it rests on one key insight: **an AI coding agent is only half of the system. The other half is the environment it acts in.** The agent literature treats the environment as part of the abstraction. The software industry has spent its energy on the agent — models, prompts, context, harnesses, loops — and left the environment as a fixed backdrop. If we engineer the environment as deliberately as we engineer the agent, we can offload many of the hardest parts of agent engineering into environment engineering. The software factory should be that environment.

A dark factory is not a room full of agents producing "software." It is a production function and a designed world, with a product, a boundary, an owner, and a verification regime. The discipline that designs them is **software factory engineering**.

## The factory is a production function

Every factory, physical or software, has four parts:

- **Inputs:** specifications, domain knowledge, reference artifacts, fixtures, environments, credentials, and model access.
- **Transformation:** agent roles, prompts, harnesses, tools, loops, review passes, and escalation paths.
- **Outputs:** not software in general, but artifacts with an expected shape — a service, a CLI, a migration, a Terraform module, a data pipeline, a test suite.
- **Control:** verification gates, acceptance criteria, observability, yield accounting, and human review points.

Most dark-factory discussion obsesses over the transformation — which model, which framework, which loop. That is the visible factory floor. The leverage is in the other three. Inputs determine whether the factory can start. Outputs determine what it is for. Control determines whether anyone can trust what comes out.

This is not a new idea. Robert W. Bemer's 1968 "machine-controlled production environment, or software factory" already included compilation, test-system construction, documentation, quality control, instrumentation, scheduling, and costing. The software factory has always meant organized production. What changed is that the machinery can now perform the steps that used to require human judgment at every turn.

## The key insight: the agent is only half the system

The standard definition is simple: an agent perceives its environment, acts on it, and receives feedback. The environment is not scenery. It defines what the agent can observe, what actions are available, what those actions cost, and what counts as success. From Russell and Norvig's agent-environment framing to the multi-agent systems literature, the definition is relational: an agent is always an agent *in* something.

The software industry's AI-coding conversation has been strangely one-sided. It has focused almost entirely on the agent: better models, better prompts, better context, better harnesses, better loops. Even the harness conversation is usually agent-side — how do we give the agent more tools, more memory, more autonomy? The environment is treated as a fixed backdrop: the repository, the build system, the ticket tracker, production.

That framing puts every hard problem inside the agent. It has to discover the build command, remember the conventions, infer the architecture, choose the safe action, recover from its own mistakes, and avoid the destructive one. Then we are surprised when it fails at things a human absorbed through years of context.

The inversion is to engineer the environment to fit the agent — as deliberately as we engineer the agent itself. Make the repository legible: one canonical build command, a reproducible development container, typed interfaces, fast tests, structured error messages, linters that encode style, and architectural boundaries that make risky areas visible. Give the agent a sandbox with least privilege, snapshots, and reversible operations. Put a paved road in front of it instead of a toolbox and a warning label.

This is not cheating. It is offloading. The agent does not need to remember the convention if a linter rejects violations. It does not need to infer the deployment sequence if there is one documented path. It does not need to be trusted never to delete the database if permissions and backups make that mistake recoverable. The difficult parts of agent engineering become ordinary environment engineering: design the constraints, feedback, and affordances so that a good-enough agent produces good work.

The payoff is durability. A prompt tweak is local and fragile. An environment improvement is reusable, testable, and shared across every agent that operates in it. It also changes how we debug. When the agent fails, the first question is no longer only "what is wrong with the model?" It is also "what did the environment fail to expose, constrain, or verify?"

This is where the factory becomes more than a pipeline. **The factory is the environment in which the coding agent operates.** The specification schema defines what the agent can see. The harness defines what it can do. The sandbox defines its blast radius. The oracle defines its feedback. The reference corpus defines what good looks like. The product-line charter defines what is out of bounds. The factory's job is to make the right action easy, the wrong action hard, and the result checkable.

Once the factory is the environment, specialization stops being an organizational preference and becomes an engineering one. A specialized factory is a specialized world. A general factory is a generic world, and a generic world forces the agent to carry every assumption we failed to encode. The more we can put into the environment, the less the agent has to hold in its head.

## Specialization is the safety mechanism

The first instinct is to build a factory that builds anything. General-purpose coding agent, good specification, test harness — surely that covers the space. It does not, for the same reason a general-purpose physical factory does not: changeover costs dominate. A factory that can stamp metal, injection-mold plastic, and bake bread is not three factories in one. It is a warehouse full of equipment that is never optimally configured for anything.

Software looks exempt because the material is text and the machinery is a model. The real constraints are elsewhere:

**Specification closure.** A specification is executable only relative to shared assumptions. A CRUD service and a flight-control loop do not share assumptions about failure, timing, state, or evidence. A general factory must encode all of them at once — making the spec language enormous — or leave them implicit, which is the ambiguity a dark factory cannot absorb.

**Verification oracles.** The factory has to know whether an artifact is correct. For a REST service, that is an API contract plus integration tests. For a data pipeline, reconciliation against source systems. For firmware, timing analysis and hardware-in-the-loop tests. These oracles are not interchangeable. A general factory either carries every oracle or has no real oracle at all.

**Toolchain depth.** Deep automation comes from tight integration with a specific framework, build system, deployment target, observability stack, and security model. General factories integrate shallowly with everything. Specialized factories integrate deeply with a few things, and depth is where reliability lives.

**Economics.** Building a factory has a fixed cost: schemas, golden artifacts, harnesses, verifiers, eval sets, and the people who maintain them. That cost amortizes over repeated production. A factory that builds one bespoke system is a project with extra steps. A factory that builds a class of artifacts repeatedly is a production function with a real cost curve.

**Accountability.** "The factory built it" is not an answer to an auditor or an incident review. A specialized factory has a defined product and a defined acceptance regime, so it can answer *what was supposed to happen* and *how we know it did*. A general factory has a story about agents.

This is why Tessl ended up with five or six factories. Debois's lesson is not that duplication is good. It is that a factory is shaped by its product. He describes aiming for "two or three paved roads" — standard shapes the organization can support — with teams that leave the road taking on the security, observability, and maintenance burden themselves.

## The product defines the factory

The rule is simple: **a software factory should be a specialized thing that produces certain types of software artifact products.**

Not "software." Not "features." Not "whatever the backlog says." Artifact products: things with a repeatable shape, a known interface, a defined quality bar, and an acceptance procedure.

Examples of product lines:

- Internal admin surfaces over a known data model, with standard auth, audit logging, and access control.
- API integration adapters with defined contracts, retry semantics, rate limits, and reconciliation tests.
- Data pipelines with schema contracts, idempotency requirements, freshness SLAs, and reconciliation against sources.
- Cloud infrastructure modules with policy-as-code checks, cost constraints, and a conformance suite.
- Developer CLIs and code generators with a command contract, exit codes, and golden-output tests.
- Evaluation and benchmark suites with reproducible tasks, graders, and statistical acceptance criteria.
- Regulated or safety-critical components, where acceptance includes traceability, certification evidence, and independent review.

Each is a different factory. They may share infrastructure — model access, sandboxing, artifact storage, telemetry, identity — but they do not share a specification language, a verification oracle, or a definition of done.

The artifact product is the right unit because it can carry a stable contract. A "feature" changes shape with every roadmap revision. A "codebase" is the accumulated output of many product lines. An artifact product has a boundary: you can say what it is, what it must do, how to test it, and when it is done.

## What factory engineering builds

Factory engineering has concrete deliverables: a product-line charter, a specification schema, a reference corpus, a harness and toolchain, verification oracles, human gates, yield telemetry, and change management. Together they form the factory's operating system — and the basis for a lifecycle.

## The software factory engineer

A new role is emerging, and it is not a renamed software engineer. The **software factory engineer** owns the production function, not the artifacts coming out of it. Their job is to make the factory reliable, measurable, and economical.

The role is easiest to define by its questions. A software engineer asks whether a piece of code works. A software factory engineer asks whether the production system reliably produces accepted artifacts at a defensible cost — and whether the organization can prove it when something goes wrong. The artifact is a product of the factory. The factory is the engineer's product.

In practice, the role combines disciplines that used to sit in different teams: product-line design (what the factory serves and refuses), specification engineering (the input language and its validation), verification engineering (oracles and acceptance evidence), platform operations (sandboxes, credentials, observability, cost, incident response), agent engineering (harnesses, context, model upgrades), and product economics (yield, cycle time, cost per accepted artifact).

That is too much for one person at scale, and the role usually starts as a small team rather than a unicorn hire. The important distinction is ownership. A platform engineer provides shared capabilities: compute, sandboxes, identity, telemetry. A software factory engineer composes those capabilities into a production function and is accountable for its yield. The platform team can say the sandbox is available. The factory engineer has to say the factory produced twelve accepted artifacts this week, rejected three specifications as out of scope, and escaped one defect that the oracle did not catch.

This is also the answer to Debois's fifteen-ticket question. Self-improvement work is not noise. It is the factory's maintenance backlog, and it needs an owner with a roadmap. Without that role, the factory either stalls or optimizes itself in directions nobody chose.

The best candidates are systems thinkers who can say precisely what class of artifact a domain supports, how it will be verified, and what the factory should refuse to build.

## The software factory development lifecycle

Factories are products, so they need a lifecycle. Most teams jump from a working demo to production and skip the stages that make a production function trustworthy. The result is predictable: a factory with no charter, an unbounded product line, a weak oracle, and no answer to who owns the tickets it generates.

A **software factory development lifecycle** (SFDL) makes those stages explicit. The names will vary by organization; the sequence and the exit criteria matter more than the vocabulary.

**1. Fit and product-line discovery.** Identify the artifact class, its expected volume, its repeatability, and whether a real oracle is feasible. Not every domain is factory-shaped. Exit: evidence that the class is repeatable and verifiable, plus a build-buy-rent decision.

**2. Charter and design.** Define outputs, refusals, interfaces, acceptance regime, risk appetite, owner, and economic targets. Exit: a charter with measurable acceptance criteria and an accountable owner.

**3. Bootstrap.** Build the specification schema, reference corpus, harness, oracles, and one paved road. Run a single artifact end to end. Exit: a golden artifact passes the factory's own acceptance procedure, with failures diagnosed by the factory rather than by a human reading a diff.

**4. Pilot and shadow operation.** Run alongside the existing human process. Measure first-pass yield, rework, defect escape rate, cycle time, and cost per accepted artifact. Do not grant merge or deploy authority yet. Exit: a baseline, a risk review, and evidence that the oracle catches the failures that matter.

**5. Operate.** Put the factory into production with telemetry, human gates, incident response, and change management. The factory itself needs on-call responsibility: when the harness breaks, the model changes, or the sandbox leaks, someone must own the response. Exit: sustained yield inside agreed thresholds, not a launch date.

**6. Evolve or fork.** Models change, dependencies drift, and product lines expand. Re-run conformance, version the acceptance standard, and decide whether to extend the existing factory or build a new one. Exit: a documented change with the oracle re-validated against the new reality.

**7. Retire.** Stop production, migrate the artifacts, archive the corpus and oracles, and revoke credentials. Exit: no orphaned artifacts, no orphaned tickets, and no factory still running because nobody remembered to turn it off.

This is a loop, not a waterfall. Pilot findings reshape the charter. Operational incidents reshape the oracles. Evolution may send a product line back to discovery. What matters is that each transition has an artifact and an owner, and that the factory's definition of "accepted" is versioned rather than inferred.

The lifecycle also gives the factory's self-generated work a place to go. The fifteen tickets about improving the factory become a roadmap item with a stage, an owner, and an exit criterion instead of a Slack thread asking who is responsible. That is the difference between a factory that improves itself coherently and one that merely produces more work.

## Dark does not mean ownerless

The most instructive part of Debois's interview is the operational mess that factories expose.

At Tessl, the factory unexpectedly took over marketing and sales tickets as well as engineering work. It produced too many features for go-to-market to keep up. Maintainers found inconsistencies between classes written by different engineers and moved more of their taste into deterministic checks — linters, verifiers, tests — so agents could detect the problem without a human commenting on a ticket. And then the factory started filing tickets about improving itself, with no owner assigned.

Debois's conclusion is that ownership is not optional: "Can you do it without? Yes. But can you redo it without ownership? No." The factory needs an owner accountable for its yield, and the organization needs a deliberate decision to create that team. Otherwise, as he puts it, the tools get handed to engineering with the instruction to figure it out, and nobody owns the production function.

Two more warnings from the interview are worth keeping:

- **Architecture is underrated.** Debois recounts an engineer who claimed he never looked at code anymore, but who could still gate risky changes because the codebase was organized by domain boundaries. The factory could detect when a change touched auth or an API. Without that architecture, the safety story collapses.
- **Dark is a spectrum, not a switch.** Risk appetite, rollback speed, data-loss windows, least-privilege controls, and deployment gates determine how far a team can safely go. "Make it work, make it work reliably, and then optimize," Debois advises. Automating a bad loop just produces bad output at scale.

Factories amplify an organization's gaps. Missing ownership, inconsistent context, weak architecture, and undefined acceptance criteria do not disappear when agents arrive. They become production incidents.

## Build factories like products

The dark-factory debate has spent a year asking whether agents can produce software. The better question is what kind of production system we are building, and what it is allowed to produce.

The answer should be specific. A factory that produces CRUD admin surfaces is different from one that produces data pipelines, and both are different from one that produces safety-critical firmware. They need different specifications, different oracles, different harnesses, different economics, and different humans at the gates.

That is not a limitation of the idea. It is the idea. A factory is powerful because it is specialized: because it knows its product, it can verify its product, and because it can verify its product, it can run dark.

The next discipline after software engineering is not prompt engineering or loop engineering alone. It is **software factory engineering**: the design of specialized production systems that reliably turn specifications into a defined class of software artifact products. The key insight is that the factory is not a pipeline the agent passes through. It is the environment the agent inhabits. Engineering that environment — its constraints, feedback, tools, and oracles — is how we offload the hard parts of agent engineering and make dark factories reliable. The teams that learn to build those environments — narrow, instrumented, owned, and measured — will get the economics of the dark factory without the failure modes of the general-purpose dream.

---

**References and related reading:**

- [Is the Future of Software Engineering a Factory? | Patrick Debois](https://www.youtube.com/watch?v=FnCnv2tGsK0) — After the Last Engineer. The five-or-six-factories story, the self-filing factory, ownership, architecture, and the spectrum from assisted coding to lights-out.
- [Software dark factories: specs in, software out](https://blog.hackspree.com/#software-dark-factories) — the StrongDM case, Shapiro's five levels, and the shift from implementation to specification.
- [The economics of the dark factory: what happens when code is free](https://blog.hackspree.com/#dark-factory-economics) — the cost structure that makes specialized production functions worth building.
- [The Factory Is Not Dead](https://blog.hackspree.com/#factory-is-not-dead) — the industrial-factory lens, the smile curve, and Bemer's 1968 software factory.
- [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) — why the oracle, not the generator, determines what a factory can safely produce.
- Robert W. Bemer, *Machine-controlled production environment*, in the [NATO Software Engineering Conference report](https://www.cs.uni.edu/~wallingf/teaching/172/resources/nato1968.pdf) (Garmisch, 1968) — the software factory as an instrumented production system, not a metaphor.
- Stuart Russell and Peter Norvig, [*Artificial Intelligence: A Modern Approach*](https://aima.cs.berkeley.edu/) — the agent-environment framing that makes the environment half of the agent abstraction.
- Michael Wooldridge, *An Introduction to MultiAgent Systems* — agents as situated actors whose behavior is defined relative to their environment.
