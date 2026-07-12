---
title: Five Habits
date: 2026-07-12
slug: scarcity-in-practice
summary: "How to apply scarcity thinking, game theory, and mechanism design to daily software engineering decisions. Five frameworks, five questions, five habits."
tags: scarcity, game-theory, mechanism-design, practice, software-engineering
series: scarcity
part: 6
---

The theory is useful. The practice is harder. Here are five frameworks for applying scarcity thinking, game theory, and mechanism design to daily software engineering decisions.
The difference between knowing a principle and applying it is the difference between owning a cookbook and cooking dinner. Most engineers who read Robbins will agree that scarcity matters. Most will return to work and debate architectures without naming the scarce resource. The knowing is easy. The doing is hard. The gap between knowing and doing is a habit. Habits are built by repetition, not by insight. The five habits below are designed to be repeated until they're automatic. The goal is not to understand scarcity. The goal is to feel its absence like a missing step on a staircase.


Alan Perlis, the first Turing Award winner, wrote: "Fools ignore complexity. Pragmatists suffer it. Some can avoid it. Geniuses remove it." Most software teams are pragmatists. They suffer complexity because avoiding it requires upfront investment and removing it requires genius. The five frameworks that follow are for pragmatists who want to suffer less. Genius is not required. Discipline is.

## 1. Name the scarce resource

Before every architectural decision, ask: what is the scarce resource? Is it developer time? Optimize for simplicity of change. Is it compute? Optimize for efficiency. Is it attention? Optimize for clarity. Is it coordination capacity? Optimize for independent deployability. The scarce resource determines the optimal trade.

Most architectural debates skip this step. Two engineers debate two architectures. Neither names the constraint. The debate is unresolvable because the constraint is unstated. State the constraint. "We are optimizing for developer time, not compute." "We are optimizing for change velocity, not operational simplicity." The debate resolves. The constraint determines the answer. Name the constraint.

## 2. Identify the game

Every situation involving multiple players is a game. Ask: who are the players? What are their strategies? What are their payoffs? What do they know that others don't? Is the game cooperative or non-cooperative? Zero-sum or non-zero-sum? Simultaneous or sequential? One-shot or repeated?

Naming the game changes how you think about it. "This is a stag hunt — we all benefit if we coordinate, but if anyone defects we all lose" is different from "this is chicken — someone needs to swerve or we crash." The game type suggests the solution. Stag hunts need coordination mechanisms. Chicken needs a pre-committed rule about who yields. Prisoner's Dilemmas need repeated interaction and reputation. Battle of the sexes needs a selection mechanism. Name the game. The solution follows.

## 3. Design the mechanism

If you don't like the equilibrium, change the game. Mechanism design works backward from desired outcomes to the rules that produce them. Ask: what outcome do I want? What rules would produce it? What information do players need? What incentives do they face? How do I make defection visible and costly?

Automated contract testing is mechanism design. "No Friday deploys" is mechanism design. A booking calendar for staging is mechanism design. A test budget per service is mechanism design. Each mechanism changes the payoffs. Changed payoffs change behavior. Design the mechanism. Don't plead for the behavior.

## 4. Account for complexity

Complexity is a cost. It consumes attention, time, and future change capacity. Every feature adds complexity. The complexity has a present cost and a future cost — the accumulated drag on every subsequent change. The future cost is invisible in the current sprint. It is visible in year three.

Account for it. When estimating a feature, include the complexity cost. "This feature will take two weeks to build and will add complexity equivalent to 5% of the current system, which will cost approximately one week per quarter in reduced velocity. The NPV at our discount rate is negative. Don't build it, or build it simpler." Most teams don't do this math. The math exists. Do it.

## 5. Use the Robbins test

For any decision, ask Robbins's four questions: What is the end? What are the means? Are they scarce? Do they have alternative uses? If yes — and it always is — the decision is economic. Name the opportunity cost. "Building this feature means not building that one." Name the trade. Make the choice explicit. Explicit choices are better than implicit ones. Implicit choices are still choices. They're just choices made without awareness that a choice was being made.

## Five habits

**Habit 1: Before debating architecture, name the scarce resource.** If you can't name it, you don't understand the trade.

**Habit 2: Before negotiating with another team, model their scarcity.** What are they optimizing for? What are they giving up? If you don't know, ask. Their answer will explain their behavior more than any personality trait.

**Habit 3: Before adding a feature, calculate the complexity budget.** How much does this cost in future velocity? Is the return worth the cost? Most features pass the "is it useful?" test. Few pass the "is it worth the complexity?" test.

**Habit 4: Before accepting a broken process, ask what game it's producing.** The process is a mechanism. The mechanism produces an equilibrium. If the equilibrium is bad, the mechanism is wrong. Change the mechanism. Don't blame the players.

**Habit 5: Before choosing a technology, ask what game it enables.** REST produces spatial coupling. NATS produces spatial decoupling. The technology is not neutral. It embeds assumptions about how players will interact. Choose the technology that produces the game you want to play.

## The meta-habit

Scarcity thinking becomes a habit. You start seeing opportunity costs everywhere. You start modeling other teams' incentives before meeting them. You start recognizing game structures in organizational conflicts. You start designing mechanisms instead of pleading for behavior. The habit is the point. The tools — economics, game theory, mechanism design — are lenses. The lenses change what you see. What you see changes what you do.

---

**This is part 6 of a 7-part series on scarcity and software.**
- [Part 1: On Scarcity](https://blog.hackspree.com/#scarcity)
- [Part 2: On Games](https://blog.hackspree.com/#scarcity-and-games)
- [Part 3: On Software Engineering Economics](https://blog.hackspree.com/#scarcity-and-software-economics)
- [Part 4: On Games in Software](https://blog.hackspree.com/#scarcity-and-software-games)
- [Part 5: On AI and Mechanism Design](https://blog.hackspree.com/#scarcity-and-mechanism-design)
- [Part 7: The Catalog of Games](https://blog.hackspree.com/#catalog-of-scarcity-games)

**References:**
- Lionel Robbins, *An Essay on the Nature and Significance of Economic Science*, Macmillan, 1932.
- Barry W. Boehm, *Software Engineering Economics*, Prentice-Hall, 1981.
