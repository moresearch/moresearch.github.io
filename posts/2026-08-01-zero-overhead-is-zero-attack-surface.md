---
title: Zero Overhead Is Zero Attack Surface
date: 2026-08-01
slug: zero-overhead-is-zero-attack-surface
summary: The supply chain attack doesn't break in — it is invited. Every dependency is a delegation of trust, every plugin an entry point, every auto-update a remote code execution channel. The industry's answer is more machinery; the only defense that compounds is subtraction. This is why zot — zero-overhead-tool — is built small on purpose.
tags: security, supply-chain, coding-agents, simplicity, developer-tools, zot
---

The xz-utils backdoor was a two-year social engineering campaign that ended in one malicious commit to a compression library. A few thousand lines, hidden inside a tarball, waiting in every major Linux distribution. The final merge was never audited by anyone who read the whole thing, because reading the whole thing was never feasible. The codebase was enormous. The reviewer was tired. The diff looked boring.

That is the supply chain attack in its pure form. It is not a breach. It is a substitution. Someone becomes the maintainer. Someone becomes the dependency. The tool you already trusted starts doing something you never asked for.

> The supply chain attack doesn't break in. It is invited. It becomes the dependency you chose.

Developers are the highest-value targets in software. The machine that builds and ships your code holds the tokens, the credentials, the source, the pipeline. Compromise the tool and you compromise everything downstream of it. This is why attackers keep coming back to the chain — not because the defenses are weak, but because the surface is enormous.

The attack surface of a modern toolchain is a geography of trust: transitive dependencies numbering in the thousands, IDE extensions that auto-update in the background, language toolchains that fetch on demand, CI images pulled at build time, build caches shared across machines, telemetry SDKs that ship your data somewhere else, and now agents that execute arbitrary code on your behalf. Each one is an entry point. The event-stream compromise injected a wallet stealer into a package with millions of weekly downloads. The polyfill.io domain sale put malware on more than a hundred thousand websites. ua-parser-js, eslint-scope, Codecov's uploader, SolarWinds Orion. The list is a tour of the trust you've delegated and the people who took it over.

> Every dependency is a maintainer who can be social-engineered. Every integration is a boundary where a substitution can hide.

Here is the uncomfortable part: the industry's response adds more of what caused the problem. SBOMs tell you what you have after you are compromised. Signature verification requires a key infrastructure that is itself a supply chain. Policy engines evaluate dependency risk using their own dependency tree. Pinning fights entropy. Every verifier needs its own verifier. Ken Thompson described the recursion in 1984: the compiler you trust may be lying, and no amount of checking escapes the base case. You cannot verify your way out of a trust problem.

> Verification below a certain size is theater. The verifier is just another dependency.

The base case of the recursion is code you can actually read. Not code that passed a scan. Not code certified by a service whose own code you have never seen. Code that fits in your head. That is the only terminal node in the trust graph — the only node that doesn't point at another node you must trust.

This is the argument for zot. zot stands for zero-overhead-tool, and the name is not a performance claim. It is a security claim. Overhead is attack surface. Every dependency is a delegation of trust you will never inspect. Every plugin is a third party running code inside your terminal. Every telemetry call is a channel that could exfiltrate what it collects. Every auto-update is a remote code execution primitive waiting for a compromise. zot keeps the overhead at zero on purpose: a small codebase, minimal dependencies, tool calls you can see before they run, no plugin marketplace shipping other people's code into your environment, no phone-home, local-first by default. The system is small enough to read — which is the only property that makes trust terminal.

> An agent you can't read is a supply chain attack you haven't met yet.

The coding agent era makes this urgent in a new way. Agents don't just run your code — they write it, edit it, execute commands with your credentials. The blast radius of a compromised agent is the blast radius of your entire trust graph. And an agent's supply chain includes every tool it invokes, every package it installs, every model it calls, every command it runs on your behalf. You cannot secure that surface by adding to it. You can only shrink it.

> An agent with a thousand dependencies is a thousand doors. An agent with ten is a door you can watch.

Some will say simplicity is a luxury, and that the ecosystem's richness is what makes tools powerful. But every extension is a boundary around code you will never read, and every boundary is a place a substitution can hide. The most powerful tool is not the one with the most integrations. It is the one whose behavior you can verify. The xz backdoor was eventually caught by a single human — a PostgreSQL developer named Andres Freund who noticed SSH was using half a second more than usual and started digging. One curious person, reading a diff, beat every automated system on the planet. That is the real defense: humans who can actually inspect the chain. You cannot buy that capability with a tool. You can only build tools small enough for it to apply.

> The best supply chain defense is a tool small enough that someone can actually read it.

Simplicity is not a constraint for zot. It is the product. Every line that isn't there is a line that can't be backdoored. Every dependency that isn't included is a maintainer who can't be compromised. Every integration that doesn't exist is a boundary that can't be substituted. The zero in zero-overhead is a security boundary.

---

**References:**

- Ken Thompson. [Reflections on Trusting Trust](https://www.cs.cmu.edu/~rdriley/487/papers/Thompson_1984_ReflectionsonTrustingTrust.pdf). Communications of the ACM, 1984. — The original statement of the trust recursion.
- [CVE-2024-3094 — xz-utils backdoor](https://en.wikipedia.org/wiki/XZ_Utils_backdoor). liblzma, 2024.
- Andres Freund. [backdoor in upstream xz/liblzma leading to ssh server compromise](https://www.openwall.com/lists/oss-security/2024/03/29/4). oss-security, March 2024.
- [Polyfill.io domain takeover](https://en.wikipedia.org/wiki/Polyfill_(programming_library)). 2024 — malware served from a trusted CDN to 100,000+ sites.
- [event-stream / flatmap-stream compromise](https://github.com/dominictarr/event-stream/issues/116). npm, 2018.
- [ua-parser-js hijack](https://github.com/faisalman/ua-parser-js/issues/536). npm, 2021.
- [eslint-scope postmortem](https://eslint.org/blog/2018/07/postmortem-for-malicious-package-publishes/). npm, 2018.
- Alex Birsan. [Dependency Confusion: How I Hacked Into Apple, Microsoft and Dozens of Other Companies](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610). 2021.
- Related: [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents) — Defense in depth against agents that escape.
- Related: [The Stack as a Design Abstraction](https://blog.hackspree.com/#the-stack-is-the-abstraction) — Each layer constrains the failure space.
