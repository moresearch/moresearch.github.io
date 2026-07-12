---
title: Lehman's Software Evolution
date: 2026-07-12
slug: lehmans-laws
summary: "Meir Lehman observed that software doesn't age — it evolves according to laws as regular as thermodynamics. E-type programs change or die. Complexity increases unless you fight it. The process is self-regulating whether you like it or not. Maintenance is not a phase. It is the process."
tags: lehman, software-evolution, complexity, entropy, laws, maintenance
---

In 1980, Meir Lehman published a paper that should have changed how the industry thinks about software maintenance. It didn't. "Programs, Life Cycles, and Laws of Software Evolution" introduced a classification of programs and a set of laws governing their evolution. The laws were derived from measurement, not opinion. They described what software *does*, not what anyone wished it did.

> "We shall, in fact, argue that the need for continuous change is intrinsic to the nature of computer usage."

Forty-six years later, the laws hold. The industry still acts surprised when its codebases become harder to change, more complex, less maintainable. Lehman explained why. The explanation is still correct. The surprise is still unwarranted.

> "The resultant evolution of software appears to be driven and controlled by human decision, managerial edict, and programmer judgement. Yet as shown by extended studies, measures of its evolution display patterns, regularity and trends that suggest an underlying dynamics."

Individual decisions feel local and independent. The aggregate is regular. The aggregate is law-like. That is Lehman's central discovery. Software evolution is not random. It is not controlled. It is regular — and the regularity persists regardless of what anyone intends.

## The program types

Lehman's first contribution was a classification. Not all programs are the same kind of thing. Treating them as if they are is the root error.

**S-type (Specification-type).** The problem has a complete, formal specification. Correctness can be proven. Changes are limited to efficiency or clarity. A sorting algorithm. The Eight Queens puzzle. S-type programs don't evolve. They are replaced.

> "Programs whose function is formally defined by and derivable from a specification."

**P-type (Problem-type).** The problem can be stated formally, but a perfect solution is infeasible. The program uses heuristics or approximations. Acceptability is judged against the real world. Chess engines. Weather prediction. P-type programs evolve as heuristics improve.

**E-type (Embedded-type).** The program mechanizes a human or societal activity. It becomes part of the world it models. Deploying the program changes user behavior, which changes requirements, which changes the program. The problem cannot be precisely formulated — it involves judgment. Operating systems. Business software. Trading systems. Your job.

> "The program has become a part of the world it models, it is embedded in it."

E-type programs *must* evolve. They have no final state.

> "E-programs change because the real-world changes... but E-programs can also be the cause of that change in the real world."

You deploy the software. Users change their behavior. The changed behavior creates new requirements. The new requirements change the software. The changed software changes behavior again. This is not a bug in the process. This is the process. You are not building a tool. You are participating in a feedback loop. The loop has no end. The loop *is* the work.

Dijkstra, from a different tradition entirely, identified the same confusion in the language we use:

> "Unfathomed misunderstanding is further revealed by the term 'software maintenance', as a result of which many people continue to believe that programs — and even programming languages themselves — are subject to wear and tear. Your car needs maintenance too, doesn't it?"

Software doesn't wear out. The environment changes around it. "Maintenance" implies restoring something to its original condition. Software evolution is about changing something to meet conditions that didn't exist when it was built. The word is wrong. The concept is wrong. The industry's entire budgeting model — "build, then maintain" — is built on the wrong concept.

## The eight laws

Lehman published the first three laws in 1974 with Belady, expanded to five in 1980, and codified all eight by 1996. Each was derived from measurement of real systems — OS/360 first, then others. These are observations of statistical regularity. You can ignore them. You cannot make them false.

### I. Continuing Change (1974)

> "A program that is used and that as an implementation of its specification reflects some other reality, undergoes continual change or becomes progressively less useful. The change or decay process continues until it is judged more cost effective to replace the system with a recreated version."

An E-type system must be continually adapted or it becomes progressively less satisfactory. The environment changes — user needs, regulations, platforms, security threats. The software changes with it or becomes irrelevant. There is no third option. You do not finish an E-type system. You stop working on it. The only question is whether you stop because it was replaced or because it was abandoned.

Brooks, independently, reached the identical conclusion in *No Silver Bullet* (1986):

> "All successful software gets changed. Software is embedded in a cultural matrix of users, laws, and hardware — all of which change continually."

Two researchers, different methods, same observation. Software that matters changes. Software that doesn't change doesn't matter. The change is not a failure mode. It is evidence that the software is doing its job — reflecting a world that moves. The cost of change is not overhead. It is the work.

Lehman was even more direct about the economic consequence:

> "Assessments of the economic viability of a program must include total lifetime costs and their life cycle distribution, and not be based exclusively on the initial development costs."

If you budget only for building and not for changing, you have budgeted for failure. The change is not optional. The budget for it is not discretionary. The economics that ignore change are wrong on their own terms. They produce numbers that look good at project start and catastrophic at year five. The numbers were always catastrophic. The accounting hid it.

### II. Increasing Complexity (1974)

> "As an evolving program is continually changed, its complexity, reflecting deteriorating structure, increases unless work is done to maintain or reduce it."

This is Lehman's most practically important law. The default direction is toward disorder. Every change increases complexity *unless explicit effort is made to reduce it*.

Lehman and Belady originally called this entropy, not complexity. The 1971 IBM report states the thermodynamic intuition directly:

> "The addition of any function not visualized in the original design will inevitably degenerate structure. Repairs also, will tend to cause deviation from structural regularity since, except under conditions of the strictest control, any repair or patch will be made in the simplest and quickest way. No search will be made for a fix that maintains structural integrity."

This is the most damning sentence in the entire Lehman corpus. Repairs are made in the simplest and quickest way. No search is made for a fix that maintains structural integrity. The deadline applies pressure. The fix is local. The structure degrades. The degradation is invisible at the time of the fix. It becomes visible later, when the next fix is harder because the structure is weaker. The cycle compounds. The entropy accumulates. Nobody intended it. The process produced it.

> "All repairs tend to destroy the structure, to increase the entropy and disorder of the system. Less and less effort is spent fixing original design flaws; more and more is spent on fixing flaws introduced in earlier fixes. As time passes, the system becomes less and less well ordered."

The effort shifts. At first, you fix the original design. Then you fix the fixes. Then you fix the fixes of the fixes. Each layer adds entropy. Each layer makes the next layer more likely. The system is not just getting more complex. It is getting more complex *in a way that accelerates further complexity*. This is a positive feedback loop driving toward disorder. Physics has a name for this. Lehman borrowed it.

In the 1976 IBM Systems Journal paper, Belady and Lehman formalized it:

> "The entropy of a system (its unstructuredness) increases with time, unless specific work is executed to maintain or reduce it."

Entropy. Not complexity. Unstructuredness. The loss of form. The drift toward chaos. Specific work is required to maintain structure. If you are not doing that specific work — if all your effort goes to features and fixes — the structure is degrading. You may not notice. The degradation is gradual. By the time it is obvious, it is expensive.

Parnas, independently, provided the mechanism for fighting Law II: information hiding. Hide volatile design decisions behind stable interfaces. Contain the change inside the module. The interface stays clean. The rest of the system doesn't accumulate the complexity of the change. Lehman described the problem. Parnas described the defense. Both were published in the early 1970s. Most production code follows neither.

Later, Parnas introduced the concept of *software aging* (1994), identifying two causes that directly mirror Lehman's Laws:

1. **Lack of movement** — failure to adapt to environmental change (Law I violation)
2. **Ignorant surgery** — changes made without proper understanding of the system (Law II accelerator)

Parnas and Lehman converge on the same insight from different angles: software doesn't physically decay. Its structure degrades through a series of individually reasonable, collectively destructive changes. Each change made sense at the time. The accumulation makes no sense at all.

Brooks separated complexity into essential and accidental in *No Silver Bullet*:

> "The complexity of software is an essential property, not an accidental one. Descriptions of a software entity that abstract away its complexity often abstract away its essence."

> "From the complexity comes the difficulty of communication among team members, which leads to product flaws, cost overruns, and schedule delays."

Essential complexity is inherent in the problem. Accidental complexity is imposed by our solutions. Law II describes the accumulation of accidental complexity over time. Each quick fix adds a little more. Unless you fight it. Brooks: "There is no silver bullet." Lehman: "Complexity increases unless you work to reduce it." Same argument. Different vocabularies.

> "The unit cost of change must initially be made as low as possible and its growth, as the system ages, minimized. Programs must be made more alterable, and the alterability maintained throughout their lifetime. The change process itself must be planned and controlled."

This is Lehman's practical prescription. Make the system alterable. Maintain alterability. Plan the change process. Control it. Most organizations do none of these. They make the system work, then react to change requests, then wonder why each change is harder than the last. The alterability was never designed in. It was assumed. The assumption was wrong.

### III. Self-Regulation (1974)

> "Program evolution is subject to a dynamics which makes the programming process, and hence measures of global project and system attributes, self-regulating with statistically determinable trends and invariances."

Lehman observed that OS/360's growth data showed patterns "typical of a self-stabilising process with positive and negative feedback loops." The rate of system growth was self-regulatory despite varying budgets, varying team sizes, varying management attitudes. The process finds its equilibrium.

> "Individual decisions may appear localised and independent, but their aggregation, moderated by many feedback relationships, produces overall system responses that are regular and often normally distributed."

This is the law that should make managers uncomfortable. You cannot accelerate software evolution by adding resources. The process has its own pace, determined by feedback loops among users, developers, and the codebase. Brooks's Law — "adding manpower to a late software project makes it later" — is a special case of Lehman's Self-Regulation. The process resists perturbation. It returns to its natural rate.

### IV. Conservation of Organizational Stability (1978)

> "During the active life of a program the global activity rate in a programming project is statistically invariant."

The amount of work a team produces per release is roughly constant. Change the team. Change the tools. The output stays approximately the same. This is not about individual productivity. It is a statistical observation about organizations. The organization has a natural throughput. It can be measured. It cannot be wished higher.

### V. Conservation of Familiarity (1978/1980)

> "During the active life of a program the release content (changes, additions, deletions) of the successive releases of an evolving program is statistically invariant."

Each release contains roughly the same amount of change. Not because anyone plans it. Because the organization can only absorb so much change at once. Ship more — quality drops, bugs increase, the next release is smaller to compensate. The average holds. Law III operating through Law V.

### VI. Continuing Growth (1991)

The functional content of an E-type system must grow to maintain user satisfaction. Users demand more features. Growth is not optional. Managing growth is the entire job. A system that doesn't grow is a system users abandon for one that does. The growth is demanded. The cost of growth is Law II. The two laws together: you must grow, and growing increases complexity. This is the tension that defines the economics of software.

### VII. Declining Quality (1996)

The quality of an E-type system declines unless it is rigorously maintained and adapted to environmental change. Quality is relative to the environment. The environment moves. What was excellent in 2019 is unmaintainable legacy in 2026, even if the code hasn't changed. The standards rose. The code stayed where it was.

Lehman on the practical consequence:

> "Top-level managerial pressure to apply life-cycle evaluation is therefore essential if a development and maintenance process is to be attained that continuously achieves desired overall balance between the short- and long-term objectives of the organization."

Short-term: ship the feature. Long-term: maintain the structure. The two are in tension. Management must enforce the balance. If management only rewards shipping, structure degrades. If management only rewards structure, nothing ships. The balance is the job. Most organizations don't recognize it as a job. They recognize shipping. They wonder why quality declines.

### VIII. Feedback System (1996)

> "The global software process that includes technical, business, marketing, user and other activities constitutes a multi-loop, multi-level feedback system. To change the characteristics of such a system requires one to consider, design or adapt and tune both forward and feedback paths to achieve the desired changes in externally visible behaviour."

Software evolution is a multi-loop, multi-agent feedback system. Users provide feedback. Developers respond. The system changes. The changes generate new feedback. The loops interact. Emergent behavior. You cannot control it. You can only participate in it.

> "Current world-wide process models and improvement activities focus primarily on the forward technical path and overlook the many feedback paths and the constraints that they impose on improvement."

Most process improvement focuses on the forward path: build better, test better, deploy better. Lehman's point is that the *feedback* paths — how users react, how the market responds, how the organization learns — are equally important and almost entirely ignored. You improved the build pipeline. You didn't improve the organization's ability to learn what to build. The forward path is faster. The feedback path is still broken. The system as a whole is not improved. It is accelerated toward the wrong destination.

## What Lehman and his peers agreed on

Lehman, Brooks, Parnas, and Dijkstra arrived at the same destination from different starting points. Lehman measured systems and derived statistical laws. Brooks managed systems and derived engineering principles. Parnas decomposed systems and derived design criteria. Dijkstra thought about systems and derived epistemological critiques. All four concluded: change is inevitable, complexity is the enemy, and the process cannot be controlled — only influenced.

Lehman in 1980:

> "Any program is a model of a model within a theory of a model of an abstraction of some portion of the world or of some universe of discourse."

A program is not the world. It is a model of a model of a theory of a model of an abstraction. That is four levels of indirection. Each level introduces error. Each level changes independently. The program must track changes across all four. This is epistemologically ambitious. It is also what every business application attempts. Lehman understood the difficulty. Most project plans don't.

Brooks in 1986:

> "The hardest single part of building a software system is deciding precisely what to build. Therefore the most important function that the software builder performs for the client is the iterative extraction and refinement of the product requirements."

The requirements are not known. They are discovered. The discovery process is iterative. Lehman's Laws are the dynamics of that iteration at scale. The iteration doesn't stop at version 1.0. It continues for the life of the system. The life of the system is the iteration.

> "There is no royal road, but there is a road." — Brooks

Dijkstra in 1972:

> "The major cause of the software crisis is that the machines have become several orders of magnitude more powerful. As long as there were no machines, programming was no problem at all; when we had a few weak computers, programming became a mild problem, and now we have gigantic computers, programming has become an equally gigantic problem."

The machines got faster. The problems got bigger. The complexity grew with the capability. Lehman's Laws describe the dynamics of that growth. More powerful machines don't reduce complexity. They enable larger systems, which are more complex. The complexity is the thing. The machine is the substrate.

Weinberg, in *The Psychology of Computer Programming* (1971), identified the human dimension that Lehman's statistical laws abstract over: programming is done by people, in organizations, under pressure. The pressure to ship degrades structure. The degradation is organizational before it is technical. Lehman's laws describe the aggregate. Weinberg describes the individual decisions that produce the aggregate. Both are necessary. Neither is sufficient alone.

## The entropy connection

In 2023, Torres, Baltes, Treude, and Wagner applied information theory to software evolution. Two entropy definitions — structural (dependency graph) and textual (compression ratio) — were tracked across 25 open source projects.

> "Both entropy measures display weak and unstable correlations with other complexity metrics."

Entropy captures something traditional metrics miss: the *information content* of the codebase and how it changes. Lehman's Law II said complexity increases unless work is done to reduce it. Lehman and Belady originally called it entropy. The Torres paper returns to the original framing and gives it measurement. Structural entropy rises when the dependency graph becomes more interconnected. Textual entropy rises when the codebase becomes more information-dense. Both rise when complexity, in Lehman's sense, is accumulating.

> "An unexpected high frequency of events where there is considerable change in the information content of the project."

These are surprisal events. Commits where the information structure jumps significantly relative to history. A refactoring. A deletion that simplifies the dependency graph. A new module that reconfigures everything. These are the commits where Law II is being obeyed or violated. Traditional metrics see size. Entropy sees structural impact. Lehman predicted these events would matter. The Torres paper gives us a way to flag them. A system's structural entropy, tracked over time, is a measure of whether Law II is being respected. Rising entropy: complexity accumulating. Stable entropy: the team is doing the work. Dropping entropy: the team is actively simplifying. Most codebases rise. Few are stable. Almost none drop. The measurement confirms what Lehman predicted in 1974.

## What to do

Lehman's Laws are not despair. They are realism. The laws describe what happens. How you respond is a choice.

**Accept that maintenance is the process.** Law I says you will change the system forever. Not because you built it wrong. Because the world changes. Stop treating maintenance as a phase after development. Development is the first phase of maintenance. Everything after is the rest of it.

**Budget for complexity reduction.** Law II says complexity increases unless you fight it. Fighting costs time and money. Budget for it. Refactoring is not a sign of past failure. It is the necessary work of preventing future failure. If your organization treats it as a luxury, your complexity is rising at a rate determined by your change velocity and your neglect of structure. Measure it. Or choose not to know. The complexity doesn't care.

**Measure entropy.** The tools exist. Dependency graph entropy. Compression-ratio entropy. Track them. Flag surprisal commits. These are the commits where the information structure changes significantly — and they are the commits your size-based review will miss. Size is easy. Structure is more important.

**Hire for taste.** All eight laws describe a system that resists management. The only counterforce is good design decisions, made early and defended persistently. That means good designers. Brooks spent his career arguing this. Lehman's Laws are the evidence Brooks was right. The process won't save you. Only taste will.

> "The central question in how to improve the software art, centers, as it always has, on people." — Brooks, 1986

---

**References:**
- M.M. Lehman, "Programs, Life Cycles, and Laws of Software Evolution," *Proceedings of the IEEE*, Vol. 68, No. 9, September 1980, pp. 1060-1076.
- M.M. Lehman, "Laws of Software Evolution Revisited," *Proceedings of the 5th European Workshop on Software Process Technology*, 1996.
- M.M. Lehman, "Feedback in the Software Evolution Process," *Information and Software Technology*, Vol. 38, 1996.
- L.A. Belady and M.M. Lehman, "A Model of Large Program Development," *IBM Systems Journal*, Vol. 15, No. 3, 1976.
- L.A. Belady and M.M. Lehman, "Programming System Dynamics or The Metadynamics of Systems in Maintenance and Growth," IBM Research Report RC3546, 1971.
- M.M. Lehman and L.A. Belady, *Program Evolution: Processes of Software Change*, Academic Press, 1985.
- Frederick P. Brooks, Jr., "No Silver Bullet: Essence and Accidents of Software Engineering," *Computer Magazine*, April 1987.
- Frederick P. Brooks, Jr., *The Mythical Man-Month*, Addison-Wesley, 1975 (Anniversary Edition, 1995).
- David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, Vol. 15, No. 12, December 1972.
- David L. Parnas, "Software Aging," *Proceedings of the 16th International Conference on Software Engineering*, 1994.
- Edsger W. Dijkstra, "The Humble Programmer," *Communications of the ACM*, Vol. 15, No. 10, October 1972 (Turing Award Lecture).
- Edsger W. Dijkstra, "On the cruelty of really teaching computing science," EWD 1036, 1988.
- Gerald M. Weinberg, *The Psychology of Computer Programming*, Van Nostrand Reinhold, 1971.
- Adriano Torres, Sebastian Baltes, Christoph Treude, Markus Wagner, "Applying Information Theory to Software Evolution," NLBSE 2023. [arXiv:2303.13729](https://arxiv.org/abs/2303.13729)
