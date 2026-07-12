---
title: On Data Variability: the gap between peak and average
date: 2026-07-12
slug: v-variability
summary: "Variability is the fluctuation in data characteristics over time. It breaks pipelines. A source that produces 1,000 events per hour suddenly produces 100,000. A column of integers suddenly contains strings. The pipeline wasn't designed for this. Variability is the reason pipelines need monitoring, alerting, and circuit breakers."
tags: data-engineering, variability, peaks, elasticity, capacity-planning
series: the-vs-of-data
part: 6
---

> Variability is not the peak. It is the gap between the peak and the average. The average is fiction. The peak is the design constraint. Designing for the average is designing for failure at the moment of maximum need.

Variability is the fluctuation in data characteristics over time. It is the dimension that breaks pipelines. A data source that produces 1,000 events per hour suddenly produces 100,000. A column that has always contained integers suddenly contains strings. A partition that has always had data is suddenly empty. The pipeline was designed for the steady state. The steady state is a lie. The world is not steady.

A food delivery platform illustrates the constraint. Order volume follows a predictable pattern: peaks at lunch and dinner, troughs in between, higher on weekends. The pipeline is sized for 2× the average peak. On Super Bowl Sunday, order volume spikes to 20× the normal peak. The ingestion pipeline, sized for 2×, falls behind. The transformation pipeline, expecting 1 million rows, receives 20 million. Dashboard queries time out. The data platform fails at the moment it is most needed — when the business wants to know how many orders were placed, how many were delivered, what the average delivery time was. The numbers are unavailable. The unavailability is the cost of designing for the average.

The engineering response is elastic infrastructure. Cloud warehouses scale compute on demand. Pipelines auto-scale to handle spikes. The infrastructure adapts to the load. The adaptation has limits: scaling takes time (minutes, not seconds), and the peak may exceed the maximum scale. The limits are the constraint. The constraint must be managed by load shedding — dropping non-critical work during spikes to preserve critical work. The load shedding is a policy. The policy must be defined before the spike. Defining it during the spike is too late.

The deeper response is defensive pipeline design. Pipelines should degrade gracefully under load — process what they can, queue what they can't, alert on what they're dropping. Pipelines should be tested against synthetic spikes — double the volume, triple the volume, ten times the volume — to find the breaking point before production does. Pipelines should have circuit breakers — if the output is anomalous (zero rows, ten times the expected count, nulls in required fields), stop the pipeline and alert. The circuit breaker prevents bad data from propagating downstream. Stopping the pipeline is better than publishing wrong numbers. Wrong numbers are worse than no numbers. No numbers prompt a question. Wrong numbers prompt a wrong decision.

*See: John Allspaw, "The Art of Capacity Planning" (O'Reilly, 2008) — the operations engineering approach to sizing for variability. Betsy Beyer et al., "Site Reliability Engineering" (O'Reilly, 2016), Chapter 13, on managing overload. Daniel Abadi et al., "The Design and Implementation of Modern Column-Oriented Database Systems" (Foundations and Trends in Databases, 2013) — why cloud elasticity changes the trade-offs.*


*This post is part of a series on [The Many Vs of Data](https://blog.hackspree.com/#many-vs-of-data), originating from Doug Laney's 2001 Gartner note. Each V names a dimension of why data is hard.*
