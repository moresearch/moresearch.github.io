---
title: Factorio & SC2: Systems Thinking
date: 2026-05-08
slug: how-factorio-and-sc2-shaped-my-systems-thinking
summary: Factorio taught me to reason about throughput, bottlenecks, and layout, while StarCraft II taught me tempo, prioritization, and hotkey discipline. Both transferred directly into effective terminal and Vim-based engineering work.
tags: systems, productivity, reflections
---

I do not think I learned systems thinking only from engineering.

A surprising amount of it came from games, especially **Factorio** and **StarCraft II**.

That does not mean games magically teach software architecture. What they did give me was repeated exposure to the exact kinds of pressure that matter in real engineering work: limited attention, constrained resources, competing priorities, incomplete information, and the need to build systems that keep working while I am busy somewhere else.

Over time, I started to realize that a lot of the habits that make me effective in terminal- and Vim-heavy environments were strengthened by those games long before I had language for them.

## Factorio taught me to think in flows, not parts

Factorio is one of the cleanest lessons I know in throughput thinking.

At first, the game looks like a construction game. Later, it becomes obvious that it is really a lesson in **flows**:

- ore becomes plates,
- plates become intermediate products,
- intermediates become higher-order assemblies,
- energy, belts, inserters, trains, and layout all constrain the whole pipeline.

The key mental shift is that local correctness is not enough. A single sub-factory can be beautifully designed and still fail the system if it starves upstream or overloads downstream.

That maps directly to software.

In a codebase, I care much more now about how information, control, and dependency pressure move through the system than about whether one module looks clever in isolation. Factorio trained my brain to look for:

1. bottlenecks,
2. wasted movement,
3. hidden coupling,
4. poor observability,
5. and scaling limits that only appear after expansion.

That is an engineering habit.

When I open a Go service or a shell pipeline, I often think about it the same way I think about a Factorio bus or train network: where is the actual choke point, what resource is really scarce, and which redesign improves throughput without adding chaos?

## Factorio also taught me to value layout as an operational decision

One of the biggest transfers from Factorio into coding is respect for layout.

In the game, layout is not decoration. Layout determines whether the factory is easy to expand, easy to debug, and easy to reason about under growth. A cramped but “efficient” build often becomes a trap later.

That same instinct helps in terminal and Vim environments.

I like tools that preserve spatial memory:

- stable file trees,
- stable keymaps,
- stable command patterns,
- stable pane layouts,
- stable text structure.

The reason is not aesthetic purity. It is operational speed. Good layout reduces context-switch cost.

Factorio makes that lesson painfully obvious because bad layout punishes you every time the system scales. So does a codebase.

## StarCraft II taught me prioritization under pressure

If Factorio trained system layout and throughput thinking, **StarCraft II** trained tempo and prioritization.

SC2 is not just about speed. It is about deciding what matters **right now** while the rest of the game keeps moving.

You cannot do everything at once. You have to:

- macro while scouting,
- spend money while defending,
- expand while preserving unit production,
- and avoid wasting attention on the wrong fight.

That feels very familiar to real engineering.

When I am coding inside Vim or a terminal-heavy workflow, the whole environment rewards the same skill: keep the main loop alive while handling interruptions. That means I am constantly asking:

1. what is the highest-leverage action in this moment,
2. what can be deferred safely,
3. what needs to stay on rhythm,
4. what signal is actually worth interrupting for.

SC2 trained my brain to stop romanticizing constant reaction. Not every alert deserves a response. Not every branch of work deserves equal attention. Good play is partly about refusing low-value actions. So is good engineering.

## Hotkeys changed how I think about tools

Both Factorio and SC2 reward compressing common actions into reliable motor patterns. That maps directly into why I like Vim and terminal workflows so much.

Once the tool becomes hotkey-native, the interaction stops feeling like “issue a command from scratch every time.” It becomes a vocabulary of rehearsed moves.

That has two effects:

First, it reduces friction. I do not want to re-decide how to move, search, select, split, grep, format, diff, or commit every few minutes.

Second, it preserves cognitive energy for the actual problem.

That is what good hotkey systems do. They move execution into muscle memory and free working memory for reasoning.

Vim is excellent at this when it clicks. The terminal is excellent at this too. You start thinking in composable verbs and operators rather than isolated GUI actions.

Games taught me to respect that style of interaction before I understood it formally.

## Map awareness became systems awareness

Another direct transfer from SC2 is the idea of map awareness.

Strong play depends on more than your current camera position. You need a model of what is happening elsewhere:

- your production,
- your expansions,
- likely enemy timings,
- vulnerable paths,
- information gaps.

In engineering terms, that becomes system awareness.

When I work effectively in a terminal environment, I am usually maintaining a rough mental map of:

- what processes are running,
- which files are authoritative,
- where the risky boundaries are,
- which commands are safe,
- what the current bottleneck is,
- and what state the repo is in.

That is not very different from strategy-game awareness. It is still about managing incomplete information across a live system.

## Terminal work feels natural to me for the same reason these games did

A good terminal workflow feels alive in the same way a strategy game does.

There is rhythm. There is structure. There is feedback. There are repeated loops. There is economy in movement. There is a constant tradeoff between local action and global awareness.

That is why I think the transition from those games into terminal/Vim-heavy coding felt natural to me. The surface domain changed, but the cognitive style did not.

I was still:

- building repeatable flows,
- reducing wasted motion,
- maintaining a global map,
- watching for bottlenecks,
- and turning frequent actions into low-friction habits.

## My real takeaway

Factorio and SC2 did not teach me software engineering directly. They taught me habits that made software engineering easier to learn deeply.

Factorio sharpened my instinct for pipelines, layout, scaling, and bottlenecks.

StarCraft II sharpened my instinct for tempo, triage, attention management, and hotkey discipline.

Together, they made terminal and Vim environments feel less like harsh tools and more like expressive systems. And I think that is why those environments still feel so productive to me now: they reward exactly the kind of system-level thinking those games trained over and over again.


Game design and systems thinking share a common structure: a set of rules, agents acting within those rules, and emergent behavior that no individual agent intended. The same structure appears in market design, protocol design, and software architecture. The engineer who studies games learns to see the rules behind the behavior. The behavior that looks like chaos is often the equilibrium of a system whose rules you haven't discovered yet. The discovery is the engineering.


> Games teach systems thinking because they are systems. Factorio is a supply chain. StarCraft is a resource allocation problem. The player who sees the system behind the graphics wins. The engineer who sees the system behind the code does too.
