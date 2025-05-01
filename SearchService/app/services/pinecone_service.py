from app.core.config import Config
from app.helpers.pinecone_helpers import initialize_pinecone

class PineconeService:
    @staticmethod
    def initialize_index():
        _, index = initialize_pinecone(
            api_key=Config.PINECONE_API_KEY,
            index_name=Config.PINECONE_INDEX_NAME,
            dimension=768
        )
        return index