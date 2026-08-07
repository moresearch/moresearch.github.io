---
title: Automatism and Vibe Coding
date: 2026-08-07
slug: automatism-and-vibe-coding
summary: "André Breton defined Surrealism as 'psychic automatism' — the dictation of thought in the absence of any control exercised by reason. A coding agent is the most literal heir of that idea ever built: it generates code by sampling a latent distribution, without deliberation — which is what the culture calls vibe coding. This post follows the automatist tradition — from Breton and Soupault's automatic writing through Masson's automatic drawing and Pollock's action painting — to the practice problem at the heart of vibe coding: what you do with the automatic output is the whole craft."
tags: agents, agentic-software-engineering, vibe-coding, automatism, surrealism, automatic-writing, jackson-pollock, andre-breton, art-history, curation, harness-engineering
---

In 1924, André Breton published the *Manifesto of Surrealism* and defined the movement with a single phrase:

> SURREALISM, n. Pure psychic automatism, by which one intends to express verbally, in writing or by any other method, the real functioning of the mind. Dictation of thought, in the absence of any control exercised by reason, and beyond any aesthetic or moral preoccupation.

Breton meant writing without deliberation — letting the mind speak in its own currents, then seeing what had been said. The Surrealists called it automatic writing, and they treated the results as messages from a deeper agency: the unconscious, the irrational, the "real functioning of the mind."

A century later, the same idea has become the most consequential mode of software production on earth. A coding agent is an automatic writer: it generates code by sampling from a latent distribution of everything it has read, without deliberation, without a plan in the way a programmer plans — and certainly "beyond any aesthetic or moral preoccupation." The culture already has a name for the purest form of this, and it is an automatist's name: vibe coding — describe what you want, accept what comes back, don't read it too closely. Automatic writing, automatic code.

**Agentic coding is automatism industrialized, and vibe coding is automatism at its most literal.** Everything strange about working with agents — the wonder, the unreliability, the editing, the anxiety — is the strangeness the Surrealists met a hundred years ago. Their answer to it is the most useful theory of vibe coding I know.

![Jackson Pollock painting on glass, photographed by Hans Namuth (1950). The automatic method, made physical — produce first, then look.](/images/pollock-namuth.jpg)

## The first automatic writers

Breton and Philippe Soupault wrote *The Magnetic Fields* in 1919, in a few weeks of trance-like haste, before the word "Surrealism" existed. The result read like nothing before it: images that did not follow from each other, syntax that kept its shape while sense dissolved. They were not trying to say something. They were trying to let the dictation happen.

The practice spread to drawing. Masson let his hand wander without an image in mind, then looked at what emerged; Miró built canvases from the same automatist improvisations; Ernst pressed frottage and scraped grattage to let the material itself suggest the image. The method was always the same: **produce automatically, then look.** The looking was not optional. It was the second half of the method — the half that made it art instead of noise.

## The automatic, controlled

Jackson Pollock belongs in this lineage. In the late 1930s, in the drawings he made during his Jungian analysis, he practiced exactly this method: letting the line run without intent, then reading the drawing for what it revealed. A decade later, the method scaled from the page to the floor, from the line to the pour. *No. 5, 1948* — the painting that would later become the most expensive ever sold at auction — is the pure product of that method: four by eight feet of enamel and paint poured from above, the automatist gesture at full scale.

![No. 5, 1948 — Jackson Pollock (1948). The automatic method made material: poured paint, fully controlled. (Low-resolution reproduction; the original is in a private collection.)](/images/no-5-1948.jpg)

> When I am in my painting, I'm not aware of what I'm doing. It is only after a sort of "get acquainted" period that I see what I have been about... the painting has a life of its own. I try to let it come through.

That is an automatist's account of his own method. But Pollock added the half the lay view always forgets:

> It is only when I lose contact with the painting that the result is a mess.

And when asked directly whether the pours were accidents: "I deny the accident. I can control the flow of paint."

The automatic was the source; control was the craft. The two were not opposites. The pour only worked while the body that produced it stayed in the loop — acting, stepping back, seeing, acting again. The same gesture, without the control, was just paint on canvas.

## What an agent is

A coding agent is the most literal automatic writer ever built, because its generation is genuinely opaque to deliberation. When the model emits code, no part of it reasons the way a programmer reasons. It samples from a learned distribution of tokens — a statistical unconscious containing the accumulated text of the world. The output issues "in the absence of any control exercised by reason," exactly as Breton prescribed, and with no aesthetic or moral preoccupation whatsoever.

The "dictation of thought" is literal too, and the thought being dictated is not the engineer's. It is the training distribution's — the aggregate of every repository, every discussion, every pattern the model absorbed. The agent is an unconscious at scale, and it writes.

This is why vibe coding feels the way automatic writing felt. The output is real, sometimes astonishing, and consistently unreliable in ways that defeat introspection. The model cannot tell you why it wrote what it wrote. Neither could Masson's hand. The interesting part is what you do next.

## The editor's craft

The Surrealists discovered within a decade that pure automatism produces mostly mediocrity. *The Magnetic Fields* was astonishing; the practice of magnetic fields, repeated, was not. What saved the movement was editing: Breton curated relentlessly — cut, revised, chose which dictations to keep — and the painters composed and corrected after the automatic pass. The art lived in the relation between the automatic and the deliberate.

The same discovery is being made, at industrial scale, about agents. The pure-automatist version of the practice — accept the output, skip the looking — is vibe coding in its purest form, and it is *The Magnetic Fields* repeated: astonishing once, mediocre as a method. The prompt is the séance — the conditions under which the dictation happens. The harness is the editing table: the tests that cut what fails, the evals that rank what survives, the sandbox that keeps the automatic from doing damage while it works. The engineer's craft has moved entirely into that relation. **You cannot make the automatic better by staring at it. You make it better by changing its conditions and editing its output.**

Automatic output also has no single author and no single intention — a dense web of small choices that no one deliberately made. Coherence has to be imposed from outside, by the spec, the contract, the harness. The spec is the architect; the agents are the automatic; the editor is you.

## Judgment without introspection

The hardest question automatism poses is: how do you know the automatic output is good? Neither the artist nor the model can answer by consulting intention — there was none. The Surrealists answered with an external test: the image — the sudden resonance of distant realities meeting. The image could not be explained, only recognized.

Surrealism's own history contains the warning, painted in the movement's other half. Magritte was the counterweight inside the group: where Breton sought the unconscious through unpremeditated marks, Magritte built each image with the deliberation of a cabinetmaker, and the tension between the two poles ran through the movement's entire history. His most famous painting is a manifesto against mistaking representation for reality:

![The Treachery of Images — René Magritte (1928–29). "Ceci n'est pas une pipe": the image is not the thing, and neither is the word. (Low-resolution reproduction; the original is at LACMA.)](/images/magritte-treachery-of-images.jpg)

*Ceci n'est pas une pipe.* Of course it is not a pipe. It is paint arranged to look like a pipe. The treachery is that it works anyway: you read it as a pipe and only the sentence beneath it — itself another representation — breaks the illusion. Magritte's whole point is that resemblance is precisely what makes the lie possible, and that representation is not identity.

This is the trap vibe coding sets. The agent's output is not engineering. It is a representation of engineering — tokens arranged to look like code, sampled from everything code has ever looked like. It resembles working software the way the painting resembles a pipe. Accept the resemblance without testing and you are doing what Magritte's sentence stops you from doing: taking the picture for the thing. The eval is the act of trying to smoke the pipe. The test is the treachery broken.

Code has the advantage that its external test is sharper: does it run, does it pass, does the system behave? The eval and the test suite are the criterion the automatist tradition never had. But a residue remains that evals cannot reach — the judgment of whether the automatic solution has the right shape, whether the surprise is a genuine find or a seductive artifact. For that you need the record of the making: the plan, the action stream, the replay — the trace of what the agent did, not just what it produced. You judge the dictation by its process as much as by its product.

## The danger of losing contact

The Surrealists' history is a caution about what happens when the relation degrades. Automatism hardened into a style; Breton spent the 1930s excommunicating members over how much control the method allowed. When the automatic is trusted wholesale, the output becomes mannerist — fluent, plausible, empty. When it is distrusted wholesale, the source dries up.

The equivalent failure in agentic coding is delegation without judgment — vibe coding practiced without the looking half. Vibe coding is fine for throwaway scripts, the automatist sketch, the one-session *Magnetic Fields*; as a permanent mode it is the loss of contact. The engineer who accepts whatever the agent produces stops exercising judgment, and judgment, once unused, decays. The trial covered on this blog found exactly this — programmers who delegated their learning to an AI scored 17% lower on what they were supposed to learn. Pollock's rule is the law of the whole practice: *the result is a mess when you lose contact with the painting.* The automatic needs the contact.

## Conclusion

Breton wanted to hear the mind speak without the interference of reason. A century later, we have built a machine that does it — at the scale of whole codebases. Automatic writing has become automatic code — vibe coding, in the culture's word — and the discipline that makes it work is the discipline the Surrealists spent a decade learning: produce automatically, then look; control the conditions, edit the output, judge by external criteria, keep the surprises worth keeping, and never lose contact.

The dictation is automatic. The engineering is not. That is the whole art.

---

**References:**

- André Breton. [Manifesto of Surrealism](https://en.wikipedia.org/wiki/Surrealist_automatism), 1924. — the definition of "psychic automatism... in the absence of any control exercised by reason."
- [Automatism — Art Term](https://www.tate.org.uk/art/art-terms/a/automatism), Tate. — the canonical definition: the term borrowed from physiology for bodily movements that are not consciously controlled (breathing, sleepwalking), via Freud's free association to Breton's dictation of thought.
- André Breton & Philippe Soupault. *The Magnetic Fields* (*Les Champs Magnétiques*), 1919. — the first automatic text.
- André Masson, Joan Miró, Max Ernst — automatic drawing, frottage, grattage, and the method of producing automatically then looking.
- [Astonishing Examples of Automatic Drawing](https://blog.artsper.com/en/a-closer-look/art-movements-en/automatic-drawing/), Artsper Magazine. — the drawing-side companion to the automatic writings: the practice of bypassing conscious control, with examples.
- Andrej Karpathy. [Vibe coding](https://en.wikipedia.org/wiki/Vibe_coding), February 2025. — the coinage of the practice this post reads as automatism's purest form: describe intent in natural language, accept the output without reading it closely.
- René Magritte. [The Treachery of Images](https://en.wikipedia.org/wiki/The_Treachery_of_Images) (*La trahison des images*), 1928–29. — "Ceci n'est pas une pipe": the deliberate pole of Surrealism and the standing warning that representation is not identity. Los Angeles County Museum of Art. Embedded above; low-resolution fair-use reproduction via Wikipedia, upscaled for display. (See also Michel Foucault, *This Is Not a Pipe*, 1973.)
- Jackson Pollock. [My Painting](https://en.wikipedia.org/wiki/Jackson_Pollock). *Possibilities I* (1947–48), ed. Harold Rosenberg. — the get-acquainted period; "the painting has a life of its own"; "it is only when I lose contact with the painting that the result is a mess."
- Pollock, "I deny the accident. I can control the flow of paint" — interview with William Wright, 1950 (recorded for radio, aired posthumously).
- [Hans Namuth photograph of Pollock painting on glass, 1950](https://commons.wikimedia.org/wiki/File:Jackson_Pollock_by_Hans_Namuth.jpg) — public domain; embedded at the top of this post.
- Jackson Pollock, [No. 5, 1948](https://en.wikipedia.org/wiki/No._5,_1948), 1948 — oil and enamel on fiberboard, 121.9 × 243.8 cm, private collection. Embedded above; low-resolution fair-use reproduction via Wikipedia, upscaled for display.
- Related: [Finding David in the Marble](https://blog.hackspree.com/#finding-david-in-the-marble) — the strike-assess loop, and the same act-then-look rhythm.
- Related: [You Don't Learn What You Delegate](https://blog.hackspree.com/#ai-impacts-skill-formation) — what delegation costs the actor's judgment.
- Related: [Harness Engineering (Martin Fowler)](https://blog.hackspree.com/#harness-engineering-fowler) — the harness as the editing table.
- Related: [Agent Harnesses Need Tasks That Fight Back](https://blog.hackspree.com/#agent-harnesses-need-tasks-that-fight-back) — the automatic needs resistance to mean anything.
- Related: [Codebases in the Era of Agentic Software Engineering](https://blog.hackspree.com/#codebases-in-the-era-of-agentic-software-engineering) — the spec is the architect, the agents are the builders.
- Related: [Conceptual Integrity and the One-Mind Rule](https://blog.hackspree.com/#brooks-design-conceptual-integrity) — coherence imposed on multi-authored output from outside.
- Related: [OpenWorker: Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer) — the record of the making as a first-class artifact.
- Related: [LLMs Can't Jump](https://blog.hackspree.com/#llms-cant-jump) — what the automatic cannot do: the leap from experience to axioms.
- Related: [Sandboxing AI Agents](https://blog.hackspree.com/#sandboxing-ai-agents) — securing the vibe coding stack, in Replit's phrase.
- Related: [Harness Engineering Best Practices for AI Agents](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) — tests are better than vibes: the looking half, made explicit.
