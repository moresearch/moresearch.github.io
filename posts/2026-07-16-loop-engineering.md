---
title: Loop Engineering is what the NATO conference asked for in 1968
date: 2026-07-16
slug: loop-engineering
summary: "Loop Engineering — designing systems that prompt agents instead of prompting them yourself — is 2026's dominant AI paradigm. The idea is not new. It is the practical realization of a principle that has been at the core of software engineering since the field was named in Garmisch, 1968: feedback is everything. Build the loop. Then get out of it."
tags: loop-engineering, agents, nato, software-engineering, feedback-loops, history, dijkstra, brooks
---

In June 2026, Peter Steinberger tweeted:

> "Here's your monthly reminder that you shouldn't be prompting coding agents anymore. You should be designing loops that prompt your agents."

Two million views. Within days, Claude Code's Boris Cherny said at a public event:

> "I don't prompt Claude anymore. I have loops running that prompt Claude and figuring out what to do. My job is to write loops."

Cherny wasn't speaking hypothetically. He reported that 100% of contributions to the Claude Code repository in the prior thirty days — 259 merged PRs — had been written by Claude Code itself, driven by loops he had designed. He deleted his IDE in November 2025 and hasn't reopened one since. The loops wake up, read GitHub issues, scan Twitter for feedback, parse Slack, decide what to build, and build it. Cherny watches the output. He does not participate in the loop.

Google's Addy Osmani gave it a name: **Loop Engineering**.

> "Loop engineering is replacing yourself as the person who prompts the agent. You design the system that does it instead."

By July, the term was everywhere — blog posts, conference talks, the obligatory LinkedIn thought leaders. The old thing was Prompt Engineering: you write better and better prompts to get better and better outputs. The new thing is Loop Engineering: you build a system that writes the prompts, runs the agent, checks the output, and feeds the result back into the next cycle. You don't talk to the agent. You talk to the loop. The loop talks to the agent. You design the loop.

The industry treated this as a breakthrough. It is not a breakthrough. It is a homecoming. Software engineering was invented to solve exactly this problem. The name was coined at a conference in 1968. The conference report said, on page 11, four sentences in: **"The need for feedback was stressed many times."** Everything since has been an attempt to build better loops.

## Garmisch, 1968

In October 1968, about fifty people from eleven countries gathered in Garmisch, Germany, for a NATO-sponsored conference. The chair was Friedrich Bauer. The topic was something nobody had a name for yet. Bauer gave it one: *software engineering*. He chose the word deliberately — provocative, aspirational, implying that software manufacture needed the same theoretical foundations and practical disciplines as the established branches of engineering. The term described a discipline that did not yet exist.

The conference report, edited by Peter Naur and Brian Randell, runs over two hundred pages. It introduced the phrase "software crisis" — the widening gap between what people wanted software to do and what they could actually build. Doug McIlroy, one of the participants, later described the editorial method — stitching together direct quotations from papers and transcribed discussions — as:

> "A triumph of misapplied quotation."

It was an affectionate jab. The report is remarkable precisely because it preserves the voices of the participants in something close to raw form. You can hear them thinking. You can hear them discovering that they all have the same problem and nobody has the solution.

The line that matters most, the one that still has not been fully absorbed, appears early:

> "The need for feedback was stressed many times."

Not discovered. Not proposed. *Stressed.* The participants already knew feedback was essential. They were stressing it because they were watching an industry ignore it. The dominant method of building software in 1968 was what J.W. Graham described during a panel on feedback through monitoring and simulation:

> "Today we tend to go on for years, with tremendous investments to find that the system, which was not well understood to start with, does not work as anticipated. We build systems like the Wright brothers built airplanes — build the whole thing, push it off the cliff, let it crash, and start over again."

Years of investment. No intermediate validation. Catastrophic late-stage discovery of fundamental problems. The feedback loop *was* the entire project lifecycle. You learned whether the system worked only when it was done. Usually it didn't. You started over. This was not a straw man. This was the state of the art in 1968. The people in the room had lived it. Some were still living it.

The conference attendees understood that this was the problem. The question was what to do about it. Kinslow pointed out that the stages everyone assumed were separate — design, production — were not separate at all:

> "The design process is iterative, as certain specifications and issues may not be realised until production of the program begins."

Some argued more aggressively: since designers make decisions during production about how to fit a design using available tools, the boundary between design and production should be *minimal*. A tight, continuous feedback loop between specification and implementation. This was 1968. They were describing continuous delivery. They were describing what we now call an inner loop. They did not have the tooling to make it real, but they had the concept. The concept was waiting fifty years for the tooling to catch up.

Alan Perlis went further. He argued that testing should not come after design — it should be woven into the act of designing. And then he said something that took three decades to fully land:

> "The critical point is that the simulation becomes the system."

The simulation becomes the system. The test harness becomes the specification. The specification is executable. The system grows through iterations of a loop where each cycle produces a working artifact that is also the specification for the next cycle. This is exactly what Test-Driven Development would become in the late 1990s. It is exactly what Loop Engineering proposes for AI agents now. The thing being looped is not a test suite. It is an agent with tools, a verifier, and a feedback mechanism. The structure is identical. The components have changed. The loop has not.

Perlis understood the significance of the report. He gave copies to his graduate students at Carnegie Mellon with the words: **"Here, read this. It will change your life."**

## The history of software engineering is the history of tightening loops

If you read the history of software engineering as a history of loops, a clear pattern emerges. Every major methodology shift was an attempt to tighten a feedback loop — to reduce the time between making a decision and discovering whether it was right.

**Structured programming (late 1960s).** Dijkstra's "Go To Statement Considered Harmful" was published in the March 1968 issue of *Communications of the ACM* — the same year as Garmisch. The argument was not about aesthetics. It was about whether you can reason about a program without running it. Dijkstra's exact reasoning is worth reading in full:

> "My second remark is that our intellectual powers are rather geared to master static relations and that our powers to visualize processes evolving in time are relatively poorly developed. For that reason we should do (as wise programmers aware of our limitations) our utmost to shorten the conceptual gap between the static program and the dynamic process, to make the correspondence between the program (spread out in text space) and the process (spread out in time) as trivial as possible."

The goto statement broke this correspondence. It made it impossible to locate yourself in the program's execution using the structure of the source text alone. You needed to also track the values of variables — but the meaning of those variables could only be understood *relative to where you were in the program*, which you couldn't determine without the variables. The loop between reading code and understanding it was broken. Structured programming restored it. The loop became local. Before: read the code, trace every possible goto, get confused, give up, run it and hope. After: read the code, understand it. The distance between the static text and the dynamic process collapsed.

This was not a stylistic preference. It was a feedback argument disguised as a language design argument. Dijkstra wanted the loop between the programmer's eye and the programmer's understanding to be as short as possible. He wanted the compiler to provide feedback — through structure, through enforceable constraints — rather than requiring the programmer to simulate execution in their head. The structured programming movement, at its core, was about tightening the comprehension loop.

**Unix philosophy (1970s).** McIlroy, the same McIlroy who called the NATO report a triumph of misapplied quotation, wrote the memo that proposed Unix pipes. His argument:

> "We should have some ways of coupling programs like garden hose — screw in another segment when it becomes necessary to massage data in another way."

Small programs that do one thing well. Text streams as universal interface. The loop: write a program, test it in isolation, compose it with others, observe the output. Each program is a closed loop. The composition is a larger loop. The feedback is immediate because the programs are small and the interface is uniform. McIlroy's pipes made the loop between idea and result as short as the time it takes to type `|`. This is the same McIlroy who appreciated the NATO report enough to gently mock it. He understood what the report was asking for. He built part of the answer.

**Test-driven development (late 1990s).** Beck's insight: write the test first, watch it fail, write the code, watch it pass, refactor. The test is the specification. The specification is executable. The loop — red, green, refactor — runs in seconds. Perlis said "the simulation becomes the system" in 1968. Beck made it mechanical thirty years later. The loop was always the idea. The infrastructure caught up.

**Continuous integration and delivery (2000s–2010s).** Fowler and Beck on CI: every commit triggers a build and a test suite. Humble and Farley on CD: every green build is potentially shippable. The loop that TDD ran in seconds on a developer's machine now runs across an entire team, continuously. The time between writing a line of code and discovering it broke something dropped from weeks to minutes. Brooks, in *No Silver Bullet* (1986), had already described the principle:

> "Using rapid prototyping as part of a planned iteration in establishing software requirements."

And:

> "Growing software organically, adding more and more function to systems as they are run, used, and tested."

He was describing the loop. He was also describing why it was hard in 1986: the tooling for rapid prototyping was primitive, the cost of iteration was high, and the organizational structures that made CI/CD possible did not yet exist. The idea was there. The infrastructure was not. The pattern repeats: first the concept, then a long wait, then the tooling.

**Agentic software engineering (2020s).** AI coding agents write the code. Other agents review it. Still others test it. The human writes the specification. The loop runs without the human in it. This is what Steinberger and Cherny and Osmani are describing. It is what StrongDM has been running in production since mid-2025. It is what Shapiro's five-level dark factory framework predicts. The loop has been progressively tightening for fifty-seven years. At the end of the tightening, the human is no longer inside the loop. The human designs the loop. The loop runs itself.

**Lehman's laws (1980).** Meir Lehman observed that large software systems must be continually adapted or they become progressively less useful. The environment changes around them. The system must change in response. This is a loop: the system observes its environment, detects drift, and adapts. Lehman formalized what the Garmisch attendees felt intuitively: software is not a product. It is a process. The process is a loop. The loop never ends. Lehman stated it as a law because it is not optional. You can ignore it only if your system is dead.

## What changed in 2026

If the loop was always the idea, what actually changed? Why did Loop Engineering become a named discipline in 2026 and not 2016 or 2006?

Three things changed at roughly the same time.

**First, the agent became capable enough to close the loop autonomously.** A loop needs a generator and a verifier. The generator produces candidates. The verifier checks them. If the generator is not good enough to produce candidates worth verifying, or the verifier is not good enough to distinguish correct from plausible, the loop does not close. You need a human inside it — providing the right prompt, inspecting the output, adjusting, trying again. Prompt Engineering was the era where the human *was* the verifier. Loop Engineering is what becomes possible when the verifier can be automated. The human moves from inside the loop to outside it, designing the verification criteria, the stopping conditions, the fallback paths. The loop runs. The human inspects only the final result, or only the failures.

This is the same transition that happened when TDD automated the test-authoring loop, and when CI automated the integration loop, and when CD automated the deployment loop. Each time, something that previously required human judgment became mechanical. Each time, the human's role shifted from *performing* the verification to *designing* the verification. The 2026 transition is different in scale, not in kind. The scope of what can be verified automatically has expanded from "does this function return the right value" to "does this pull request correctly implement the specification." That is a large expansion. But the structure — human designs the verification, machine runs it — is the same structure Perlis described in 1968.

**Second, the infrastructure matured.** Claude Code shipped `/loop`, `/goal`, and dynamic workflows. OpenClaw shipped worktree-isolated agent runners. The major platforms converged on the same primitives: time-based triggers, isolated workspaces, project memories, MCP connectors, sub-agents. Steinberger followed his viral tweet with a concrete example:

> "Tell codex to maintain your repos, wake up every 5 minutes and direct work to threads."

This is not a metaphor. It is a command. The infrastructure exists. The loop is no longer a research project. It is an engineering discipline with known primitives, known failure modes, known design patterns. Osmani identified five core components: automations, worktrees, skills, plugins/connectors, and sub-agents. Cherny's workflow — loops reading GitHub issues, scanning feedback, parsing Slack, dispatching work — is a composition of these components. The loop is an architecture. The architecture has primitives. The primitives are documented.

**Third, the economics inverted.** For most of software engineering history, the constraint was implementation. You could specify faster than you could build. The bottleneck was the rate at which intent could be translated into working code. With capable coding agents, the constraint moved from implementation to specification. You can build faster than you can specify. The expensive thing is no longer writing code. The expensive thing is writing precise enough specifications that the agent produces what you actually want.

The loop is the mechanism that discovers the gap between your specification and your intent. Each turn of the loop reveals something you forgot to say, an edge case you didn't consider, an assumption you didn't state. The loop does not just produce code. It produces *knowledge about what you actually want*. This is exactly what Kinslow was describing in 1968: "certain specifications and issues may not be realised until production of the program begins." The discovery that your specification was wrong is not a failure of the process. It is the process. The loop is the machine that converts specification errors into knowledge.

## The loop is a capital investment

There is an economic principle here that is easy to miss. A prompt is a labor expense. You pay it every time. A loop is a capital investment. You pay to build it once. It pays you back over every turn.

The distinction is the same one that separates craftsmanship from manufacturing. A craftsperson makes each item by hand. Quality depends on skill, attention, and energy. A manufacturer builds a production line. The line has a higher upfront cost. It requires design, tooling, calibration. Once built, it produces items at marginal cost approaching zero. The craftsman's advantage is flexibility. The manufacturer's advantage is scale.

Prompt Engineering is craftsmanship. Each interaction is hand-tuned. The quality of the output is a function of the quality of the prompt, which is a function of the skill and attention of the prompter. This works well at low volume. It breaks at scale. When your agent generates a hundred candidate solutions and you need to evaluate all of them, you cannot hand-craft a hundred evaluation prompts. You need a loop.

Loop Engineering is manufacturing. You invest upfront in designing the generator-verifier-refiner pipeline, the stop conditions, the fallback paths, the state management. Once built, the loop runs at marginal cost approaching the API bill. You pay the design cost once. You pay the compute cost per turn. You pay zero attention cost. Your attention is the scarcest resource. The loop conserves it.

This is why Steinberger called it a "monthly reminder." The economics keep shifting toward loops and away from prompts. Every month, the agents get better, the infrastructure gets more capable, and the case for remaining inside the loop gets weaker. The reminder is monthly because the threshold keeps moving. What required a human in the loop last month might not this month. The capital investment that made sense today would have been premature a year ago and will feel obvious a year from now.

## The loop is the system

0xCodez synthesized the emerging practice into a fourteen-step roadmap. The central thesis:

> "Self-improvement is a property of the system, not the model — build the system."

The model's weights are fixed. The model does not learn during use. The model does not remember what worked last time unless you build memory around it. The *system* accumulates. STATE files record what was tried and what succeeded. Skills capture reusable patterns. Eval loops measure whether the output is getting better or worse. Each run writes lessons to memory. The next run inherits sharper context. The model is unchanged. The framework around it gets sharper. That is the self-improvement. Not the model learning. The framework accumulating.

This was always true, even before AI. A team with a good CI pipeline and a bad codebase will, over time, improve the codebase. The pipeline provides feedback. The feedback drives improvement. A team with a good codebase and no pipeline will, over time, degrade the codebase. The absence of feedback allows drift. The codebase is not self-healing. The pipeline does the healing. The pipeline is the loop. The loop is the system. The system improves because it is a loop.

The same principle applies to agentic systems. A single agent with a powerful model but no loop produces one output and stops. The output is as good as the model can make it in one shot. A weaker model wrapped in a well-designed loop — generate, verify, refine, repeat — can outperform a stronger model used once. The loop provides the improvement the model cannot provide on its own. The model is a component. The loop is the system. The system outperforms the component.

Osmani captured the relationship between the loop and its components:

> "Loop Engineering sits one floor above the harness. The harness runs on a timer, it spawns little helpers, and it feeds itself."

The harness is the scaffolding. The loop is the design that gives the scaffolding purpose. The harness wakes up, spawns agents, collects their output. The loop decides what to do with the output: was it good enough? Should we try again with a different approach? Should we escalate to the human? The harness is mechanical. The loop is architectural. You can build a harness without understanding loops. You will just have a very expensive timer.

This is why Loop Engineering matters beyond the buzzword. It is not about whether you prompt by hand or design loops. It is about where you invest your design effort. Prompt Engineering invests in the input to the model. Loop Engineering invests in the structure around the model. The second investment compounds. The first does not. A better prompt produces a better output once. A better loop produces better outputs forever.

Dijkstra understood this in 1968. The goto statement was "just too primitive; it is too much an invitation to make a mess of one's program." He was not arguing against power. He was arguing for structure over ad-hoc power. A goto can do anything. That is the problem. A loop with well-designed stop conditions, verification gates, and state management can also do anything — but it does it reliably, repeatably, improvable. The difference between a goto and a for-loop is the difference between a prompt and a loop-engineered system. Both produce output. One produces output you can reason about. One produces output you can improve.

## The trap

There is a trap. Osmani warned about it. The loop can become a machine for accelerating decline.

> "Loops can also accelerate decline if used to avoid understanding."

If you don't understand what your loop is doing, the loop will drift. Each turn of the loop compounds small errors. The output gets worse. The verification gates, if they are weak, pass the degraded output through. The next turn starts from degraded input. The loop becomes a decay spiral. You discover the problem only when the output is visibly broken — by which point the loop has been producing subtly wrong results for weeks.

This is not a new problem. It is the same problem Graham described in 1968: "build the whole thing, push it off the cliff, let it crash, and start over again." The difference is that the loop accelerates the cycle. A waterfall project takes years to crash. A loop can degrade in hours. The loop tightens feedback in both directions — it tightens the feedback that improves the system, and it tightens the feedback that degrades it. A well-designed loop amplifies good decisions. A poorly-designed loop amplifies bad ones. The loop is amoral. It does not care what it amplifies.

The answer is not to avoid loops. The answer is to verify the verifier — to have meta-loops that check whether the primary loop is still producing useful output, to have escalation paths that flag anomalies for human review, to have circuit breakers that halt the loop when confidence drops below a threshold. The loop needs guardrails. The guardrails are themselves loops. The architecture is recursive. The recursion bottoms out in human judgment. The human does not need to be in every loop. The human needs to be at the escape hatch.

This is the hardest lesson of Loop Engineering. Removing yourself from the loop does not mean removing yourself from responsibility. It means relocating your attention to the places where it matters most: designing the verification criteria, monitoring the escape hatches, inspecting the anomalies. The loop handles the routine. You handle the exceptions. The exceptions are where the value is.

## From Garmisch to the loop

The NATO conference ended with a set of recommendations. Better tools. Better education. Better management practices. Better theoretical foundations. They did not recommend Loop Engineering because the term did not exist and the infrastructure did not exist and the agents did not exist. But the idea was there. The need for feedback was stressed many times. The Wright brothers method was recognized as unsustainable. Perlis described the loop where testing and designing interlace and the simulation becomes the system. Kinslow described the loop where production reveals what specification missed. The entire conference was an argument for the loop. The argument was correct in 1968. It took fifty-seven years for the infrastructure to catch up.

What happened in those fifty-seven years was the progressive construction of the machinery that makes the loop practical. Structured programming made loops local. Unix made loops composable. TDD made loops executable. CI/CD made loops continuous. Agentic engineering made loops autonomous. Each step removed a human from some part of the loop. Each step tightened the time between action and feedback. Each step moved the industry closer to what the Garmisch attendees described as necessary but impossible with the tools of their time.

The tools are no longer the constraint. The constraint is the willingness to design the system instead of just operating inside it. Steinberger's tweet was not a technical insight. It was a restatement of page eleven of the NATO report, fifty-seven years later, in the language of the platform that finally made it possible. Feedback is everything. Build the loop. Then get out of it.

The loop has always been the idea. The loop is finally buildable. The question now is whether we will build loops that make us better, or loops that make us faster at being wrong. That choice — better or faster, understanding or avoidance, structure or chaos — is the same choice the Garmisch attendees faced. The technology has changed. The choice has not.
