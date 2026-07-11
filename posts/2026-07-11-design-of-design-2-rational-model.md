---
title: "Brooks on design, part 2: the rational model is wrong"
date: 2026-07-11
slug: brooks-design-rational-model
summary: "Fred Brooks argues that the rational model of design — requirements first, design second, implementation third — is fundamentally wrong. Design is empirical, iterative, and the process itself discovers what the client actually wants."
tags: design, fred-brooks, rational-model, empiricism, iteration, waterfall
---

If part 1 of Brooks's argument is about *who* should design — one mind, or at most two — part 2 is about *how* design actually works. His conclusion is blunt: the dominant model of design in software engineering is not just a simplification. It is a misrepresentation.

The rational model — called the waterfall model in its software variant — says: gather requirements, produce a design, implement, test, ship. Each phase finishes before the next begins. The design proceeds logically from known premises to a correct conclusion. It is clean. It is orderly. It is wrong.

> "The Waterfall Model is wrong and harmful; we must outgrow it. What is wrong is that it is an essentially rational model, and for wicked problems, the rational model is simply the wrong model."

He doesn't hedge. Not "sometimes inappropriate." Not "useful in certain domains." *Wrong and harmful.* The problem is structural: the model demands decisions at the point of maximum ignorance — the beginning — and then forbids revisiting them. That is not how any complex design has ever succeeded.

## The fatal flaw

The rational model assumes the designer knows the goal at the start. This assumption is demonstrably false.

In Herbert Simon's classic formulation, design is systematic search through a combinatorial space. You have goals, utility functions, constraints, and resources. You search for a solution that satisfies the constraints and maximizes the utility. If the space is tractable, the result is optimal. The theory is beautiful.

The practice is that **nobody knows the goal** — not the designer, not the client. Brooks captures this in the line every experienced designer recognizes: *"That's what I asked for, but that's not what I want."* A client cannot articulate what they need until they see something. Once they see it, they can tell you what's wrong. But they cannot tell you what's right before anything exists. This is not a failure of communication. It is a property of design problems.

The act of designing changes the designer's understanding of the problem. The new understanding changes the requirements. The new requirements change the design. This is co-evolution, and it is not a bug in the rational model. It is a fact about reality that the model cannot accommodate.

> "Requirements and design co-evolve. The act of designing changes the designer's understanding of the problem. As the design emerges, the requirements change. This is not failure; it is discovery."

Brooks names this the co-evolution model. Requirements do not sit fully formed in the client's mind, waiting for extraction. They are *produced* through the act of designing. Each iteration teaches the client something about what they actually need. The process ends not when the design matches the requirements, but when the two have reached mutual stability — when further iteration produces diminishing returns on new understanding.

Peter Naur reached the same conclusion from a different direction. In *Programming as Theory Building* (1985), Naur argued that a program is not its code but the theory its builders hold of the problem it solves. That theory cannot be extracted upfront; it is built through the act of designing. Brooks's co-evolution and Naur's theory-building are one insight from two traditions: the real product of design is understanding, and understanding emerges through the work.

## The empiricist alternative

Brooks does not just critique. He names his position:

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

This is not the Brooks of *The Mythical Man-Month*. The younger Brooks believed in planning — the famous "plan to throw one away" was about doing the rational process twice, as if better requirements gathering would close the gap the first time. The older Brooks knows better. The first plan was never going to be right. No amount of upfront analysis would have made it right. The only path to a good design runs through being wrong, noticing, and trying again.

Compare this with the dominant methodology of Brooks's early career: Niklaus Wirth's *stepwise refinement* (1971). Wirth argued that you decompose a problem into sub-problems, then refine each until it is trivial to code. The method works beautifully — when you already understand the problem well enough to decompose it correctly on the first try. Brooks's response, implicit but clear: stepwise refinement is the rational model in miniature. It assumes the decomposition is known. It never is. The real decomposition emerges through building, testing, and discovering which boundaries were wrong.

The empiricist method has six steps:

1. **Understand the problem domain** — study the users, the context, the constraints
2. **Design something** — produce a candidate design, knowing it will be wrong
3. **Build an early prototype** — make it concrete enough to test
4. **Test it with real users** — watch what they do, not what they say
5. **Iterate** — use what you learned to improve the design
6. **Build incrementally** — grow the system in steps, testing at each step

> "The prototype is the pivot of the design process. It makes ideas concrete and thereby falsifiable. A prototype that fails teaches more than a specification that pleases."

The prototype is not a draft. It is a question posed to reality: does this work? does this make sense? does it solve the right problem? A failed prototype is not wasted effort — it is the fastest mechanism for learning what the requirements actually are. Specifications cannot do this because they are unfalsifiable. Only a running system can be wrong in a way that teaches something.

This is not agile methodology, though the philosophical roots overlap. It is a deeper claim about the nature of design knowledge. Design knowledge is empirical, not deductive.

> "Formal methods — proving programs correct — represent rationalism's last stand in software. They work in principle for small, well-specified modules. They cannot scale to large, complex, evolving systems. No other design discipline even attempts formal correctness proofs. Architects do not prove buildings will stand; they build them and test them."

You cannot derive a good design from axioms. You must discover it through interaction with the problem, the users, and the constraints. This is true of buildings, bridges, airplanes, and software. The rationalist dream — design from first principles, correct by construction — survives only in computer science departments. Every other design discipline abandoned it centuries ago. Brooks's entire body of work is an argument that software should join them.

In *No Silver Bullet* (1986), he argued there is no single breakthrough that will eliminate the essential difficulty of software design. In *Computer Architecture: Concepts and Evolution* (1997, with Gerrit Blaauw), he showed that even computer instruction sets — among the most rigorously specified artifacts ever designed — evolved through trial, error, and market selection, not deduction from first principles. The empiricist conviction is not a late-career softening. It is the thread that runs through everything he wrote.

## Constraints are friends

One of Brooks's most counterintuitive insights: a problem with no constraints has no criteria for excellence. When anything is possible, nothing is good.

Constraints reduce the search space. They make the problem tractable. A designer without constraints faces infinite possibility and is paralyzed by it. A designer with clear constraints — budget, schedule, weight, power, compatibility, regulatory requirements — has a defined field. The creativity lies in finding an elegant solution within the boundaries, not in pretending the boundaries do not exist.

This inverts the conventional wisdom of Brooks's era. C.A.R. Hoare famously warned that "premature optimization is the root of all evil" — one constraint, performance, should not dominate design too early. Brooks goes further. Constraints are not deferred evils; they are the conditions that make design possible. Without them, you have infinite search space and no way to distinguish good from bad. With them, the designer's job goes from impossible to merely hard.

Brooks draws a practical line:

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

This is aimed at clients and stakeholders who confuse requirements with implementation. "Use React" is not a requirement. "Renders at 60fps on mobile devices" is. "Works offline" is. "Accessible to screen readers" is. The how is the designer's problem. The what is the client's — and the client often does not know the what until they see a candidate how. The process of proposing a design and having the client react is how the what gets discovered. The client says "that's not what I want," the designer says "what specifically is wrong," and the answer refines the requirements. This is not failure. It is the process working.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

This is the designer's real job — not translating requirements into blueprints, but helping the client discover what the requirements are. Every hour spent clarifying the problem saves ten hours of building the wrong solution. Tools change. This does not.

Brooks illustrates the alternative with a cautionary tale. A military helicopter project had spent months carefully negotiating requirements. At the final meeting, someone added: "It shall be capable of flying itself across the Atlantic." This contradicted every prior constraint — weight, range, cost, everything. But it was now in the document. The rational model treats all documented requirements as equally valid. It has no mechanism for saying "this requirement, despite being written down, is absurd." The feedback loop is broken. The helicopter was never built.

## User models: better wrong than vague

One of Brooks's most practical recommendations sounds obvious and is almost never done: write down your explicit model of the user. Who are they? What do they know? What do they need? What are their constraints?

The model will be wrong. That is fine.

> "Better a precise model, even if wrong, than a vague one. A precise model exposes its assumptions and invites correction; a vague one is unfalsifiable and thus unhelpful."

When you write "the user is a domain expert who uses the command line daily," everyone can see the assumption and challenge it. When you write nothing, everyone fills in their own implicit model and nobody realizes they disagree. The user model makes assumptions explicit. Once explicit, they become testable. Once testable, they become correctable. This is the empiricist method applied to the most important unknown in any design: who it is for. UX designers arrived at personas independently. Brooks got there from engineering. Same logic, different path.

## How expert designers go wrong

Novices make technical mistakes — they forget edge cases, choose the wrong algorithm, build something that doesn't scale. These are easy to spot and easy to fix.

Experts make a different kind of mistake. They produce designs that are **comprehensively, systematically wrong**. The design is internally consistent, well-executed, and solves the wrong problem. Or it solves the right problem using assumptions from a previous era. Or it optimizes for a constraint that no longer exists.

Brooks's explanation is the paradigm trap. An expert has deep experience in one paradigm — one way of seeing problems and solutions. That depth is their strength.

> "The expert's very expertise becomes a liability when the paradigm shifts. The habits that served him well in one era mislead him systematically in the next. He is not making small errors; he is solving yesterday's problem with today's tools."

When the paradigm shifts — mainframes to minicomputers, monoliths to microservices, deterministic to probabilistic systems — the expert's intuition becomes systematically misleading. They are solving yesterday's problem with today's tools and calling it design. And because the design is internally coherent, nobody spots the error until it is too late.

The defense is the empiricist method. An expert who tests designs against reality — who prototypes, watches users, iterates on evidence — will discover that their assumptions are wrong. An expert who designs from first principles and never tests will ship a beautiful, coherent, wrong system. The industry is full of both kinds. You can tell which is which by asking: when did you last change your mind about a design decision because of something a user did?

## The divorce of design

Brooks identifies a trend that has only accelerated since 2010: the separation of designers from implementers, and implementers from users.

> "The designer who does not build, and the builder who does not use, are both crippled."

This is not career advice. It is epistemology. If you don't build, you don't know whether your design works. If you don't use, you don't know what "works" means. Each handoff — designer to implementer to user — is a site where knowledge is lost. When the designer is also the builder and the user, feedback is immediate and brutal. A bad design decision shows up in the workshop that afternoon.

Modern software has broken this connection. Designers produce specifications. Implementers produce code. Users produce bug reports. Each handoff loses information, accountability, and the rapid feedback that drives empirical improvement. The architect who never writes code designs abstractions that are elegant on paper and unbuildable in practice. The developer who never talks to users builds features that are technically impressive and functionally useless.

Brooks's prescription is simple: shorten the feedback loops. Designers should build. Builders should use. Users should be in the room. The organizational structures that prevent this are design problems in their own right — and they are the hardest kind, because the people who need to solve them are the people the structures benefit.

---

**This is part 2 of a 3-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the one-mind rule](/posts/brooks-design-conceptual-integrity)
- [Part 3: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
