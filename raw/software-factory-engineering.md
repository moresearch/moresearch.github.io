---
title: Software Factory Engineering
date: 2026-10-17
slug: software-factory-engineering
summary: "Dark factories are not buildings where software happens; they are production functions. The discipline that designs those production functions — software factory engineering — is distinct from software engineering, and its first rule is that a factory must be specialized: it should produce a defined class of software artifact products, not software in general."
tags: software-factory, dark-factory, agents, software-engineering, factory-engineering, specialization, verification, economics
---

The dark factory argument has been won. The remaining argument is about what a factory is.

Over the past year the evidence has accumulated: agents can write code, review code, test code, and iterate on failures without a human in the loop. StrongDM reported running a three-engineer dark factory since mid-2025, with rules blunt enough to be a manifesto — code must not be written by humans, code must not be reviewed by humans. Dan Shapiro's five levels gave the industry a ladder, and level five gave it a destination: specs in, software out. The debate has moved from *whether* this works to *where* it works, *how well*, and *who is accountable when it does not*.

But "specs in, software out" hides the hard part. A specification is not a raw material that any factory can accept. Software is not a single product that any factory can emit. The thing between the input and the output — the production system itself — has to be designed, instrumented, and maintained. That is a different discipline from software engineering, and it is the discipline this post is about.

Call it **software factory engineering**: the engineering of the production system that turns specifications into accepted software artifacts.

A dark factory is not a room full of agents. It is a production function. Factory engineering is the work of defining that function precisely enough that it can be run, measured, and trusted.

## The production function, not the metaphor

Every factory, physical or software, has the same four parts:

- **Inputs.** In a software factory: specifications, domain knowledge, reference artifacts, fixtures, schemas, environments, credentials, model access, and the accumulated corpus of prior accepted work.
- **Transformation.** The agent roles, prompts, harnesses, tools, loops, review passes, test runs, and escalation paths that convert inputs into candidate artifacts.
- **Outputs.** Not "software" in the abstract, but specific artifacts with an expected shape: a service, a CLI, a migration, a Terraform module, a data pipeline, a test suite, a documentation set.
- **Control.** The verification gates, acceptance criteria, observability, yield accounting, and human review points that decide whether a candidate output becomes an accepted output.

Most dark factory discussion concentrates on the transformation — which model, which agent framework, which loop. That is the visible part, the factory floor. The leverage is elsewhere. The inputs determine whether the factory can start. The outputs determine what the factory is for. The control layer determines whether anyone can trust what comes out.

Factory engineering is the design of all four, together, as one system.

This is not a new idea wearing an AI hat. Robert W. Bemer's 1968 "machine-controlled production environment, or software factory" was already a production-system proposal: compilation, test-system construction, documentation, dependency graphs, quality control, instrumentation, scheduling, and costing. The software factory has always meant organized production, not merely automated typing. What changed is that the machinery can now do the parts that previously required human judgment at every step.

## The generality trap

The first instinct is to build a factory that builds anything. Give it a good enough specification, a general-purpose coding agent, and a test harness, and it should be able to produce any software artifact. This is the most natural idea in the world, and it is the wrong default.

A general-purpose software factory fails for the same reason a general-purpose physical factory fails. Changeover costs dominate. A factory that can stamp metal, injection-mold plastic, and bake bread is not three factories in one; it is a warehouse full of equipment that is never optimally configured for anything. Physical manufacturing solved this with specialization and product lines: a semiconductor fab, an injection-molding shop, a bakery. Each is optimized around a narrow class of product, its materials, its tolerances, its quality regime, and its economics.

Software appears to escape this constraint because the "material" is text and the "machinery" is a model. That appearance is misleading. The real constraints are not physical. They are:

**Specification closure.** A specification is executable only relative to a shared body of assumptions. A CRUD service and a flight-control loop do not share assumptions about failure, timing, state, or evidence. A general factory must either encode all of those assumptions at once — which makes the specification language enormous — or leave them implicit — which is exactly the ambiguity that dark factories cannot absorb.

**Verification oracles.** The factory must know whether an artifact is correct. For a REST service, the oracle is an API contract plus integration tests. For a data pipeline, it is reconciliation against source systems and invariants over outputs. For firmware, it is timing analysis, hardware-in-the-loop tests, and certification evidence. These oracles are not interchangeable. A general factory either carries every oracle or has no real oracle at all.

**Toolchain depth.** Deep automation comes from tight integration with a specific toolchain: the framework, the build system, the deployment target, the observability stack, the security model. General factories integrate shallowly with everything. Specialized factories integrate deeply with a few things, and depth is where reliability lives.

**Economics.** Building a factory has a fixed cost: specification schemas, golden artifacts, harnesses, verifiers, evaluation sets, and the human expertise to maintain them. That cost is amortized over repeated production. If the factory produces one bespoke system, it is a project with extra steps. If it produces a class of artifacts repeatedly, it is a production function with a real cost curve.

**Accountability.** "The factory built it" is not an answer to a regulator, an auditor, or an incident review. A specialized factory has a defined product, a defined acceptance regime, and therefore a defined answer to *what was supposed to happen* and *how we know it did*. A general factory has a story about agents.

None of this means a factory can only ever produce one artifact. It means a factory should be designed around a **product line**: a family of artifacts that share a shape, a toolchain, a verification regime, and a set of failure modes. The product line is the unit of specialization.

## A factory is defined by its product

The clearest way to say this is: **a software factory should be a specialized thing that produces certain types of software artifact products.**

Not "software." Not "features." Not "whatever the backlog says." Artifact products: things with a repeatable shape, a known interface, a defined quality bar, and an acceptance procedure.

Consider what counts as a product line:

- **Internal admin surfaces.** CRUD screens over a known data model, with standard auth, audit logging, pagination, and access control. Highly templated, strongly verifiable, high volume.
- **API integration adapters.** A connector to a third-party service with a defined contract, retry semantics, rate limits, and reconciliation tests. The product is the adapter, not the integration project.
- **Data pipelines.** Ingestion, transformation, and publication jobs with schema contracts, idempotency requirements, freshness SLAs, and reconciliation against sources.
- **Cloud infrastructure modules.** Terraform or Pulumi modules with policy-as-code checks, cost constraints, and a conformance suite. The artifact is a module, and its acceptance is a plan plus policy results.
- **Migrations.** Schema changes, data backfills, and cutover runbooks with reversibility and validation. The product line is the migration pattern, not the individual migration.
- **Developer tooling.** CLIs, linters, code generators, and test harnesses with a command contract, exit codes, and golden-output tests.
- **Documentation and reference material.** Generated from a source of truth, checked for link integrity, coverage, and consistency with the artifact it describes.
- **Evaluation and benchmark suites.** Tasks, graders, and result reports with reproducibility requirements and statistical acceptance criteria.
- **Static sites and content products.** A build pipeline, a rendering contract, accessibility checks, and link validation.
- **Regulated or safety-critical components.** The product line is narrower still, because the acceptance regime includes traceability, certification evidence, and independent review.

Each of these is a different factory. They share a philosophy and some infrastructure. They do not share a specification language, a verification oracle, or a definition of done.

The artifact product is the right unit because it is the smallest thing that can carry a stable contract. A "feature" is a slice of a product and changes shape with every roadmap revision. A "codebase" is the accumulated output of many product lines. An artifact product has a boundary: you can say what it is, what it must do, how to test it, and when it is done. That boundary is what makes lights-out operation possible.

## Why specialization is what makes darkness possible

A dark factory removes the human who used to absorb ambiguity. In a traditional team, an underspecified request is patched by judgment, context, and conversation. In a dark factory, the same ambiguity becomes a defect, a rework cycle, or a confident wrong answer. Specialization is how you reduce ambiguity to a manageable surface.

**Specialization makes the specification language finite.** If the factory only produces API adapters, the specification can be a schema: endpoint, auth method, retry policy, rate limits, data mapping, error taxonomy, reconciliation rules. The agent does not need to infer what kind of thing it is building. The space of valid inputs is bounded, and the factory can reject inputs that fall outside it.

**Specialization makes verification real.** For a known product line, you can build oracles that are stronger than "the tests pass." You can compare against a reference implementation, check invariants, run conformance suites, replay production traffic, and measure drift. A general factory can only run generic tests, which is why general factories tend to produce artifacts that pass generic tests and fail in specific ways.

**Specialization compounds toolchain knowledge.** Every product line accumulates its own harness: the right build commands, the right fixtures, the right local environment, the right deployment checks, the right failure playbook. That knowledge lives in the factory, not in a prompt. It is reusable, versioned, and testable. A general factory resets to general knowledge on every run.

**Specialization makes the economics work.** The fixed cost of a factory is paid once and amortized over the product line. A CRUD factory that produces fifty admin surfaces is a different economic object from a general factory that produces five unrelated systems. The first has a yield curve. The second has a series of custom projects with an agent in the middle.

**Specialization makes accountability concrete.** When a factory produces a known artifact, you can define what "accepted" means before production starts. You can audit the specification, the oracle, and the accepted artifact as a chain. You can answer the question that enterprise buyers and regulators actually ask: not "did an AI write this?" but "what evidence do you have that this is correct, and who owns the standard of correctness?"

This is the paradox of dark factories. The more autonomous the factory, the narrower it has to be. Autonomy is not generality. Autonomy is the removal of the human who used to fill gaps, and the only way to remove that human safely is to shrink the gaps.

## What factory engineering actually involves

If a factory is a product line plus a production system, then factory engineering has a recognizable set of deliverables:

**A product-line charter.** What artifact does this factory produce? What is in scope, what is explicitly out? What are the interfaces and quality attributes? What regulations or internal policies apply? The charter is the factory's constitution, and it should be short enough that every agent role can be evaluated against it.

**A specification schema.** The input language of the factory. It should be structured enough to validate, expressive enough to capture the product line's real decisions, and narrow enough that an agent cannot wander outside the product line without the factory noticing.

**A reference corpus.** Golden artifacts, exemplars, and anti-exemplars. The factory needs to know what good looks like and what failure looks like. This corpus is the factory's institutional memory, and it is more valuable than any individual prompt.

**A harness and toolchain.** The build, test, lint, security, and deployment machinery the agents operate. The harness is the factory floor. It should be versioned, reproducible, and owned by someone accountable for its reliability.

**Verification oracles.** The tests, invariants, reconciliations, policy checks, and review procedures that decide acceptance. The oracle is the factory's quality function. If the oracle is weak, the factory's yield number is fiction.

**Human gates.** Dark does not mean unattended. It means humans are not in the inner loop of production. Humans still own the charter, the schema, the oracle, exceptions, and the decision to change the product line. The gate should be placed where judgment is irreducible, not where it is merely habitual.

**Telemetry and yield accounting.** Accepted artifacts per unit time, first-pass yield, rework rate, defect escape rate, cost per accepted artifact, specification coverage, and changeover cost. A factory that cannot report these numbers is a demo, not a factory.

**Change management.** Models change, dependencies change, toolchains change, and the product line changes. The factory needs a way to detect drift, re-run its own conformance suite, and retire or upgrade components without silently changing the definition of accepted.

This is ordinary engineering discipline applied to an unusual production system. It is not glamorous. It is the difference between a factory and a very expensive autocomplete.

## The factory is a product, too

The deepest consequence of taking factory engineering seriously is that the factory itself becomes a product with a lifecycle.

It has a version. It has customers — the teams whose artifacts it produces. It has a roadmap, a support burden, and a deprecation path. It has an owner who is accountable for its yield. It has a boundary: the set of artifact products it is designed to produce, and the set it explicitly refuses.

That boundary is a feature. A factory that refuses work outside its product line is not inflexible; it is honest about where its verification is real. The alternative — accepting everything and hoping the general-purpose harness catches the difference — is how dark factories produce plausible garbage at scale.

Organizations that run multiple product lines will end up with a portfolio of factories. Some will share infrastructure: model access, sandboxing, artifact storage, telemetry, identity, and security review. That shared layer is a platform, not a factory. The distinction matters. A platform provides capabilities. A factory owns a production function. Confusing the two produces a platform team that thinks it is building factories and product teams that think they are using one.

The portfolio also creates a real decision: when does a new product line justify a new factory, and when should an existing factory be extended? The answer is not "when the backlog asks for it." It is when the new artifacts share enough of the specification language, the toolchain, the oracle, and the failure modes that the existing factory's verification remains meaningful. When those diverge, extending the old factory silently weakens its acceptance standard. That is the moment to build a new one.

## Build factories like products

The dark factory debate has spent a year asking whether agents can produce software. The more useful question is now: what kind of production system are we building, and what is it allowed to produce?

The answer should be specific. A factory that produces CRUD admin surfaces is a different thing from a factory that produces data pipelines, and both are different from a factory that produces safety-critical firmware. They need different specifications, different oracles, different harnesses, different economics, and different humans at the gates.

That is not a limitation of the idea. It is the idea. A factory is powerful precisely because it is specialized: because it knows its product, it can verify its product, and because it can verify its product, it can run dark.

The next discipline after software engineering is not prompt engineering, loop engineering, or harness engineering alone. It is software factory engineering: the design of specialized production systems that reliably turn specifications into a defined class of software artifact products. The teams that learn to build those factories — narrow, instrumented, accountable, and measured — will get the economics of the dark factory without the failure modes of the general-purpose dream.

---

**References and related reading:**

- [Software dark factories: specs in, software out](https://blog.hackspree.com/#software-dark-factories) — the StrongDM case, Shapiro's five levels, and the shift from implementation to specification.
- [The economics of the dark factory: what happens when code is free](https://blog.hackspree.com/#dark-factory-economics) — the cost structure that makes specialized production functions worth building.
- [The Factory Is Not Dead](https://blog.hackspree.com/#factory-is-not-dead) — the industrial-factory lens, the smile curve, and Bemer's 1968 software factory.
- [Loop Engineering is what the NATO conference asked for in 1968](https://blog.hackspree.com/#loop-engineering) — the loop as the unit of agentic production.
- [Harness Engineering (Martin Fowler)](https://blog.hackspree.com/#harness-engineering-fowler) and [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) — why the harness and the oracle determine what a factory can safely produce.
- Robert W. Bemer, *Machine-controlled production environment*, in the [NATO Software Engineering Conference report](https://www.cs.uni.edu/~wallingf/teaching/172/resources/nato1968.pdf) (Garmisch, 1968) — the software factory as an instrumented production system, not a metaphor.
