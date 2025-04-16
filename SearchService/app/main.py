import uvicorn as uvicorn
from app.rabbitmq.consumer import listen_for_updates
import logging
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.routes import product_search
import threading
from app.services.pinecone_service import PineconeService
from app.services.embedding_service import EmbeddingService
from contextlib import asynccontextmanager
from app.dependencies.container import container

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s"
)

@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Loading SentenceTransformer model...")
    container.model = EmbeddingService.load_model()
    logging.info("Model loaded.")

    logging.info("Initializing Pinecone index...")
    container.index = PineconeService.initialize_index()
    logging.info("Pinecone index initialized.")

    def run_consumer():
        try:
            listen_for_updates(container.model, container.index)
        except Exception as e:
            logging.error(f"RabbitMQ consumer crashed: {e}")
    
    threading.Thread(target=run_consumer, daemon=True).start()

    yield

    logging.info("Shutting down the application.")

app = FastAPI(lifespan=lifespan)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allow all origins, modify for specific domains
    allow_credentials=True, 
    allow_methods=["*"],  # Allow all HTTP methods
    allow_headers=["*"],  # Allow all headers
)

app.include_router(product_search.router)

if __name__ == "__main__":
    uvicorn.run(app, host="localhost", port=8000)