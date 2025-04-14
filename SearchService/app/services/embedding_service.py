from sentence_transformers import SentenceTransformer

class EmbeddingService:
    _model = None

    @classmethod
    def get_model(cls):
        if cls._model is None:
            # TODO: try rokn/slovlo-v1 model
            cls._model = SentenceTransformer("sentence-transformers/multi-qa-mpnet-base-cos-v1")
        return cls._model