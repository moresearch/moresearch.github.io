---
title: "CUPID: Properties, Not Principles"
date: 2026-08-03
slug: cupid-for-joyful-coding
summary: "Dan North's CUPID looks like five principles to replace SOLID's five. That is the misreading. CUPID's real move is replacing the idea of a principle itself — bounded sets of rule-followers become centred sets with a direction of travel. Read as rules, CUPID is SOLID with a better haircut. Read as properties, it is a compass."
tags: dan-north, cupid, solid, software-design, principles, properties, composability, unix-philosophy, predictability, idiomatic-code, domain-driven-design, joyful-coding, simplicity, empathy
---

Dan North's [*CUPID: for joyful coding*](https://dannorth.net/blog/cupid-for-joyful-coding/) (2022) starts with the reader's body memory. As a rookie he cracked open a large C codebase, expecting to drown; within minutes he was deep in a nest of calls and knew exactly where the bug was. Structure, naming, and flow so obvious they felt like architecture he already knew.

The vocabulary escalates: Fowler's "code that humans can understand" is too low a bar; Gabriel's *habitability* — "to change it comfortably and confidently" — gets closer; the word North actually wants is *joyful*. If you work in code, the codebase is your user experience, programmed in by people you have never met, one of whom may be future you.

CUPID is his account of what makes code joyful, a five-letter backronym aimed at SOLID: **Composable**, **Unix philosophy**, **Predictable**, **Idiomatic**, **Domain-based**. It is usually filed under "the anti-SOLID," and that is the misreading. CUPID's contribution is not five new principles. It is the argument that the idea of a principle is the problem.

## The move: properties over principles

North started by trying to replace each SOLID letter with a better one and quickly concluded the frame itself was broken:

> "Principles are like rules: you are either compliant or you are not. This gives rise to 'bounded sets' of rule-followers and rule-enforcers rather than 'centred sets' of people with shared values."

Bounded set: in or out, compliant or violating, pass or fail. Centred set: closer or further, with a direction of travel that is always clear, and no one ever "outside." Principles give you judgment day; properties give you a compass. They are chosen to be practical ("there is never a 'done'"), human ("what it *feels like* to work with code"), and layered (obvious to a beginner, deep for the experienced). Every letter is best read as a direction, not a target.

## Composable — plays well with others

Three heuristics, none of them laws. **Small surface area:** a narrow, opinionated API has less to learn and less to go wrong — but too narrow, and "knowing the right combination" becomes tacit knowledge; there is a sweet spot between fragmented and bloated. **Intention-revealing:** a component you can discover and assess quickly — the tutorial ladder of 2, 10, and 30 minutes. **Minimal dependencies:** "a logging library" is really a dependency on a specific version, and version is where incompatibilities break.

> "More is not necessarily better; it is all trade-offs."

A principle says *minimize dependencies*. A property says *dependencies are a cost you can move, and the skill is knowing which way lowers the total*. The difference is whether the guidance survives contact with a real codebase.

## Unix philosophy — does one thing well, from the outside

"Doing one thing well" sounds like the Single Responsibility Principle, and North is careful about why it is not:

> "The former is about how you use code, and the latter is about the internals of the code itself."

SRP is inside-out: "one and only one reason to change." Unix is outside-in: a specific, comprehensive purpose visible from the call site. `ls` does not know anything about files — `stat` provides the data, `ls` only renders it; pipes compose such commands into pipelines, each a narrow, complete, outside-visible contract.

Then the attack on SRP, the most useful passage in the essay. "One reason to change" is trivially refutable — a single line changes for security, compliance, dependencies, operations. The real damage is the *artificial seams*: report content and format change together, so SRP's demand to separate them makes every new field a chore of chaining identical fields across files; UI components suffer the same split. Applied as a rule, SRP imports accidents — seams nobody's problem forced on you.

> The inside-out question ("what could make this change?") generates seams. The outside-in question ("what would a user of this call it?") generates boundaries. Seams are things to maintain. Boundaries are things to use.

## Predictable — does what you expect

Predictability is "a generalization of testability": code should **behave as expected**, be **deterministic**, and be **observable**. Behave as expected — the first of Kent Beck's four rules, holding even with no tests: the intended behaviour is obvious from structure and naming. Deterministic splits into robust (covers what you know), reliable (same result every time), resilient (survives what you don't). Observable is the control-theory word — internal state inferable from outputs — only possible if designed in. North's ladder: instrumentation, telemetry, monitoring, alerting. "Most software does not even get past step 1."

## Idiomatic — feels natural, because empathy

"The greatest programming trait is empathy; empathy for your users; empathy for support folks; empathy for future developers; any of whom may be future you." Idiomatic code is empathy compiled into style — matching the language's idioms so the reader's context switches are free. Opinionated languages help (Go's `gofmt` makes all code look the same; Python's Zen: "one—and preferably only one—obvious way to do it"); multi-paradigm ones (Perl's TIMTOWTDI, Ruby, JavaScript) let five ways to iterate a sequence coexist, each adding cognitive load. Where the language has no consensus, the team must supply one: shared formatting, linting, Architecture Decision Records.

> "Your learning curve for a technology will likely be shorter-lived than any code you write in it."

Read it again: the person writing the code is a temporary visitor to its style; the code outlives them. "Reads well to me right now" is the wrong bar — North calls idiomatic writing "writing code for someone else."

## Domain-based — the solution models the problem

The last property is the deepest: the solution domain should model the problem domain in language and structure. A surname is not a `string[30]`; money is not a float. Type the domain — `Surname`, `Money` with its `Currency` and `Amount` — and the cognitive distance between what you write and what it does collapses. North's criterion, stated once:

> "A casual observer cannot tell whether people are discussing the code or the domain."

Structure next, where CUPID is most contrarian. Framework scaffolds impose an *a priori* structure — the Rails skeleton's `app/models`, `app/views`, `app/controllers`, `app/helpers`, `app/jobs` — scattering one semantic unit across half a dozen directories: a patient-record change touches a model, a view, a controller, a helper, each in a different folder. North's proposal is radical in its mildness: structure by the domain, not the framework — `patient_history`, `appointments`, `staffing`, `compliance` as the top level. A codebase grouped by framework role is a codebase whose architecture was decided by the framework's author, not by the one mind that understands the domain.

## Why it holds together

The properties are mutually reinforcing: composable and comprehensive — doing one thing well — "is like a reliable friend"; idiomatic code "feels familiar even though you have never seen it before"; predictable code "gives you spare cycles to concentrate on surprises elsewhere"; domain-based code "minimises the cognitive distance from need to solution."

> "Moving code towards the 'centre' of any of these properties leaves it better than you found it."

That is the whole philosophy in one sentence — and the sentence that cannot survive translation back into principles: "leaves it better than you found it" has no compliance check. You can only point at a codebase and say: closer to the centre than last month, and here is the direction to keep travelling.

## The test

Read as rules, CUPID is SOLID with a better haircut: five more checkboxes, five more fights in code review. Read as properties, it is a different instrument. Five questions, no pass or fail:

- What would it take to reuse this outside its home? (Composable)
- What is the one thing this does, and can I see that from outside? (Unix philosophy)
- If I change this, how hard is it to know what happens next? (Predictable)
- Would the person who inherits this recognize it as theirs? (Idiomatic)
- Does this read like the domain, or like the framework? (Domain-based)

None of these has a yes-or-no answer. That is not a weakness. It is the point.

> SOLID asks: are you compliant? CUPID asks: which way is better? A principle is a verdict. A property is a compass. Joy is not a property you can check off — it is what it feels like to be moving in the right direction.

---

**References:**

- Dan North. [CUPID: for joyful coding](https://dannorth.net/blog/cupid-for-joyful-coding/). Feb 10, 2022. — The five properties: Composable, Unix philosophy, Predictable, Idiomatic, Domain-based; properties over principles; the three "properties of properties" (practical, human, layered).
- Martin Fowler. [Refactoring: Improving the Design of Existing Code](https://martinfowler.com/books/refactoring.html). Addison-Wesley, 1999. — "Any fool can write code that a computer can understand. Good programmers write code that humans can understand."
- Richard P. Gabriel. [Habitability and Piecemeal Growth](https://www.dreamsongs.com/Files/PatternsOfSoftware.pdf), in *Patterns of Software*. Oxford University Press, 1996. — "Habitability is the characteristic of source code that enables people to understand its construction and intentions and to change it comfortably and confidently."
- Michael Feathers. [Working Effectively with Legacy Code](https://www.informit.com/store/working-effectively-with-legacy-code-9780131177055). Prentice Hall, 2004. — Characterization tests; "When a system goes into production, in a way, it becomes its own specification."
- [The Zen of Python, PEP 20](https://peps.python.org/pep-0020/). — "There should be one—and preferably only one—obvious way to do it."
- [Effective Go](https://go.dev/doc/effective_go) and [gofmt](https://go.dev/blog/gofmt). — Idiomatic Go as a shipped, enforced standard.
- [TIMTOWTDI](https://en.wikipedia.org/wiki/There%27s_more_than_one_way_to_do_it). — Perl's "there is more than one way to do it."
- Michael Nygard. [Documenting Architecture Decisions](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions). 2011. — Architecture Decision Records for local idioms.
- Robert C. Martin. *Agile Software Development, Principles, Patterns, and Practices*. Prentice Hall, 2002. — The Single Responsibility Principle ("one and only one reason to change"), the inside-out view North argues against.
- Dan North. [Code in the Language of the Domain](https://97-things-every-x-should-know.gitbooks.io/97-things-every-programmer-should-know/content/en/thing_14/), in *97 Things Every Programmer Should Know*. O'Reilly, 2010.
- [Ruby on Rails](https://rubyonrails.org/) — the generated skeleton app whose type-based directory structure North critiques from a domain-based view.

**Related on this site:**

- [The Unix Philosophy Is the Only Software Engineering Theory That Works](https://blog.hackspree.com/#unix-philosophy) — the outside-in view, taken seriously as a theory.
- [Accidental Complexity Is the Only Complexity You Can Remove](https://blog.hackspree.com/#brooks-accidental-complexity) — SRP's artificial seams as imported accidents; framework scaffolding as the tar pit.
- [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) — minimal dependencies as the only defense that compounds.
- [Taste as Conceptual Integrity](https://blog.hackspree.com/#taste-conceptual-integrity-brooks) — idiomatic code as one-mind judgment, written for the person who inherits it.
- [Simplicity Is the Prerequisite for Reliability](https://blog.hackspree.com/#simplicity-reliability-dijkstra) — predictability and observability as structural, not aesthetic.
