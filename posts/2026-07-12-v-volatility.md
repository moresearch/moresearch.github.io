---
title: Data Volatility
date: 2026-07-12
slug: v-volatility
summary: "Volatility is the rate at which data loses relevance. Stock prices matter for milliseconds. Medical records matter for decades. The half-life of data value determines how much to invest in its quality and how long to keep it."
tags: data-engineering, volatility, retention, time-to-live, data-lifecycle
series: the-vs-of-data
part: 10
---

> Volatility is not about how fast data changes. It is about how fast data dies. Every dataset has a half-life. After the half-life, the storage cost exceeds the remaining value. Keeping data past its half-life is not preservation. It is waste.

Volatility is the rate at which data loses relevance over time. It is the dimension that determines retention policy. High volatility: the data is valuable for milliseconds, then worthless. Low volatility: the data is valuable for decades. The half-life of data value — the time after which the data is half as useful as it was when created — determines how much to invest in its quality and how long to store it. A dataset with a half-life of one day should be stored on cheap storage, minimally curated, and deleted after a month. A dataset with a half-life of ten years should be stored on reliable storage, carefully curated, and retained indefinitely.

A stock trading system illustrates high volatility. Real-time price data is valuable for milliseconds — the window for executing an arbitrage. After the window closes, the data's value drops to near zero for trading. But it remains valuable for backtesting and compliance — the dual value curve: immediate and high, then residual and low. The curve determines the storage architecture: in-memory for the trading window, archival for compliance. The architecture is tiered because the volatility is dual. A single tier would either be too expensive (keeping millisecond-old data on fast storage forever) or too slow (keeping compliance queries waiting for archival retrieval).

Most organizations have no retention policy. The absence of a policy is a policy — keep everything forever. Forever is expensive. The storage cost is linear. The value of the oldest data is near zero. The gap between cost and value widens every year. The gap is waste. The waste is invisible because the storage bill is aggregated.

The engineering response is automated data lifecycle management. Define retention rules: transactional data kept for 7 years (compliance), web analytics kept for 2 years, ML training features kept for 6 months or until model retraining. Enforce the rules automatically — pipelines that delete or archive data when it exceeds its retention period. The enforcement must be auditable — if data is deleted, there must be a record of what was deleted, when, and under what policy. The audit trail is the defense against the accusation of destroying evidence. The accusation is rare. The defense is necessary.

Gordon Moore's 1965 paper predicted that transistor density would double every two years. The prediction held. Storage became cheap enough that keeping everything seemed rational. The rationality was an illusion. Storage is cheap per gigabyte. The total cost — storage plus maintenance plus cognitive load — is not cheap. The total cost accumulates. The accumulation is the reason retention policies exist. The policy is the recognition that data has a finite useful life. The useful life is shorter than the organizational memory. The data outlives its usefulness. The outliving is the problem. The deletion is the solution.

*See: Apache Iceberg, "Table Spec: Snapshots and Time Travel" (Iceberg Documentation) — on the table format features that make retention management programmable. Netflix Tech Blog, "Evolution of Data Lifecycle Management at Netflix" (2021) — a real-world architecture for automated data retention. Gordon Moore, "Cramming More Components onto Integrated Circuits" (Electronics, 1965).*


*This post is part of a series on [The Many Vs of Data](https://blog.hackspree.com/#many-vs-of-data), originating from Doug Laney's 2001 Gartner note. Each V names a dimension of why data is hard.*
