from sentence_transformers import SentenceTransformer, CrossEncoder

class EmbeddingService:
    @staticmethod
    def load_model():
        return SentenceTransformer("rokn/slovlo-v1")
    @staticmethod
    def load_reranker():
        return CrossEncoder("cross-encoder/ms-marco-MiniLM-L-6-v2")