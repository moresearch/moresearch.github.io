---
title: "RBF: The Collateral Is Your Code"
date: 2026-08-03
slug: revenue-based-financing-ai
summary: "The collateral for a revenue-based financing loan is your telemetry. Lenders underwrite the revenue your software produces, so the terms are engineering outcomes: margin is your architecture, revenue quality is your data model, the repayment schedule is your usage curve. Build revenue legibility like a system and the company gets financed like a system."
tags: revenue-based-financing, rbf, startups, saas, ai, financing, software-engineering, gross-margin, usage-based-pricing, data-quality, telemetry, underwriting, non-dilutive, engineering-economics
---

The collateral for a revenue-based financing (RBF) loan is not an asset. It is your telemetry — the data your software produces about the revenue it generates. The lender does not underwrite your company; it underwrites a revenue stream, read through your APIs, your metering, your books. Everything the lender can know, your code told them. Everything it cannot trust, your code hid. That makes RBF the financing instrument most legible to software engineers, and the one whose terms are engineering outcomes: gross margin is your architecture, revenue quality is your data model, the repayment schedule is your usage curve. The cost of capital is a performance metric, set by the codebase.

The thesis: **RBF is the financing instrument whose terms are engineering outcomes — margin, usage, data quality. In the age of AI, the company that builds its revenue legibility like a system gets financed like a system — and the code you write sets the price of your money.**

## The mechanics, for engineers

The terms are a protocol. A lender advances capital sized against your recurring revenue — commonly several months of it. You repay a fixed percentage of monthly revenue, typically 5–8% (lenders range 2–10%), until cumulative repayments reach the advance plus a multiple — typically 1.5–2.5x — or the term limit passes, usually three to five years, whichever comes first; in most structures the balance is forgiven at the limit, so the cap is a ceiling, not a promise. No equity, no board seat, no fixed payment that can sink a bad month; strong months pay more, weak months less. Lenders may still file a lien or take a personal guarantee at the small end — the telemetry is the underwriting collateral, not the only legal one.

| | Equity | Venture debt | RBF |
|---|---|---|---|
| Cost | dilution + board seat | interest + warrants | % of revenue, up to a cap |
| Underwriting | narrative + forecast | sponsor + balance sheet | live revenue data |
| Collateral | none | company assets | your telemetry |
| Speed | months | weeks | weeks |
| Governance | board, control | covenants | minimal |
| What the builder controls | roadmap + burn | unit economics | margin, usage, data quality |

Reported global volume surpassed $9.8 billion in 2025, with more than a hundred active lenders — Lighter Capital, Capchase, ClearCo, Pipe, Wayflyer, Founderpath, Recur Club — nearly all underwriting through the same stack a modern product already runs: bank feeds, payment processors, accounting APIs.

## The arithmetic nobody does

The cap is the price, and duration sets the rate. A $1M advance at a 1.8x cap costs $800,000: roughly 40% a year over 24 months, half that over 48. Sellers quote the percentage of revenue, never the annualized cost, because the rate depends on growth — the one number the lender cannot know and the roadmap determines. And the arithmetic cuts against fast growth: repayment is a share of revenue, so the cap is hit fastest when revenue grows fastest. Concretely, $100k MRR, 6% repayment, 1.8x cap: at 10% monthly growth the cap lands in about three years (~25% annualized); at 3% growth it takes more than six (~half that, and the term limit may forgive the rest). Same loan, same cap — the difference is the growth you ship. **The rate was never in the contract; it was in the roadmap.** The instrument is cheapest for the companies that need capital least, and the fastest-growing product is the one where equity would have been cheapest. The equity comparison is a valuation question; the rate is a growth question; growth is an engineering question.

## Where software engineering decides the terms

The term sheet is written in your code, in four places.

**Gross margin is an architecture decision.** For an AI product, margin is set in the serving layer — routing, batching, caching, prompt compression. Every cost-per-request cut is an improvement to the company's collateral; every point of margin recovered is a point the repayment cannot consume. Teams that treat margin as a per-feature budget get financed differently from teams that report it quarterly.

**Revenue quality is a data problem.** Recurring vs one-off, cohort retention, churn — outputs of your metering, billing, and analytics systems. "Is this revenue recurring?" is an instrumentation decision. Clean revenue data is the company's credit file.

**Usage-based pricing is fragile revenue.** The AI economy's default pricing is the most legible and the most fragile: readable in real time, cancelable in real time — a config change, a model switch, an agent contract lapse. Contracted MRR survives board meetings; usage revenue survives until the next API call. The lender prices the mix.

**The repayment schedule is your usage curve.** The loan is serviced by your product's throughput. Every revenue-generating agent call is a drip into the repayment — free-tier usage is not collateral. A product that grows paid usage grows its own repayment capacity.

## The lender is a software system

The counterparty is a system, not a banker: live revenue feeds, ML credit models over churn and usage curves, automated revenue-quality checks and repayment; diligence is an API integration. The lender's model asks what your analytics stack asks — *is this usage real?* — and in the AI age the artifacts of real and manufactured usage look alike: agent-driven signups, AI-inflated usage, subsidized pilots, card stacking. Revenue-quality detection is an adversarial ML problem. Lenders who do not build it will underwrite the AI economy's subprime; engineers who build products their own users can game will finance at the gaming rate.

## The fit test

RBF fits engineering profiles, not just business models.

**Perfect fit — a high-margin API product with clean metering.** 80% gross margin, twelve months of revenue history, the data self-emitted. Take $1–3M, buy growth, repay from the usage the growth creates.

**Wrong fit — the compute and training layer.** Capex sinks with delayed revenue; there is nothing to underwrite. Equity, strategic investment, and cloud credits remain the instruments.

**Tricky fit — thin margins and pass-through revenue.** Below ~60–70% gross margin, or revenue that is resold compute and tokens, the repayment eats real cash flow and the lender cannot trust the stream.

## What the honest critics say

Three objections survive scrutiny, each an engineering problem before a finance one. **The cap can cost more than the equity you saved** — on a fast-growing product, 1.8x can exceed the dilution of a modest round; do the math against the real growth rate. **Revenue is the tax base and the north star** — RBF takes a cut of the metric you optimize, hardest in the hypergrowth quarter when it repays fastest. **Revenue can be manufactured, more easily in the AI age** — revenue quality (recurring, diversified, contract-backed, human-signed) becomes the new credit score, a data-quality problem before a credit one.

## The engineering takeaway

Treat financability like a feature. Build revenue as an API — clean metering, honest telemetry, a pricing curve that does not hide churn — and the company gets cheaper to fund, because the lender's model converges faster and the collateral documents itself. The financing and the product have merged: the loan is serviced by usage, the lender underwrites the telemetry, the repayment schedule is the usage curve. A company is not a story told to investors; it is a system emitting evidence, and the quality of that evidence prices its money.

## The test

Six questions, asked the way an engineer reviews a design:

- What is the annualized cost of the cap at your actual growth rate? (If you cannot answer, you are signing a rate you do not know.)
- What is gross margin per feature, and who owns it?
- Is revenue recurring or usage-based, and does your metering make the difference visible?
- Is the capital for growth rather than R&D capex?
- Does the cap cost less, in real dollars, than the equity you would give up?
- Will the revenue still be there, and still real, in month 24 — and would your own data prove it?

The first question is the one sellers never ask; the last is the one the AI age added. The revenue that finances the AI product must exist without the AI — if the only thing generating it is the same machinery that reports it, the lender is underwriting a mirror, and the engineering team built the mirror.

> RBF prices what you have built, not what you might build — right for a product that makes more with less. The collateral is your code, the repayment is your usage, and the cost of the loan is whatever you failed to compute about your own growth. Instrument what you want financed; do the arithmetic before the lender does.

---

**References:**

- [Revenue-based financing — definition and terms (5–8% of monthly revenue; 1.5–2.5x caps)](https://www.investopedia.com/terms/r/revenue-based-financing.asp). *Investopedia.*
- [Global RBF market surpassed $9.8B in 2025, 129+ active lenders](https://www.hubspot.com/sales/revenue-based-financing). *HubSpot Sales Blog.*
- Lighter Capital — [the largest dedicated RBF lender for SaaS](https://www.lightercapital.com/).
- [Capchase — subscription and revenue financing](https://www.capchase.com/).
- [Pipe — the revenue exchange](https://pipe.com/).
- [Founderpath — RBF for bootstrapped SaaS](https://www.founderpath.com/).
- [Recur Club — revenue-based financing for growing companies](https://recurclub.com/).
- [Plaid — the data layer RBF underwriting runs on](https://plaid.com/).
- [Stripe — revenue data via API](https://stripe.com/).

**Related on this site:**

- [Engineering Is Art and Philosophy, Grounded in Economic Law](https://blog.hackspree.com/#engineering-is-economics) — the margin structure that makes AI startups financeable, as an engineering property.
- [Task Automation Economics](https://blog.hackspree.com/#task-automation-economics) — why an agent run is not automation, and how the AI product's economics differ.
- [Data Quality and the Problem of Truth](https://blog.hackspree.com/#data-engineering-quality) — revenue quality is a data-quality problem; truth is borrowing capacity.
- [Always-On Agents: State, Memory, and the Governance Gap](https://blog.hackspree.com/#always-on-agents) — the agentic operations layer that makes startup finance legible.
- [Lending Protocols](https://blog.hackspree.com/#fi-lending) — credit is 3,800 years old; the Code of Hammurabi knew the trade-offs RBF rediscovered.
