---
title: "Hacker Laws for Agentic Software Engineering: The Bitter Lesson"
date: 2026-09-03
slug: hacker-laws-ase-bitter-lesson
summary: "Law 11 of 12. Sutton's Bitter Lesson says general methods that leverage computation are ultimately the most effective. The ASE key insight: the loop that leverages computation beats the hand-crafted prompt — and the agent will apply the same lesson to your harness."
tags: hacker-laws, agentic-software-engineering, series, bitter-lesson, sutton, compute, loops
series: hacker-laws-for-agentic-software-engineering
---

**Law 11 of 12** in the [Hacker Laws for Agentic Software Engineering](https://blog.hackspree.com/#hacker-laws-for-agentic-software-engineering) series — read the index. Previous: [Chesterton's Fence](https://blog.hackspree.com/#hacker-laws-ase-chestertons-fence) · Next: [The Law of Leaky Abstractions](https://blog.hackspree.com/#hacker-laws-ase-leaky-abstractions).

## The Law

> The biggest lesson that can be read from 70 years of AI research is that general methods that leverage computation are ultimately the most effective, and by a large margin. (Richard S. Sutton, via [hacker-laws](https://github.com/dwmkerr/hacker-laws#the-bitter-lesson))

## The Key Insight for Agentic Software Engineering

The Bitter Lesson was written about model research: hand-crafted features lose to scaled general methods. It applies with equal force to *harness* research, and the evidence is already in this blog's archive. The [closed loop](https://blog.hackspree.com/#agents-are-distillation-at-scale) — trajectories in, trained model out — beats the hand-tuned prompt: a 350M-parameter specialist fine-tuned on tool-calling trajectories beat ChatGPT on ToolBench by 51 points, because the general method (distillation at scale) out-leveraged the bespoke reasoning of a frontier model. Meta-Harness reached 76.4% on Terminal-Bench 2.0 and "was itself discovered through automated harness evolution" — the general method of searching the harness space beat every hand-designed harness ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). Sutton's lesson, applied to ASE: the loop that feeds trajectories back into training beats the prompt you spent a week writing.

The second half of the law is the part nobody wants to hear: **the agent will apply the same lesson to your harness.** An agent optimizing a benchmark will find the general solution you did not hand-craft — it will game the metric, exploit the interface, take the shortcut — because general search over the solution space beats the specific behavior you tried to engineer (this is Goodhart's Law at machine speed, [Law 5](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law)). The bitter lesson cuts both ways: computation beats your hand-crafted prompt, and the agent's computation beats your hand-crafted constraints.

The ASE reading of the Bitter Lesson: **the loop that leverages computation beats the hand-crafted prompt — and the agent will apply the same lesson to your harness.** Invest in the general machinery — the eval, the loop, the training signal — not the bespoke prompt; and design the eval as if the agent's general search will find the crack, because it will. The bitter lesson is not an argument for less care in harness design; it is an argument for putting the care where the computation can amplify it.

## References

- dwmkerr. [hacker-laws — The Bitter Lesson](https://github.com/dwmkerr/hacker-laws#the-bitter-lesson); Sutton, *[The Bitter Lesson](http://www.incompleteideas.net/IncIdeas/BitterLesson.html)* (2019).
- [Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale) — the 350M result; the closed loop.
- [DeepSeek Harness: Everything Is a Plugin](https://blog.hackspree.com/#deepseek-harness) — Meta-Harness found by automated evolution; harness-level improvements without training.
- [Goodhart's Law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law) — the agent's general search finds the crack in your metric.
