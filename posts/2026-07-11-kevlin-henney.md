---
title: "Kevlin Henney: the books and the talks"
date: 2026-07-11
slug: kevlin-henney
summary: "Kevlin Henney writes, speaks, and thinks about code as if it matters — because it does. A tour through his books, his best-known talks, and the recurring ideas that make him essential."
tags: kevlin-henney, software-design, patterns, coding-habits, talks
---

Kevlin Henney is the kind of thinker the software industry doesn't produce enough of: someone who has written production code since the late 1980s, who co-authored two volumes of the POSA pattern series, who edited the *97 Things* books that sit on thousands of desks, and who gives conference talks that get half a million views not because they're slick but because they're *right*.

He lives in Bristol. He consults, trains, writes, and speaks. His development interests are programming, people, and practice — patterns, architecture, languages, and processes. If you've watched a software conference talk on YouTube, you've probably seen his name in the sidebar. If you haven't clicked it yet, this post is your nudge.

## The books

### Pattern-Oriented Software Architecture, Volumes 4 and 5 (2007)

Co-authored with Frank Buschmann and Doug Schmidt, these two volumes extend the POSA series — the definitive catalogue of architectural and design patterns for software systems. Volume 4 covers *A Pattern Language for Distributed Computing*. Volume 5 is *On Patterns and Pattern Languages* — the meta-level work on what patterns are, how they compose, and how to write them.

These are not books you read cover to cover. They are books you consult when you're building something distributed and you want to know what patterns have already been discovered, named, and documented by people who learned the hard way. Henney's contribution to the POSA series places him in the direct lineage of the Gang of Four, Buschmann, and the pattern movement that shaped how we talk about software architecture.

### 97 Things Every Programmer Should Know (2010)

Henney edited this — and it is a masterclass in editing. 97 short essays from practitioners including Uncle Bob Martin, Scott Meyers, Dan North, Linda Rising, Neal Ford, Dave Farley, and many others. Each essay is two to three pages. Each makes one point. None wastes your time.

Henney's own contributions include:

- **"Comment Only What the Code Cannot Say"** — the code says what it does. The comment should say why. If the code doesn't say what it does clearly, fix the code, not the comment.
- **"Test for Required Behavior, Not Incidental Behavior"** — test the contract, not the implementation. If you test the implementation, you lock it in place and make refactoring a test-breaking exercise.
- **"Test Precisely and Concretely"** — vague tests pass when they should fail. Specific tests fail only when something is actually wrong.
- **"Program with GUTs"** — Good Unit Tests. The acronym is memorable. The practice is not.
- **"Uncheck Your Exceptions"** — checked exceptions sound like a good idea. In practice, they create coupling across layers and force catch blocks that do nothing. Henney was saying this before it was popular. Now it's obvious.
- **"Name the Date"** — don't write `getDate()`. Write `getExpiryDate()` or `getPublicationDate()`. The extra word is not noise. It is the difference between understanding and guessing.

### 97 Things Every Java Programmer Should Know (2020)

Co-edited with Trisha Gee. Same format, Java-specific lens. Henney's editorial voice — practical, opinionated, allergic to dogma — shapes both volumes. If you've read one, you know the format. If you haven't, you're missing the most efficient way to absorb collective wisdom from several dozen experienced programmers in one sitting.

Henney also contributed to *97 Things Every Software Architect Should Know*.

### Other writing

He has been a columnist for *Better Software*, *The Register*, *C/C++ Users Journal*, *JavaSpektrum*, *C++ Report*, *EXE*, *Overload*, and *Application Development Advisor*. If there was a magazine about programming in the 1990s or 2000s, Henney probably wrote for it. Most of those columns are still sharper than the average blog post published this morning.

## The talks

Henney's talks are the core of his public presence. He has over 500,000 views on the GOTO Conferences channel alone. He speaks at NDC, BuildStuff, ACCU, C++Online, OOP, and dozens of other conferences. His style is dry, precise, and quietly devastating. He doesn't shout. He doesn't need to. The ideas do the work.

### Seven Ineffective Coding Habits of Many Programmers

His most-watched talk. Originally presented at NDC and BuildStuff in 2014. Still relevant. Still uncomfortable. The seven habits:

1. **Noisy code.** Low signal-to-noise ratio. Comments that say `i = i + 1; /* add one to i */`. Code that buries its intent under boilerplate. The cure: maximize signal, minimize noise. If a comment repeats what the code says, delete the comment.

2. **Unsustainable spacing.** Visual dishonesty. Code whose layout doesn't match its structure. Indentation that lies. Brace placement that obscures. Style is not personal preference. Style is what makes the structure visible to the next reader. The next reader is often you, six months later, at 2am.

3. **Lego naming.** Agglutination. `controllerFactoryFactory`. `processValidateManager`. More words is not more meaning. Henney calls this labeling, not naming. A name should distinguish, not catalogue. If your class name has five words, three of them are probably noise.

4. **Under-abstraction.** Primitive obsession. Ten-parameter functions. Code that talks to data structures instead of domain concepts. Alan Perlis: "If you have ten parameters, you probably missed some." Henney: use a tag cloud on your codebase. If your domain words don't appear, your abstraction level is wrong.

5. **Unencapsulated state.** Mutable internals exposed. "Wearing your underwear on the outside." Your API should make correct usage easy and incorrect usage impossible. If someone can corrupt your object's state from outside, they will. Not out of malice. Out of not knowing the rules. If the rules aren't enforced by the interface, they don't exist.

6. **Getters and setters.** Mindlessly generated by IDEs. `getX()` and `setX()` as a reflex. In English, "get" implies side effects — "I get married" changes your state. Programming borrowed the word and used it for side-effect-free queries. The mismatch is not trivial. "When it is not necessary to change, it is necessary not to change." — Lucius Cary, via Henney.

7. **Uncohesive tests.** One test class per production class. `TestFoo` for `Foo`. Tests that mirror method structure instead of expressing behavior. Complexity arises from how methods interact, not from individual methods. Tests should correspond to scenarios, not to methods. Nat Pryce and Steve Freeman: "For tests to drive development they must do more than just test that code performs its required functionality: they must clearly express that required functionality to the reader."

> "A common fallacy is to assume authors of incomprehensible code will somehow be able to express themselves lucidly and clearly in comments."

### Small Is Beautiful

Henney argues that small code is not just aesthetically pleasing — it is less likely to contain bugs, easier to test, easier to delete, and easier to understand. The evidence for this is overwhelming and universally ignored. We measure lines of code written. We should measure lines of code deleted. The best commit you can make is the one with more red than green.

The argument connects to the broader minimalism that runs through Henney's work. Small functions. Small classes. Small modules. Small interfaces. Every boundary is a place where complexity can enter. Every additional line is a place where a bug can hide. You cannot eliminate bugs by adding code. You can only eliminate them by removing the code they live in.

### Code as Risk

Every line of code is a liability. It must be tested, maintained, understood, migrated, and eventually deleted. The more code you have, the more risk you carry. Code is not an asset. It is an inventory cost. The asset is the solved problem. The code is what you paid to solve it. Confusing the two is why organizations treat growing codebases as progress. Progress is solving more problems with less code. Everything else is accumulation.

Henney connects this to technical debt but goes further. Technical debt implies you borrowed against the future and will pay it back. Henney's point is darker: most code is not debt. It is inventory. It sits there, depreciating, requiring maintenance, generating no value. You're not paying it back. You're paying to keep it around.

### Old Is the New New

The ideas that work are old. The ideas that are trendy are often old ideas with new names. Henney traces modern "innovations" back to papers from the 1970s, 1960s, and earlier. Microservices? Distributed computing patterns, documented decades ago. Event sourcing? Older than most programmers. Functional programming? Lisp is from 1958. The industry doesn't lack ideas. It lacks memory.

This talk is a corrective to the amnesia that drives software hype cycles. Every generation rediscovers the same principles and gives them new names. Henney doesn't mock this. He just points at the original sources and suggests reading them. The past is not obsolete. It is under-read.

### The Way the Future Was

A meditation on what we thought software would become versus what it became. The predictions that were wrong. The predictions that were right for reasons nobody expected. Henney's historical depth — he has been in the industry since the late 1980s — gives him perspective that younger speakers can't replicate. He saw the promises made. He saw which ones were kept.

Henney also speaks on Question-Led Development (asking better questions rather than jumping to answers), on software architecture as a verb not a noun, on the relationship between code and design, and on the craft of programming as a discipline that rewards deliberate practice. His conference schedule is a tour of the ideas that matter, delivered by someone who has thought about them longer than most developers have been alive.

## The recurring ideas

Across the books and the talks, a few themes recur:

**Code is communication.** It communicates to the machine, yes. But more importantly, it communicates to other programmers. If they can't read it, it doesn't matter that it works. Readability is not a nice-to-have. It is the primary property of maintainable software.

**Small is better.** Every line carries cost. Less code is less risk, less maintenance, less surface area for bugs. The measure of productivity is not lines written. It is problems solved per line.

**Names matter.** Naming is the hardest problem in computer science for a reason. A good name distinguishes, clarifies, and lasts. A bad name obscures, confuses, and spreads — because every reference to the thing reinforces the wrong name.

**History matters.** The answers to most current questions were published decades ago. The industry's amnesia is expensive. Reading old papers and old books is not nostalgia. It is efficiency. Why rediscover what Dijkstra already explained in 1972?

**Habits are invisible until examined.** Most programmers code the way they do because they learned it by imitation and never questioned it. Henney's method is to surface the habit, name it, and ask: does this actually help? Usually, the answer is no. Usually, the habit persists because nobody asked.

**Test behavior, not implementation.** Tests that mirror code structure are fragile. Tests that express required behavior are durable. The difference is whether you'd keep the test if you completely rewrote the implementation. If the answer is no, the test is coupled to the implementation. Decouple it or delete it.

## Why Henney matters

The software industry has no shortage of loud voices. It has a shortage of precise ones. Henney's talks don't rely on theatrics. They rely on clear thinking, historical depth, and an eye for the habit you didn't know you had. His books are edited with the same rigor — every essay earns its place, every sentence carries weight.

He is not famous in the way that tech celebrities are famous. He is famous in the way that a good teacher is famous: the people who know his work know it deeply. If you've never watched one of his talks, start with *Seven Ineffective Coding Habits*. If you've never read one of his books, start with *97 Things Every Programmer Should Know*. Both will take less than an hour. Both will change how you see your own code. That is the definition of time well spent.

---

**References:**
- Kevlin Henney and Frank Buschmann, Doug Schmidt, *Pattern-Oriented Software Architecture, Volumes 4 & 5*, Wiley, 2007.
- Kevlin Henney (ed.), *97 Things Every Programmer Should Know*, O'Reilly, 2010.
- Kevlin Henney and Trisha Gee (eds.), *97 Things Every Java Programmer Should Know*, O'Reilly, 2020.
- *Seven Ineffective Coding Habits of Many Programmers*, NDC/BuildStuff, 2014. [InfoQ](https://www.infoq.com/presentations/7-ineffective-coding-habits/)
- GOTO Conferences YouTube channel — [Kevlin Henney talks](https://www.youtube.com/@GOTO-)
