---
title: "Empirical Software Engineering: What the Studies Actually Say"
date: 2026-07-20
slug: empirical-se-what-studies-say
summary: "Code review works. Sleep deprivation destroys performance. The language you choose probably doesn't matter. And a famous study claiming Haskell prevents bugs turned out to be wrong. What happens when you actually read the papers."
tags: empirical-se, evidence, code-review, human-factors, replication
series: empirical-software-engineering
series_order: 2
---

The first post made the case that we know almost nothing. This post is about what we actually do know — the handful of findings that have survived scrutiny, replication, and the self-correction mechanisms of science.

## What works: code review

The strongest, most replicated finding in empirical software engineering is that **code review finds bugs.** It consistently catches a large portion of defects in reasonable time. This is not surprising — having a second person look at the code works for the same reason having a second person look at anything works. But the studies add texture.

**File position matters.** A 2022 study (*First Come First Served: The Impact of File Position on Code Review*) found that reviewers give more scrutiny to files appearing earlier in a changeset. Later files get less attention. If you want a file reviewed thoroughly, put it first.

**Not all review is equal.** Google's internal study of their own review practices (*Modern Code Review: A Case Study at Google*) found that review effectiveness varies enormously with reviewer expertise, changeset size, and tooling. The practice is effective, but the variance is large. "Do code review" is like "eat healthy" — the general advice is correct, the specific execution determines most of the outcome.

**Most review catches style, not logic.** The SmartBear study (*Best Kept Secrets of Peer Code Review*) found a roughly 3:1 ratio of style issues to actual bugs. Review catches formatting problems and naming issues far more reliably than it catches subtle control-flow errors. This is not an argument against review. It's an argument for understanding what review actually provides.

> Code review is the closest thing empirical SE has to a settled finding: it works, consistently, across contexts. Everything else is weaker.

## What works: basic testing

Yuan et al.'s *Simple Testing Can Prevent Most Critical Failures* (USENIX OSDI '14) is the kind of study that makes you reconsider everything you think about sophistication.

The authors examined real-world catastrophic production failures — the kind that take down services and make the front page of Hacker News — and asked: what testing would have caught this? The answer, in the majority of cases, was **trivial testing.** Not property-based testing. Not formal verification. Not chaos engineering. Just basic unit tests and integration tests on the code paths that failed.

> The finding is humbling: most catastrophic failures would have been prevented by practices we already know how to do and mostly don't do consistently. The problem is not lack of sophisticated techniques. The problem is lack of consistent application of simple ones.

## What works: not being exhausted

The human-factors literature is unambiguous. Sleep deprivation degrades cognitive performance measurably and dramatically. Extended overtime in construction produces *negative* total productivity after a few weeks (CURT, NIOSH). In game development, mandated crunch correlates with *worse* project outcomes (Tozour, The Game Outcomes Project). Fucci et al. demonstrated that a single night of sleep deprivation measurably degrades novice programmers' performance on coding tasks.

> The effect sizes on human factors dwarf the effect sizes on any technical choice. You will get more improvement from letting your team sleep than from switching languages, frameworks, or methodologies. The evidence for this is overwhelming. The adoption rate is abysmal.

## What doesn't work: the language doesn't care about you

In 2014, Ray et al. published *A Large Scale Study of Programming Languages and Code Quality in GitHub.* It appeared to show that functional languages — Haskell, Clojure, Scala — produced fewer defects than procedural and scripting languages. The paper made the rounds. It was cited as proof that type systems prevent bugs, that functional programming is safer, that language choice is a significant lever for quality.

Then Berger et al. replicated the study. They controlled for project size, developer experience, and other confounds the original had missed. **The effect largely vanished.** Languages that appeared safer were mostly being used by more experienced developers on smaller projects. When you controlled for that, the language effect became negligible.

> This is how science is supposed to work: study finds effect, replication corrects it, knowledge advances. In software engineering, the original study became folklore and the replication was ignored. The lesson is not "language studies are hard." The lesson is "don't believe a single study, ever."

## What we don't know: TDD, types, pair programming

The evidence for TDD is mixed and weak. Wayne personally likes it but acknowledges the data doesn't strongly support it. The evidence for static typing preventing bugs is, in his words, that one study "found no clear evidence it helps — or hurts." Pair programming shows some benefit in some studies, but the effect is smaller and less consistent than code review.

> None of this means these practices are worthless. It means the evidence doesn't support being dogmatic about them. Anyone who tells you a practice is mandatory or immoral is making a claim they cannot back with data. They are selling you vibes dressed as expertise.

## The meta-lesson

The most important finding in empirical software engineering is not about any specific practice. It's about the structure of knowledge itself.

Most things we believe about software development come from authority, anecdote, and marketing. A small fraction come from studies. Of those studies, a meaningful portion fail to replicate. The findings that survive are humbler than we want them to be: fundamentals, consistently applied, under conditions of adequate sleep and manageable stress, produce better outcomes than heroics.

> The evidence doesn't tell you which framework to use. It tells you to review the code, test the critical paths, and go to bed at a reasonable hour. Everything else is speculation. Some of it is good speculation. None of it is science.

---

*Based on Hillel Wayne's talk [What We Know We Don't Know](https://www.hillelwayne.com/talks/ese/). Part 1: [Nothing Is Real](#empirical-se-nothing-is-real) · Part 3: [How to Think Empirically](#empirical-se-how-to-think).*
