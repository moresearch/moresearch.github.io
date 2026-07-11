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

> "Requirements and design co-evolve. The act of designing changes the designer's understanding of the problem. As the design emerges, the requirements change. This is not failure; it is discovery."

Brooks calls this the co-evolution model. It is not that requirements exist fully formed in the client's mind, waiting to be extracted. The requirements are *produced* through the act of designing. Each iteration of the design teaches the client something about what they actually need. The process ends not when the design matches the requirements, but when the requirements and the design have co-evolved to a point of mutual stability.

This echoes Peter Naur's argument from the same period. In *Programming as Theory Building* (1985), Naur argued that a program is not its code but the theory its builders hold of the problem it solves. That theory is built through designing. It cannot be extracted upfront. Brooks's co-evolution model and Naur's theory-building are the same insight from different traditions: the real product of design is understanding, and understanding emerges through the act of designing.

## The empiricist alternative

Brooks declares himself:

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

This is not the Brooks of *The Mythical Man-Month*. The younger Brooks believed in planning — the famous "plan to throw one away" was about doing the rational process twice. The older Brooks believes the first plan was never going to be right, no matter how carefully you made it. The shift from "plan better" to "build, test, learn" is the intellectual journey of the book.

This is the methodology that drove the scientific revolution, and Brooks argues it is the only methodology that works for complex design. Compare this with Niklaus Wirth's *stepwise refinement* (1971) — the dominant methodology of the same era. Wirth argued you should decompose a problem into sub-problems, then refine each until it is trivial to code. Brooks's response: stepwise refinement assumes you know the decomposition upfront. You don't. The decomposition emerges through building and testing. Wirth's method works for well-understood problems. Brooks's empiricism works for everything else.

The steps are:

1. **Understand the problem domain** — study the users, the context, the constraints
2. **Design something** — produce a candidate design, knowing it will be wrong
3. **Build an early prototype** — make it concrete enough to test
4. **Test it with real users** — watch what they do, not what they say
5. **Iterate** — use what you learned to improve the design
6. **Build incrementally** — grow the system in steps, testing at each step

> "The prototype is the pivot of the design process. It makes ideas concrete and thereby falsifiable. A prototype that fails teaches more than a specification that pleases."

The prototype is not a draft of the final product. It is a question posed to reality: does this work? does this make sense? does this solve the problem? A failed prototype is not wasted effort — it is the fastest way to discover what the requirements actually are.

This is not agile methodology, though it shares the same philosophical roots. It is a deeper claim about the nature of design knowledge. Design knowledge is empirical, not deductive.

> "Formal methods — proving programs correct — represent rationalism's last stand in software. They work in principle for small, well-specified modules. They cannot scale to large, complex, evolving systems. No other design discipline even attempts formal correctness proofs. Architects do not prove buildings will stand; they build them and test them."

You cannot derive a good design from axioms. You must discover it through interaction with the problem, the users, and the constraints. This is true of buildings, bridges, airplanes, and software. The rationalist dream — design from first principles, correct by construction — is alive only in computer science departments. Every other design discipline abandoned it centuries ago.

This empiricist conviction runs through all of Brooks's work. In *No Silver Bullet* (1986), he argued there is no single breakthrough that will eliminate the essential difficulty of software design. In *Computer Architecture: Concepts and Evolution* (1997, with Gerrit Blaauw), he showed that even the most rigorously specified designs — computer instruction sets — evolved through trial, error, and market selection, not deduction from first principles.

## Constraints are friends

One of Brooks's most counterintuitive insights: a problem with no constraints has no criteria for excellence. When anything is possible, nothing is good.

Constraints reduce the design space. They make the problem tractable. A designer without constraints is paralyzed by infinite possibility. A designer with clear constraints — budget, schedule, weight, power, compatibility, regulatory requirements — has a defined playing field. The creativity comes from finding an elegant solution within the constraints, not from ignoring them.

This inverts the conventional wisdom of Brooks's era. Hoare famously warned that "premature optimization is the root of all evil" — a constraint (performance) should not dominate design decisions too early. Brooks goes further: constraints are not just necessary evils to be deferred. They are the conditions that make design possible at all. Without them, you have infinite search space and no way to evaluate a candidate. With them, the designer's job goes from impossible to merely hard.

Brooks goes further:

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

This is a critique of clients — and internal stakeholders — who confuse requirements with implementation. "Use React" is not a requirement. "Must render at 60fps on mobile devices" is a requirement. "Must work offline" is a requirement. "Must be accessible to screen readers" is a requirement. The how is the designer's problem. The what is the client's — and the client often does not know the what until they see a candidate how.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

This is the designer's real job. Not translating requirements into blueprints. Helping the client discover what the requirements are. Every hour spent clarifying the problem saves ten hours of building the wrong solution. The tools change. This doesn't.

This connects back to the empiricist model. The process of proposing a design and having the client react to it is how the what gets discovered. The client says "that's not what I want" and the designer says "what specifically is wrong?" and the answer refines the requirements.

> "Brooks tells of a military helicopter whose requirements committee, after months of careful work, added as a final requirement: 'It shall be capable of flying itself across the Atlantic.' This contradicted every prior constraint. But it was in the requirements document, so it had to be honored. The rational model has no defense against a late-breaking absurdity, because it treats all requirements as equally valid once documented."

This is not failure. It is the process working as designed — but only if the process allows requirements to be challenged. The rational model doesn't. Requirements are inputs. Design is output. The feedback loop is broken.

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

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
