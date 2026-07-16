---
title: "S&R: What Netflix Learned Running Search and Recommendation Side by Side"
date: 2026-07-16
slug: search-recommendation-netflix
summary: Deep dive into Netflix's search and recommendation systems — Elasticsearch/Flink indexing pipelines, three-tier serving architecture, the Foundation Model for personalized recommendation, and UniCoRn's unified approach to search and recs.
tags: [search, recommendation, netflix, elasticsearch, collaborative-filtering, foundation-model, unicorn, series]
series: search-recommendation
---

*S&R stands for Search & Recommendation. A deep dive into how Netflix builds and maintains both search and recommendation systems at global scale — and what happened when they tried to unify them.*

No company illustrates the search–recommendation distinction better than Netflix. They run both systems at global scale, on the same catalog, for the same users — and they've published extensively about each. This part examines both stacks, then looks at UniCoRn, their 2024 attempt to unify them.

Reed Hastings, Netflix's co-founder, articulated the recommendation goal simply: "If the Starbucks secret is a smile when you get your latte, ours is that the website adapts to the individual's taste." The emotional target — "Netflix gets me" — is fundamentally different from search's target of "Netflix found what I asked for." Recommendation is about identity. Search is about utility.

## Netflix Search: Elasticsearch, Flink, and Reverse Queries

Netflix Studio Search indexes federated GraphQL data across the entire content production pipeline. The architecture is built on **Elasticsearch** (running Apache Lucene's BM25 under the hood) with a custom query DSL [1].

```python
# Conceptual: Netflix's Studio Search indexing pipeline

# 1. Events stream into Kafka from applications and CDC
# 2. Apache Flink processes consume events, enrich via GraphQL, sink to Elasticsearch
# 3. A custom ANTLR-based query parser translates SQL-like DSL → Elasticsearch queries
# 4. Search returns entity keys only; results are hydrated via federated GraphQL
# 5. Authorization is late-binding, translated into Boolean filters AND-ed with the query

# Simplified: what a search request looks like
class NetflixSearchRequest:
    query: str            # e.g., "movies shooting in Mexico with Arnold Schwarzenegger"
    filters: dict         # e.g., {"content_type": "movie", "status": "in_production"}
    user_id: str          # for late-binding authorization
    facets: list[str]     # e.g., ["genre", "release_year", "language"]

class NetflixSearchResponse:
    entity_keys: list[str]       # just IDs — lean index
    facets: dict[str, list]      # aggregation results
    total_count: int

# Results are hydrated by the client via GraphQL:
# query { nodes(ids: $entity_keys) { title, synopsis, cast { name }, ... } }
```

Their **Asset Management Platform (AMP)** indexes over 7TB of digital media metadata. The hard lesson: their original design (one index per asset type, ~900 indices, 16,200 shards) caused CPU hotspots because shard sizes ranged from thousands to millions of documents. The fix — time-bucket-based indices with uniform sizes — dropped CPU from 70% to 10% [2].

```python
# AMP's metadata model: handling 1000+ asset types with heterogeneous schemas
# Solution: nested metadata field with typed value columns

es_document = {
    "asset_id": "asset_12345",
    "created_at": "2026-01-15T10:30:00Z",
    "metadata": [
        {"key": "resolution", "string_value": "4K"},
        {"key": "runtime_minutes", "long_value": 142},
        {"key": "color_space", "string_value": "Rec.2020"},
        {"key": "file_size_gb", "double_value": 287.5},
        {"key": "has_subtitles", "boolean_value": True},
        {"key": "shoot_date", "date_value": "2025-11-03"},
    ]
}

# Query: "find all 4K assets shot after October 2025 larger than 200GB"
# This becomes nested Elasticsearch queries matching both key AND value fields
```

Netflix also uses **Percolate Queries** for reverse search — matching documents to queries instead of queries to documents [3]. When a production asset changes (e.g., "movie shooting in Mexico City without a key role assigned"), percolation identifies which saved searches match, enabling targeted notifications.

```python
# Conceptual: percolation for production monitoring
class ReverseSearch:
    """Store queries as documents, match new data against them."""

    def index_saved_search(self, search_id: str, criteria: dict):
        """A creative executive saves: 'alert me when a thriller shoots in Thailand'."""
        es_query = translate_to_elasticsearch(criteria)
        es.index(index="saved_searches", id=search_id,
                 body={"query": es_query, "owner": "exec_42"})

    def percolate(self, asset_data: dict) -> list[str]:
        """When a new asset arrives, find all searches that match it."""
        result = es.percolate(index="saved_searches",
                             body={"doc": asset_data})
        return [match["_id"] for match in result["matches"]]

# This is the kind of problem that only exists when you think of search
# as a retrieval infrastructure problem, not just ranking.
```

## Netflix Recommendation: Three Tiers, One Goal

Netflix's 2013 tech blog post describing the three-tier serving architecture remains the conceptual backbone [4]:

```python
class NetflixRecommendationPipeline:
    """
    Three-tier architecture: offline, nearline, online.

    Offline: batch model training, feature pre-computation (hours/days)
    Nearline: event-triggered async processing (seconds/minutes)
    Online: real-time scoring with strict latency SLAs (milliseconds)
    """

    def offline_training(self):
        """Run nightly: train models, pre-compute features."""
        # Matrix factorization / foundation model training
        # Compute item-item similarity matrices
        # Pre-compute user embeddings for active users
        # Generate candidate sets per user cluster
        pass

    def nearline_update(self, user_id: str, event: str):
        """Triggers on user action: update recs after a viewing session."""
        if event == "finished_watching":
            # Update user embedding based on the completed title
            # Refresh the "Because You Watched" row
            # Adjust candidate generation weights
            pass

    def online_serve(self, user_id: str, context: dict, k: int = 40) -> list[str]:
        """Called on every page load. Must return in <200ms."""
        # Lookup pre-computed candidates
        # Score with real-time context (time of day, device, session)
        # Re-rank for diversity and freshness
        # Return the rows for the homepage
        pass
```

But the models themselves evolved dramatically. Their 2021 *AI Magazine* article contains a sobering finding: **when only user–item interaction data is available, properly tuned non-deep-learning baselines remain competitive** [5]. The power of deep learning emerged only when heterogeneous features (metadata, context, images, text) were incorporated.

This is a crucial point that connects back to the search–recommendation distinction: **search has always been multi-modal** — query text, document text, links, anchor text, click data, freshness. Recommendation was historically impoverished: just a ratings matrix. The deep learning era equalized this by giving recommendation systems the same heterogeneous-feature diet that search had enjoyed since the 2000s.

### The Foundation Model (2025)

In 2025, Netflix published their most ambitious recommendation paper: a **Foundation Model for Personalized Recommendation** that treats user interaction histories as sequences and uses autoregressive next-token prediction [6]:

```python
class NetflixFoundationModel(torch.nn.Module):
    """
    Autoregressive model: user actions → next item prediction.

    Inspired by GPT but for user behavior instead of text.
    """

    def __init__(self, vocab_size: int, embed_dim: int = 1024,
                 num_layers: int = 24, num_heads: int = 16):
        super().__init__()
        self.token_embedding = torch.nn.Embedding(vocab_size, embed_dim)
        self.positional_encoding = PositionalEncoding(embed_dim)

        # Sparse attention for long sequences (hundreds of interactions)
        self.transformer_blocks = torch.nn.ModuleList([
            SparseTransformerBlock(embed_dim, num_heads)
            for _ in range(num_layers)
        ])

        # Multi-token prediction: predict next n items, not just one
        self.output_heads = torch.nn.ModuleList([
            torch.nn.Linear(embed_dim, vocab_size) for _ in range(5)
        ])

        # Auxiliary prediction heads as regularizers
        self.genre_head = torch.nn.Linear(embed_dim, num_genres)
        self.language_head = torch.nn.Linear(embed_dim, num_languages)

    def forward(self, token_ids: torch.Tensor, mask: torch.Tensor
                ) -> tuple[torch.Tensor, torch.Tensor, torch.Tensor]:
        """Predict next items, genre, and language from interaction history."""
        x = self.token_embedding(token_ids)
        x = self.positional_encoding(x)

        for block in self.transformer_blocks:
            x = block(x, mask)

        # Next-item predictions (positions 1-5 ahead)
        item_preds = [head(x) for head in self.output_heads]
        # Auxiliary predictions
        genre_pred = self.genre_head(x.mean(dim=1))
        lang_pred = self.language_head(x.mean(dim=1))

        return item_preds, genre_pred, lang_pred

# Key design decisions:
# - User actions are "tokenized" via an analog of Byte Pair Encoding
# - Interaction tokens include watch duration, device, locale, time, item metadata
# - Cold-start titles: metadata+ID embedding with attention-based mixing weighted
#   by entity "age" — new titles lean on metadata, established ones on ID embeddings
# - Scaling laws confirmed: more data, more parameters, longer context → better
```

Hastings revealed why behavioral data trumps stated preferences: "What happens is, when we rate, we're meta-cognitive about quality — that's sort of our aspirational self. It works out much better, to please people, to look at the actual choices that they make." Users say they want documentaries; they watch reality TV. A recommender that trusts stated preferences fails. A search engine that second-guesses the query fails. This is the central tension between the two problems.

It's also why Hastings famously said Netflix "competes with sleep" — not just with other streaming services. The recommender's job is not to satisfy a stated need but to command attention in a world of infinite alternatives. Search competes with ignorance. Recommendation competes with every other possible way to spend time. Different competitions, different rules, different systems.

## The Convergence: UniCoRn (2024)

Before UniCoRn, Netflix ran three separate models for search, homepage recommendations, and "More Like This." UniCoRn replaced all three [7].

```python
class UniCoRn(torch.nn.Module):
    """
    Unified Contextual Recommender: one model, three tasks.

    The insight: treat task identity as a first-class feature.
    """

    def __init__(self, user_vocab_size: int, item_vocab_size: int,
                 embed_dim: int = 512):
        super().__init__()
        self.user_embedding = torch.nn.Embedding(user_vocab_size, embed_dim)
        self.item_embedding = torch.nn.Embedding(item_vocab_size, embed_dim)
        self.query_encoder = TextEncoder(embed_dim)  # encodes search query text
        self.task_embedding = torch.nn.Embedding(3, embed_dim)  # search/recs/similar

        # Shared transformer backbone
        self.shared_backbone = TransformerEncoder(
            d_model=embed_dim, num_layers=6, num_heads=8
        )

        # Task-specific output heads
        self.search_head = torch.nn.Linear(embed_dim, 1)        # P(click | query)
        self.recs_head = torch.nn.Linear(embed_dim, 1)          # P(watch | profile)
        self.similar_head = torch.nn.Linear(embed_dim, 1)       # P(watch | source item)

    def forward(self, user_id: int, item_ids: torch.Tensor,
                query_text: str | None, source_entity_id: int | None,
                task_type: int) -> torch.Tensor:
        """
        task_type: 0 = search, 1 = homepage recs, 2 = More Like This

        Missing context imputation:
        - Search: source_entity_id = null (no source entity)
        - Recs: query_text = null (no explicit query)
        - Similar: query_text = entity title text (imputed)
        """
        user_vec = self.user_embedding(user_id)
        item_vecs = self.item_embedding(item_ids)
        task_vec = self.task_embedding(task_type)

        # Impute missing contexts
        query_vec = (self.query_encoder(query_text)
                     if query_text
                     else torch.zeros_like(user_vec))
        source_vec = (self.item_embedding(source_entity_id)
                      if source_entity_id is not None
                      else torch.zeros_like(user_vec))

        # Shared representation
        combined = torch.cat([user_vec, item_vecs, query_vec, source_vec, task_vec],
                             dim=-1)
        hidden = self.shared_backbone(combined)

        # Task-specific scoring
        if task_type == 0:
            return self.search_head(hidden)
        elif task_type == 1:
            return self.recs_head(hidden)
        else:
            return self.similar_head(hidden)

# Results:
# +7% lift for search tasks
# +10% lift for recommendations tasks
# Fewer models to maintain, shared learnings across tasks
```

The critical detail: UniCoRn succeeded not by ignoring the distinction but by making the model **explicitly aware of it**. `task_type` tells the model which behavior to invoke. The imputation strategy respects the structural difference. The separate output heads optimize for different objectives because, at Netflix, a successful search is not the same thing as a successful recommendation.

Netflix also documented the central tension: **personalization can overpower query relevance**. If a user searches for "documentaries about World War II" and the recommender knows the user loves romantic comedies, should it show *The Notebook*? Obviously not. UniCoRn added personalization incrementally, with guardrails [8].

---

---

## Open Questions

1. **UniCoRn achieved +7% for search and +10% for recommendations by making task identity explicit.** What would the numbers be if they had collapsed the distinction entirely — one model, one objective, no task_type feature? Would it outperform the separate models at all?

2. **Netflix confirmed that scaling laws apply to recommendation foundation models.** Is there a point where a large enough recommender, trained on enough user behavior, internalizes search as a special case — without being told?

3. **Percolate queries (reverse search) solve a problem that only exists when you think of search as infrastructure, not ranking.** What other search primitives are we missing because we default to the ranked-list mental model?

4. **Hastings said Netflix competes with sleep.** If recommendation is competing for attention and search is competing with ignorance, what happens when an LLM can do both? Does the system compete with everything?

**References**

1. Netflix Technology Blog. [*Netflix Studio Search: Using Elasticsearch and Apache Flink to Index Federated GraphQL Data*](https://netflixtechblog.com/studio-search-using-elasticsearch-and-apache-flink-to-index-federated-graphql-data-7d77cad7b7c8). March 2022.

2. Netflix Technology Blog. [*Elasticsearch Indexing Strategy in Asset Management Platform (AMP)*](https://netflixtechblog.com/elasticsearch-indexing-strategy-in-asset-management-platform-amp-99332231e541). 2023.

3. InfoQ. [*Netflix Uses Elasticsearch Percolate Queries to Implement Reverse Searches Efficiently*](https://www.infoq.com/news/2024/04/netflix-percolate-queries/). April 2024.

4. Netflix Technology Blog. [*System Architectures for Personalization and Recommendation*](https://netflixtechblog.com/system-architectures-for-personalization-and-recommendation-e081aa94b5d8). March 2013.

5. Harald Steck et al. [*Deep Learning for Recommender Systems: A Netflix Case Study*](https://doi.org/10.1609/aimag.v42i3.18140). AI Magazine, 42(3): 7–18, 2021.

6. Netflix Technology Blog. [*Foundation Model for Personalized Recommendation*](https://netflixtechblog.com/foundation-model-for-personalized-recommendation-1a0bd8e02d39). March 2025.

7. Moumita Bhattacharya et al. [*Joint Modeling of Search and Recommendations Via an Unified Contextual Recommender (UniCoRn)*](https://dl.acm.org/doi/fullHtml/10.1145/3640457.3688034). RecSys 2024.

8. Sudarshan Lamkhede and Christoph Kofler. [*Recommendations and Results Organization in Netflix Search*](https://dl.acm.org/doi/abs/10.1145/3460231.3474602). RecSys 2021.

---

