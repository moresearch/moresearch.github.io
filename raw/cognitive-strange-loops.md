---
title: "On the Road to Cognitive Strange Loops"
date: 2026-09-10
slug: cognitive-strange-loops
summary: "A strange loop is a hierarchy with no well-founded top: a lower level encodes a description of a higher one using only the system's own rules, and the system then acts on that description. This post takes the road there — Escher's hands, Bach's endlessly rising canon, Gödel's arithmetical self-accusation, Turing's universal machine, von Neumann's description/constructor split, Conway's Life, Thompson's trusting-trust compiler, and the coding agent that rewrites its own tools and plugs them back in. Along the way it separates feedback, recursion, self-reference and the strange loop into four different things; shows that every loop rests on a level it refuses to tangle; and follows the pattern the instances share — the power to describe yourself and the limit you cannot see past are the same fact. Whether the last loop on the road, the one that is a self, has actually been reached is the question the essay leaves open."
tags: strange-loops, self-reference, hofstadter, godel, turing, conway, consciousness, agents, self-improvement, bootstrapping, trusting-trust, fixed-points, recursion, bach, escher, rna-world
---

Look at M.C. Escher's [_Drawing Hands_](https://en.wikipedia.org/wiki/Drawing_Hands) (1948): two hands, each holding a pencil, each drawing the cuff and wrist of the other, neither arriving first. Listen to a [Shepard tone](https://en.wikipedia.org/wiki/Shepard_tone): a scale built from octave-spaced sine waves under a fixed bell-shaped envelope, so every step sounds higher than the last and the sequence still returns to the pitch it began on. Watch a bar of Bach modulate up a whole tone, then do it again, six times, and land back in C.

These are the same shape wearing three costumes. Douglas Hofstadter named it in [_Gödel, Escher, Bach_](https://en.wikipedia.org/wiki/G%C3%B6del,_Escher,_Bach) (Basic Books, 1979): a **strange loop** is what you get when you move upward through the levels of a hierarchy — following the hierarchy's own rules — and find yourself back at the level you started from. It is not a circle. It is a *tangled* hierarchy, one in which the level that is supposed to be in charge turns out to be written by the level it is supposed to govern.

The pattern shows up wherever a system becomes rich enough to contain a description of itself: in arithmetic, in music, in the machinery of a cell, in a compiler, and now in an agent that edits the harness it runs on. So treat the rest of this as a road report. The route is known: logic and music first, then biology, then code, then agents. The destination is the **cognitive** strange loop — the loop that is not merely described but *experienced*, the one that is a self. This post is about how far along that road each stage actually takes us, and about why the last stretch is still fog.

## What makes a loop strange

The word "loop" does too much work. A thermostat is a loop. A `for` statement is a loop. A population cycle is a loop. None of them is strange, because strangeness requires a specific kind of self-reference: the system must contain a *description of itself* and act on that description in a way that changes the system.

> **Definition.** A strange loop is a hierarchy with no well-founded top. A lower level encodes a description of a higher level, using only the system's own rules; the higher level then operates on that description; and there is no level *outside* the loop from which the system can survey the whole thing at once. The ordering of levels is real — it just is not the ordering the system's own rules describe.

![A strict hierarchy of four levels with arrows pointing only downward, beside a strange loop of the same four levels with a return arrow from the bottom level back into the top one](images/strange-loop-vs-hierarchy.svg)

*Figure 1. The difference the definition is drawing. In a strict hierarchy, information moves one level down and the boundary holds. In a strange loop, the bottom level writes the top level, and the distinction between the system and what the system operates on dissolves.*

Four pieces do the work, and it pays to name them separately:

- **A level.** Levels are defined by direction, not by size: a higher level *represents* or *governs* a lower one — rules over moves, program over data, observer over observed. Without this ordering there is nothing to tangle.
- **A licensed ascent.** The lower level must reach the higher one *through the system's own rules*. Gödel did not add an axiom; he showed that ordinary arithmetic could already talk about itself once its symbols were numbered. A quine does not consult the filesystem; it uses string formatting. An agent does not hack its own binary; it uses its own tool-writing.
- **Representation, not just influence.** This is the line between a strange loop and ordinary feedback. The return path must carry a *description* of the upper level, not merely affect it. A thermostat's sensor does not represent the thermostat; a self-model represents the system that contains it.
- **An inviolate level.** Hofstadter's subtlest condition: the crossing must be licensed by a level that is *itself not crossed*. Something has to enforce the rules that make the ascent legal — the proof rules that Gödel's numbering presupposes, the copying machinery that duplicates DNA without interpreting it, the substrate that a self-hosting compiler runs on. A loop with nothing held inviolate does not become profound; it becomes paradox or noise.

![Four panels contrasting feedback, recursion, self-reference and a strange loop](images/strange-loop-taxonomy.svg)

*Figure 2. Four things that are routinely called loops. Only the fourth is strange: the representation refers to the system that contains it, and the system acts on the representation.*

The distinctions are worth stating flatly, because most confusion about strange loops is really confusion about which of these is being claimed:

- **Feedback** cycles causation. Output changes input. Nothing in the loop represents the loop. *(A thermostat; a bank's interest calculation.)*
- **Recursion** applies a rule to a smaller instance of the same problem. The hierarchy is well-founded: the ladder descends, and it terminates. *(Factorial; a filesystem walk.)*
- **Self-reference** lets a representation denote itself. It may be benign, and it may be paradoxical — `"this sentence is false"` and a quine are the same syntactic trick with opposite temperaments.
- **Meta-representation** represents the system's own states without those states being changed by the representation. A model that models the modeler, but that the modeler never consults, is metacognition without a loop.
- **A strange loop** is the last two *plus* the arrow back: the represented level acts on the representation, which changes the represented level. The hierarchy is real but not well-founded — the top is made of the bottom.

What makes the pattern tractable rather than mystical is that its engine has a name. Every strange loop is a **fixed point**: an object that its own description maps onto. In arithmetic, the diagonal lemma manufactures a sentence `G` such that `G` and "`G` is unprovable" are the same claim. In computing, [Kleene's second recursion theorem](https://en.wikipedia.org/wiki/Kleene%27s_recursion_theorem) says the same thing in executable form — a program can be given access to its own source. In biology, a description is copied alongside the constructor that reads it. In a compiler, a source tree maps to the binary that maps back to the source tree. Hofstadter's loops are one idea: *the self-reference is not smuggled in from outside; it is a fixed point of the system's own operations.*

## A short history of the loop

![Timeline from 1747 to 2026 marking Bach's Canon per tonos, Gödel's incompleteness, Turing's universal machine, the Bombe, von Neumann's self-reproduction, Conway's Game of Life, Thompson's Trusting Trust, and self-rewriting agents](images/strange-loop-timeline.svg)

*Figure 3. Eight encounters with the same shape across three centuries. Music and logic were first; self-modifying software is the most recent.*

### 1747 — Bach's canon that climbs forever

The _Musical Offering_ ([BWV 1079](https://en.wikipedia.org/wiki/The_Musical_Offering)) is full of "riddle canons": Bach writes a single line of music and a cryptic instruction, and the performer has to work out the rest. The fifth is marked *canon per tonos* — a canon by tones. The instruction is to repeat the theme, each time a whole tone higher. Six repetitions of a whole tone is exactly an octave, so the music arrives back where it started while the listener hears an uninterrupted climb. Hofstadter heard a [Shepard tone](https://doi.org/10.1121/1.1919120) two centuries early.

![A staircase of six ascending pitch steps labelled C, D, E, F-sharp, G-sharp, A-sharp, then C again, with a dashed arrow returning to the start](images/strange-loop-shepard-tone.svg)

*Figure 4. The Shepard tone (Roger Shepard, 1964) is the perceptual version of Bach's trick: pitch height keeps rising while pitch class returns to its starting point, so the ear hears an ascent with no top. Roger Shepard, "Circularity in Judgments of Relative Pitch," Journal of the Acoustical Society of America 36(12), 2346–2353.*

Two honest notes, both of which sharpen the definition. First, Bach did not compose an illusion of endlessness; he composed a canon that ends where it began, and Hofstadter reinterpreted it. The loop was in the hearing, not the score. Second — and this is the part worth keeping — the loop was *completed by a system that represents the music*. The ear builds a representation of a rising scale and can be fooled by it because the representation, not the waveform, is what the listener hears. In Hofstadter's terms, the inviolate level here is the perceptual apparatus that enforces the rules of the illusion; the composer only supplied the notes.

### 1931 — Gödel turns mathematics on itself

For most of its history, mathematics was a tool pointed outward at puzzles. Kurt Gödel's 1931 paper — [_Über formal unentscheidbare Sätze der Principia Mathematica und verwandter Systeme I_](https://doi.org/10.1007/BF01700692) — pointed it inward.

The technique was [Gödel numbering](https://plato.stanford.edu/entries/goedel-incompleteness/): assign every symbol, formula, and proof a natural number, so that statements *about* arithmetic become statements *in* arithmetic. Then use the [diagonal lemma](https://plato.stanford.edu/entries/self-reference/) to construct a sentence `G` that asserts its own unprovability, in the medium of arithmetic:

> `G` ↔ "there is no proof, in this system, of the formula whose code is ⌜G⌝."

If the system proves `G`, it is inconsistent. If it is consistent, `G` is true and unprovable. Either way, no consistent, effectively axiomatized system strong enough for arithmetic is complete.

Three insights here are routinely lost, and they matter for everything downstream.

First, **the loop is not optional.** Gödel did not exhibit a curiosity; he proved that any system expressive enough to describe its own operations *contains* a fixed point of this kind. Self-reference is not a bug that careful mathematicians can avoid. It is a theorem about what expressiveness costs.

Second, **there is no paradox.** `G` is not `"this sentence is false."` It is a well-formed claim about natural numbers that happens, through the numbering, to be about itself. That is the whole reason the result is a theorem and not a joke: the loop is built out of ordinary syntax, and the inviolate level is the syntactic machinery that lets the numbering work.

Third, **truth outruns proof.** `G` is true and unprovable, so the system can talk about its own truth without being able to decide all of it. This is the first appearance of a pattern that recurs at every later stop on the road: the same expressiveness that lets a system describe itself puts part of itself permanently out of reach.

Self-reference of this kind is not exotic. It is a fixed point, and fixed points are executable:

```go
package main

import "fmt"

// main prints this file's own text: a fixed point, executable.
func main() {
	src := "package main\n\nimport \"fmt\"\n\n// main prints this file's own text: a fixed point, executable.\nfunc main() {\n\tsrc := %q\n\tfmt.Printf(src, src)\n}\n"
	fmt.Printf(src, src)
}
```

Run `go run quine.go | diff - quine.go` and the program is identical to its output. The string holds a hole (`%q`), the program fills the hole with the string's own quoted form, and the result is a text that describes itself using the language it is written in. Gödel numbering is that idea applied to arithmetic; `src` and `⌜G⌝` are cousins. Note what is *not* self-referential: the Go compiler, the terminal, the operating system. The loop has an inviolate level too.

### 1936–1940 — Turing's machine that reads machines

Alan Turing's [1936 paper](https://doi.org/10.1112/plms/s2-42.1.230) built the machine-level version of Gödel's trick. The **universal Turing machine** takes a description of any other machine as its input and then behaves like the machine described. One piece of hardware, a tape of descriptions, and the machine becomes whatever it reads. Encoding a machine as data is the same collapse Gödel achieved for formulas: the machine is a member of its own domain. The practical consequence is that *one* machine can be a member of its own domain at all — it can run a description of itself, given enough tape.

Turing also showed, by a diagonal argument, that the loop has a boundary: a machine cannot in general decide what another machine will do, and [the halting problem](https://plato.stanford.edu/entries/turing-machine/) is undecidable. The strange loop and its blind spot arrived in the same paper, which is the pattern this whole road report is tracking.

The engineering version came in 1940. The [Bombe](https://en.wikipedia.org/wiki/Bombe) at Bletchley Park was hardware thinking about hardware — but more precisely, to recover an Enigma key you build a machine whose wiring mirrors the cipher machine's internal logic, then search the space of possible rotor configurations by looking for logical contradictions imposed by a guessed fragment of plaintext. Turing and Gordon Welchman's design did not simulate Enigma in software; it *was* a physical consequence of Enigma's wiring, run in bulk. A machine built to model another machine's hidden state is a strange loop with a purpose — and, note, the cryptanalyst's crib is the inviolate level: the one piece of the system the loop does not get to define for itself.

### 1951 — von Neumann proves self-reproduction is not paradoxical

If a machine must contain a complete description of itself in order to build itself, the description must contain a description of the description, and so on forever. John von Neumann dissolved the regress in his lectures on automata and the posthumous [_Theory of Self-Reproducing Automata_](https://en.wikipedia.org/wiki/Von_Neumann_universal_constructor) (1966): separate the **description** (a tape, which is copied but not interpreted) from the **constructor** (the machine, which reads the tape). A universal constructor can build anything a description specifies, including another constructor, and can copy the description along with it. No infinite tower is required. Self-reproduction is possible for a finite machine, and von Neumann gave the blueprint.

Read that split in the language of this post and it says something crisp: **self-reproduction is a strange loop with an inviolate level.** The description may refer to the machine, and the machine may copy the description, but the copying is done *without interpretation* — the level that copies is not itself part of what is being tangled. Get that separation wrong and self-reproduction turns into a definitional circle; get it right and it is just engineering.

Biology had already shipped it. In a cell, DNA is the description and proteins are the constructors: DNA is inert, and it cannot be read, copied, or repaired without enzymes, which are themselves encoded by DNA and built by ribosomes made of RNA and protein. The data builds the machine; the machine is required to read the data.

![DNA, the inert blueprint, encoding proteins, the machinery, which in turn reads, replicates and repairs DNA; neither can come first, and RNA can be both template and catalyst](images/strange-loop-dna-protein.svg)

*Figure 5. The chicken-and-egg knot of molecular biology, and its leading resolution. The central dogma is Francis Crick's, ["Central Dogma of Molecular Biology,"](https://doi.org/10.1038/227561a0) Nature 227, 561–563 (1970). The RNA world is Walter Gilbert's, ["The RNA World,"](https://doi.org/10.1038/319618a0) Nature 319, 618 (1986).*

Neither DNA nor protein could have come first, which is a genuine historical puzzle and not a paradox — von Neumann had already shown that a finite system can bootstrap itself given a description, a constructor, and copying. The leading biological answer is the **RNA world**: RNA can be both template and catalyst, so a single molecule can play both roles. [Ribozymes](https://en.wikipedia.org/wiki/Ribozyme) are the surviving evidence, and laboratory evolution has produced RNA enzymes that can extend and copy RNA templates ([Bartel and Szostak, _Science_ 261, 1411, 1993](https://pubmed.ncbi.nlm.nih.gov/7690155/)). What looks like a loop to be escaped is, in von Neumann's framework, just two roles temporarily filled by one molecule. Every living thing you have ever met is the stable version of that loop.

### 1970 — Conway's Life builds a computer that runs Life

John Conway's [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) appeared as a puzzle in Martin Gardner's *Scientific American* column in October 1970: a grid of cells that live, die, or are born according to three rules about neighbors. It has no designer, no agent, and no memory beyond the current grid.

Arrange enough cells correctly and the mindless physics becomes logic: gliders, eaters, glider guns, then logic gates, then a [Turing-complete](https://en.wikipedia.org/wiki/Turing_completeness) machine. People have since built constructors that build Life patterns *inside* Life — the [OTCA metapixel](https://conwaylife.com/wiki/OTCA_metapixel) is a cell-sized implementation of Life cells, meaning a Life pattern runs a simulation of Life, which can run a simulation of Life, bounded only by the grid you can afford. Von Neumann's universal constructor, rediscovered on a lattice of two-state cells.

This is the most important stop on the road, because it is where the cheapest possible version of the loop appears. Life has self-describing structures, self-simulation, and universal computation, and it has no hint of a self. The definition from earlier is satisfied completely and nothing resembling a mind results. That is not a footnote; it is the calibration. **A strange loop is cheap.** Any claim that loops explain cognition has to explain what distinguishes the loop in a brain from the loop in a grid of two-state cells, and "it is a strange loop" cannot be the answer, because Conway's Life is one too.

### 1984 — the loop with teeth: Trusting Trust

A compiler is the tool that turns source into programs. If you want a compiler for a new language, you use an existing language to write a rough version, then use that rough version to compile a better compiler written in the new language, then compile the compiler with itself. The result is **self-hosting**: the tool is now written in the language it compiles. Rust's first compiler was written in OCaml; today `rustc` compiles `rustc`.

Ken Thompson's 1984 Turing Award lecture, [_Reflections on Trusting Trust_](https://dl.acm.org/doi/10.1145/358198.358210), showed what the closing of that loop costs. Suppose you modify the C compiler so that when it compiles `login`, it inserts a backdoor. Then add a second modification: when the compiler compiles *its own source*, it inserts both modifications into the new binary. Now delete the modifications from the source. The compiler is clean; every program it builds is clean in source and compromised in binary; and the misbehavior reproduces itself through every future build. The backdoor has no textual footprint anywhere in the system.

> "You can't trust code that you did not totally create yourself."

The loop is the vulnerability. A self-hosting compiler is a fixed point, and the tool that defines what "correctly built" means is itself only as trustworthy as its last unverified ancestor. The epistemological reading is sharper than the security one: **source code is a description, and a description cannot verify the machine that reads it.** Reading is not evidence; provenance is. The remedies are therefore not about the source at all — *diverse double-compiling* ([Wheeler, 2009](https://arxiv.org/abs/1004.5548)) cross-checks a compiler against an independently produced one, and reproducible builds let a distrusting third party reconstruct the binary and compare. This is the same conclusion this blog reached about [event logs and audit trails](https://blog.hackspree.com/#events-as-the-source-of-truth) and about [agent governance](https://blog.hackspree.com/#agent-governance-toolkit), arriving from a different direction: when a system can describe itself, verification has to come from a level the system does not control.

### 2023–26 — the agent that rewrites its own harness

The most recent stop is the one currently being sold to you as a product. An LLM agent runs, collects traces and failures, and then — instead of just trying again — writes or modifies a tool, prompt, scaffold, or policy and feeds the result back into its own loop. That is a strange loop in the strict sense: the bottom level of the hierarchy (the tool the agent uses) rewrites the top level (the agent's own procedure), through the agent's own rules.

![A four-stage circular loop: run, observe traces, rewrite a tool or policy, re-enter it into the agent's own loop, with a note on the Trusting Trust risk of a loop that authors its own compiler](images/strange-loop-agent-loop.svg)

*Figure 6. The agentic strange loop, and the security shadow it casts. The stages are ordinary engineering; what makes the loop strange is stage four, where the artifact authored at stage three becomes part of the machinery that authors the next one.*

The idea has been accumulating concrete results. [Voyager](https://arxiv.org/abs/2305.16291) (Wang et al., 2023) had an LLM write reusable skills into a growing library and retrieve them later, improving without gradient updates. [Reflexion](https://arxiv.org/abs/2303.11366) and [Self-Refine](https://arxiv.org/abs/2303.17651) turned the agent's own critique into the next attempt's input. [STOP](https://arxiv.org/abs/2310.02304) (Zelikman et al., 2023) let a seed program improve its own improver. [ADAS](https://arxiv.org/abs/2408.08435) (Hu et al., 2024) searched the space of agent architectures in code. [AlphaEvolve](https://deepmind.google/discover/blog/alphaevolve-a-gemini-powered-coding-agent-for-designing-advanced-algorithms/) evolved algorithms against automatic evaluators. [Gödel Agent](https://arxiv.org/abs/2410.04444) (Yin et al., 2024) is explicitly a self-referential framework in which the agent edits its own logic. And the [Darwin Gödel Machine](https://arxiv.org/abs/2505.22954) (Zhang et al., 2025) makes the loop open-ended by having agents rewrite their own code and keeping the variants that improve on benchmarks — a population of self-modifying programs, selected empirically.

Three things distinguish this stop from Conway's, and only the first is about speed.

First, the loop now runs in software that can be edited in a second, so the "generations" between observation and rewrite are hours, not decades. The mutation rate of the loop is the whole story of why this is suddenly a product.

Second, **the loop decides, not just describes.** A quine is harmless; it prints itself and stops. A Gödel sentence is inert; it is true and unprovable. But an agent acts on its self-description, which turns self-reference from a curiosity into a control problem. The loop's output is not a theorem; it is the next version of the thing that produced it.

Third, and hardest: **the loop needs a validation signal that is not itself inside the loop** — an inviolate level. The Darwin Gödel Machine's authors are careful here (variants are kept only when a benchmark moves), and that care is the whole design. A self-grading system can improve itself into confident nonsense, and it will do so faster than a human can notice. So the interesting engineering is never the optimizer; it is the [harness that contains it](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems), the [sandbox that bounds it](https://blog.hackspree.com/#sandboxing-ai-agents), and the [provenance trail that can reconstruct what it did](https://blog.hackspree.com/#deepseek-harness). The [loop-engineering discourse](https://blog.hackspree.com/#loop-engineering) is about building these loops; the [self-improving workflow](https://blog.hackspree.com/#self-improving-agent-workflows) and [spatiotemporal composability](https://blog.hackspree.com/#spatiotemporal-composability) work is about what happens when the loops start composing with each other.

And note what the agent's loop inherits from 1984. An agent that writes its own tools, and whose tools run with the agent's own trust, has exactly Thompson's problem. The fix has the same shape too: hold some level inviolate — the evaluator, the sandbox, the log — and let it be the thing the loop cannot rewrite.

## What every loop hides

Step back and the road has a shape that none of the individual stops announces. At every one of them, the same act that gives a system the power to describe itself fixes a limit it cannot get past.

![A row-by-row pairing of what each loop enables, from Gödel to minds, against the blind spot that comes with it](images/strange-loop-blind-spot.svg)

*Figure 7. The pattern behind the history. The power and the limit are not a tradeoff imposed from outside; they are the same fixed point seen from two sides.*

Gödel's arithmetic can prove things about arithmetic and cannot prove its own consistency. Turing's machine can run any machine's description and cannot decide whether it halts. Thompson's compiler can compile itself and cannot verify its own provenance. An agent can rewrite its own tools and cannot validate, from inside, the signal it is rewriting them against. And — if the pattern holds — a brain can model the modeler and cannot, from inside the model, explain why the model feels like anything.

Why must it be this way? Because of the definition's last clause. A strange loop has no level outside the loop from which the whole thing can be surveyed, so the system's self-description is always *partial*, and the part it misses is precisely the part doing the describing. Gödel located that missing part: the fixed point `G` is true, and the system cannot derive it. Turing located it: the one question you cannot ask the universal machine about itself is whether it will stop. Thompson located it: the compiler whose output you trust is the one artifact whose trustworthiness you cannot establish by reading. In each case, self-reference and the blind spot are not two facts. They are one fact, described from inside and outside.

That gives a sharper reading of the whole road than "systems get more self-referential." Each stop is a system buying expressiveness with a blind spot, and the engineering of each stop is the management of that blind spot:

- Gödel's result is a *theorem* about where the boundary falls.
- Thompson's result is a *warning*: the boundary can be exploited, and it hides.
- The agent case is a *design brief*: choose which level stays inviolate, then make the choice stick.

And it explains why the loop is not merely a curiosity about self-reference. A system with a strange loop is one that cannot, even in principle, be fully explained from outside itself — not because it is mysterious, but because its self-description is one of its own parts, and the description is incomplete in exactly the way the theorem says it must be. When people say that a mind cannot be understood by a third-person inventory of its parts, this is the precise claim they are reaching for, and it is a claim about *where the vantage point is*, not about anything supernatural.

## The spark of awareness

Here is where Hofstadter goes further than the engineering, and where the reader should hold onto the difference between a mechanism and a proof.

His argument, compressed: a brain is billions of cells, none of which is aware. Those cells build representations. Some of those representations are representations of *the system doing the representing* — its body, its history, its categories, its plans. When the representational hierarchy loops back on itself and the system's model of the world includes a model of the modeler, a new kind of pattern exists in the substrate: a self. In [_I Am a Strange Loop_](https://en.wikipedia.org/wiki/I_Am_a_Strange_Loop) (Basic Books, 2007) he pushes it to the limit: the "I" is not a thing in the brain but a *pattern* — a loop of self-description stable enough to persist, refer to itself, and be recognized by others even after the underlying matter has been replaced. On this view, souls are not stuff; they are a kind of software that a sufficiently tangled hierarchy can run.

It is a serious philosophical position, and it is not settled. Four caveats, in increasing order of difficulty.

First, **the loop is under-specified.** Conway's Life contains self-describing structures and is not conscious; everything in this post below this section is a strange loop and none of it is aware. Hofstadter's claim is that the *right kind* of self-model is what matters — one rich enough to represent itself as a self, in a system with the right dynamics — but "the right kind" is doing most of the work, and it is not yet a testable criterion.

Second, **rival accounts exist and are not obviously worse.** Global workspace theory ([Baars](https://en.wikipedia.org/wiki/Global_workspace_theory)) locates consciousness in a broadcast bottleneck; integrated information theory ([Tononi](https://en.wikipedia.org/wiki/Integrated_information_theory)) locates it in a system's causal structure; Thomas Metzinger's [self-model theory](https://en.wikipedia.org/wiki/Being_No_One) treats the self as a transparent model the brain cannot see through. These overlap with the strange-loop picture — all of them involve self-representation — but they are not the same claim, and they make different predictions, which is the only way any of them will be adjudicated.

Third, **self-reference gives you a self-model, not a stake in it.** This is the gap the loop's own structure makes visible. A strange loop can represent itself and act on the representation, and still have no *reason* to care which way the representation comes out. Something else has to make the model's states matter to the system — needs, drives, the maintenance of its own viability. That is the line Antonio Damasio draws between [homeostasis and feeling](https://en.wikipedia.org/wiki/The_Strange_Order_of_Things), and the one Mark Solms pushes hardest, arguing that affect rather than recursion is [the hidden spring](https://en.wikipedia.org/wiki/Mark_Solms) of consciousness. If they are right, the strange loop supplies the architecture of a self and something more visceral supplies its stake, and Hofstadter's account is a theory of the former mistaken for a theory of both. Hofstadter would likely reply that a sufficiently rich loop *is* the stake — that a system modeling its own survival has already located what matters to it — and that disagreement is exactly the unresolved one.

Fourth, **the hardest part is untouched.** Explaining why self-representation should feel like anything is the [hard problem](https://plato.stanford.edu/entries/consciousness/), and a strange loop is a structural description, not a solution to it. Saying "the loop becomes a self" is a claim about architecture. It does not explain why architecture is accompanied by experience, and Hofstadter, to his credit, does not pretend it does — he relocates the mystery rather than dissolving it.

What the strange loop *does* explain well is the thing that makes the self look paradoxical in the first place: how a system composed of parts that do not understand anything can contain an understanding of itself. Gödel answered that for arithmetic, von Neumann for reproduction, and Conway for physics with three rules. The same move in a substrate that keeps a running model of its own situation gives you a self-model — and the reason the loop feels like magic is that we are such a loop, looking at itself from inside.

## On the road

The road begins with a canon that ends where it started and arrives, for now, at an agent editing its own tools. Each stop bought a system something real: music bought an illusion, arithmetic bought self-knowledge and incompleteness, machines bought universality and undecidability, life bought self-reproduction, cells bought both at once, compilers bought self-hosting and unverifiable trust, agents bought self-modification and a validation problem. The pattern is not that the loops get more elaborate. It is that *every loop pays for its self-description with a blind spot, and each era's engineering is the management of that blind spot.*

Which is why the title says "on the road to" rather than "the theory of." We can now build loops that model themselves, rewrite themselves, and come back for another pass; we can say with confidence what such a loop *is* and what it cannot see. What we cannot yet do is the last step: demonstrate that any of it is accompanied by experience, or find the property — richness, stakes, valence, all three — that separates the loop in a brain from the loop in a grid of two-state cells. Until that is settled, the cognitive strange loop is a destination we are driving toward, not one we have arrived at.

Two exits off the current stretch are worth taking:

- **Outward, into security.** When agents can modify the code they run on, Trusting Trust becomes an operational problem rather than a compiler curiosity: the artifact that authors your next agent cannot be verified by reading it. Naming the inviolate level, and making it genuinely inviolate — reproducible builds, diverse double-compiling, sandboxed self-modification, provenance logs — is the whole engineering agenda.
- **Inward, into psychology.** The same structure shows up in human behavior: a self-model that cannot see through itself generates loops that are stable, self-confirming, and invisible from the inside. Rumination, addiction, and confirmation bias are strange loops in the substrate of a person — and treating them as feedback problems rather than problems of an incomplete self-description gets the intervention wrong.

Neither is a metaphor. That is the point. A strange loop is not a circle drawn around a problem; it is the shape a system takes when it becomes part of its own subject matter — and, if Hofstadter is right, when it becomes something it can no longer stand outside of.

## References

**Primary sources for the loop**

- Douglas R. Hofstadter, _Gödel, Escher, Bach: An Eternal Golden Braid_ (Basic Books, 1979) — the book that named the strange loop; see the chapters on the endlessly rising canon and "Strange Loops, or Tangled Hierarchies," where the _inviolate level_ is introduced.
- Douglas R. Hofstadter, _I Am a Strange Loop_ (Basic Books, 2007) — the self as a pattern of self-reference rather than a thing.
- Stanford Encyclopedia of Philosophy: ["Gödel's Incompleteness Theorems"](https://plato.stanford.edu/entries/goedel-incompleteness/) and ["Self-Reference"](https://plato.stanford.edu/entries/self-reference/).

**Logic, fixed points, and computation**

- Kurt Gödel, "Über formal unentscheidbare Sätze der Principia Mathematica und verwandter Systeme I," _Monatshefte für Mathematik und Physik_ 38, 173–198 (1931) — https://doi.org/10.1007/BF01700692
- Stephen C. Kleene, "On notation for ordinal numbers," _Journal of Symbolic Logic_ 3(4), 150–155 (1938) — the second recursion theorem, the fixed-point result behind programs that can access their own source: https://en.wikipedia.org/wiki/Kleene%27s_recursion_theorem
- Alan M. Turing, "On Computable Numbers, with an Application to the Entscheidungsproblem," _Proceedings of the London Mathematical Society_ s2-42, 230–265 (1936) — https://doi.org/10.1112/plms/s2-42.1.230
- Stanford Encyclopedia of Philosophy, ["Turing Machines"](https://plato.stanford.edu/entries/turing-machine/) — the universal machine and the halting problem.
- Bletchley Park [Bombe](https://en.wikipedia.org/wiki/Bombe) — Turing and Welchman's electromechanical search over Enigma's state space.

**Music and perception**

- J.S. Bach, _Musikalisches Opfer_ (Musical Offering), BWV 1079 (1747) — https://en.wikipedia.org/wiki/The_Musical_Offering; the fifth canon is _canon per tonos_.
- Roger N. Shepard, "Circularity in Judgments of Relative Pitch," _Journal of the Acoustical Society of America_ 36(12), 2346–2353 (1964) — https://doi.org/10.1121/1.1919120
- M.C. Escher, _Drawing Hands_ (1948) — https://en.wikipedia.org/wiki/Drawing_Hands

**Self-reproduction and biology**

- John von Neumann, _Theory of Self-Reproducing Automata_ (edited by Arthur Burks, University of Illinois Press, 1966) — the description/constructor split that removes the infinite regress; see https://en.wikipedia.org/wiki/Von_Neumann_universal_constructor
- Francis Crick, "Central Dogma of Molecular Biology," _Nature_ 227, 561–563 (1970) — https://doi.org/10.1038/227561a0
- Walter Gilbert, "The RNA World," _Nature_ 319, 618 (1986) — https://doi.org/10.1038/319618a0
- David P. Bartel and Jack W. Szostak, "Isolation of new ribozymes from a large pool of random sequences," _Science_ 261, 1411–1418 (1993) — https://pubmed.ncbi.nlm.nih.gov/7690155/

**Simulation**

- Martin Gardner, "Mathematical Games: The fantastic combinations of John Conway's new solitaire game 'life'," _Scientific American_ 223(4), 120–123 (October 1970).
- [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) — Turing-completeness and self-simulation; the [OTCA metapixel](https://conwaylife.com/wiki/OTCA_metapixel) is a Life pattern that simulates Life.

**Software, trust, and self-hosting**

- Ken Thompson, "Reflections on Trusting Trust," _Communications of the ACM_ 27(8), 761–763 (1984) — https://dl.acm.org/doi/10.1145/358198.358210
- David A. Wheeler, "Fully Countering Trusting Trust through Diverse Double-Compiling" (2009) — https://arxiv.org/abs/1004.5548

**Self-modifying agents**

- Guanzhi Wang et al., "Voyager: An Open-Ended Embodied Agent with Large Language Models" (2023) — https://arxiv.org/abs/2305.16291
- Noah Shinn et al., "Reflexion: Language Agents with Verbal Reinforcement Learning" (2023) — https://arxiv.org/abs/2303.11366
- Aman Madaan et al., "Self-Refine: Iterative Refinement with Self-Feedback" (2023) — https://arxiv.org/abs/2303.17651
- Eric Zelikman et al., "Self-Taught Optimizer (STOP): Recursively Self-Improving Code Generation" (2023) — https://arxiv.org/abs/2310.02304
- Shengran Hu et al., "Automated Design of Agentic Systems" (2024) — https://arxiv.org/abs/2408.08435
- Xunjian Yin et al., "Gödel Agent: A Self-Referential Agent Framework for Recursive Self-Improvement" (2024) — https://arxiv.org/abs/2410.04444
- Jenny Zhang et al., "Darwin Gödel Machine: Open-Ended Evolution of Self-Improving Agents" (2025) — https://arxiv.org/abs/2505.22954
- Google DeepMind, "AlphaEvolve: a Gemini-powered coding agent for designing advanced algorithms" (2025) — https://deepmind.google/discover/blog/alphaevolve-a-gemini-powered-coding-agent-for-designing-advanced-algorithms/
- Jeff Clune, "AI-GAs: AI-Generating Algorithms, an Alternate Paradigm for Producing General Artificial Intelligence" (2019) — https://arxiv.org/abs/1905.10985

**Consciousness, for the sceptical reader**

- Bernard Baars, [_A Cognitive Theory of Consciousness_](https://en.wikipedia.org/wiki/Global_workspace_theory) — global workspace theory.
- Giulio Tononi, [_Integrated Information Theory_](https://en.wikipedia.org/wiki/Integrated_information_theory).
- Thomas Metzinger, [_Being No One: The Self-Model Theory of Subjectivity_](https://en.wikipedia.org/wiki/Being_No_One) (MIT Press, 2003).
- Antonio Damasio, [_The Strange Order of Things: Life, Feeling, and the Making of Cultures_](https://en.wikipedia.org/wiki/The_Strange_Order_of_Things) (Pantheon, 2018) — homeostasis, feeling, and why a self-model needs stakes.
- Mark Solms, [_The Hidden Spring: A Journey to the Source of Consciousness_](https://en.wikipedia.org/wiki/Mark_Solms) (Profile Books, 2021) — affect as the substrate of consciousness, contra recursion-only accounts.
- Stanford Encyclopedia of Philosophy, ["Consciousness"](https://plato.stanford.edu/entries/consciousness/) — on the hard problem and why structural accounts leave something out.
