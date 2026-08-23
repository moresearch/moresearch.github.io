---
title: You Don't Learn What You Delegate
date: 2026-08-03
slug: ai-impacts-skill-formation
summary: "A pre-registered randomized trial (n=52) of Python programmers learning a new async library found AI assistance cut quiz scores 17% (d=0.74, p=0.01) while saving no significant time. Errors are the curriculum; pasted code is not. Six AI-use personas split the group: delegators scored 24–39%, conceptual inquirers 65–86%. Summary infographic below."
tags: ai, llm, skill-formation, learning, software-engineering, randomized-controlled-trial, pre-registration, cognitive-offloading, overreliance, trio, anthropic, empirical-software-engineering, human-ai-collaboration
---

A rare thing: a [pre-registered randomized controlled trial](https://osf.io/w49e7) of how AI assistance changes what workers learn. Shen and Tamkin (Anthropic), *[How AI Impacts Skill Formation](https://ar5iv.labs.arxiv.org/html/2601.20245v2)*. Fifty-two professional and freelance Python programmers were randomized, asked to learn a library they had never used (Trio, an async I/O library), complete two coding tasks, and then take a quiz on what they actually learned. One group had a GPT-4o chat assistant that could write the entire correct solution; the other had no AI at all.

![Summary infographic — How AI Impacts Skill Formation (via LinkedIn, August 2026)](https://blog.hackspree.com/images/skill-formation-linkedin-summary.jpg)

The headline, in two numbers: **the AI group scored 17% lower on the quiz — about two grade points (Cohen's d = 0.738, p = 0.010) — and gained no significant time in return.** The productivity miracle measured in prior work (55.5% faster with Copilot; 26.8% more pull requests) did not show up when the task required learning. That is the paper's contribution, and the thesis of this post: **AI assistance doesn't just complete the task — it removes the part of the task that teaches. You don't learn what you delegate, and the trial shows exactly which behaviors turn delegation into learning and which into a 17% toll.**

## The experiment

Participants made this a hard test of "AI helps novices most": more than a year of Python, weekly use, prior AI experience, never Trio ([Fig 17](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x15.png)). A warm-up task, then 35 minutes for two Trio tasks: a concurrent timer and a record-retrieval function with error handling. Then a 14-question, 27-point quiz over 7 Trio concepts in three skill types — conceptual understanding, code reading, debugging ([Fig 20](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x18.png)); code-writing was excluded, since syntax is what AI fixes cheapest.

![thumb:Fig 3 - Experiment interface with AI assistant panel](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Task-AI.png)
![thumb:Fig 4 - Learning task and comprehension check flow](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/task_flow.png)

The treatment assistant was no toy: GPT-4o with access to the participant's code, able to produce the full correct solution to both tasks. Four pilot studies tuned the design — the first platform had 35% non-compliance (controls using AI anyway); a later pilot ([Fig 5](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x3.png)) showed the effect so strongly (d=1.7) that the pre-registration assumed a conservative d=0.85 for power. Integrity controls included pledges and screen recordings of every participant ([Figs 21–24](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Pledge-NoAI.png); [Figs 25–27](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Task-No-AI.png) show both platforms).

![thumb:Fig 17 - Participant distribution](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x15.png)

## Results: the 17% toll

The main results ([Fig 6](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x4.png)): a 4.15-point gap on 27 points — 17%, two grade points — with no significant difference in completion time; the effect survives controlling for warm-up speed (d=0.725, p=0.016). [Fig 7](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x7.png) shows the control group scoring higher at every experience level, and [Fig 8](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x6.png) shows where the damage concentrates: the largest gap is in **debugging**, the smallest in code reading. Both groups read code; only the control group had to fix it.

![thumb:Fig 6 - Main study: task time and quiz score (n=52)](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x4.png)
![thumb:Fig 8 - Score breakdown by question type](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x6.png)

That debugging gap is the paper's most important sentence in disguise. The industry case for human oversight of AI-written code rests on humans being able to debug what AI produces — and this trial suggests the mechanism that builds that skill is exactly what AI assistance removes.

## Why there was no speedup

The qualitative analysis explains the missing speedup: heterogeneity. Watching every screen recording, the researchers found six AI-interaction personas ([Fig 11](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x9.png)), with completion times and quiz scores that barely overlap:

- **AI Delegation** (n=4) — handed the whole task to the AI, finished fastest (19.5 vs 23 min), scored 24–39%.
- **Progressive AI Reliance** (n=4) — started with a question, then delegated everything; lost the second task's concepts entirely.
- **Iterative AI Debugging** (n=4) — used AI to fix code without understanding it; slowest of the low scorers.
- **Generation-Then-Comprehension** (n=2) — generated code, then asked follow-ups to understand it. 65–86%.
- **Hybrid Code-Explanation** (n=3) — asked for code *with* explanations, and read them.
- **Conceptual Inquiry** (n=7) — asked only conceptual questions, hit errors, resolved them independently. Second-fastest overall.

![thumb:Fig 11 - The 6 AI interaction personas with outcomes](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x9.png)

The gap between worst and best usage is the real story: 24–39% vs 65–86% on the same quiz with the same assistant. The difference is not whether AI was used; it is what happened after the AI answered. The time side confirms it: some participants spent up to 11 minutes composing queries and asked 15 questions ([Fig 12](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x10.png)); debugging-heavy query mixes correlated with slower times and lower scores ([Figs 18–19](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x16.png)). And the paste-vs-type analysis ([Fig 13](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x11.png)) has a twist: pasting AI output was fastest, manually copying it was as slow as control — but the two scored the same on the quiz. *Spending time typing doesn't build understanding.*

![thumb:Fig 13 - Pasting vs manual copying of AI output](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x11.png)

## Errors are the curriculum

The mechanism the paper lands on is the most old-school finding in it: **errors teach.** The median control participant hit three errors; the median AI participant hit one ([Fig 15](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x13.png)). Of the twelve error-free completers, only two were control. The errors that matter are the Trio-specific ones — `RuntimeWarning` (a coroutine never awaited), `TypeError` (a coroutine passed where an async function was expected) — because they force exactly the conceptual knowledge the quiz tests ([Fig 14](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x12.png)). The control group did not learn *despite* the errors; the errors were the lesson plan. Active coding time tells the same story from the other side ([Fig 16](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x14.png)): AI shifted time from writing to reading AI output. The AI group's own feedback confirms it — they felt "lazy" with "gaps in (their) understanding."

![thumb:Fig 15 - Errors by condition: control meets the concepts](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x13.png)

## What this means

Three implications, stated as plainly as the paper allows.

**First, the chat interface is the *best* case.** A chat assistant at least forces the user to compose a query — some spent six minutes on a single one, and that thinking correlates with learning. An agentic tool that writes, runs, and fixes code itself removes even that. If this trial is the lower bound on cognitive offloading, agentic settings are below the floor.

**Second, the supervision argument inverts.** "AI writes code, humans verify" assumes humans have debugging skill to spend. This trial suggests that skill is built by the exact experience AI removes. Companies adopting AI-assisted onboarding for juniors are not just changing throughput; they may be choosing which generation holds the verification skill.

**Third, the dark-factory argument applies to skills, not just complexity.** When automation removed factory labor, complexity moved into supervision. When AI removes implementation, the question is where the *learning* goes — and this trial's answer is: out of the worker, unless the worker keeps engaging. Delegation is a productivity strategy and a learning strategy simultaneously; you cannot have both.

The six personas are the practical deliverable: the difference between treating the model as a chauffeur and as a tutor.

> The productivity view asks: how much did AI finish? The skill view asks: what did you keep? This trial shows the two questions now have different answers. The 17% is what you pay when you confuse them.

## All figures

Every figure from the paper, thumbed. Substantive results first; study artifacts last.

![thumb:Fig 1 - Overview of results: skills down, time flat](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x1.png)
![thumb:Fig 2 - Motivation: novices and AI in the workplace](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x2.png)
![thumb:Fig 5 - Pilot Study D: time and quiz](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x3.png)
![thumb:Fig 7 - Task time and quiz by years of experience](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x5.png)
![thumb:Fig 9 - Self-reported enjoyment and learning](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x7.png)
![thumb:Fig 10 - Self-reported task difficulty](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x8.png)
![thumb:Fig 12 - AI interaction time and query count](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x10.png)
![thumb:Fig 14 - All errors by type](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x12.png)
![thumb:Fig 16 - Active coding time vs quiz score](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x14.png)
![thumb:Fig 18 - Queries vs completion time](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x16.png)
![thumb:Fig 19 - Queries vs quiz score](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x17.png)
![thumb:Fig 20 - Example evaluation question types](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/x18.png)
![thumb:Fig 21 - Control pledge: no AI](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Pledge-NoAI.png)
![thumb:Fig 22 - Treatment pledge](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Instructions-AI-A.png)
![thumb:Fig 23 - Control instructions](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-No-AI-Instructions-B.png)
![thumb:Fig 24 - Treatment instructions](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Instructions-AI-B.png)
![thumb:Fig 25 - Task platform, control condition](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Task-No-AI.png)
![thumb:Fig 26 - Task platform, AI condition](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Task-AI.png)
![thumb:Fig 27 - Interacting with the AI assistant](https://ar5iv.labs.arxiv.org/html/2601.20245/assets/img/SC-Task-AI-Assistant.png)

---

**References:**

- Judy Hanwen Shen, Alex Tamkin. [How AI Impacts Skill Formation](https://ar5iv.labs.arxiv.org/html/2601.20245v2). arXiv:2601.20245v2, 2026. — Pre-registered RCT of AI assistance on learning the Trio async library; quiz deficit d=0.738, p=0.010; the six interaction personas.
- [Study pre-registration](https://osf.io/w49e7) and [annotated transcripts](https://github.com/safety-research/how-ai-impacts-skill-formation).
- [Summary infographic of the study (LinkedIn, August 2026)](https://media.licdn.com/dms/image/v2/D4D22AQEFQ9DawOphRA/feedshare-shrink_800/B4DZ_G1e_DGkAc-/0/1785747337417?e=1787184000&v=beta&t=OWm6l_TOpYkVO0wMJvrlSvBofiGvG-6UW1jrzvVrY24) — embedded at the top of this post.
- S. Peng et al. [The Impact of AI on Developer Productivity: Evidence from GitHub Copilot](https://arxiv.org/abs/2302.06590), 2023. — 55.5% faster task completion.
- Z. Cui et al. [The Effects of Generative AI on High Skilled Work: Evidence from Three Field Experiments with Software Developers](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4945566), 2024. — 26.8% productivity boost.
- T. Wu et al. [The Value of AI Assistance? Evidence from the Performance and Learning of Knowledge Workers](https://www.nber.org/papers/w33041), 2024. — performance gains did not persist after AI was removed.
- Related: [The Dark Factory Doesn't Eliminate Complexity — It Moves It](https://blog.hackspree.com/#dark-factory-complexity) — AI removes the work; the question is where the skill goes.
- Related: [Empirical Software Engineering: What the Studies Actually Say](https://blog.hackspree.com/#empirical-se-what-studies-say) — why RCTs in software are rare, and how to read them.
- Related: [Empirical Software Engineering: How to Think Empirically](https://blog.hackspree.com/#empirical-se-how-to-think) — pre-registration, effect sizes, and the discipline behind this trial.
- Related: [LLMs Can't Jump](https://blog.hackspree.com/#llms-cant-jump) — what models cannot do is exactly what the control group had to learn.
- Related: [Always-On Agents: State, Memory, and the Governance Gap](https://blog.hackspree.com/#always-on-agents) — agentic AI removes the query-composition step this paper found valuable.
