from typing import Optional
from sentence_transformers import SentenceTransformer, CrossEncoder
from pinecone import Index

class Container:
    model: Optional[SentenceTransformer] = None
    index: Optional[Index] = None
    reranker: Optional[CrossEncoder] = None

container = Container()