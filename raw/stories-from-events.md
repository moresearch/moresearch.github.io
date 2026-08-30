---
title: "Make Stories from Events: Why Event-Sourced Systems Need a Narrative Layer"
date: 2026-08-30
slug: stories-from-events
summary: "The brain does not experience reality raw — it segments it into scenes, edits it into stories, and matches those stories against a library of familiar scripts. Event-sourced systems keep only the raw footage. This post argues for projecting a story layer over the event log, and for comparing stories to stories instead of events to events."
tags: [event-sourcing, cognitive-science, story-abstraction, cqrs, sagas, similarity, case-based-reasoning]
---

## The shapes, and the story

In 1944, Fritz Heider and Marianne Simmel showed 114 subjects a crude animated film: a small triangle, a big triangle, a small circle, and a rectangle with a doorlike flap, moving against a plain background. Asked what they saw, almost nobody described the physical motion of the shapes. The overwhelming majority "interpreted the picture in terms of actions of animated beings, chiefly of persons." To some the shapes were a love triangle; to others, a family drama; to still others, an episode of bullying. Nobody was instructed to narrate. The brain could not watch the shapes and *not* make a story.

Your event-sourced system is in exactly that position. Its append-only log is a stream of shapes: `order.cancelled`, `payment.failed`, `email.sent` — precise, immutable, timestamped, replayable, true. And like the triangle's coordinates, each event is true and meaningless. Meaning is not stored in the events. It exists only in the story you cut from them. That is why scientists studying the brain's response to movies keep arriving at the same design: the brain does not record experience; it **edits reality into a story**, and it reads new situations by matching them against stories it already knows ([Scientific American, "The brain doesn't just experience reality. It edits it into a story"](https://www.scientificamerican.com/article/the-brain-doesnt-just-experience-reality-it-edits-it-into-a-story/)). The systems we build should do the same.

## The brain is a story machine

The neuroscience here is not metaphor. Over the past decade, researchers put people inside fMRI scanners and decoded their brain activity while they watched movies. Three findings matter for system design.

**Brains synchronize on stories.** In 2004, Uri Hasson's lab scanned five people watching half an hour of *The Good, the Bad and the Ugly* — the sort of complex, dynamic stimulus that was considered untractable for fMRI ([Hasson et al., *Science* 303, 2004](https://pubmed.ncbi.nlm.nih.gov/15016991/)). All five brains responded *synchronously* to shifts in scenery, dialogue, and plot, "as if operated by the same neural story-watching program." A movie drives a shared, predictable structure of neural activity across unrelated people. An event log does the same — if you build the narrative layer to expose it.

**Memory is an edit, not a recording.** In a follow-up study, Janice Chen's group scanned people watching — then recalling — the first episode of *Sherlock* ([Chen et al., *Nature Neuroscience* 20, 2017](https://doi.org/10.1038/nn.4450)). The neural pattern during recall was *not* the same as during viewing. The differences were systematic and similar across viewers: the brain was editing the footage as it stored it, compressing the episode into its causal shape. What you remember is the cut, not the rushes.

**The brain carves experience into scenes and matches them against scripts.** Christopher Baldassano's re-analysis of the same data showed the brain segmenting the episode into scenes — striking shifts in prefrontal-cortex activity lined up exactly with perceived scene changes. "At a scene change, you recognize, okay, now a new thing is starting." This is the known psychological phenomenon of **event segmentation**: people draw boundaries at shifts in place, time, or situation, and those segments are "a natural product of the brain." And the brain does not build these segments from scratch: back in the 1970s, psychologists hypothesized that people hold *scripts* for familiar situations (the grocery-store run: take a cart, walk the aisles, queue, pay). Baldassano puts the scale on it — "an adult brain holds hundreds of thousands of scripts for expected scenarios, a huge neuronal library of story outlines crafted from experience" — and sums up the operating principle in one sentence: **"Your brain is not built to just record the pixels that are coming in from the screen or from the movie. You are trying to match this into something that you know."**

One more result, the one that matters most for similarity: in Baldassano's studies, the more faithfully an experience matched a stored script, the more detailed the memory of it. Matching against a known story is not a cheap shortcut — **it is the mechanism of comprehension and retention**. Jesse Andrews, Pixar co-writer of *Luca*, puts the whole research program into one line: "When we say 'narrative,' all we are really saying is 'simplification.' It's a simplification in the name of understanding."

## Event sourcing gives you footage, not a film

None of this says event sourcing is wrong. Event sourcing is the right decision *precisely because* it keeps the raw footage: an immutable, replayable, auditable record of everything that happened, in order, with no interpretive damage done on the way in. That is the correct end of the camera.

But the discipline of storing only facts has a cost: the facts do not interpret themselves. An event log is the unedited rushes — frame-accurate, and without a cut. Consider two events:

```text
2026-08-01 09:14  order.cancelled   order=o-4821  reason=changed-mind
2026-08-01 09:14  order.cancelled   order=o-4821  reason=changed-mind
```

Identical payloads, seconds apart, same causal *content* — but one is a customer wavering and the other is a system double-firing after a retry. The events cannot tell you which is which. The story can: *the same story* (one episode, one protagonist, one intent) vs *two stories colliding* (a genuine cancellation racing a replay). This is the Heider–Simmel lesson in production: identical shapes, different stories, and the shapes alone are mute.

Systems that stop at the log are asking the frames to be the movie. Replay, audit, point queries, outbox delivery — these are frame-level operations, and they are necessary. But they are the *cinematographer's* operations. The *director's* operations — What happened to this customer over the quarter? Which run of this process fails, and where in the arc? Which past incidents is this incident like? — are story-level, and no amount of frame-level tooling answers them.

## How to cut the film: making stories from events

The brain's three mechanisms map cleanly onto three engineering mechanisms, all of them projections over the log:

| Brain mechanism (the science) | Log mechanism (what to build) |
|---|---|
| Event segmentation — scene changes marked by prefrontal shifts | Correlation and causation ids, saga/process instances, aggregate lifecycles, session boundaries |
| Scripts — a library of hundreds of thousands of story outlines | A story-type taxonomy (archetypes) with typed slots, learned from the log and curated by the domain |
| Viewpoint-guided extraction — critic vs wedding planner write different stories from the same restaurant scene | Keyed story projections: one per viewpoint, or stories with explicit viewpoint labels |

Concretely, a **story** is a bounded, goal-directed episode cut from the event stream: a start trigger, a set of actors, a causal chain of events, an outcome, and a time span. The cutting is a projection — it consumes the log and writes story records, idempotently and replayably, and it can always be rebuilt. The log stays the source of truth; the story store is disposable interpretation.

```go
// Story is the unit of narrative: a bounded, goal-directed episode
// cut from the event log. It is always a projection — re-derivable
// from events, never the source of truth.
type Story struct {
	ID       string            // correlation id: the scene boundary
	Type     string            // archetype: "customer-churn", "order-fulfillment"
	Actors   []string          // protagonists: aggregate or entity ids
	Goal     string            // what the episode was trying to do
	Events   []string          // event ids in causal order: the plot
	Outcome  string            // success, failure, abandoned
	Span     [2]time.Time      // scene boundaries, from segmentation
	Features map[string]float64 // numeric features for story-level similarity
}
```

Where do the scene boundaries come from? In order of preference:

1. **Explicit process boundaries** — the ones already in your log and underused: saga and process-manager instance ids, correlation ids, causation ids, transaction and outbox batch ids. A saga instance *is* a scene; its id is the prefrontal shift.
2. **Aggregate boundaries (DDD)** — an aggregate's lifecycle is a minimal story: one protagonist, a coherent goal, a causal chain of state transitions. Aggregate id = protagonist id.
3. **Implicit segmentation** — for streams with no process marker: state-transition bursts, idle gaps, session boundaries, change-point detection on state and rate. This is the fuzzy end of the spectrum; prefer 1 and 2 wherever the domain gives them to you.

Each story closes when its causal thread terminates — an outcome event, a timeout, a boundary marker. The reducer is boring in the right way:

```go
// Reduce folds the log into stories, grouping by scene boundary and
// advancing the plot one event at a time. Deterministic and replayable.
func Reduce(events []Event) []Story {
	stories := map[string]*Story{}
	for _, e := range events {
		s, ok := stories[e.CorrelationID]
		if !ok {
			s = newStory(e) // start trigger: the scene opens
			stories[e.CorrelationID] = s
		}
		s.Events = append(s.Events, e.ID)
		s.Outcome = advance(s.Outcome, e) // the plot moves, frame by frame
	}
	return toSlice(stories)
}
```

The story projection is where you are allowed to be lossy — the summary field, the typed slots, the outcome label are the "edit." The log keeps the pixels; the story keeps the cut.

## Why "story similar to story" beats "event similar to event"

This is the claim worth arguing carefully, because it inverts a very natural instinct. Most teams building retrieval, dedup, or anomaly detection over event-sourced data start with **event-to-event similarity** — embed the events, or hash them, or cluster them. The argument here is that the level is wrong. Similarity is only meaningful at the level where causality and outcomes live, and that level is the story. Five reasons, in increasing order of importance.

**1. Points cannot be similar to each other; paths can.** An event is a point in time; a story is a path through the state space. Point-level similarity has no access to the path — it sees the type and the payload and nothing else. Two `payment.failed` events are *trivially* identical syntactically and *causally unrelated* semantically; event similarity is the worst of both worlds. Story similarity, by contrast, is structural: it compares arcs, actors, obstacles, and outcomes. It is the difference between saying "two triangles moved the same way" and "these are both the love-triangle story."

**2. Composition does not commute.** Similarity does not survive the trip down. If events `a₁` and `a₂` are similar, and `b₁` and `b₂` are similar, nothing follows about the stories `a₁→b₁` and `a₂→b₂` — the same-type events can appear in opposite stories (a churn story and a retention story both contain `email.sent`), and different-type events can instantiate the same archetype (churn via failed renewal, churn via support hell). Event-level indexing gives you a bag of matched facts with no causal thread; story-level structure decomposes into event structure, so it gives you the thread and the facts together.

**3. The brain's own answer: match against what you know.** Baldassano's line is the whole argument: "you are trying to match this into something that you know." The brain does not compute similarity between pixels; it matches experiences against stored story outlines, and the fidelity of that match is what predicts memory. When you retrieve "a story similar to this story," you retrieve a *template with a causal structure* — and a template is actionable: you can predict the next scene, replay the remedy that worked last time, generalize the counterfactual ("this churn story deviates from the archetype at step 3"). That is case-based reasoning, and it is the mechanism the brain runs. Retrieving "an event similar to this event" retrieves a fact with no affordance. A fact does not tell you what comes next; a story does.

**4. Deviation is sequence-level, not point-level.** Anomaly detection over events flags rare events — and rare events are mostly benign noise, while the dangerous event is usually individually ordinary. The signal lives in the *arc*: "this incident is the standard failure story except the retry storm happens at step 4 instead of never." You can only say that once you have a story to deviate from. Story-level similarity turns anomaly detection into *deviation from archetype* — structural, explainable, and precisely the "neural scaffold for a familiar scenario" the brain relies on when it encounters a new-but-similar situation.

**5. Outcomes are story-level, so similarity must be too.** Churn, conversion, incident resolution, SLA breach — every metric you actually care about is an outcome of a story, not an attribute of an event. Similarity for the purpose of prediction, retrieval, or root-cause should live at the level of the thing you are predicting. Indexing at the event level optimizes for a level that has no outcome attached to it; you are comparing frames while the scoreboard lives in the cut.

There is also a practical, current-LLM reason: retrieval-augmented reasoning over event logs drowns in grain. Stories are the right chunk — dense, causally coherent, few. Feeding an agent ten thousand raw events and asking it to find the similar incident is like asking a human to find the love triangle in the coordinates. Feed it the story records and the match is immediate, because the representation matches the reasoning machinery — the same reason the brain edits memory before storing it.

## Where the story layer can mislead

The brain's strategy is not free, and neither is the engineering one. The lossy edit is the point, but it is also the failure mode.

- **The narrative fallacy.** A story layer can produce confident, coherent, *wrong* explanations — the storytelling brain has a documented talent for this (rumination, hindsight bias, false confidence in the edited cut). The defense is structural: **the log is ground truth, the story store is disposable**. Stories must always be re-derivable from events; when a story and the log disagree, the log wins. The moment the story becomes the source of truth, you have turned an append-only audit trail into a mutable interpretation — the exact thing event sourcing exists to prevent.
- **Viewpoint decides the story.** The wedding-planner result — subjects given a role remembered a different story from the same restaurant scene — cuts both ways. A projection with no explicit viewpoint smuggles one in silently. Decide the viewpoints you index (customer journey, process health, incident response) and label them.
- **Archetypes go stale.** The script library is built from experience and the domain drifts. Story types need the same versioning and deprecation discipline as any schema, or the similarity search quietly starts comparing against obsolete templates.
- **Segmentation of boundary-less streams is fuzzy.** Idle-gap and change-point heuristics overfit to noise. Where the domain gives you saga ids, use them; treat implicit segmentation as the last resort, and audit its boundaries.

The tradeoff to remember is the same one the brain accepts: the edit loses detail and buys comprehension. That is not a bug in your data model — it is the abstraction that makes generalization possible at all.

## Simplification in the name of understanding

Thirty thousand years of cave painting, ten thousand years of oral storytelling, seventy-five years of the Heider–Simmel film, a decade of decoding movie-watching brains — they all converge on one design principle: segment experience into scenes, edit the footage into a story, and match new experience against the stories you already know. Event sourcing gives you the footage — the unedited, immutable truth. That is worth keeping, exactly as it is.

But the footage is not the film, and the film is where the meaning lives. Cut your log into stories, project them beside it, and compare stories to stories. The brain has spent its entire evolutionary budget proving that is the abstraction at which the world becomes legible.
