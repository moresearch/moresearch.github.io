---
title: Why algorithmic trading is a game worth playing
date: 2026-07-12
slug: algorithmic-trading-books
summary: "Algorithmic trading is the discipline of encoding market beliefs into executable strategies and testing them against reality. It is the closest thing to pure applied epistemology in finance. Every trade is a hypothesis. Every P&L is a test result. This post explains why the game is worth playing and the nine books that teach you how."
tags: algorithmic-trading, books, quantitative-finance, market-microstructure, education
---

Algorithmic trading is not about speed. It is not about colocation, microwave towers, or the arms race for nanoseconds. That is high-frequency trading — one branch of a larger discipline. Algorithmic trading is the broader practice: encoding market beliefs into executable strategies and testing them against reality. Every trade is a hypothesis. Every P&L is a test result. The market is the laboratory. The laboratory never closes.

The game is worth playing for three reasons. First, it forces intellectual honesty. A backtest that looks good and a live strategy that loses money cannot both be right. The market's verdict is final. Second, it forces you to understand markets at a level that passive investing never requires. You must understand order flow, liquidity, market impact, adverse selection. You must understand what you're trading against and why your edge exists. Third, it is composable with everything else you know. Statistics, computer science, game theory, behavioral economics — every intellectual tool you have finds application in trading. The trading sharpens the tools. The tools sharpen the trading.

This post is about the books that teach you to play. Nine books. Each teaches a different part of the game. Each matters for a different reason. The reading order builds understanding from the ground up.

## 1. Larry Harris — Trading and Exchanges: Market Microstructure for Practitioners (2003)

Harris was the SEC's chief economist. His book is the best single-volume education in how markets actually work. 600 pages. Assumes nothing. Explains everything: order types, priority rules, trading costs, market maker obligations, dealer markets vs. auction markets, transparency, fragmentation, regulation.

The lesson: **the market is not a black box that converts orders into executions. It is a specific set of rules and participants with specific incentives. The rules determine who wins and who loses. Understand the rules before you trade.**

Harris organizes markets along two dimensions: the trading session (continuous vs. call auction) and the market structure (order-driven vs. quote-driven vs. brokered). Every real market is a hybrid. The NYSE is continuous and order-driven but runs call auctions at the open and close. The corporate bond market is quote-driven and brokered. Crypto AMMs are a form that didn't exist when Harris wrote — formula-driven markets. The taxonomy is portable. Add new forms. The principles remain.

A trader who hasn't read Harris doesn't know what a limit order is, how it interacts with other orders in the book, why the spread exists, or what happens to their order when they click submit. They are trading blind. The blindness is expensive.

## 2. Maureen O'Hara — Market Microstructure Theory (1995)

Where Harris tells you *what* happens, O'Hara tells you *why*. The book traces the evolution of microstructure theory: from inventory models (market makers set spreads to manage inventory) to information-based models (market makers set spreads to protect against informed traders) to strategic trader models (traders anticipate each other's strategies).

The key insight: **the bid-ask spread is not a transaction cost. It is an information cost.** The market maker posts a bid and an ask, offering to trade with anyone. Some of those anyones know more than the market maker. The spread compensates for the expected loss to informed traders. This is why spreads widen during volatility (more uncertainty → more adverse selection), why small-cap stocks have wider spreads (less public information → more private information), and why AMM fees are higher for volatile pairs. The explanation is the same across markets. O'Hara gives you the explanation.

## 3. Ernest P. Chan — Algorithmic Trading: Winning Strategies and Their Rationale (2013)

Chan's book is the practical starting point for someone who wants to build and test trading strategies. Where the theory books give you models, Chan gives you: here is how to backtest a mean-reversion strategy. Here is how to source data. Here is why your backtest looks great and will lose money live.

The lesson: **backtesting is hard. Most backtests are overfit. The overfit is invisible to the person who ran the backtest because they want the strategy to work. The only defense is out-of-sample testing, paper trading, and the humility to accept that a strategy that backtests perfectly will lose money in production.**

Chan's most valuable chapter catalogs backtesting pitfalls: look-ahead bias (using data not available at trade time), survivorship bias (testing on assets that survived), data-snooping bias (testing many strategies and reporting the winner). Each bias inflates backtest returns. Each is present in most amateur backtests. Each is avoidable. Avoiding them is the discipline that separates algorithmic trading from gambling dressed in Python.

The book covers mean reversion, momentum, pairs trading, and ETF arbitrage. The specific strategies are dated — most stopped working as more capital chased them. The principles of strategy development, backtesting, and risk management are permanent. Chan teaches you to fish. The fish you catch are your own.

## 4. Álvaro Cartea, Sebastian Jaimungal, José Penalva — Algorithmic and High-Frequency Trading (2015)

This is the graduate textbook. It assumes Harris and O'Hara. It applies stochastic optimal control to execution, market making, and liquidity provision. The mathematics is advanced — SDEs, Hamilton-Jacobi-Bellman equations, impulse control. The reward is a unified framework for thinking about trading as optimization under uncertainty.

The lesson: **trading is an optimization problem. The uncertainty has structure — price dynamics, order arrival dynamics, market impact. The structure can be modeled. The model can be solved. The solution is a strategy.**

The gap between the continuous-time model and discrete market reality is where the practitioner's edge lives. Cartea gives you the model. You must adapt it to the reality of gas costs, discrete blocks, and adversarial MEV searchers. The adaptation is the work.

## 5. Marco Avellaneda and Sasha Stoikov — High-Frequency Trading in a Limit Order Book (2008)

A 25-page paper, not a book. It is on this list because it is the most influential single publication in algorithmic market making. Every market-making bot in production descends from it.

The model: a market maker chooses bid and ask quotes to maximize expected utility. The state: cash, inventory, price. The control: how far from the mid-price to place quotes. The solution: skew quotes away from inventory. Long inventory → sell more aggressively. Short inventory → buy more aggressively. Spread widens with volatility and risk aversion.

The paper is implementable. The calibration is harder than the implementation. The calibration is the edge. The edge erodes as more people implement the model. The erosion is why the field keeps moving. Read the paper. Implement it. Learn what happens when your calibration is wrong. Then build something better.

## 6. Andrew Lo — Adaptive Markets: Financial Evolution at the Speed of Thought (2017)

Lo's book is not about strategies. It is about the framework within which all strategies operate. The efficient market hypothesis says prices reflect all available information. Lo says: markets are adaptive systems populated by boundedly rational agents competing for profits. The competition produces efficiency as an emergent property, not a static condition. Efficiency is approached, never reached.

The lesson: **strategies work, attract capital, compress their own edge, and stop working. The cycle is evolutionary. The strategist must continuously find new edges. The search is the career.**

Lo integrates neuroscience, evolutionary biology, and behavioral economics into a unified theory. The theory explains why arbitrages exist, persist, and disappear. The explanation is evolutionary. The evolution is driven by competition. The competition is getting smarter. The smarter competition is the subject of the next book.

## 7. Peter Bernstein — Against the Gods: The Remarkable Story of Risk (1996)

Bernstein tells the history of risk management from the Renaissance to modern finance. The mathematicians, philosophers, and gamblers who invented probability, statistics, and derivatives pricing. Not a trading book. The intellectual infrastructure that makes trading possible.

The lesson: **risk is quantifiable. The quantification of risk — modeling the future as a probability distribution rather than the will of the gods — is the defining intellectual achievement of modern capitalism. Every trading strategy is a bet on a distribution. Understanding the distribution is the edge. Understanding that the distribution is an estimate, and estimates are wrong, is the meta-edge.**

Bernstein's history ends before 2008. The crisis validated his thesis in reverse: quantification breeds overconfidence. The model is not the territory. The territory contains fat tails that the model missed. The fat tails are the subject of the next book.

## 8. Nassim Nicholas Taleb — Dynamic Hedging: Managing Vanilla and Exotic Options (1997)

Taleb was an options market maker before he was a public intellectual. *Dynamic Hedging* is a collection of lessons from actually hedging options in real markets. Delta hedging works until it doesn't. Gamma risk kills you when volatility spikes. The tail is where the money is lost. The tail is thicker than Black-Scholes assumes.

The lesson: **the market is not log-normal. Tails are fat. Hedging strategies that assume thin tails fail when the tail arrives. The tail arrives more often than the model predicts. Respect the tail.**

Out of print and expensive. Worth it. Taleb's later books — *Fooled by Randomness*, *The Black Swan* — are philosophical expansions. Read *Dynamic Hedging* for the trading. The trading book is better.

## 9. Michael Lewis — Flash Boys: A Wall Street Revolt (2014)

Lewis tells the story of Brad Katsuyama and IEX — traders who discovered the U.S. stock market was rigged for speed and built an exchange to fix it. The book is a thriller. It is also the best introduction to the reality of modern market infrastructure for someone who has never thought about it.

The lesson: **the market is not a level playing field. Speed is an advantage. The advantage is purchased through infrastructure. The infrastructure cost creates barriers. The barriers concentrate profits. The concentration is structural.**

*Flash Boys* is controversial. HFT firms argue Lewis misunderstood market making. The controversy is beside the point. The book's contribution is making visible the physical infrastructure of trading — the data centers, the fiber, the matching engines — that determines who wins. The infrastructure was invisible before Lewis. It is not invisible now.

## The reading order

**Start with Harris.** Understand the mechanism. You cannot trade what you don't understand.

**Then O'Hara.** Understand *why* the mechanism works the way it does. The theory makes the practice intelligible.

**Then Chan.** Learn the discipline of turning an idea into a backtest and a backtest into a live strategy. Make the mistakes Chan warns about. Learn from them.

**Then Cartea.** The mathematics of optimization. Not everyone needs this. The people who do know who they are.

**Then Avellaneda-Stoikov.** Short. Implementable. The bridge from theory to a working market-making bot. Build it. Watch it lose money because your calibration is wrong. Fix the calibration.

**Then Lo.** The evolutionary framework. Understand why your edge will erode and why that's normal and what to do about it.

**Then Bernstein.** The intellectual history. Understand where the tools came from. Respect the minds that built them.

**Then Taleb.** The reminder that the tools have limits. The limits are where you lose money. The reminder is uncomfortable. The discomfort is productive.

**Then Lewis.** The human story. The market is not an abstraction. It is people, infrastructure, incentives, and politics. Lewis makes it human. The humanity matters.

The bookshelf is a curriculum. The curriculum is a career. The career is a continuous search for edges that haven't yet been competed away. The search is the game. The game is worth playing.

---

**References:**
- Larry Harris, *Trading and Exchanges: Market Microstructure for Practitioners*, Oxford University Press, 2003.
- Maureen O'Hara, *Market Microstructure Theory*, Blackwell, 1995.
- Ernest P. Chan, *Algorithmic Trading: Winning Strategies and Their Rationale*, Wiley, 2013.
- Álvaro Cartea, Sebastian Jaimungal, José Penalva, *Algorithmic and High-Frequency Trading*, Cambridge University Press, 2015.
- Marco Avellaneda and Sasha Stoikov, "High-Frequency Trading in a Limit Order Book," *Quantitative Finance*, 2008.
- Andrew Lo, *Adaptive Markets: Financial Evolution at the Speed of Thought*, Princeton University Press, 2017.
- Peter Bernstein, *Against the Gods: The Remarkable Story of Risk*, Wiley, 1996.
- Nassim Nicholas Taleb, *Dynamic Hedging: Managing Vanilla and Exotic Options*, Wiley, 1997.
- Michael Lewis, *Flash Boys: A Wall Street Revolt*, W.W. Norton, 2014.
- Related posts: [DEX trading series](https://blog.hackspree.com/#dex-trading-order-book), [On Scarcity](https://blog.hackspree.com/#scarcity)
