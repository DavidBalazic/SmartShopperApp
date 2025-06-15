from app.dependencies.container import container

def get_model():
    if container.model is None:
        raise RuntimeError("Model not initialized")
    return container.model

def get_index():
    if container.index is None:
        raise RuntimeError("Pinecone index not initialized")
    return container.index

def get_reranker():
    if container.reranker is None:
        raise RuntimeError("Reranker not initialized")
    return container.reranker