---
title: "Brooks on design, part 7: great designs come from great designers — not great processes"
date: 2026-07-11
slug: brooks-design-great-designers
summary: "Brooks argues that great designs come from great designers, not great processes — and that organizations must deliberately grow, protect, and retain design talent rather than smothering it in process."
tags: design, fred-brooks, great-designers, engineering-culture, mentorship
---

The preceding six parts covered conceptual integrity, the one-mind rule, protecting the designer, the critique of the rational model, the empiricist alternative, and the forces that undermine empiricism. This final part addresses the uncomfortable conclusion: **great designs come from great designers, not from great processes.**

This runs through all of Brooks's major works. *The Mythical Man-Month* (1975): the insistence on a chief architect. *No Silver Bullet* (1986): better tools alone cannot produce better design — only better designers can. *Computer Architecture: Concepts and Evolution* (1997, with Blaauw): every great machine architecture traces to one or two minds. *The Design of Design* (2010): the organizing thesis.

## Definition: process vs. talent

**Process vs. talent.** Process prevents bad design from shipping. It cannot produce good design. For that, you need individual brilliance.

> "Great designs come from great designers. Process can make a good design better; it cannot make a bad designer good."

The modern organization invests heavily in process: code review, design review, RFCs, ADRs, retrospectives. These are not bad. They are supporting structures, not substitutes. Process raises the floor. Talent raises the ceiling. Conflating the two is the central mistake of engineering management.

Knuth took a decade from *The Art of Computer Programming* to invent TeX. Alan Kay described the best software as "a single person's vision carried through." Hoare, in his 1980 Turing lecture, warned against believing that "better tools and more systematic methods can replace the need for individual insight." Licklider's "Man-Computer Symbiosis" (1960) argued the goal was coupling human intuition with machine power. The whole generation agreed: process supports. Talent creates.

## Definition: process sequencing

**Process sequencing.** Let the designer design first. Apply process second. The order matters.

> "The trick is to hold process off long enough to permit great design to occur, so that the lesser issues can be debated once the great design is on the table — rather than smothering it in the cradle."

If you convene a review committee before a strong design exists, the committee produces a compromise. If you let a strong designer produce a coherent design first, then subject it to review, the review can find weaknesses without diluting the vision.

Most organizations do the reverse. Committee first, compromise second, individual execution third. Conceptual integrity is already lost at step one.

## Definition: the dual ladder

**Dual ladder.** Technical and managerial career paths should be parallel, equally compensated, equally respected.

The problem: organizations promote their best technical people into management. The best engineer becomes the manager. The organization loses a great designer and gains a mediocre manager.

> "The dual ladder is everywhere espoused and nowhere practiced. The managerial ladder remains the path to power, prestige, and pay. The technical ladder is too often a consolation prize — a place to park brilliant people who won't or can't manage."

Making it real requires equal compensation (Distinguished Engineer = VP), visible strategic voice, protection from budget cuts, and staffing with exemplars. If the numbers differ, the ladder is not real.

## Growing designers

Design talent can be grown:

**Recruit for design sense.** Look at what the candidate has built. Does it show integrity? Do the abstractions work? Past performance predicts future performance.

> "Hire for demonstrated design ability, not for interview prowess. Look at what the candidate has built. Does it show conceptual integrity? Do the abstractions make sense? Would you want to use it?"

**Provide mentors.** Taste transmits through apprenticeship.

> "Design judgment is not taught; it is caught. The master-apprentice relationship — explicit in architecture, mostly informal in software — is how taste transmits. A junior working under a senior absorbs not just technique but judgment: what to optimize for, what to ignore, when to fight and when to concede."

**Rotate experiences.** Breadth builds pattern recognition.

> "The designer who has worked in only one domain, with one set of constraints, has a narrow base of pattern recognition. Breadth of experience — across domains, technologies, and constraints — builds the repertoire from which great designs are synthesized."

**Protect from managing.**

> "The proper office of the manager is to protect his great designers from managing."

A designer's highest-value activity is designing. Every hour in a status meeting is an hour of lost design work. The manager's job is managing everything else.

## Esthetics, exemplars, and case studies

**Esthetics in design.** "Clean" architectures and "elegant" APIs are not metaphors. They denote real properties: fitness, coherence, economy — judged by those with taste.

> "We speak of 'clean' machines, 'elegant' languages, 'beautiful' proofs. These are not mere metaphors. They denote a real property: the fitness, coherence, and economy of a design, as judged by those with taste."

Developing taste: study others' styles, practice in another's style, revise for consistency, choose designers with demonstrated taste.

**Exemplar gap.** Software designers do not study exemplars the way architects study buildings or composers study scores.

> "Architects study buildings; composers study scores; writers study books; painters visit museums. Software designers too rarely study existing software designs, and when they do, it is to learn how to use the system, not to understand why it was designed as it was."

Great designs contain lessons that cannot be conveyed in principles — only in the specific sequence of decisions that produced a specific result.

**Case studies.** Brooks draws on System/360, OS/360, his beach house ("View/360"), a kitchen renovation, and a book design.

> "Why an 8-bit byte? Why 32-bit words? Why a clean separation between architecture and implementation? Each of these decisions was made by a small group, debated intensely, and then locked. Once locked, they could not be reopened. This discipline — decide, commit, don't revisit — is the only way to ship a coherent architecture."

**Second-system effect.** The most dangerous system a designer ever builds is the second one.

> "The second is the most dangerous system a designer ever builds. Having succeeded with the first, he loads the second with every feature he omitted from the first, every embellishment suggested by users, every capability the hardware now permits. The result is a bloated, over-budget, late disaster." — *The Mythical Man-Month*, 1975

OS/360 was Brooks's cautionary example. JCL was the scar:

> "The job control language JCL for OS/360 is, in my opinion, the worst computer language ever devised, a triumph of committee design over conceptual integrity."

> "I have designed in five media: computer architecture, software, houses, books, and organizations. This book draws lessons from all five. The principles of design are independent of the medium."

If the same principles govern a computer architecture and a kitchen renovation, you have shown something deeper than a software methodology. You have shown a theory of design.

The question the book leaves you with: does your organization take these claims seriously? A single person with real design authority? Design as empirical discovery, not rational derivation? Deliberately growing designers, not hiring for coding speed and hoping taste comes along?

Most answer no. Brooks spent six decades explaining why that is a mistake.

---

**This is part 7 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity — the most important property](/posts/brooks-design-conceptual-integrity)
- [Part 2: Why one mind must rule the design](/posts/brooks-design-one-mind-rule)
- [Part 3: How to protect designers from their organizations](/posts/brooks-design-protecting-designer)
- [Part 4: The waterfall model is wrong and harmful](/posts/brooks-design-rational-model)
- [Part 5: Build, test, iterate — the empiricist method](/posts/brooks-design-empiricist-alternative)
- [Part 6: Why experts design the wrong thing beautifully](/posts/brooks-design-experts-divorce)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
