---
title: "It's Adam Back!"
date: 2026-08-03
slug: its-adam-back
summary: Five people have been named Satoshi Nakamoto. All five collapsed. The case that Adam Back is Satoshi is stronger than every one of them — and the strongest evidence is the thing that looks like it clears him: the emails, the denials, and a career that never paused. The citation, the signals, the resume, the denial. Five exhibits, all pointing one way.
tags: adam-back, satoshi-nakamoto, bitcoin, hashcash, cypherpunk, proof-of-work, cryptography, identity, pseudonymity, forensics, blockstream, cypherspace
---

Five people have been named Satoshi Nakamoto. All five times, the answer collapsed.

**Dorian Nakamoto.** Newsweek's March 2014 cover, "The Face Behind Bitcoin," on little more than a shared name. Satoshi's account answered the same week: "I am not Dorian Nakamoto."

**Hal Finney.** The first person ever to receive a bitcoin, the most technically credible suspect, the man who built RPOW. Finney denied it and died in 2014; the case died with him.

**Nick Szabo.** Bit gold, the closest thing to Bitcoin before Bitcoin, a writing style that felt right. He denied it; stylometry never closed it.

**Craig Wright.** He didn't deny it; he sued everyone who said he wasn't. In March 2024 the UK High Court ruled in *COPA v Wright* that he is not Satoshi and had forged his evidence. The only man to claim the identity in court lost it in court.

**Peter Todd.** HBO's *Money Electric* (2024) built its finale on a forum-post coincidence. He denied it; the theory is remembered as the moment the genre jumped the shark.

Five verdicts, five collapses. The pattern is the point: the field keeps looking at suspects instead of at citations. The whitepaper names its sources — and one of them fits everything. The inventor of Bitcoin's core mechanism. British. A distributed-systems PhD. A founding-generation cypherpunk, denying it with the consistency of a man who has rehearsed the answer for fifteen years.

The thesis: **Adam Back is the only suspect whose entire career was the design document for Bitcoin.** The strongest evidence against him is his own denial. The case, exhibit by exhibit.

## Exhibit A: the citation

The whitepaper's first technical move is proof-of-work, and proof-of-work is not Satoshi's invention. It is Adam Back's:

> "To implement a distributed timestamp server on a peer-to-peer basis, we will need to use a proof-of-work system similar to Adam Back's Hashcash [6], rather than newspaper or Usenet posts."

The mechanism is described in Satoshi's own words, and the words are hashcash's:

> "The proof-of-work involves scanning for a value that when hashed, such as with SHA-256, the hash begins with a number of zero bits."

Back invented hashcash in 1997 as an anti-spam device: force the sender of an email to burn CPU on a partial hash collision. The 2002 papers are where it becomes Bitcoin's economics — "Hashcash - A Denial of Service Counter-Measure" and, the title that should be on a plaque, "Hashcash - Amortizable Publicly Auditable Cost-Functions." Publicly auditable cost functions *is* mining. Bitcoin's difficulty adjustment, its halving, its "one-CPU-one-vote" consensus — all of it is the hashcash paper extended from spam defense to money.

Back's own homepage says it plainly: hashcash is "the mining function in bitcoin," and under "Bitcoin Related," "how bitcoins uses hashcash fractional difficulty, automated inflation control."

Read that second line again. The inventor of the mining function publishing, on his own page, a note about how Bitcoin uses its fractional difficulty and inflation control. Every other suspect had to learn proof-of-work. Exactly one person in the world did not.

## Exhibit B: the emails

In August 2008 — weeks before the whitepaper went public — someone using the name Satoshi Nakamoto emailed Adam Back. Five emails from the correspondence are now in the public record, entered as exhibits in *COPA v Wright*. Satoshi referenced Back's hashcash paper and said he was preparing to release a whitepaper.

For the defense, this is the whole case in one exhibit: the inventor of hashcash would not email himself about hashcash in 2008.

For the prosecution, read it as the most careful piece of stagecraft in the case. A real Satoshi needed hashcash to be the foundation of the whitepaper — the mechanism, the citation, the legitimacy. He also needed, at some future point, to be able to say: *look, I contacted the inventor before I published; I'm just a builder on his work.* The August 2008 email is that alibi, insurance taken out before the crime.

And look at what Back did with the email. He kept it, and built his company's brand around it: Blockstream produced a commercial dramatizing a young Adam Back reading the historic email from Satoshi. The man who, if he were anyone else, would be the greatest living witness to Bitcoin's creation made the email a marketing prop. People do not make their alibis into commercials.

## Exhibit C: the signals

Everything about Satoshi is UK-shaped. The genesis block — January 3, 2009 — embeds a UK newspaper headline: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks." The first block of the first cryptocurrency is timestamped with a London front page, chosen on the day it was printed.

Satoshi's posting hours cluster in the British evening. His writing mixes British spelling and idioms with American expressions — the profile of a Brit steeped in US tech culture. Adam Back is British, with a University of Exeter PhD (1995) in distributed systems, on the Cypherpunks mailing list since the mid-1990s, running the UK crypto-politics pages on his own site. He has maintained since 1995 exactly the profile the signals describe.

## Exhibit D: the resume is the design document

Open [cypherspace.org/adam](http://cypherspace.org/adam/) — the page this post is named after — and read it as a requirements list:

- **hashcash** — proof-of-work. ✓ (Bitcoin's consensus)
- **credlib** — "credential library with chaum and brands ecash/credentials" — anonymous e-cash. ✓ (Bitcoin's premise)
- **The Eternity Service** — a censorship-proof document store, *Phrack*, 1997. ✓ (Bitcoin's promise)
- **Cypherspace** — "distributed data haven project." ✓ (the same promise, bigger)
- **Cebolla** — IP anonymity. ✓ (Bitcoin's network layer)
- **The Crypto Hacks page** — breaking Netscape's SSL challenges (the second in 32 hours), extracting the NSA's keys from Microsoft's CAPI. ✓ (the conviction that the old money systems are broken)

The man who spent the 1990s breaking the security of the financial internet, exposing government backdoors, and writing the theory of publicly auditable cost functions is the only candidate who does not need to be explained. Everyone else on the suspect list was a scholar of one or two of these domains. Back authored all of them.

Then note what he did next. Between hashcash (1997) and 2014, Back published no major protocol design. In 2014 — the year he founded Blockstream — he co-authored "Enabling Blockchain Innovations with Pegged Sidechains." The first new Back protocol in seventeen years arrived exactly when Bitcoin needed its next layer, from the CEO of a company whose entire business is Bitcoin. The two 17-year gaps are the same gap.

## Exhibit E: the candidate pool

Bitcoin required simultaneous mastery of proof-of-work, public-key cryptography, P2P networking, incentive design, monetary economics, and C++. Every named suspect was strong in three or four: Finney had crypto, code, and RPOW but no monetary economics; Szabo had economics and bit gold but no proof-of-work innovation and no C++ track record; Wright had none of it, which is why the court shredded him.

Back is the one candidate strong in all six, at inventor level in one of them. And that person's next act was to build a company on the protocol, which is the most Satoshi thing anyone has ever done.

## The denial

Now the exhibit that looks like the whole answer: Adam Back has said, repeatedly, in public, for fifteen years, that he is not Satoshi.

A real Satoshi must deny. **The asset is destroyed.** Bitcoin's value is not the million-plus coins in the Patoshi wallets; it is the neutrality. The protocol became the industry's because it had no owner, no spokesperson, no CEO. "Satoshi is Adam Back, CEO of Blockstream" converts the most valuable neutral protocol on earth into a company's project, overnight. The mystery is not a bug in Satoshi's plan; it is the plan.

**The exposure is existential.** A pseudonymous creator of a trillion-dollar asset, discovered, faces every regulator, every litigant, every tax authority, every hacker. The Wright trial showed what happens to people who merely *claim* the identity; what happens to someone who actually has it is the reason the identity was built to be unclaimable.

**The performance is uniform.** Back has denied it in the same flat, unvarying way for fifteen years — no outrage, no legal threats, no "let me prove it." He denies it the way a man denies a rumor he has decided to enjoy.

And here is the part that keeps the case alive. The tell is not the denial; it is what Back does around it. His homepage caches Satoshi's deleted Wikipedia article — under the note "cache of Satoshi Nakamoto's wikipedia page which the editors deleted??", the double question mark doing visible work. His company dramatizes the moment he received Satoshi's email. He has played the "I'm not Satoshi" line as a bit for a decade. A man falsely accused, whose company would benefit from the rumor dying, keeps a candle lit on it. Either he is innocent and enjoys the free publicity — or he is guilty and enjoys the only freedom he has, which is telling the truth in the way nobody believes it.

## The case for the defense

Honesty requires the strongest counter-argument, and it is the emails again, from the other side. The natural reading of the August 2008 correspondence is that a real inventor contacted the real author of his foundational reference, cited him properly, and the two men exchanged professional notes. That is the ordinary world's explanation, and it requires nothing special.

The whitepaper over-credits — Dai, Szabo, Finney, Haber and Stornetta, Merkle — unusually generous prior-art for a man about to change the world. Inventors who build on published work cite it; self-citing inventors over-cite. Both readings survive. Stylometry has matched Satoshi to half the field and convicted no one. And the deep counter: the deception would be enormous — fifteen years, sustained through his own company's marketing, in a field that scrutinizes him constantly. That is either impossible, or it is exactly the discipline of a man who already kept the biggest secret in technology for two years while everyone in his field discussed it in front of him.

## Verdict

The case cannot be closed, and that is the strongest evidence of all. Satoshi engineered the identity to survive forensic pressure: anonymous email, anonymous registration, no PGP key, pattern-averaging posting hours, and a silence chosen at exactly the moment Bitcoin stopped being a hobby. The protocol was designed so that its creator could never be proven — and so that anyone could build on it.

You cannot prove Adam Back is Satoshi. Neither can anyone else. But every other suspect needs a theory to explain how they did it. Back is the only one who needs a theory to explain how he did not.

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

**Related on this site:**

- [The Decentralized AI-Agent Experience](https://blog.hackspree.com/#freenet) — the cypherpunk lineage: Eternity Service → Freenet → the distributed future.
- [Agentic Markets: Mechanism Design and Network Economics](https://blog.hackspree.com/#mechanism-design-and-network-economics-for-agentic-markets) — incentive design is Bitcoin's real invention; the hardest problem Satoshi had to solve.
