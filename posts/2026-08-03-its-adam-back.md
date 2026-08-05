---
title: It's Adam Back!
date: 2026-08-03
slug: its-adam-back
summary: "Five people have been named Satoshi Nakamoto. All five collapsed. The case that Adam Back is Satoshi is stronger than every one of them — and the strongest evidence is the thing that looks like it clears him: the emails, the denials, and a career that never paused. The citation, the signals, the resume, the denial. Five exhibits, all pointing one way."
tags: adam-back, satoshi-nakamoto, bitcoin, hashcash, cypherpunk, proof-of-work, cryptography, identity, pseudonymity, forensics, blockstream, cypherspace
---

Five people have been named Satoshi Nakamoto. All five times, the answer collapsed.

**Dorian Nakamoto.** Newsweek's 2014 cover, on little more than a shared name. Satoshi's account answered: "I am not Dorian Nakamoto."

**Hal Finney.** The first person to receive a bitcoin, the man who built RPOW. He denied it and died in 2014; the case died with him.

**Nick Szabo.** Bit gold, the closest thing to Bitcoin before Bitcoin. He denied it; stylometry never closed it.

**Craig Wright.** He sued everyone who said he wasn't. In March 2024 the UK High Court ruled in *COPA v Wright* that he is not Satoshi and had forged his evidence.

**Peter Todd.** HBO's *Money Electric* (2024) built its finale on a forum-post coincidence; he denied it, and the theory jumped the shark.

Five verdicts, five collapses. The pattern is the point: the field keeps looking at suspects instead of at citations. The whitepaper names its sources — and one of them fits everything: the inventor of Bitcoin's core mechanism, British, a distributed-systems PhD, a founding-generation cypherpunk, denying it with the consistency of a man who has rehearsed the answer for fifteen years.

The thesis: **Adam Back is the only suspect whose entire career was the design document for Bitcoin — and the strongest evidence against him is his own denial.**

## Exhibit A: the citation

The whitepaper's first technical move is proof-of-work, and proof-of-work is not Satoshi's invention. It is Adam Back's:

> "To implement a distributed timestamp server on a peer-to-peer basis, we will need to use a proof-of-work system similar to Adam Back's Hashcash [6]."

Back invented hashcash in 1997 as an anti-spam device: force the sender of an email to burn CPU on a partial hash collision. The 2002 papers turn it into Bitcoin's economics — "Hashcash - A Denial of Service Counter-Measure" and, the title that should be on a plaque, "Hashcash - Amortizable Publicly Auditable Cost-Functions." Publicly auditable cost functions *is* mining. Bitcoin's difficulty adjustment, halving, and "one-CPU-one-vote" consensus are the hashcash paper extended from spam defense to money.

Back's own homepage says it plainly: hashcash is "the mining function in bitcoin," and under "Bitcoin Related," "how bitcoins uses hashcash fractional difficulty, automated inflation control." The inventor of the mining function publishing a note on how Bitcoin uses its fractional difficulty — every other suspect had to learn proof-of-work. Exactly one person in the world did not.

## Exhibit B: the emails

In August 2008, weeks before the whitepaper, someone using the name Satoshi Nakamoto emailed Adam Back. Five emails from the correspondence are now in the public record, entered in *COPA v Wright*; Satoshi referenced hashcash and said he was preparing to release a whitepaper.

For the defense, this is the whole case in one exhibit: the inventor of hashcash would not email himself about hashcash.

For the prosecution, read it as stagecraft. A real Satoshi needed hashcash as the whitepaper's foundation — and needed, at some future point, to be able to say: *look, I contacted the inventor before I published.* The August 2008 email is that alibi, insurance taken out before the crime.

And look at what Back did with the email. He kept it, and built his company's brand around it: Blockstream produced a commercial dramatizing a young Adam Back reading the historic email from Satoshi. The man who, if he were anyone else, would be the greatest living witness to Bitcoin's creation made the email a marketing prop. People do not make their alibis into commercials.

## Exhibit C: the signals

Everything about Satoshi is UK-shaped. The genesis block — January 3, 2009 — embeds a UK newspaper headline: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks." His posting hours cluster in the British evening; his writing mixes British spelling and idioms with American expressions — the profile of a Brit steeped in US tech. Adam Back is British, a University of Exeter distributed-systems PhD (1995), on the Cypherpunks list since the mid-1990s. He has maintained since 1995 exactly the profile the signals describe.

## Exhibit D: the resume is the design document

Open [cypherspace.org/adam](http://cypherspace.org/adam/) — the page this post is named after — and read it as a requirements list:

- **hashcash** — proof-of-work. ✓ (Bitcoin's consensus)
- **credlib** — "chaum and brands ecash/credentials" — anonymous e-cash. ✓ (Bitcoin's premise)
- **The Eternity Service** — a censorship-proof document store, *Phrack*, 1997. ✓ (Bitcoin's promise)
- **Cebolla** — IP anonymity. ✓ (Bitcoin's network layer)
- **The Crypto Hacks page** — breaking Netscape's SSL challenges (the second in 32 hours), extracting NSA keys from Microsoft's CAPI. ✓ (the conviction that the old money systems are broken)

The man who spent the 1990s breaking the financial internet, exposing government backdoors, and writing the theory of publicly auditable cost functions is the only candidate who does not need to be explained. Then note what he did next: between hashcash (1997) and 2014, no major protocol design. In 2014 — the year he founded Blockstream — he co-authored the sidechains paper. The first new Back protocol in seventeen years arrived exactly when Bitcoin needed its next layer, from the CEO of a company whose entire business is Bitcoin. The two 17-year gaps are the same gap.

## Exhibit E: the candidate pool

Bitcoin required simultaneous mastery of proof-of-work, public-key crypto, P2P networking, incentive design, monetary economics, and C++. Every named suspect was strong in three or four: Finney had crypto and RPOW but no monetary economics; Szabo had economics and bit gold but no proof-of-work innovation and no C++; Wright had none of it. Back is the one candidate strong in all six, at inventor level in one — and his next act was to build a company on the protocol, which is the most Satoshi thing anyone has ever done.

## The denial

Adam Back has said, repeatedly, for fifteen years, that he is not Satoshi. A real Satoshi must deny. **The asset is destroyed:** Bitcoin's value is not the coins; it is the neutrality. "Satoshi is Adam Back, CEO of Blockstream" converts the most valuable neutral protocol on earth into a company's project overnight. **The exposure is existential:** a pseudonymous creator of a trillion-dollar asset faces every regulator, litigant, and hacker; the Wright trial showed what happens to people who merely claim the identity. **The performance is uniform:** Back denies it in the same flat, unvarying way — no outrage, no legal threats, no "let me prove it."

And here is the part that keeps the case alive. The tell is not the denial; it is what Back does around it. His homepage caches Satoshi's deleted Wikipedia article, under the note "cache of Satoshi Nakamoto's wikipedia page which the editors deleted??", the double question mark doing visible work. His company dramatizes the moment he received Satoshi's email. He has played the "I'm not Satoshi" line as a bit for a decade. A man falsely accused, whose company would benefit from the rumor dying, keeps a candle lit on it.

## The case for the defense

The strongest counter-argument is the emails, from the other side: a real inventor contacted the real author of his foundational reference, cited him properly, and the two exchanged professional notes. The whitepaper over-credits — Dai, Szabo, Finney, Haber and Stornetta, Merkle — unusually generous prior art. Both readings survive; stylometry has convicted no one. And the deep counter: the deception would be enormous — fifteen years, sustained through his own company's marketing. That is either impossible, or it is exactly the discipline of a man who already kept the biggest secret in technology for two years while everyone in his field discussed it in front of him.

## Verdict

The case cannot be closed, and that is the strongest evidence of all. Satoshi engineered the identity to survive forensic pressure — anonymous email, no PGP key, pattern-averaging hours, a silence chosen at the moment Bitcoin stopped being a hobby. The protocol was designed so its creator could never be proven, and so anyone could build on it.

You cannot prove Adam Back is Satoshi. Neither can anyone else. But every other suspect needs a theory of how they did it; Back is the only one who needs a theory of how he did not.

> The citation. The emails. The signals. The resume. The denial. Five exhibits, none conclusive, all pointing one way. It's Adam Back — and the denial is part of the evidence, not the answer.

---

**References:**

- [Adam Back's home page — cypherspace.org/adam](http://cypherspace.org/adam/) — hashcash ("used as the mining function in bitcoin"), the Satoshi Wikipedia cache, the publications list, the Crypto Hacks page, the Exeter PhD.
- [Satoshi Nakamoto. *Bitcoin: A Peer-to-Peer Electronic Cash System*](https://bitcoin.org/bitcoin.pdf). — "a proof-of-work system similar to Adam Back's Hashcash [6]"; "proof-of-work is essentially one-CPU-one-vote."
- Adam Back. [Hashcash - A Denial of Service Counter-Measure](http://www.hashcash.org/papers/hashcash.pdf), 2002; and [Hashcash - Amortizable Publicly Auditable Cost-Functions](http://www.hashcash.org/papers/amortizable.pdf), 2002.
- Adam Back. [The Eternity Service](https://www.phrack.org/issues/51/6.html), *Phrack* 7(51), Sep 1997.
- Adam Back et al. [Enabling Blockchain Innovations with Pegged Sidechains](https://blockstream.com/sidechains.pdf), Oct 2014.
- [COPA v Wright — UK High Court judgment](https://www.judiciary.uk/judgments/copa-v-wright/), March 14, 2024. — Wright is not Satoshi; the Satoshi–Back email correspondence entered into the record.
- Leah McGrath Goodman. [The Face Behind Bitcoin](https://www.newsweek.com/2014/03/14/face-behind-bitcoin-247957.html), *Newsweek*, March 2014.
- [Money Electric: The Bitcoin Mystery](https://www.hbo.com/money-electric-the-bitcoin-mystery), HBO, October 2024. — the Peter Todd conclusion.
- [Bitcoin genesis block](https://en.bitcoin.it/wiki/Genesis_block) — the embedded Times headline, January 3, 2009.
- Related: [The Decentralized AI-Agent Experience](https://blog.hackspree.com/#freenet) — the cypherpunk lineage: Eternity Service → Freenet → the distributed future.
- Related: [Agentic Markets: Mechanism Design and Network Economics](https://blog.hackspree.com/#mechanism-design-and-network-economics-for-agentic-markets) — incentive design is Bitcoin's real invention; the hardest problem Satoshi had to solve.
