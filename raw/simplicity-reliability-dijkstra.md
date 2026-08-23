---
title: Simplicity Is the Prerequisite for Reliability
date: 2026-07-22
slug: simplicity-reliability-dijkstra
summary: Dijkstra's A Discipline of Programming argued that correctness proofs must be developed alongside programs — and that simplicity is not an aesthetic preference but a practical necessity for making those proofs possible.
tags: simplicity, reliability, dijkstra, correctness, formal-methods, software-engineering
---

In 1976, Edsger Dijkstra published a book that opened with a claim most programmers would find alien today: that the primary task of a programmer is not to write programs, but to construct proofs. The programs are a byproduct. The proofs are the work.

*A Discipline of Programming* is 217 pages. It introduces the weakest precondition calculus, the guarded command language, and a methodology for deriving programs from their specifications by developing the correctness proof slightly ahead of the code. The book is difficult. It is also the most sustained argument ever written for the proposition that simplicity and reliability are the same property considered from different angles.

## The argument

Dijkstra's central claim is that you cannot verify correctness after the fact. You must construct it alongside the program, with the proof leading and the code following. His own formulation of the conclusion is worth quoting at length, because the precision of the language is the argument:

> "It does not suffice to design a mechanism of which we hope that it will meet its requirements, but that we must design it in such a form that we can convince ourselves — and anyone else for that matter — that it will, indeed, meet its requirements. And, therefore, instead of first designing the program and then trying to prove its correctness, we develop correctness proof and program hand in hand. (In actual fact, the correctness proof is developed slightly ahead of the program: after having chosen the form of the correctness proof we make the program so that it satisfies the proof's requirements.)"

This is not a claim about process. It is a claim about the relationship between a program and the reasoning that justifies its existence. If you write the program first and attempt to verify it afterward, you are attempting to reconstruct the reasoning that would have produced the program had the program been derived from its specification. The reconstruction is harder than the derivation would have been, because the program contains implementation decisions that were made without being constrained by the proof. You are now trying to discover whether those unconstrained decisions happen to be correct. The probability that they all are, in a program of non-trivial size, is remote.

> The practical consequence: programs derived from their proofs are reliably shorter and clearer than programs written forward from intuition, because the proof forces you to eliminate everything that is not necessary to establish the postcondition.

The mechanism Dijkstra proposed for this derivation is the **weakest precondition calculus**. For any program statement and desired postcondition — a logical formula describing what must be true after the statement executes — the weakest precondition is the least constrained precondition that guarantees the statement will terminate in a state satisfying the postcondition. The calculus provides transformation rules for each construct in the language. To derive a program, you start from the postcondition and work backward through the rules until you reach a precondition you can satisfy. The program that emerges from this process is correct by construction. It cannot be otherwise, because it was built to satisfy the proof at each step.

The language Dijkstra designed for expressing these derivations is deliberately minimal — alternation and repetition with guarded commands, no recursion, no complex features. His justification is instructive: "The point is that I felt no need for them in order to get my message across, viz. how a carefully chosen separation of concerns is essential for the design of in all respects, high-quality programs: the modest tools of the mini-language gave us already more than enough latitude for nontrivial, yet satisfactory designs." The language is minimal because minimality is the point. Every construct you add to a language expands the space of programs that can be written and therefore the space of programs that must be verified. A smaller language makes the verification task smaller.

## Why simplicity is not an aesthetic preference

Dijkstra identified several ideas in the book that he described as "elusive" — ideas that should take root in the mind of a programmer but typically don't, because the industry treats them as matters of style rather than as structural requirements.

The first of these is simplicity. Dijkstra's argument is not that simple programs are nicer to read. It is that simple programs are the only programs for which correctness proofs are feasible. Complexity is not merely unpleasant; it is a barrier to verification. A program that is too complex to reason about is a program whose correctness cannot be established. Whether it actually works is unknown, and testing — Dijkstra's most famous observation — "reveals only the presence of errors, not their absence." A tested program that passes all its tests is a program for which no known inputs produce incorrect outputs. It is not a correct program. The distinction is not philosophical. It is the difference between a bridge that has survived every load it has encountered and a bridge that has been shown, by structural analysis, to withstand every load it could encounter.

The second is elegance. Dijkstra used the word, which makes engineers uncomfortable, but he meant something precise. An elegant solution is one in which the proof of correctness is natural — where the formal argument flows without contortion, because the program structure mirrors the logical structure of the specification. Elegance is not decoration. It is evidence that the derivation worked.

> In Dijkstra's framework, simplicity, elegance, and reliability are not three properties. They are one property described three ways. A simple program is one whose correctness proof is manageable. An elegant program is one whose correctness proof is natural. A reliable program is one whose correctness proof exists.

## The connection to the agent era

Agents produce programs by statistical prediction. They do not construct proofs. They do not derive programs from specifications. They generate tokens that are probable given their training data, and the programs that result from this process have no accompanying reasoning that justifies why they are correct. They might be correct. They might not be. There is no way to know from the output alone, because the output was not produced by a process designed to establish correctness.

This makes Dijkstra's argument more urgent, not less. If programs are increasingly produced by entities that cannot reason about their correctness, then the verification burden shifts entirely to the infrastructure surrounding those entities. The harness must supply the reasoning the agent cannot perform. The verification layer must establish what the generation process cannot guarantee.

Dijkstra argued that the proof must lead and the program must follow. In the agent era, the proof still has to exist. The question is who — or what — constructs it, and whether the program is constrained by it or merely inspected by it afterward. If the agent produces code and a separate system verifies it, the verification is still an attempt to reconstruct the reasoning that would have produced the program had it been derived correctly. The reconstruction is harder than the derivation. The probability that it succeeds on every change, at the speed agents generate changes, is remote.

> Dijkstra's methodology does not scale to the agent era in its original form. No one is going to derive weakest preconditions for agent-generated code at review time. But the principle — that correctness must be constructed, not inspected into existence — does not become false because it becomes harder to satisfy. It becomes more expensive to ignore.

![A Discipline of Programming — Edsger Dijkstra, 1976. 217 pages. The most sustained argument ever written for the proposition that simplicity and reliability are the same property.](/images/discipline-of-programming.jpg)

---

**References:**

- Dijkstra, E. W. (1976). *A Discipline of Programming.* Prentice-Hall. — The source text: weakest preconditions, guarded commands, and the argument that correctness proofs must lead and programs must follow.
- Dijkstra, E. W. (1972). "The Humble Programmer." *Communications of the ACM.* — The earlier lecture that established the tone: programming is inherently difficult, and the only way to manage that difficulty is through disciplined simplicity.
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first) — What OpenBSD teaches about correctness as the prerequisite for security.
- Related: [Taste as Conceptual Integrity](https://blog.hackspree.com/#taste-conceptual-integrity) — Brooks on the property Dijkstra's methodology was designed to preserve.
- Related: [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) — The verification imperative in the agent era.
