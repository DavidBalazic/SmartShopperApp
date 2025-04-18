from app.services.embedding_service import EmbeddingService
from app.services.pinecone_service import PineconeService
from app.helpers.pinecone_helpers import query_from_pinecone
from typing import Optional
import logging

def query_products(
    query: str,
    model,
    index,
    store: Optional[str] = None,
    namespace: str = "products",
    top_k: int = 10

):
    results = query_from_pinecone(
        query=query,
        index=index,
        model=model,
        namespace=namespace,
        top_k=top_k,
        include_metadata=True
    )
    
    logging.info(f"Query: {query} results: {results}")
    
    return [
        match for match in results
        if match.score >= 0.3 and (store is None or match.metadata.get("store").lower() == store.lower())
    ]