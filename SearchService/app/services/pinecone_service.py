from app.core.config import Config
from app.helpers.pinecone_helpers import initialize_pinecone

class PineconeService:
    _index = None

    @classmethod
    def get_index(cls):
        if cls._index is None:
            _, cls._index = initialize_pinecone(
                api_key=Config.PINECONE_API_KEY,
                index_name=Config.PINECONE_INDEX_NAME,
                dimension=768
            )
        return cls._index
