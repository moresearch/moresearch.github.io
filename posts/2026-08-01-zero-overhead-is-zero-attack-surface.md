---
title: Zero Overhead Is Zero Attack Surface
date: 2026-08-01
slug: zero-overhead-is-zero-attack-surface
summary: The xz backdoor was caught by one human reading a diff. That capability has to be designed for. zot — the zero-overhead coding agent — is built that way on purpose: four tools, eleven dependencies, no telemetry, no marketplace, extensions opt-in. Overhead is attack surface. Subtraction is the only supply chain defense that compounds.
tags: security, supply-chain, coding-agents, simplicity, developer-tools, zot
---

The xz-utils backdoor was a two-year social engineering campaign that ended in one malicious commit to a compression library. A few thousand lines hidden in a tarball, waiting inside every major Linux distribution. It was caught by a single human — PostgreSQL developer Andres Freund — who noticed SSH was taking half a second longer than usual and started reading a diff. One curious person beat every automated defense on the planet.

The supply chain attack is not a breach. It is a substitution. Someone becomes the maintainer. Someone becomes the dependency. The tool you already trusted starts doing something you never asked for.

> The supply chain attack doesn't break in. It is invited. It becomes the dependency you chose.

Developers are the highest-value targets in software. The machine that builds and ships your code holds the tokens, the credentials, the source, the pipeline. Compromise the tool and you compromise everything downstream. This is why attackers keep coming back to the chain — not because the defenses are weak, but because the surface is enormous.

The attack surface of a modern toolchain is a geography of trust: transitive dependencies numbering in the thousands, IDE extensions that auto-update in the background, language toolchains that fetch on demand, CI images and actions pulled at build time, telemetry SDKs that ship your data elsewhere — and now agents that execute arbitrary code with your credentials. Each one is an entry point. The event-stream compromise injected a wallet stealer into a package with millions of weekly downloads. The polyfill.io domain sale put malware on more than a hundred thousand websites. ua-parser-js, eslint-scope, Codecov's uploader, SolarWinds Orion, and the 2025 tj-actions GitHub Action compromise that leaked CI secrets from repositories running it. The list is a tour of the trust you've delegated and the people who took it over.

> Every dependency is a maintainer who can be social-engineered. Every integration is a boundary where a substitution can hide.

Here is the uncomfortable part: the industry's response adds more of what caused the problem. SBOMs tell you what you have after you are compromised. Signature verification requires a key infrastructure that is itself a supply chain. Policy engines evaluate dependency risk using their own dependency tree. Pinning fights entropy. Every verifier needs its own verifier. Ken Thompson described the recursion in 1984: the compiler you trust may be lying, and no amount of checking escapes the base case. You cannot verify your way out of a trust problem.

> Verification below a certain size is theater. The verifier is just another dependency.

The base case of the recursion is code you can actually read — the only terminal node in the trust graph, the only node that doesn't point at another node you must trust. This is the entire argument for zot, and it's why zot is built the way it is.

**What zot actually is.** zot — "zero-overhead-tool" — is a coding agent harness: "lightweight and written in Go," per its own README, ~300 stars in its first four months. It ships as one static binary. Its entire tool surface is four tools: read, write, edit, bash. Its module graph contains eleven dependencies. Read that again: a complete coding agent — interactive TUI, JSON-RPC subprocess mode, a Telegram bridge, built-in providers for thirty-plus model families from Anthropic and OpenAI to DeepSeek, Gemini, Copilot, and local Ollama — on eleven packages, every one a boring, battle-tested utility, and not one telemetry or analytics dependency in the graph. Most developer tools burn through more dependencies before they print their first error message.

> The entire tool surface of zot is four tools. The entire dependency graph is eleven packages. Overhead is a choice.

Every design decision in zot reads like a supply chain checklist:

- **No runtime, no node_modules.** One static binary. There is no dependency tree to hydrate, no package manager to pwn, no lockfile to poison. The install script verifies the release's SHA-256 against `checksums.txt` before the binary ever touches disk.
- **Extensions are opt-in, not a marketplace.** Extensions run as subprocesses speaking JSON-RPC over stdio — any language, no SDK required — and none are installed by default. There is no extension marketplace and no registry you silently subscribe to. The portable-agent format, zotfiles, explicitly lists "indexed registry distribution, installation, signatures, bundled executable extensions" as not yet implemented. The market hasn't been built. That is the point.
- **No telemetry in the graph.** The entire module graph contains no analytics or crash-reporting package. There is nothing to exfiltrate and no channel built in to carry it.
- **Credentials are treated as executable.** zot stores auth in `auth.json` with mode 0600 and documents it as executable configuration: anyone who can modify it can make zot run a program as you. So API keys pulled from commands run the program directly — no shell, no `!` prefixes, no string interpolation — and the output is cached in memory and never written to disk. The honesty is the security property.
- **Bounded blast radius by default.** A jail mode confines the agent's tools to the working directory and can be set as the default for new sessions. Skills declare their own allowed tools and bash permission patterns. Guardrail extensions can intercept tool calls mid-flight.

Each of these is a supply chain decision the industry would normally answer with more machinery. zot answers with less.

> An agent with a thousand dependencies is a thousand doors. An agent with four tools and eleven dependencies is a door you can watch.

The coding agent era makes this urgent in a new way. Agents don't just run your code — they write it, edit it, execute commands with your credentials. An agent's supply chain includes every tool it invokes, every package it installs, every model API it calls. The OWASP LLM Top 10 now lists supply chain vulnerabilities as one of the ten systemic risks of LLM applications. When the tool has agency, its supply chain is your supply chain. You cannot secure that surface by adding to it. You can only shrink it.

Some will say simplicity is a luxury — that the ecosystem's richness is what makes tools powerful. But every extension is a boundary around code you will never read, and every boundary is where a substitution hides. The most powerful tool is not the one with the most integrations; it is the one whose behavior you can verify. The xz backdoor was caught by one person reading a diff. That capability has to be designed for: tools small enough that a curious human can actually read them.

> The best supply chain defense is a tool small enough that someone can actually read it.

Simplicity is not a constraint for zot. It is the product. Every line that isn't there can't be backdoored. Every dependency that isn't included is a maintainer who can't be compromised. Every integration that doesn't exist is a boundary that can't be substituted. The zero in zero-overhead is a security boundary.

---

**References — supply chain:**

- Ken Thompson. [Reflections on Trusting Trust](https://www.cs.cmu.edu/~rdriley/487/papers/Thompson_1984_ReflectionsonTrustingTrust.pdf). Communications of the ACM, 1984. — The original statement of the trust recursion.
- [CVE-2024-3094 — xz-utils backdoor](https://en.wikipedia.org/wiki/XZ_Utils_backdoor). liblzma, 2024.
- Andres Freund. [backdoor in upstream xz/liblzma leading to ssh server compromise](https://www.openwall.com/lists/oss-security/2024/03/29/4). oss-security, March 2024.
- [Polyfill.io domain takeover](https://en.wikipedia.org/wiki/Polyfill_(programming_library)). 2024 — malware served from a trusted CDN to 100,000+ sites.
- [event-stream / flatmap-stream compromise](https://github.com/dominictarr/event-stream/issues/116). npm, 2018.
- [ua-parser-js hijack](https://github.com/faisalman/ua-parser-js/issues/536). npm, 2021.
- [eslint-scope postmortem](https://eslint.org/blog/2018/07/postmortem-for-malicious-package-publishes/). npm, 2018.
- CISA. [Codecov Supply Chain Compromise (AA21-131A)](https://www.cisa.gov/news-events/cybersecurity-advisories/aa21-131a). 2021.
- CISA. [Emergency Directive 21-01 — SolarWinds Orion (AA20-352A)](https://www.cisa.gov/news-events/cybersecurity-advisories/aa20-352a). 2020.
- [tj-actions/changed-files GitHub Action compromise (GHSA-jhrg-gw98-xc3h)](https://github.com/advisories/GHSA-jhrg-gw98-xc3h). 2025 — CI secrets exfiltrated from repositories using the action.
- Alex Birsan. [Dependency Confusion: How I Hacked Into Apple, Microsoft and Dozens of Other Companies](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610). 2021.
- OWASP. [Top 10 for Large Language Model Applications — LLM07 Supply Chain](https://genai.owasp.org/llm-top-10/). 2025.

**References — zot:**

- [zot — yet another coding agent harness, lightweight and written in Go](https://github.com/patriceckhart/zot). MIT, Go 1.25+.
- [zot README](https://github.com/patriceckhart/zot#readme) — one static binary, four tools (read, write, edit, bash), 30+ providers, extensions opt-in.
- [docs/extensions.md](https://github.com/patriceckhart/zot/blob/main/docs/extensions.md) — subprocess extensions over JSON-RPC, any language, none installed by default.
- [docs/zotfiles.md](https://github.com/patriceckhart/zot/blob/main/docs/zotfiles.md) — portable agents; registry distribution and signatures explicitly not implemented.
- [docs/rpc.md](https://github.com/patriceckhart/zot/blob/main/docs/rpc.md) — JSON-RPC subprocess mode with token auth via `ZOTCORE_RPC_TOKEN`.
- [docs/skills.md](https://github.com/patriceckhart/zot/blob/main/docs/skills.md) — skills declare allowed tools and bash permission patterns.
- [go.mod](https://github.com/patriceckhart/zot/blob/main/go.mod) — eleven dependencies total, no telemetry or analytics package in the graph.
- [install.sh](https://github.com/patriceckhart/zot/blob/main/install.sh) — downloads latest release and verifies SHA-256 against `checksums.txt`.
- Related: [Sandboxes Are Hard](https://blog.hackspree.com/#sandboxing-ai-agents) — Defense in depth against agents that escape.
- Related: [The Stack as a Design Abstraction](https://blog.hackspree.com/#the-stack-is-the-abstraction) — Each layer constrains the failure space.
