from sentence_transformers import SentenceTransformer

class EmbeddingService:
    @staticmethod
    def load_model():
        # TODO: try rokn/slovlo-v1 model
        return SentenceTransformer("sentence-transformers/multi-qa-mpnet-base-cos-v1")