---
title: "S&R: Practical Architecture — Building Systems That Get It Right"
date: 2026-07-16
slug: search-recommendation-practical-architecture
summary: Part 7 of 7. Practical guidance for building systems that handle both search and recommendation: real-life analogies, error cost analysis, five architectural principles, and a reference hybrid architecture sketch.
tags: [search, recommendation, architecture, system-design, llm, practical, series]
series: search-recommendation
---

*S&R stands for Search & Recommendation. Part 7 of a seven-part series. Practical architecture guidance for building systems that handle both search and recommendation — correctly.*

Over six parts, we've traced the fifty-year history of search, the thirty-year history of recommendation, Netflix's dual-stack architecture, how four major companies handle the boundary, and how LLMs transform both fields. This final part translates that history into practical guidance.

## The Distinction in Everyday Life

Before the architecture, a mental model. The search–recommendation distinction isn't an engineering abstraction — it's visible in how we make decisions every day.

**The Restaurant Menu (Search) vs. The Chef's Tasting Menu (Recommendation).** When you scan a menu, you're doing search: you have an intent ("I feel like pasta") and you're scanning a structured catalog. Your satisfaction depends on whether the menu accurately represents what the kitchen delivers. A mistake — "I ordered carbonara, got bolognese" — is a retrieval error. When the chef sends out a tasting menu, you're receiving recommendation: the chef built a model of you and is making predictions. A mistake — "the chef brought me mushrooms, which I hate" — is a modeling error. Different failures, different fixes.

**The Grocery List (Search) vs. The Recipe Suggestion (Recommendation).** You walk into a store with a list: milk, eggs, bread. Search. The store's app notices your cart and says: "With those ingredients, you could make shakshuka. Want me to show you cumin and paprika?" Recommendation. The deviation from the list is the *point* in one case. It's a *failure* in the other.

**The Error Cost Difference.** A search error is a precision failure: "I asked for X and got Y." The user immediately knows something is wrong. The error is visible and attributable. A recommendation error is a relevance failure: "The system showed me something I don't care about." The user doesn't know if the system is broken or if they're just having a bad experience. The error is invisible and erodes trust cumulatively.

This is why Peter Norvig's 80% rule applies to search but not recommendation. If a search engine gets 80% right, users happily ignore the other 20% — they can see what went wrong. If a recommender gets 80% right, users notice the 20% wrong more than the 80% right, because every wrong suggestion is an interruption — screen space that could have shown something they'd love.

## Five Principles for Building Both

**1. Make task identity a first-class feature.**

```python
class TaskAwareModel(torch.nn.Module):
    """
    Whether you use Netflix's approach (separate output heads on a shared backbone),
    Spotify's approach (LLM router dispatching to separate systems), or DoorDash's
    approach (separate retrieval with shared embeddings) — the model must know
    whether it's doing search or recommendation.
    """

    def __init__(self, shared_backbone, num_tasks: int):
        super().__init__()
        self.backbone = shared_backbone
        self.task_embedding = torch.nn.Embedding(num_tasks, embedding_dim)
        self.heads = torch.nn.ModuleDict({
            "search": SearchHead(),
            "recommendation": RecommendationHead(),
            "similar_items": SimilarItemsHead(),
        })

    def forward(self, inputs, task_type: str):
        task_encoding = self.task_embedding(TASK_IDS[task_type])
        hidden = self.backbone(inputs, task_encoding)
        return self.heads[task_type](hidden)

# The task_type signal propagates through the entire model.
# It tells the model: "optimize for relevance to query" vs.
# "optimize for long-term engagement." These are different instructions.
```

**2. Don't let personalization overpower query relevance.**

Netflix learned this the hard way with UniCoRn. They added personalization to search incrementally, with explicit guardrails. The fully personalized model improved both tasks — but only after careful tuning to ensure search results remain *relevant to the query* even as they benefit from personalization signals.

```python
def safe_personalized_search(query: str, user_profile: dict,
                             base_ranker, personalization_model,
                             relevance_threshold: float = 0.7) -> list[Item]:
    """Personalize search results without overriding query relevance."""
    # Stage 1: Get relevance-scored candidates
    candidates = base_ranker.retrieve(query, k=200)

    # Stage 2: Apply personalization re-ranking
    personalized = personalization_model.rerank(candidates, user_profile)

    # Stage 3: Enforce relevance floor
    # A result that is NOT relevant to the query should never outrank
    # a result that IS relevant, regardless of personalization score.
    for i, item in enumerate(personalized):
        if item.relevance_score < relevance_threshold:
            personalized[i].final_score *= 0.1  # severe penalty

    personalized.sort(key=lambda x: x.final_score, reverse=True)
    return personalized[:10]
```

**3. Treat the user's posture as a design constraint.**

Search users lean forward; recommendation users lean back. This affects latency budgets, UI, and the acceptable cost of being wrong.

| Constraint | Search | Recommendation |
|---|---|---|
| Latency SLA | <100ms per keystroke | <200ms per page load |
| Error visibility | User sees it immediately | User may never notice — or slowly churn |
| Query volume | Every keystroke | Every page view |
| Freshness requirement | Near-real-time (new docs) | Daily batch (new items) |
| UI expectation | Explicit, controllable | Ambient, delightful |

**4. Share infrastructure where it makes sense — separate where it doesn't.**

```python
# SHARED: embedding stores, feature stores, model training pipelines
shared_infrastructure = {
    "feature_store": "Feast",           # same features, different consumers
    "embedding_index": "FAISS/Milvus",  # same ANN, different query patterns
    "model_registry": "MLflow",         # same versioning, different models
    "experiment_platform": "A/B tests", # same framework, different metrics
}

# SEPARATE: retrieval indices, ranking objectives, evaluation
separate_infrastructure = {
    "search_index": "Elasticsearch/BM25",     # inverted index, text-optimized
    "recs_index": "ANN over user-item space", # vector index, behavior-optimized
    "search_objective": "NDCG/MRR",           # relevance to query
    "recs_objective": "retention/discovery",  # long-term satisfaction
    "search_eval": "explicit relevance judgments",  # human-labeled
    "recs_eval": "A/B test on member retention",    # behavioral
}
```

**5. LLMs are infrastructure, not a replacement for retrieval.**

The most successful production deployments use LLMs for content understanding, query intent classification, and feature generation *offline*, while keeping online retrieval in purpose-built low-latency systems:

```python
class HybridLLMRetrievalPipeline:
    """LLMs enrich — they don't replace."""

    def offline_enrichment(self, catalog: list[Item]):
        """LLM generates rich profiles. Runs nightly. Costs amortized."""
        for item in catalog:
            item.llm_profile = self.llm.describe(item)       # rich description
            item.llm_embedding = self.encoder.encode(item.llm_profile)
            item.llm_tags = self.llm.extract_tags(item)      # structured metadata

    def online_search(self, query: str, k: int = 10) -> list[Item]:
        """No LLM in the request path. Sub-100ms."""
        # LLM-generated embeddings and tags are already indexed.
        # This is just BM25 + ANN retrieval + cross-encoder re-rank.
        return self.retrieval_pipeline.search(query, k)

    def online_recommendations(self, user_id: str, k: int = 10) -> list[Item]:
        """No LLM in the request path. Sub-200ms."""
        # LLM-generated carousel intents are pre-computed.
        # This is just embedding lookup + ANN + re-rank.
        return self.recs_pipeline.recommend(user_id, k)
```

## A Reference Architecture

Bringing everything together, here's what a system that handles both search and recommendation looks like in 2026:

```python
class SearchAndRecommendationSystem:
    """
    Reference architecture for a system that does both.

    Shared: embeddings, feature store, LLM enrichment pipeline.
    Separate: retrieval indices, ranking heads, objectives, evaluation.
    Task-aware: every model knows which mode it's in.
    """

    def __init__(self):
        # Shared offline enrichment (LLM — runs nightly)
        self.content_enricher = LLMContentEnricher()
        self.embedding_encoder = GeminiEncoder(dim=256)

        # Separate retrieval indices
        self.search_index = ElasticsearchBM25()     # text-optimized
        self.recs_index = MilvusANN()               # user-item optimized

        # Shared backbone, task-specific heads
        self.model = TaskAwareTwoTower(
            shared_backbone=TransformerEncoder(layers=6),
            tasks=["search", "recommendations", "similar_items"]
        )

        # Separate evaluation
        self.search_eval = SearchEvaluator(metric="ndcg@10")
        self.recs_eval = RecsEvaluator(metric="retention_30d")

    def nightly_batch(self):
        """Run once per day: LLM content enrichment, embedding refresh."""
        for item in self.catalog:
            item.profile = self.content_enricher.describe(item)
            item.embedding = self.embedding_encoder.encode(item.profile)

        self.search_index.rebuild()
        self.recs_index.rebuild()
        self.content_enricher.generate_carousels()  # DoorDash-style memory blocks

    def serve_search(self, query: str, user_id: str) -> SearchResult:
        """Online search: sub-100ms, no LLM in path."""
        candidates = self.search_index.retrieve(query, k=200)
        scored = self.model.score(candidates, task="search",
                                  query=query, user_id=user_id)
        return self.search_reranker.apply(scored, user_id)

    def serve_homepage(self, user_id: str) -> list[Carousel]:
        """Online recommendations: sub-200ms, no LLM in path."""
        carousels = self.carousel_store.lookup(user_id)
        filled_carousels = []
        for carousel in carousels:
            items = self.recs_index.retrieve(carousel.embedding, k=20)
            scored = self.model.score(items, task="recommendations",
                                      user_id=user_id)
            filled_carousels.append(Carousel(carousel.title, scored[:10]))
        return filled_carousels
```

## Conclusion: The Coin and the Mint

In 1992, Belkin and Croft asked whether information retrieval and information filtering were two sides of the same coin. Thirty-four years later, the question has a sharper answer: **they share a mint.**

The mathematical machinery — vector spaces, embedding learning, transformer attention, contrastive objectives — is increasingly shared. Netflix's UniCoRn, Spotify's intent router, DoorDash's content embeddings, Pinterest's multi-task two-tower models — all exploit this commonality. As Greg Linden showed at Amazon, the algorithms that retrieve known items and the algorithms that surface unknown ones succeed for different reasons. As Reed Hastings understood, the emotional contract of "Netflix gets me" is not the same as "Netflix found what I searched for." As Karen Spärck Jones knew in 1999, systems that assist human users cannot replace them — they can only narrow the gap. As Peter Norvig warned, the gap between "here are some suggestions" and "here is what you need" demands a higher standard of trust.

The arrival of LLMs makes these distinctions *more* important, not less. When a single model can retrieve, rank, recommend, and explain — all in natural language — the question is no longer "can we unify?"

The question is: **"do we know which one we're doing right now?"**

The answer must be yes. Because search competes with ignorance. Recommendation competes with sleep. They are not the same fight.

---

**References**

1. Nicholas J. Belkin and W. Bruce Croft. [*Information Filtering and Information Retrieval: Two Sides of the Same Coin?*](https://dl.acm.org/doi/10.1145/138859.138861). Communications of the ACM, 35(12): 29–38, 1992.

2. Francesco Ricci, Lior Rokach, and Bracha Shapira (editors). [*Recommender Systems Handbook*, 3rd Edition](https://link.springer.com/book/10.1007/978-1-0716-2197-4). Springer, 2022.

3. Yutao Zhu et al. [*Large Language Models for Information Retrieval: A Survey*](https://arxiv.org/abs/2308.07107). ACM TOIS, 2024.

4. Yongqi Li et al. [*A Survey of Generative Search and Recommendation in the Era of Large Language Models*](https://arxiv.org/abs/2404.16924). arXiv:2404.16924, 2024.

---

*Previous: [Part 6 — How LLMs Reshape Both Fields](/search-recommendation-llms)* · *Series start: [Part 1 — The Fundamental Distinction](/search-recommendation-fundamental-distinction)*
