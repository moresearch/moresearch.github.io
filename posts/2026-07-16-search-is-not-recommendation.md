---
title: Search Is Not Recommendation — And Why the Distinction Matters More Than Ever
date: 2026-07-16
summary: Search and recommendation are often conflated, especially now that LLMs blur the boundary. This post traces each field's history, techniques, and their convergence at Netflix, Spotify, DoorDash, and Pinterest — arguing that the distinction remains essential for building systems that work.
tags: [search, recommendation, information-retrieval, collaborative-filtering, llm, netflix, machine-learning, survey]
---

Search and recommendation are often described as "two sides of the same coin." Both match users with items. Both rank results. Both drive discovery. Netflix, Spotify, DoorDash, Pinterest, and Airbnb build both. The phrase appears in conference papers and engineering blog posts alike.

It is also wrong — or at least, incomplete enough to be dangerous.

Search and recommendation are different problems. They emerged from different research communities, solved different user needs, developed different mathematical frameworks, and measure success differently. Confusing them leads to systems that are bad at both: search results that drift into irrelevance under the weight of personalization, and recommendation feeds that fail to respect explicit intent.

The arrival of Large Language Models makes the distinction sharper, not softer. When a single transformer can be fine-tuned to retrieve documents, rank products, generate recommendations, and explain its reasoning, the temptation is to collapse everything into one model, one pipeline, one problem. Netflix's UniCoRn and Spotify's agentic search router do something close to that — but they succeed precisely because they respect the distinction, not because they ignore it.

Consider three everyday scenarios:

1. You walk into a library and ask the librarian, "Where are the books on the history of the Byzantine Empire?" The librarian walks you to a specific shelf. That's **search**.

2. You walk into the same library and the librarian says, "I notice you've been reading a lot about medieval trade routes. We just got a new book on the Silk Road you might enjoy." That's **recommendation**.

3. You walk into the library, glance around, and say, "I'm not sure what I want. Something historical but not too heavy, maybe with a good story?" The librarian pauses, then suggests three books and asks which one sounds right. That's **the boundary** — where search and recommendation blur.

These three scenarios feel different. They engage different cognitive postures, demand different system designs, and tolerate different kinds of errors. The first failure looks like "this librarian doesn't know where the books are." The second looks like "this librarian doesn't understand my taste." The third looks like a conversation — and it's exactly the mode that LLMs are making possible at scale.

This post traces the history, the mathematics, the architectures, and the convergence. It uses Netflix as the central case study — the company that runs perhaps the world's most sophisticated search *and* recommendation systems side by side — and draws on engineering blogs from Spotify, DoorDash, Pinterest, Airbnb, and others to show how the best teams think about the boundary.

---

## The Fundamental Distinction

Before the history, the definitions. The difference between search and recommendation can be stated in one sentence:

> **Search is query-driven retrieval. Recommendation is preference-driven filtering.**

But that sentence unpacks into a richer set of dimensions:

| Dimension | Search | Recommendation |
|---|---|---|
| **User intent** | Active: user formulates a query | Passive: system surfaces items unprompted |
| **Input representation** | Short, explicit query text | Implicit user profile (history, behavior, context) |
| **Core operation** | Query–document matching | User–item matching |
| **Information need** | Known: "I want to find X" | Unknown: "Show me what I might like" |
| **Evaluation** | Precision, recall, NDCG, MRR | RMSE, AUC, CTR, retention, discovery |
| **Cold start** | Query/document side (new terms, new docs) | User/item side (new users, new items) |
| **Serendipity** | Undesirable (should return what was asked for) | Desirable (should surface unexpected gems) |
| **Personalization** | Optional; applied at re-ranking | Core to the entire pipeline |
| **Theoretical roots** | Information Retrieval (library science, linguistics) | Collaborative Filtering (HCI, machine learning) |

The key insight is that **the user's posture differs**. In search, the user leans forward — they have a goal. In recommendation, the user leans back — they are open to suggestion. The system architectures that serve these two postures optimize for fundamentally different things.

Peter Norvig captured the consequence of this difference with characteristic precision. When asked about the shift from retrieval to proactive assistance, he noted:

> "With information retrieval, anything over 80% recall and precision is pretty good — not every suggestion has to be perfect, since the user can ignore the bad suggestions. With assistance, there is a much higher barrier."

The bar is different because the contract is different. In search, the user is the arbiter — they scan results, skip what's irrelevant, and reformulate if necessary. The system is a tool the user *operates*. In recommendation, the system is the arbiter — it decides what the user sees. If it's wrong, the user doesn't skip a result; they lose trust in the system itself. This is why the distinction matters in production: getting search wrong loses a query. Getting recommendation wrong loses a user.

Belkin and Croft put this more formally in their 1992 survey of information filtering: search systems index items (documents) and match against ad-hoc queries from anonymous users. Filtering (recommendation) systems index user profiles and match against incoming items over time, for known users whose preferences are modeled [1].

The theoretical frameworks emerged from different communities — library science and computational linguistics for search, human-computer interaction and machine learning for recommendation — and those different origins shaped everything that followed.

---

## Part I: The History of Search (Information Retrieval)

### The Boolean Era (1960s–1970s)

The earliest computerized search systems were built for libraries and intelligence agencies. They used Boolean logic: a query was a logical expression of terms, and documents either matched or did not. There was no ranking — just set intersection. This worked for trained librarians who could craft precise queries. It failed for everyone else.

The fundamental problem was what we now call the **vocabulary mismatch problem**: users describe their information needs using different words than authors used to write. Searching for "tragic love story" returns nothing for Shakespeare's "star-crossed lovers." Boolean retrieval is blind to semantics.

Karen Spärck Jones, who invented inverse document frequency (IDF) — the weighting scheme that would become half of TF-IDF — understood this better than anyone. In a 1999 reflection on the field she helped create, she wrote: "Classical document retrieval thus falls in the class of AI tasks that assist the human user but cannot, by definition, replace them." Even in 1999, she saw that the irreducible ambiguity of human language meant retrieval would always be an approximation. The vocabulary mismatch problem is not a bug to be fixed — it is the fundamental condition of language, and every generation of retrieval technology, from TF-IDF to BERT to RAG, is an attempt to narrow the gap without ever closing it.

### The Vector Space Model (1975)

Gerard Salton and his collaborators at Cornell changed everything. In 1975, they proposed the **Vector Space Model (VSM)**, representing both documents and queries as sparse vectors in a high-dimensional term space [2]. A document was no longer a set of words — it was a point in a mathematical space where similarity could be computed.

The breakthrough was the weighting scheme: **TF-IDF**. Term Frequency (TF) measured how often a term appeared in a document — more frequent terms are more important. Inverse Document Frequency (IDF) discounted terms that appeared everywhere — common words are less discriminative. The product gave each term a weight that captured both local importance and global rarity.

```
TF-IDF(t, d, D) = tf(t, d) × log(N / df(t))
```

Where `tf(t,d)` is the frequency of term `t` in document `d`, `N` is the total number of documents, and `df(t)` is the number of documents containing `t`.

Similarity between query and document was computed via cosine similarity — the angle between their TF-IDF vectors. This worked remarkably well, was computationally efficient with inverted indices, and remains a strong baseline fifty years later.

The VSM didn't solve the vocabulary mismatch problem, but it gave us a mathematical framework for thinking about it.

### Probabilistic Retrieval and BM25 (1970s–1990s)

Parallel to Salton's vector space work, a different tradition emerged: probabilistic relevance models. Stephen Robertson, Karen Spärck Jones, and collaborators at City, University of London asked a different question: given a document and a query, what is the probability that the document is relevant?

The answer was the **Binary Independence Model** and, eventually, the **Okapi BM25** ranking function. Published in 1994, BM25 became the dominant text retrieval function for the next two decades [3]. Its formula:

```
BM25(D, Q) = Σ IDF(qᵢ) × [f(qᵢ, D) × (k₁ + 1)] / [f(qᵢ, D) + k₁ × (1 − b + b × |D| / avgdl)]
```

The critical innovations over TF-IDF were:
1. **Diminishing returns for term frequency**: a term appearing 10 times is not 10× as important as appearing once. The saturation curve (controlled by `k₁`, typically 1.2–2.0) captures this.
2. **Document length normalization**: longer documents aren't inherently more relevant. The `b` parameter (typically 0.75) controls how much length is normalized away.

BM25 is derived from a probabilistic relevance framework, not heuristics. It remains the default scoring function in Elasticsearch and Lucene — which means it powers the search infrastructure at Netflix, Spotify, DoorDash, and most of the web.

### PageRank and the Web (1998)

When the web arrived, content-based retrieval hit a wall. Any document can link to any other. Brin and Page's insight at Stanford was that the link graph itself carried relevance information [4]. PageRank treated links as votes: a page linked to by many important pages is itself important. The score propagates through the graph via a random-surfer model:

```
PR(p) = (1 − d) / N + d × Σ PR(pᵢ) / L(pᵢ)
```

Where `d` is the damping factor (typically 0.85), `L(pᵢ)` is the number of outbound links, and the sum runs over all pages linking to `p`.

PageRank didn't replace content-based retrieval — it augmented it. A modern search engine computes hundreds of features (BM25, PageRank, proximity, freshness, click-through rate, spam score) and feeds them into a learned ranking function.

### Learning to Rank (2000s–2010s)

The next leap was treating ranking as a supervised machine learning problem. Given a query, a set of candidate documents with relevance labels, and hundreds of hand-crafted features, learn a function that orders them optimally.

The dominant approaches were:
- **Pointwise**: predict relevance score for each query–document pair independently (regression, classification)
- **Pairwise**: learn to order pairs of documents correctly (RankNet, RankSVM)
- **Listwise**: optimize the entire ranked list directly (LambdaRank, LambdaMART)

**LambdaMART** — a gradient-boosted decision tree trained with a listwise LambdaRank objective — became the industry standard. By 2010, every major search engine used some variant of this approach. Microsoft, Yahoo, and Google published extensively on it [5].

### The Neural Turn (2013–2020)

Three papers changed search again:

1. **Word2Vec** (Mikolov et al., 2013): showed that dense, low-dimensional word vectors could capture semantic relationships. "king − man + woman ≈ queen" entered the lexicon.
2. **BERT** (Devlin et al., 2018): introduced deeply contextualized bidirectional embeddings. A word's representation now depends on the words around it. This was transformative for query–document matching — "bank" in "river bank" versus "investment bank" finally had different vectors.
3. **Dense Passage Retrieval (DPR)** (Karpukhin et al., 2020): used a BERT-based bi-encoder architecture — one encoder for queries, one for passages — trained with contrastive learning. Dense retrieval combined with FAISS-based approximate nearest neighbor search significantly outperformed BM25 on top-k accuracy.

The modern search stack is hybrid: BM25 (or learned sparse retrieval like SPLADE) for first-stage candidate generation, followed by a neural cross-encoder for re-ranking. BM25 is not obsolete — it is repurposed as the fast, cheap first stage of a multi-stage pipeline.

---

## Part II: The History of Recommendation

### Tapestry and the Birth of Collaborative Filtering (1992)

While search emerged from library science, recommendation emerged from office work. In 1992, researchers at Xerox PARC built **Tapestry**, an email filtering system that let users annotate messages and write queries referencing others' annotations — "show me emails that Bob found interesting" [6].

They coined a term: **collaborative filtering**. The insight was that relevance is social. If people with similar tastes found something useful, you probably will too. Tapestry required users to write explicit queries, which limited its reach. But the idea was planted.

### GroupLens and Automated Collaborative Filtering (1994)

The automation came from the **GroupLens** project at the University of Minnesota. Paul Resnick, John Riedl, Joseph Konstan, and colleagues built a system that automatically predicted how much a user would like a Usenet article based on ratings from similar users [7].

The architecture was elegant:

```python
# User-based collaborative filtering (simplified)
def predict_rating(user, item, ratings_matrix):
    # Find users similar to the target user (Pearson correlation)
    neighbors = find_k_nearest_neighbors(user, k=50)
    # Weighted average of neighbors' ratings for this item
    numerator = sum(sim(user, n) * ratings_matrix[n][item] for n in neighbors)
    denominator = sum(abs(sim(user, n)) for n in neighbors)
    return numerator / denominator
```

GroupLens didn't just build a research prototype — they spawned **Net Perceptions**, a company that served Amazon, CDnow, and others. In 2010, the team received the **ACM Software System Award** for showing how "a distributed set of users could receive personalized recommendations by sharing ratings."

### Amazon Item-to-Item Collaborative Filtering (2003)

The GroupLens approach — user-based collaborative filtering — had a scaling problem. Computing user-user similarity is O(N²) in the number of users, and it must be recomputed as users rate more items.

Amazon's Greg Linden, Brent Smith, and Jeremy York published a different approach in 2003: **item-to-item collaborative filtering** [8]. Instead of finding similar users, find similar items:

```python
# Item-based collaborative filtering
def build_similarity_matrix(ratings_matrix):
    # Pre-compute item-item similarities offline
    for item_i in items:
        for item_j in items:
            # Only consider users who rated both
            common_users = users_who_rated(item_i) ∩ users_who_rated(item_j)
            sim(item_i, item_j) = cosine_similarity(
                [ratings[u][item_i] for u in common_users],
                [ratings[u][item_j] for u in common_users]
            )

def recommend(user):
    # Online: look up items similar to what the user liked
    liked_items = get_items_rated_highly(user)
    candidates = []
    for item in liked_items:
        candidates.extend(similar_items[item])
    return rank_and_filter(candidates)
```

The key advantage: the O(N²) item-to-item similarity matrix is computed offline. At serving time, recommendations are a simple lookup.

Linden later explained why this approach beat search-based methods in practice: "Rather than matching the user to similar customers, item-to-item collaborative filtering matches each of the user's purchased and rated items to similar items, then combines those similar items into a recommendation list." Search-based approaches — constructing queries from a user's purchase history to find items with similar keywords or categories — produced recommendations that were either too general (bestsellers) or too narrow (more books by the same author). Collaborative filtering discovered cross-category connections: people who bought *Into Thin Air* also bought *The Perfect Storm*. No keyword match connects a mountaineering disaster to a fishing boat tragedy. Only behavior does.

Amazon reported that recommendations drove a measurable fraction of revenue — a finding that validated the commercial importance of the search/recommendation distinction. Search helps users find what they know they want. Recommendation helps them discover what they didn't know existed. Both drive revenue, but through completely different mechanisms.

### The Netflix Prize (2006–2009)

If Amazon's paper showed that recommendation had commercial value, the Netflix Prize showed that it was a serious mathematical discipline.

In October 2006, Netflix published a dataset — 100 million ratings from 480,000 users on 17,770 movies — and offered $1 million to any team that could improve their Cinematch algorithm's RMSE by 10%. Three years of frantic innovation followed.

The winning approach — and the one that reshaped the field — was **matrix factorization**. Instead of computing user-user or item-item similarities directly, learn low-dimensional vector representations (embeddings) for every user and item:

```python
# Matrix factorization via SGD (simplified)
def train_matrix_factorization(ratings, k=50, epochs=100, lr=0.01, lambda_reg=0.02):
    # Initialize user and item latent factor matrices
    P = np.random.normal(0, 0.1, (num_users, k))  # user factors
    Q = np.random.normal(0, 0.1, (num_items, k))  # item factors

    for epoch in range(epochs):
        for u, i, r in ratings:
            error = r - np.dot(P[u], Q[i])
            P[u] += lr * (error * Q[i] - lambda_reg * P[u])
            Q[i] += lr * (error * P[u] - lambda_reg * Q[i])

    return P, Q

def predict(u, i):
    return np.dot(P[u], Q[i])
```

Yehuda Koren's **SVD++** fused explicit ratings with implicit feedback (what you browsed, not just what you rated) and bias terms (some users rate everything high; some movies are universally beloved). The winning ensemble — BellKor's Pragmatic Chaos — achieved RMSE 0.8556, a 10.07% improvement over Cinematch [9].

The Netflix Prize established matrix factorization as the dominant paradigm for nearly a decade. It also revealed the field's key tension: the metric that drove the competition (RMSE on withheld ratings) doesn't actually measure whether users are satisfied. More on that later.

### The Deep Learning Era (2016–2020)

YouTube's 2016 paper *"Deep Neural Networks for YouTube Recommendations"* marked the moment deep learning entered production recommendation at scale [10]. The architecture was two-stage:

1. **Candidate generation**: a deep neural network narrowed millions of videos to hundreds, using watch history, search history, demographics, and video embeddings as features.
2. **Ranking**: a separate DNN scored candidates using hundreds of features, optimizing for expected watch time rather than clicks — a critical insight that aligned the objective with what YouTube actually cared about.

The feature engineering that followed — Wide & Deep, DeepFM, Deep Interest Network (DIN) — pushed recommendation from collaborative filtering into a regime where hundreds of heterogeneous signals (user demographics, item metadata, time, device, context, cross-features) were ingested by deep networks optimized for specific business objectives.

By 2020, the standard industry pipeline had crystalized into three stages: **retrieval** (fast, cheap, high recall), **ranking** (expensive ML, high precision), and **re-ranking** (diversity, freshness, business rules).

---

## Part III: Netflix — Both Problems Side by Side

No company illustrates the search–recommendation distinction better than Netflix. They run both systems at global scale, on the same catalog, for the same users — and have published extensively about each.

### Netflix Search: Query-Driven Retrieval at Scale

Netflix Studio Search indexes federated GraphQL data across the entire content production pipeline. Their architecture, described in a 2022 Netflix Tech Blog post and subsequent InfoQ coverage, is built on **Elasticsearch** (which runs Apache Lucene's BM25 under the hood) with a custom SQL-like query DSL [11][12].

The pipeline works as follows:
- Application events and Change Data Capture (CDC) events stream into **Kafka** topics
- **Apache Flink** processes consume these events, enrich data via federated GraphQL queries, and sink documents into Elasticsearch
- A custom ANTLR-based query parser translates a SQL-like DSL into proper Elasticsearch queries, handling nested JSON documents transparently
- Search returns matching entity keys only; results are hydrated via federated GraphQL queries for any needed data
- Authorization is late-binding: at query time, a centralized policy server evaluates constraints, translated into additional Boolean filters AND-ed with the user's query

Separately, Netflix's **Asset Management Platform (AMP)** indexes over 7TB of digital media metadata. They learned hard lessons about Elasticsearch shard sizing: their original design (one index per asset type, ~900 indices, 16,200 shards) caused CPU hotspots because shard sizes ranged from thousands to millions of documents. Switching to time-bucket-based indices with uniform sizes dropped CPU utilization from 70% to 10% [13].

Netflix also uses **Elasticsearch Percolate Queries** for reverse search — instead of matching documents to queries, match queries to documents. When a production asset changes (e.g., "movie shooting in Mexico City without a key role assigned"), percolation identifies which saved searches match, enabling targeted notifications. This is the kind of problem that only exists when you think of search as a retrieval infrastructure problem, not just ranking [14].

### Netflix Recommendation: From Collaborative Filtering to Foundation Models

Netflix's recommendation story is better known, but worth retracing for the contrast it provides with search. Reed Hastings, Netflix's co-founder, articulated the goal with a simple analogy: "If the Starbucks secret is a smile when you get your latte, ours is that the website adapts to the individual's taste." The emotional target — "Netflix 'gets me'" — is fundamentally different from search's target of "Netflix found what I asked for." Recommendation is about identity. Search is about utility.

Hastings revealed an even deeper insight at the 2018 TED conference, explaining why Netflix shifted from relying on what users *say* they like (star ratings) to tracking what they *actually watch*:

> "What happens is, when we rate, we're meta-cognitive about quality — that's sort of our aspirational self. It works out much better, to please people, to look at the actual choices that they make."

This is a problem that search doesn't have. In search, the user explicitly states their intent — the query is the ground truth. In recommendation, the ground truth is hidden. Users say they want documentaries; they watch reality TV. A recommender that trusts stated preferences over revealed preferences fails. A search engine that second-guesses the query fails.

This is also why Hastings famously said "We compete with sleep" — not just with other streaming services. The recommender's job is not to satisfy a stated need but to command attention in a world of infinite alternatives. Search competes with ignorance. Recommendation competes with every other possible way to spend time.

The **three-tier serving architecture** (offline, nearline, online) described in Netflix's 2013 tech blog post remains the backbone [15]:
- **Offline**: batch model training, feature pre-computation
- **Nearline**: asynchronous event-triggered processing (e.g., update recommendations immediately after a viewing session)
- **Online**: real-time scoring with strict latency SLAs

But the models themselves evolved dramatically. Netflix's 2021 *AI Magazine* article *"Deep Learning for Recommender Systems: A Netflix Case Study"* contains a sobering finding: when only user–item interaction data is available — the classic collaborative filtering setting — properly tuned non-deep-learning baselines remained competitive with deep networks [16]. The power of deep learning emerged only when heterogeneous features (metadata, context, images, text) were incorporated.

This is a crucial point that connects back to the search–recommendation distinction: **search has always been multi-modal** (query text, document text, links, anchor text, click data, freshness). Recommendation was historically impoverished — just a ratings matrix. The deep learning era equalized this by giving recommendation systems the same multi-modal, heterogeneous-feature diet that search had enjoyed since the 2000s.

In 2025, Netflix published what may be their most ambitious recommendation paper yet: a **Foundation Model for Personalized Recommendation** [17]. The architecture treats user interaction histories as sequences — analogous to how GPT treats text — and uses autoregressive next-token prediction to model user behavior:

- User actions (plays, ratings, searches, browsing) are "tokenized" via an analog of Byte Pair Encoding
- A transformer model predicts the next item the user will engage with
- Auxiliary prediction heads predict genres, languages, and metadata — acting as regularizers
- Sparse attention mechanisms extend the context window to hundreds of interactions
- Cold-start new titles use a metadata+ID embedding combination with an attention-based mixing layer weighted by entity "age"

Crucially, Netflix confirmed that **scaling laws apply** to recommendation foundation models — larger datasets, more parameters, and longer context windows yield consistent improvements.

### The Convergence: UniCoRn (2024)

The most important paper for understanding the search–recommendation boundary at Netflix is **UniCoRn** (Unified Contextual Recommender), presented at RecSys 2024 [18]. Before UniCoRn, Netflix ran three separate models for:
1. Search (query → video)
2. Homepage recommendations (user profile → video)
3. "More Like This" (video → video)

UniCoRn replaces all three with a single transformer-based deep learning model. The insight is in *how* they unified them: by treating **task type as a context feature**. The input representation becomes:

```
[user_id, query (or null), country, source_entity_id (or null), task_type]
```

For search tasks where there is no source entity, null imputation is used. For video-video recommendations where there is no query, the title text of the source entity is imputed. The shared base model learns cross-task patterns, while task-specific output heads predict the right objective for each surface (click probability for search, watch probability for recommendations).

The results: **+7% lift for search, +10% lift for recommendations** compared to separate models.

But here is the critical detail: UniCoRn succeeded not by *ignoring* the distinction between search and recommendation, but by making the model **explicitly aware of it**. The `task_type` feature tells the model which behavior to invoke. The imputation strategy respects the structural difference (search has queries, recommendation has source entities). The separate output heads optimize for different objectives because, at Netflix, a successful search is not the same thing as a successful recommendation.

Netflix also documented the central tension: **personalization can overpower query relevance**. If a user searches for "documentaries about World War II" and the recommender knows the user loves romantic comedies, should it show *The Notebook*? Obviously not. UniCoRn took an incremental approach — semi-personalized first (user clustering as a feature), then fully personalized with fine-tuned user/item representations. The process required careful guardrails to ensure search results remain *relevant to the query* even as they benefit from personalization [19].

This is the tension that makes the distinction matter. When you collapse search into recommendation, you risk showing users what they usually like instead of what they explicitly asked for.

---

## Part IV: How the Industry Handles the Boundary

Netflix isn't the only company navigating this. Several engineering blogs document different approaches to the same problem.

### Spotify: Intent-Based Routing with LLMs

Spotify's 2025 paper *"You Say Search, I Say Recs"* describes an agentic approach to query understanding [20]. An LLM-based router classifies user intent:

- **Navigational queries** ("find song X") → traditional search (Elasticsearch + BM25 + Voyager ANN)
- **Exploratory queries** ("new indie rock releases," "Italian 80s disco nostalgia") → recommendation APIs leveraging collaborative filtering
- **Mixed intent** → both paths fire in parallel, results fused into a structured SERP

The router is a small, fine-tuned LLM (distilled from a larger teacher) achieving 60% latency reduction and 99% cost reduction. Critically, only the query goes to the router — not user features — maximizing cache hit rates. User features are injected downstream in the retrieval/ranking tools, which is where personalization belongs.

Results versus regular search: +115% for finding similar artists, +91% for new music discovery, +25% for broad music searches.

Spotify also explored a unified generative approach using **Semantic IDs** — discrete tokens derived from item embeddings that serve as autoregressive prediction targets [21]. Their finding: task-specific Semantic IDs optimized for search don't generalize to recommendation, and vice versa. A multi-task bi-encoder achieves a Pareto-optimal trade-off, but the boundary persists — you can't optimize one embedding space for both tasks simultaneously without compromise.

### DoorDash: Content-First Embeddings, Task-Aware Retrieval

DoorDash published a series of blog posts in 2024–2025 describing their modernization of search and recommendation across restaurants, grocery, and retail verticals [22][23]. Their environment makes the search–recommendation distinction unusually sharp. Consider two DoorDash users:

- **Alice** searches for "spicy Sichuan noodles near me." She has a specific craving. The system's job is retrieval: match her query to relevant restaurants, rank by distance and quality, and don't substitute her specific intent with what she usually orders.

- **Bob** opens the app with no particular craving. He's ordered Thai, pizza, and sushi in the past month. The system's job is recommendation: infer what he might want tonight, balance familiarity with discovery, and surface options he wouldn't have searched for.

If DoorDash treats Alice's search like Bob's browse — injecting personalization that overpowers her query — she gets Thai food when she wanted Sichuan noodles. If they treat Bob's browse like Alice's search — requiring an explicit query — he gets a blank screen when he wanted inspiration. Both users churn. For different reasons.

DoorDash's approach is instructive because clicks alone don't distinguish these cases. A click on a Sichuan noodle soup doesn't tell you whether the user wanted spicy food specifically or noodles specifically — or whether they were just hungry and clicked the first thing they saw.

Their solution is **content-first embeddings**: LLMs generate rich, standardized profiles for every merchant and item on the platform, covering ingredients, preparation, cuisine, dietary attributes, and context. These profiles are then embedded using off-the-shelf encoders (they selected `gemini-embedding-001` with 256-dimensional Matryoshka embedding). The key finding: upgrading profile quality yielded +31% improvement versus only +6% from a better encoder on raw metadata. Data quality dominates model choice.

For recommendation, DoorDash developed **Consumer Memory Blocks** — typed, composable, evidenced representations of user state (long-running preferences, behavioral patterns, household context, brand affinities) serialized as JSON prompts for batch LLM calls. The generated carousel titles and search intents are embedded offline and retrieved online via Milvus ANN. No LLM runs in the online serving path — cost is amortized across the refresh interval. Results: +2.4% order rate [24].

### Airbnb: Embeddings That Bridge Search and Discovery

Airbnb's KDD 2018 paper on real-time personalization for search ranking is a classic of the genre [25]. They trained listing embeddings on over 800 million search click sessions using a skip-gram model adapted from word2vec. The innovations were domain-specific: negative sampling within the same market (users mostly search within a single city, so negatives drawn globally are trivially easy to reject), and the booked listing is always treated as the context being predicted regardless of its position in the click sequence.

The embeddings are used as real-time features in the ranking model: similarity of a candidate listing to listings the user recently clicked (`EmbClickSim`), skipped (`EmbSkipSim`), or spent time viewing (`EmbLastLongClickSim`). Five embedding features ranked among the top 20 among all 104 model features. The result: +2.27% offline NDCG lift and statistically significant booking gains.

Airbnb's 2019 follow-up on applying deep learning to search documented instructive failures [26]. Listing ID embeddings overfit because listings have limited bookings. Multi-task learning (booking + long views) increased long views but bookings stayed flat — long views correlated with but were orthogonal to bookings (expensive or unusual listings get looked at but not booked). These are the kinds of lessons that only emerge when you treat search ranking as its own problem with its own objective landscape.

### Pinterest: Two-Tower Architectures and Multi-Task Embeddings

Pinterest's engineering blog and publications describe an evolving two-tower architecture deployed across homefeed recommendation and search [27]. The user tower encodes long-term history, context, and real-time sequences via transformers. The item tower encodes category, description, and image features. At serving time, the dot product between user and item embeddings produces a relevance score, and ANN search via their in-house Manas system (HNSW-based) retrieves candidates.

For **search** specifically, Pinterest's OmniSearchSage (WWW 2024) extended the two-tower approach into a multi-task, multi-entity framework where a single unified query embedding retrieves across pins, products, and related queries simultaneously [28]. A teacher-student distillation approach followed: an 8B-parameter Llama cross-encoder (teacher) achieved +12–20% improvement over BERT baselines, then a bi-encoder student was trained on 100× more data from daily search logs for production serving. Pin embeddings are pre-computed offline; query embeddings benefit from 85% cache hit rates.

The common pattern across all of these companies is not unification for its own sake — it is **awareness of task identity**. Every system that successfully spans search and recommendation does so by encoding *which task it's doing* as a first-class signal.

---

## Part V: How LLMs Change Both — And Why the Distinction Still Matters

The arrival of LLMs — GPT-4, Claude, Gemini, and their open-source counterparts — is the most significant development in both information retrieval and recommendation since BERT. But the way LLMs affect each field is different, and understanding the difference is essential for engineering.

### How LLMs Transform Search

For search, LLMs operate at four levels:

1. **Query understanding**: LLMs disambiguate, expand, and reformulate queries. "Show me movies like Inception but funnier" becomes a structured search intent with entity extraction and constraint parsing.

2. **Document understanding**: LLMs generate richer document representations — summaries, key phrases, entity tags, embeddings — that transcend the vocabulary mismatch problem by operating in semantic space.

3. **Retrieval-Augmented Generation (RAG)**: Instead of returning a ranked list, the system retrieves relevant documents and generates a synthesized answer grounded in them. This is the dominant LLM-search paradigm in 2025.

4. **Generative retrieval**: The most radical approach — the LLM directly generates document identifiers without an explicit retrieval index. The model's parameters *are* the index. This is still experimental but advancing rapidly [29].

The core search user need — "I have a question, find me the answer" — aligns naturally with LLM capabilities. The LLM doesn't replace retrieval; it augments it at every stage of the pipeline.

### How LLMs Transform Recommendation

For recommendation, the transformation is more structural:

1. **Feature engineering automation**: LLMs generate rich item descriptions and user profiles from sparse structured data. DoorDash's content embeddings and Consumer Memory Blocks are the paradigmatic example — the LLM produces the fuel, but the retrieval and ranking engines are still purpose-built.

2. **Generative recommendation**: Instead of scoring a pre-defined candidate set, the LLM autoregressively generates item tokens. Meta's HSTU, Kuaishou's OneRec, and Google's TIGER treat recommendation as a sequence-to-sequence problem. The model is the index, the retriever, and the ranker — all in one [30].

3. **Conversational recommendation**: Multi-turn dialogue where the LLM elicits preferences, proposes items, and adapts based on feedback. This collapses the boundary between search and recommendation: the user might start with a search-like query ("I want a sci-fi movie") and, through conversation, the system shifts into recommendation mode ("based on what you said about liking cerebral plots, you might enjoy...").

4. **Agentic recommendation**: LLM-powered agents that plan, use tools, maintain memory, and reason across multiple steps. Spotify's intent router, DoorDash's carousel generation pipeline, and the ARAG framework (SIGIR 2025) all embody this pattern [31].

### Why the Distinction Survives

Despite this convergence, the distinction remains essential for four reasons:

**1. Objective alignment.** A search system optimizing for engagement will show users addictive content instead of answering their query. A recommendation system optimizing for precision will produce an echo chamber. The objective functions *must* be different, and that difference propagates through the entire architecture — what data you collect, what loss you optimize, what you measure online.

**2. Evaluation.** Search quality is measured against explicit relevance judgments: NDCG, MRR, precision@K. Recommendation quality is measured against user behavior: retention, discovery, session length, long-term satisfaction. These metrics are correlated but not identical. A system that is jointly optimized without awareness of task will drift toward whichever signal is strongest — usually engagement.

**3. The serendipity gradient.** In search, serendipity is a bug. If a user types "The Godfather" and you show them "Goodfellas," you've failed — no matter how good Goodfellas is. In recommendation, serendipity is a feature. The entire point is surfacing things the user wouldn't have known to search for. Collapsing both into one model without task conditioning loses this distinction.

    Imagine wanting to rewatch *The Godfather*. You type it into the search bar. The system returns *Goodfellas* first because "people who watch The Godfather also love Goodfellas." That's a broken search experience. Now imagine browsing your Netflix homepage on a Friday night. The system shows you *The Godfather* — which you've already seen four times. That's a broken recommendation experience. Both failures happen when teams optimize for the wrong thing. The first team optimized for engagement (clicks on "also loved") instead of relevance (did we return the exact match?). The second team optimized for precision (only showing items similar to known likes) instead of discovery (showing items the user might love but hasn't seen).

**4. Infrastructure.** Search requires inverted indices, real-time query parsing, and sub-100ms latency at high QPS. Recommendation requires user profile stores, embedding indices, and offline batch pipelines amortized across daily refresh cycles. The infrastructure shapes the architecture, and the architecture shapes what the model can do.

A 2024 survey of generative search and recommendation in the LLM era puts it well: the fields share a common mathematical framework — matching entities across representation spaces — but differ in the *nature of the mismatch*. Search deals with query–document mismatch (different words for the same concept). Recommendation deals with user–item mismatch (users and items live in completely different semantic spaces). The first is a lexical/semantic gap. The second is a modality gap. Solving them requires different tools, even when those tools share a common transformer backbone [32].

---

## Part VI: The Distinction in Everyday Life

The search–recommendation distinction isn't just an engineering concern. It's visible in how we make decisions every day. Recognizing it clarifies what kind of system you're building — and what it should be good at.

### The Restaurant Menu (Search) vs. The Chef's Tasting Menu (Recommendation)

When you sit down at a restaurant and scan the menu, you're doing **search**. You have an intent — "I feel like pasta" — and you're scanning a structured catalog for matches. The menu is an inverted index: organized by category (appetizers, mains, desserts), each dish has attributes (ingredients, price, dietary tags). Your satisfaction depends on whether the menu accurately represents what the kitchen can deliver.

When the chef sends out a tasting menu — "trust me, you'll love this" — you're receiving a **recommendation**. The chef has built a model of what you might enjoy (based on dietary restrictions you mentioned, the season, what's freshest today) and is making predictions. Your satisfaction depends on whether the chef's model of you is accurate — and whether they can surprise you in a good way.

A mistake in the first scenario: "I ordered the carbonara and got the bolognese." A mistake in the second: "The chef brought me a dish centered on mushrooms, which I hate." The first is a retrieval error — the kitchen had the wrong item mapped to your query. The second is a modeling error — the chef's understanding of you was wrong. Different failures, different fixes.

### The Grocery List (Search) vs. The Recipe Suggestion (Recommendation)

You walk into a grocery store with a list: milk, eggs, bread, tomatoes. That's **search** — you know what you need, you need to find it efficiently. The store's job is to make items findable: clear signage (an index), logical layout (a taxonomy), and fast checkout (low latency). Any deviation from the list is a failure — "I wanted tomatoes and came back with bell peppers" is a bad outcome, even if you like bell peppers.

Now imagine the store's app notices what's in your cart and says: "With those ingredients, you could make a great shakshuka. You'd need cumin, paprika, and feta — want me to show you where they are?" That's **recommendation**. It takes what it knows about you (your cart, your purchase history) and suggests something you didn't know you wanted. The deviation from the list is the *point*.

### The Difference in Error Cost

These analogies reveal something crucial: search errors and recommendation errors have different *shapes*.

A **search error** is a precision failure: "I asked for X and got Y." The user immediately knows something is wrong, because they had a specific expectation. The error is visible and attributable.

A **recommendation error** is a relevance failure: "The system showed me something I don't care about." The user doesn't know if the system is broken or if they're just having a bad experience. The error is invisible and erodes trust cumulatively.

This is why Peter Norvig's 80% bar applies to search but not recommendation. If a search engine gets 80% of results right, users happily ignore the other 20% — they can see what went wrong and scroll past it. If a recommender gets 80% right, users notice the 20% wrong more than the 80% right, because every wrong suggestion is an interruption — an item that took up screen space that could have shown something they'd love.

### Why "Two Sides of the Same Coin" Is Misleading

The coin metaphor implies you can flip from one to the other by changing perspective. You can't. The data is different (queries vs. behavioral histories). The latency profiles are different (sub-100ms real-time vs. batch-refreshed). The error tolerance is different (visible and dismissable vs. invisible and trust-eroding). The user posture is different (active and goal-directed vs. passive and open).

The better metaphor is a **restaurant**: search is the menu, recommendation is the tasting menu, and the LLM-powered hybrid is a waiter who listens to what you're in the mood for, knows what the kitchen does well tonight, and helps you decide — sometimes by pointing at the menu (search), sometimes by making a suggestion you wouldn't have thought of (recommendation), and always by knowing which mode you're in right now.

---

## Part VII: Practical Guidance

If you're building a system that involves both search and recommendation, here are the architectural principles that emerge from the literature:

**1. Make task identity a first-class feature.** Whether you use Netflix's approach (separate output heads on a shared backbone), Spotify's approach (LLM router dispatching to separate systems), or DoorDash's approach (separate retrieval pipelines with shared embeddings), the model must know whether it's doing search or recommendation.

**2. Don't let personalization overpower query relevance.** The Netflix experience is instructive. They added personalization to search incrementally, with explicit guardrails. The fully personalized UniCoRn improved both search and recommendations, but it required careful tuning to ensure search results remain relevant to the query.

**3. Treat the user's posture as a design constraint.** Search users lean forward; recommendation users lean back. This affects not just ranking but UI, latency budgets, and the acceptable cost of being wrong. A search error ("I searched for X and got Y") is more salient to users than a recommendation error ("this row isn't for me").

**4. Use shared infrastructure where it makes sense — separate where it doesn't.** Embedding stores, feature stores, and model training pipelines can be shared. Retrieval indices, ranking objectives, and evaluation frameworks should be task-specific.

**5. LLMs are infrastructure, not a replacement for retrieval.** The most successful deployments — DoorDash, Spotify, Pinterest — use LLMs for content understanding, query intent classification, and feature generation *offline*, while keeping online retrieval and ranking in purpose-built, low-latency systems. The LLM enriches; it doesn't replace.

---

## Conclusion

In 1992, Belkin and Croft asked whether information retrieval and information filtering were two sides of the same coin. Thirty-four years later, the question has a sharper answer: they share a mint. The mathematical machinery — vector spaces, embedding learning, transformer attention, contrastive objectives — is increasingly shared. Netflix's UniCoRn, Spotify's intent router, DoorDash's content embeddings, Pinterest's multi-task two-tower models — all exploit this commonality to reduce technical debt and improve both tasks simultaneously.

But sharing machinery is not sharing purpose. As Greg Linden showed at Amazon, the algorithms that retrieve known items and the algorithms that surface unknown ones succeed for different reasons. As Reed Hastings understood, the emotional contract of "Netflix gets me" is not the same as "Netflix found what I searched for." As Karen Spärck Jones knew in 1999, systems that assist human users cannot replace them — they can only narrow the gap between what the user says and what the user means. And as Peter Norvig warned, the gap between "here are some suggestions" and "here is what you need" demands a higher standard of trust.

The arrival of LLMs makes these distinctions *more* important, not less. When a single model can retrieve, rank, recommend, and explain — all in natural language — the question is no longer "can we unify?" The question is "do we know which one we're doing right now?"

The answer must be yes. Because search competes with ignorance. Recommendation competes with sleep. They are not the same fight.

---

**References**

1. Nicholas J. Belkin and W. Bruce Croft. [*Information Filtering and Information Retrieval: Two Sides of the Same Coin?*](https://dl.acm.org/doi/10.1145/138859.138861). Communications of the ACM, 35(12): 29–38, 1992.

2. Gerard Salton, Anita Wong, and Chung-Shu Yang. [*A Vector Space Model for Automatic Indexing*](https://dl.acm.org/doi/10.1145/361219.361220). Communications of the ACM, 18(11): 613–620, 1975.

3. Stephen E. Robertson, Steve Walker, Susan Jones, Micheline Hancock-Beaulieu, and Mike Gatford. [*Okapi at TREC-3*](https://trec.nist.gov/pubs/trec3/papers/city.ps.gz). Proceedings of the Third Text REtrieval Conference (TREC-3), 1994.

4. Sergey Brin and Lawrence Page. [*The Anatomy of a Large-Scale Hypertextual Web Search Engine*](https://doi.org/10.1016/S0169-7552(98)00110-X). Computer Networks and ISDN Systems, 30(1–7): 107–117, 1998.

5. Christopher J.C. Burges. [*From RankNet to LambdaRank to LambdaMART: An Overview*](https://www.microsoft.com/en-us/research/publication/from-ranknet-to-lambdarank-to-lambdamart-an-overview/). Microsoft Research Technical Report MSR-TR-2010-82, 2010.

6. David Goldberg, David Nichols, Brian M. Oki, and Douglas Terry. [*Using Collaborative Filtering to Weave an Information Tapestry*](https://dl.acm.org/doi/10.1145/138859.138867). Communications of the ACM, 35(12): 61–70, 1992.

7. Paul Resnick, Neophytos Iacovou, Mitesh Suchak, Peter Bergstrom, and John Riedl. [*GroupLens: An Open Architecture for Collaborative Filtering of Netnews*](https://dl.acm.org/doi/10.1145/192844.192905). Proceedings of CSCW 1994, pp. 175–186.

8. Greg Linden, Brent Smith, and Jeremy York. [*Amazon.com Recommendations: Item-to-Item Collaborative Filtering*](https://doi.org/10.1109/MIC.2003.1167344). IEEE Internet Computing, 7(1): 76–80, 2003.

9. Yehuda Koren, Robert Bell, and Chris Volinsky. [*Matrix Factorization Techniques for Recommender Systems*](https://doi.org/10.1109/MC.2009.263). IEEE Computer, 42(8): 30–37, 2009.

10. Paul Covington, Jay Adams, and Emre Sargin. [*Deep Neural Networks for YouTube Recommendations*](https://dl.acm.org/doi/10.1145/2959100.2959190). Proceedings of RecSys 2016, pp. 191–198.

11. Netflix Technology Blog. [*Netflix Studio Search: Using Elasticsearch and Apache Flink to Index Federated GraphQL Data*](https://netflixtechblog.com/studio-search-using-elasticsearch-and-apache-flink-to-index-federated-graphql-data-7d77cad7b7c8). March 2022.

12. InfoQ. [*Netflix Studio Search: Using Elasticsearch and Apache Flink to Index Federated GraphQL Data — News Coverage*](https://www.infoq.com/news/2022/04/netflix-studio-search/). April 2022.

13. Netflix Technology Blog. [*Elasticsearch Indexing Strategy in Asset Management Platform (AMP)*](https://netflixtechblog.com/elasticsearch-indexing-strategy-in-asset-management-platform-amp-99332231e541). 2023.

14. InfoQ. [*Netflix Uses Elasticsearch Percolate Queries to Implement Reverse Searches Efficiently*](https://www.infoq.com/news/2024/04/netflix-percolate-queries/). April 2024.

15. Netflix Technology Blog. [*System Architectures for Personalization and Recommendation*](https://netflixtechblog.com/system-architectures-for-personalization-and-recommendation-e081aa94b5d8). March 2013.

16. Harald Steck, Linas Baltrunas, Ehtsham Elahi, Dawen Liang, Yves Raimond, and Justin Basilico. [*Deep Learning for Recommender Systems: A Netflix Case Study*](https://doi.org/10.1609/aimag.v42i3.18140). AI Magazine, 42(3): 7–18, 2021.

17. Netflix Technology Blog. [*Foundation Model for Personalized Recommendation*](https://netflixtechblog.com/foundation-model-for-personalized-recommendation-1a0bd8e02d39). March 2025.

18. Moumita Bhattacharya et al. [*Joint Modeling of Search and Recommendations Via an Unified Contextual Recommender (UniCoRn)*](https://dl.acm.org/doi/fullHtml/10.1145/3640457.3688034). Proceedings of RecSys 2024.

19. Sudarshan Lamkhede and Christoph Kofler. [*Recommendations and Results Organization in Netflix Search*](https://dl.acm.org/doi/abs/10.1145/3460231.3474602). Proceedings of RecSys 2021, pp. 577–579.

20. Spotify Research. [*You Say Search, I Say Recs: A Scalable Agentic Approach to Query Understanding and Exploratory Search at Spotify*](https://research.atspotify.com/2025/9/you-say-search-i-say-recs-a-scalable-agentic-approach-to-query-understanding). September 2025.

21. Spotify Research. [*Semantic IDs for Generative Search and Recommendation*](https://www.research.atspotify.com/2025/9/semantic-ids-for-generative-search-and-recommendation). September 2025.

22. DoorDash Engineering. [*Using LLMs to Build Content Embeddings for Search and Recommendations*](https://careersatdoordash.com/blog/doordash-llms-to-build-content-embeddings-for-search-and-recommendations/). 2025.

23. DoorDash Engineering. [*When GenAI Meets Personalization: Powering DoorDash's Next-Generation Homepage Experience*](https://careersatdoordash.com/blog/doordashs-next-generation-homepage-genai/). 2025.

24. DoorDash Engineering. [*Offline LLMs, Online Personalization: Generating Carousels at DoorDash*](https://careersatdoordash.com/blog/doordash-offline-llms-online-personalization-generating-carousels/). 2025.

25. Mihajlo Grbovic and Haibin Cheng. [*Real-time Personalization using Embeddings for Search Ranking at Airbnb*](https://dl.acm.org/doi/10.1145/3219819.3219885). Proceedings of KDD 2018, pp. 311–320.

26. Malay Haldar, Mustafa Abdool, et al. [*Applying Deep Learning to Airbnb Search*](https://arxiv.org/abs/1810.09591). Proceedings of KDD 2019.

27. Pinterest Engineering. [*Advancements in Embedding-Based Retrieval at Pinterest Homefeed*](https://medium.com/pinterest-engineering/advancements-in-embedding-based-retrieval-at-pinterest-homefeed-176a3c03df64). 2024.

28. Pinterest Engineering. [*LLM-Enhanced Search Relevance at Pinterest*](https://medium.com/pinterest-engineering/llm-enhanced-search-relevance-at-pinterest-2025). 2025.

29. Yongqi Li et al. [*A Survey of Generative Search and Recommendation in the Era of Large Language Models*](https://arxiv.org/abs/2404.16924). arXiv:2404.16924, 2024.

30. Shashank Gupta et al. [*Generative Recommendation: A Survey of Models, Systems, and Industrial Advances*](https://www.techrxiv.org/doi/full/10.36227/techrxiv.176523089.94266134/v2). TechRxiv, 2025.

31. Chi Zhang et al. [*ARAG: Agentic Retrieval Augmented Generation for Personalized Recommendation*](https://arxiv.org/abs/2506.21931). Proceedings of SIGIR 2025.

32. Zhuang Liu et al. [*A Comprehensive Survey on LLM-Powered Recommender Systems: From Discriminative, Generative to Multi-Modal Paradigms*](https://ieeexplore.ieee.org/abstract/document/11129085). IEEE Access, 2024.

---

*See also: [On Rule Engines — From RETE to MCP](/on-rule-engines-rete-to-mcp) — for another exploration of how symbolic and neural approaches converge without collapsing into identity.*
