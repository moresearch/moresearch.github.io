---
title: Henney's Microservices
date: 2026-07-11
slug: kevlin-henney
summary: "Kevlin Henney never wrote a microservices book. He wrote the books you should have read before adopting microservices. Coupling is shared knowledge. Architecture is measured by cost of change. Modularity is in the relationships, not the services."
tags: kevlin-henney, microservices, modularity, coupling, software-architecture
---

Kevlin Henney never wrote a book called *Building Microservices*. He wrote the books you should have read *before* adopting microservices — the POSA pattern volumes, *97 Things Every Programmer Should Know*, and two decades of columns and talks on modularity, coupling, and the cost of change.

The microservices movement adopted Conway's Law and forgot Parnas. Adopted bounded contexts and forgot coupling. Adopted independent deployability and forgot that deployment independence requires interface stability — and interface stability requires getting the boundaries right. Henney's work explains what went wrong. It also explains what to do instead.

## What a module is

Before microservices, there were modules. Before modules were a deployment unit, they were a design decision. Henney, drawing on Parnas and the pattern literature, defines modularity in terms of *knowledge* — not code, not deployment, not team boundaries.

> "The modularity of a system cannot be evaluated by examining designs of individual modules in isolation. The goal of modular design is to simplify the relationships between components of a system."

This sentence indicts most microservices architectures. Teams design services in isolation — one team per service, optimizing locally — and wonder why the system is a distributed mess. You cannot evaluate modularity by looking at one service. You evaluate it by looking at what happens when one service changes. If the change stays inside the service, the boundary is in the right place. If it propagates, the boundary is decoration. The modularity is not in the service. It is in the relationships *between* services.

> "Coupling is the aspect of a system that defines what knowledge is shared between components of a system. Different ways of coupling components share different types and amounts of knowledge. Some will increase complexity, while others will contribute to modularity."

Coupling is shared knowledge. Not shared libraries. Not shared databases. Not shared deployment pipelines. Shared *knowledge*. When Service A knows the internal data format of Service B, they are coupled — regardless of whether they communicate via HTTP, gRPC, or a message queue. The transport is irrelevant. The knowledge is the coupling. If Service B changes its schema and Service A breaks, the coupling is real. If Service B changes and Service A doesn't notice, the coupling was imaginary. Most organizations don't know which kind they have. They find out in production. The finding is expensive.

## The test of architecture

Henney channels Grady Booch for the definition that should be printed above every microservices whiteboard:

> "All architecture is design but not all design is architecture. Architecture represents the significant design decisions that shape a system, where significant is measured by cost of change."

Cost of change. That is the criterion. If changing a decision requires rewriting one service, it was local design. If it ripples across seven services, it was architectural — and you got the boundary wrong. Most microservices failures are boundary failures. The services work individually. The boundaries between them are in the wrong places. When the business requirement changes, the change crosses five service boundaries and requires coordinated deploys. The architecture didn't reduce the cost of change. It increased it. You distributed a system to make changing your mind *more* expensive. That is not progress.

> "The aggressive pursuit of LCHC (low coupling and high cohesion) can ensure that the effect of change is simplified and isolated, rather than traumatic and global. LCHC also simplifies testing, building, versioning, experimentation, optimization, team organization, and pretty much any other development activity you can think of."

LCHC is the goal. Microservices are one mechanism for achieving it — but only if the boundaries are right. LCHC simplifies everything Henney lists: testing, building, versioning, experimentation, optimization, team organization. If your microservices migration made any of these harder, you didn't achieve LCHC. You achieved distribution without modularity. That is the worst of both worlds: the complexity of a distributed system with the coupling of a monolith. You kept the monolith's problems and added network latency.

> "It is possible to simplify the structure of software without losing effective options. It is even possible to do so and increase your options. Now, that sounds worthwhile: simpler and more flexible."

Simpler. More flexible. After your migration, is the system both? If yes, the boundaries are correct. If no — if it's more complex, harder to change, harder to reason about — the modularity didn't improve. The distribution was cosmetic. The coupling was the real structure all along.

## Boundaries under uncertainty

Henney's deepest architectural insight is about decisions you don't yet know how to make:

> "When a design decision can reasonably go one of two ways, an architect needs to take a step back. Instead of trying to decide between options A and B, the question becomes 'How do I design so that the choice between A and B is less significant?' The most interesting thing is not actually the choice between A and B, but the fact that there is a choice between A and B."

The existence of a choice is a signal. It means the boundary is volatile. Don't pick A or B. Design the interface so that A and B are interchangeable behind it. This is Parnas's information hiding applied to architectural decisions: *hide the decision behind a stable interface, and the choice becomes less significant because you can change your mind later without telling anyone.*

Applied to microservices: don't ask "should this be one service or two?" Ask "how do I design the interface so that whether it's one or two doesn't matter to the callers?" If you can split a service without the callers updating, the interface was right. If splitting forces every caller to redeploy, the interface was coupled to the deployment topology. That coupling is the problem. The service count is a distraction. The interface stability is everything.

> "Although we cannot predict the future with any certainty, it is still possible to write code that is graceful and accommodating — rather than troublesome and resistant — in the face of change."

Graceful in the face of change. Not "scales to a million users." Not "uses the latest framework." Can you change your mind about a decision without rewriting the system? If yes, the architecture works. If no, the architecture is decoration with good slides.

## Simplicity before generality

The principle that should govern every microservices platform decision:

> "The best route to generality is through understanding known, specific examples and focusing on their essence to find an essential common solution. Simplicity through experience rather than generality through guesswork."

Build the specific service first. Make it work. Then, when you have a *second* specific service that resembles the first, find the common essence and extract it. This is the opposite of how most microservices platforms are built. The platform team designs a general-purpose service template with twenty configuration options, five deployment modes, and an abstraction layer that handles every possible use case — most of which never occur. That's not architecture. That's prophecy. Prophecy is usually wrong.

> "A common problem in component frameworks, class libraries, foundation services, and other infrastructure code is that many are designed to be general purpose without reference to concrete applications. This leads to a dizzying array of options and possibilities that are often unused, misused, or just not useful."

The microservices platform designed before any services existed. The shared library that abstracts every database. The common logging framework with seventeen configuration levels. These are not solutions. They are options factories. They produce optionality, not functionality. Most of the options are never used. They exist because someone thought they might be needed. They are inventory. Inventory has cost. The cost is paid by every team that has to understand the platform before they can build the service.

> "Favoring simplicity before generality acts as a tiebreaker between otherwise equally viable design alternatives. When there are two possible solutions, favor the one that is simpler and based on concrete need rather than the more intricate one that boasts of generality."

> "The trick to achieving generality is, somewhat counterintuitively, to make the code specific enough to be fit for purpose."

Specific enough to work. General enough to adapt. Most microservices platforms miss both: too general to be useful for the specific case, too specific to adapt when the case changes. The middle is where Henney operates. It's harder. It's worth it.

## The seven habits, distributed

Henney's most-watched talk — *Seven Ineffective Coding Habits of Many Programmers* — examines patterns acquired by imitation and never questioned. Each habit has a microservices equivalent. Each equivalent is common. Each is wrong.

### 1. Noisy code

> "Comments should provide additional information that is not readily obtainable from the code itself. They should never parrot the code."

> "A common fallacy is to assume authors of incomprehensible code will somehow be able to express themselves lucidly and clearly in comments."

Distributed: Swagger docs that repeat the endpoint name. Architecture decision records that say "we chose Kafka because it scales." Documentation that restates what the API already says. Noise is noise at any scale. A team that can't write a clear function won't write a clear service contract. The scale changes. The skill doesn't. If you can't modularize a monolith, you can't modularize microservices. You'll distribute the spaghetti.

### 2. Visual dishonesty

Code layout must reflect code structure. If indentation lies, the reader is confused. Distributed: the architecture diagram shows clean bounded contexts. The runtime call graph shows a complete mesh. The diagram lies. The system's structure is in the call graph, not the whiteboard. Visual dishonesty at the service level is more dangerous because the evidence is harder to see. You have to instrument the system to discover what the code would have told you directly.

### 3. Lego naming

> "More words is not more meaning."

Distributed: `CustomerProfileServiceManagerFactory`. `OrderProcessingOrchestrationController`. `EnterprisePaymentGatewayAdapterImpl`. Lego naming at the service level catalogues rather than distinguishes. A service name should say what domain function it performs, not what design patterns it contains. The patterns are implementation. The domain function is the contract.

### 4. Under-abstraction

Primitive obsession scaled to services. Services that exchange raw JSON with no schema. Services that pass database IDs instead of domain references. "REST" APIs that are thin wrappers over CRUD. If your service API mirrors your database schema with HTTP verbs, you have not abstracted. You have exposed a tunnel into your storage. The coupling is total. Every caller is now coupled to your schema. Every schema change is a breaking change. The API didn't abstract the storage. It published it.

### 5. Unencapsulated state

Henney's analogy: wearing your underwear on the outside. Distributed: services that share a database. Services that read each other's tables. Services that depend on another service's internal state transitions rather than its published events. If Service A queries Service B's database directly, Service B has no encapsulation. Its state is public. Any schema change in B breaks A. The services are not independent. They are co-dependent with extra network hops and a false sense of modularity.

### 6. Getters and setters

> "When it is not necessary to change, it is necessary not to change." — Lucius Cary

IDE-generated getters and setters become framework-generated CRUD endpoints. `GET /customers/{id}`, `PUT /customers/{id}`, `DELETE /customers/{id}`. Generated. Never questioned. The API exposes mutation where mutation should be a domain event (do you DELETE a customer or *close their account*?). The API exposes internal structure because the framework generated it from the entity. This is not an API. It is an ORM with a public URL and a false sense of RESTfulness.

### 7. Uncohesive tests

> "For tests to drive development they must do more than just test that code performs its required functionality: they must clearly express that required functionality to the reader." — Nat Pryce & Steve Freeman

One test class per class becomes one integration test suite per service. Tests that verify internal implementation rather than external behavior. When you refactor a service's internals, the tests break — not because behavior changed, but because tests were coupled to implementation. The same anti-pattern, distributed. Test behavior. Not structure. At every scale. A test suite that prevents refactoring is not a safety net. It is a cage.

## Code is inventory, not an asset

> "Less code, more software. Less code = less bugs."

> "Programmers write code: a formal plan of the software, expressing its intent in maximal detail. Software is the end product: in execution it is what the user perceives, interacts with and experiences. Sometimes this difference can be significant."

Code is the plan. Software is the product. Confusing them is why organizations measure productivity in lines written and celebrate growing codebases. A growing codebase is a growing liability. Every line must be tested, maintained, understood, migrated, deleted. The asset is the solved problem. The code is what it cost. A microservices migration that produces more total lines of code than the monolith it replaced is not an improvement. It is inventory growth. The problem didn't get bigger. The solution did.

> "Duplicate code has a bad smell and violates the DRY principle... Contrary to benighted management belief, more is not better in this case. Every problem has an optimal code size: too short and the code is cryptic line noise; too long and you cannot see the wood for the trees."

Optimal size exists. Not zero. Not whatever ships. Somewhere in between. Finding it is the discipline. Most teams never look. They add code until the feature works, then stop. The codebase grows monotonically. Nobody's job is deletion. In a microservices world, the cost of excess code is amplified: every service carries its own surplus, and the surplus in each service compounds across the system. Ten services, each 20% larger than needed, is not 20% waste. It is 20% waste multiplied by ten different maintenance, testing, and deployment pipelines. The overhead is multiplicative. The benefit is imaginary.

## The books behind the argument

Henney's microservices-relevant work spans two decades:

**Pattern-Oriented Software Architecture, Volumes 4 & 5** (2007, with Frank Buschmann and Doug Schmidt). Volume 4: *A Pattern Language for Distributed Computing*. Volume 5: *On Patterns and Pattern Languages*. Read Volume 4 before you split another service. The patterns you need — distribution, communication, coordination — are already named and documented. You are discovering them the hard way, expensively, in production. The patterns are free. The downtime is not.

**97 Things Every Programmer Should Know** (2010, editor). Henney's six contributions are micro-essays in modularity:

- **"Comment Only What the Code Cannot Say."** The code says what. The comment says why. At the service level: the API says what. The documentation says why the API exists. Don't document the endpoint. Document the domain decision that created it.
- **"Test for Required Behavior, Not Incidental Behavior."** Test the contract, not the implementation. At the service level: test the API contract, not the internal implementation. If an internal refactor breaks the test suite, the tests were coupled. Decouple them.
- **"Test Precisely and Concretely."** Vague tests are worse than no tests — they give false confidence. A service integration test that passes with a mock is testing the mock, not the service.
- **"Name the Date."** Don't write `getDate()`. Write `getExpiryDate()`. Don't name a service `data-processor`. Name it `invoice-generator`. The extra word is not noise. It is the contract.
- **"Program with GUTs."** Good Unit Tests. Test code is code. It deserves review, refactoring, the same standards. Service tests that nobody reads are not tests. They are rituals.
- **"Uncheck Your Exceptions."** Checked exceptions create coupling. At the service level: checked exceptions are shared error types that force every caller to depend on every callee's internal error taxonomy. That coupling crosses service boundaries. It shouldn't.

**97 Things Every Java Programmer Should Know** (2020, co-editor with Trisha Gee). Same format, Java lens. The editorial voice is unchanged: practical, specific, allergic to dogma.

## What Henney knew

The microservices movement made modularity a deployment concern. Henney's work — and the pattern literature he co-authored — makes modularity a *knowledge* concern. The deployment follows the knowledge boundary. Not the other way around.

If you split by knowledge — by what each module knows and hides — the deployment topology emerges naturally. Each knowledge boundary becomes a service boundary. Each hidden decision becomes an internal implementation. Each stable interface becomes an API contract. The services are modular because the knowledge is modular. The deployment is a consequence.

If you split by database table, by team size, or by "one service per aggregate," the knowledge boundaries don't match the service boundaries. Knowledge leaks across services. Coupling becomes the runtime reality, regardless of what the diagram shows. The services are not modular. They are fragments of a monolith, communicating over a network, with all the original coupling intact and new failure modes added.

Henney never wrote *Building Microservices*. He wrote the theory that explains why most of them fall over. Coupling is shared knowledge. Architecture is measured by cost of change. Modularity lives in the relationships, not the services. Simpler and more flexible is the only test that matters. Most microservices migrations fail that test. Henney's work explains why. It also explains how to pass it. The books are short. The ignorance is expensive.

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


Infrastructure choice is engineering choice. The protocol, the format, the runtime — each is a decision that shapes everything built on top of it. DOT for pipelines, Temporal vs DBOS for durable execution, NATS for messaging. The choice determines the coupling, the scaling, the failure modes. The engineer who treats infrastructure as a commodity gets the failure modes of the default. The engineer who treats it as a design decision gets the failure modes they chose.
