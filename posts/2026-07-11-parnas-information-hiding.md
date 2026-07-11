---
title: "information hiding"
date: 2026-07-11
slug: parnas-information-hiding
summary: "David Parnas's 1972 paper changed how we think about modules. The criterion for decomposition is not what the system does. It is what you want to hide when things change."
tags: design, parnas, information-hiding, modules, software-architecture
---

In 1972, David Parnas published a paper with a dry title and a wet explosive inside. "On the Criteria to Be Used in Decomposing Systems into Modules" sounds like something you'd skim at a faculty meeting. It is not. It is the paper that taught software engineers why some decompositions age well and others rot.

The argument: **modularization is not about dividing up the work. It is about hiding design decisions that will change.**

Every module should hide one design decision from the others. The interface is what stays stable. The implementation is what you're allowed to change without telling anyone. This sounds obvious now. In 1972, nobody was doing it. Most still aren't.

## Two ways to slice the same system

Parnas demonstrates with the KWIC (Key Word In Context) index — a system that takes lines of text, produces all circular shifts of each line, sorts them alphabetically, and formats the output. Simple enough to implement in an afternoon. Complex enough to teach a lesson that has lasted fifty years.

He compares two modularizations:

**Conventional decomposition.** Five modules, each doing one step of the process. Input module reads. Circular shift module rotates. Alphabetizer sorts. Output formats. A master control sequences them. The steps are clear. The data flows predictably. This is how most programmers would build it. This is how most programmers still build it.

**Information-hiding decomposition.** Five modules, each hiding one design decision. Input format hidden behind an interface that delivers lines. Circular shift algorithm hidden behind an interface that produces shifts. Sorting algorithm hidden. Output format hidden. The master control is gone, replaced by module cooperation through stable interfaces.

The conventional decomposition is easier to explain. The information-hiding decomposition is easier to change.

## What happens when things change

Parnas applies a test: what modules change when a single design decision changes?

Change the input format. Conventional: input module, circular shift module, alphabetizer, output, master control. Everything touches the data format. Everything breaks. Information-hiding: one module. The input module's interface stays the same. Nobody else knows the format changed. Nobody else cares.

Change the sorting algorithm. Conventional: alphabetizer, output — at minimum. The sort order propagates to formatting. Information-hiding: one module. The alphabetizer's interface stays the same. The rest of the system hasn't even heard of the sort algorithm it was using. It doesn't need to.

Change the output format. Conventional: master control, output, probably alphabetizer if the sort depends on how lines are represented. Information-hiding: one module. The output module.

Add a new feature — say, filtering out stop words before sorting. Conventional: changes ripple across the pipeline. New module wedged in. Data format extended. Everyone downstream needs to know. Information-hiding: one new module, implementing the same interface as the old one, delegating what it doesn't change. The rest of the system sees the same interface. It doesn't know the module behind it changed. That's the point.

> "We propose instead that one begins with a list of difficult design decisions or design decisions which are likely to change. Each module is then designed to hide such a decision from the others."

The criterion is not function. The criterion is volatility. Ask not what the module does. Ask what it protects you from when the world changes.

## Why this was radical

In 1972, modularization meant dividing the program by what it did, in the order it did it. Step one, module one. Step two, module two. The flowchart was the architecture. Parnas pointed out that the flowchart is the least stable thing about any system. The steps change. The order changes. What stays stable are the fundamental design decisions — the data representation, the algorithm choice, the hardware interface, the output format — and those should be the basis of modularization, precisely because they will change.

This inverted the logic of decomposition. You don't modularize by what the system does. You modularize by what you want to be able to change without telling anyone. A module is not a step in a process. It is a secret keeper. It knows something the other modules don't, and it exposes only what they need.

> "Every module is characterized by its knowledge of a design decision which it hides from all others. Its interface or definition was chosen to reveal as little as possible about its inner workings."

Interface as revelation of the minimum. Implementation as concealment of the volatile. This is the definition of a good API. It is also the definition of a good organizational boundary. It is also the definition of a healthy relationship. Parnas accidentally described how to structure everything.

## The consequences

The paper's impact is hard to overstate, partly because it succeeded so completely that its ideas became invisible. Information hiding is now just "how you design modules." The concept was absorbed into object-oriented programming — objects as units that hide data behind methods — and into microservices — services as units that hide databases behind APIs — and into every architecture document that talks about "separation of concerns."

But most implementations get the form without the function. They hide data behind getters and setters and call it encapsulation. Parnas would not be impressed. A getter that returns the internal data structure is not hiding anything. It is exposing the decision with extra syntax. True information hiding means the caller does not know the representation. They cannot know. The interface is designed to make it impossible to tell. If you change the data structure and the caller's code still compiles, you did information hiding. Otherwise, you did theater.

> "The criteria for module decomposition should be based on minimizing the propagation of change."

Minimizing. Not eliminating. Some changes propagate — you cannot hide everything. But every change that crosses a module boundary is a failure of the modularization. The ideal module boundary prevents the change from crossing. The realistic one minimizes how far it goes.

## The connection to Brooks

This is the paper I kept referencing in the Brooks on Software Design series. Parnas and Brooks worked in the same era, on similar problems, and reached compatible conclusions from different angles. Brooks argued that conceptual integrity requires one mind — one designer controlling the interfaces. Parnas showed what those interfaces should hide: the volatile decisions.

Brooks without Parnas: one mind, but no criterion for what to modularize. Parnas without Brooks: the right modularization, but no mechanism for achieving it across a system. Together: one mind deciding what to hide, and everyone else building behind stable interfaces. That is the architecture of every system that has aged well.

The paper is from 1972. It is fifty-four years old. It is still correct. Most production systems still violate its basic principle. This is either comforting or depressing. Parnas would probably say both. He was like that.

---

**Reference:** David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, Vol. 15, No. 12, December 1972, pp. 1053-1058. [PDF](https://prl.khoury.northeastern.edu/img/p-tr-1971.pdf)
