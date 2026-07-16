---
title: "S&R: How LLMs Reshape Both Fields — And Where They Don't"
date: 2026-07-16
slug: search-recommendation-llms
summary: Part 6 of 7. LLMs transform search and recommendation in different ways — RAG, generative retrieval, conversational recommendation, agentic pipelines. But the distinction survives in objective alignment, evaluation, serendipity tolerance, and infrastructure.
tags: [search, recommendation, llm, rag, generative-retrieval, agentic, survey, series]
series: search-recommendation
---

*S&R stands for Search & Recommendation. Part 6 of a seven-part series. How LLMs reshape both fields — differently — and why the fundamental distinction survives the generative era.*

The arrival of LLMs is the most significant development in both information retrieval and recommendation since BERT. But the way LLMs affect each field is different, and understanding the difference is essential for engineering.

Two recent surveys capture the scope. Zhu et al. (2024) organize LLM-IR integration into four roles: query rewriter, retriever, reranker, and reader [1]. Li et al. (2024) provide the cross-cutting view — framing both search and recommendation as instances of generative retrieval and identifying where they converge and where they diverge [2].

The key insight: **LLMs don't eliminate the distinction. They reveal it at a higher level of abstraction.** Both fields are moving toward generative paradigms, but the *thing being generated* is different. Search generates answers from documents. Recommendation generates item predictions from user histories.

## How LLMs Transform Search

LLMs operate at four levels in the search pipeline:

```python
# Level 1: Query Understanding — LLMs disambiguate and expand queries
def llm_query_understanding(raw_query: str, llm) -> dict:
    """Transform natural language into structured search intent."""
    prompt = f"""
    Parse this search query into structured intent:
    Query: "{raw_query}"

    Return JSON with: intent_type, entities, constraints, expansions.
    """
    response = llm.generate(prompt)
    return json.loads(response)

# "Show me movies like Inception but funnier"
# → {intent_type: "similar_items_with_constraint",
#    entities: ["Inception"],
#    constraints: ["genre: comedy", "tone: lighter"],
#    expansions: ["mind-bending heist films", "comedic thrillers"]}

# Level 2: Document Understanding — LLMs generate richer representations
def llm_enrich_document(doc: str, llm) -> dict:
    """Generate structured metadata that transcends keyword matching."""
    prompt = f"""
    For this document, generate:
    1. A 2-sentence summary
    2. 5-10 key phrases
    3. Entity tags (people, places, concepts)
    4. A semantic embedding-friendly description

    Document: {doc[:4000]}
    """
    return llm.generate_structured(prompt)

# Level 3: RAG — Retrieve then generate a synthesized answer
def rag_search(query: str, retriever, llm, k: int = 5) -> str:
    """The dominant LLM-search paradigm in 2025."""
    # Retrieve relevant documents
    docs = retriever.retrieve(query, k=k)

    # Generate answer grounded in retrieved documents
    context = "\n\n".join(f"[{i+1}] {doc.text}" for i, doc in enumerate(docs))
    prompt = f"""
    Answer the query using ONLY the provided documents.
    Cite sources by number.

    Query: {query}

    Documents:
    {context}
    """
    return llm.generate(prompt)

# Level 4: Generative Retrieval — the model IS the index
# The most radical approach: the LLM directly generates document IDs
# without an explicit retrieval index. Still experimental.
```

The core search user need — "I have a question, find me the answer" — aligns naturally with LLM capabilities. The LLM augments retrieval at every stage without replacing it.

## How LLMs Transform Recommendation

For recommendation, the transformation is more structural:

```python
# Level 1: Feature Engineering Automation
# LLMs generate rich profiles from sparse structured data.
# DoorDash's content embeddings and Consumer Memory Blocks are paradigmatic:
# the LLM produces the fuel, retrieval and ranking remain purpose-built.

def llm_generate_item_profile(item: dict, llm) -> str:
    """Produce a rich narrative that a standard encoder can embed."""
    return llm.generate(f"Describe this item for a recommendation system: {item}")

# Level 2: Generative Recommendation
# Instead of scoring candidates, the LLM autoregressively generates item tokens.
# Meta's HSTU, Kuaishou's OneRec, Google's TIGER are examples.

class GenerativeRecommender:
    """Recommendation as sequence-to-sequence: predict next-item tokens."""
    def recommend(self, user_history: list[int], llm, k: int = 10) -> list[int]:
        prompt = f"User interaction history (item IDs): {user_history}\nPredict next items:"
        # LLM autoregressively generates item IDs
        return llm.generate_tokens(prompt, max_tokens=k)

# Level 3: Conversational Recommendation
# Multi-turn dialogue where the system elicits preferences and adapts.

def conversational_recommendation_loop(user_id: str, llm, catalog):
    """Search and recommendation blur into conversation."""
    context = {"history": get_user_history(user_id)}
    for turn in range(5):
        user_input = get_user_response()
        if is_search_like(user_input):    # "I want a sci-fi movie"
            results = catalog.search(user_input)
            context["mode"] = "search"
        else:                              # "something cerebral, not too long"
            results = catalog.recommend(context)
            context["mode"] = "recommendation"
        llm_response = llm.generate_response(results, context)
        show_to_user(llm_response)

# Level 4: Agentic Recommendation
# LLM-powered agents that plan, use tools, maintain memory, reason.
# ARAG (SIGIR 2025): 4 specialized agents — User Understanding, NLI,
# Context Summary, Item Ranker — achieve +42.1% NDCG@5 over vanilla RAG.
```

## Why the Distinction Survives

Despite this convergence, the distinction remains essential for four reasons:

**1. Objective alignment.** A search system optimizing for engagement shows addictive content instead of answering the query. A recommendation system optimizing for precision produces an echo chamber.

```python
# Wrong: one objective for both
def bad_unified_objective(model_output, labels):
    return cross_entropy(model_output, labels)  # what are we even optimizing?

# Right: task-aware objectives
def search_objective(predictions, relevance_labels):
    return ndcg_loss(predictions, relevance_labels)  # did we retrieve the right thing?

def recommendation_objective(predictions, engagement_labels):
    return binary_cross_entropy(predictions, engagement_labels)  # will user engage?
    # Plus: diversity bonus, freshness decay, exploration budget, long-term retention proxy
```

**2. Evaluation.** Search: NDCG, MRR, precision@K against relevance judgments. Recommendation: retention, discovery, session length, long-term satisfaction. These are correlated but not identical.

**3. The serendipity gradient.** In search, serendipity is a bug. In recommendation, it's a feature.

```python
def demonstrate_serendipity_gradient():
    """The same behavior is a bug in one context and a feature in the other."""

    # SEARCH: User types "The Godfather"
    # System returns "Goodfellas" because "people who watch The Godfather
    # also love Goodfellas"
    # → BROKEN. The user wanted The Godfather. Return The Godfather.

    # RECOMMENDATION: User browses homepage on Friday night
    # System shows "The Godfather" — which the user has already seen 4 times
    # → BROKEN. The user wants something new. Show Goodfellas.

    # Both failures happen when teams optimize for the wrong thing.
    # The first optimized for engagement instead of relevance.
    # The second optimized for precision instead of discovery.
```

**4. Infrastructure.** Search: inverted indices, real-time query parsing, sub-100ms latency at high QPS. Recommendation: user profile stores, embedding indices, offline batch pipelines amortized across daily refresh cycles.

```python
# Search infrastructure
class SearchServing:
    latency_sla: float = 0.100  # seconds — user is waiting
    qps: int = 10_000           # every keystroke is a query
    freshness: str = "real-time"  # new documents must be searchable immediately
    index_type: str = "inverted"  # keyword → posting list

# Recommendation infrastructure
class RecServing:
    latency_sla: float = 0.200  # seconds — page load, more forgiving
    qps: int = 1_000            # one request per page view
    freshness: str = "daily"      # embeddings recomputed nightly
    index_type: str = "ANN"      # vector → nearest neighbors
```

## The Survey Papers Agree

The Li et al. (2024) survey on generative search and recommendation puts it well: the fields share a common mathematical framework — matching entities across representation spaces — but differ in the *nature of the mismatch* [2]. Search deals with query–document mismatch: different words for the same concept. Recommendation deals with user–item mismatch: users and items live in completely different semantic spaces. The first is a lexical/semantic gap. The second is a modality gap. Solving them requires different tools, even when those tools share a common transformer backbone.

The Gupta et al. (2025) survey on generative recommendation reinforces this: industrial systems like TIGER, LIGER, OneRec, and HSTU all treat recommendation as its own generative problem with its own tokenization, its own evaluation, and its own cold-start challenges [3]. None of them use the same architecture unmodified for search.

LLMs are infrastructure. They don't eliminate the distinction between search and recommendation — they make it more important to get right.

---

**References**

1. Yutao Zhu, Huaying Yuan, Shuting Wang, et al. [*Large Language Models for Information Retrieval: A Survey*](https://arxiv.org/abs/2308.07107). ACM Transactions on Information Systems, 2024.

2. Yongqi Li et al. [*A Survey of Generative Search and Recommendation in the Era of Large Language Models*](https://arxiv.org/abs/2404.16924). arXiv:2404.16924, 2024.

3. Shashank Gupta et al. [*Generative Recommendation: A Survey of Models, Systems, and Industrial Advances*](https://www.techrxiv.org/doi/full/10.36227/techrxiv.176523089.94266134/v2). TechRxiv, 2025.

4. Chi Zhang et al. [*ARAG: Agentic Retrieval Augmented Generation for Personalized Recommendation*](https://arxiv.org/abs/2506.21931). SIGIR 2025.

5. Zhuang Liu et al. [*A Comprehensive Survey on LLM-Powered Recommender Systems*](https://ieeexplore.ieee.org/abstract/document/11129085). IEEE Access, 2024.

---

*Previous: [Part 5 — How Spotify, DoorDash, Airbnb, and Pinterest Navigate the Boundary](/search-recommendation-industry-boundary) · Next: [Part 7 — Practical Architecture: Building Systems That Get It Right](/search-recommendation-practical-architecture)*
