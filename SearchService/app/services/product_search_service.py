from app.services.embedding_service import EmbeddingService
from app.services.pinecone_service import PineconeService
from app.helpers.pinecone_helpers import query_from_pinecone
from typing import Optional
import logging

def query_products(
    query: str,
    model,
    index,
    reranker,
    store: Optional[str] = None,
    namespace: str = "products",
    top_k: int = 30
    ):
    # Get top_k from Pinecone
    results = query_from_pinecone(
        query=query,
        index=index,
        model=model,
        namespace=namespace,
        top_k=top_k,
        include_metadata=True
    )
    
    logging.info(f"Query: {query} results: {results}")
    
    # Filter by Pinecone score
    results = [match for match in results if match.score > 0.7555]

    logging.info(f"Filtered by Pinecone score > 0.755: {results}")
    
    # Return empty list if no results found
    if not results:
        return []
    
    # Prepare (query, name) pairs for reranking
    rerank_pairs = [
        (query, match.metadata.get("name", ""))
        for match in results
    ]
    
    # Predict reranker scores
    reranker_scores = reranker.predict(rerank_pairs)

    logging.info(f"Reranker scores: {reranker_scores}")

    # Combine and filter by reranker > 0
    reranked = [
        (match, score)
        for match, score in zip(results, reranker_scores)
        if score > 0
    ]

    logging.info(f"Reranker positive matches: {reranked}")

    # If no matches were kept by reranker, return empty list
    if not reranked:
        return []
    
    # Sort by reranker score descending
    reranked.sort(key=lambda x: x[1], reverse=True)

    # Apply store filter if needed
    filtered = [
        match
        for match, score in reranked
        if store is None or match.metadata.get("store", "").lower() == store.lower()
    ]

    logging.info(f"After store filter matches: {filtered}")
    
    return filtered