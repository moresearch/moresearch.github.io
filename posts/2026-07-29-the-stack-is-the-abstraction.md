---
title: The Stack as a Design Abstraction
date: 2026-07-29
slug: the-stack-is-the-abstraction
summary: Design is a search problem. The stack constrains the search problem of design. Brooks called the designer's vision conceptual integrity — the ideal system in the mind's eye. The stack makes that vision manifest. This matters doubly in the Agentic Software Engineering era: AI coding agents are non-deterministic search engines. The stack does not make them deterministic — nothing can. It ensures all paths through the space converge on the designer's intent.
tags: software-engineering, tech-stack, abstractions, architecture, ai-agents, design
---

Design is a search problem. The space of all possible systems is infinite. You cannot search it. You need constraints. Christopher Alexander called this the search for "good fit" between form and context — and argued, in *Notes on the Synthesis of Form*, that constraints are the mechanism that makes fit discoverable. Without them, the search problem of design never converges.

The stack is the constraint. A database narrows persistence. A framework narrows routing. A message queue narrows communication. Each layer says: the answer is in this region, not that one. PostgreSQL eliminates document stores. Rails eliminates manual HTTP handling. Temporal eliminates ad-hoc retry logic. The stack does not add options. It removes them. That is the point.

> The stack is not a collection of tools. It is a set of constraints on the design space. What remains is small enough to search.

Choose PostgreSQL, and you commit to schemas, migrations, ACID. Choose MongoDB, and you commit to documents and eventual consistency. This is not a performance decision. It is a search decision — which region of the design space will you explore? DHH understood this: convention over configuration is not about keystrokes. It is about eliminating the design search. Every Rails convention removes a decision. Rails doesn't make you faster. It makes the design space smaller. A good stack choice makes the region small enough to navigate and rich enough to contain the answer. A bad one eliminates the answer before you start looking.

Beyond search, the stack makes intent manifest. Dijkstra argued the intellectual manageability of software depends on levels of abstraction — each level a self-contained world. Brooks called the result *conceptual integrity* — one design voice, every part consistent, the design existing in the mind's eye before code is written. The stack is what carries it from mind to machine. The designer thinks at the level of the problem. The stack translates. When the vocabulary is consistent, the design is coherent. When the design is coherent, the intent survives.

> The stack translates intent into code. The designer's intent is the input. The system is the output. The stack is the compiler.

The stack constrains technology. Patterns constrain structure. Brooks gave three principles for conceptual integrity, all negative: Propriety (no immaterial features), Orthogonality (no unnecessary coupling), Generality (no artificial limits). Buschmann et al. extended this to architecture in *Pattern-Oriented Software Architecture* (1996): Layers constrains dependencies, MVC constrains component roles, Pipes and Filters constrains data flow. Each pattern eliminates a region of the structural design space. You do not decide to use Layers and then decide how to organize dependencies. Layers *is* the decision.

Stack and patterns compose. PostgreSQL constrains what you store. Layers constrains how storage is accessed. Rails constrains how requests are handled. MVC constrains how responses are rendered. Each constraint narrows the design search. Together they form a corridor through the design space — a path a hundred designers would independently discover because the constraints make it obvious. Brooks called this the design in the mind's eye, made visible. The constraints don't prevent the design. They reveal the designer's intent.

> The stack constrains the materials. Patterns constrain the form. Together they converge the design search on the designer's intent. Without them, every stone contains every statue. With them, the stone contains David — and the designer sees him before the first strike.

This matters doubly for AI coding agents. An agent does not search creatively. It searches the space you define. "Build me a backend" — it drowns. "Add an endpoint to this Rails controller that queries this PostgreSQL table and returns JSON" — it writes idiomatic code instantly. The difference is the stack. It turns the unbounded into the bounded.

> Coding agents excel at well-constrained problems. The stack is what constrains them. Without it, every answer is possible. With it, one answer is right.

This inverts the Unix philosophy. "Do one thing well" gives you thin tools and expects the programmer to compose them. The agent cannot compose — it searches. Hickey's distinction: Unix tools are simple but require the programmer to *complect* them into a system. Thick stacks are pre-composed. Thin abstractions — Express, raw SQL, manual deploys — give the agent a vast space with no idiomatic gradient. Thick abstractions — Rails, an ORM, a PaaS — give the agent a narrow space where competence is the only option. The question is answered before the agent writes a line: how much of the search space am I willing to eliminate in advance?

## Yes, but

Every strong thesis has a shadow. Here is mine.

**Over-constraint and drift.** Wheeler: every problem can be solved by another level of indirection. Henney: except too many. A stack can over-constrain — each layer adds a surface to learn, a dependency to maintain. When the constraint is wrong, it excludes the answer. I have watched teams contort PostgreSQL schemas around document-shaped data. The contortion is expensive. But the alternative — no stack, no constraint — is worse. A team with no stack drifts. Every developer picks their own abstraction. The codebase becomes a museum of decisions, none coherent together. Lock-in is honest. Drift is lock-in you didn't choose. I'll take the former. The designer's job: the minimum set of layers that makes the design search tractable. Then stop.

**Abstractions leak. Thin abstractions leak more.** Yes, the query planner will choose the wrong index. Yes, the framework will resist the feature that doesn't fit. But these failures are *legible*. The stack gave you a vocabulary for them. You know what an N+1 query is because the ORM named it. You know where to put the escape hatch because the framework told you where conventions live. Thin abstractions fail silently. Raw SQL doesn't tell you the query is slow — it just runs slow. Express doesn't tell you the middleware order is wrong — it just routes wrong. Thick abstractions fail with error messages. I'll take error messages over silence any day.

**Training bias is temporary.** Yes, an agent writes better Rails than it writes your custom ORM. That is a reason to use Rails, not a reason to avoid frameworks. The training distribution will shift. Agents will see more stacks, more conventions, more idioms. The gap between "thick enough to constrain" and "popular enough to be known" will widen. Betting that thin stacks will always produce better agent output because they're more common in training data is betting against the direction of the field. The field is moving toward thicker abstractions, not thinner ones. Rails won. Django won. Next.js is winning. The agent's training distribution is not static. Design for where it will be.

**Senior developers don't need guardrails.** That's true. But the system will outlast them. The senior who builds a bespoke Express architecture will leave, and the agent — or the junior — will inherit it. The stack you chose for yourself is the stack your successors will be constrained by, whether you intended it or not. A thick stack is an explicit constraint, documented and debated. A thin stack is an implicit constraint, embedded in the idiosyncratic architecture of the one person who understood it. I know which one I'd rather inherit. I know which one the agent would rather inherit. The agent cannot ask the senior why they chose this abstraction. It can only read the code. Give it code that speaks a language it knows.

> Brooks said great designs come from great designers. He was talking about the designer's eye — the ability to see the system before it is built. The stack is how that eye reaches the code. The default should be a thick stack — a constrained space, a coherent vocabulary, a surface the agent can read. Freedom is what you reach for when the constraint has failed, not what you start with. Conceptual integrity is expensive. Make the stack pay for it.

## Determinism and search

AI coding agents are non-deterministic search engines. Each prompt launches a search through the design space. The same prompt, run twice, takes a different path. An unconstrained agent asked to build a backend searches a different region each time — Express today, Fastify tomorrow, raw Node the next. Each is a valid solution. None is predictable. The agent is searching correctly. The problem is not the agent. The problem is the space.

The stack solves this by constraining not the path, but the region. Inside Rails, the agent still searches non-deterministically. Different variable names. Different query ordering. Different middleware arrangement. What it cannot do is choose Express. The region boundary is the stack. The search is free within it. Every path stays within the Rails region. Every path realizes the designer's intent. The stack does not make the agent deterministic — nothing can. It makes the agent's non-determinism *harmless*.

> The agent searches. The stack bounds the design search. The path varies. The region is fixed. The designer's intent converges. Non-determinism becomes variation within a predictable envelope — today, and tomorrow, when a different agent searches the same codebase.

Spolsky's [Development Abstraction Layer](https://www.joelonsoftware.com/2006/04/11/the-development-abstraction-layer-2/) argued management insulates programmers from everything not code. The stack insulates the designer's intent — Brooks' vision in the mind's eye — from everything not the problem, including the agent that searches it. Both succeed when invisible. The designer's job is to make them disappear.

## Open questions

**How much of the problem must you understand before choosing the stack?** A stack constrains the design search. Choose too early, and the constraint eliminates the answer. Choose too late, and the codebase has already drifted into a museum of ad-hoc decisions. When is the right moment to commit?

**How do you test whether the stack still fits?** The problem evolves. The stack that constrained the design search perfectly at version 1 may exclude the answer at version 3. What is the signal that it's time to change the constraint — to swap a layer, adopt a new pattern, or remove an abstraction that no longer earns its place?

**Who decides the stack?** Brooks said design must come from one mind or very few. Does the stack decision require the same? Or is stack choice inherently a team decision — a social contract about which constraints everyone will accept — and therefore demands broad buy-in?

**How do you balance thickness against legibility for the agent?** A thick stack constrains the agent's search. But every layer is also a surface the agent must learn. At what point does thickness become opacity — the agent drowning not in possibility but in abstraction?

**What is the half-life of a stack decision?** Some constraints are permanent. Some expire when the team grows, the problem shifts, or the agent's training distribution changes. How do you know which is which? Which layers do you expect to replace, and which do you expect to outlast you?

**Can a stack be too thin for humans and too thick for agents simultaneously?** Human designers want freedom. Agents want constraints. The stack that satisfies both may not exist. Do you optimize for the human or the agent? Does the answer change over time as agents improve?

**What is the minimum viable stack?** The minimum set of layers that makes the design search tractable — and then stop. Every layer beyond the minimum is a tax. How do you know when you've reached the minimum? What is the test?

---

**References:**

- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press. — Design as search; satisficing as stopping condition.
- Christopher Alexander. (1964). *Notes on the Synthesis of Form*. Harvard University Press. — Design as the search for "good fit" between form and context; constraints as the mechanism that makes fit discoverable.
- Fred Brooks. (2010). *The Design of Design*. Addison-Wesley. — Conceptual integrity; the rational model vs. the empirical model; style as trained heuristic.
- David Heinemeier Hansson. [Convention over Configuration](https://en.wikipedia.org/wiki/Convention_over_configuration). Ruby on Rails. — Conventions as decisions removed from the design space.
- Joel Spolsky. (2006). [The Development Abstraction Layer](https://www.joelonsoftware.com/2006/04/11/the-development-abstraction-layer-2/). — The organization as platform; abstraction as invisibility.
- Rich Hickey. (2011). [Simple Made Easy](https://www.infoq.com/presentations/Simple-Made-Easy/). — Simplicity vs. complexity; complect vs. compose; why abstractions must not entangle.
- Edsger W. Dijkstra. (1972). [The Humble Programmer](https://www.cs.utexas.edu/~EWD/transcriptions/EWD03xx/EWD340.html). *Communications of the ACM*, 15(10). — Levels of abstraction as the intellectual manageability of software.
- Frank Buschmann, Regine Meunier, Hans Rohnert, Peter Sommerlad, Michael Stal. (1996). *Pattern-Oriented Software Architecture: A System of Patterns*. Wiley. — Architectural patterns as constraints on structural design; Layers, MVC, Pipes and Filters as search-space eliminators.
- Kevlin Henney. [From Mechanism to Method: Generic Decoupling](https://accu.org/journals/overload/12/60/henney_308/). *Overload*, 60. — Abstraction quality: good abstractions remove the right details; corollary to Wheeler's indirection aphorism.
- Related: [Finding David in the Marble](https://blog.hackspree.com/#finding-david-in-the-marble) — Design as search, constraints as preconditions.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment) — Why the stack must not surprise the agent.
