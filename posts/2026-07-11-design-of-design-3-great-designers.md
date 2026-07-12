---
title: Brooks on Software Design Series: great designers, not great processes
date: 2026-07-11
slug: brooks-design-great-designers
summary: "Great designs come from great designers, not great processes. Organizations must grow, protect, and retain design talent. Your agile coach can't save you. Only taste can."
tags: design, fred-brooks, great-designers, mentorship
---

Parts 1-6: conceptual integrity, one mind, protection, rational model critique, empiricism, forces against it. The dominoes all fall one way. The conclusion is uncomfortable. The industry has spent fifty years avoiding it.

Conclusion: **great designs come from great designers, not from great processes.** Your Jira workflow has never designed anything. Your RFC template has never had an idea. Your retrospective has never produced insight — it produced action items. Action items are not designs. They are evidence that a meeting occurred.

This runs through all four Brooks books. *The Mythical Man-Month* (1975): the chief architect. *No Silver Bullet* (1986): tools alone cannot produce better design. *Computer Architecture* (1997, with Blaauw): every great machine traces to one or two minds. *The Design of Design* (2010): the organizing thesis. Four books, one argument. The industry read them and built SAFe anyway. Brooks is patient. Brooks is dead. The argument remains.

## Definitions

**Process vs. talent.** Process prevents bad design from shipping. It cannot produce good design. Process is a seatbelt. It will not drive the car. It will not choose the destination. It will only prevent some of the damage when the driver — who is still necessary — makes a mistake.

> "Great designs come from great designers. Process can make a good design better; it cannot make a bad designer good."

Modern organizations invest in process: code review, design review, RFCs, ADRs, retrospectives. Supporting structures, not substitutes. Process raises the floor. Talent raises the ceiling. Conflating them is the central mistake. You cannot review your way to brilliance. You can only review your way to adequacy. Adequacy at scale is still adequacy. The world has enough adequate software. Nobody remembers who built it.

Knuth took a decade for TeX — one mind, one vision, one beautifully typeset result. Alan Kay: best software is "a single person's vision carried through." Hoare (1980 Turing lecture): "the most dangerous error is believing better tools can replace individual insight." Licklider (1960): couple human intuition with machine power. The generation agreed. Process supports. Talent creates. This is not controversial among people who have built great things. It is only controversial among people who manage them and wish the management were sufficient.

**Process sequencing.** Designer designs first. Process applies second. Order matters. Skip this and you get compromise dressed in review comments, which is still compromise.

> "The trick is to hold process off long enough to permit great design to occur, so that the lesser issues can be debated once the great design is on the table — rather than smothering it in the cradle."

Committee before design = compromise. Design before committee = review finds weaknesses without diluting vision. Most organizations reverse it. Committee first, compromise second, individual execution third. Integrity lost at step one. The rest is expensive theater with good catering.

**Dual ladder.** Technical and managerial paths: parallel, equal compensation, equal respect. Every company claims to have this. Almost none do. You can tell by asking one question: does your Distinguished Engineer report to a VP who controls their compensation? If yes, the ladder is a label. Labels are cheap. That's why everyone has one.

> "The dual ladder is everywhere espoused and nowhere practiced. The managerial ladder remains the path to power, prestige, and pay. The technical ladder is too often a consolation prize."

Make it real: equal pay (DE = VP), visible strategic voice, budget protection, staff with exemplars. Otherwise it's a parking lot for brilliant people you don't want to lose but don't want to empower. They know. They're brilliant.

## Growing designers

**Recruit for design sense.** Look at what the candidate built. Past performance predicts. Whiteboard interviews predict nothing except whiteboard interview performance, which is a skill nobody needs after the interview.

> "Hire for demonstrated design ability, not for interview prowess. Look at what the candidate has built. Does it show conceptual integrity? Do the abstractions make sense?"

**Mentor.** Taste transmits through apprenticeship. Not docs. Not talks. Not "lunch and learns." Through working next to someone who has it and watching how they think.

> "Design judgment is not taught; it is caught. The master-apprentice relationship is how taste transmits. A junior absorbs judgment: what to optimize for, what to ignore, when to fight."

**Rotate.** Breadth builds pattern recognition. One domain, one technology, one problem space = one pattern. You cannot synthesize from one example any more than you can learn a language from one word.

> "The designer who has worked in only one domain has a narrow base. Breadth builds the repertoire from which great designs are synthesized."

**Protect from managing.** The reward for great design is promotion to management. This is like rewarding a great chef by making them run the restaurant. Different skills. Different outcomes. The chef stops cooking. The food gets worse. Everyone notices except the person who promoted the chef.

> "The proper office of the manager is to protect his great designers from managing."

Highest-value activity is designing. Every hour in meetings is lost design work. Count your best designer's meetings. More than five per week? You're burning your best asset for scheduling convenience. The calendar is a furnace. Your talent is the fuel.

## Esthetics, exemplars, cases

**Esthetics.** "Clean" and "elegant" are not metaphors. They denote fitness, coherence, economy — judged by those with taste. Taste is real. Taste is testable. Taste is not subjective in the way that matters. You can be wrong about elegance. People wrong about elegance are usually wrong about estimates too, for the same reason: they can't see the shape of the thing.

> "We speak of 'clean' machines, 'elegant' languages, 'beautiful' proofs. These are not mere metaphors. They denote a real property: the fitness, coherence, and economy of a design."

Develop taste: study styles, practice another's, revise for consistency, hire for demonstrated taste. There is no shortcut. Taste is accumulated judgment. Judgment is accumulated mistakes. Mistakes are accumulated by doing. Go do. Make mistakes. Learn to see them before you make them. That's taste.

**Exemplar gap.** Software designers don't study exemplars like architects study buildings. Architects visit buildings. Composers study scores. Most software designers have not read the source code of the systems they admire. They use them. They don't study them. That's like becoming a novelist by only reading book reviews. You'll learn what's popular. You won't learn how sentences work.

> "Architects study buildings; composers study scores; writers study books; painters visit museums. Software designers too rarely study existing designs, and when they do, it is to learn how to use the system, not why it was designed as it was."

Great designs contain lessons not conveyable in principles — only in the decisions that produced a result. Read TeX's source. Read the Unix kernel. Read anything by people who thought harder than you. It's all there, waiting. Almost nobody does it. The people who do become the designers everyone else envies.

**Case studies.** System/360, OS/360, beach house, kitchen, book design. Brooks uses his own work. Use yours. Study your wins and your disasters. Both have more to teach than any methodology book.

> "Why an 8-bit byte? Why 32-bit words? Why separate architecture from implementation? Each decision made by a small group, debated intensely, then locked. This discipline — decide, commit, don't revisit — ships coherent architecture."

**Second-system effect.** The most dangerous system is the second. Version one succeeds. Version two gets every feature cut from version one. Version two sinks. Everyone is surprised. Nobody should be. This has been documented since 1975. It still happens. It will happen on your next project unless you actively prevent it.

> "The second is the most dangerous system a designer ever builds. Having succeeded with the first, he loads the second with every feature he omitted. The result is a bloated, over-budget, late disaster." — *The Mythical Man-Month*, 1975

OS/360 was the example. JCL was the scar: "the worst computer language ever devised, a triumph of committee design over conceptual integrity." Brooks wrote this about his own project. That is intellectual honesty. Most of us blame the tools, the timeline, the market. He blamed himself and wrote four books. The least we can do is read one of them.

> "I have designed in five media: computer architecture, software, houses, books, and organizations. The principles of design are independent of the medium."

Same principles govern a computer architecture and a kitchen renovation. If your design principles can't survive contact with a kitchen, they weren't principles. They were platform-specific habits dressed up as wisdom. The test of a principle is whether it holds when the medium changes. Brooks tested his across five media. How many have you tested yours across?

Does your organization take these claims seriously? Design authority? Empirical discovery? Growing designers? Most answer no. Brooks spent six decades explaining the mistake. We spent six decades nodding and building process instead. The book is on the shelf. It's short. You have time. The question is whether you have the nerve.

---

[← Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 3](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 4](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).
