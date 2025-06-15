import logging
from pinecone import Pinecone, ServerlessSpec

def initialize_pinecone(api_key, index_name, dimension):
    pc = Pinecone(api_key=api_key)
    if index_name not in pc.list_indexes().names():
        logging.info(f"Creating index {index_name}")
        pc.create_index(
            name=index_name,
            dimension=dimension,
            metric="cosine",
            spec=ServerlessSpec(cloud="aws", region="us-east-1"),
        )
    index = pc.Index(name=index_name)
    return pc, index

def get_document_embeddings(texts, model):
    return model.encode(texts, convert_to_tensor=False).tolist()

def get_document_embedding(text, model):
    # Convert text to lowercase for consistency
    text = text.lower()
    return get_document_embeddings([text], model)[0]

def get_query_embeddings(texts, model):
    return model.encode(texts, convert_to_tensor=False).tolist()

def get_query_embedding(text, model):
    # Convert text to lowercase for consistency
    text = text.lower()
    return get_query_embeddings([text], model)[0]

def query_from_pinecone(query, index, model, namespace, top_k=3, include_metadata=True):
    # get embedding from THE SAME embedder as the documents
    query_embedding = get_query_embedding(query, model)

    return index.query(
        vector=query_embedding,
        top_k=top_k,
        namespace=namespace,
        include_metadata=include_metadata, 
    ).get("matches")