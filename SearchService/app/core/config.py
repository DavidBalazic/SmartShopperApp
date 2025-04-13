import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    RABBITMQ_HOST = os.getenv('RABBITMQ_HOST')
    RABBITMQ_QUEUE = os.getenv('RABBITMQ_QUEUE')
    PINECONE_API_KEY = os.getenv('PINECONE_API_KEY')
    PINECONE_INDEX_NAME = os.getenv('PINECONE_INDEX_NAME')
    GRPC_SERVER_HOST = os.getenv('GRPC_SERVER_HOST')