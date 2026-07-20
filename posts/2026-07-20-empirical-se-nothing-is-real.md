---
title: "Empirical Software Engineering: Nothing Is Real"
date: 2026-07-20
slug: empirical-se-nothing-is-real
summary: "Hillel Wayne's talk on empirical software engineering opens with a confession: we know almost nothing with scientific certainty about how to build software. The TDD wars, the language debates, the methodology fights — none of them have the evidence they claim. That's not depressing. It's liberating."
tags: empirical-se, evidence, methodology, software-engineering, hillel-wayne
series: empirical-software-engineering
series_order: 1
---

In 2014, David Heinemeier Hansson declared "TDD is dead." Robert C. Martin, a decade earlier, had called TDD a professional moral imperative — "you are unprofessional if you do not practice TDD." Two of the most influential voices in software. Two positions that cannot both be true. Zero empirical studies cited by either of them.

This is not a post about TDD. It's a post about the fact that the most important methodological debate in modern software engineering was fought entirely on the basis of authority, anecdote, and vibes.

Hillel Wayne's talk *What We Know We Don't Know* is an introduction to empirical software engineering — the study of what actually works in programming, using data, controlled studies, and peer review instead of intuition. His opening thesis is uncomfortable: **nobody has the secret knowledge.** Not Uncle Bob. Not DHH. Not the tech lead who swears by microservices. Not the staff engineer who says monoliths are the only sane choice. Nobody actually knows.

> "Nothing is real, we don't understand what we're doing, and the only way to write good software is to stop drinking coffee. Burn it all down." — Hillel Wayne, the actual description of his talk

## The problem with intuition

Software engineering runs on intuition. Someone tries a practice. It feels better. They tell their team. The team adopts it. A blog post is written. A conference talk is given. A book is published. Within five years, the practice is orthodoxy, and anyone who questions it is unprofessional.

This is how we got TDD as a moral imperative. This is how we got "microservices are the default architecture." This is how we got "dynamic typing is for prototypes, static typing is for production." None of these statements are supported by evidence. All of them are supported by the confidence of the people saying them.

> The strongest opinions in software engineering are held by the people who have done the least systematic investigation. The people who have actually run the studies sound nothing like them.

Wayne points to the TDD debate as the archetypal case. Two camps, each utterly certain, neither citing evidence. The empirical literature on TDD is, in Wayne's words, "iffy." He personally likes TDD. But the data doesn't strongly support it. And the data doesn't strongly refute it either. The honest answer is that we don't know — and that answer is unacceptable to both sides of the debate.

## The COST of scale

The talk opens with a paper that should humble anyone who has ever reached for a distributed system because "we need to scale."

McSherry et al., in *Scalability! But at what COST?* (USENIX HotOS '15), showed that single-laptop implementations often outperform large distributed systems when measured by total computational cost. The distributed system is faster in wall-clock time, sure. But the laptop required zero nodes, zero network, zero orchestration. The distributed system's advantage vanishes when you measure total work done, not time elapsed.

> The paper is a controlled demolition of the reflex to scale before you need to. Most systems don't need to be distributed. Most systems that are distributed don't need to be. The decision to distribute is almost never made empirically. It's made because distributed systems are cool, and cool is not a performance metric.

## Why we don't know

Wayne's talk isn't nihilistic. It's diagnostic. The reason we know so little is not that software is uniquely unstudyable. It's that we haven't built the empirical culture.

Construction engineering has CURT — a consortium that studies the effect of overtime on project outcomes, producing reports that show (with data) that extended overtime yields *negative* productivity. Software engineering has blog posts.

> Every other profession that deals with complex systems under uncertainty has developed empirical traditions. Nursing studies what works. Teaching studies what works. Law studies what works. Software engineering studies what gets GitHub stars.

The talk is a call to build that culture. To read papers. To run studies. To replicate findings. To be suspicious of anyone — including yourself — who is certain about what works without evidence.

## The liberating truth

Wayne's deepest point is that the lack of empirical certainty is not depressing. It's liberating.

When someone tells you that you are unprofessional for not using TDD, you now know that they don't actually know. When someone tells you that static types prevent bugs, you now know the evidence is weak and conflicting. When someone tells you that pair programming is the answer, you now know the studies show some benefit but smaller and less reliable than code review.

> Nobody has the secret knowledge. The loudest voices in the room are not the most informed. They're the most confident — and confidence is negatively correlated with accuracy once you control for expertise.

The honest answer to most software engineering questions is "it depends" followed by layers of caveats. That answer is unsatisfying. It's also true. And building a discipline that can tolerate unsatisfying truths is harder than building one that rallies around confident falsehoods. But it's the only kind of discipline worth building.

---

*Based on Hillel Wayne's talk [What We Know We Don't Know: Empirical Software Engineering](https://www.hillelwayne.com/talks/ese/). Part 2: [What the Studies Actually Say](#empirical-se-what-studies-say) · Part 3: [How to Think Empirically](#empirical-se-how-to-think).*
