---
title: On the Many Vs of Data
date: 2026-07-12
slug: many-vs-of-data
summary: "In 2001, Doug Laney wrote a Gartner note describing three dimensions of the data challenge: Volume, Velocity, Variety. It was not meant to launch a framework. It launched a framework. The Vs became the vocabulary for describing why data is hard."
tags: data-engineering, big-data, volume, velocity, variety, laney
series: the-vs-of-data
part: 0
---

In 2001, Doug Laney, an analyst at Gartner, wrote a research note titled "3D Data Management: Controlling Data Volume, Velocity, and Variety." The note was internal. It described three dimensions of the growing data management challenge. It was not meant to launch a framework. It launched a framework.

The "three Vs of big data" became the standard vocabulary for describing why data was getting harder to manage. The framework was simple enough to remember, flexible enough to extend. Over the following two decades, practitioners and researchers added more Vs — each naming a specific dimension of the challenge, each adding a word to the vocabulary.

The genius of the Vs is not their precision. It is their utility. They give engineers a shared language for diagnosing why a data project is struggling. Is the volume overwhelming the storage tier? Is the velocity exceeding the pipeline's throughput? Is the variety creating an unmanageable number of pipelines? Is the veracity undermining trust in the numbers? The Vs name the dimensions of difficulty. Naming them is the first step to managing them.

This series explores each V in depth. Each post asks: what does this V mean, why does it matter, what is the engineering constraint it imposes, and how do you respond?

1. [Data Volume](https://blog.hackspree.com/#v-volume) — how much you have, and how much you can afford to keep
2. [Data Velocity](https://blog.hackspree.com/#v-velocity) — how fast data arrives, and how fast you must decide
3. [Data Variety](https://blog.hackspree.com/#v-variety) — the diversity of formats, and the cost of integration
4. [Data Veracity](https://blog.hackspree.com/#v-veracity) — trustworthiness, and the gap between data and truth
5. [Data Value](https://blog.hackspree.com/#v-value) — the only V that justifies the platform's existence
6. [Data Variability](https://blog.hackspree.com/#v-variability) — fluctuation over time, and why pipelines break
7. [Data Visualization](https://blog.hackspree.com/#v-visualization) — the bridge between data and decision
8. [Data Viscosity](https://blog.hackspree.com/#v-viscosity) — resistance to movement, and technical debt in the supply chain
9. [Data Virality](https://blog.hackspree.com/#v-virality) — how data usage spreads through people
10. [Data Volatility](https://blog.hackspree.com/#v-volatility) — the half-life of data value

The Vs are not a taxonomy. They are a diagnostic tool. When a data project is struggling, ask: which V is the problem? The answer points to the solution.

---

**Reference:** Doug Laney, "3D Data Management: Controlling Data Volume, Velocity, and Variety," Gartner, 2001.
