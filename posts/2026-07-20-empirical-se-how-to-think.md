---
title: "Empirical Software Engineering: How to Think Empirically"
date: 2026-07-20
slug: empirical-se-how-to-think
summary: "How to read a study, spot a confound, and build a personal practice of empiricism. The methods, the reading list, and the intellectual posture that Wayne's talk recommends."
tags: empirical-se, methodology, critical-thinking, studies, hillel-wayne
series: empirical-software-engineering
series_order: 3
---

Part 1 argued that we know almost nothing. Part 2 covered what we do know. This post is about how to think — the methods, the habits, and the reading list for building a personal practice of empirical software engineering.

## The three methods

Wayne organizes empirical SE into three research methodologies. Each answers a different kind of question. Each has different strengths and failure modes.

**Controlled trials** manipulate one variable and measure the outcome. Häregård & Kruger's study on syntax highlighting: give two groups the same code with different color schemes, measure comprehension. Strongest for establishing causation. Weakest for ecological validity — lab conditions are not production conditions. A developer reading code in a study is not a developer fixing a production outage at 3am.

**Natural experiments** exploit real-world variation that approximates random assignment. Yuan et al.'s study of catastrophic failures: the researchers didn't cause the failures. They analyzed them after the fact, asking what would have prevented each one. High ecological validity — these were real failures with real consequences. Lower control — you can't randomize which services fail in production.

**Observational studies** measure what people actually do and look for correlations. Meyer et al.'s study of developer productivity perceptions: ask developers what makes them feel productive, correlate with what they actually do. Reveals practices at scale. Cannot establish causation. Developers who use TDD might produce better code, or developers who produce better code might be the kind of developers who use TDD. The study can't tell you which.

> The method determines the claim. A controlled trial can say "X causes Y." An observational study can only say "X and Y co-occur." Most industry blog posts can only say "I tried X and it felt good." Know which kind of claim you're reading.

## How to read a study

Wayne's talk implies a method for reading empirical research that most software engineers never learned:

1. **What's the methodology?** Controlled trial, natural experiment, or observational? This determines what the study can claim.

2. **What are the confounds?** What else varies with the variable being studied? The Ray et al. language study: functional languages correlated with smaller projects and more experienced developers. Those are confounds. They explain the effect better than the language does.

3. **Has it been replicated?** A single study is a data point, not a conclusion. The replication is the signal. No replication, no confidence.

4. **What's the effect size?** Statistical significance is not practical significance. A language might "significantly" reduce defects by 0.3%. That's not worth switching languages for.

5. **Who funded it?** This isn't cynicism. It's basic epistemic hygiene. A study funded by a vendor is not necessarily wrong. But it deserves more scrutiny than an independent replication.

> Most software engineers read zero studies. Most of the ones who read one study stop at step one. The skill is reading enough studies, across methodologies, with attention to confounds and replications, to form a tentative and revisable picture of what might be true.

## The self-correction of science

The Ray et al. → Berger et al. replication saga is the centerpiece of Wayne's talk, and he uses it to make a point about how knowledge actually advances.

The original study found an effect. A replication found the effect was confounded. The original authors engaged with the replication. Knowledge advanced. This is science working as designed. In software engineering, the original study became "functional languages prevent bugs" and entered the permanent folklore. The replication was a paper that most people never read.

> The replication is the immune system. Without it, the field accumulates unquestioned findings until the body of "knowledge" is mostly noise. Wayne's 6,000-word writeup of the saga is titled "This is How Science Happens." The subtitle is implicit: this is how science happens *when it's allowed to.*

## The reading list

Wayne recommends four books and two ongoing sources:

**Books:**
- *Making Software* (Oram & Wilson) — a curated collection of empirical SE papers with practitioner-oriented summaries. The best entry point.
- *Leprechauns of Software Engineering* (Bossavit) — how folklore becomes fact in software, and how to spot it.
- *The Programmer's Brain* (Hermans) — what cognitive science tells us about reading, writing, and debugging code.
- *Teaching Tech Together* (Greg Wilson) — how people learn programming, based on educational research.

**Ongoing:**
- [*It Will Never Work In Theory*](https://neverworkintheory.org/) — a blog that summarizes empirical SE papers for practitioners. Short, readable, rigorous.
- ACM Digital Library and arXiv — the primary sources, if you want to read the papers themselves.

> The reading list is not long. It's not meant to be. The barrier to being more empirical than 99% of software engineers is reading four books and following one blog. The bar is on the floor.

## The posture

Wayne's deepest lesson is not about any specific study. It's about intellectual posture.

The empirical posture is: *I don't know, but I can find out, and until I do I will hold my beliefs lightly.* It's the opposite of the posture that dominates software engineering — the tech lead who knows The Right Way, the consultant who sells The Methodology, the influencer who declares that The Old Way Is Dead.

> The empirical posture is less charismatic. It doesn't keynote. It doesn't trend. It says "the evidence is mixed, the effect sizes are small, and it depends." That sentence will never go viral. It is also the most honest sentence available in most software engineering debates.

Being empirical doesn't mean reading every paper. It means knowing that papers exist. It means being suspicious of certainty. It means asking "how do we know that?" when someone declares a practice mandatory or obsolete. It means holding your own preferences lightly — liking TDD because it feels good while acknowledging the evidence is weak. It means being the person in the room who says "I'm not sure" and means it as a strength, not a weakness.

> The goal is not to know everything. The goal is to know the difference between what you know and what you believe.

---

*Based on Hillel Wayne's talk [What We Know We Don't Know](https://www.hillelwayne.com/talks/ese/). Part 1: [Nothing Is Real](#empirical-se-nothing-is-real) · Part 2: [What the Studies Actually Say](#empirical-se-what-studies-say).*
