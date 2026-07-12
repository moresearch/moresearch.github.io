---
title: "entropy as a code metric"
date: 2026-07-12
slug: entropy-software-evolution
summary: "A 2023 paper applies information theory to software evolution. Two entropy measures track how codebases change. Both capture something traditional complexity metrics miss: the information content of a codebase, and the moments when it surprises you."
tags: entropy, information-theory, software-evolution, complexity, metrics
---

In 2023, Torres, Baltes, Treude, and Wagner published a small paper with a large idea: what if you measured software complexity not by counting lines, branches, or dependencies, but by measuring the *information content* of the code itself?

The paper is eight pages. It is an early-stage investigation. It is also the kind of work that makes you reconsider what "complexity" means. Most complexity metrics — cyclomatic complexity, lines of code, coupling between objects — measure structural properties of code. They tell you how big the code is, how many paths through it, how interconnected its parts. They don't tell you how much *information* it contains, or how that information changes over time.

Entropy does.

## Two kinds of entropy

The paper applies two definitions of entropy to code:

**Structural entropy.** Treat the codebase as a graph. Each file is a node. Each dependency (import, include, reference) is an edge. The entropy of the graph measures how uniformly connected the system is. A system where every file depends on every other file has maximum structural entropy — maximum disorder. A system where files are organized into clean, independent modules with minimal cross-dependencies has lower entropy. Same functionality. Different information structure.

**Textual entropy.** Treat the code as text. Compress it. The compressed size, relative to the uncompressed size, is a measure of how much information the text contains. Redundant code compresses well — low entropy. Code where every line carries distinct information compresses poorly — high entropy. This is not about quality. It is about information density. A verbose codebase with repetitive patterns has lower textual entropy than a dense, terse one. Neither is inherently better. The metric just tells you which you have.

> "Both entropy measures display weak and unstable correlations with other complexity metrics."

This is the key finding. Entropy is not just another complexity metric in disguise. It captures something different. Traditional complexity metrics measure *how complicated the code is to work with*. Entropy measures *how much information the code contains*. These are related but distinct. A system can be information-dense without being complex. A system can be simple in structure but bloated in information content — verbose, repetitive, full of boilerplate. Traditional metrics see simplicity. Entropy sees bloat.

> "Structural and textual entropy generally are highly correlated."

The two entropy measures, though derived completely differently — one from graph structure, one from text compression — move together over time. When structural entropy rises, textual entropy tends to rise with it. When the codebase becomes more interconnected, it also tends to become more information-dense. The two measures converge on the same underlying signal: the system is accumulating information. Whether that information is *valuable* — whether it represents solved problems or accumulated cruft — is a separate question. Entropy doesn't judge. It measures.

## Surprisal: when the codebase jumps

The most intriguing finding is about outlier events. The authors tracked entropy over the commit history of 25 open source projects and looked for commits where the entropy change was unusually large — where the information content of the entire codebase shifted significantly in a single change.

> "An unexpected high frequency of events where there is considerable change in the information content of the project."

These are surprisal events. A commit lands that is not just large — large commits happen — but *structurally unexpected* relative to the prior history. The information content of the codebase changes in a way that the historical trend didn't predict. A major refactoring. A new module that reconfigures the dependency graph. A deletion that simplifies the structure dramatically. The entropy doesn't just change. It jumps.

The authors suggest these outliers could form the basis of a "surprisal" metric — a way to automatically flag commits that are historically unusual, regardless of their size or author. A commit that adds 5,000 lines to an established pattern might be large but not surprising. A commit that deletes 50 lines and rewires the dependency graph might be small but deeply surprising. Traditional metrics see size. Entropy sees the rewiring. Size is easy to measure. Structure change is harder. Surprisal captures structure change.

## Why this matters

Software complexity metrics have a long history of being simultaneously useful and misleading. Cyclomatic complexity correlates with bug density — but so does file size, and the two are correlated with each other, and disentangling cause from correlation is hard. Coupling metrics tell you about interconnections but not about whether those interconnections are stable or volatile. Lines of code is the most commonly used metric and the most commonly abused — it measures size, not complexity, and is gamed by the simple act of writing verbose code.

Entropy offers something different: a measure of information content that is not gamed by style, not reducible to size, and not correlated with traditional metrics. It is a *complementary* signal. A codebase whose traditional complexity is low but whose entropy is high is a codebase that is simple in structure but dense in information. That might be a well-factored system with concise code. Or it might be a terse, unreadable mess. The entropy tells you to look. It doesn't tell you what you'll find.

A codebase whose traditional complexity is high but whose entropy is low is highly interconnected but information-sparse. Lots of boilerplate. Lots of repetition. The structure is tangled but the information is thin. This is the enterprise Java project with fifteen layers of indirection and the same pattern repeated in every package. Complex to navigate. Simple in information content. The metrics agree that something is wrong. They disagree about what.

## The connection to Henney

This paper lands directly on themes explored in the Brooks and Henney series on this blog. Henney argued that less code is less risk — but also that there is an optimal code size, not zero. Entropy gives you a way to measure where you are on that curve. A codebase with high textual entropy is dense with information. Whether that information is essential complexity (the problem is hard) or accidental complexity (the code is unclear) is a judgment call. The metric flags the density. The human makes the call.

Brooks argued that conceptual integrity — the system feeling like one mind designed it — is the most important property. A system designed by committee will have higher structural entropy: more interconnections, less clean modularity, more uniform distribution of dependencies. The entropy didn't cause the lack of integrity. It measured it. The measurement is not the problem. But it makes the problem visible in a way that individual code reviews might miss. A code review sees one file. Entropy sees the whole graph.

Parnas argued that modules should hide design decisions likely to change. A system that follows Parnas's criterion will have *lower* structural entropy in its dependency graph — dependencies flow to stable interfaces, not to volatile internals. A system that ignores Parnas will have higher entropy, with dependencies sprawling across module boundaries. Entropy doesn't replace Parnas. It quantifies what happens when you ignore him.

## The limits

The paper is eight pages. It studies 25 projects. The correlations with traditional metrics are weak and unstable — which is simultaneously the most interesting finding and a signal that more work is needed. Weak correlation means entropy is measuring something distinct. Unstable correlation means the relationship changes across projects and over time. Neither is a clean result. Both are honest. The authors don't oversell. They present initial evidence and suggest directions. That is what a workshop paper should do.

The larger question is whether entropy can become *actionable*. A metric that tells you something interesting about your codebase is academically valuable. A metric that tells you *which commit to investigate* — flagging surprisal events for review, tracking entropy trends as an early warning of architectural decay — is practically valuable. The paper points toward the second but doesn't reach it. The path is visible. The work remains to be done.

## What to take from it

If you track metrics for your codebase — and you should — entropy is probably not among them. It should be. The tools exist. The theory is sound. The signal is distinct from everything else you're measuring. You can compute structural entropy from your dependency graph. You can compute textual entropy from compression ratios. Both are cheap. Both capture something your existing metrics miss.

The paper's most actionable insight is about surprisal. Flag commits where the information content of the codebase changes unusually. These are the commits where something structural happened — a refactoring, a deletion, a rearchitecture. They are the commits most likely to contain bugs and most likely to be overlooked by review processes focused on size. Size is a proxy for risk. Surprisal is closer to the thing itself. A commit that changes the structure is riskier than a commit that adds lines to an existing structure. Size-based review thresholds miss this. Entropy-based thresholds would catch it. The industry hasn't adopted them yet. It will. The paper is the first step.

---

**Reference:** Adriano Torres, Sebastian Baltes, Christoph Treude, Markus Wagner. "Applying Information Theory to Software Evolution." NLBSE 2023. [arXiv:2303.13729](https://arxiv.org/abs/2303.13729)
