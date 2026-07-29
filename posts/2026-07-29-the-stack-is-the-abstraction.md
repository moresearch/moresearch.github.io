---
title: The Stack as a Design Abstraction
date: 2026-07-29
slug: the-stack-is-the-abstraction
summary: Design is a search problem. The stack is the set of constraints that make the search tractable. This matters doubly in the Agentic Software Engineering era: AI coding agents are non-deterministic. The stack absorbs the variance — it constrains the agent's output to a predictable region where all paths lead to the same design. The stack turns non-deterministic agents into deterministic contributors.
tags: software-engineering, tech-stack, abstractions, architecture, ai-agents, design
---

Design is a search problem. The space of all possible systems is infinite. You cannot search it. You need constraints. Christopher Alexander called this the search for "good fit" between form and context — and argued, in *Notes on the Synthesis of Form*, that constraints are the mechanism that makes fit discoverable. Without them, the search never converges.

The stack is the constraint. A database narrows persistence. A framework narrows routing. A message queue narrows communication. Each layer says: the answer is in this region, not that one. PostgreSQL eliminates document stores. Rails eliminates manual HTTP handling. Temporal eliminates ad-hoc retry logic. The stack does not add options. It removes them. That is the point.

> The stack is not a collection of tools. It is a set of constraints on the design space. What remains is small enough to search.

Choose PostgreSQL, and you commit to schemas, migrations, ACID. Choose MongoDB, and you commit to documents and eventual consistency. This is not a performance decision. It is a search decision — which region of the design space will you explore? DHH understood this: convention over configuration is not about keystrokes. It is about eliminating search. Every Rails convention removes a decision. Rails doesn't make you faster. It makes the design space smaller. A good stack choice makes the region small enough to navigate and rich enough to contain the answer. A bad one eliminates the answer before you start looking.

A good stack makes intent manifest. Dijkstra argued that the intellectual manageability of software depends on levels of abstraction — each level a self-contained world the programmer can reason about without descending into the one below. Brooks called this *conceptual integrity* — one design voice, every part consistent. The stack provides both: the levels and the voice. Persistence is schemas and migrations. Routing is endpoints and handlers. Durable execution is workflows and steps. The designer thinks at the level of the problem. The stack translates. When the vocabulary is consistent, the design is coherent. When the design is coherent, the intent survives.

> The stack translates intent into code. The designer's intent is the input. The system is the output. The stack is the compiler.

The stack constrains technology. Patterns constrain structure. Buschmann, Meunier, Rohnert, Sommerlad, and Stal formalized this in *Pattern-Oriented Software Architecture* (1996): architectural patterns are not implementations. They are constraints on the arrangement of components. Layers says dependencies flow down. MVC says model, view, controller — separate them. Pipes and Filters says each step transforms data independently. Each pattern eliminates a region of the structural design space. You do not decide to use Layers and then decide how to organize dependencies. Layers *is* the decision. The pattern made it before you arrived.

Stack and patterns compose. PostgreSQL constrains what you store. Layers constrains how storage is accessed. Rails constrains how requests are handled. MVC constrains how responses are rendered. Each constraint narrows the search. Together they form a corridor through the design space — a path that a hundred designers would independently discover because the constraints make it the obvious path. This is what it means to find David in the marble. The constraints don't prevent the design. They reveal it.

> The stack constrains the materials. Patterns constrain the form. Together they make the search converge. Without them, every stone contains every statue. With them, the stone contains David — and the designer knows where to strike.

This matters doubly for AI coding agents. An agent does not search creatively. It searches the space you define. "Build me a backend" — it drowns. "Add an endpoint to this Rails controller that queries this PostgreSQL table and returns JSON" — it writes idiomatic code instantly. The difference is the stack. It turns the unbounded into the bounded.

> Coding agents excel at well-constrained problems. The stack is what constrains them. Without it, every answer is possible. With it, one answer is right.

This inverts the Unix philosophy. "Do one thing well" gives you thin tools and expects the programmer to compose them. The agent cannot compose — it searches. Hickey's distinction: Unix tools are simple but require the programmer to *complect* them into a system. Thick stacks are pre-composed. Thin abstractions — Express, raw SQL, manual deploys — give the agent a vast space with no idiomatic gradient. Thick abstractions — Rails, an ORM, a PaaS — give the agent a narrow space where competence is the only option. The question is answered before the agent writes a line: how much of the search space am I willing to eliminate in advance?

## Yes, but

Every strong thesis has a shadow. Here is mine.

**Too many layers.** Wheeler: every problem can be solved by another level of indirection. Henney: except too many levels of indirection. Abstractions are neither good nor bad — their quality depends on what they remove. A poor abstraction "leaves too much in or removes the wrong details." A stack can over-constrain. Each layer adds a surface to learn, a dependency to maintain, a convention that might conflict. Too little elimination leaves the search unbounded. Too much buries the answer under abstractions the problem didn't need. The designer's job: the minimum set of layers that makes the search tractable. Then stop.

**Lock-in is real. So is drift.** A thick stack constrains the search. When the constraint is wrong, it excludes the answer. I have watched teams contort a relational schema around document-shaped data because they picked PostgreSQL before they understood the problem. The contortion is expensive. The rewrite is more expensive. But the alternative — no constraint, no stack, just composable pieces — is worse. A team with no stack does not avoid lock-in. It drifts. Every developer picks their own abstraction. The codebase becomes a museum of decisions, each reasonable in isolation, none coherent together. At least lock-in is honest. Drift is lock-in you didn't choose. I'll take the former.

**Abstractions leak. Thin abstractions leak more.** Yes, the query planner will choose the wrong index. Yes, the framework will resist the feature that doesn't fit. But these failures are *legible*. The stack gave you a vocabulary for them. You know what an N+1 query is because the ORM named it. You know where to put the escape hatch because the framework told you where conventions live. Thin abstractions fail silently. Raw SQL doesn't tell you the query is slow — it just runs slow. Express doesn't tell you the middleware order is wrong — it just routes wrong. Thick abstractions fail with error messages. I'll take error messages over silence any day.

**Training bias is temporary.** Yes, an agent writes better Rails than it writes your custom ORM. That is a reason to use Rails, not a reason to avoid frameworks. The training distribution will shift. Agents will see more stacks, more conventions, more idioms. The gap between "thick enough to constrain" and "popular enough to be known" will widen. Betting that thin stacks will always produce better agent output because they're more common in training data is betting against the direction of the field. The field is moving toward thicker abstractions, not thinner ones. Rails won. Django won. Next.js is winning. The agent's training distribution is not static. Design for where it will be.

**Senior developers don't need guardrails.** That's true. But the system will outlast them. The senior who builds a bespoke Express architecture will leave, and the agent — or the junior — will inherit it. The stack you chose for yourself is the stack your successors will be constrained by, whether you intended it or not. A thick stack is an explicit constraint, documented and debated. A thin stack is an implicit constraint, embedded in the idiosyncratic architecture of the one person who understood it. I know which one I'd rather inherit. I know which one the agent would rather inherit. The agent cannot ask the senior why they chose this abstraction. It can only read the code. Give it code that speaks a language it knows.

> I am not arguing that thick stacks are always right. I am arguing that the burden of proof has shifted. The default should be a thick stack — a constrained search space, a coherent vocabulary, a surface the agent can read. The default should not be freedom. Freedom is what you reach for when the constraint has failed, not what you start with. Freedom is expensive. Make the stack pay for it.

## Determinism

There is one more reason the stack matters now, and it is the reason that makes everything else in this post urgent.

AI coding agents are non-deterministic. The same prompt produces different code each time. Without constraints, the variance is unbounded. An agent asked to build a backend might return Express one day, Fastify the next, raw Node the third. Each is correct. None is predictable. The agent's non-determinism becomes the system's instability.

The stack absorbs the variance. When the agent operates inside Rails, the answer space collapses. There are only so many ways to define a migration, only so many ways to write a controller. The agent may still vary — different variable names, slightly different query structures — but the variance is bounded by the stack's vocabulary. The code is different each time. The design is the same.

> The stack turns non-deterministic agents into deterministic contributors. The prompt may vary. The model may vary. The output stays within the lines. That is what a stack is for.

This is the defining problem of the Agentic Software Engineering era. We are introducing non-deterministic processes into a discipline that depends on predictability. Builds must be reproducible. Behavior must be testable. The system must behave the same way twice. An agent that writes different code each time threatens all three. The stack is the countermeasure. It says: the agent may choose any path through this space, but the space is small enough that all paths lead to the same design. The details differ. The architecture converges. The stack makes the agent's non-determinism safe.

> In the Agentic Software Engineering era, the stack is not a productivity tool. It is a stability guarantee. It ensures that an agent's output is not just correct, but coherent — not just working, but idiomatic — not just today, but tomorrow, when a different agent with a different model modifies the same codebase and must find code it can read.

Spolsky's [Development Abstraction Layer](https://www.joelonsoftware.com/2006/04/11/the-development-abstraction-layer-2/) argued management insulates programmers from everything not code. The stack insulates the design from everything not the problem — including the agent that builds it. Both succeed when invisible. Both fail when they demand attention. The designer's job is to make them disappear.

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
