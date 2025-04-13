from app.services.embedding_service import EmbeddingService
from app.services.pinecone_service import PineconeService
from app.helpers.pinecone_helpers import query_from_pinecone

def query_products(q: str, namespace: str = "products", top_k: int = 10):
    model = EmbeddingService.get_model()
    index = PineconeService.get_index()
    
    response = query_from_pinecone(
        query=q,
        index=index,
        model=model,
        namespace=namespace,
        top_k=top_k,
        include_metadata=True
    )
    
    return response.get("matches", [])