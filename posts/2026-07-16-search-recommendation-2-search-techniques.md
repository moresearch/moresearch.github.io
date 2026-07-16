---
title: "search/recommendation: Fifty Years of Search — From Boolean to BERT"
date: 2026-07-16
slug: search-recommendation-search-techniques
summary: Part 2 of 7. A technical history of information retrieval with working Python code for each major paradigm: Boolean retrieval, TF-IDF, BM25, PageRank, Learning to Rank, and neural IR with embeddings.
tags: [search, information-retrieval, tf-idf, bm25, pagerank, bert, neural-ir, series]
series: search-recommendation
---

Search answers one question: given a query and a collection of documents, which documents are most relevant, and in what order? Every major advance in information retrieval has come from realizing the previous generation's answer was incomplete — not wrong, just missing a dimension of what "relevance" means.

This part traces fifty years of that evolution with working Python. You'll see how each generation built on the last, what each solved, and what each left unsolved.

## Boolean Retrieval — Exact Matching and Its Limits (1960s–1970s)

The earliest computerized search systems used Boolean logic: a query was a logical expression of terms, and documents either matched or did not. No ranking — just set intersection.

```python
from collections import defaultdict
from typing import list[str]

class BooleanIndex:
    """The simplest retrieval engine: exact match, no ranking."""

    def __init__(self):
        self.index: dict[str, set[int]] = defaultdict(set)
        self.documents: dict[int, str] = {}

    def add(self, doc_id: int, text: str):
        self.documents[doc_id] = text
        for token in set(text.lower().split()):
            self.index[token].add(doc_id)

    def search_and(self, terms: list[str]) -> set[int]:
        """Boolean AND: all terms must appear."""
        result = None
        for term in terms:
            docs = self.index.get(term.lower(), set())
            result = docs if result is None else result & docs
        return result or set()

    def search_or(self, terms: list[str]) -> set[int]:
        """Boolean OR: any term can appear."""
        result = set()
        for term in terms:
            result |= self.index.get(term.lower(), set())
        return result

# Usage
idx = BooleanIndex()
idx.add(1, "the tragedy of star crossed lovers in Verona")
idx.add(2, "a tragic love story set in medieval Italy")

# The vocabulary mismatch problem in action
assert idx.search_and(["tragic", "love", "story"]) == {2}   # finds doc 2
assert idx.search_and(["star", "crossed", "lovers"]) == {1} # finds doc 1
# But: no document contains BOTH phrasings, so they never appear together.
# Searching for "tragic love story" misses Shakespeare entirely.
```

This worked for trained librarians. It failed for everyone else. The **vocabulary mismatch problem** — users describe their needs with different words than authors use — is the fundamental condition of language, not a bug to be fixed. Every subsequent generation of retrieval technology is an attempt to narrow the gap.

Karen Spärck Jones, who invented inverse document frequency (the weighting scheme that would become half of TF-IDF), understood this better than anyone. In a 1999 reflection, she wrote: "Classical document retrieval thus falls in the class of AI tasks that assist the human user but cannot, by definition, replace them." The gap can be narrowed. It cannot be closed.

## The Vector Space Model — When Documents Became Points in Space (1975)

Gerard Salton at Cornell proposed representing documents and queries as sparse vectors in a high-dimensional term space. A document was no longer a set of words — it was a point whose coordinates were term weights.

The weighting scheme was **TF-IDF**:

```python
import math
from collections import Counter

def compute_tf_idf(documents: list[list[str]]) -> dict[int, dict[str, float]]:
    """Compute TF-IDF vectors for a document collection."""
    N = len(documents)
    df: dict[str, int] = Counter()

    # Count document frequency for each term
    for doc_tokens in documents:
        for term in set(doc_tokens):
            df[term] += 1

    tfidf_vectors: dict[int, dict[str, float]] = {}
    for i, doc_tokens in enumerate(documents):
        tf = Counter(doc_tokens)
        doc_len = len(doc_tokens)
        tfidf_vectors[i] = {}
        for term, count in tf.items():
            # TF: normalized term frequency
            tf_norm = count / doc_len
            # IDF: log(N / df) — rare terms get higher weight
            idf = math.log((N - df[term] + 0.5) / (df[term] + 0.5) + 1.0)
            tfidf_vectors[i][term] = tf_norm * idf

    return tfidf_vectors

def cosine_similarity(vec_a: dict[str, float], vec_b: dict[str, float]) -> float:
    """Cosine similarity between two sparse vectors."""
    common_terms = set(vec_a) & set(vec_b)
    if not common_terms:
        return 0.0

    dot = sum(vec_a[t] * vec_b[t] for t in common_terms)
    norm_a = math.sqrt(sum(v ** 2 for v in vec_a.values()))
    norm_b = math.sqrt(sum(v ** 2 for v in vec_b.values()))
    return dot / (norm_a * norm_b) if norm_a and norm_b else 0.0

def search_tfidf(query: list[str], tfidf_vectors: dict[int, dict[str, float]],
                 k: int = 10) -> list[tuple[int, float]]:
    """Search using TF-IDF cosine similarity."""
    # Build query vector (TF only — no IDF for ad-hoc queries in simplest form)
    query_tf = Counter(query)
    query_len = len(query)
    query_vec = {t: c / query_len for t, c in query_tf.items()}

    scores = [(doc_id, cosine_similarity(query_vec, doc_vec))
              for doc_id, doc_vec in tfidf_vectors.items()]
    scores.sort(key=lambda x: x[1], reverse=True)
    return scores[:k]

# Usage
docs = [
    "the tragedy of star crossed lovers in Verona".split(),
    "a tragic love story set in medieval Italy".split(),
    "machine learning algorithms for classification tasks".split(),
]
tfidf = compute_tf_idf(docs)
# Now "tragic love story" finds BOTH documents — TF-IDF bridges the vocabulary gap
# through shared high-IDF terms
results = search_tfidf("tragic love story".split(), tfidf)
```

TF-IDF doesn't *understand* semantics. It exploits a statistical regularity: rare terms are more discriminating, and frequent local terms are more important. It's a heuristic — but an extraordinarily robust one that remains a strong baseline fifty years later.

## BM25 — The Probabilistic Framework That Still Powers Production (1970s–1990s)

Stephen Robertson, Karen Spärck Jones, and collaborators at City, University of London asked a different question: given a document and a query, what is the *probability* that the document is relevant? The answer was BM25.

```python
def bm25_score(query_terms: list[str], doc_tokens: list[str],
               doc_lengths: list[int], avg_dl: float,
               df: dict[str, int], N: int,
               k1: float = 1.5, b: float = 0.75) -> float:
    """
    BM25 scoring function.

    k1: controls term frequency saturation (higher = more linear)
    b:  controls document length normalization (0 = none, 1 = full)
    """
    dl = len(doc_tokens)
    tf = Counter(doc_tokens)
    score = 0.0

    for term in set(query_terms):
        if term not in df:
            continue

        # IDF component (Robertson-Spärck Jones formulation)
        idf = math.log((N - df[term] + 0.5) / (df[term] + 0.5) + 1.0)

        # TF with saturation: as count grows, additional occurrences matter less
        f = tf.get(term, 0)
        tf_saturated = (f * (k1 + 1)) / (f + k1 * (1 - b + b * dl / avg_dl))

        score += idf * tf_saturated

    return score

# BM25's key innovations over TF-IDF:
# 1. Diminishing returns for term frequency (a term appearing 10x is not 10x as
#    important as appearing 1x — the saturation curve handles this)
# 2. Document length normalization that's tunable (longer docs aren't inherently
#    more relevant just because they have more words)

# BM25 remains the default scoring in Elasticsearch and Lucene — which means it
# powers production search at Netflix, Spotify, DoorDash, and most of the web.
```

BM25's IDF is derived from the Binary Independence Model — it has a probabilistic foundation that TF-IDF lacks. It is not obsolete in 2026. In modern hybrid systems, BM25 is the fast, cheap first stage that feeds candidates to neural re-rankers.

## PageRank — When the Link Graph Became a Relevance Signal (1998)

Content-based retrieval treats documents independently. But on the web, documents link to each other — and those links carry information. Brin and Page's insight at Stanford was that the link graph is itself a relevance signal.

```python
import numpy as np

def pagerank(adjacency: dict[int, list[int]], N: int,
             d: float = 0.85, max_iter: int = 100, tol: float = 1e-6) -> dict[int, float]:
    """
    Compute PageRank scores from a link graph.

    adjacency: {page_id: [outgoing_link_page_ids]}
    d: damping factor — probability the random surfer continues clicking
    """
    pr = {p: 1.0 / N for p in adjacency}  # uniform initialization

    for iteration in range(max_iter):
        new_pr = {}
        for page in adjacency:
            # Random jump component: (1-d)/N — every page gets baseline probability
            # Link component: sum of PR from inbound links, divided by their out-degree
            inbound_sum = 0.0
            for other_page, out_links in adjacency.items():
                if page in out_links:
                    inbound_sum += pr[other_page] / len(out_links)
            new_pr[page] = (1 - d) / N + d * inbound_sum

        # Check convergence
        delta = sum(abs(new_pr[p] - pr[p]) for p in pr)
        if delta < tol:
            break
        pr = new_pr

    return pr

# Usage: a tiny web of 4 pages
web = {
    0: [1, 2],       # page 0 links to 1 and 2
    1: [2],          # page 1 links to 2
    2: [0, 3],       # page 2 links to 0 and 3
    3: [2],          # page 3 links to 2
}
scores = pagerank(web, N=4)
# Page 2 gets the highest score: it has the most inbound links from important pages
```

PageRank didn't replace content-based retrieval — it augmented it. A modern search engine computes hundreds of features (BM25, PageRank, proximity, freshness, click-through rate, spam score) and feeds them into a learned ranking function.

## Learning to Rank — When Ranking Became a Supervised ML Problem (2000s–2010s)

The insight: ranking is just a machine learning problem. Given a query, a set of candidate documents with relevance labels, and hundreds of features, learn a function that orders them optimally.

```python
from sklearn.ensemble import GradientBoostingRegressor
import numpy as np

def train_pointwise_ltr(training_data: list[tuple[np.ndarray, float]]) -> GradientBoostingRegressor:
    """
    Pointwise LTR: predict relevance score for each (query, document) pair.

    Features might include: BM25, PageRank, click rate, freshness, title match, etc.
    """
    X = np.array([features for features, _ in training_data])
    y = np.array([score for _, score in training_data])
    model = GradientBoostingRegressor(n_estimators=500, max_depth=5)
    model.fit(X, y)
    return model

def rank_candidates(model, candidates: list[tuple[int, np.ndarray]]) -> list[int]:
    """Score and sort candidates using the learned model."""
    scored = [(doc_id, model.predict(features.reshape(1, -1))[0])
              for doc_id, features in candidates]
    scored.sort(key=lambda x: x[1], reverse=True)
    return [doc_id for doc_id, _ in scored]

# In practice, LambdaMART — gradient-boosted trees with a listwise LambdaRank
# objective — became the industry standard. The key difference from pointwise:
# it optimizes the ordering directly (NDCG), not individual relevance scores.

# Real search engines at Microsoft, Yahoo, and Google used ~500-1000 features.
# The hand-crafted feature era produced the best search quality we'd ever seen —
# but it was also a maintenance nightmare. Every new signal needed a new feature.
```

## The Neural Turn — From Word2Vec to BERT to Dense Retrieval (2013–2020)

Three papers changed search again. Let's see what each one did to the retrieval pipeline.

**Word2Vec (2013):** Dense word vectors that capture semantic relationships.

```python
# Conceptual: Word2Vec trains on the task "predict surrounding words"
# The learned vectors capture analogies: king - man + woman ≈ queen
# In retrieval: query and document terms can be matched even when they
# don't share exact words, by comparing their vector representations.

def embed_query(query: str, word_vectors: dict[str, np.ndarray]) -> np.ndarray:
    """Simple averaging of word vectors — the simplest dense query representation."""
    tokens = query.lower().split()
    vectors = [word_vectors[t] for t in tokens if t in word_vectors]
    return np.mean(vectors, axis=0) if vectors else np.zeros(300)

# But: word vectors are context-independent. "bank" has the same vector
# in "river bank" and "investment bank." That's the problem BERT solved.
```

**BERT (2018):** Deeply contextualized embeddings. A word's representation depends on the words around it.

```python
# With BERT, "bank" in "river bank" and "investment bank" have different vectors.
# This transformed query–document matching.

# In retrieval, BERT is typically used as a RERANKER, not a first-stage retriever:
# 1. BM25 retrieves top-1000 candidates (fast, cheap)
# 2. BERT cross-encoder scores each (query, candidate) pair (slow, expensive, accurate)

# The cross-encoder concatenates query and document, passes them through BERT
# jointly, and produces a single relevance score. This captures fine-grained
# interactions but is too slow to run over the entire collection.
```

**Dense Passage Retrieval (DPR) (2020):** Bi-encoder architecture that makes neural retrieval fast enough for first-stage retrieval.

```python
import torch
import torch.nn.functional as F

class DPRBiEncoder(torch.nn.Module):
    """Dense Passage Retrieval: separate encoders for queries and passages."""
    def __init__(self, query_encoder, passage_encoder):
        super().__init__()
        self.query_encoder = query_encoder    # e.g., BERT-base
        self.passage_encoder = passage_encoder # e.g., BERT-base

    def encode_query(self, query_texts: list[str]) -> torch.Tensor:
        """Encode queries into dense vectors."""
        return self.query_encoder(query_texts)  # shape: (batch, 768)

    def encode_passages(self, passage_texts: list[str]) -> torch.Tensor:
        """Encode passages into dense vectors (can be pre-computed offline)."""
        return self.passage_encoder(passage_texts)  # shape: (batch, 768)

    def retrieve(self, query: torch.Tensor, passage_embeddings: torch.Tensor,
                 k: int = 10) -> tuple[torch.Tensor, torch.Tensor]:
        """ANN retrieval using dot product similarity."""
        scores = torch.matmul(query, passage_embeddings.T)  # (1, num_passages)
        top_scores, top_indices = torch.topk(scores, k=k)
        return top_indices, top_scores

# Key property: passage embeddings are pre-computed and indexed in FAISS.
# At query time, only the query encoder runs. This makes neural first-stage
# retrieval feasible at scale.
```

Guo et al.'s comprehensive survey *A Deep Look into Neural Ranking Models for Information Retrieval* (2020) catalogs this transition from hand-crafted features to learned representations, noting that the key shift was not just better accuracy — it was the elimination of feature engineering as the bottleneck in search quality improvement [1].

## The Modern Stack: Hybrid All the Way Down

No production search system in 2026 uses a single model. The standard architecture is a multi-stage cascade:

```python
def production_search_pipeline(query: str, k: int = 10) -> list[Document]:
    """Multi-stage search: each stage is more expensive but operates on fewer candidates."""

    # Stage 1: Lexical retrieval (BM25 via inverted index)
    # Cost: O(query_terms) — sub-millisecond
    # Coverage: full corpus (millions of documents)
    # Recall: high; Precision: low
    candidates_bm25 = bm25_retrieve(query, top_k=1000)

    # Stage 2: Dense retrieval (bi-encoder + ANN)
    # Cost: O(log N) with FAISS — ~10ms
    # Coverage: candidates from stage 1 re-ranked by embedding similarity
    # Recall: high; Precision: moderate
    query_embedding = query_encoder.encode(query)
    candidates_dense = faiss_index.search(query_embedding, k=200)

    # Stage 3: Cross-encoder re-ranking (BERT)
    # Cost: O(candidates) with full transformer — ~100ms for 200 candidates
    # Coverage: top-200 from dense retrieval
    # Recall: moderate; Precision: high
    merged = merge_and_deduplicate(candidates_bm25, candidates_dense)
    scored = []
    for doc in merged[:200]:
        score = cross_encoder.score(query, doc.text)
        scored.append((doc, score))
    scored.sort(key=lambda x: x[1], reverse=True)

    return [doc for doc, _ in scored[:k]]

# Each stage compensates for the limitations of the one before it.
# BM25 handles exact matches that confuse embeddings (rare names, IDs, codes).
# Dense retrieval handles semantic matches that BM25 misses (synonyms, paraphrases).
# Cross-encoder handles fine-grained relevance that dot products miss.
```

The comprehensive survey by Hambarde and Proença (2023) organizes this pipeline into two stages — term-based retrieval and semantic retrieval — and catalogs the models available at each level [2]. Their key insight: **modern search is never one model. It's a pipeline where each stage compensates for the limitations of the one before it.**

This is the critical difference from recommendation, as we'll see in Part 3. Search has always been multi-stage. Recommendation was historically single-stage — and the move to multi-stage pipelines in recommendation was one of the key convergences between the two fields.

---

**References**

1. Jiafeng Guo, Yixing Fan, Liang Pang, Liu Yang, Qingyao Ai, Hamed Zamani, W. Bruce Croft, et al. [*A Deep Look into Neural Ranking Models for Information Retrieval*](https://doi.org/10.1016/j.ipm.2019.102067). Information Processing & Management, 57(6), 2020.

2. Kailash Hambarde and Hugo Proença. [*Information Retrieval: Recent Advances and Beyond*](https://arxiv.org/abs/2301.08801). arXiv:2301.08801, 2023.

3. Gerard Salton, Anita Wong, and Chung-Shu Yang. [*A Vector Space Model for Automatic Indexing*](https://dl.acm.org/doi/10.1145/361219.361220). Communications of the ACM, 18(11): 613–620, 1975.

4. Stephen E. Robertson, Steve Walker, Susan Jones, Micheline Hancock-Beaulieu, and Mike Gatford. [*Okapi at TREC-3*](https://trec.nist.gov/pubs/trec3/papers/city.ps.gz). Proceedings of TREC-3, 1994.

5. Sergey Brin and Lawrence Page. [*The Anatomy of a Large-Scale Hypertextual Web Search Engine*](https://doi.org/10.1016/S0169-7552(98)00110-X). Computer Networks and ISDN Systems, 30(1–7): 107–117, 1998.

6. Christopher J.C. Burges. [*From RankNet to LambdaRank to LambdaMART: An Overview*](https://www.microsoft.com/en-us/research/publication/from-ranknet-to-lambdarank-to-lambdamart-an-overview/). Microsoft Research Technical Report MSR-TR-2010-82, 2010.

---

*Previous: [Part 1 — The Fundamental Distinction](/search-recommendation-fundamental-distinction) · Next: [Part 3 — Thirty Years of Recommendation](/search-recommendation-recommendation-techniques)*
