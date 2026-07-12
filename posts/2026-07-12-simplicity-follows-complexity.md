---
title: Simplicity does not precede complexity, but follows it
date: 2026-07-12
slug: simplicity-follows-complexity
summary: "Alan Perlis wrote: 'Simplicity does not precede complexity, but follows it.' You do not start simple. You start messy and refine. The simple thing is the complex thing, understood. Git rebase, refactoring, and clean interfaces are not first drafts. They are the artifact of having already been wrong."
tags: perlis, simplicity, complexity, git, refactoring, software-design
---

Alan Perlis, the first Turing Award winner, wrote a set of epigrams on programming. Number 24 reads:

> "Simplicity does not precede complexity, but follows it."

The sentence is small. The insight is deep. You do not start with simplicity. You start with the mess. The simplicity comes after — after you have built the complex thing, understood it, and removed everything that wasn't necessary. The clean interface wasn't designed. It was discovered by building the dirty one and noticing which parts mattered.

This is not how we teach software engineering. We teach: design the clean thing first. Think before you code. Get the architecture right upfront. The teaching is aspirational. The practice is different. You cannot design the clean thing first because you don't understand the problem well enough to know what "clean" means. Clean is a property of understanding. Understanding comes from building. Building produces complexity. The complexity teaches you what to remove. The removal produces simplicity. The simplicity follows the complexity. It never preceded it.

## Git history as the proof

Your git history is the most honest record of how software is actually made. The first draft of a branch is a mess. Commits are named "WIP," "fix," "try again," "actually fix," "fix the fix." Files change in ways that make no sense in isolation. The commit that added the feature also broke the tests. The commit that fixed the tests also reformatted unrelated code. The history is a log of discovery, not a narrative. Discovery is messy. Messy produces complex histories.

Then, before merging, you run `git rebase -i`. You squash the "fix" commits into the commit they were fixing. You reorder the logical changes into a sequence that tells a story. You split the commit that did two things into two commits that each do one thing. You rewrite the commit messages to explain why, not what. The result is a clean history. Four commits. Each does one thing. Each has a clear message. The sequence makes sense. The history reads like it was planned.

It was not planned. It was discovered. The clean history is the artifact of having already been wrong. The rebase is the removal of the evidence of confusion. The evidence was real. The confusion was productive. The clean history is the simplicity that followed the complexity.

This is Perlis's epigram applied to software configuration management. You cannot produce the clean history first. You produce the messy history, learn what the story actually is, and then rewrite the history to tell that story. The rewrite is not dishonest. It is editorial. The editor removes what doesn't serve the narrative. The narrative wasn't known when the first draft was written. It was discovered by writing the draft.

## The dirty branch as research

A branch is a research project. You don't know what you'll find. You have a hypothesis: "I think I can add this feature by modifying these three files." The hypothesis is wrong. The three files become seven. One of them requires a refactor you didn't anticipate. The refactor breaks a test in an unrelated module. You fix the test. You discover that the feature doesn't work the way the spec described because the spec didn't account for an edge case you found while implementing. You adjust. The branch grows. The commit count climbs. The messages get shorter. "wip," "ugh," "ok actually working now."

This is not failure. This is research. The research produced complexity. The complexity is the evidence that you learned something. The learning is the value. The messy history is the record of the learning. It is not fit for public consumption. It is not meant to be. It is your lab notebook. The notebook is messy. The paper you publish is clean. The paper is the rebased branch. The notebook is the original history. You need both. The notebook is how you got there. The paper is what you found.

The industry norm of squashing entire branches into a single commit is the extreme form of this. One commit. One message. All the evidence of discovery, erased. The squash is too aggressive. It removes the intermediate logic — the sequence of insights that produced the final result. A future reader who encounters a bug in this code wants to see the commits that added it, not a single massive diff. The rebase preserves the logic. The squash preserves only the outcome. Outcome without logic is harder to debug. Logic without cleanup is harder to read. The art is in the middle: enough rebase to tell the story, not so much that the story disappears.

## Refactoring as the same pattern

Refactoring is the Perlis pattern applied to code structure. You do not design the clean abstraction first. You design the abstraction that works. It is messy. Methods are too long. Classes have too many responsibilities. The interface exposes implementation details. The code works. The code is ugly.

You refactor. You extract methods. You split classes. You hide implementation behind the interface. The result is clean. The clean result looks like it was designed. It was not designed. It was discovered by writing the ugly version, understanding which parts were ugly, and removing the ugly. The ugly taught you what clean meant. Clean meant "what remains after you remove the accidental complexity." You couldn't identify the accidental complexity until you built the accidental complexity. Building it was the research. Removing it was the refactoring. The refactored code is the simplicity that followed the complexity.

> "There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult." — Tony Hoare

The first method is more difficult because it requires you to see the simplicity before you've built the complexity. Nobody can do this reliably. The people who appear to do it are people who have built similar complexity before and are remembering what they learned. They're not starting from simplicity. They're starting from the simplicity that followed the complexity they built five years ago on a different project. Experience is the accumulated simplicity that followed accumulated complexity. The senior engineer who designs the clean thing first is not designing. They are remembering. The memory is of complexity they once built and later simplified. The simplification became intuition. The intuition looks like foresight. It is hindsight, internalized.

## The false simplicity of premature design

The danger of Perlis's epigram is the danger of misunderstanding it. It does not mean "don't try to be simple." It means "don't expect to be simple on the first try." The first try will be complex because the first try is where you learn what the problem actually is. The learning requires complexity. The complexity is the tuition. You pay it. Then you simplify. The simplification is the return on the tuition.

Premature simplicity is the attempt to skip the tuition. You design the clean architecture before you've written any code that exercises it. The architecture is clean. It is also wrong. It is wrong because it was designed against an understanding of the problem that the designer didn't have yet. The designer thought they understood. They understood the spec. The spec is not the problem. The problem is what you discover when you try to implement the spec and find that the spec didn't account for the database migration, the cache invalidation, the race condition, the legacy API that returns XML, the user who needs the opposite of what the spec describes. The clean architecture didn't account for any of this because the designer hadn't encountered it. The designer hadn't encountered it because they hadn't built the complex thing. They tried to start with simplicity. Simplicity does not precede complexity.

The teams that ship ugly code that works and then refactor are following Perlis. The teams that design beautiful architectures that never ship are violating him. The first group pays the tuition. The second group avoids the tuition and never graduates.

## The rebase as editorial craft

The rebase is the moment when the research becomes the story. The researcher becomes the editor. The editor's job is to remove everything that doesn't serve the reader. The reader is the future maintainer — possibly you, six months from now, at 2am, trying to understand why this code does what it does.

A good rebase does several things. It groups related changes into cohesive commits. It orders commits so that each one is a logical step that builds on the previous. It writes messages that explain why the change was made, not what the diff contains. It removes false starts, dead ends, and debugging code that served the researcher but would confuse the reader. The result is a history that tells a story. The story is true. It is not the whole truth. The whole truth includes the four hours you spent chasing a bug that turned out to be a typo. That truth is not useful to the reader. The editor removes it.

The rebase is not lying. It is editing. The difference between editing and lying is whether the published version misleads. A history that hides a security vulnerability is lying. A history that squashes the commit where you tried three different approaches and none worked is editing. The approaches that didn't work taught you something. The something is in the final commit message. The approaches themselves are not. They served their purpose. Their purpose was to teach you. You learned. The reader needs the lesson, not the curriculum.

## The art of the squash

When to squash, when to preserve, when to split — these are editorial judgments. They require taste. The taste is developed by reading other people's histories and noticing which ones helped you understand and which ones didn't.

Squash when the intermediate commits are noise. "wip," "fix typo," "try again" — these add no information to the reader. Squash them into the commit they were iterating toward. The iteration was real. The reader doesn't need to see it.

Preserve when the intermediate commits are logical steps. "Extract the payment interface" followed by "Implement the Stripe adapter" followed by "Add payment confirmation email" — these are three distinct decisions. Each can be understood in isolation. Each might need to be reverted independently. Preserve them as separate commits.

Split when a single commit does two unrelated things. "Add the payment flow and also reformat the entire codebase" — these should be two commits. The reformatting is noise in the payment commit. The payment logic is noise in the reformat commit. Split them. The reader of the payment commit wants to understand payment. The reader of the reformat commit wants to understand reformatting. Neither wants both.

These judgments are Perlis applied to history. The clean history is the simplicity that followed the complexity of the original branch. The original branch was the research. The rebase is the simplification. The simplified history is what you merge. It didn't precede the messy one. It followed it.

## The general principle

Perlis's epigram generalizes beyond git. It applies to every creative act that produces a structured artifact. The first draft of a novel is messy. The published version is clean. The clean version didn't precede the messy one. The messy one taught the author what the novel was about. The author rewrote it to be about that. The rewrite was the simplicity that followed the complexity.

The first version of an API is messy. It exposes implementation details, has inconsistent naming, handles errors differently in different endpoints. The second version is clean. The second version didn't precede the first. The first taught the designer which parts callers actually used, which parts were confusing, which errors were common. The second version removed what wasn't needed, standardized what was, and hid what shouldn't have been exposed. The clean API is the artifact of having watched real users struggle with the messy one.

The first architecture of a system is messy. The monolith has responsibilities that should be separate, dependencies that should be inverted, data that should be owned. The second architecture — the microservices extraction, the refactored modules, the clean interfaces — didn't precede the first. The first taught the team where the boundaries actually were. The boundaries weren't visible until the code was written. The code was complex. The complexity revealed the boundaries. The boundaries enabled the simplicity.

> "Fools ignore complexity. Pragmatists suffer it. Some can avoid it. Geniuses remove it." — Alan Perlis

The genius doesn't start with simplicity. The genius builds the complex thing, understands it completely, and then removes the complexity that was never necessary. The removal is the genius. The removal is visible. The complex thing that preceded it is not. The genius looks like someone who started simple. They didn't. They started complex and removed everything that didn't earn its place. The removal is the art. The art is invisible. The simplicity is what remains.

---

**References:**
- Alan Perlis, "Epigrams on Programming," *SIGPLAN Notices*, Vol. 17, No. 9, September 1982.
- Related posts: [Git is a Unix tool](https://blog.hackspree.com/#git-unix-philosophy), [The Unix philosophy](https://blog.hackspree.com/#unix-philosophy), [Brooks on Software Design](https://blog.hackspree.com/#brooks-design-conceptual-integrity), [Engineering is art and philosophy, grounded in economic law](https://blog.hackspree.com/#engineering-is-economics)
