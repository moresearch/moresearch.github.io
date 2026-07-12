---
title: "Lehman's Laws"
date: 2026-07-12
slug: lehmans-laws
summary: "Meir Lehman observed that software doesn't just age — it evolves according to laws as regular as thermodynamics. E-type programs change or die. Complexity increases unless you fight it. The process is self-regulating whether you like it or not."
tags: lehman, software-evolution, complexity, entropy, laws
---

In 1980, Meir Lehman published a paper that should have changed how the industry thinks about software maintenance. It didn't. The paper — "Programs, Life Cycles, and Laws of Software Evolution" — introduced a classification of programs and a set of laws governing their evolution. The laws were based on measurement, not opinion. They described what software *does*, not what anyone wished it did.

Forty-six years later, the laws hold. The industry still acts surprised when its codebases become harder to change, more complex, and eventually unmaintainable. Lehman explained why this happens. The explanation is still correct. The surprise is still unwarranted.

## The program types

Lehman's first contribution was a classification. Not all programs are the same kind of thing. Treating them as if they are is the root error of most software process thinking.

**S-type (Specification-type).** The problem has a complete, formal specification. Correctness can be proven. Changes are limited to efficiency or clarity. Examples: sorting algorithms, the Eight Queens puzzle, a calculator. S-type programs don't evolve. They are replaced.

> "Programs whose function is formally defined by and derivable from a specification."

**P-type (Problem-type).** The problem can be stated formally, but a perfect solution is computationally infeasible. The program uses heuristics, approximations, or simplifications. Acceptability is judged against the real world. Examples: chess engines, weather prediction, scheduling. P-type programs evolve as heuristics improve.

**E-type (Embedded-type).** The program mechanizes a human or societal activity. It becomes part of the world it models. The act of deploying the program changes user behavior, which changes requirements, which changes the program. The problem cannot be precisely formulated — it involves judgment. Examples: operating systems, business software, inventory management, trading systems. E-type programs *must* evolve. They have no final state.

> "Programs that mechanize a human or societal activity and are embedded in the real world."

E-type programs are where Lehman's Laws apply. They are also essentially all of the software that anyone pays anyone to work on. S-type programs are undergraduate assignments. E-type programs are your job.

## The eight laws

Lehman published the first three laws in 1974, expanded to five in 1980, and codified all eight by 1996. Each was derived from measurement of real systems — OS/360, and later others. These are not principles. They are observations of statistical regularity. They describe what happens. You can ignore them. You cannot make them false.

### I. Continuing Change (1974)

> "A program that is used and that as an implementation of its specification reflects some other reality, undergoes continual change or becomes progressively less useful. The change or decay process continues until it is judged more cost effective to replace the system with a recreated version."

An E-type system must be continually adapted or it becomes progressively less satisfactory. The environment changes — user needs, regulations, platforms, security threats. The software either changes with it or becomes irrelevant. There is no third option. You do not finish an E-type system. You stop working on it.

Brooks, independently, reached the same conclusion in *No Silver Bullet* (1986):

> "All successful software gets changed. Software is embedded in a cultural matrix of users, laws, and hardware — all of which change continually."

Two researchers, different methods, same observation. Software that matters changes. Software that doesn't change doesn't matter. The change is not a sign of failure. It is a sign that the software is doing its job — reflecting a world that itself is changing. The cost of change is not a bug in the process. It is the process.

### II. Increasing Complexity (1974)

> "As an evolving program is continually changed, its complexity, reflecting deteriorating structure, increases unless work is done to maintain or reduce it."

This is Lehman's most practically important law. Every change to a system increases its complexity *unless explicit effort is made to reduce it*. The default direction is toward disorder. The default is entropy.

Lehman originally used the word "entropy" for this law, not "complexity." He changed the term later to align with the complexity research of the time, but the original intuition was thermodynamic: systems drift toward disorder unless energy is expended to maintain order. Software is a system. It obeys the same logic. Every commit that adds a feature without simplifying the structure increases total complexity. Every quick fix that works around an existing abstraction rather than extending it properly increases complexity. Every deadline that forces "we'll clean it up later" increases complexity. Later never comes.

Parnas, Lehman's contemporary, provided the mechanism for fighting Law II: information hiding. If you hide volatile design decisions behind stable interfaces, you contain the complexity increase. The change happens inside the module. The interface stays clean. The rest of the system doesn't accumulate the complexity of the change. Parnas was not directly responding to Lehman, but the fit is exact. Lehman described the problem. Parnas described the defense. Both were published in the early 1970s. Neither is widely followed.

> "I believe the hard part of building software to be the specification, design, and testing of this conceptual construct, not the labor of representing it and testing the fidelity of the representation." — Brooks, *No Silver Bullet*, 1986

Brooks separated complexity into essential and accidental. Essential complexity is inherent in the problem. Accidental complexity is imposed by our solutions. Law II describes the accumulation of accidental complexity over time. Every change adds a little more. Unless you fight it.

### III. Self-Regulation (1974)

> "Program evolution is subject to a dynamics which makes the programming process, and hence measures of global project and system attributes, self-regulating with statistically determinable trends and invariances."

Lehman observed that OS/360's growth data showed patterns "typical of a self-stabilising process with positive and negative feedback loops." The rate of system growth was self-regulatory despite varying budgets, varying team sizes, varying management attitudes, and changing release intervals. The process finds its own equilibrium regardless of what management intends.

This is the law that should make managers uncomfortable. You cannot accelerate software evolution by adding resources. The process has its own pace, determined by feedback loops among users, developers, and the codebase. Brooks's Law — "adding manpower to a late software project makes it later" — is a special case of Lehman's Self-Regulation. The process resists external perturbation. It returns to its natural rate.

### IV. Conservation of Organizational Stability (1978)

> "During the active life of a program the global activity rate in a programming project is statistically invariant."

The amount of work a team can produce per release is roughly constant. You can change the team. You can change the tools. The output stays the same. This is not a claim about individual productivity. It is a statistical observation about organizations. The organization has a natural throughput. It can be measured. It cannot be wished higher.

### V. Conservation of Familiarity (1978/1980)

> "During the active life of a program the release content (changes, additions, deletions) of the successive releases of an evolving program is statistically invariant."

Each release contains roughly the same amount of change. Not because anyone plans it that way, but because the organization can only absorb so much change at once. If you try to ship more, the quality drops, the bugs increase, and the next release is smaller to compensate. The system self-corrects. The average holds. This is Law III operating through Law V.

### VI. Continuing Growth (1991)

The functional content of an E-type system must grow to maintain user satisfaction. Users demand more features over time, not fewer. The system that doesn't grow is the system that users abandon for one that does. Growth is not optional. Managing growth is the entire job.

### VII. Declining Quality (1996)

The quality of an E-type system declines unless it is rigorously maintained and adapted to environmental change. Quality is not static. It decays. What was "high quality" in 2019 is "unmaintainable legacy" in 2026, even if the code hasn't changed. The environment changed around it. The standards rose. The code stayed where it was. Quality is relative to the environment. The environment moves.

### VIII. Feedback System (1996)

Software evolution is a multi-loop, multi-agent feedback system. Users provide feedback. Developers respond. The system changes. The changes generate new feedback. The loops interact. The system's behavior emerges from these interactions, not from any single decision or plan. You cannot control software evolution. You can only participate in the feedback loops and hope your participation makes things better rather than worse. This is the law that makes all the other laws irreducible: they are properties of a feedback system, not of a plan.

## What Lehman and Brooks agreed on

Lehman and Brooks arrived at the same destination from different starting points. Lehman measured systems and derived statistical laws. Brooks managed systems and derived engineering principles. Both concluded that software evolution is inevitable, complexity is the enemy, and the process cannot be controlled — only influenced.

Brooks in 1986:

> "The complexity of software is an essential property, not an accidental one. Descriptions of a software entity that abstract away its complexity often abstract away its essence."

Lehman in 1980:

> "The observed fact that the number of decisions driving the process of evolution, the many feedback paths, the checks and balances of organizations, human interactions in the process, reactions to usage, the rigidity of program code, all combine to yield statistically regular behavior."

Complexity is essential. Evolution is inevitable. The process is self-regulating. These are not problems to solve. They are conditions to live with. The industry has spent decades trying to solve them — with better tools, better processes, better architectures — and has succeeded mainly in changing the form of the complexity, not its presence. Brooks predicted this. Lehman measured it.

## The entropy connection

In 2023, Torres, Baltes, Treude, and Wagner published a small paper applying information theory to software evolution. They applied two definitions of entropy — structural (dependency graph) and textual (compression ratio) — to the commit histories of 25 open source projects.

> "Both entropy measures display weak and unstable correlations with other complexity metrics."

Entropy captures something that traditional complexity metrics miss: the *information content* of the codebase, and how it changes. Lehman's Law II said complexity increases unless work is done to reduce it. The entropy paper provides a way to *measure* that increase — not just at release boundaries, but at every commit. Structural entropy rises when the dependency graph becomes more interconnected. Textual entropy rises when the codebase becomes more information-dense. Both rise when complexity, in Lehman's sense, is accumulating.

> "An unexpected high frequency of events where there is considerable change in the information content of the project."

These are surprisal events. Commits where the information structure of the codebase jumps significantly relative to its history. A major refactoring. A deletion that simplifies the dependency graph. A new module that reconfigures the architecture. These are the commits where Law II is either being obeyed — complexity reduced — or violated spectacularly — complexity added in a single large structural change. Traditional metrics see the size of the commit. Entropy sees the structural impact. Lehman's laws predicted these events would matter. The entropy paper gives us a way to flag them.

Lehman originally called Law II "entropy" before changing the term to "complexity." The Torres paper returns to the original framing and gives it measurement. A system's structural entropy, tracked over time, is a direct measure of whether Law II is being respected or ignored. If entropy is rising, complexity is accumulating. If entropy is stable, the team is doing the work to maintain structure. If entropy drops — the rarest and most encouraging case — the team is actively simplifying. Most codebases rise. Few are stable. Almost none drop. The measurement confirms what Lehman predicted in 1974.

## What to do

Lehman's Laws are not a call to despair. They are a call to realism. The laws describe what happens. How you respond is a choice.

**Accept that change is not failure.** Law I says you will be changing the system forever. This is not because you built it wrong. This is because the world changes. Stop treating maintenance as a phase that follows development. Development is the first phase of maintenance. Everything after is the rest of it.

**Budget for complexity reduction.** Law II says complexity increases unless you fight it. Fighting it costs time and money. Budget for it. Refactoring is not a sign of past failure. It is the necessary work of preventing future failure. If your organization treats refactoring as a luxury, your complexity is rising. It is rising at a rate determined by your change velocity and your neglect of structure. You can measure this. If you don't, you're choosing ignorance, not safety.

**Measure entropy.** The tools exist. Structural entropy from your dependency graph. Textual entropy from compression ratios. Track them. Flag surprisal commits — the ones where the information structure changes dramatically. These are the commits most likely to contain architectural change and most likely to be overlooked by review processes that flag size. Size is easy to measure. Structure change is more important.

**Hire for taste.** All eight laws describe a system that resists management. The only counterforce is good design decisions, made early and defended persistently. That means good designers. Brooks spent his career arguing this. Lehman's Laws are the evidence that Brooks was right. The process won't save you. Only taste will.

> "The central question in how to improve the software art, centers, as it always has, on people." — Brooks, *No Silver Bullet*, 1986

---

**References:**
- M.M. Lehman, "Programs, Life Cycles, and Laws of Software Evolution," *Proceedings of the IEEE*, Vol. 68, No. 9, September 1980, pp. 1060-1076.
- M.M. Lehman, "Laws of Software Evolution Revisited," *Proceedings of the 5th European Workshop on Software Process Technology*, 1996.
- Frederick P. Brooks, Jr., "No Silver Bullet: Essence and Accidents of Software Engineering," *Computer Magazine*, April 1987.
- Frederick P. Brooks, Jr., *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975 (Anniversary Edition, 1995).
- David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, Vol. 15, No. 12, December 1972.
- Adriano Torres, Sebastian Baltes, Christoph Treude, Markus Wagner, "Applying Information Theory to Software Evolution," NLBSE 2023. [arXiv:2303.13729](https://arxiv.org/abs/2303.13729)
