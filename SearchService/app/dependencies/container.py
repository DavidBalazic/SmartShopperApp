from typing import Optional
from sentence_transformers import SentenceTransformer
from pinecone import Index

class Container:
    model: Optional[SentenceTransformer] = None
    index: Optional[Index] = None

container = Container()