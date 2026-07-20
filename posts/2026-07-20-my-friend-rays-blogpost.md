---
title: "My Friend Ray's Blog Post Hit the Top of Hacker News, and It Deserved To"
date: 2026-07-20
slug: my-friend-rays-blogpost
summary: "Ray Myers' post on the Zig/Bun/Anthropic controversy is the kind of writing that takes a year of reading to produce and a lifetime of clear thinking to pull off. 1,546 points on HN. It earned every one."
tags: friend, writing, ai, software-engineering, reference-density
---

Ray and I meet semi-regularly to discuss everything software. The meetups are the kind where you show up without an agenda and leave three hours later with a notebook full of things to read and a slightly rearranged understanding of the field. Ray has a habit of tracking down the reference behind the reference — the paper the blog post cited, the study the paper was based on, the replication that corrected the study. Talking to him is like being handed a curated reading list in real time.

Last week, his blog post [*Zig Creator Calls Spade a Spade, Anthropic Blows Smoke*](https://raymyers.org/post/zig-creator-calls-spade-a-spade/) hit the top of Hacker News. 1,546 points. 784 comments. It stayed there.

## What the post is about

If you missed the controversy: Anthropic's Claude helped Bun rewrite itself from Zig to Rust. The rewrite was presented as a technical decision driven by memory safety concerns. Andrew Kelley, the creator of Zig, responded with characteristic directness — pointing out that the codebase was a mess due to engineering decisions, including overusing AI agents, and that Zig was being blamed for problems it didn't cause.

Ray's post is the definitive analysis of what actually happened. But that undersells it.

## AI is not enough

Ray's core insight — the one that organizes the entire post — is that **every element of the story that Bun and Anthropic present as evidence that AI replaces engineers actually demonstrates the opposite.** The million-line agentic PR? Required human review and cleanup that the AI couldn't do. The automated language migration? Left semantic gaps that only a human could close. The speed of the rewrite? Achieved by a team working 90-hour weeks, not by AI magic.

> The post's thesis isn't "AI is bad." It's "AI is not enough." Every layer of the story, examined carefully, reveals a place where human judgment, human design, and human communication were the difference between success and failure. The AI was a tool. The humans were the engineers. The story Anthropology wants to tell — that software engineering is becoming obsolete — is contradicted by the very case study they chose to prove it.

## The reference density

Here is the thing that makes Ray's writing distinctive, and the thing I want to call out specifically.

Most blog posts cite a tweet and a Wikipedia article and call it a day. Ray's post cites **38 references**, including:

- The DARPA TRACTOR program's formal evaluation report
- TigerBeetle's TigerStyle documentation and simulation testing methodology
- A Buddhist framework for right speech, applied to engineering communication
- The Ghostty AI policy and RedMonk's analysis of generative AI governance in open source
- Ed Zitron's financial analysis of Anthropic's profitability problem
- Hillel Wayne's empirical software engineering talk
- Marianne Bellotti's *Kill It With Fire* and Dr. Cat Hicks' *The Psychology of Software Teams*

Each reference is not dropped for credibility. Each one is woven into the argument. The DARPA report isn't cited to show Ray reads government documents. It's cited because the TRACTOR program ran into the exact same semantic-gap problem that Bun's AI rewrite ran into, and the program's evaluation report documented it. The Buddhist right-speech framework isn't cited for intellectual decoration. It's cited because the question at the heart of the controversy — when is blunt truth-telling ethical? — has been analyzed for millennia, and the framework helps.

> The post is the output of a process that most people never see: reading widely, tracking references to their source, building a mental map of how ideas connect, and then writing only when the map is dense enough to support an argument. The post took a week to write. The reading took a career.

## Why the post earned its spot

Hacker News is fickle. The front page is a random variable with a tech-industry-shaped distribution. But some posts earn their spot through sheer substance. Ray's is one of them.

The post does three things simultaneously, and does all three well:

1. **It tells the story** — who said what, when, and why it matters
2. **It analyzes the incentives** — Anthropic needs a narrative to justify a $1 trillion valuation ambition; the Bun rewrite was marketing as much as engineering
3. **It builds a counter-narrative** — not "AI is bad" but "AI is not enough," supported by references that would take a year to track down and study independently

Most blog posts do one of these. The good ones do two. Ray's does all three, with prose that is clear without being simplistic and rigorous without being academic.

## Go read it

If you work in software and haven't read [Ray's post](https://raymyers.org/post/zig-creator-calls-spade-a-spade/), you should. If you don't work in software but care about how technology companies construct narratives, you should read it too. The post is about Zig and Rust and Anthropic and Bun. But the lesson is about something bigger: how to think clearly when the loudest voices in the room have $132 billion worth of reasons to be heard.

Ray and I will discuss it at our next meetup. He'll probably have tracked down five more references by then. That's how he works.

---

*[Zig Creator Calls Spade a Spade, Anthropic Blows Smoke](https://raymyers.org/post/zig-creator-calls-spade-a-spade/) by Ray Myers. [Hacker News discussion](https://news.ycombinator.com/item?id=48889637) (1,546 points, 784 comments).*
