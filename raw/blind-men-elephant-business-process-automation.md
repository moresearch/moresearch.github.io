---
title: "The blind men and the elephant: why every role sees only part of the process"
date: 2026-08-28
slug: blind-men-elephant-business-process-automation
summary: "The oldest parable in operations is a story about roles, not blindness: sales sees the customer, finance sees the cash, ops sees the floor, support sees the tickets — and every one of them is telling the truth about a part of the same animal. No role can see the whole end-to-end process, because the whole does not exist inside any role. That is the real problem with business process automation: most projects automate one role's elephant and call it the process. Why roles are structurally, incentive-wise, and temporally blind to the big picture — and what automation looks like when you stop trusting a single view."
tags: [automation, business-process-automation, process-mining, roles, organizations, rpa, agents, conways-law, goodharts-law, essay, reflection]
---

![The blind men and the elephant — every role reports the part it touches, and none sees the whole.](/images/blind-men-elephant.jpg)

Six blind men are led to an elephant and asked to describe it. The man at the trunk calls it a snake. The man at the tusk calls it a spear. The man at the ear calls it a fan, the man at the leg a pillar, the man at the side a wall, the man at the tail a rope. Each report is accurate. Each man is telling the truth about what he touched. And the aggregate is wrong — not because anyone lied, but because *the whole is not contained in any of the parts*.

The parable is usually read as a lesson about humility: don't mistake your part for the whole. But there is a sharper reading, and it is the one that matters for business process automation: **the elephant is what the parable is really about, and no role can see it from where it stands.** Every organization is six blind men standing around one process, each role touching a different part of the same animal — and every automation project begins by trusting one of them.

## The big picture is an end-to-end process, and no one owns it

**The elephant is not a department, a system, or a tool. The elephant is the end-to-end process: order-to-cash, hire-to-retire, ticket-to-resolution — the whole flow from the event that starts it to the outcome that ends it.** And here is the structural fact every organization lives with: the big picture does not live inside any role. It lives *between* the roles, in the handoffs.

Walk the parts of a single order-to-cash process and name who touches what:

- **Sales** sees the customer, the opportunity, the pipeline. For sales, the process is "win the deal" — and it ends the moment the contract is signed.
- **Legal and credit** see the contract, the risk, the terms. The process for them is "make this deal safe."
- **Finance** sees the invoice, the payment, the cash. The process for them is "recognize revenue and get paid."
- **Fulfillment and operations** see the picking, the shipping, the inventory. The process for them is "get the product out the door."
- **Support** sees the tickets, the returns, the complaints. The process for them is "handle what went wrong."
- **The customer** sees only the output — the product that arrives, the invoice that shows up — and is the only one who experiences the process as one continuous thing.

None of these roles is wrong. Sales is not lying when it says the deal is won; the elephant's trunk really is a snake when you are standing at the trunk. But the end-to-end process — the thing the company is actually running — is none of these views. It is the sequence of handoffs between all of them: contract → credit check → invoice → shipment → payment → reconciliation → collection. The big picture is the flow *through* the roles, and no single role is positioned to see it.

## Why each role cannot see the whole

**The blindness is not a character flaw; it is a position. Four mechanisms guarantee that every role sees only part of the elephant, and all four are structural, not personal.**

**Position.** A role sits *inside* the flow at one point. Information moves past it the way the elephant moves past a man standing at its side: you feel a surface, never the whole body. The sales rep never sees the invoice; the invoice is upstream or downstream of where they stand. This is not an information-technology problem — you could give sales full access to every system in the company and the sales role would still not see the process, because the *role* is defined by the part it plays, and roles see what their part touches.

**Incentive.** Every role is measured on its part, and the KPI becomes the view. Sales is measured on pipeline; finance on days-sales-outstanding; ops on throughput; support on resolution time. The deeper effect is [Goodhart's law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law): a role does not just *see* its metric, it optimizes it — and the metric is the part, never the whole. The call-center agent who is measured on handle time sees every call as a handle-time problem. The finance team measured on DSO sees every invoice as a DSO problem. Each role's elephant is shaped by what it is rewarded for touching.

**Tooling.** A role's systems and dashboards render its slice. The CRM shows sales the pipeline; the ERP shows finance the ledger; the WMS shows ops the warehouse. When someone asks "how does the process work?", each role opens its own system and describes what it renders. The tool reinforces the position: the dashboard is the part, and it is very convincing, because it is *accurate* — it shows exactly the part that role touches.

**Time.** A role sees its moment in the lifecycle, not the lifecycle. The process takes weeks end-to-end; every role lives in the hours where it plays its part. The exception that took three weeks to resolve is invisible to the role that caused it on day one, because that role has already moved on to the next elephant. The parts are not just spatially separated — they are temporally separated, which is why nobody can "look around" and see the whole even if they wanted to.

## Why automation built from one role's view fails

**An automation project picks a role to believe — the one who sponsors it, the one who owns the pain — and then automates that role's elephant as if it were the process.** This is where the parable stops being a metaphor and becomes a failure mode, because the automation inherits the role's blindness and then locks it in at machine speed:

- **Automate what sales sees** and you get a beautiful quote-to-contract tool whose invoices go out wrong, because invoicing is not in sales's view.
- **Automate what finance sees** and you get perfect invoicing that ships against unapproved credit terms, because credit is not in finance's view.
- **Automate what the RPA vendor's screen-recording sees** and you get the tail: the keystrokes that twitch at the end of a process whose real structure lives upstream, in handoffs and rules the bot never touched.
- **Automate what the agent sees** and you get the newest version of the error: [an agent run is not automation](https://blog.hackspree.com/#task-automation-economics) — an agent makes a fresh decision from its context window each time, and a context window is a hand on one part of the elephant. An agent given the billing queue will optimize the billing queue, with complete confidence and complete blindness to the cost it pushes into collections.

The failure is not in the automation. The failure is in the trust: the project believed one role's account of the animal. [Conway's law](https://blog.hackspree.com/#hacker-laws-ase-conways-law) then makes the mistake durable — the systems you build mirror the roles you consulted, so an organization that automates department-shaped views ships department-shaped systems, and the elephant ends up assembled from mismatched parts with the handoffs automated nowhere. And because the [chain's throughput is set by its weakest link](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law), "automating the part" does not speed the process up; it moves the bottleneck to the next role's handoff. Order entry automated to ten seconds; invoicing still a three-day queue. The elephant did not get faster. One of its parts got faster and the rest of the animal got more frustrated.

## The whole is only visible in the handoffs

**The big picture is not a bigger version of any role's view. It is a different object entirely: the flow of work between the roles — and it is visible only where the parts meet.** This is the deepest consequence of the parable and the most useful one for automation. You cannot see the whole elephant by giving any single man a better view of his part, because the whole is not composed of parts stacked together; it is composed of the *transitions* between them — the handoffs, the waits, the exceptions, the rework.

That is exactly why [process mining](https://blog.hackspree.com/#process-mining-with-python-and-solving-real-world-data-science-tasks) matters more than any single automation tool: event logs are the only artifact that records the transitions. Every system the process touches writes a timestamped record of its own behavior, and process mining replays those logs across systems to reconstruct the actual end-to-end flow — the real paths, the real handoffs, the real waits, the real exceptions. It does not ask any role to describe the process, because no role can. It reads the elephant from its own motion, and it shows what every role's dashboard hides: the big picture lives in the handoffs, and the handoffs are where processes actually fail.

Process mining has its own limit, and honesty requires naming it: an event log is the elephant's *shadow*, not the elephant. It records behavior, not intent — the invoice disputed because the contract was ambiguous shows up as a long edge in the graph, and the log cannot tell you why. So the whole is visible only to a *combination*: the measured flow, plus the humans from each role who can interpret their part of the shadow. The elephant-walker is not a role that sees everything; the elephant-walker is the person who assembles all the roles' partial truths into one animal — and for the first time, process mining gives them accurate parts to assemble.

## What automation looks like when no role is trusted with the whole

**The fix is not to find the role that sees everything — there is no such role. The fix is to stop building automation from any single role's view.** Concretely:

- **End-to-end ownership before end-to-end automation.** Give the process an owner whose job is the handoffs, not a department — the person whose elephant is the whole flow, and who reports that the tail is attached to something large.
- **Event logs as ground truth, not role interviews.** If the project cannot produce a conformance report of the actual flow, it does not know what it is automating; it only knows what one role believes.
- **Cross-role review before anything ships.** Every automation must answer three questions: what role's view did we build this from, what does it optimize, and whose downstream elephant does it break? Asked before the code is written.
- **Widen the agent's aperture like the humans'.** The harness-engineering move is the same move as process mining: give the agent [the whole system, not the polite demo](https://blog.hackspree.com/#good-harnesses-watch-the-whole-operating-system) — [real repositories](https://blog.hackspree.com/#coding-agent-harnesses-need-real-repositories), [real browsers](https://blog.hackspree.com/#harnesses-need-real-browsers-not-polite-demos), [tasks that fight back](https://blog.hackspree.com/#agent-harnesses-need-tasks-that-fight-back). An agent automating from one context window is the newest blind man, and the most dangerous, because it is confident, fast, and cannot see past its part.
- **Keep a human whose job is the whole.** [The elephant is a system](https://blog.hackspree.com/#i-pencil), and a system cannot be seen from inside one of its parts — the dark-factory dream ([the complexity](https://blog.hackspree.com/#dark-factory-complexity), [the economics](https://blog.hackspree.com/#factory-is-not-dead)) does not delete this job. It makes the job be only this job.

## Key insight

**The point of the parable is not that the men are blind — it is that each of them is right, and the elephant is still missing.** Every role in an organization is telling the truth about the part it touches: sales really does win the deal, finance really does get paid, ops really does ship, support really does resolve. The whole is not in any of their reports, and no amount of honesty, data, or goodwill inside a single role can produce it. Business process automation fails when it trusts one role's elephant; it works when it treats the big picture as something no role can see alone — measured across the handoffs, assembled from every role's partial truth, and automated only once the whole animal is in view. Automation is not the goal. Seeing the elephant is the goal, and the only way to see it is together.

## References

- The parable — Jain, Hindu, and Buddhist versions; popularized in English by John Godfrey Saxe's poem "The Blind Men and the Elephant" (1873) — https://en.wikipedia.org/wiki/Blind_men_and_an_elephant
- Wil van der Aalst, *Process Mining: Data Science in Action* — https://www.springer.com/gp/book/9783662498507
- IEEE Task Force on Process Mining, Process Mining Manifesto (2011) — https://www.tf-pm.org/resources/manifesto
- Image: LinkedIn — https://media.licdn.com/dms/image/v2/D4E10AQGddprHhTvV2Q/image-shrink_800/B4EaBHUnZ5GgAg-/0/1787902982366
- Archive: [process mining with Python](https://blog.hackspree.com/#process-mining-with-python-and-solving-real-world-data-science-tasks), [why an agent run is not automation](https://blog.hackspree.com/#task-automation-economics), [good harnesses watch the whole operating system](https://blog.hackspree.com/#good-harnesses-watch-the-whole-operating-system), [Conway's law](https://blog.hackspree.com/#hacker-laws-ase-conways-law), [Goodhart's law](https://blog.hackspree.com/#hacker-laws-ase-goodharts-law), [Amdahl's law](https://blog.hackspree.com/#hacker-laws-ase-amdahls-law), [dark factories: complexity](https://blog.hackspree.com/#dark-factory-complexity), [the factory is not dead](https://blog.hackspree.com/#factory-is-not-dead), [I, Pencil](https://blog.hackspree.com/#i-pencil)
