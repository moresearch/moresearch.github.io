---
title: "Brooks on design, part 7: great designers, not great processes"
date: 2026-07-11
slug: brooks-design-great-designers
summary: "Great designs come from great designers, not great processes. Organizations must grow, protect, and retain design talent rather than smothering it in process."
tags: design, fred-brooks, great-designers, mentorship, engineering-culture
---

Parts 1-6 covered conceptual integrity, the one-mind rule, protecting the designer, the rational model critique, empiricism, and forces that undermine it. Conclusion: **great designs come from great designers, not from great processes.**

This runs through all four of Brooks's books. *The Mythical Man-Month* (1975): the chief architect. *No Silver Bullet* (1986): better tools alone cannot produce better design. *Computer Architecture* (1997, with Blaauw): every great machine traces to one or two minds. *The Design of Design* (2010): the organizing thesis.

## Definitions

**Process vs. talent.** Process prevents bad design from shipping. It cannot produce good design.

> "Great designs come from great designers. Process can make a good design better; it cannot make a bad designer good."

Modern organizations invest in process: code review, design review, RFCs, ADRs, retrospectives. Supporting structures. Not substitutes. Process raises the floor. Talent raises the ceiling. Conflating them is the central mistake.

Knuth took a decade from *The Art of Computer Programming* to invent TeX — one mind's full attention. Alan Kay: the best software is "a single person's vision carried through." Hoare (1980 Turing lecture): "the most dangerous error is believing that better tools can replace individual insight." Licklider's "Man-Computer Symbiosis" (1960): couple human intuition with machine power. The generation agreed. Process supports. Talent creates.

**Process sequencing.** Let the designer design first. Apply process second. Order matters.

> "The trick is to hold process off long enough to permit great design to occur, so that the lesser issues can be debated once the great design is on the table — rather than smothering it in the cradle."

Committee before design = compromise. Design before committee = review can find weaknesses without diluting vision. Most organizations reverse it. Committee first, compromise second. Integrity lost at step one.

**Dual ladder.** Technical and managerial paths: parallel, equal compensation, equal respect.

> "The dual ladder is everywhere espoused and nowhere practiced. The managerial ladder remains the path to power, prestige, and pay. The technical ladder is too often a consolation prize."

Make it real: equal pay (DE = VP), visible strategic voice, budget protection, staff with exemplars. Numbers differ = not real.

## Growing designers

**Recruit for design sense.** Look at what the candidate built. Does it show integrity? Past performance predicts.

> "Hire for demonstrated design ability, not for interview prowess. Look at what the candidate has built. Does it show conceptual integrity? Do the abstractions make sense? Would you want to use it?"

**Provide mentors.** Taste transmits through apprenticeship.

> "Design judgment is not taught; it is caught. The master-apprentice relationship is how taste transmits. A junior absorbs not just technique but judgment: what to optimize for, what to ignore, when to fight."

**Rotate experiences.** Breadth builds pattern recognition.

> "The designer who has worked in only one domain has a narrow base of pattern recognition. Breadth of experience builds the repertoire from which great designs are synthesized."

**Protect from managing.**

> "The proper office of the manager is to protect his great designers from managing."

Highest-value activity is designing. Every hour in status meetings is lost design work.

## Esthetics, exemplars, case studies

**Esthetics.** "Clean" architectures and "elegant" APIs are not metaphors. They denote real properties: fitness, coherence, economy.

> "We speak of 'clean' machines, 'elegant' languages, 'beautiful' proofs. These are not mere metaphors. They denote a real property: the fitness, coherence, and economy of a design, as judged by those with taste."

Develop taste: study styles, practice in another's style, revise for consistency, hire for demonstrated taste.

**Exemplar gap.** Software designers don't study exemplars the way architects study buildings.

> "Architects study buildings; composers study scores; writers study books; painters visit museums. Software designers too rarely study existing software designs, and when they do, it is to learn how to use the system, not to understand why it was designed as it was."

Great designs contain lessons not conveyable in principles — only in the specific decisions that produced a specific result.

**Case studies.** System/360, OS/360, beach house ("View/360"), kitchen renovation, book design.

> "Why an 8-bit byte? Why 32-bit words? Why a clean separation between architecture and implementation? Each decision made by a small group, debated intensely, then locked. Once locked, they could not be reopened. This discipline — decide, commit, don't revisit — is the only way to ship a coherent architecture."

**Second-system effect.** The most dangerous system is the second one.

> "The second is the most dangerous system a designer ever builds. Having succeeded with the first, he loads the second with every feature he omitted from the first. The result is a bloated, over-budget, late disaster." — *The Mythical Man-Month*, 1975

OS/360 was the cautionary example. JCL was the scar:

> "The job control language JCL for OS/360 is, in my opinion, the worst computer language ever devised, a triumph of committee design over conceptual integrity."

> "I have designed in five media: computer architecture, software, houses, books, and organizations. The principles of design are independent of the medium."

If the same principles govern a computer architecture and a kitchen renovation, you have a theory of design, not just a software methodology.

Does your organization take these claims seriously? One person with real design authority? Design as empirical discovery? Deliberately growing designers? Most answer no. Brooks spent six decades explaining the mistake.

---

**Part 7 of 7.** [← Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce)

**References:** Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).
