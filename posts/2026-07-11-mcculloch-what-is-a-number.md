---
title: "McCulloch's question: what is a number, that a man may know it?"
date: 2026-07-11
slug: mcculloch-what-is-a-number
summary: "Warren McCulloch's 1960 lecture traces a single question — asked at age 19 — through Augustine, Duns Scotus, Hume, Russell, Turing, and the neural circuitry of the brain, ending with a probabilistic logic for fallible neurons."
tags: mcculloch, cybernetics, neural-networks, logic, epistemology, history-of-science
---

In 1917, a nineteen-year-old Warren McCulloch was called into the office of the Quaker philosopher Rufus Jones at Haverford College.

"Warren," said Jones, "what is thee going to be?"

"I don't know."

"And what is thee going to do?"

"I have no idea; but there is one question I would like to answer: What is a number, that a man may know it, and a man, that he may know a number?"

Jones smiled. "Friend, thee will be busy as long as thee lives."

Forty-three years later, McCulloch delivered the Alfred Korzybski Memorial Lecture under that same title. He had been busy. The lecture is a compressed intellectual autobiography — part history of philosophy, part neuroscience manifesto, part mathematical logic — and it traces how one man spent a lifetime trying to ground epistemology in the physics and chemistry of the brain.

## The two halves of the question

The title is recursive for a reason. McCulloch's argument is that you cannot answer "what is a number?" without also answering "what is a knower?" The two questions are one question. Any theory of number that does not account for the biological system that apprehends number is incomplete. Any theory of mind that cannot account for how a physical system grasps mathematical truth is insufficient.

This is not a casual framing. It is the central methodological commitment of McCulloch's career: **reduce epistemology to an experimental science**. Not by philosophizing about knowledge, but by understanding the neural circuitry that produces it. The lecture is the story of that attempt.

## The philosophical arc: from Augustine to Hume

McCulloch traces the Western approach to number through four theological principles, each of which maps onto an epistemological position.

**The eternal verities.** Augustine, around 500 AD: "7 and 3 are 10; 7 and 3 have always been 10; 7 and 3 at no time and in no way have ever been anything but 10; 7 and 3 will always be 10." These are truths independent of time and place — ideas in the Mind of God, which we can understand but never fully comprehend. Augustine's examples are drawn from arithmetic, geometry, and logic, but he includes what we would now call the laws of Nature. The history of Western science — Galileo, Newton, Einstein, the tensor invariant — would have been no shock to his theology.

**Authority.** The scholastic reliance on texts and their interpretations. McCulloch moves quickly through this, noting the persistent questions of textual corruption and translation, but his real interest is in what came next.

**The shift to experiment.** Roger Bacon insisted that the eternal verities and the authorities must be tested not only by reason — the old meaning of "experiment" — but by looking again at Nature. Natural law began to grow. Duns Scotus, the last great scholastic defender of realistic logic, formalized the three roads to truth: deduction (from rules and cases to facts), induction (from cases and facts to rules), and abduction (from rules and facts to the hypothesis that the fact is a case under the rule). Abduction is the breeding place of scientific ideas, of intuition, of insight. McCulloch would return to this at the end of his life as the unsolved problem.

**The nominalist break.** William of Ockham, "greatest of nominalists," severed logic from empirical science. His demand that no conclusion contain what was not in the premises barred two of the three roads to truth and reduced the third — deduction — to vacuous tautology. Logic decayed. Science was born. Law usurped the throne of Theology, and Science began to usurp the throne of Law.

Then came Hume. At twenty-three, Hume had shown that only in logic and arithmetic can we argue through any number of steps, because only here do we have the proper test of equality: "When a number hath a unit answering to an unit of the other we pronounce them equal." One-to-one correspondence.

## What a number is

Bertrand Russell, whom McCulloch credits as the first to thank Hume properly, gave the definition: **a number is the class of all classes that can be put into one-to-one correspondence with it.** The number 7 is the class of all classes that can be put into one-to-one correspondence with the days of the week.

McCulloch accepts this definition and then makes a critical observation: the numbers 1 through 6 are perceptibles. Experiments on many animals — birds, rats, primates — show this. A creature can distinguish 3 from 4 without counting. These are, in Ockham's phrase, natural terms — shared with the beasts. All larger integers are countables, arrived at by a symbolic process: putting pebbles in pots, cutting notches in sticks, establishing one-to-one correspondences. These are conventional terms — "tricks for setting things into one-to-one correspondence" — that grew out of communication, out of logos.

The definition of number depends on two foundations: the perception of small whole numbers (a biological given) and a symbolic process of one-to-one correspondence (a cultural achievement). This is the answer to the first half of the question.

## What a man is

The second half is harder. McCulloch's attempt to answer it consumed his career.

In 1923, he attempted to write a logic of transitive verbs and failed. The problem was too hard, the available logic too primitive. He pivoted to a different approach: invent a "least psychic event" — a psychon — with four properties. It either happened or it didn't. It happened only if its bound cause had happened. It proposed this fact to subsequent psychons. And these could be compounded to produce more complex propositions.

In 1929 it dawned on him: these events might be the all-or-none impulses of neurons, combined by convergence onto the next neuron to yield propositional complexes. A neuron fires if and only if its input conditions are met. That firing implies its antecedent. It signals this to downstream neurons. Networks of such units compute logical functions of their inputs.

This was the seed of the 1943 paper with Walter Pitts, "A Logical Calculus of the Ideas Immanent in Nervous Activity" — one of the founding documents of neural network theory, artificial intelligence, and computational neuroscience.

## The 1943 paper and its consequences

The paper proved several things simultaneously. First, that networks of neurons computing simple logical functions could extract any configuration of signals from their input. Second, because Gödel had arithmetized logic, and Turing had shown that a simple machine could compute any computable number, **nets of neurons were equivalent to Turing machines**. The brain was a computer, and computers could be brains — not metaphorically, but formally, in the structure of the proof.

But the paper did more. Pitts's modulo mathematics allowed them to analyze circuits with closed paths — reverberating loops — and to set up a theory of memory. A memory is a temporal invariant: given an event at one time, and its regeneration at later times, one knows that there was an event of that kind. In logical notation: (∃x)(ψx). There exists some x such that x was a ψ. Given this and negation — for which neural inhibition suffices — you get the lower predicate calculus with equality, which was recently proved to be a sufficient logical framework for all of mathematics.

The brain, in other words, contains the logical machinery for mathematics. This is not a metaphor. It is a claim about circuit structure.

McCulloch and Pitts followed this with "How We Know Universals" (1947), which generalized the mechanism. Any object — any universal — is an invariant under some group of transformations. A square is an invariant under 90-degree rotations. A face is an invariant under changes in expression, angle, and lighting. The neural net need only compute averages over the group of transformations to recognize the universal. The mechanism is general.

## Von Neumann's problems: reliability from unreliable parts

The second half of the lecture concerns a set of problems posed by John von Neumann, who had absorbed the McCulloch-Pitts model and used it in teaching computing machine theory. Von Neumann wanted to know: how can a brain made of fallible neurons compute reliably?

His attempt — "Toward a Probabilistic Logic" — made three assumptions, each of which was fatal. He assumed failures were absolute (not dependent on signal strength or threshold). He assumed neurons had only two inputs. He assumed each computed the same single function. Under these constraints, reliable computation from unreliable parts is nearly impossible.

McCulloch and his collaborators spent years dismantling these assumptions one by one. With Leo Verbeek, he showed that the error probability of a net can be reduced to the error probability of a single output neuron, and that this can be further reduced by parallel output neurons. With Manuel Blum, he proved that excitation, inhibition, and inhibitory interaction between afferent fibers are necessary and sufficient for neurons to compute their logical functions across the full range from coma to convulsion. With Eugene Prange, he showed that neurons with multiple inputs and controlling signals can compute the vast majority of possible logical functions — the great logical redundancy of the system buying reliability.

The brain does not work despite its unreliable components. It works *because* of the redundancy that unreliable components make necessary. The logic is probabilistic — not a logic of probabilities (where the arguments are uncertain but the logical operations are certain), but a probabilistic logic (where the functions themselves are infected by chance). This is a thing "of which Aristotle never dreamed."

## The Venn functions

McCulloch closes with a practical tool he developed for teaching this logic to neurologists and psychiatrists. Using a simplified Venn diagram notation — an X with four quadrants, each marked with a jot (true), blank (false), or probability p — he created a visual calculus for probabilistic logic that a twelve-year-old could learn in minutes.

The point is not the notation. The point is that the logic of fallible neurons is tractable. It can be taught. It can be computed. It can be programmed. The tools exist to reason rigorously about systems whose components misbehave.

## The unanswered question

McCulloch ends with a confession. The problem of insight — of intuition, of invention, of abduction — is still unsolved. A child learns at least one logical particle ("neither" or "not both") from ostension alone. How? We do not know. Tarski thought we lacked a fertile calculus of relations of more than two relata. McCulloch, at sixty-two, felt too old to tackle it: "Too bad — I'm too old. I may live to see the youngsters do it."

He did not. The problem of how a physical system generates insight — how it abduces a hypothesis from a rule and a fact — remains open. McCulloch would have been delighted by the progress in machine learning, but he would also have noted that we still do not have a satisfactory theory of abduction stated in terms of neural circuitry. The question he asked Rufus Jones in 1917 is not fully answered.

## Why this lecture matters

McCulloch's lecture is easy to misread. It looks like a ramble through the history of philosophy, some personal anecdotes, and a technical appendix on probabilistic logic. It is actually a unified argument:

1. **Numbers are relations of one-to-one correspondence**, built on a biological foundation (the perception of small quantities) and a cultural one (the symbolic process of counting).
2. **A knower is a neural system** whose circuitry implements the lower predicate calculus with equality — sufficient for all of mathematics — using neurons as propositional units and reverberating loops as memory.
3. **The logic must be probabilistic** because the components are fallible, but the system can be made arbitrarily reliable through redundancy.
4. **The unsolved problem is abduction** — how a physical system generates hypotheses — and it is the problem that connects epistemology to neuroscience.

The lecture is McCulloch's intellectual testament. He died nine years later. The question in its title is still worth sitting with: not because we lack answers, but because the answers we have reveal how much we still do not understand about how a physical system — a brain, a body, a network of fallible cells — comes to know that 7 and 3 are 10, and always have been, and always will be.

## The question, sixty-five years later

In February 2025, a paper appeared on arXiv with a title that would have made McCulloch smile: **"What is a Number, That a Large Language Model May Know It?"** Raja Marjieh, Veniamin Veselovsky, Thomas Griffiths, and Ilia Sucholutsky picked up McCulloch's question and pointed it at the system that has come closest to realizing his vision of a thinking machine.

Their finding is both elegant and troubling. LLMs represent numbers through an **entangled representation** — a blend of string-based similarity (Levenshtein edit distance) and numerical magnitude (log-linear distance). The digit sequence "911" activates both its string properties (it looks like "191" and "119") and its quantity properties (it is close to 910 and 912). These two representations are not cleanly separable in the model's latent embeddings. Context can reduce the entanglement but cannot eliminate it.

This is not how humans know numbers. McCulloch's distinction between perceptibles (1–6, shared with beasts, grounded in dedicated neural circuitry for quantity) and countables (larger integers, requiring symbolic convention and one-to-one correspondence) describes a cognitive architecture with clean boundaries. The biological brain has specialized circuits for numerosity. The LLM has a single set of weights that must simultaneously learn that digits are symbols with edit distances and quantities with magnitudes. The two forms of knowledge bleed into each other.

The practical consequences are real. The paper shows that representational confusion propagates into downstream decisions. Ask an LLM to reason about numbers and it may be influenced by how similar the digit strings *look* rather than what they *mean*. This is a category error in silicon — exactly the kind of failure that McCulloch's careful distinction between natural terms and conventional terms was designed to prevent.

The deeper point is that McCulloch's question was always about **architecture**. What kind of system can know a number? His answer was specific: a system with propositional units (neurons), temporal invariants (memory loops), and a mechanism for computing invariants under groups of transformations (universals). The LLM is a different architecture — a stack of attention layers trained on next-token prediction — and it knows numbers differently as a result. The entanglement that Marjieh et al. document is not a bug in the training data. It is a consequence of the architecture.

McCulloch would have been fascinated. He would have recognized the paper as empirical epistemology in his own tradition — asking, experimentally, how a particular kind of knower represents number. And he would have noted, with some satisfaction, that sixty-five years after his lecture, the question he asked Rufus Jones still organizes a research program.

> "What is a number, that a man may know it, and a man, that he may know a number?" — Warren McCulloch (1961)

---

**Reference:** Warren S. McCulloch, "What Is a Number, that a Man May Know It, and a Man, that He May Know a Number?" Alfred Korzybski Memorial Lecture, 1960. Published in *General Semantics Bulletin*, No. 26/27, 1960, pp. 7–18. [Full text (PDF)](https://www.nsl.com/k/parry/mcculloch_what-is-a-number.pdf).

**Related:** Raja Marjieh, Veniamin Veselovsky, Thomas L. Griffiths, Ilia Sucholutsky. ["What is a Number, That a Large Language Model May Know It?"](https://arxiv.org/abs/2502.01540) arXiv:2502.01540, February 2025.
