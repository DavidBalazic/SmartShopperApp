from sentence_transformers import SentenceTransformer
from app.helpers.pinecone_helpers import initialize_pinecone
from app.core.config import Config

model = SentenceTransformer("sentence-transformers/multi-qa-mpnet-base-cos-v1")
pc, index = initialize_pinecone(
    api_key=Config.PINECONE_API_KEY,
    index_name=Config.PINECONE_INDEX_NAME,
    dimension=768
)