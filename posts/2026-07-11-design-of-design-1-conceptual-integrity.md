---
title: "Brooks on design, part 1: conceptual integrity and the one-mind rule"
date: 2026-07-11
slug: brooks-design-conceptual-integrity
summary: "Fred Brooks' 'The Design of Design' argues that the most important property of any system is conceptual integrity — and that integrity requires a single mind, or at most two, controlling the design."
tags: design, fred-brooks, conceptual-integrity, software-architecture, collaboration
---

In 2010, thirty-five years after *The Mythical Man-Month*, Fred Brooks published *The Design of Design: Essays from a Computer Scientist*. It is a quieter book than its predecessor — less urgent, more reflective — and it was met with less fanfare. That is a mistake. It contains the most distilled thinking on what makes design work from someone who spent six decades designing across five media: computer architecture, software, houses, books, and organizations.

The book's central argument is that **conceptual integrity is the most important consideration in system design**. Everything else — process, tools, team structure, schedules — is in service of this property, or should be.

## What conceptual integrity is

Brooks defines conceptual integrity through three principles he first articulated in *The Mythical Man-Month*:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas."

The sentence is a thesis statement for his entire career. Note the tradeoff: omit useful features to preserve coherence. This is the hard part. Most designers agree with conceptual integrity in principle and violate it in practice the moment someone proposes a good feature that doesn't fit.

**Orthogonality.** Concepts should not overlap or conflict. Each function in the system should be achievable in exactly one way. If there are two ways to do something, users must learn both, must decide which to use, and will encounter inconsistencies between them. Orthogonality is not about minimalism — it is about non-redundancy. A design can be rich and still orthogonal.

**Propriety.** The design should include only what is necessary, and what is included should be transparent to the user. Propriety is parsimony with a purpose: every element earns its place. This is harder than it sounds. It requires saying no — repeatedly, to smart people with good arguments — and having the authority to make the no stick.

**Generality.** Within its chosen scope, the design should be complete. It should not have arbitrary limits that the user must work around. A design with generality handles the cases the designer anticipated and the ones they did not, because the primitives compose. Generality without orthogonality produces a kitchen-sink system. Orthogonality without generality produces a toy.

A system with all three properties has conceptual integrity. The user forms a coherent mental model. They can make predictions about how the system will behave in situations they have not yet encountered. The system feels like it came from one mind — because it did.

## The Reims Cathedral

Brooks uses architecture to make the point visual. The Reims Cathedral was built over eight generations of architects. The original architect drew the plan, and every successor — across centuries — stuck to it. The result is a unified work, coherent in every detail, that happens to have been built by many hands.

Most cathedrals are not like this. Most cathedrals have conflicting concepts: a Gothic nave with a Renaissance facade, a Baroque chapel grafted onto a Romanesque transept. Each generation of builders had a vision and nobody had the authority to enforce continuity. The result is architectural chaos — individually interesting, collectively incoherent.

Software systems are cathedrals built over decades. The systems with conceptual integrity — Unix, Lisp, Go, the original Macintosh — feel like Reims. The systems without it feel like the architectural equivalent of a patchwork quilt. Everyone who has maintained enterprise software knows the second kind.

## The one-mind rule

The mechanism that produces conceptual integrity is simple and unpopular: **the design must proceed from one mind, or from a very small number of agreeing resonant minds.**

Brooks is explicit about this. Great designs are attributed to individuals or pairs. Not committees. Not teams. Not "the community." The list of counterexamples — designs that emerged from large groups and achieved integrity — is vanishingly short.

> "The Design of Design sharpens the earlier contention: the design must represent the vision of one designer or, at most, a pair."

He doubles down. The 1975 position was that the *implementation* could be done by teams but the *architecture* needed one mind. The 2010 position eliminates even that qualification. If two people work in genuine resonance — finishing each other's thoughts, sharing a mental model so deeply they function as one — that can work. Three cannot.

This is not an argument against collaboration. Brooks devotes two chapters to collaboration and is careful to distinguish its uses. Teams are valuable for requirements elicitation — more people means more perspectives, more edge cases, more domain knowledge surfaced. Teams are valuable for exploring the design space — brainstorming, design competitions, alternatives analysis. Teams are valuable for implementation — decomposing the work, parallelizing effort, reviewing each other's code.

But the *conceptual design* — the core set of abstractions, the primitives and their relationships, the mental model presented to the user — must be controlled by one person. Or at most two, working in tight resonance.

The reason is structural. Every additional mind involved in the conceptual design introduces a new set of assumptions, preferences, and mental models. These must be reconciled. The reconciliation process produces compromises — and each compromise chips away at conceptual integrity. The result is a design that offends no one and satisfies no one. The committee didn't make the design better. It made it safer. And safety is the enemy of coherence.

## The cost of collaboration

Brooks quantifies this in terms that will be familiar to readers of *The Mythical Man-Month*. Adding people to a design effort has costs:

**Learning cost.** Each new person must acquire the shared vision. If the vision lives in one person's head, transferring it to *n* people takes *n × l* effort, not *l*. This is the teaching burden of the original designer, and it does not scale.

**Communication cost.** *n* people have *n(n−1)/2* communication paths between them. For a design team of 5, that is 10 paths. For 10 people, 45 paths. For 100 people, 4,950 paths. Each path is a channel through which misunderstandings flow and compromises are negotiated.

**Change control cost.** As the number of contributors grows, the cost of changing any design decision grows with it. More people must be consulted, more objections must be addressed, more downstream effects must be traced. The design calcifies.

This is the basis of Brooks's Law — "adding people to a late software project makes it later" — but applied to design rather than implementation. Adding designers to a design effort dilutes the design.

## The protection of the designer

If conceptual integrity requires one mind, and organizations contain many minds, the practical question becomes: how do you protect the designer?

Brooks's answer is rooted in his IBM experience. The System/360 architecture — one of the most influential computer architectures ever designed — was controlled by a small group with Brooks and Gene Amdahl at the center. They had the authority to say no. More importantly, they were *protected* from the organizational forces that would have diluted the design: feature requests from field sales, compatibility demands from existing customers, performance optimizations that would have compromised the clean abstraction.

This protection is not about giving designers autonomy. It is about giving them the *right kind* of constraint: clear overriding objectives, a schedule with urgency, and the freedom to make design decisions within those bounds. The worst outcome — and the most common — is a designer with no real authority, overseen by a committee that can override any decision but takes responsibility for none.

## What this means for teams today

The one-mind rule is uncomfortable in an industry that valorizes collaboration. But Brooks is not arguing that teams are bad. He is arguing that teams need a design authority, and that authority must be real.

The practical implications are clear. Every project of any complexity should have a single person — or a tight pair — who owns the conceptual design. That person reviews every interface, every abstraction, every user-visible decision. They have the authority to say no without escalation. They are not the team lead, not the engineering manager, not the product manager — though they may wear those hats too. They are the design owner, and their primary responsibility is conceptual integrity.

This role is hard to fill. It requires someone who can hold the entire system in their head, who has taste, who is willing to say no repeatedly to smart people, and who the organization trusts enough to give real authority. The industry does not train for this role. It does not reward it in hiring pipelines. It does not protect it in organizational structures. And then it wonders why most systems feel like patchwork quilts.

Brooks's argument, forty years after *The Mythical Man-Month* and sixteen years after *The Design of Design*, is still waiting to be taken seriously.

> "The building of a design, indeed, is the forcing of the will of one upon the stuff of the world."

A design is not a consensus. It is an imposition. The designer imposes coherence on a medium that has no opinion about coherence. This is uncomfortable language — "forcing," "will," "one upon the world" — and Brooks means it to be. Design is an act of authority.

---

**This is part 1 of a 3-part series on Fred Brooks' *The Design of Design*.**
- [Part 2: The rational model is wrong](/posts/brooks-design-rational-model)
- [Part 3: Growing great designers](/posts/brooks-design-great-designers)

**Reference:** Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
