---
title: Parnas's Information Hiding
date: 2026-07-11
slug: parnas-information-hiding
summary: "David Parnas's 1972 paper introduced information hiding: modularize around design decisions likely to change, not around processing steps. The interface reveals as little as possible."
tags: design, parnas, information-hiding, modules, software-architecture
---

In 1972, David Parnas published a paper that changed how programmers think about modules. The title is dry — "On the Criteria to Be Used in Decomposing Systems into Modules." The content is a controlled explosion.

> "The effectiveness of a 'modularization' is dependent upon the criteria used in dividing the system into modules."

Before Parnas, modularization meant dividing the program by what it did, in the order it did it. Step one, module one. Step two, module two. The flowchart was the architecture. Parnas pointed out that this criterion — "major steps in the processing" — is the least stable thing about any system.

> "Note, however, that nothing is said about the criteria to use in dividing the system into modules. Because the decision to divide a system into n modules of a given size does not determine the decomposition, this paper will discuss that issue."

Having n modules is not enough. The *criterion* for drawing the boundaries is everything. The same system with the same number of modules can be flexible or brittle depending on where you put the cuts. Most programmers were putting the cuts in the wrong place. Most still are.

## The KWIC index

Parnas illustrates with a KWIC (Key Word In Context) index — a system that takes lines of text, produces all circular shifts, sorts them, and formats output. Small enough to understand in one sitting. Complex enough to teach a fifty-year lesson.

Two modularizations, same functionality:

**Modularization 1: conventional.** Five modules by processing step. Input. Circular Shift. Alphabetizer. Output. Master Control. The data flows predictably. The interfaces expose data structures: character packing, pointer conventions, core formats. Every module knows how the data is stored. This is how most programmers would build it.

> "In the first decomposition the criterion used was make each 'major step' in the processing a module."

**Modularization 2: information hiding.** Also five modules. Line Storage hides character packing behind access functions. Input hides the input format. Circular Shifter hides the shift algorithm behind access functions identical to Line Storage. Alphabetizer hides the sort. Output hides the format. Each module's interface reveals nothing about how it works inside.

> "The second decomposition was made using 'information hiding' as a criteria. The modules no longer correspond to steps in the processing."

> "Every module in the second decomposition is characterized by its knowledge of a design decision which it hides from all others. Its interface or definition was chosen to reveal as little as possible about its inner workings."

Interface as minimum revelation. Implementation as concealed volatility. A module is not a step in a process. It is a secret keeper. It knows something the others don't, and it tells them only what they need.

## The change propagation test

Parnas applies the test that matters: what modules change when a single design decision changes?

Input format changes. Modularization 1: every module touches the data. Input, Circular Shift, Alphabetizer, Output, Master Control — all break. Modularization 2: one module. The Input module's interface is stable. Nobody else knows the format changed.

Sort algorithm changes. Modularization 1: Alphabetizer and Output break. Modularization 2: one module. The Alphabetizer's interface holds. The rest of the system hasn't heard of the sort algorithm it was using and doesn't need to.

Storage representation changes — packed characters four to a word, or linked lists, or a tree. Modularization 1: every module breaks. Modularization 2: one module. Line Storage. Nothing else even knows how characters are stored.

> "It is by looking at changes such as these that we can see the differences between the two modularizations."

The criterion is not function. It is volatility. Ask not what the module does. Ask what it protects you from when the world changes.

## Why this was radical

Parnas inverted the logic of decomposition. You don't modularize by what the system does. You modularize by what you want to be able to change without telling anyone.

> "We propose instead that one begins with a list of difficult design decisions or design decisions which are likely to change. Each module is then designed to hide such a decision from the others. Since, in most cases, design decisions transcend time of execution, modules will not correspond to steps in the processing."

This is the core of the paper. Begin with the volatile. Wrap each volatile decision in a module whose interface survives the change. The execution order is secondary. The change isolation is primary.

Dijkstra, Parnas's contemporary, had already argued for separation of concerns as a structuring principle. But Dijkstra was thinking about intellectual manageability — can you reason about one part without holding the whole in your head? Parnas added the operational criterion: can you change one part without touching the others? The two criteria converge. What you can change without touching is what you can reason about in isolation. What you can reason about in isolation is what you can build independently.

> "The major progress in the area of modular programming has been the development of coding techniques and assemblers which allow one module to be written with little knowledge of the code used in another module... but its use has not resulted in the expected benefits."

The tools existed. The assemblers could do separate compilation. Modules could be swapped without full rebuilds. The mechanics worked. The designs didn't. Because the criterion was wrong.

## The efficiency objection

Parnas was honest about the cost. Information-hiding modularizations, implemented conventionally as subroutines, would be less efficient than the conventional decomposition. More function calls. More indirection. More boundaries to cross.

> "The unconventional decomposition, if implemented with the conventional assumption that a module consists of one or more subroutines, will be less efficient in most cases."

He sketched an alternative: a preprocessor that inlines module accesses at compile time. The modularization stays clean. The runtime code stays fast. This was 1971. He was describing what we now call zero-cost abstractions. C++ would take twenty more years to get there. Rust would take forty.

Hoare, in his work on data abstraction, was developing compatible ideas from a different direction — proving modules correct through their specifications. Parnas was less interested in proof than in adaptability. Hoare wanted to know the module was right. Parnas wanted to know you could change your mind about what was right without rewriting everything. Both were necessary. Neither was sufficient alone.

## What "information hiding" actually means

The term has been so thoroughly absorbed that it now means almost nothing. Most "encapsulation" in production code is theater. A getter that returns the internal data structure is not hiding anything. It is exposing the decision with extra syntax.

> "Its interface or definition was chosen to reveal as little as possible about its inner workings."

*As little as possible.* Not "as little as convenient." Not "as little as the framework supports." As little as the problem allows while still being useful. If you change the data structure and the caller's code still compiles, you did information hiding. If the caller imports the type, you didn't. If the caller casts to the internal type, you really didn't. If the caller wrote their own parser for your output format, you have lost control of your own design decision and someone else is now coupled to a choice you didn't even know you were making.

Wirth, with stepwise refinement, had given programmers a method for designing top-down. Parnas gave them a reason to design bottom-up — from the volatile decisions outward. The two methods are compatible but the emphasis is different. Wirth asked: what are the steps? Parnas asked: what are the secrets? The steps change. The secrets change too, but when they do, you want them contained.

## The connection to Brooks

This is the paper referenced throughout the Brooks on Software Design series. Parnas and Brooks worked the same era and reached compatible conclusions from different premises.

Brooks: "The building of a design is the forcing of the will of one upon the stuff of the world." One mind controls the interfaces. Parnas: those interfaces exist to hide the volatile decisions. One mind decides *what to hide*. The rest of the system builds against stable interfaces and doesn't need to know.

Brooks without Parnas: one mind, but no criterion for where to draw the modular boundaries. Parnas without Brooks: the right criterion, but no mechanism for enforcing it across a system. Together: one designer, one set of secrets, one set of stable interfaces. That is the architecture of every system that has aged well. That is also the architecture of almost no system you actually work on.

The paper is from 1971. It is fifty-five years old. It was right then. It is right now. Most production code still violates its central principle — modularizing by step, not by secret, and wondering why changes ripple. This is either comforting or damning. Parnas would probably note that both reactions are consistent with his theory. He'd hidden which one he meant. That's how you know the man understood his own idea.

---

**Reference:** David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, Vol. 15, No. 12, December 1972, pp. 1053-1058. [PDF](https://prl.khoury.northeastern.edu/img/p-tr-1971.pdf)
