---
title: "Building AI Factories: What a Stanford Lecture Reveals About Where the Money Goes"
date: 2026-07-19
slug: ai-infrastructure-economics-stanford-msande435
summary: "Stanford MS&E 435's Class #3, featuring Crusoe CEO Chase Lochmiller, lays bare the physical economics of AI: $60M per megawatt, 2.1 GW campuses, and why the semiconductor layer captures 75% of AI revenue."
tags: ai, infrastructure, economics, data-centers, energy, semiconductors, stanford
---

Apoorv Agrawal's Stanford course MS&E 435 — *Economics of the AI Supercycle* — has been quietly accumulating tens of thousands of views on YouTube. The course's central thesis is that generative AI breaks the traditional software playbook: unlike past tech cycles where value migrated upward to applications, AI's market structure is an **inverted triangle**. Roughly 75% of the ~$350 billion in new AI ecosystem revenue has gone straight to the semiconductor layer. Application margins hover between 0% and 30%.

Class #3 of the course, featuring **Chase Lochmiller** — co-founder and CEO of Crusoe — is where the abstraction hits the concrete. Lochmiller is the lead infrastructure developer for Project Stargate's Abilene, Texas site. His lecture is a masterclass in the physical economics of AI: what it actually costs to build the factories that produce tokens, where the bottlenecks live, and why the value chain looks the way it does.

## The numbers that reset the conversation

Lochmiller shared the per-megawatt economics that underpin the entire AI infrastructure buildout:

| Layer | Cost per MW |
|---|---|
| Data center shell + power plant (CAPEX) | ~$20M |
| GPU hardware ($30M) | |
| InfiniBand/ROCE networking ($4M) | |
| CPU + storage ($3M) | |
| **IT hardware subtotal** | **~$40M** |
| **Total upfront per MW** | **~$60M** |

At gigawatt scale, that's roughly **$60 billion** all-in for a single campus.

The revenue side splits into two tiers. Infrastructure-only leasing (powered shell, cooling, connectivity) generates about $15M per MW annually — a four-year payback. Managed compute clusters, where the operator hosts models and runs API services on top of the hardware, generate about $30M per MW annually — a two-year payback. The difference between those two numbers is the margin that accrues to whoever controls the compute stack rather than just the real estate.

> The economic gradient points upward: the more of the stack you operate, the faster your capital comes back. But the capital required to play at the top of the stack is orders of magnitude larger.

## Project Stargate: what "largest buildout in human history" looks like

The Abilene campus is the lecture's concrete anchor. Lochmiller described it as "probably the largest buildout of infrastructure in human history." The numbers bear that out:

- **1.2 to 2.1 gigawatts** of power capacity — roughly equivalent to powering two cities the size of Denver
- **1,200 acres**, with eight data center buildings planned
- **Up to 400,000 GPUs** (NVIDIA Blackwell GB200 NVL72 racks)
- **7,000–9,000 workers on-site daily**, in a city of 120,000
- **America's largest private substation**

Crusoe broke ground in June 2024 on land that was, in Lochmiller's words, "dirt and mesquite trees." The first two buildings started then. The next six started in February 2025. Crusoe beat competitors' fastest bids of 2.5 years by delivering in roughly 12 months.

The speed matters economically. Every month a gigawatt-scale facility sits idle is tens of millions in capital costs with zero revenue. The bottleneck is not concrete — it's transformers, switchgear, skilled electrical labor, and the permitting timelines that govern grid interconnection. Lochmiller noted ~9,000 workers are needed on-site daily for builds of this scale, with acute shortages of electricians and welders.

## The water myth and the cooling reality

One detail Lochmiller addressed directly: each data hall contains about 1 million gallons of cooling water. Headlines routinely convert that number into a consumption figure. But modern facilities use closed-loop systems — once filled, the annual water consumption equals roughly one average household. The water sits in the loop; it doesn't evaporate at scale.

> The million-gallon figure is a stock, not a flow. Confusing the two turns an engineering detail into a misleading narrative.

This matters because water access is becoming a siting constraint. Communities that would welcome the tax base of a data center campus may resist on water-use grounds. Getting the facts right changes which sites are viable.

## Crusoe's origin: from wasted methane to AI infrastructure

The most unexpected part of the lecture was Crusoe's origin story. Lochmiller was previously a quant portfolio manager using deep learning for financial trading. Crusoe began by capturing **flared and waste methane from oil fields** — gas that would otherwise be burned off or vented — to power modular, shipping-container data centers. The initial monetization was Bitcoin mining, but the plan was always to pivot to AI infrastructure once the compute demand materialized.

This origin shaped Crusoe's strategy in two ways. First, it forced the company to think about energy *before* compute — to site facilities where energy is abundant and cheap rather than where fiber is densest. Second, it forced vertical integration. Crusoe built its own manufacturing arm, **Crusoe Industries**, to produce electrical equipment like power distribution centers in 20 weeks, versus the industry standard of 100 weeks. When you're competing against Eaton and Schneider on timelines, owning your supply chain is not a luxury.

## "Across the Meter": the energy strategy that enables the economics

Traditional data center siting follows fiber and network proximity — Northern Virginia, Santa Clara, Ashburn. Crusoe's thesis is the inverse: follow the energy. West Texas has abundant curtailed wind power — turbines that are routinely shut off because transmission lines are congested and power prices go negative.

Crusoe's "Across the Meter" approach co-locates data centers with wind, solar, battery storage, and natural gas generation. The facility draws power directly from generation assets rather than through congested transmission. Excess power is sold back to the grid. For AI training workloads and most inference, latency is irrelevant — a data center in Abilene serves a model as well as one in Ashburn.

This strategy exploits a structural inefficiency: the U.S. has abundant generation capacity in places where transmission infrastructure is decades behind. Building a data center at the generation source bypasses the transmission bottleneck entirely. The economic arbitrage is the spread between the locational marginal price of power at the generation node and the price at the load center — which can be multiples.

## The investment implications Lochmiller shared

Lochmiller offered a candid set of market views:

- **Short-term bullish, long-term bearish** on legacy electrical equipment companies (Eaton, Schneider). The current buildout is a demand shock they cannot meet, but the supply response — from Crusoe's own manufacturing arm and competitors — will erode their pricing power over time.
- **Bullish on solid-state transformers** and power electronics innovation. The transformer is the least-innovated component in the electrical grid. That changes when a 2.1 GW campus needs transformers that don't exist in standard catalogs.
- **Bullish on space-based data centers** as a long-term play: optical interconnectivity, zero permitting, and unlimited solar power. Thermal management in vacuum remains unsolved, but Lochmiller treated it as an engineering problem rather than a science fiction premise.

## Why the inverted triangle persists

The lecture makes Agrawal's inverted triangle thesis concrete. Nvidia's 75% gross margins are not a temporary anomaly caused by supply shortage. They are structural: every new AI user burns GPU compute, and the GPU supplier captures the rent. The infrastructure layer — where Crusoe operates — is the next most concentrated. The application layer, where ChatGPT earns roughly $10 per user annually against Alphabet's ~$100, is where the margin compression lives.

Lochmiller's numbers explain why. At $60M per MW upfront, the capital barrier to entering the infrastructure layer is enormous. Once built, a gigawatt campus has pricing power because the alternative — building another one — takes years and billions. The semiconductor layer has even higher barriers: Nvidia's CUDA moat, the multi-year lead time on advanced packaging capacity at TSMC, the trillion-dollar cumulative R&D investment in the GPU architecture.

The application layer has none of these defenses. Switching costs between chatbots are near zero. Model quality gaps compress with each release cycle. And the "active work" bottleneck Agrawal identifies — AI tools require users to formulate queries and engage actively, unlike the passive consumption of social feeds — limits organic growth in ways that don't apply to prior platform cycles.

## What the lecture leaves open

Lochmiller's lecture is deliberately about the physical layer — concrete, steel, copper, silicon, and electrons. It does not answer the question of whether the application layer eventually finds a monetization model that closes the gap. Agrawal's broader course argues that advertising is the likely path: AI platforms see deep, logged-in intent signals that search and social cannot match, making AI-delivered ads potentially more valuable per impression than any existing channel.

But that thesis is speculative. What is not speculative is that the physical layer is being built at a scale and speed that has no precedent in civilian infrastructure. The lecture makes clear that this is not a cloud cycle replay. The cloud cycle eventually inverted — software captured the surplus once the infrastructure was built. The AI cycle may stay inverted much longer because the inference compute cost is structural rather than temporary.

For founders, the implication is that building at the application layer without a thesis about monetization — and specifically about monetization that scales with compute cost — is building on someone else's margin. For investors, the near-term bet remains silicon and the infrastructure that powers it. For everyone else, Lochmiller's lecture is a reminder that the most important numbers in AI are not parameter counts or benchmark scores. They are dollars per megawatt, weeks per transformer, and gigawatts per campus.

---

**Source:** [Stanford MS&E 435: Economics of the AI Supercycle, Class #3](https://www.youtube.com/watch?v=sRvrXL83N-c) — Chase Lochmiller (Crusoe CEO), Spring 2026, taught by Apoorv Agrawal. Course website: [mse435.stanford.edu](https://mse435.stanford.edu/). Additional context from Josipa Majic Predin's [Forbes coverage](https://www.forbes.com/sites/josipamajic/2026/04/25/a-stanford-lecture-explains-why-ai-value-gets-trapped-in-chips/) of the course.
