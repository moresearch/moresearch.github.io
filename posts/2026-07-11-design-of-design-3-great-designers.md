---
title: "Brooks on design, part 3: growing great designers"
date: 2026-07-11
slug: brooks-design-great-designers
summary: "Brooks argues that great designs come from great designers, not great processes — and that organizations must deliberately grow, protect, and retain design talent rather than smothering it in process."
tags: design, fred-brooks, great-designers, engineering-culture, mentorship
---

The first two parts of this series covered Brooks's arguments about conceptual integrity and the empiricist model of design. This third part addresses the uncomfortable conclusion that follows from both: **great designs come from great designers, not from great processes.**

This is an unfashionable claim. The modern engineering organization invests heavily in process: code review, design review, architecture review, RFCs, ADRs, planning poker, retrospectives. These are not bad practices. But Brooks argues they are supporting structures, not substitutes. Process can prevent bad design from shipping. It cannot produce good design. For that, you need individual brilliance — recognized, nurtured, and protected.

> "Great designs come from great designers. Process can make a good design better; it cannot make a bad designer good."

The two sentences are in tension with everything the industry has built since 2010. We have invested in process — agile, code review, CI/CD, architecture review boards — as if better process produces better design. Brooks says it doesn't. Process raises the floor. Talent raises the ceiling. Conflating the two is the central mistake of engineering management.

## Process is not a substitute for talent

Brooks's position is blunt:

> "The trick is to hold process off long enough to permit great design to occur, so that the lesser issues can be debated once the great design is on the table — rather than smothering it in the cradle."

The sequence matters. If you convene a design review committee before a strong design exists, the committee will produce a compromise — a design that offends no one and satisfies no one. If you let a strong designer produce a coherent design first, and *then* subject it to review, the review can identify weaknesses, surface missed requirements, and test assumptions without diluting the core vision.

This is not a call to abolish process. It is a call to sequence process correctly: **design first, review second**. Most organizations do the reverse. They start with a committee, produce a compromise, and then hand it to an individual to "execute" — at which point conceptual integrity is already lost.

The practical implication is that organizations need to identify their best designers and give them space to design before the process machinery engages. This requires trust. It requires management that can distinguish "this person needs more process" from "this person has a track record of producing coherent designs and should be given room." And it requires the organizational discipline to protect the designer from the forces that will try to pull the design apart.

## The dual ladder

Brooks was one of the early advocates for the dual ladder — the idea that technical and managerial career paths should be parallel, equally compensated, and equally respected. In *The Design of Design*, he returns to this theme with forty more years of evidence.

The problem is well-known: organizations promote their best technical people into management. The best engineer becomes the engineering manager. The best architect becomes the head of architecture. This removes them from hands-on design work — the thing they were best at — and puts them in a role they may be mediocre at. The organization loses a great designer and gains a mediocre manager.

The dual ladder is the solution in principle. In practice, Brooks observes, it almost never works as intended.

> "The dual ladder is everywhere espoused and nowhere practiced. The managerial ladder remains the path to power, prestige, and pay. The technical ladder is too often a consolation prize — a place to park brilliant people who won't or can't manage."

The managerial ladder is always more prestigious, better compensated, and seen as the "real" path to influence. The technical ladder becomes a consolation prize for people who cannot or will not manage.

Brooks's advice for making the dual ladder real:

- **Compensate equally.** A Distinguished Engineer should make what a VP makes. If the numbers are different, the ladder is not real.
- **Make it visible.** Senior technical people should be in the room for strategic decisions. Their opinion should carry weight equal to senior managers. If they are not in the room, the ladder is not real.
- **Protect it from erosion.** When budgets tighten, the technical ladder should not be the first thing cut. If it is, the ladder is theater.
- **Staff it with exemplars.** The people on the technical ladder should be visibly great. If the Distinguished Engineer title goes to people who are merely tenured rather than brilliant, the ladder loses credibility.

## Growing designers deliberately

Brooks argues that design talent is not purely innate. It can be grown, and organizations that fail to grow it deliberately will be limited by whatever talent walks in the door.

His prescriptions:

**Recruit for design sense, not just interview performance.** The standard technical interview measures algorithm knowledge, coding speed, and system design on a whiteboard. None of these measure design taste.

> "Hire for demonstrated design ability, not for interview prowess. Look at what the candidate has built. Does it show conceptual integrity? Do the abstractions make sense? Would you want to use it? Past performance in design is the best predictor of future performance."

Brooks recommends looking at past work: what has the candidate actually designed? Does it show conceptual integrity? Do the abstractions make sense? Would you use it?

**Provide mentors.** Great designers learn from other great designers.

> "Design judgment is not taught; it is caught. The master-apprentice relationship — explicit in architecture, mostly informal in software — is how taste transmits. A junior working under a senior absorbs not just technique but judgment: what to optimize for, what to ignore, when to fight and when to concede."

The master-apprentice relationship — explicit in architecture, implicit in software — is how design judgment gets transmitted. A junior designer working under a senior one absorbs not just knowledge but taste: what to optimize for, what to ignore, when to fight and when to concede.

**Rotate experiences.** A designer who has only built one kind of system in one domain with one set of constraints has a narrow foundation.

> "The designer who has worked in only one domain, with one set of constraints, has a narrow base of pattern recognition. Breadth of experience — across domains, technologies, and constraints — builds the repertoire from which great designs are synthesized."

Rotating across domains — from infrastructure to product, from backend to frontend, from systems to tools — builds the pattern recognition that distinguishes good designers from great ones. Brooks himself designed in five media; he didn't think it was a coincidence.

**Protect them from management.** This is the hardest one. The reward for being a great designer is usually more meetings, more management responsibility, and less time designing. Breaking this cycle requires active intervention: someone in the organization with the authority to say "no, this person's highest-value activity is designing, not managing."

> "The proper office of the manager is to protect his great designers from managing."

The line is vintage Brooks — aphoristic, hierarchical, and correct. A great designer's highest-value activity is designing. Every hour they spend in a status meeting is an hour the organization loses their best work. The manager's job is not to manage the designer. It is to manage *everything else* so the designer can design.

## Esthetics and style

Brooks includes a chapter on esthetics in technical design — a topic that makes many engineers uncomfortable. But his argument is practical: even invisible designs have an aesthetic dimension.

> "We speak of 'clean' machines, 'elegant' languages, 'beautiful' proofs. These are not mere metaphors. They denote a real property: the fitness, coherence, and economy of a design, as judged by those with taste."

We talk about "clean" architectures, "elegant" APIs, "beautiful" abstractions. These are not metaphors. They are judgments about the coherence and fitness of a design, and they are made by people with taste.

His prescription for developing taste:

1. **Study other designers' styles intentionally.** Read code by people known for design quality. Study APIs that are widely admired. Ask: what makes this good? What decisions did the designer make? What did they choose *not* to include?
2. **Practice in another's style.** Write code in the style of a designer you admire. This forces you to make explicit what you would otherwise absorb implicitly. You learn more about a style by attempting to reproduce it than by passively appreciating it.
3. **Revise for stylistic consistency.** A design that is partly elegant and partly clumsy is worse than a design that is consistently adequate. The inconsistency distracts and confuses. Go back through your designs looking specifically for elements that do not belong.
4. **Choose designers with demonstrated good taste.** When building a design team, taste matters as much as technical skill. A team of brilliant implementers with no taste will produce a system that works correctly and feels wrong.

## The exemplar gap

One of Brooks's sharpest criticisms of software engineering as a discipline is that we do not study exemplars.

> "Architects study buildings; composers study scores; writers study books; painters visit museums. Software designers too rarely study existing software designs, and when they do, it is to learn how to use the system, not to understand why it was designed as it was."

Architects study buildings. Composers study scores. Writers study books. Software designers... write software. They rarely study existing systems in depth, and when they do, it is usually to understand how to use them, not why they were designed the way they were.

This is a loss. Great designs contain lessons that cannot be conveyed in principles or patterns — only in the specific sequence of decisions that produced a specific result. Reading the source code of the Unix kernel, studying the architecture of TeX, understanding why Go's concurrency model works the way it does — these are the software equivalents of an architect studying the Pantheon.

Brooks argues that this gap is partly cultural and partly practical. Culturally, software has a bias toward building over studying. Practically, software designs are less accessible than buildings — you cannot walk through a codebase the way you can walk through a cathedral. But the gap is real, and it limits the field's ability to accumulate design knowledge.

## The case studies

*The Design of Design* is structured around case studies, and they are worth reading in full. Brooks draws on:

- **IBM System/360** — the architecture that established the modern mainframe. Brooks was the project manager, and his account of the design decisions is a master class in making irreversible decisions under uncertainty.

> "Why an 8-bit byte? Why 32-bit words? Why a clean separation between architecture and implementation? Each of these decisions was made by a small group, debated intensely, and then locked. Once locked, they could not be reopened. This discipline — decide, commit, don't revisit — is the only way to ship a coherent architecture."

The decisions — 8-bit bytes, 32-bit words, the architecture/implementation split — shaped computing for decades. They were not obviously correct at the time. They became correct because a small team with real authority made them and refused to reopen them.
- **IBM OS/360** — the operating system that taught Brooks about the second-system effect and the limits of coordination. He is candid about its failures:

> "The job control language JCL for OS/360 is, in my opinion, the worst computer language ever devised, a triumph of committee design over conceptual integrity."

The second-system effect — the tendency to over-engineer the successor to a successful first system — is one of Brooks's most enduring contributions.

> "The second is the most dangerous system a designer ever builds. Having succeeded with the first, he loads the second with every feature he omitted from the first, every embellishment suggested by users, every capability the hardware now permits. The result is a bloated, over-budget, late disaster."

OS/360, with its thousand-man team and its committee-designed interfaces, was his cautionary example. JCL was the scar he carried from that project.
- **His beach house** — designed by Brooks himself, an amateur architect. The house — which he calls "View/360" in a nod to IBM — is a case study in how a designer works in an unfamiliar medium, and how the principles of conceptual integrity, empiricism, and iteration transfer across domains.
- **His kitchen renovation** — a smaller case study that illustrates the difference between designing for yourself (where you are the user, the client, and the builder) and designing for others.
- **A book design** — Brooks co-authored *Computer Architecture: Concepts and Evolution* and designed its structure. The case study shows how book design — chapter organization, cross-referencing, indexing — follows the same principles as software design.

The case studies are uneven — the house and kitchen chapters feel self-indulgent to some readers — but they serve Brooks's purpose. He is demonstrating, not just asserting, that design principles are invariant across media. The same lessons that produced System/360 also produce a good beach house.

> "I have designed in five media: computer architecture, software, houses, books, and organizations. This book draws lessons from all five. The principles of design are independent of the medium."

This is the book's methodological bet. If you can show that the same principles that govern a computer architecture also govern a kitchen renovation, you have shown something deeper than a software engineering methodology. You have shown a theory of design.

## The argument, taken seriously

*The Design of Design* is not a book you read for techniques. It is a book you read for perspective. Brooks's central claims — conceptual integrity requires one mind, the rational model is wrong, great designs come from great designers — are simple enough to state in a sentence each. The book earns its length by working through the implications of each claim across decades of practice.

The question the book leaves you with is whether your organization takes these claims seriously. Do you have a single person with real design authority? Do you treat design as an empirical process of discovery, or a rational process of derivation? Do you deliberately grow designers, or do you hire for coding speed and hope design sense comes along?

Most organizations answer no to all three. Brooks spent six decades explaining why that is a mistake. The book is still waiting for the industry to catch up.

---

**This is part 3 of a 3-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the one-mind rule](/posts/brooks-design-conceptual-integrity)
- [Part 2: The rational model is wrong](/posts/brooks-design-rational-model)

**Reference:** Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
