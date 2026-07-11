---
title: "Brooks on design, part 4: the rational model is wrong"
date: 2026-07-11
slug: brooks-design-rational-model
summary: "Fred Brooks argues that the rational model of design — requirements first, design second, implementation third — is fundamentally wrong. You don't know the goal at the start, and requirements co-evolve with the design."
tags: design, fred-brooks, rational-model, waterfall, co-evolution, peter-naur
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


The next part takes up the positive alternative: if we can't reason our way to a correct design, what do we do instead?

---

**This is part 4 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the Reims Cathedral](/posts/brooks-design-conceptual-integrity)
- [Part 2: The one-mind rule](/posts/brooks-design-one-mind-rule)
- [Part 3: Protecting the designer](/posts/brooks-design-protecting-designer)
- [Part 5: The empiricist alternative](/posts/brooks-design-empiricist-alternative)
- [Part 6: How experts go wrong and the divorce of design](/posts/brooks-design-experts-divorce)
- [Part 7: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.
