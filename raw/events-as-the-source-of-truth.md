---
title: "Events as the Source of Truth: Why Blockchains Are Relevant"
date: 2026-09-07
slug: events-as-the-source-of-truth
summary: "Oscar van der Leij's event-sourcing essay makes the case that every state change should be an immutable event in an append-only log — a full, tamper-evident history your database keeps, not a current-state snapshot that forgets how it got there. Read as an engineering claim, the essay's own adjectives — immutable, tamper-evident, append-only — are promises a single-operator store cannot enforce, only honor. That gap is exactly where blockchains become relevant: a chain is event sourcing run by keepers who do not trust each other, where append-only stops being a convention and becomes a structural property. This post maps the essay's principles onto chains, walks a hash-chained event log in Go, and ends with the trust-boundary test for when you need consensus, a hash anchor, or just Postgres."
tags: event-sourcing, blockchain, audit-trail, append-only, cqrs, projections, database-architecture, immutability, trust, essay
---

A governance officer asks: "Can you show us a full audit trail of who changed what, and when?" The development team's answer, writes Oscar van der Leij in [_Event Sourcing: A History Lesson Your Database Actually Wants_](https://www.architectviewmaster.com/blog/event-sourcing-a-history-lesson-your-database-actually-wants/), is always some version of "we store the current state, so we can see what it is now, but not how it got there." His essay is the standard fix: Event Sourcing, the pattern where an application stores every state change as an immutable event in an append-only log and derives current state by replaying those events. He recommends it, he says, "not as an architectural curiosity, but as a practical response to a requirement I kept encountering: full, tamper-evident, queryable history of every state change in the application."

Notice the adjectives in that requirement. *Tamper-evident.* *Immutable.* *Append-only.* The essay's whole architecture is built to deliver them — and yet the implementation it walks through, a local SQLite event store in C#, cannot actually enforce any of them. It can only honor them. That gap between the property an audit trail claims and the property a database can enforce is not a flaw in the essay; it is the essay's unfinished thought. When the history has to be believed by someone who does not trust the system that wrote it, the append-only log needs a keeper the writer does not control. That is what a blockchain is. A blockchain is event sourcing executed by keepers who do not trust each other, where append-only stops being a convention and becomes a structural property.

This post reads the essay carefully, then follows the sentence it leaves hanging.

## What the essay teaches

The essay's core move is to reframe persistence. A traditional database "stores the current truth": update the address and the old one "disappears into the digital ether." That works until someone asks how the system got where it is. The essay's analogy is a meeting-room whiteboard: you erase the old content and write the new state, so the board can always show the current truth and never the meeting. Event Sourcing "is like replacing that whiteboard with a roll of paper that keeps feeding forward. Nothing gets erased." The current state stays visible; so does everything that led to it. "Your database becomes a ledger, a chronicle, a time machine."

Four principles carry the pattern:

- **Events as the source of truth.** State is derived; events are canonical. Each event records what happened (a business fact in the past tense: `OrderPlaced`, `PaymentProcessed`), when it happened, who made it happen, and the details needed to replay it. "These are historical records, as unchangeable as yesterday's weather." And events are the contract you make with the future: "once written, they never change."
- **The append-only log.** Events are never updated or deleted. Made an accounting error? "Don't change the original transaction. Add an 'AccountingCorrectionApplied' event." The mistake and its correction both survive, "which is often exactly what auditors and regulators want to see."
- **State reconstruction through replay.** Current state is not (primarily) stored; it is rebuilt by replaying the event stream. The essay's C# aggregate is a `switch` over event types — `OrderPlaced` sets `Status = Placed`, `OrderShipped` sets `Shipped` — which is a state machine folded into code, the same shape this site traced through the workflow-engine argument in [Every workflow is an FSM](https://blog.hackspree.com/#every-workflow-is-an-fsm).
- **Temporal queries.** Because the full history exists, you can ask what the state was at any earlier time — replay to a date, not just to now. "This temporal dimension transforms debugging, auditing, and analytics from guesswork into precision."

The essay is equally clear about the costs, which is what makes it a good reference rather than a manifesto. Replaying ten thousand events on every read is expensive, so you take periodic **snapshots** — current-state captures that serve as replay starting points, tagged with the event version they include. Reading from the log for every query is expensive, so you build **projections** — denormalized read models that answer queries instantly and stay in sync with the stream: "Same data, two different shapes. That is CQRS in practice." And you should not apply the pattern everywhere: simple CRUD without audit requirements, teams without event-driven experience, and hot read paths are all reasons to skip it. The deciding question, he says, is the one he now asks at the start of every project: "Will anyone ever need to prove what happened, who did it, and when?"

His closing claim is the load-bearing one: "The audit trail is a structural property of the system, and retrofitting it onto a CRUD database after the fact is far more expensive than building it in from the start."

## The property the demo does not enforce

Now look closely at what the essay's own implementation delivers. The demo's `EventStore` is a SQLite table with an `INSERT`-only API and an index on `(AggregateId, EventVersion)`. The audit trail prints straight from that table — "there is no separate logging system, no trigger, no after-the-fact reconstruction: the history is the data." All true, and all true *by convention*. The table permits `UPDATE` and `DELETE` like any other table. Nothing in the schema, the file, or the machine stops an administrator with write access — or an attacker with the application's credentials, or a `DROP TABLE` followed by a restore from an older backup — from rewriting, deleting, or truncating the history. "Immutable" here means "the application agreed not to mutate," and "tamper-evident" means "tampering would require violating an application-level rule," not "tampering would be detected by the system." A hash chain would make even local tampering detectable:

```go
package eventlog

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
)

// Event is one immutable fact: what happened, when, and by whom. Actor and
// At are self-asserted — the event says who did it and when; nothing yet
// proves it.
type Event struct {
	Type     string          `json:"type"`
	Actor    string          `json:"actor"`
	At       int64           `json:"at"`
	Data     json.RawMessage `json:"data"`
	PrevHash string          `json:"prevHash"` // digest of the previous event
	Hash     string          `json:"hash"`     // digest of this event
}

// digest covers every field except Hash, so an event never signs itself.
func digest(e Event) string {
	body, _ := json.Marshal(struct {
		Type     string          `json:"type"`
		Actor    string          `json:"actor"`
		At       int64           `json:"at"`
		Data     json.RawMessage `json:"data"`
		PrevHash string          `json:"prevHash"`
	}{e.Type, e.Actor, e.At, e.Data, e.PrevHash})
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[:])
}

// Append links the new event to the log head and inserts it. The store only
// ever INSERTs: correcting a mistake means appending a compensating event.
func Append(s *Store, e Event) (Event, error) {
	head, err := s.Head()
	if err != nil {
		return Event{}, err
	}
	e.PrevHash = head
	e.Hash = digest(e)
	return e, s.Insert(e)
}

// Verify walks the log in order and recomputes every digest. Editing any
// event, deleting any link, or reordering any pair breaks the chain here —
// tampering becomes detectable, though nothing yet stops a determined
// operator from rewriting the whole file and re-chaining it.
func Verify(s *Store) error {
	prev := ""
	for _, e := range s.All() { // All returns events in insertion order
		if e.PrevHash != prev || e.Hash != digest(e) {
			return fmt.Errorf("event log tampered at %s", e.Hash)
		}
		prev = e.Hash
	}
	return nil
}

// Store is deliberately boring: rows are appended in sequence and never
// updated or deleted. Keeping that discipline is the application's job; the
// schema cannot enforce it — which is precisely the essay's tension.
type Store struct {
	mu   sync.Mutex
	rows []Event
}

func (s *Store) Head() (string, error) {
	if len(s.rows) == 0 {
		return "", nil
	}
	return s.rows[len(s.rows)-1].Hash, nil
}

func (s *Store) Insert(e Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rows = append(s.rows, e)
	return nil
}

func (s *Store) All() []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Event, len(s.rows))
	copy(out, s.rows)
	return out
}
```

The hash chain upgrades the story from "we promised not to edit" to "an edit breaks the chain and Verify fails." That is a real improvement — it is what turns an audit *log* into an audit *trail* — but notice what it does not do. The actor field is still self-asserted: any code that can call `Append` can claim to be anyone. And the whole chain lives in one file under one operator's control, so a determined operator can rewrite the history *and* re-chain it: delete an embarrassing event, recompute every subsequent hash, and present the file as untampered. Verification detects amateur tampering. It cannot detect tampering by the party who owns the store — and the party who owns the store is exactly the party whose history an auditor, a regulator, or a legal team is trying to verify.

This is the sentence the essay leaves hanging. The author asked for "tamper-evident, queryable history" and built "history with a hash chain your own team could still rewrite." The requirement was never fully satisfiable by a single-operator database, because the threat model of a compliance audit includes the operator. "Will anyone ever need to prove what happened, who did it, and when?" — prove *to whom*? If the answer is "to people who trust us," the SQLite store is fine. If the answer is "to people who might not," no amount of application discipline closes the gap. The history must be kept, or at least vouched for, by someone other than the writer.

## Blockchain is event sourcing across a trust boundary

Every piece of the essay's architecture has a counterpart in a blockchain, and recognizing the mapping is the fastest way to see why the technology is relevant to this pattern rather than adjacent to it:

- **Events as the source of truth.** A blockchain does not store its "current truth" as an authority separate from history. Its state is derived by replaying the log: every full node re-executes the ordered history of blocks to rebuild the state, exactly as the essay's `Order.FromEvents` re-executes events. The block is `OrderPlacedEvent`; the account balance is `order.Status`; the state at height N is the temporal query.
- **Append-only as a structural property.** This is the big one. Blocks are hash-linked (each commits to its predecessor, and through it, to the entire history), and the log is replicated across validators who do not trust each other. Editing one block breaks every descendant hash and every honest replica detects it. Rewriting history therefore requires rewriting it on a majority of independent keepers — corrupting the consensus, not one database. Immutability stops being a convention the application honors and becomes a property the network enforces. The essay's whiteboard-with-a-paper-roll becomes a paper roll nobody owns alone: everyone holds a copy, so nobody can privately erase.
- **Who and when.** The essay's event carries a `UserId` string that any writer can set. On a chain, the "who" is a public key: the event is only accepted if it carries a signature the network can verify against that key, and the writer's right to append at all is decided by the chain's rules, not by whoever holds the database connection. The "when" is the consensus-ordered sequence — the block height — not an application clock. The actors and timestamps an auditor reads are not self-asserted; they are asserted under rules the auditor can check.
- **Temporal queries.** Historical state at any height is a first-class concept, and since every block commits to the state root that resulted from replaying up to that point, the answer to "what did this look like six months ago" is a commitment that can be verified against a stored root — not a guess reconstructed from whatever events happened to survive.
- **Snapshots and projections.** A chain's per-block state root is the essay's snapshot: current state plus the version (height) it includes, used as a replay starting point. Wallets, explorers, and indexers are the essay's projections — read models derived from the log, answering queries instantly, eventually consistent with it. "Same data, two different shapes" is not just CQRS; it is the standard shape of blockchain infrastructure.

So the relevance thesis is precise: **event sourcing makes one organization's history a first-class citizen; a blockchain makes history first-class across organizations.** Everything the essay wants inside a single trust boundary — complete history, compensating events, temporal queries, derived read models — a chain provides across a trust boundary, at the price of consensus. When the events in question all belong to one company and the proof's audience trusts that company, Postgres plus a hash chain is the right-sized answer and a chain is waste. When the events are jointly produced by parties that must be able to audit each other — two firms settling trades, a hospital and an insurer disputing a claim, a supplier and a buyer arguing about provenance, two organizations' agents executing a contract — each party keeping its own append-only log recreates the original problem at the boundary: in a dispute, each side can rewrite its own history, and both sides know it. A shared log with no single keeper is the only history both sides can be asked to accept. That is what blockchains are for. The article's own audit question — "who changed what, and when" — is answered by event sourcing within a company, and by a chain between companies.

This is the same claim this site made about blockchains and agentic software from the mechanism side: [blockchains matter because they make rules visible and programmable rather than opaque](https://blog.hackspree.com/#why-blockchains-matter-again-for-agentic-software). The event-sourcing lens adds the evidential side. A blockchain is not primarily a token ledger; it is a shared, append-only, tamper-evident history that no participant controls, plus rules about who may append. If your problem is "we need a history that outsiders will believe," you have described a blockchain-shaped requirement whether or not you use the word.

## The trust-boundary test

The honest engineering question is not "should I use event sourcing" or "should I use blockchain" but *where the boundary of trust runs*. Three configurations cover most of the design space:

**Inside one trust boundary — a plain event store.** The essay's SQLite/Postgres design, optionally hash-chained. Tamper-evident to your own auditors in the amateur-tampering sense, cheap, fast, boring. This is where most of the essay's use cases belong: your orders, your claims, your inventory, your auditors inside your org.

**History you own, proof you share — hash anchoring.** When outsiders must be able to verify that your history has not been rewritten since a certain time, you do not need to move your events onto a chain; you need to commit their fingerprints somewhere neither you nor your auditor controls. Periodically compute a digest of the current log head and append it to a public chain (or a shared ledger). From that moment on, any later rewrite of your local log breaks the anchored digest, and the anchor is timestamped and preserved by people who are not you. This is the pattern behind RFC 3161 timestamping and [OpenTimestamps](https://opentimestamps.org/): the log stays in your database, the *proof* leaves it. It is the cheapest way to buy "tamper-evident to outsiders" without paying for consensus on every event. It works because the essay's snapshots give you a natural anchor point: every snapshot is a digest of the history up to a version, and that digest is exactly what you notarize.

**History jointly produced and jointly audited — a shared ledger.** When multiple parties must *append* to the same history without trusting each other's databases — settlement, provenance, interorganizational workflows, agents transacting across companies — an anchor is not enough, because each party still maintains its own log and can still disagree about what the other appended. Then the log itself must be the shared artifact: a permissioned shared ledger among known counterparties, or a public chain where the value of neutral history justifies the cost. The essay's "when to think twice" list has a direct analogue here: consensus is slow and expensive, so this configuration is for histories whose integrity is worth more than their throughput — the settlement layer, not the shopping cart.

## Why this matters for agents

The reference essay was written about humans asking humans for audit trails. The reason the topic keeps compounding is that the askers are increasingly programs. Autonomous agents emit events continuously — every tool call, every decision, every transfer — and when two agents from different organizations interact, the situation is the two-company dispute at machine speed: each side's internal log is self-serving, neither side can inspect the other's database, and there is no phone call and no court deadline to force the question. This site has argued that [the event log, not the plan, is the honest memory of an agent](https://blog.hackspree.com/#stories-from-events): graphs describe intentions, logs describe executions. The blockchain relevance follows from that argument's logical endpoint — when the execution spans organizations, the log that records it must not be owned by either side, or neither side can trust it. An agent's audit trail is only as honest as the party that keeps it; a shared append-only history is the only form in which "what happened" survives contact between agents that do not trust each other. The essay's deciding question — "will anyone ever need to prove what happened, who did it, and when?" — becomes, in an agentic economy, a question asked at machine frequency by counterparties that cannot take anyone's word. The answer event sourcing gives within a system, and blockchains give between systems, is the same answer the essay gives its governance officer: the history is the data. The difference is who gets to rewrite it.

## Key insight

The essay is right that state without history is a whiteboard, and that the audit trail should be "a structural property of the system" rather than a retrofit. The unfinished sentence is *whose* system. "Immutable," "append-only," and "tamper-evident" are properties an application can promise but only a keeper outside the writer's control can enforce. Within one trust boundary, event sourcing — with or without a local hash chain — is the right and complete answer, and adding consensus would be over-engineering. Across a trust boundary, the same pattern needs its log kept by people who are not the writer, which is the definition of a blockchain: event sourcing whose append-only property is enforced by replication and consensus instead of by convention, and whose actors sign instead of assert. The spectrum between the two — plain log, anchored log, shared ledger — is chosen by one question, which is the essay's own question sharpened: will anyone ever need to prove what happened, who did it, and when — *to someone who does not trust you?* If no, Postgres. If yes, an anchor. If the parties must write the history together, a chain. The history lesson your database actually wants ends where your trust boundary does.

---

**References:**

- van der Leij, O. [_Event Sourcing: A History Lesson Your Database Actually Wants_](https://www.architectviewmaster.com/blog/event-sourcing-a-history-lesson-your-database-actually-wants/) — Architect Viewmaster, 10 June 2026. The main reference; all unattributed quotes above are from this essay (the whiteboard-and-paper-roll analogy, the audit question, the CQRS and snapshot strategies, and the deciding question in the conclusion).
- Fowler, M. [_Event Sourcing_](https://martinfowler.com/bliki/EventSourcing.html) — martinfowler.com bliki, 2005. The canonical definition of the pattern: events as the primary record of a system's state.
- [_OpenTimestamps_](https://opentimestamps.org/) — the client-side hash-anchoring scheme (Bitcoin blockchain as a public timestamping calendar), the canonical example of "log stays with you, proof leaves."
- IETF. [_RFC 3161: Internet X.509 Public Key Infrastructure Time-Stamp Protocol_](https://www.rfc-editor.org/rfc/rfc3161) — the standardized form of the same anchoring idea.
- Related: [Make Stories from Events](https://blog.hackspree.com/#stories-from-events) — the event log as honest memory, and why raw events need a narrative layer; the complementary half of this argument.
- Related: [Blockchains for Agentic Software](https://blog.hackspree.com/#why-blockchains-matter-again-for-agentic-software) — blockchains as mechanisms with visible, programmable, auditable rules; the rules side of the same coin.
- Related: [Every workflow is an FSM. Not every FSM is a workflow.](https://blog.hackspree.com/#every-workflow-is-an-fsm) — state reconstruction as replay, event history as the state machine's truth, and the versioning discipline shared by event-sourced aggregates and replay engines.
- Related: [Mechanism Design and Network Economics for Agentic Markets](https://blog.hackspree.com/#mechanism-design-and-network-economics-for-agentic-markets) — rules as the design surface when autonomous parties must coordinate without trusting each other.


Engineering treats every system as a set of dependencies and asks which ones you can afford to own. Event sourcing is the discovery that history is not a dependency — it is the asset you delete when you store only current state. The append-only log is the engineering answer to "what happened," and the blockchain is the engineering answer to "what happened, when the people asking do not trust the people who kept the record." Same pattern, wider boundary. The boundary is the specification.


> Immutable means nothing unless someone you do not control can check it. Event sourcing makes history the source of truth inside one trust boundary; a blockchain is what the same idea becomes across one. The question is never "append-only or not." It is who else holds the roll of paper.
