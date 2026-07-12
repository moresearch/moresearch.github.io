---
title: The economics of data
date: 2026-07-12
slug: data-engineering-economics
summary: "Data is not free. It costs money to collect, store, transform, and serve. It generates value when used and costs when ignored. The economics of data engineering is the economics of an asset that depreciates faster than any other: data that was valuable yesterday may be worthless today, but it still costs money to store."
tags: data-engineering, economics, cost, build-vs-buy, data-value
series: data-engineering
part: 6
---

Data is an asset. Data is also a liability. The distinction is whether the data generates more value than it costs to maintain. Most organizations treat all data as an asset. Most data is a liability. The difference between the two is the difference between a data platform that pays for itself and a data platform that is a cost center the CFO wants to cut.

## The cost structure

The cost of data has four components. Each is measurable. Few organizations measure them.

**Ingestion cost.** The engineering time to build and maintain data pipelines. The compute cost of running them. The cost of the ingestion tools (Fivetran, Airbyte) or the infrastructure they run on. Ingestion cost scales roughly linearly with the number of data sources.

**Storage cost.** The cost of storing data in the warehouse or data lake. Storage is cheap — $20-30 per TB per month in cloud warehouses. Storage is also unbounded — data accumulates indefinitely. A table that cost $100/month to store last year costs $200/month this year. The growth is compound. The compound growth is invisible until someone looks at the storage bill.

**Transformation cost.** The compute cost of running transformations. Every `dbt run` consumes warehouse credits. Every hourly refresh of the dashboard tables consumes credits. Transformation cost scales with data volume and refresh frequency. A dashboard that refreshes every hour costs 24 times more than a dashboard that refreshes once daily. The refresh frequency is a business decision. The cost is an engineering consequence.

**Serving cost.** The compute cost of querying the data. Every dashboard load, every ad-hoc query, every ML training job consumes warehouse credits. Serving cost scales with the number of users and the complexity of their queries. A poorly written query that scans the entire table costs more than a query that uses partitions and filters. The query author doesn't see the cost. The engineering team pays the bill.

The total cost is the sum of these four. For a typical mid-size organization, the breakdown is roughly: 15% ingestion, 25% storage, 35% transformation, 25% serving. The exact numbers vary. The pattern is consistent: most of the cost is in transformation and serving, not storage. Storage is the scapegoat. Transformation is the real expense.

## Why data projects fail

Data projects fail for economic reasons disguised as technical reasons. The project was "too complex." The pipeline "couldn't scale." The data "wasn't reliable." These are symptoms. The cause is that the project's costs exceeded its benefits, and nobody measured either.

A data project has benefits: faster decisions, better decisions, new revenue from data products, reduced cost of manual reporting. The benefits are diffuse — they accrue to many people across the organization. They are hard to measure. They are easy to overstate in the business case and impossible to verify after the fact.

A data project has costs: engineering time, compute, storage, ongoing maintenance. The costs are concentrated — they are paid by the data engineering team. They are easy to measure. The data engineering team knows exactly how much the pipeline costs to build and run. They are rarely asked.

The asymmetry — diffuse benefits, concentrated costs, neither measured accurately — produces predictable outcomes. The project is approved based on overstated benefits. The project is built. The costs exceed expectations. The benefits are invisible because nobody is measuring them. The project is declared a failure. The postmortem blames the technology. The technology was not the problem. The economics were the problem. The economics were never made explicit.

## Build vs. buy

The build-vs-buy decision in data engineering is the same as in software engineering: build if the internal cost is less than the external cost, adjusted for risk, control, and strategic value. The data-specific twist: data tools have extreme economies of scale. A managed service (Fivetran, Snowflake, dbt Cloud) amortizes its development cost across thousands of customers. An internal pipeline tool amortizes its cost across one organization. The managed service is almost always cheaper for common use cases — ingesting from standard sources, transforming with SQL, serving dashboards. The internal tool is justified only when the use case is unique to the organization or when the strategic value of control exceeds the cost premium.

The trap: organizations underestimate the maintenance cost of internal tools. The initial build is a month. The maintenance is years — bug fixes, feature requests, onboarding documentation, operational support. The maintenance cost is proportional to the number of users and the rate of change of the tool's dependencies. The number of users grows. The rate of change increases. The maintenance cost compounds. The internal tool that seemed cheaper than Fivetran at year one is more expensive than Fivetran by year three. The cost has shifted from the vendor to the internal team. The shift is invisible because internal labor is a fixed cost. The fixed cost is already paid. The marginal cost of asking the internal team to maintain another tool appears to be zero. It is not zero. It is opportunity cost — the features they could have built instead of maintaining the ingestion tool.

## The value of data

Data has value when it is used to make a decision that produces a better outcome than the decision that would have been made without it. The value is the difference between the outcome with data and the outcome without. The value is measurable in principle. It is almost never measured in practice.

Data that is collected but never used has negative value — it cost money to collect, store, and maintain, and it generated zero benefit. The storage cost accumulates. The maintenance cost accumulates. The value is zero. The net is negative. The negativity is invisible because the costs are aggregated into the data platform budget and the zero benefit is never calculated. The calculation would require asking: "what decision did this dataset inform, and what was the outcome?" Nobody asks. The data accumulates. The costs accumulate. The value is zero. The net is negative.

The most valuable data in an organization is often the data that doesn't exist yet — the data that would answer a question the business has been asking for months but nobody has had the time to pipeline. The value is latent. The latency is a prioritization failure. The prioritization failure is an economic failure. The economic failure is that the cost of building the pipeline was compared to the engineering time required, not to the value of the decisions it would inform. The comparison was never made. The pipeline was never built. The decisions were made without data. The outcomes were worse than they could have been. The difference is the cost of not building the pipeline. The cost is real. It is unmeasured.

## The depreciation problem

Data depreciates faster than any other asset. Customer behavior data from 2020 has limited relevance to customer behavior in 2026. The relevance decays over time. The storage cost is constant. The value declines. The crossover point — when the storage cost exceeds the remaining value — arrives faster than organizations expect. Most organizations never delete data. The data accumulates. The storage cost grows. The average value of the stored data declines. The decline is the depreciation.

Depreciation should be accounted for. Data that is older than its useful life should be archived or deleted. The useful life depends on the domain: transaction records (7+ years, for compliance), web analytics (2 years), ML training features (6 months, or until the model is retrained). The retention policy should be explicit. The policy should be enforced automatically. The enforcement should be a pipeline that deletes old data. Most organizations have no retention policy. The absence of policy is a policy — keep everything forever. Forever is expensive.

## The discipline

The discipline of data economics is the discipline of measuring what matters. Measure the cost of every pipeline: build cost, run cost, maintenance cost. Measure the value of every dataset: what decisions does it inform, what outcomes does it improve? Compare the two. Keep the datasets where value exceeds cost. Kill the rest. The killing is the discipline. The discipline is rare. The rarity is why data platforms are cost centers. The cost center is a choice. The choice is to not measure.

---

**References:**
- Benn Stancil, "The Data Platform Cost Model," Mode Blog, 2020.
- Martin Kleppmann, *Designing Data-Intensive Applications*, O'Reilly, 2017.
- Related posts: [Scarcity Rules Everything](https://blog.hackspree.com/#scarcity), [No Free Lunch](https://blog.hackspree.com/#scarcity-and-software-economics), [Engineering is economics](https://blog.hackspree.com/#engineering-is-economics)
