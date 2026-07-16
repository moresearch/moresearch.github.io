---
title: "S&R: Why Guessing What People Want Is Harder Than Finding What They Asked For"
date: 2026-07-16
slug: search-recommendation-recommendation-techniques
summary: A technical history of recommender systems with working Python code for each major paradigm: user-based CF, item-based CF, matrix factorization (SVD++), two-tower neural models, and the Netflix Prize story.
tags: [search, recommendation, collaborative-filtering, matrix-factorization, netflix-prize, deep-learning, series]
series: search-recommendation
---

*S&R stands for Search & Recommendation. We trace thirty years of recommender system techniques — from GroupLens to Foundation Models — with working Python code for each paradigm.*

If search is about matching what the user says to what exists, recommendation is about guessing what the user wants before they say it — and, in the hardest cases, before they even know they want it. This is a fundamentally harder information problem. In search, the user tells you what they want and you try to find it. In recommendation, you infer what they want from behavior they may not even be conscious of.

Wu et al. (2023), in their comprehensive survey of neural recommendation models published in IEEE TKDE, organize the field into two broad families: models that use only interaction data (collaborative filtering) and models that incorporate side information (content, context, sequences) [1]. The progression from one to the other mirrors the search field's own evolution — from impoverished signals to rich, multi-modal representations. But the starting point is different. Search began with text and added behavior. Recommendation began with behavior and added text.

This part traces that evolution with working Python.

## Tapestry — The First Collaborative Filtering System (1992)

In 1992, researchers at Xerox PARC built Tapestry, an email filtering system that let users annotate messages and write queries referencing others' annotations. They coined the term **collaborative filtering**. The insight was that relevance is social: if people with similar tastes found something useful, you probably will too.

```python
# Tapestry's conceptual model: manual, query-based CF
# "Show me emails that Bob found interesting"

class TapestryFilter:
    def __init__(self):
        self.annotations: dict[str, dict[int, str]] = {}  # user -> {msg_id -> annotation}

    def annotate(self, user: str, msg_id: int, label: str):
        """Bob annotates a message as 'interesting' or 'boring'."""
        self.annotations.setdefault(user, {})[msg_id] = label

    def query(self, annotator: str, label: str) -> list[int]:
        """'Show me all messages Bob found interesting'."""
        return [msg_id for msg_id, ann in self.annotations.get(annotator, {}).items()
                if ann == label]

# Tapestry required users to write explicit queries. Powerful for power users,
# unusable for everyone else. The automation was coming.
```

## GroupLens — Automating Collaborative Filtering at Scale (1994)

The GroupLens project at the University of Minnesota automated the process. Users rated Usenet articles (1–5 stars), and the system automatically predicted ratings for unread articles based on similar users' ratings.

```python
import numpy as np
from collections import defaultdict

def user_based_cf(ratings: dict[int, dict[int, float]],
                  user_id: int, item_id: int, k: int = 50) -> float:
    """
    User-based collaborative filtering: predict a user's rating for an item
    based on how similar users rated it.

    ratings: {user_id: {item_id: rating}}
    """
    if user_id not in ratings:
        return 3.0  # global mean for cold start

    # Step 1: Find users who rated this item
    co_raters = [(other, ratings[other][item_id])
                 for other in ratings
                 if item_id in ratings[other] and other != user_id]

    if not co_raters:
        return np.mean(list(ratings[user_id].values()))  # user's mean rating

    # Step 2: Compute similarity between target user and each co-rater
    similarities = []
    for other, _ in co_raters:
        sim = pearson_similarity(ratings[user_id], ratings[other])
        similarities.append((other, sim))

    similarities.sort(key=lambda x: x[1], reverse=True)
    neighbors = similarities[:k]

    # Step 3: Weighted average of neighbors' ratings
    weighted_sum = sum(sim * ratings[neighbor][item_id] for neighbor, sim in neighbors
                       if sim > 0)
    norm = sum(abs(sim) for _, sim in neighbors if sim > 0)

    return weighted_sum / norm if norm > 0 else np.mean(list(ratings[user_id].values()))


def pearson_similarity(user_a: dict[int, float], user_b: dict[int, float]) -> float:
    """Pearson correlation between two users' rating vectors."""
    common_items = set(user_a) & set(user_b)
    if len(common_items) < 3:
        return 0.0  # not enough overlap

    mean_a = np.mean([user_a[i] for i in common_items])
    mean_b = np.mean([user_b[i] for i in common_items])

    num = sum((user_a[i] - mean_a) * (user_b[i] - mean_b) for i in common_items)
    den_a = np.sqrt(sum((user_a[i] - mean_a) ** 2 for i in common_items))
    den_b = np.sqrt(sum((user_b[i] - mean_b) ** 2 for i in common_items))

    return num / (den_a * den_b) if den_a and den_b else 0.0
```

GroupLens spawned Net Perceptions, a company that served Amazon, CDnow, and others. In 2010, the team won the ACM Software System Award. But user-based CF had a scaling problem: computing user-user similarity is O(N²) in the number of users.

## Amazon Item-to-Item CF — Scaling to Millions (2003)

Amazon's Greg Linden, Brent Smith, and Jeremy York flipped the problem. Instead of finding similar *users*, find similar *items*:

```python
def item_based_cf(ratings: dict[int, dict[int, float]],
                  user_id: int, k: int = 10) -> list[tuple[int, float]]:
    """
    Item-based collaborative filtering: recommend items similar to what
    the user already liked.

    The expensive part — computing the item-item similarity matrix —
    runs offline, not at serving time.
    """
    # Build item-item similarity matrix (OFFLINE — runs daily)
    item_sim = build_item_similarity_matrix(ratings)

    # Online: look up items similar to what the user liked (MILLISECONDS)
    user_ratings = ratings.get(user_id, {})
    liked = [(item, r) for item, r in user_ratings.items() if r > 3.5]

    candidates: dict[int, float] = defaultdict(float)
    total_weight: dict[int, float] = defaultdict(float)

    for item, rating in liked:
        for similar_item, sim in item_sim.get(item, {}).items():
            if similar_item not in user_ratings:  # don't recommend what they've rated
                candidates[similar_item] += sim * rating
                total_weight[similar_item] += abs(sim)

    # Normalize
    scored = [(item, candidates[item] / total_weight[item])
              for item in candidates if total_weight[item] > 0]
    scored.sort(key=lambda x: x[1], reverse=True)
    return scored[:k]


def build_item_similarity_matrix(ratings: dict[int, dict[int, float]]
                                 ) -> dict[int, dict[int, float]]:
    """O(N_items² × N_users) — run offline, results cached."""
    # Transpose: item -> {user: rating}
    item_users: dict[int, dict[int, float]] = defaultdict(dict)
    for user_id, user_ratings in ratings.items():
        for item_id, rating in user_ratings.items():
            item_users[item_id][user_id] = rating

    items = list(item_users.keys())
    sim_matrix: dict[int, dict[int, float]] = defaultdict(dict)

    for i, item_a in enumerate(items):
        for item_b in items[i + 1:]:
            common = set(item_users[item_a]) & set(item_users[item_b])
            if len(common) < 5:
                continue
            sim = cosine_similarity(
                {u: item_users[item_a][u] for u in common},
                {u: item_users[item_b][u] for u in common}
            )
            if sim > 0:
                sim_matrix[item_a][item_b] = sim
                sim_matrix[item_b][item_a] = sim

    return sim_matrix
```

Greg Linden later explained why this beat search-based approaches: "Rather than matching the user to similar customers, item-to-item collaborative filtering matches each of the user's purchased and rated items to similar items." Search-based methods — constructing queries from purchase history to find items with similar keywords — produced recommendations that were either too general (bestsellers) or too narrow (more books by the same author). Collaborative filtering discovered cross-category connections: people who bought *Into Thin Air* also bought *The Perfect Storm*. No keyword match connects a mountaineering disaster to a fishing boat tragedy. Only behavior does.

## The Netflix Prize — Matrix Factorization Takes Over (2006–2009)

In 2006, Netflix offered $1M to any team that could improve their recommendation algorithm by 10%. The winning approach reshaped the field: **matrix factorization**.

```python
class MatrixFactorization:
    """Learn user and item latent factors via stochastic gradient descent."""

    def __init__(self, num_users: int, num_items: int, k: int = 50,
                 lr: float = 0.01, lambda_reg: float = 0.02):
        # Initialize latent factor matrices randomly
        self.P = np.random.normal(0, 0.1, (num_users, k))  # user factors
        self.Q = np.random.normal(0, 0.1, (num_items, k))  # item factors
        self.bu = np.zeros(num_users)  # user biases
        self.bi = np.zeros(num_items)  # item biases
        self.mu = 0.0                  # global mean
        self.lr = lr
        self.lambda_reg = lambda_reg

    def fit(self, ratings: list[tuple[int, int, float]], epochs: int = 100):
        """Train via SGD."""
        self.mu = np.mean([r for _, _, r in ratings])

        for epoch in range(epochs):
            np.random.shuffle(ratings)
            total_loss = 0.0

            for u, i, r in ratings:
                # Prediction: global mean + user bias + item bias + latent interaction
                pred = self.mu + self.bu[u] + self.bi[i] + np.dot(self.P[u], self.Q[i])
                error = r - pred
                total_loss += error ** 2

                # SGD updates with L2 regularization
                self.bu[u] += self.lr * (error - self.lambda_reg * self.bu[u])
                self.bi[i] += self.lr * (error - self.lambda_reg * self.bi[i])

                # Update latent factors
                pu_old = self.P[u].copy()
                self.P[u] += self.lr * (error * self.Q[i] - self.lambda_reg * self.P[u])
                self.Q[i] += self.lr * (error * pu_old - self.lambda_reg * self.Q[i])

            if epoch % 20 == 0:
                rmse = np.sqrt(total_loss / len(ratings))
                print(f"Epoch {epoch}: RMSE = {rmse:.4f}")

    def predict(self, u: int, i: int) -> float:
        return self.mu + self.bu[u] + self.bi[i] + np.dot(self.P[u], self.Q[i])

    def recommend(self, u: int, rated_items: set[int], k: int = 10) -> list[int]:
        """Generate top-k recommendations for user u."""
        scores = [(i, self.predict(u, i))
                  for i in range(len(self.Q))
                  if i not in rated_items]
        scores.sort(key=lambda x: x[1], reverse=True)
        return [item for item, _ in scores[:k]]

# Usage with MovieLens-100K style data
# ratings = [(user, item, rating), ...]
# mf = MatrixFactorization(num_users=943, num_items=1682, k=50)
# mf.fit(ratings, epochs=100)
# recommendations = mf.recommend(user_id=42, rated_items={item for _, item, _ in ratings if _ == 42})
```

Yehuda Koren's **SVD++** extended this by incorporating implicit feedback — what you browsed, not just what you rated — and achieved the winning RMSE of 0.8556. The Netflix Prize established matrix factorization as the dominant paradigm for nearly a decade.

It also revealed the field's central tension: **the metric that drove the competition (RMSE on withheld ratings) doesn't actually measure whether users are satisfied**. Reed Hastings articulated this later: "When we rate, we're meta-cognitive about quality — that's sort of our aspirational self. It works out much better, to please people, to look at the actual choices that they make." Users say they want documentaries; they watch reality TV. A recommender that trusts stated preferences over revealed preferences fails.

## The Deep Learning Era — YouTube and the Multi-Stage Pipeline (2016–2020)

YouTube's 2016 paper marked deep learning's entry into production recommendation. The architecture was two-stage:

```python
class YouTubeCandidateGenerator(torch.nn.Module):
    """YouTube-style candidate generation: narrow millions → hundreds."""

    def __init__(self, vocab_sizes: dict[str, int], embedding_dim: int = 256):
        super().__init__()
        # Embedding layers for categorical features
        self.embeddings = torch.nn.ModuleDict({
            name: torch.nn.Embedding(size, embedding_dim)
            for name, size in vocab_sizes.items()
        })
        # Concatenated embeddings → hidden layers
        input_dim = len(vocab_sizes) * embedding_dim
        self.layers = torch.nn.Sequential(
            torch.nn.Linear(input_dim, 1024),
            torch.nn.ReLU(),
            torch.nn.Linear(1024, 512),
            torch.nn.ReLU(),
            torch.nn.Linear(512, 256),
        )
        # Output: softmax over all video IDs (treated as classes)
        self.output = torch.nn.Linear(256, vocab_sizes['video_id'])

    def forward(self, features: dict[str, torch.Tensor]) -> torch.Tensor:
        """Predict the next video a user will watch."""
        embedded = [self.embeddings[name](features[name])
                    for name in self.embeddings]
        concat = torch.cat(embedded, dim=1)
        hidden = self.layers(concat)
        return self.output(hidden)

# Key innovation: the model was trained on ALL YouTube watches, including those
# on embedded players. The training objective was next-video prediction —
# treating recommendation as an extreme multi-class classification problem.

# But the real insight was about WHAT to optimize. YouTube didn't optimize for
# clicks — they optimized for expected watch time. Clicks are easy to game
# (clickbait thumbnails). Watch time is harder to fake and better aligned
# with user satisfaction.
```

The three-stage pipeline that emerged became industry standard:

```python
def production_recsys_pipeline(user_id: int, user_features: dict,
                               all_items: list[int], k: int = 10) -> list[int]:
    """Standard three-stage recommendation pipeline."""

    # Stage 1: Candidate Generation — millions → thousands
    # Multiple parallel generators: collaborative filtering, trending,
    # new releases, content-based. High recall, low cost per item.
    candidates_cf = cf_retriever.retrieve(user_id, k=500)
    candidates_trending = trending_retriever.retrieve(k=200)
    candidates_similar = similar_items_retriever.retrieve(
        user_features['last_watched'], k=200)
    candidates = list(set(candidates_cf + candidates_trending + candidates_similar))

    # Stage 2: Ranking — thousands → hundreds
    # Deep neural network scores each candidate using hundreds of features.
    # High precision, moderate cost per item.
    ranked = ranker.score(user_id, candidates, user_features)
    ranked.sort(key=lambda x: x.score, reverse=True)
    top = ranked[:200]

    # Stage 3: Re-Ranking — hundreds → tens
    # Apply diversity, freshness boost, business rules, exploration.
    # Low cost per item (post-processing), high impact on user experience.
    final = re_ranker.apply(top, user_features,
                            diversity_factor=0.3,
                            freshness_boost=1.2,
                            max_same_genre=3)
    return [item.id for item in final[:k]]
```

By 2020, the search field had always been multi-stage. Recommendation caught up — and that convergence is one reason the two fields are so often conflated.

---

---

## Open Questions

1. **Matrix factorization dominated recommendation for a decade after the Netflix Prize, but the Prize's metric (RMSE) didn't measure user satisfaction.** What would a recommendation competition look like today if the metric were retention, not rating prediction? Could we even run one?

2. **Hastings' insight — that users' stated preferences differ from their revealed preferences — has uncomfortable implications.** Should recommenders ever *ignore* what users explicitly tell them? When is the aspirational self a feature rather than noise?

3. **The deep learning era equalized recommendation and search by giving both access to heterogeneous features.** But search had decades of multi-modal infrastructure first. Did recommendation catch up, or did it just inherit search's architecture without adapting it?

4. **Item-to-item CF at Amazon discovered cross-category connections no keyword match could find.** What connections are today's models missing because they optimize for engagement rather than surprise?

**References**

1. Le Wu, Xiangnan He, Xiang Wang, Kun Zhang, and Meng Wang. [*A Survey on Accuracy-Oriented Neural Recommendation: From Collaborative Filtering to Information-Rich Recommendation*](https://doi.org/10.1109/TKDE.2022.3145690). IEEE TKDE, 35(5): 4425–4445, 2023.

2. David Goldberg, David Nichols, Brian M. Oki, and Douglas Terry. [*Using Collaborative Filtering to Weave an Information Tapestry*](https://dl.acm.org/doi/10.1145/138859.138867). Communications of the ACM, 35(12): 61–70, 1992.

3. Paul Resnick et al. [*GroupLens: An Open Architecture for Collaborative Filtering of Netnews*](https://dl.acm.org/doi/10.1145/192844.192905). CSCW 1994.

4. Greg Linden, Brent Smith, and Jeremy York. [*Amazon.com Recommendations: Item-to-Item Collaborative Filtering*](https://doi.org/10.1109/MIC.2003.1167344). IEEE Internet Computing, 7(1): 76–80, 2003.

5. Yehuda Koren, Robert Bell, and Chris Volinsky. [*Matrix Factorization Techniques for Recommender Systems*](https://doi.org/10.1109/MC.2009.263). IEEE Computer, 42(8): 30–37, 2009.

6. Paul Covington, Jay Adams, and Emre Sargin. [*Deep Neural Networks for YouTube Recommendations*](https://dl.acm.org/doi/10.1145/2959100.2959190). RecSys 2016.

---

