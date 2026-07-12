---
title: Lending Protocols
date: 2026-07-12
slug: fi-lending
summary: "Lending is the oldest financial instrument. The Code of Hammurabi regulated interest rates in 1754 BC. Compound and Aave automated lending on-chain with overcollateralized pools and algorithmically set interest rates. The functional origin is 3,800 years old. The code is seven years old. The principle is the same."
tags: defi, lending, compound, aave, credit
series: crypto-fi-instruments
part: 2
---

Lending is the oldest financial instrument. Before there were equities, before there were bonds, before there were derivatives, there were loans. A farmer borrows seed grain after a bad harvest. A merchant borrows capital for a voyage. A king borrows to finance a war. The lender provides resources now in exchange for repayment with interest later. The interest compensates for the time value of money and the risk of default.

The Code of Hammurabi, inscribed in Babylon around 1754 BC, regulated lending. It set maximum interest rates — 33.3% for grain loans, 20% for silver loans. It required loans to be witnessed. It specified penalties for fraudulent lending practices. The code is the oldest surviving legal text. Its largest single subject is lending. Lending is that old. Lending is that central to civilization.

## The functional origin: the moneylender

In medieval Europe, lending was constrained by usury laws — the Catholic Church prohibited charging interest on loans. The prohibition created a market niche filled by Jews, who were not subject to canon law, and by Lombards, Italian merchants who developed bills of exchange that disguised interest as exchange-rate differentials. The moneylender lent at interest and was vilified for it. The vilification was moral. The function was economic. The economy needed credit. The moneylender provided it. The moral opposition was overcome by economic necessity.

The Medici family, in 15th-century Florence, systematized lending into banking. They took deposits, made loans, transferred funds across Europe through bills of exchange, and financed trade, government, and the arts. The Medici bank was the largest financial institution in Europe. Its innovation was scale: lending as an institution rather than an individual activity. Deposits funded loans. Loans generated interest. Interest funded expansion. The virtuous cycle was the model for modern banking.

The model had a structural vulnerability: loans were not fully collateralized. The Medici bank lent to kings and princes who could not be compelled to repay. The bank failed in 1494 when its largest borrowers defaulted. The failure was the consequence of unsecured lending to sovereigns. The lesson: collateral matters. The lesson was learned by DeFi lending protocols 500 years later.

## DeFi lending: Compound and Aave

Compound, launched in 2018, introduced the pool-based lending model to DeFi. Lenders deposit tokens into a pool. Borrowers borrow from the pool. Interest rates are set algorithmically: as utilization (the percentage of the pool that is borrowed) increases, the interest rate increases. The rate curve incentivizes equilibrium — when borrowing demand is high, rates rise, attracting more deposits and discouraging borrowing. When demand is low, rates fall. The algorithm is the market maker for credit.

All loans are overcollateralized. A borrower must deposit more value than they borrow. If the collateral value falls below the liquidation threshold, the position is liquidated: anyone can repay the loan and claim the collateral at a discount. The liquidation incentive ensures that positions are closed before they become undercollateralized. The overcollateralization eliminates credit risk. The trade-off: capital efficiency. Overcollateralized lending cannot expand the credit supply. It can only recycle existing capital. The limitation is the subject of undercollateralized lending protocols, which are still experimental.

Aave, launched in 2020, extended Compound's model with flash loans (borrow and repay in one transaction, zero collateral), rate switching (stable vs. variable), and multi-asset pools. Aave is now the largest DeFi lending protocol, with billions in total value locked.

## The interest rate model

The core innovation of DeFi lending is the algorithmic interest rate. Traditional lending uses credit scores, relationship banking, and manual underwriting. DeFi lending uses a utilization curve: borrow rate = base rate + (utilization rate × multiplier), with a "kink" at optimal utilization where the slope steepens. The curve is transparent. The curve is enforced by code. The curve eliminates the need for credit assessment. The collateral substitutes for the credit score.

The utilization curve is a market mechanism. When utilization is low, rates are low — capital is abundant, borrowers are scarce. When utilization is high, rates rise sharply — capital is scarce, borrowers are abundant. The sharp rise above the kink prevents the pool from being fully utilized, which would prevent depositors from withdrawing. The kink is the safety valve. The safety valve is algorithmic. The algorithm is the lender of last resort.

## The reference

Sidney Homer and Richard Sylla, *A History of Interest Rates* (1963, 4th edition 2005). The definitive history of lending from ancient Mesopotamia to the modern era. The book documents 5,000 years of interest rates across civilizations, tracing the evolution of credit from temple loans in Babylon to the Eurodollar market. The data shows that interest rates reflect the intersection of time preference, risk, and institutional structure. DeFi lending changes the institutional structure. Time preference and risk remain. The book is the context. The context is essential.

---

**References:**
- Sidney Homer and Richard Sylla, *A History of Interest Rates*, Wiley, 4th edition, 2005.
- Robert Leshner and Geoffrey Hayes, "Compound: The Money Market Protocol," 2019.
- Aave, "Aave Protocol Whitepaper," 2020.
- Related posts: [Flash Loans](https://blog.hackspree.com/#dex-trading-flash-loans), [Stablecoins](https://blog.hackspree.com/#fi-stablecoins)
