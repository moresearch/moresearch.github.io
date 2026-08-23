---
title: Accidental Complexity Is the Only Complexity You Can Remove
date: 2026-08-02
slug: brooks-accidental-complexity
summary: In No Silver Bullet, Brooks split software difficulty into essence and accidents, then argued the essence dominates. Forty years of evidence point the other way — and the distinction, not the conclusion, is what should have survived.
tags: fred-brooks, no-silver-bullet, complexity, accidental-complexity, essential-complexity, software-engineering, simplicity, design
---

In April 1987, *IEEE Computer* published Fred Brooks's essay "No Silver Bullet: Essence and Accidents of Software Engineering." It may be the most quoted and least read document in the history of the field. Everyone knows its conclusion — "There is no silver bullet" — because it has been the standard reply to every promised miracle since. Almost nobody remembers its actual subject, which was not miracles at all. The essay's real contribution was a distinction: between the complexity that is *essential* to the problem and the complexity that is *accidental* to the solution. Forty years later, that distinction has turned out to be more durable, and more useful, than the prediction it was built to support.

The abstract states the conclusion with characteristic compression:

> "There is no single development, in either technology or management technique, which by itself promises even one order-of-magnitude improvement within a decade in productivity, in reliability, in simplicity."

But the sentence that carries the essay is the one that defines the terms:

> "All software construction involves essential tasks, the fashioning of the complex conceptual structures that compose the abstract software entity, and accidental tasks, the representation of these abstract entities in programming languages and the mapping of these onto machine languages within space and speed constraints."

Essential complexity is the difficulty that lives in the problem — the conceptual construct itself, its interlocking data structures, relationships, and invariants. Accidental complexity is the difficulty that lives in the *means of production*: the languages, tools, platforms, conventions, and accumulated infrastructure through which the construct is forced to pass. Brooks, following Aristotle, called them the "essence" and the "accidents":

> "Following Aristotle, I divide them into essence—the difficulties inherent in the nature of the software—and accidents—those difficulties that today attend its production but that are not inherent."

The word "today" is doing enormous work in that sentence. Accidental complexity is contingent. It is a property of the moment's tooling, not of the problem. It can in principle be eliminated by a better way of doing things — and every time the industry has actually gotten dramatically better, that is what happened.

## The arithmetic Brooks did, and the claim he made

Brooks did not merely assert that the essence dominates. He supplied an argument, and it was an arithmetic one:

> "How much of what software engineers now do is still devoted to the accidental, as opposed to the essential? Unless it is more than 9/10 of all effort, shrinking all the accidental activities to zero time will not give an order of magnitude improvement."

That is the whole essay in miniature. A "silver bullet" is defined as a tenfold improvement. If the accidental part is half the effort, eliminating it entirely buys a factor of two. If it is nine-tenths, eliminating it entirely buys a factor of ten. Brooks's empirical guess in 1986 was that the accidental share was below nine-tenths — so no single advance, however complete, could cross the order-of-magnitude threshold. Hence no silver bullet.

Notice what follows from this framing. The claim is not that accidental complexity is small. The claim is that *even if you eliminated every scrap of it*, the improvement would be bounded by how much of your effort is actually essential. The conclusion is hostage to a ratio Brooks estimated by hand. He was explicit about the uncertainty — the "unless" is doing the work — and the ratio was never measured. It was a guess, offered in a paragraph, in a field where guesses about ratios have a way of becoming doctrine.

## The four properties of the essence

Brooks identified four properties of software that make the essential difficulty irreducible: **complexity, conformity, changeability, and invisibility**.

**Complexity.** Software entities are more complex for their size than perhaps any other human construct, because no two parts are alike — above the statement level, repetition is a smell and gets factored out. The parts interact nonlinearly:

> "Software systems have orders of magnitude more states than computers do."

And then the sentence the whole distinction rests on:

> "The complexity of software is an essential property, not an accidental one. Hence descriptions of a software entity that abstract away its complexity often abstract away its essence."

**Conformity.** Here Brooks is subtlest, and the subtlety matters for everything that comes later. The complexity of conformity is *not* inherent to the software; it is imposed from outside:

> "Much of the complexity he must master is arbitrary complexity, forced without rhyme or reason by the many human institutions and systems to which his interfaces must conform."

An interface defined by a legacy mainframe, a regulation, or a competitor's format is arbitrary. It could have been any shape. But the software that must serve it inherits that arbitrariness, and no redesign of the software alone can remove it.

**Changeability.** "All successful software gets changed." Software embodies function, and function is the part most exposed to the pressure of change. Unlike a building, where cost dampens the whims of the changer, software is "pure thought-stuff, infinitely malleable" — so it gets changed constantly, by users discovering new uses and by the machines it must outlive.

**Invisibility.** "The reality of software is not inherently embedded in space." It has no geometric representation the way a floor plan captures a building. Diagramming software yields several superimposed directed graphs — control, data, dependency, time — which are "usually not even planar, much less hierarchical."

These four are the essence: the irreducible difficulty of getting a complex, arbitrary, changing, invisible construct right. Brooks's conclusion follows:

> "I believe the hard part of building software to be the specification, design, and testing of this conceptual construct, not the labor of representing it and testing the fidelity of the representation."

## What the silver bullets actually did

The most instructive part of the essay is its review of the candidates — because Brooks shows, case by case, that every genuine advance in software productivity had attacked the *accidental* part, not the essential one.

**High-level languages.** The most powerful stroke for productivity, reliability, and simplicity — credited with at least a factor of five:

> "To the extent that the high-level language embodies the constructs wanted in the abstract program and avoids all lower ones, it eliminates a whole level of difficulty that the programmer would otherwise have to master."

The programmer was not solving a harder problem with Fortran. They were freed from bits, registers, conditions, branches, channels, and disks — from the accidents of the machine. That is why the first transition from machine language to a high-level language produced the huge payoff, and why each subsequent language improvement pays less: the accidents are fewer each time, because the previous round removed them.

**Object-oriented programming and Ada.** Brooks gives these their due — "Each removes one more accidental difficulty from the process, allowing the designer to express the essence of his design without having to express large amounts of syntactic material that add no new information content" — and then delivers the sentence that should be carved over every architecture department:

> "Such advances can do no more than to remove all the accidental difficulties from the expression of the design. The complexity of the design itself is essential; and such attacks make no change whatever in that."

**AI, expert systems, automatic programming, graphical programming, verification, environments, workstations.** Each gets a section and each is shown to be either an attack on accidents (verification, environments) or, in the case of AI and automatic programming, a promise to automate the *essential* task — which Brooks judged impossible, because the essence is not a representation problem, it is a thinking problem.

The pattern is the point. Every historical leap in software productivity — from assembly to high-level languages, from batch to time-sharing, from unstructured to modular — was a removal of accidents. Brooks's own history supports a stronger claim than the one he made: the gains in this industry have *always* come from removing accidental complexity, and the residual accidental complexity is precisely where the next gain is hiding.

> The essence is the part of the difficulty you cannot remove. The accidents are the part you can. Every order of magnitude software has ever gained came from the second kind.

## The disagreement: Out of the Tar Pit

Brooks's essay was answered, twenty years later, by Ben Moseley and Peter Marks in "Out of the Tar Pit" (2006) — a paper that opens with the same diagnosis and then disagrees with the ratio:

> "Complexity is the single major difficulty in the successful development of large-scale software systems. Following Brooks we distinguish accidental from essential difficulty, but disagree with his premise that most complexity remaining in contemporary systems is essential."

The disagreement is not about the distinction. It is about which side of the line most of the complexity actually sits on:

> "We disagree. Complexity itself is not an inherent (or essential) property of software (it is perfectly possible to write software which is simple and yet is still software), and further, much complexity that we do see in existing software is not essential (to the problem)."

Their test is clean: a complexity is essential only if the team would have to contend with it "even in the ideal world" — with perfect infrastructure and no performance constraint. By that test, the majority of what contemporary systems carry — the handling of state, the explicit management of control flow, the plumbing — is accidental:

> "We believe that the major contributor to this complexity in many systems is the handling of state and the burden that this adds when trying to analyse and reason about the system."

This is the point at which the forty-year argument actually lives. Brooks guessed the accidental share was below nine-tenths and therefore irrelevant to order-of-magnitude claims. Moseley and Marks argue the accidental share is the *dominant* share — which is why the paper is titled after the tar pit: the tar is mostly of our own manufacture. If they are right, then the search for silver bullets was never the error. The error was concluding that the essence dominates without measuring it.

The empirical record since 1986 has been kinder to Moseley and Marks than to Brooks. The industry has not lacked for order-of-magnitude removals of accidents: garbage collection eliminated entire classes of memory-management errors; managed runtimes eliminated whole toolchains of manual build and deploy; static typing eliminated whole classes of runtime failures; source control, package managers, and CI removed categories of coordination work that once consumed weeks. Each was an accident removed. And each removal revealed that the next layer down was also, in substantial part, accidental.

## Where accidental complexity lives today

The most useful way to read Brooks's distinction is as a diagnostic: for any piece of difficulty in a system, ask whether it is forced by the problem or by the choices made along the way. Run that test over a modern codebase and the accidental fraction is hard to keep below nine-tenths.

**The toolchain itself.** A build system exists to manage the accidents created by the previous build system. The dependencies in a modern application are a geography of other people's accidents: transitive packages, platform shims, polyfills, and compatibility layers, each of which solved one problem by introducing the possibility of a hundred. The 2016 left-pad incident — the unpublishing of an eleven-line npm package taking down thousands of projects — was an accident of the *packaging* layer, not of anyone's essential problem. The software supply chain attacks of the last decade live in exactly this sediment. Every dependency is accidental complexity you have chosen to import, and the attacks are the price of the import. The supply chain argument for minimalism is the accidental-complexity argument wearing a security hat.

**The architecture.** The microservices wave was, at bottom, an attempt to escape accidental complexity that the industry had built for itself — the coordination cost of large codebases, the fear of the monolith. It then manufactured a new layer of accidents: distributed transactions, network partitions, service discovery, observability, deployment matrices. The pattern is so common it has a name, and it is Brooks's pattern: each generation's solution is the next generation's accidental complexity. When Segment published "Goodbye Microservices" in 2018, documenting its move from a microservices zoo back to a modular monolith, the headline finding was not about ideology — it was that the distributed systems complexity had become the dominant cost, and none of it was essential to the business problem. They were removing accidents.

**The abstraction layer.** Joel Spolsky's "Law of Leaky Abstractions" (2002) states the mechanism precisely: "All non-trivial abstractions, to some degree, are leaky." An abstraction that worked perfectly would have removed its accidents permanently. A leaky abstraction — TCP over satellite links, a filesystem over NFS, an ORM over a join — keeps the accidents it was supposed to remove and adds new ones on top. The complexity is not gone. It has moved into the leak.

Rich Hickey's "Simple Made Easy" (2012) made the same point with different words. The industry optimizes for *easy* — familiar, convenient, near at hand — while *simplicity* — unbraided, unentangled, each thing doing one thing — is what actually reduces complexity. Easy is frequently the delivery mechanism for accidental complexity: the framework that makes today's task trivial by braiding in a thousand obligations that must be paid later. Ousterhout's *A Philosophy of Software Design* (2018) gives the definition that should close the argument: "Complexity is anything related to the structure of a software system that makes it hard to understand and modify the system." By that definition, most of the structural difficulty in contemporary systems is accidental, because most of it comes from choices — and choices can be unmade.

Dijkstra got there first, in 1972:

> "The tools we use have a profound (and devious!) influence on our thinking habits, and, therefore, on our thinking abilities."

And, in 1988, on the economics of the whole problem:

> "Simplicity is a great virtue but it requires hard work to achieve it and education to appreciate it. And to make matters worse: complexity sells better."

Complexity sells better. That sentence is the reason the accidental share never falls on its own: there is no market pressure toward removing accidents, because accidents are usually invisible to the buyer. They are paid for later, in maintenance, in onboarding, in incidents — in the tar.

## The one ratio Brooks got right by accident

Here is the irony that makes the essay worth reading today. Brooks listed AI among the candidate silver bullets and dismissed it, as he dismissed everything, on the grounds that the essence cannot be automated. What the 2020s actually demonstrated is subtler: AI did not remove the essential difficulty — it collapsed the *accidental* part, which is precisely the part Brooks defined as removable. The labor of representing a conceptual construct in a programming language, and mapping it onto machines — his own definition of the accidental task — is what generative coding tools have made nearly free.

> The essence is still there. The accidents are what the AI ate.

That is why the residual bottleneck in the AI era looks exactly like Brooks's list of essential properties. Specification is the bottleneck — the conceptual construct must still be designed, and designing it is thinking, not typing. Conformity is the bottleneck — the arbitrary interfaces of the world do not become less arbitrary because code is generated. Changeability is the bottleneck — generated code inherits the pressure of change at higher velocity. Invisibility is the bottleneck — an agent cannot hold the whole directed graph of a system in view, and neither can its context window. The four essential properties are now the entire difficulty, because the accidents that used to hide them have been removed. This site has argued the same thing from two directions: that dark factories simply move the complexity upstream into specification and downstream into validation, and that the irreplaceable human contribution is the perception of conceptual integrity — the one-mind judgment of whether the whole still coheres.

Brooks's conclusion — no silver bullet — was a prediction about a ratio he guessed. The prediction's status is now genuinely contested: if accidental complexity is the dominant share, as the evidence increasingly suggests, then a tool that eliminates accidents *is* the order-of-magnitude event, and the industry just lived through it. But his distinction has survived the prediction, because the distinction is not an empirical claim. It is a classification, and it is the classification that does the work:

> The essence is the complexity the problem forces on you. The accidents are the complexity you chose. You cannot remove the first. You can only remove the second. Therefore the only complexity you can actually remove is accidental complexity — and removing it is the entire history of progress in this field.

The practical consequence is not a technology. It is a discipline of subtraction. Every dependency is a choice. Every framework is a choice. Every layer of indirection is a choice. Every one of them imported accidents that will be paid for later, and each one was imported because the accident it removed was more visible than the accidents it added. Brooks's essay is usually cited as the reason progress is impossible. Read correctly, it is the reason progress is possible at all — one removal at a time.

> There is no royal road, but there is a road.

---

**References — Brooks:**

- Frederick P. Brooks Jr. [No Silver Bullet: Essence and Accidents of Software Engineering](https://worrydream.com/refs/Brooks-NoSilverBullet.pdf). *Information Processing '86* (IFIP World Congress, 1986); reprinted in *IEEE Computer* 20(4), April 1987. — The distinction between essential and accidental difficulty; the four essential properties (complexity, conformity, changeability, invisibility); the 9/10 arithmetic; the review of candidate silver bullets.
- Brooks, F. P. *The Mythical Man-Month: Essays on Software Engineering.* Addison-Wesley, 1975; 20th Anniversary Edition, 1995 (includes "No Silver Bullet" and "'No Silver Bullet' Refired"). — The tar pit; Brooks's Law; the second-system effect; the programming systems product.
- Brooks, F. P. *The Design of Design: Essays from a Computer Scientist.* Addison-Wesley, 2010. — Conceptual integrity, the one-mind rule, the divorce of design from implementation.
- Brooks & Blaauw. *Computer Architecture: Concepts and Evolution.* Addison-Wesley, 1997.
- [Wikipedia — No Silver Bullet](https://en.wikipedia.org/wiki/No_Silver_Bullet) · [The Mythical Man-Month](https://en.wikipedia.org/wiki/The_Mythical_Man-Month) · [Brooks's law](https://en.wikipedia.org/wiki/Brooks%27s_law) · [Second-system effect](https://en.wikipedia.org/wiki/Second-system_effect) · [The Design of Design](https://en.wikipedia.org/wiki/The_Design_of_Design)

**References — the disagreement and the state of the argument:**

- Ben Moseley & Peter Marks. [Out of the Tar Pit](https://curtclifton.net/papers/MoseleyMarks06a.pdf). BCS Software Practice Advancement, 2006. — "Complexity is the single major difficulty in the successful development of large-scale software systems." Disagrees with Brooks's premise that most complexity is essential; state and control as the major accidental causes.
- Edsger W. Dijkstra. [The Humble Programmer](https://www.cs.utexas.edu/users/EWD/transcriptions/EWD03xx/EWD340.html). Turing Award Lecture, 1972. — "The tools we use have a profound (and devious!) influence on our thinking habits."
- Edsger W. Dijkstra. [On the cruelty of really teaching computing science](https://www.cs.utexas.edu/~EWD/transcriptions/EWD10xx/EWD1036.html). EWD1036, 1988. — "Simplicity is a great virtue... complexity sells better."
- Joel Spolsky. [The Law of Leaky Abstractions](https://www.joelonsoftware.com/2002/11/11/the-law-of-leaky-abstractions/). Joel on Software, 2002. — "All non-trivial abstractions, to some degree, are leaky."
- Rich Hickey. [Simple Made Easy](https://www.infoq.com/presentations/Simple-Made-Easy/). Strange Loop, 2011 / InfoQ. — The distinction between simple and easy; complexity as interleaving of things.
- John Ousterhout. [A Philosophy of Software Design](https://web.stanford.edu/~ouster/cgi-bin/book.php). Yaknyam Press, 2018. — "Complexity is anything related to the structure of a software system that makes it hard to understand and modify the system."
- Nancy G. Leveson & Clark S. Turner. [An Investigation of the Therac-25 Accidents](http://sunnyday.mit.edu/papers/therac.pdf). *IEEE Computer* 26(7), 1993. — The cost of mishandled state in a safety-critical system.
- Melvin E. Conway. [How Do Committees Invent?](http://www.melconway.com/Home/Committees_Paper.html). *Datamation*, 1968. — Conway's Law: systems mirror the communication structures that build them.
- [C2 Wiki — Essential Vs Accidental Complexity](http://wiki.c2.com/?EssentialVsAccidentalComplexity).

**References — accidental complexity in the modern era:**

- [The left-pad incident](https://en.wikipedia.org/wiki/Left-pad_incident). npm, March 2016 — an eleven-line package's removal broke thousands of projects; an accident of the packaging layer.
- [Goodbye Microservices: From 100s of problem children to 1 superstar](https://segment.com/blog/goodbye-microservices/). Segment Engineering, 2018 — moving back from a microservices zoo to a modular monolith to remove distributed-systems accidents.
- Lehman, M. M. [Lehman's laws of software evolution](https://en.wikipedia.org/wiki/Lehman%27s_laws_of_software_evolution). — "Complexity increases unless work is done to reduce it"; the evolution-side formulation of Brooks's argument.

**Related on this site:**

- [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) — the property that makes a system feel like one mind designed it.
- [Taste as Conceptual Integrity](https://blog.hackspree.com/#taste-conceptual-integrity-brooks) — why "taste" is the perception of coherence, and the human skill the AI era leaves standing.
- [Lehman's Laws](https://blog.hackspree.com/#lehmans-laws) — complexity accumulates unless you fight it; Brooks and Lehman, same argument.
- [Dark Factory Complexity](https://blog.hackspree.com/#dark-factory-complexity) — AI removes implementation accidents and pushes the complexity upstream and downstream.
- [Simplicity Is the Prerequisite for Reliability](https://blog.hackspree.com/#simplicity-reliability-dijkstra) — Dijkstra's argument that simplicity is structural, not aesthetic.
- [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) — every dependency is a boundary where a substitution can hide; subtraction as the only defense that compounds.
- [Unix Philosophy](https://blog.hackspree.com/#unix-philosophy) — the original discipline of doing one thing well.
