import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    GRPC_SERVER_HOST = os.getenv('GRPC_SERVER_HOST')