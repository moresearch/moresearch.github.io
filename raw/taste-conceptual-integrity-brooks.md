---
title: Taste as Conceptual Integrity
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: When engineers say someone has "taste," what they mean is that person can perceive conceptual integrity. Fred Brooks spent a career explaining why that is the most important property of any designed thing.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

Ray's experiment is effective because the sentence he asks us to imagine contains its own refutation. If a building collapses and the subsequent investigation reveals that an inspector objected and was overruled, we do not describe what happened as the inspector's taste being ignored. We describe it as the inspector's professional judgment being overridden. The word "taste" does not merely fail to capture the gravity of the situation — it actively miscategorizes it. It moves the event from the domain of engineering accountability, where the question is whether the objection was correct, to the domain of aesthetics, where the question is whether the objection was agreeable. These are different categories of judgment, and conflating them is a category error.

The reason the word persists despite this flaw is that it gestures at something real. When a senior engineer looks at a proposed change and says "this does not feel right," they are not stating a preference. They are reporting a perception. They have detected something about the system before they have isolated what it is or produced evidence for its existence. The question is what, exactly, they are detecting.

The most precise answer was given by Fred Brooks. The property being perceived is **conceptual integrity**.

## The property

![The Design of Design — Fred Brooks's final book, and his least read](/images/design-of-design.jpg)

Brooks defined conceptual integrity as the property of a system that feels as though one mind designed it. In *The Mythical Man-Month* (1975) he wrote:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas."

Thirty-five years later, in *The Design of Design*, he reduced the claim to its most concentrated form:

> "Most great works have been made by one mind. The exceptions have been made by two minds."

Neither statement is a claim about aesthetics. Both are structural claims about the conditions under which coherence enters a designed object. Brooks's argument, developed across both books, is that design does not parallelize. Every additional mind added to a design introduces assumptions that differ from the assumptions already present. Those assumptions must be reconciled. Reconciliation requires compromise. Each compromise reduces coherence, because the original set of design ideas was internally consistent and the compromise introduces an element that was not part of that set. Brooks put the point economically: "Many hands make light work — Often. But many hands make more work — Always." You can distribute the labor of implementation across as many people as the coordination overhead permits. You cannot distribute the labor of deciding what the thing is. The design must proceed from one point of decision, or at most two in what Brooks called genuine resonance — a pair who share a mental model so completely that either can speak for the architecture. Three is already a committee, and committees produce settlements, not coherence.

If this argument is accepted, a definition follows: **taste is the ability to perceive conceptual integrity.** It is the faculty by which an engineer detects that a system has maintained coherence or lost it — that the parts compose or have begun to diverge, that the whole still speaks with one voice or has acquired a second. The senior engineer who says "this does not feel right" is perceiving a violation of conceptual integrity before they can name the specific violation. The feeling is not the argument. It is the reason to begin looking for one.

## The positive case

Brooks's canonical example is Reims Cathedral. Ground was broken in 1211. Structural work was completed by 1275, and decorative work continued into the 1460s. Across those two and a half centuries, four master masons — Jean d'Orbais, Jean-le-Loup, Gaucher of Reims, and Bernard de Soissons — directed the construction. Their names were inscribed in a labyrinth set into the nave floor. The labyrinth was not a signature in the modern sense; it was a public oath to a design larger than any single lifetime. Each man bound himself to the constraints established by his predecessor, and each successor did the same.

The coherence of the result was not achieved by consensus. It was achieved by constraint. Reims was among the first buildings to use stones of standardized dimensions, which reduced the degrees of freedom available to each successive builder. The structural system — four-part rib vaults producing arcades of identical pillars rather than the alternating pillars and piers of earlier Gothic — made the architectural rhythm self-enforcing. Once the pillar spacing and vault geometry were fixed, any architect who altered them would have broken the structural logic of the building. The plan had mechanical authority. It did not require that each successor agree with it; it required only that each successor could not alter it without producing visible damage.

The result is a building in which you cannot identify the transition from one architect to the next. Two centuries of evolving Gothic fashion — bar tracery in the transept roses giving way to full Rayonnant in the west façade — read not as competing visions but as variation within a theme. The first architect's taste set the constraints. The constraints outlived him because they were structural, not advisory.

![Reims Cathedral — four architects, two and a half centuries, one coherent result](/images/reims-cathedral.jpg)

## The negative case

If Reims demonstrates what happens when taste has authority, the Bradley Fighting Vehicle demonstrates what happens when it does not.

The Bradley began as a light troop carrier. Over 17 years and at a cost of $14 billion, successive groups of stakeholders added requirements. Armor advocates wanted more survivability. Infantry commanders wanted more troop capacity. Generals wanted more firepower. Contractors wanted larger contracts. Each of these requests was defensible when evaluated against the objectives of the stakeholder who made it. The problem is not that any individual request was unreasonable. The problem is that there was no single mind with the authority to evaluate each request against the coherence of the whole — to say, this feature is reasonable on its own terms, and it will break the design, so the answer is no.

The result was a vehicle that Sergeant Fanning, a character in the 1998 HBO film *The Pentagon Wars*, describes with the precision of someone who has spent years watching the logic play out:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

![The Pentagon Wars (1998) — 17 years, $14 billion, and a machine that did nothing well](/images/pentagon-wars.jpg)

Every clause in that sentence corresponds to a stakeholder requirement that was defensible when considered in isolation. Together they describe a machine that performed no role adequately because it was required to perform every role simultaneously. The film documents that live-fire testing was manipulated to conceal the vehicle's failures; that the officers responsible for the manipulation were promoted; and that Lt. Colonel James Burton — the one person in the story who acted as though coherence mattered, who forced an honest test over the objections of his chain of command — was eventually forced into retirement.

The Bradley is the limiting case of what Brooks described as "many good but independent and uncoordinated ideas." It was not produced by incompetence or bad faith. It was produced by a structural condition: a decision-making process in which every stakeholder could add a requirement and no single mind could reject one. The output was not a design. It was a negotiated settlement. And it is the reason the word "taste" must be defined precisely. If taste means personal preference, then ignoring it is reasonable. If taste means the perception of conceptual integrity, then ignoring it is how you produce the Bradley.

## The organizational conditions

Brooks did not stop at naming the property. On IBM System/360, he, Gene Amdahl, and Gerrit Blaauw implemented the organizational structures required to preserve it. The principles are distributed across both books. Considered together, they form a set of necessary conditions.

**One mind, or two in genuine resonance.** Brooks and Blaauw achieved a working relationship in which either could speak for the architecture — a state Brooks described as resonance and glossed with what may be the most human observation in either book: "And two is indeed a magic number for collaborations; marriage was a brilliant invention and has a lot to be said for it." Three people cannot maintain this state. At three, the dynamic shifts from resonance to negotiation, and negotiation produces settlements rather than designs.

**Real veto power.** The architect must be able to refuse a feature and have the refusal stand. Brooks described this as saying no "repeatedly, to smart people with good arguments." Advisory veto — in which the architect may object but a person higher in the reporting chain may override the objection — is indistinguishable in practice from having no veto at all. The vice president approves the feature. It ships. The coherence of the system degrades. The vice president eventually moves to a different organization. The architect remains, responsible for the accumulated consequences of decisions they were not permitted to make.

**The separation of design from implementation.** Brooks devoted a chapter of *The Design of Design* to what he called the divorce of design — the progressive separation, beginning around the 16th century, of the act of specifying from the act of making. On System/360, this principle was operationalized as a small architecture team that defined what the system was and a large implementation team that built how it worked. "The architecture team must be protected; the implementation team must be coordinated." Organizations that assign both functions to the same people discover that neither is performed adequately. The architects, drawn into the demands of implementation, cease to think about systemic coherence. The implementers, asked to make architectural decisions, optimize for the local context at the expense of the whole.

**Protection from organizational forces.** The System/360 architecture team was deliberately insulated from field sales, who wanted features for specific customers; from engineering, who wanted optimizations that would have compromised the cleanliness of the abstractions; and from customers, who demanded backward compatibility with their existing systems. Each of these demands was legitimate. Accommodating any significant fraction of them collectively would have destroyed the system's coherence. The insulation was not a privilege extended to the architects. It was a structural precondition for the design work to be possible at all.

**Career paths that reward refusal.** Brooks implied this condition throughout both books without stating it as directly as the evidence warrants. In most organizations, agreeing to requests accumulates political capital and refusing them spends it. If the incentive structure penalizes the person who protects coherence, then taste may exist within the organization — individual engineers may perceive what is happening to the system — but it will be systematically powerless. The Bradley provides the limiting case: the officers who manipulated the tests were promoted; the officer who forced an honest evaluation was removed from service.

## Open questions

Brooks's conditions assume a human designer producing at human speed. Coding agents break that assumption. They generate tokens that are statistically probable given their training distribution and context window. They do not generate coherent designs, because coherence across the entire surface area of a system is not a statistical property — it is a structural property that must be imposed by an entity capable of holding the whole in view. An agent can write a function that compiles and passes its tests while violating the design philosophy that every adjacent function observes. It does so without awareness, because the design philosophy is not legible in the text the agent was trained on. The rationale lived in the head of the person with the authority to say no.

This raises questions Brooks did not live to address:

**Does taste become the irreducible human contribution, or does the role become impossible?** An architect can veto a human engineer's output at the speed of code review. An architect cannot veto at the speed of token generation. If an agent produces a Bradley-scale incoherence in an afternoon — and the Bradley took 17 years — maintaining conceptual integrity may require constraints so deeply embedded in the generation pipeline that the architect's role shifts from reviewer to tool-builder. Whether that role still satisfies Brooks's definition of the designer is an open question.

**Can taste be automated?** If taste is the ability to perceive violations of conceptual integrity, and conceptual integrity is a property of the system considered as a whole, then automating taste requires building a system capable of forming a model of the whole. Current agents cannot do this. It is unclear whether the limitation is architectural — larger context windows, better retrieval — or categorical. If coherence can only be perceived by holding the entire set of design decisions in view simultaneously, and the set grows faster than any context window can expand, then taste may be the capacity that survives automation because it is the capacity that automation cannot reach.

**If an agent generates 500 lines and a human keeps 30, who wrote the feature?** The question is not rhetorical. If the human's contribution is selecting which output maintains the system's coherence, then the job is not prompt engineering. It is architecture in Brooks's sense: one mind enforcing conceptual integrity against a force that produces incoherence at speed. The title matters because the authority matters. You cannot enforce coherence if your role is understood by the organization as "the person who writes the prompts."

---

Ray's experiment contains its own answer, provided the term is defined correctly. "The building fell down after the inspector's taste was ignored" fails because it miscategorizes professional judgment as personal preference. Reformulate: the Bradley became a death trap because the people capable of perceiving conceptual integrity lacked the authority to enforce it. That sentence identifies a faculty, a property, and a failure mode. It describes an event for which someone can be held accountable.

Brooks supplied the property, the evidence, and the organizational conditions under which the faculty can function. What he could not supply — what no one can yet supply — is an account of whether taste, defined this way, survives an era in which code is generated faster than any human can perceive whether it coheres. That question is now operational. We will have an answer whether we want one or not.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — Conceptual integrity as the most important consideration in system design; the one-mind rule; the architect's authority to say no.
- Brooks, F. P. (2010). *The Design of Design.* — "Most great works have been made by one mind. The exceptions have been made by two minds." "Many hands make light work — Often. But many hands make more work — Always." The divorce of design from implementation.
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)
