---
title: "Kevlin Henney: the books, the talks, and why modularity is not microservices"
date: 2026-07-11
slug: kevlin-henney
summary: "Kevlin Henney writes, speaks, and thinks about code, patterns, and modularity. His work is the theoretical foundation most microservices adopters skipped. Time to read it."
tags: kevlin-henney, software-design, patterns, modularity, microservices
---

Kevlin Henney is the kind of thinker the software industry doesn't produce enough of: someone who has written production code since the late 1980s, who co-authored two volumes of the POSA pattern series, who edited the *97 Things* books, and who gives conference talks that get half a million views because they're *right* — not slick, not loud, just correct at a level that makes you uncomfortable about your own code.

He lives in Bristol. He consults, trains, writes, and speaks. His interests span programming, people, and practice. If you've watched a software conference talk on YouTube, his name has been in the sidebar. If you haven't clicked it, this post is the nudge. If you're building microservices, this post is mandatory reading. Most of what passes for microservices architecture violates principles Henney documented decades ago. The industry skipped the foundations and built the tower. The tower is leaning.

## The architecture of modularity

Henney's work on patterns and modularity forms the theoretical spine that microservices — done well — rest on. The problem is that most microservices adopters never read the theory. They split services by database table, called it "bounded context," and declared victory. Henney would not be impressed.

> "The modularity of a system cannot be evaluated by examining designs of individual modules in isolation. The goal of modular design is to simplify the relationships between components of a system."

This is the sentence that indicts most microservices architectures. Teams design services in isolation — one team per service, each optimizing locally — and then wonder why the system as a whole is a distributed mess. The modularity is not in the services. The modularity is in the relationships *between* the services. If you have twelve services that all call each other in a complete graph, you do not have a modular system. You have a distributed monolith with network latency and serialization overhead. The modularity didn't improve. The deployment got harder.

> "Coupling is the aspect of a system that defines what knowledge is shared between components of a system. Different ways of coupling components share different types and amounts of knowledge. Some will increase complexity, while others will contribute to modularity."

Coupling is shared knowledge. Not shared code. Not shared libraries. Shared *knowledge*. When Service A knows the internal data format of Service B, they are coupled — regardless of whether they communicate via HTTP, gRPC, or a message queue. The transport is irrelevant. The knowledge is the coupling. If Service B changes its schema and Service A breaks, the coupling is real. If Service B changes its schema and Service A doesn't notice, the coupling was imaginary. Most organizations don't know which kind they have because they never tested. They will find out in production.

> "The aggressive pursuit of LCHC (low coupling and high cohesion) can ensure that the effect of change is simplified and isolated, rather than traumatic and global. LCHC also simplifies testing, building, versioning, experimentation, optimization, team organization, and pretty much any other development activity you can think of."

LCHC is the goal. Microservices are one mechanism for achieving it — but only if you get the boundaries right. The boundaries are everything. Henney, channeling Grady Booch, defines architecture precisely:

> "All architecture is design but not all design is architecture. Architecture represents the significant design decisions that shape a system, where significant is measured by cost of change."

The cost of change. That is the criterion. If changing a decision requires rewriting one service, it was a local design decision. If changing it ripples across seven services, it was an architectural decision — and you either got the boundary wrong or you failed to protect it. Most microservices failures are boundary failures. The services work individually. The boundaries between them are in the wrong places. When the business requirement changes — and it will — the change crosses five service boundaries and requires coordinated deploys. The architecture didn't reduce the cost of change. It increased it. You built a distributed system to make changing your mind *more* expensive. That is not progress.

> "It is possible to simplify the structure of software without losing effective options. It is even possible to do so and increase your options. Now, that sounds worthwhile: simpler and more flexible."

Simpler and more flexible. This is the test. After your microservices migration, is the system simpler and more flexible than the monolith it replaced? If the answer is no — and it often is — you didn't improve the modularity. You just distributed the monolith across a network. The modularity was the goal. The distribution was a side effect. Confusing them is the central mistake of the microservices era.

## Architecture as managed uncertainty

One of Henney's deepest insights is about architectural decision-making under uncertainty — and it directly applies to the "how many services?" question that paralyzes microservices adopters.

> "When a design decision can reasonably go one of two ways, an architect needs to take a step back. Instead of trying to decide between options A and B, the question becomes 'How do I design so that the choice between A and B is less significant?' The most interesting thing is not actually the choice between A and B, but the fact that there is a choice between A and B."

The existence of a choice is a signal. It means the boundary is volatile. Instead of picking A or B, design the interface so that A and B are interchangeable behind it. This is Parnas's information hiding applied to architectural decisions: hide the decision behind a stable interface, and the choice becomes less significant because you can change your mind later without telling anyone.

Applied to microservices: the question is not "should this be one service or two?" The question is "how do I design the interface so that whether it's one or two doesn't matter to the callers?" If you can switch from one to two without the callers changing, you got the interface right. If a service split forces every caller to update, the interface was wrong. The interface was coupled to the deployment topology. That coupling is the problem. The number of services is a distraction.

> "Although we cannot predict the future with any certainty, it is still possible to write code that is graceful and accommodating — rather than troublesome and resistant — in the face of change."

Graceful in the face of change. That is the definition of good architecture. Not "scales to a million users." Not "uses the latest framework." Not "has a beautiful diagram." Can you change your mind about a decision without rewriting the system? If yes, the architecture is working. If no, the architecture is decoration.

## Simplicity before generality

Henney's most quoted principle is a direct challenge to the way most platforms and frameworks are built — and to the way most microservices platforms encourage over-generalization.

> "The best route to generality is through understanding known, specific examples and focusing on their essence to find an essential common solution. Simplicity through experience rather than generality through guesswork."

Build the specific thing first. Make it work. Then, when you have a second specific thing that resembles the first, find the common essence and extract it. This is the opposite of how most microservices platforms are designed. The platform team builds a general-purpose service template with twenty configuration options, five deployment modes, and an abstraction layer that handles every possible use case — most of which never occur. The template is general. It is also heavy, complex, and wrong in ways that only become visible when you try to use it for something the designers didn't anticipate.

> "A common problem in component frameworks, class libraries, foundation services, and other infrastructure code is that many are designed to be general purpose without reference to concrete applications. This leads to a dizzying array of options and possibilities that are often unused, misused, or just not useful."

> "Favoring simplicity before generality acts as a tiebreaker between otherwise equally viable design alternatives. When there are two possible solutions, favor the one that is simpler and based on concrete need rather than the more intricate one that boasts of generality."

Simplicity first. Generality second. Use before reuse. This is not laziness. It is epistemological humility. You don't know what the general case looks like until you've seen several specific cases. Guessing at generality produces abstractions that fit nothing well and everything poorly. The microservices platform that was designed before any services existed? That's not architecture. That's prophecy. Prophecy is usually wrong.

> "The trick to achieving generality is, somewhat counterintuitively, to make the code specific enough to be fit for purpose."

Specific enough to work. General enough to adapt. Finding that line is the discipline. Most teams err on one side or the other — over-specific code that hardcodes assumptions and breaks when anything changes, or over-general frameworks that handle everything and solve nothing. The middle is where Henney lives. It's harder to get to. It's worth it.

## The seven ineffective habits

Henney's most-watched talk examines coding habits acquired by imitation and never questioned. Each habit has architectural consequences when scaled to a distributed system.

### 1. Noisy code

> "Comments should provide additional information that is not readily obtainable from the code itself. They should never parrot the code."

At the microservices level: documentation that repeats what the API already says. Swagger docs that restate the endpoint name. Architecture decision records that say "we chose Kafka because it scales." The signal-to-noise ratio is the same whether you're looking at a function or a system. Noise is noise. It obscures. It doesn't inform.

> "A common fallacy is to assume authors of incomprehensible code will somehow be able to express themselves lucidly and clearly in comments."

The service equivalent: assuming that a team that can't design a clean module boundary will somehow design a clean service boundary. The scale changes. The skill doesn't. If you can't modularize a monolith, you can't modularize microservices. You'll just distribute the spaghetti.

### 2. Unsustainable spacing / visual dishonesty

Code layout must reflect code structure. If the indentation lies, the reader is confused. At the system level: if the service diagram shows clean boundaries but the runtime shows a complete call graph, the diagram lies. The architecture document is visual dishonesty. The system's actual structure is in the call graph, not the whiteboard.

### 3. Lego naming

> "More words is not more meaning."

`CustomerProfileServiceManagerFactory`. `OrderProcessingOrchestrationController`. Lego naming at the service level produces names that catalogue rather than distinguish. A service name should tell you what domain function it performs, not what design patterns it implements. The patterns are implementation details. The domain function is the interface.

### 4. Under-abstraction

Primitive obsession at the function level becomes primitive obsession at the service level. Services that exchange raw JSON with no schema. Services that pass database IDs rather than domain references. Services that expose their internal table structure through "REST" APIs that are really thin wrappers over CRUD. If your service API looks like your database schema with HTTP verbs, you have not abstracted. You have exposed. That's not an API. That's a tunnel into your storage. The coupling is total.

### 5. Unencapsulated state

Henney's analogy: wearing your underwear on the outside. At the service level: services that share a database. Services that read each other's tables. Services that depend on another service's internal state transitions rather than its published events. If Service A queries Service B's database directly, Service B has no encapsulation. Its state is public. Any change to B's schema breaks A. The services are not independent. They are co-dependent with extra network hops.

### 6. Getters and setters

> "When it is not necessary to change, it is necessary not to change." — Lucius Cary, via Henney

Mindless getter/setter generation becomes mindless CRUD endpoint generation. `GET /customers/{id}`, `PUT /customers/{id}`, `PATCH /customers/{id}`, `DELETE /customers/{id}`. Generated. Never questioned. The API exposes mutation when mutation should be impossible (do you really need to DELETE a customer, or do you need to *close an account*?). The API exposes internal structure because the framework generated it from the entity. This is not REST. This is an ORM with a public URL.

### 7. Uncohesive tests

> "For tests to drive development they must do more than just test that code performs its required functionality: they must clearly express that required functionality to the reader." — Nat Pryce & Steve Freeman

One test class per production class becomes one integration test per service. Tests that verify internal implementation rather than external behavior. When you refactor a service's internals, the tests break — not because the behavior changed, but because the tests were coupled to the implementation. The same anti-pattern, distributed. Test behavior. Not structure. At every scale.

## Code is not an asset

One of Henney's most bracing arguments:

> "Less code, more software. Less code = less bugs."

> "Programmers write code: a formal plan of the software, expressing its intent in maximal detail. Software is the end product: in execution it is what the user perceives, interacts with and experiences. Sometimes this difference can be significant."

Code is the plan. Software is the product. Confusing them — treating the plan as the asset — is why organizations measure productivity in lines of code and celebrate growing codebases. A growing codebase is not an achievement. It is a growing liability. Every line must be tested, maintained, understood, migrated, and eventually deleted. The asset is the solved problem. The code is what it cost to solve it.

> "Duplicate code has a bad smell and violates the DRY principle... Contrary to benighted management belief, more is not better in this case. Every problem has an optimal code size: too short and the code is cryptic line noise; too long and you cannot see the wood for the trees."

There is an optimal size. Not zero. Not infinite. Somewhere in between, where the code is minimal enough to understand and complete enough to work. Finding that point is the discipline. Most teams never look for it. They add code until the feature works, then stop. The codebase grows monotonically. Nobody's job is deletion. It should be.

## The books

### Pattern-Oriented Software Architecture, Volumes 4 and 5 (2007)

Co-authored with Frank Buschmann and Doug Schmidt. Volume 4: *A Pattern Language for Distributed Computing*. Volume 5: *On Patterns and Pattern Languages*. These are not introductory texts. They are for practitioners who have already encountered the problems that patterns solve and need the vocabulary to reason about them. Volume 4 is directly relevant to anyone building distributed systems — i.e., anyone building microservices. Volume 5 is the meta-level: what makes a pattern a pattern, how patterns compose, how to write them. Henney's contribution places him in the direct lineage of the Gang of Four and the pattern movement. Read Volume 4 before you split another service. The patterns you need are already named and documented. You're discovering them the hard way.

### 97 Things Every Programmer Should Know (2010)

Henney edited this. 97 essays. Two to three pages each. One point per essay. No filler. His own six contributions are worth reading first:

- **"Comment Only What the Code Cannot Say."** The code says what. The comment says why. If the code doesn't say what clearly, fix the code, not the comment. A comment that explains bad code is a commitment to keeping the bad code.

- **"Test for Required Behavior, Not Incidental Behavior."** Test the contract. Not the implementation. If you test the implementation, refactoring becomes a test-breaking exercise and you stop refactoring. That is how codebases rot.

- **"Test Precisely and Concretely."** Vague tests pass when they should fail. Specific tests fail only when something is wrong. A test that passes in the presence of a bug is worse than no test — it gives false confidence.

- **"Program with GUTs."** Good Unit Tests. The name is a pun. The practice is not. Tests are code. They deserve the same care, the same review, the same refactoring. Test quality is system quality.

- **"Uncheck Your Exceptions."** Checked exceptions create coupling across layers. They force catch blocks that do nothing. Henney was arguing this before it was consensus. Now it's obvious. The debate is over. Henney won.

- **"Name the Date."** Don't write `getDate()`. Write `getExpiryDate()`. The extra word is not noise. It is the difference between understanding and guessing. Names are the primary interface between a programmer and a codebase. Every ambiguous name is a small tax paid by every future reader.

### 97 Things Every Java Programmer Should Know (2020)

Co-edited with Trisha Gee. Same format, Java-specific. The editorial voice is the same: practical, opinionated, allergic to received wisdom.

Henney also contributed to *97 Things Every Software Architect Should Know* and wrote columns for *Better Software*, *The Register*, *C/C++ Users Journal*, *JavaSpektrum*, *C++ Report*, *EXE*, *Overload*, and *Application Development Advisor*. If a programming magazine existed in the 1990s or 2000s, he probably wrote for it. Those columns are still sharper than most things published this morning.

## The other talks

Beyond *Seven Ineffective Coding Habits*, Henney's talk catalogue is deep:

**Small Is Beautiful.** Less code is empirically better. Smaller modules, smaller functions, smaller interfaces. The evidence is overwhelming and universally ignored. Measure deletions, not additions. The best commit has more red than green.

**Code as Risk.** Every line is a liability. Code is inventory cost, not an asset. The asset is the solved problem. The code is what you paid. Confusing them is why organizations celebrate growing codebases. Growth is not progress if the growth is in inventory, not in solved problems.

**Old Is the New New.** Microservices, event sourcing, functional programming — the ideas that are "new" are old ideas with new names. The industry doesn't lack ideas. It lacks memory. The original papers are still correct. The original authors are still worth reading. Henney traces modern trends back to their 1970s sources. It is humbling. It should be.

**The Way the Future Was.** A meditation on prediction, kept and broken. The promises the industry made. The ones it kept. The ones it forgot it made. Henney has been building software since the late 1980s. He saw the promises. He remembers which ones were real.

**Question-Led Development.** Asking better questions rather than jumping to answers. Most design debates are not about tradeoffs. They are about people who haven't asked the same question answering different questions and wondering why they disagree. Clarify the question. The answer often follows.

**The Architecture of Uncertainty.** How to design when you don't know what will change. The answer: make the choice less significant by hiding it behind a stable interface. This is Parnas updated for the distributed systems era. It is also the best advice you will get about microservices boundaries.

## Why Henney is the prerequisite for microservices

The microservices movement adopted Conway's Law and forgot Parnas. It adopted bounded contexts and forgot coupling. It adopted independent deployability and forgot that deployment independence requires interface stability — and interface stability requires getting the boundaries right.

Henney's work connects all of this. Patterns give you the vocabulary for boundaries. Information hiding gives you the criterion for where to put them. LCHC gives you the measure of whether they work. Architecture as cost-of-change gives you the test of whether the system improved. Simplicity before generality gives you the discipline to stop at the right size.

Most microservices adopters never read any of this. They split services by entity, called it domain-driven, and moved on. The system is more complex, harder to change, and harder to understand than the monolith it replaced. The deployment is independent. Nothing else is. That is failure dressed in modern infrastructure. Henney's work explains why. It also explains what to do instead. The books are short. The talks are free. The ignorance is expensive.

---

**References:**
- Kevlin Henney, Frank Buschmann, Doug Schmidt, *Pattern-Oriented Software Architecture, Volumes 4 & 5*, Wiley, 2007.
- Kevlin Henney (ed.), *97 Things Every Programmer Should Know*, O'Reilly, 2010.
- Kevlin Henney and Trisha Gee (eds.), *97 Things Every Java Programmer Should Know*, O'Reilly, 2020.
- *Seven Ineffective Coding Habits of Many Programmers*, NDC/BuildStuff, 2014. [InfoQ](https://www.infoq.com/presentations/7-ineffective-coding-habits/)
- *The Architecture of Uncertainty*, Agile Singapore, 2013.
- *Simplicity Before Generality, Use Before Reuse*, [Artima](https://www.artima.com/weblogs/viewpost.jsp?thread=351149), 2012.
- *From Mechanism to Method: Generic Decoupling*, [Overload 60](https://accu.org/journals/overload/12/60/henney_308), 2004.
- GOTO Conferences YouTube — [Kevlin Henney](https://www.youtube.com/@GOTO-)
