import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    RABBITMQ_HOST = os.getenv('RABBITMQ_HOST', 'localhost')
    RABBITMQ_QUEUE = os.getenv('RABBITMQ_QUEUE', 'product-queue')
    PINECONE_API_KEY = os.getenv('PINECONE_API_KEY')
    PINECONE_INDEX_NAME = os.getenv('PINECONE_INDEX_NAME')
    GRPC_SERVER_HOST = os.getenv('GRPC_SERVER_HOST')