---
title: Software Factory Engineering
date: 2026-10-17
slug: software-factory-engineering
summary: "Dark factories are production functions, not rooms full of agents. Patrick Debois's account of building five or six factories at Tessal shows why: a factory has to be specialized around a class of software artifact products, owned like a product, and verified like a production line."
tags: software-factory, dark-factory, agents, software-engineering, factory-engineering, specialization, verification, patrick-debois
---

Ask whether software engineering is becoming a factory and the answer is probably yes. The more useful question is: *which* factory?

Patrick Debois — one of the people who gave the industry DevOps — has been building software factories at Tessal. In a recent interview, [*Is the Future of Software Engineering a Factory?*](https://www.youtube.com/watch?v=FnCnv2tGsK0), he describes what happened when the whole organization was told to build one: "we ended up with five or six different factories." Not because the team wanted duplication, but because the work was not the same. As he puts it, "not all projects are similar." The factories differed by repository shape, tooling, context, and workflow.

[![Patrick Debois interviewed about software factories: five or six factories, ownership, and what it takes to survive production](images/software-factory-engineering-video.jpg)](https://www.youtube.com/watch?v=FnCnv2tGsK0)

Then came the message that captures the whole problem. One morning, someone asked on Slack: the factory had created fifteen tickets about improving itself. *Who owns this?*

That question is the subject of this post. A dark factory is not a room full of agents producing "software." It is a production function, and production functions need a product, a boundary, an owner, and a verification regime. The discipline that designs them is **software factory engineering**.

## The factory is a production function

Every factory, physical or software, has four parts:

- **Inputs:** specifications, domain knowledge, reference artifacts, fixtures, environments, credentials, and model access.
- **Transformation:** agent roles, prompts, harnesses, tools, loops, review passes, and escalation paths.
- **Outputs:** not software in general, but artifacts with an expected shape — a service, a CLI, a migration, a Terraform module, a data pipeline, a test suite.
- **Control:** verification gates, acceptance criteria, observability, yield accounting, and human review points.

Most dark-factory discussion obsesses over the transformation — which model, which framework, which loop. That is the visible factory floor. The leverage is in the other three. Inputs determine whether the factory can start. Outputs determine what it is for. Control determines whether anyone can trust what comes out.

This is not a new idea. Robert W. Bemer's 1968 "machine-controlled production environment, or software factory" already included compilation, test-system construction, documentation, quality control, instrumentation, scheduling, and costing. The software factory has always meant organized production. What changed is that the machinery can now perform the steps that used to require human judgment at every turn.

## Specialization is the safety mechanism

The first instinct is to build a factory that builds anything. General-purpose coding agent, good specification, test harness — surely that covers the space. It does not, for the same reason a general-purpose physical factory does not: changeover costs dominate. A factory that can stamp metal, injection-mold plastic, and bake bread is not three factories in one. It is a warehouse full of equipment that is never optimally configured for anything.

Software looks exempt because the material is text and the machinery is a model. The real constraints are elsewhere:

**Specification closure.** A specification is executable only relative to shared assumptions. A CRUD service and a flight-control loop do not share assumptions about failure, timing, state, or evidence. A general factory must encode all of them at once — making the spec language enormous — or leave them implicit, which is the ambiguity a dark factory cannot absorb.

**Verification oracles.** The factory has to know whether an artifact is correct. For a REST service, that is an API contract plus integration tests. For a data pipeline, reconciliation against source systems. For firmware, timing analysis and hardware-in-the-loop tests. These oracles are not interchangeable. A general factory either carries every oracle or has no real oracle at all.

**Toolchain depth.** Deep automation comes from tight integration with a specific framework, build system, deployment target, observability stack, and security model. General factories integrate shallowly with everything. Specialized factories integrate deeply with a few things, and depth is where reliability lives.

**Economics.** Building a factory has a fixed cost: schemas, golden artifacts, harnesses, verifiers, eval sets, and the people who maintain them. That cost amortizes over repeated production. A factory that builds one bespoke system is a project with extra steps. A factory that builds a class of artifacts repeatedly is a production function with a real cost curve.

**Accountability.** "The factory built it" is not an answer to an auditor or an incident review. A specialized factory has a defined product and a defined acceptance regime, so it can answer *what was supposed to happen* and *how we know it did*. A general factory has a story about agents.

This is why Tessal ended up with five or six factories. Debois's lesson is not that duplication is good. It is that a factory is shaped by its product. He describes aiming for "two or three paved roads" — standard shapes the organization can support — with teams that leave the road taking on the security, observability, and maintenance burden themselves.

## The product defines the factory

The rule is simple: **a software factory should be a specialized thing that produces certain types of software artifact products.**

Not "software." Not "features." Not "whatever the backlog says." Artifact products: things with a repeatable shape, a known interface, a defined quality bar, and an acceptance procedure.

Examples of product lines:

- Internal admin surfaces over a known data model, with standard auth, audit logging, and access control.
- API integration adapters with defined contracts, retry semantics, rate limits, and reconciliation tests.
- Data pipelines with schema contracts, idempotency requirements, freshness SLAs, and reconciliation against sources.
- Cloud infrastructure modules with policy-as-code checks, cost constraints, and a conformance suite.
- Migrations with reversibility, backfill validation, and cutover runbooks.
- Developer CLIs and code generators with a command contract, exit codes, and golden-output tests.
- Documentation sets generated from a source of truth and checked for coverage and link integrity.
- Evaluation and benchmark suites with reproducible tasks, graders, and statistical acceptance criteria.
- Static sites and content products with a rendering contract, accessibility checks, and link validation.
- Regulated or safety-critical components, where the acceptance regime includes traceability, certification evidence, and independent review.

Each is a different factory. They may share infrastructure — model access, sandboxing, artifact storage, telemetry, identity — but they do not share a specification language, a verification oracle, or a definition of done.

The artifact product is the right unit because it can carry a stable contract. A "feature" changes shape with every roadmap revision. A "codebase" is the accumulated output of many product lines. An artifact product has a boundary: you can say what it is, what it must do, how to test it, and when it is done.

## What factory engineering builds

If a factory is a product line plus a production system, factory engineering has concrete deliverables:

- **A product-line charter:** what the factory produces, what is explicitly out of scope, and which quality and regulatory regimes apply.
- **A specification schema:** structured enough to validate, expressive enough to capture real decisions, narrow enough to reject work outside the product line.
- **A reference corpus:** golden artifacts, exemplars, and anti-exemplars — the factory's institutional memory.
- **A harness and toolchain:** the versioned build, test, lint, security, and deployment machinery the agents operate.
- **Verification oracles:** tests, invariants, reconciliations, policy checks, and review procedures that decide acceptance. If the oracle is weak, the yield number is fiction.
- **Human gates:** placed where judgment is irreducible, not where it is habitual.
- **Telemetry and yield accounting:** first-pass yield, rework rate, defect escape rate, cost per accepted artifact, specification coverage, changeover cost.
- **Change management:** a way to detect drift, re-run conformance, and upgrade models or dependencies without silently changing what "accepted" means.

This is ordinary engineering discipline applied to an unusual production system. It is the difference between a factory and a very expensive autocomplete.

## Dark does not mean ownerless

The most instructive part of Debois's interview is the operational mess that factories expose.

At Tessal, the factory unexpectedly took over marketing and sales tickets as well as engineering work. It produced too many features for go-to-market to keep up. Maintainers found inconsistencies between classes written by different engineers and moved more of their taste into deterministic checks — linters, verifiers, tests — so agents could detect the problem without a human comment on a ticket. And then the factory started filing tickets about improving itself, with no owner assigned.

Debois's conclusion is that ownership is not optional: "Can you do it without? Yes. But can you redo it without ownership? No." The factory needs an owner accountable for its yield, and the organization needs a deliberate decision to create that team. Otherwise, as he puts it, the tools get handed to engineering with the instruction to figure it out, and nobody owns the production function.

This is also where the line between platform and factory matters. A platform provides shared capabilities: sandboxes, proxies, a skills registry, observability, security scanning. A factory owns a production function. Confusing the two produces a platform team that thinks it is building factories and product teams that think they are using one. Debois describes the eventual convergence of the enablement team and the factory maintenance team into something like a platform team — but for the software factory, not for the cloud.

Two more warnings from the interview are worth keeping:

- **Architecture is underrated.** Debois recounts an engineer who claimed he never looked at code anymore, but who could still gate risky changes because the codebase was organized by domain boundaries. The factory could detect when a change touched auth or an API. Without that architecture, the safety story collapses.
- **Dark is a spectrum, not a switch.** Risk appetite, rollback speed, data-loss windows, least-privilege controls, and deployment gates determine how far a team can safely go. "Make it work, make it work reliably, and then optimize," Debois advises. Automating a bad loop just produces bad output at scale.

Factories amplify an organization's gaps. Missing ownership, inconsistent context, weak architecture, and undefined acceptance criteria do not disappear when agents arrive. They become production incidents.

## Build factories like products

The dark-factory debate has spent a year asking whether agents can produce software. The better question is what kind of production system we are building, and what it is allowed to produce.

The answer should be specific. A factory that produces CRUD admin surfaces is different from one that produces data pipelines, and both are different from one that produces safety-critical firmware. They need different specifications, different oracles, different harnesses, different economics, and different humans at the gates.

That is not a limitation of the idea. It is the idea. A factory is powerful because it is specialized: because it knows its product, it can verify its product, and because it can verify its product, it can run dark.

The next discipline after software engineering is not prompt engineering or loop engineering alone. It is **software factory engineering**: the design of specialized production systems that reliably turn specifications into a defined class of software artifact products. The teams that learn to build those factories — narrow, instrumented, owned, and measured — will get the economics of the dark factory without the failure modes of the general-purpose dream.

---

**References and related reading:**

- [Is the Future of Software Engineering a Factory? | Patrick Debois](https://www.youtube.com/watch?v=FnCnv2tGsK0) — After the Last Engineer. The five-or-six-factories story, the self-filing factory, ownership, architecture, and the spectrum from assisted coding to lights-out.
- [Software dark factories: specs in, software out](https://blog.hackspree.com/#software-dark-factories) — the StrongDM case, Shapiro's five levels, and the shift from implementation to specification.
- [The economics of the dark factory: what happens when code is free](https://blog.hackspree.com/#dark-factory-economics) — the cost structure that makes specialized production functions worth building.
- [The Factory Is Not Dead](https://blog.hackspree.com/#factory-is-not-dead) — the industrial-factory lens, the smile curve, and Bemer's 1968 software factory.
- [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) — why the oracle, not the generator, determines what a factory can safely produce.
- Robert W. Bemer, *Machine-controlled production environment*, in the [NATO Software Engineering Conference report](https://www.cs.uni.edu/~wallingf/teaching/172/resources/nato1968.pdf) (Garmisch, 1968) — the software factory as an instrumented production system, not a metaphor.
