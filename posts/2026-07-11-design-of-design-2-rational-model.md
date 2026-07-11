---
title: "Brooks on design, part 2: the rational model is wrong"
date: 2026-07-11
slug: brooks-design-rational-model
summary: "Fred Brooks argues that the rational model of design — requirements first, design second, implementation third — is fundamentally wrong. Design is empirical, iterative, and the process itself discovers what the client actually wants."
tags: design, fred-brooks, rational-model, empiricism, iteration, waterfall
---

If part 1 of Brooks's argument is about *who* should design — one mind, or at most two — part 2 is about *how* design actually works. And his conclusion is that the dominant model of design in software engineering is wrong.

The rational model — sometimes called the waterfall model in its software variant — says: gather requirements, produce a design, implement the design, test the implementation, ship. Each phase completes before the next begins. The design is rational in the sense that it proceeds logically from known premises to a correct conclusion.

Brooks calls it "wrong and harmful":

> "The Waterfall Model is wrong and harmful; we must outgrow it. What is wrong is that it is an essentially rational model, and for wicked problems, the rational model is simply the wrong model."

He doesn't hedge. Not "sometimes inappropriate." Not "useful but limited." *Wrong and harmful.* The waterfall model is not a simplification of how design works — it is a misrepresentation. Following it damages projects because it demands decisions at the point of maximum ignorance.

## The rational model's fatal flaw

The rational model makes one assumption that Brooks argues is demonstrably false: that the designer knows the goal at the start.

In Herbert Simon's formulation, design is systematic search through a combinatorial space. You have goals, utility functions, constraints, and resources. You search the space for a solution that satisfies the constraints and maximizes the utility. This is rational, systematic, and — if the space is tractable — optimal.

The problem is that in real design, **you do not know the goal**. The client does not know it either. Brooks captures this in the line that every designer recognizes: *"That's what I asked for, but that's not what I want."*

The client cannot articulate what they want until they see something. Once they see it, they can tell you what is wrong with it. But they cannot tell you what is right before they see it. This is not a failure of the client. It is a property of design problems: the requirements are discovered through the process of designing. The act of designing changes the designer's understanding of the problem, which changes the requirements, which changes the design.

This co-evolution of problem and solution is not a bug in the rational model. It is a fact about the world that the rational model cannot accommodate.

## The empiricist alternative

Brooks declares himself:

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

This is not the Brooks of *The Mythical Man-Month*. The younger Brooks believed in planning — the famous "plan to throw one away" was about doing the rational process twice. The older Brooks believes the first plan was never going to be right, no matter how carefully you made it. The shift from "plan better" to "build, test, learn" is the intellectual journey of the book.

This is the methodology that drove the scientific revolution, and Brooks argues it is the only methodology that works for complex design. The steps are:

1. **Understand the problem domain** — study the users, the context, the constraints
2. **Design something** — produce a candidate design, knowing it will be wrong
3. **Build an early prototype** — make it concrete enough to test
4. **Test it with real users** — watch what they do, not what they say
5. **Iterate** — use what you learned to improve the design
6. **Build incrementally** — grow the system in steps, testing at each step

This is not agile methodology, though it shares the same philosophical roots. It is a deeper claim about the nature of design knowledge. Design knowledge is empirical, not deductive. You cannot derive a good design from axioms. You must discover it through interaction with the problem, the users, and the constraints.

## Constraints are friends

One of Brooks's most counterintuitive insights: a problem with no constraints has no criteria for excellence. When anything is possible, nothing is good.

Constraints reduce the design space. They make the problem tractable. A designer without constraints is paralyzed by infinite possibility. A designer with clear constraints — budget, schedule, weight, power, compatibility, regulatory requirements — has a defined playing field. The creativity comes from finding an elegant solution within the constraints, not from ignoring them.

Brooks goes further:

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

This is a critique of clients — and internal stakeholders — who confuse requirements with implementation. "Use React" is not a requirement. "Must render at 60fps on mobile devices" is a requirement. "Must work offline" is a requirement. "Must be accessible to screen readers" is a requirement. The how is the designer's problem. The what is the client's — and the client often does not know the what until they see a candidate how.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

This is the designer's real job. Not translating requirements into blueprints. Helping the client discover what the requirements are. Every hour spent clarifying the problem saves ten hours of building the wrong solution. The tools change. This doesn't.

This connects back to the empiricist model. The process of proposing a design and having the client react to it is how the what gets discovered. The client says "that's not what I want" and the designer says "what specifically is wrong?" and the answer refines the requirements. This is not failure. It is the process working as designed.

## User models: better wrong than vague

Brooks makes a practical recommendation that sounds obvious and is almost never done: write down your explicit model of the user. Who are they? What do they know? What do they need? What are their constraints?

The model will be wrong. That is fine.

> "Better a precise model, even if wrong, than a vague one. A precise model exposes its assumptions and invites correction; a vague one is unfalsifiable and thus unhelpful."

A precise, wrong model surfaces its own assumptions. When you write "the user is a domain expert who uses the command line daily," everyone can see the assumption and challenge it. When you write nothing, everyone fills in their own implicit model and nobody realizes they disagree.

This is the same logic behind personas in UX design, but Brooks arrived at it from engineering rather than design research. The user model is a tool for making implicit assumptions explicit. Once explicit, they become testable. Once testable, they become correctable. This is the empiricist method applied to the most important unknown in any design: who is it for?

## How expert designers go wrong

One of Brooks's most striking observations is about expertise. Novices make technical mistakes — they forget edge cases, they choose the wrong algorithm, they build something that does not scale. These mistakes are easy to spot and easy to fix.

Experts make a different kind of mistake. They produce designs that are **comprehensively, systematically wrong**. The design is internally consistent, well-executed, and solves the wrong problem. Or it solves the right problem using assumptions from a previous era. Or it optimizes for a constraint that no longer exists.

Brooks's explanation is the paradigm trap. An expert designer has deep experience in one paradigm — one way of seeing problems and solutions. That depth is their strength. It is also their vulnerability.

> "The expert's very expertise becomes a liability when the paradigm shifts. The habits that served him well in one era mislead him systematically in the next. He is not making small errors; he is solving yesterday's problem with today's tools."

When the paradigm shifts — from mainframes to minicomputers, from monoliths to microservices, from deterministic to probabilistic systems — the expert's accumulated intuition becomes systematically misleading. They are not making small errors. They are solving yesterday's problem with today's tools and calling it design.

The empiricist method is the defense against this. An expert who tests their designs against reality — who prototypes, who watches users, who iterates based on evidence — will discover that their assumptions are wrong. An expert who designs from first principles and never tests will ship a beautiful, coherent, wrong system. The industry is full of examples of the second kind.

## The divorce of design

Brooks identifies a trend that has only accelerated since 2010: the separation of designers from implementers and implementers from users.

Brooks puts the point with characteristic bluntness:

> "The designer who does not build, and the builder who does not use, are both crippled."

This is not a suggestion about career development. It is a claim about epistemology. If you don't build, you don't know whether your design works. If you don't use, you don't know what "works" means. Each handoff in the chain — designer to implementer to user — is a site where knowledge is lost. The only fix is to collapse the chain into one person, or at least ensure that everyone in it has done every role at some point. When the designer is also the builder and the user, feedback is immediate and brutal. A bad design decision shows up in the workshop that afternoon.

Modern software development has broken this connection. Designers produce specifications. Implementers produce code. Users produce bug reports. Each handoff is a loss of information, a loss of accountability, and a loss of the rapid feedback that drives empirical improvement. The architect who never writes code designs abstractions that are elegant on paper and unbuildable in practice. The developer who never talks to users builds features that are technically impressive and functionally useless.

Brooks's prescription is to shorten the feedback loops. Designers should build. Builders should use. Users should be in the room. The organizational structures that prevent this — and there are many — are design problems in their own right.

---

**This is part 2 of a 3-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the one-mind rule](/posts/brooks-design-conceptual-integrity)
- [Part 3: Growing great designers](/posts/brooks-design-great-designers)

**Reference:** Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
