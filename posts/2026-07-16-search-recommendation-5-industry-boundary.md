---
title: "S&R: Where Should You Draw the Line Between Search and Recommendation?"
date: 2026-07-16
slug: search-recommendation-industry-boundary
summary: How four leading engineering organizations handle the search–recommendation boundary in production: Spotify's LLM intent router, DoorDash's content-first embeddings and Consumer Memory Blocks, Airbnb's listing embeddings for real-time search personalization, and Pinterest's two-tower architectures.
tags: [search, recommendation, spotify, doordash, airbnb, pinterest, two-tower, embeddings, series]
series: search-recommendation
---

*S&R stands for Search & Recommendation. How Spotify, DoorDash, Airbnb, and Pinterest navigate the boundary between search and recommendation in production — with code for each architecture.*

Netflix isn't the only company navigating the search–recommendation boundary. Spotify, DoorDash, Airbnb, and Pinterest each handle it differently — and their engineering blogs document the trade-offs. This part examines four production architectures.

## Spotify: Intent-Based Routing with LLMs

Spotify's 2025 paper *"You Say Search, I Say Recs"* describes an LLM-based router that classifies user intent and dispatches accordingly [1]:

```python
from enum import Enum

class QueryIntent(Enum):
    NAVIGATIONAL = "navigational"    # "find song X" → search
    EXPLORATORY = "exploratory"      # "Italian 80s disco nostalgia" → recs
    MIXED = "mixed"                  # both paths, parallel → fused results

class SpotifyIntentRouter:
    """
    LLM-based router: classify intent, route to search or recommendation.

    Key design decisions:
    - Only the query goes to the router (NOT user features) → max cache hits
    - User features go to downstream tools → personalization where it belongs
    - Small distilled LLM: 60% latency reduction, 99% cost reduction vs teacher
    """

    def __init__(self, llm_model):
        self.llm = llm_model  # small fine-tuned model, p75 latency ~450ms

    def classify(self, query: str) -> QueryIntent:
        """Classify query intent. Cached aggressively — no user features."""
        if cached := self.cache.get(query):
            return cached
        intent = self.llm.classify(query)  # "navigational", "exploratory", "mixed"
        self.cache.set(query, intent)
        return intent

    def route(self, query: str, user_features: dict) -> SearchResult:
        intent = self.classify(query)

        if intent == QueryIntent.NAVIGATIONAL:
            return self.search_backend.search(query, user_features)
            # Elasticsearch + BM25 + Voyager ANN

        elif intent == QueryIntent.EXPLORATORY:
            return self.recs_backend.recommend(query, user_features)
            # Collaborative filtering + content-based + Voyager ANN

        else:  # MIXED — both paths in parallel, fuse results
            search_result = self.search_backend.search_async(query, user_features)
            recs_result = self.recs_backend.recommend_async(query, user_features)
            return self.fuse(search_result, recs_result)

# Results vs. regular search:
# +115% for finding similar artists
# +91% for new music discovery
# +25% for broad music searches

# Spotify's Unified Embedding Infrastructure (Voyager):
# HNSW-based, 10x faster than Annoy at same recall.
# Powers both search AND recommendation — shared infrastructure,
# task-specific usage patterns.
```

Spotify's key finding: **Semantic IDs optimized for search don't generalize to recommendation**, and vice versa. A multi-task bi-encoder achieves a Pareto-optimal trade-off, but you can't optimize one embedding space for both tasks without compromise [2].

## DoorDash: Content-First Embeddings, Task-Aware Retrieval

DoorDash's environment is intent-driven and transactional. A click tells you less than you think:

```python
# DoorDash's core insight: clicks are poor proxies for semantics.
# A click on a Sichuan noodle soup doesn't distinguish spicy preference
# from noodle preference — or from just being hungry.

# Their solution: LLMs generate rich content profiles, embeddings follow.

class DoorDashContentPipeline:
    """Content-first: LLM profiles → off-the-shelf encoder → ANN retrieval."""

    def generate_item_profile(self, item: dict) -> str:
        """LLM produces a standardized narrative for every item."""
        prompt = f"""
        Describe this menu item for a food recommendation system:
        Name: {item['name']}
        Category: {item['category']}
        Ingredients: {item['ingredients']}
        Preparation: {item['preparation']}
        Cuisine type: {item['cuisine']}

        Cover: ingredients, preparation method, cuisine attributes,
        dietary properties (spicy, vegetarian, etc.), eating context,
        flavor profile.
        """
        return self.llm.generate(prompt)

    def embed(self, profile: str) -> np.ndarray:
        """Encode the profile with an off-the-shelf encoder."""
        return self.encoder.encode(profile)  # gemini-embedding-001, 256-dim MRL

    def retrieve(self, query_embedding: np.ndarray, k: int = 100) -> list[int]:
        """ANN search over pre-computed item embeddings."""
        return self.milvus_index.search(query_embedding, k)

    # Key finding: upgrading profile quality → +31% improvement.
    # Upgrading the encoder on raw metadata → only +6%.
    # Data quality dominates model choice.

# Results:
# -3.65% null search rate, +0.66% CVR, +0.072% 7D active customers
```

For recommendations, DoorDash developed **Consumer Memory Blocks** — typed, composable representations of everything known about a user:

```python
class ConsumerMemoryBlock:
    """
    Structured, namespaced user state serialized as JSON for LLM prompt input.

    Properties:
    - Composable: different use cases request different sub-blocks
    - Evidenced, not inferred: derived from observed behavior with provenance
    - Extensible: new sub-blocks without downstream disruption
    """

    def build(self, user_id: str, sub_blocks: list[str]) -> dict:
        blocks = {}
        if "long_term_preferences" in sub_blocks:
            blocks["long_term"] = {
                "cuisines": ["Thai", "Sichuan", "Italian"],
                "dietary": ["prefers spicy", "avoids dairy"],
                "price_range": "$$-$$$",
                "avg_order_value": 42.50,
            }
        if "behavioral_patterns" in sub_blocks:
            blocks["patterns"] = {
                "order_days": ["Fri", "Sat"],
                "peak_time": "19:00-21:00",
                "group_orders": True,
                "repeat_rate": 0.35,
            }
        return blocks

    def to_prompt(self, blocks: dict) -> str:
        """Serialize blocks as compact JSON for the LLM carousel generator."""
        return json.dumps(blocks)

# Generated carousels are embedded OFFLINE, retrieved via Milvus ONLINE.
# No LLM in the request path — cost amortized across the refresh interval.
# Results: +2.4% order rate.
```

## Airbnb: Embeddings That Bridge Search and Discovery

Airbnb's KDD 2018 paper trained listing embeddings on 800 million search click sessions using skip-gram adapted from word2vec [3]:

```python
class AirbnbListingEmbeddings:
    """Domain-specific embedding training for travel search."""

    def generate_training_pairs(self, sessions: list[list[int]]) -> list[tuple[int, int]]:
        """
        Key innovations over standard word2vec:

        1. In-market negative sampling: users search within a single city.
           Negatives drawn globally are trivially easy to reject.
           Draw negatives from the SAME MARKET for harder discrimination.

        2. Booked listing as global context: the booked listing is ALWAYS
           treated as the context being predicted, regardless of position in
           the click sequence. The booking is the signal. Everything else is noise.
        """
        pairs = []
        for session in sessions:
            booked_id = session[-1]  # last item is the booking
            for clicked_id in session[:-1]:
                pairs.append((clicked_id, booked_id))
        return pairs

    def compute_serving_features(self, user_id: int, candidates: list[int],
                                 embeddings: dict[int, np.ndarray]) -> dict[str, np.ndarray]:
        """
        Embedding-based features for real-time ranking.

        Five embedding features ranked among the TOP 20 of 104 total features.
        """
        user_recent = self.get_recent_interactions(user_id)
        return {
            # Similarity to listings the user recently clicked
            "EmbClickSim": self.mean_sim(user_recent['clicked'], candidates, embeddings),
            # Similarity to listings the user skipped
            "EmbSkipSim": self.mean_sim(user_recent['skipped'], candidates, embeddings),
            # Similarity to the last listing the user spent significant time on
            "EmbLastLongClickSim": self.mean_sim([user_recent['last_long_click']],
                                                  candidates, embeddings),
        }

# Results: +2.27% offline NDCG, statistically significant booking gain online.
# Similar listing recommendations: +20% CTR over prior algorithm.
```

Airbnb's 2019 follow-up on deep learning documented instructive failures: listing ID embeddings overfit (too few bookings per listing), and multi-task learning for bookings + long views increased views but not bookings — because expensive listings get looked at but not booked [4].

## Pinterest: Two-Tower Architectures for Homefeed and Search

Pinterest's two-tower model powers both homefeed and search:

```python
class PinterestTwoTower(torch.nn.Module):
    """Two-tower: user tower and item tower, dot-product scoring."""

    def __init__(self, user_input_dim: int, item_input_dim: int, hidden: int = 256):
        super().__init__()
        # User tower: long-term history, context, real-time sequences
        self.user_tower = torch.nn.Sequential(
            torch.nn.Linear(user_input_dim, hidden * 4),
            torch.nn.ReLU(),
            MaskNet(hidden * 4, num_blocks=3),   # bitwise feature crossing
            torch.nn.Linear(hidden * 4, hidden),
        )
        # Item tower: category, description, image features
        self.item_tower = torch.nn.Sequential(
            torch.nn.Linear(item_input_dim, hidden * 4),
            torch.nn.ReLU(),
            torch.nn.Linear(hidden * 4, hidden),
        )

    def forward(self, user_features, item_features):
        user_emb = F.normalize(self.user_tower(user_features), dim=1)
        item_emb = F.normalize(self.item_tower(item_features), dim=1)
        return (user_emb * item_emb).sum(dim=1)  # dot product

# Item embeddings pre-computed offline, indexed in Manas (HNSW-based ANN).
# User tower runs once per request at serving time.
# This decoupling is the key to ~3ms median latency at 300K QPS.
```

For search, Pinterest extended this into **OmniSearchSage** — a multi-task, multi-entity framework where a single unified query embedding retrieves pins, products, and related queries simultaneously [5]. A teacher-student distillation approach followed: an 8B Llama cross-encoder (teacher, +12–20% improvement) → bi-encoder student trained on 100× more data from daily search logs (85% query cache hit rate).

## The Pattern Across All Four Companies

Every system that successfully spans search and recommendation does so by encoding **which task it's doing** as a first-class signal. Spotify uses an explicit router. DoorDash uses separate retrieval pipelines with shared embeddings. Airbnb uses the same embeddings but different ranking features for search vs. discovery. Pinterest uses the same two-tower but different objectives for homefeed vs. search.

The common pattern: **unification where it reduces cost, separation where it preserves correctness**.

---

---

## Open Questions

1. **Spotify found that Semantic IDs optimized for search don't generalize to recommendation, and vice versa.** Is this a fundamental property of the two tasks, or an artifact of how we train embeddings? What would a *truly* unified embedding space look like?

2. **DoorDash's finding — data quality (+31%) dominates model choice (+6%) — echoes a broader truth.** How much of the search-vs-recommendation gap is actually a *data quality* gap in disguise?

3. **Every successful system in this survey encodes task identity as a first-class signal.** Is the search/recommendation distinction a permanent architectural feature, or a temporary crutch that better models will eventually absorb?

4. **Airbnb's failed multi-task experiment (bookings + long views) is a warning about proxy objectives.** What other proxy objectives are we optimizing across the industry that don't actually measure what we think they measure?

**References**

1. Spotify Research. [*You Say Search, I Say Recs: A Scalable Agentic Approach to Query Understanding and Exploratory Search*](https://research.atspotify.com/2025/9/you-say-search-i-say-recs-a-scalable-agentic-approach-to-query-understanding). September 2025.

2. Spotify Research. [*Semantic IDs for Generative Search and Recommendation*](https://www.research.atspotify.com/2025/9/semantic-ids-for-generative-search-and-recommendation). September 2025.

3. Mihajlo Grbovic and Haibin Cheng. [*Real-time Personalization using Embeddings for Search Ranking at Airbnb*](https://dl.acm.org/doi/10.1145/3219819.3219885). KDD 2018.

4. Malay Haldar et al. [*Applying Deep Learning to Airbnb Search*](https://arxiv.org/abs/1810.09591). KDD 2019.

5. DoorDash Engineering. [*Using LLMs to Build Content Embeddings for Search and Recommendations*](https://careersatdoordash.com/blog/doordash-llms-to-build-content-embeddings-for-search-and-recommendations/). 2025.

6. DoorDash Engineering. [*Offline LLMs, Online Personalization: Generating Carousels at DoorDash*](https://careersatdoordash.com/blog/doordash-offline-llms-online-personalization-generating-carousels/). 2025.

7. Pinterest Engineering. [*Advancements in Embedding-Based Retrieval at Pinterest Homefeed*](https://medium.com/pinterest-engineering/advancements-in-embedding-based-retrieval-at-pinterest-homefeed-176a3c03df64). 2024.

---

