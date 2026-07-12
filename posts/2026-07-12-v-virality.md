---
title: Data Virality
date: 2026-07-12
slug: v-virality
summary: "Virality is the rate at which data usage spreads through an organization. A dataset that one team finds useful spreads to other teams. The spread is the measure of value. A dataset that nobody knows about is worthless, regardless of its quality."
tags: data-engineering, virality, adoption, data-catalog, discoverability
series: the-vs-of-data
part: 9
---

> Virality is not about how fast data moves through pipelines. It is about how fast data moves through people. A dataset that nobody knows about has no value. Discoverability is the prerequisite for value. Without it, the best dataset in the world is invisible.

Virality is the rate at which data usage spreads through an organization. It is the dimension that separates successful data platforms from unsuccessful ones. A dataset that one team finds useful spreads to other teams. The recruiting team discovers the employee movement dataset and uses it for time-to-fill metrics. The finance team discovers it and uses it for headcount forecasting. The facilities team discovers it and uses it for office space planning. The dataset spreads because it answers questions that multiple teams have. The spread is the evidence of value.

The opposite of virality is invisibility. The data exists. It is accurate, timely, well-modeled. Nobody knows about it. The team that built it uses it. Nobody else does. The dataset is technically successful and organizationally irrelevant. The irrelevance is not the fault of the data. It is the fault of the platform. The platform made no provision for discoverability — no catalog, no documentation, no mechanism for users to find datasets they didn't already know existed. The platform is a library with no catalog. The books are there. You cannot find them.

The engineering response to the virality problem is the data catalog. A catalog — Alation, Atlan, DataHub, Amundsen — indexes all data assets across the organization. It shows lineage — where data came from and where it goes. It shows ownership — who is responsible. It shows quality metrics — when the dataset was last validated. It shows usage — who queries it, how often. The catalog makes data discoverable. Discoverability is the precondition for virality. You cannot use data you cannot find.

But discoverability is not enough. The data must also be understandable. A dataset with cryptic column names and no documentation is discoverable but unusable. The user finds it, opens it, sees `col_37 VARCHAR`, and closes it. The data dictionary — column descriptions, business definitions, example values — is the mechanism for understandability. The data quality dashboard — freshness, completeness, accuracy metrics — is the mechanism for trust. Discoverability, understandability, trust. All three are required for virality. Most platforms have one. Some have two. Few have all three. The ones that have all three have viral data.

*See: Prukalpa Sankar, "The Data Maturity Curve" (Atlan, 2021) — on the stages of data culture from fragmented to viral. Michelle Casbon et al., "Data Governance: The Definitive Guide" (O'Reilly, 2021) — on how governance enables rather than restricts data adoption. Alation, "The Data Catalog: The Foundation of Data Culture" (Whitepaper, 2020).*


*This post is part of a series on [The Many Vs of Data](https://blog.hackspree.com/#many-vs-of-data), originating from Doug Laney's 2001 Gartner note. Each V names a dimension of why data is hard.*
