---
title: "Brooks on design, part 2: why one mind must rule the design"
date: 2026-07-11
slug: brooks-design-one-mind-rule
summary: "Fred Brooks argues that conceptual integrity requires a single mind — or at most a resonant pair — controlling the design. Committees produce compromises, not coherence."
tags: design, fred-brooks, conceptual-integrity, one-mind, collaboration
---

Part 1 established what conceptual integrity is and why it matters. This part is about the mechanism that produces it — and why that mechanism is so uncomfortable.

The rule is simple and unpopular: **the design must proceed from one mind, or from a very small number of agreeing resonant minds.**

## The case for one mind

Brooks is explicit. Great designs are attributed to individuals or pairs. Not committees. Not teams. Not "the community." The list of counterexamples — designs that emerged from large groups and achieved integrity — is, by his accounting, empty.

> "The Design of Design sharpens the earlier contention: the design must represent the vision of one designer or, at most, a pair."

In *The Mythical Man-Month* (1975), Brooks allowed that *implementation* could be done by teams if the *architecture* belonged to one mind. By 2010, he had sharpened even that. The architecture itself, he now argues, can have at most two authors — and only if they are genuinely resonant.

> "Two people can serve as one mind only if they are in genuine resonance — finishing each other's thoughts, sharing a mental model so deeply that each knows what the other would decide. Three cannot."

Resonance at this depth is rare. Brooks found it with Gerrit Blaauw on System/360. They shared a mental model so complete that either could speak for the architecture. Most designer pairs never achieve this. The default case is not a resonant pair but a committee — and committees produce what committees always produce.

## The committee problem

> "Design by committee produces designs that offend no one and satisfy no one. Each member's wish list is accommodated, each objection smoothed over, until the result is a feature-laden compromise lacking any coherent vision."

The mechanism is structural. Every additional mind brings new assumptions, preferences, and mental models. These must be reconciled. Reconciliation produces compromises. Each compromise chips away at conceptual integrity. The committee didn't make the design better. It made it safer. Safety is the enemy of coherence.

Brooks is careful to say what this is *not*. It is not an argument against collaboration in general. Teams are valuable for requirements elicitation — more perspectives surface more edge cases. Teams are valuable for exploring the design space — brainstorming, design competitions, alternatives analysis. Teams are valuable for implementation — decomposing work, parallelizing effort, reviewing code.

But the *conceptual design* — the core abstractions, the primitives and their relationships, the mental model presented to the user — must be controlled by one person. Or at most two.

## The cost of collaboration

Brooks quantifies this in terms familiar to readers of *The Mythical Man-Month*:

**Learning cost.** Each new person must acquire the shared vision. If the vision lives in one designer's head, transferring it to *n* people takes *n × l* effort — the teaching burden of the original designer. It does not scale. For a team of ten, the designer spends more time teaching than designing.

**Communication cost.** *n* people have *n(n−1)/2* communication paths. For five designers, ten paths. For ten, forty-five. For a hundred, 4,950. Each path is a channel through which misunderstandings flow and compromises are negotiated. The communication overhead grows quadratically while design throughput grows, at best, linearly.

**Change control cost.** As contributors multiply, changing any design decision becomes exponentially harder. More people to consult, more objections to address, more downstream effects to trace. The design calcifies. What began as a living vision becomes a frozen specification that nobody can change and nobody fully owns.

> "Adding manpower to a late software project makes it later." — *The Mythical Man-Month*, 1975

Brooks's Law was originally about scheduling — the n(n−1)/2 communication paths that make adding people counterproductive on a late project. But the same mathematics applies to design. Adding designers dilutes the design. Each additional designer is not just a new contributor. They are a new source of divergence from the conceptual center.

In *No Silver Bullet* (1986), Brooks extended this logic to the full software process. He distinguished *essential* complexity — inherent in the problem, no matter what tools you use — from *accidental* complexity — imposed by our tools and methods. The rational model treats all complexity as accidental, as if better process or more people could eliminate it. Brooks argued that essential complexity remains no matter what you do. Conceptual integrity is how you manage it: one voice in the design, even if a hundred hands build the implementation.

## The structural necessity

The one-mind rule is not a personality preference. It is a structural requirement. A design is a set of interdependent decisions. Each decision constrains the others. When different people make different decisions, the constraints conflict. The system acquires multiple personalities. The user suffers.

This is why Brooks insists that even the *idea* of "design by community" is incoherent. A community has no single vision. It has many visions, overlapping and conflicting. The output of a community process is not a design. It is a negotiated settlement. Settlements can govern societies. They cannot produce coherent software.

The practical implication: every project of any complexity needs a single person — or a resonant pair — who owns the conceptual design. That person reviews every interface, every abstraction, every user-visible decision. They have the authority to say no without escalation. They are the design owner. Their primary responsibility is conceptual integrity. Everything else follows from this.

Melvin Conway, Brooks's IBM contemporary, observed in 1968 that "organizations which design systems are constrained to produce designs which are copies of the communication structures of these organizations." Conway's Law is usually cited as a warning. Read alongside Brooks, it becomes something else: a proof that if you want a coherent design, you need a coherent design organization — which means one mind. A committee's communication structure is a complete graph. Its design output will be, too.

The next part takes up the practical question that follows: how do you protect that person from the organizational forces that will try to dilute their design?

---

**This is part 2 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the Reims Cathedral](/posts/brooks-design-conceptual-integrity)
- [Part 3: Protecting the designer](/posts/brooks-design-protecting-designer)
- [Part 4: The rational model is wrong](/posts/brooks-design-rational-model)
- [Part 5: The empiricist alternative](/posts/brooks-design-empiricist-alternative)
- [Part 6: How experts go wrong and the divorce of design](/posts/brooks-design-experts-divorce)
- [Part 7: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
