---
title: Software dark factories: specs in, software out
date: 2026-07-07
slug: software-dark-factories
summary: "The dark factory model — where humans write specs and AI agents handle everything else — is not a thought experiment. StrongDM is already running one. Here's what that means for how we build software."
tags: ai, agents, software-engineering, dark-factory, automation
---

In the 1980s, the Japanese robotics company FANUC built a factory where robots manufactured other robots. No human workers. No lights — because nobody was there to need them. The machines just ran.

That image — a dark, humming factory floor producing goods around the clock with zero people in the loop — has haunted manufacturing ever since. Now it has arrived in software.

The term "dark factory" was adapted to software development by Dan Shapiro in January 2026, who laid out a five-level framework for AI-assisted coding. Level 0 is hand-written code. Level 5 is: **specs go in, software comes out.** Three weeks after Shapiro's post, StrongDM revealed they had been running a dark factory internally since mid-2025. This is not a thought experiment. It is running in production.

## The five levels to lights-out

Shapiro's framework maps cleanly to the self-driving car levels. It is worth walking through, because each step describes a working mode that exists today.

**Level 0 — Manual.** Hand-written code. AI is absent or an afterthought. This is fading fast even among skeptics.

**Level 1 — Task delegation.** AI handles discrete, on-command tasks: generate unit tests, scaffold a component, write a docstring. The human drives; AI is a tool.

**Level 2 — Pair programming.** Real-time collaboration between developer and AI. The human guides direction, AI generates code. Shapiro estimates roughly 90% of "AI-native" developers operate here. This is the current default.

**Level 3 — Code review.** The relationship inverts. AI authors the code; humans review diffs and approve PRs. The developer becomes a manager, not a maker. This is where the psychological shift happens — you stop thinking "I write code" and start thinking "I specify behavior, inspect output, and approve."

**Level 4 — Spec-driven development.** Humans write detailed specifications — behaviors, acceptance criteria, edge cases — and hand them to AI agents. Hours later, humans check outputs against specs and tests. The developer becomes a product manager. The unit of work is no longer a pull request; it is a specification document.

**Level 5 — The dark factory.** Specs go in. Software comes out. AI agents write the code, other AI agents review it, still others test it. Agents iterate on failures autonomously. The human role is exclusively defining *what* to build and *why*. The *how* is fully automated.

## StrongDM is already doing it

The jump from theory to practice came fast. Three weeks after Shapiro's post, StrongDM — an infrastructure access company — went public with an internal dark factory they had been running since mid-2025. The details are striking:

- **Team size:** three engineers.
- **Rules:** "Code must not be written by humans" and "Code must not be reviewed by humans."
- **Process:** Engineers write prose specifications covering edge cases, error handling, and acceptance criteria. AI agents generate the code. Other AI agents review it. Still others test it. Agents iterate on failures autonomously. Humans touch only the specification and validation layers.
- **Spend benchmark:** If you aren't spending at least $1,000 per engineer per day on AI compute, "you have room for improvement."

Simon Willison, co-creator of Django, visited the StrongDM team and described their approach as "very convincing." His takeaway: the critical investment was not in better AI models but in better *specifications* and test coverage. The quality of the output was a direct function of the quality of the input.

This is the inversion that matters. In a traditional team, you hire for coding skill and hope the specification thinking comes along with it. In a dark factory, specification thinking *is* the job. Coding is a downstream implementation detail handled by machines.

## The bottleneck moves upstream

For decades, the primary constraint on software output was typing speed — or more precisely, the rate at which a human could translate intent into code, handle edge cases, write tests, and iterate through review. Dark factories move the bottleneck from implementation to specification.

This has consequences:

**Vague requirements become instantly visible.** When a human team receives an underspecified ticket, they fill in the gaps with judgment, experience, and hallway conversations. When an AI agent receives an underspecified spec, it produces exactly what you asked for — and you discover at validation time that what you asked for was wrong. The feedback loop is shorter and more brutal. Ambiguity that would have been absorbed by a senior engineer's intuition now produces a broken build.

**Specification becomes a first-class engineering discipline.** Writing a spec that an AI can execute against is not the same as writing a Jira ticket. It requires defining behaviors, acceptance criteria, edge cases, and error handling with enough precision that a machine — with no context, no judgment, no hallway conversations — can produce working software. This is a skill. It can be learned. And in a dark factory world, it is the skill that determines output quality.

**Senior expertise concentrates differently.** Architecture, system design, security, UX — the knowledge that feeds into specifications — becomes more valuable, not less. What becomes obsolete is hand-coding CRUD endpoints or writing boilerplate authentication flows. The senior engineer stops being a high-throughput typist and becomes a high-precision spec writer and validator.

## The agency model flips

For software agencies, the dark factory model rewrites the business equation. The traditional model sells developer hours. More work means more billable hours — a linear relationship between output and headcount. In a dark factory, the constraint is not hours but clarity of thought. A small team running a dark factory pipeline can match the output of a much larger traditional team because the typing is free.

The new agency pitch becomes: you are not paying us to write code. You are paying us to define exactly what should be built, validate that it was built correctly, and own the outcome. The value is in the specification craftsmanship and the validation rigor, not in the keystrokes. This is a better business — higher margins, faster delivery, cleaner differentiation — but it requires agencies to sell something they are not used to selling: their thinking, not their typing.

## The risks are real

None of this is free of problems.

**Quality at scale is unproven.** StrongDM's experiment is one team, one domain, one set of constraints. Multiple analyses show AI-generated code carries higher defect rates than human-written code. Dark factories demand extraordinarily robust automated testing and validation to compensate for the absent human review layer. CodeRabbit and similar companies are building tools to address this gap, but the tooling is young.

**Specification debt replaces technical debt.** Flawed specs produce flawed outputs faster and more confidently than human teams would. The code "works" according to its tests, but the tests were generated from flawed specs. You can end up with a system that passes every automated check and is still wrong — at scale, at speed. Debugging *why* requires going back to the spec, which was written by a human who may or may not still be on the project.

**Trust and compliance are open questions.** For enterprise buyers concerned with security, compliance, and maintainability, "no human ever reviewed this code" is not a selling point. The dark factory model will need to build trust — through audit trails, verification tooling, and proven track records — before it is acceptable in regulated environments.

**The name might be too good.** "Dark factory" is visceral, memorable, and slightly ominous. That is part of its power. It is also part of the risk — the term invites reaction before understanding. Expect pushback from people who hear "no humans" and think "no accountability."

## What happens next

The trajectory is clear even if the timeline is not. AI models improve with each release. Agent tooling matures rapidly. The economics are compelling: StrongDM's three-person team plus high compute costs still dramatically undercuts the cost of a traditional team building the same output. And a large percentage of professional software development applies well-understood patterns to business problems — CRUD, auth, API plumbing, dashboard construction. This kind of work is factory-shaped whether we like the metaphor or not.

The open questions are about trust, quality, and adoption curves. Most organizations will spend years at Levels 3 and 4 before approaching true dark factory operations. The tooling needs to catch up. The specification craft needs to develop as a discipline. And the industry needs to figure out what "accountable" means when no human touched the code.

But the direction of travel is set. The dark factory is not science fiction. It is running right now, with real users, on real infrastructure. The only question is how fast it spreads.

---

**References:**

- Dan Shapiro, [The Five Levels: from spicy autocomplete to the software factory](https://danshapiro.com/blog/2026/01/the-five-levels-from-spicy-autocomplete-to-the-software-factory/) (January 2026)
- Dark Factory, [What is dark factory software development?](https://darkfactory.dev/blog/what-is-dark-factory-software-development)
- Simon Willison's commentary on the StrongDM dark factory approach (referenced in the Dark Factory post above)
- FANUC's lights-out factory — the manufacturing origin of the "dark factory" term


Systems design is the core engineering discipline. Every system — whether a dark factory, an agent governance framework, or a software architecture — involves the same set of decisions: what are the components? what are their interfaces? what changes do we hide? what stays stable? The engineer who can answer these questions can design any system. The domain provides the constraints. The principles provide the method.


> A dark factory is not a tool. It is a production function. The inputs are specifications. The outputs are software. The factory is the mechanism. The mechanism determines the economics.
