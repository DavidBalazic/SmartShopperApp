import uvicorn as uvicorn
from app.rabbitmq.consumer import listen_for_updates
import logging
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.routes import product_search
import threading
from app.services.pinecone_service import PineconeService
from app.services.embedding_service import EmbeddingService

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s"
)

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Add your frontend origin here
    allow_credentials=True, 
    allow_methods=["*"],  # Allow all HTTP methods
    allow_headers=["*"],  # Allow all headers
)

app.include_router(product_search.router)

@app.on_event("startup")
def startup_event():
    logging.info("Loading SentenceTransformer model...")
    model = EmbeddingService.get_model()
    logging.info("Model loaded.")

    logging.info("Initializing Pinecone index...")
    index = PineconeService.get_index()
    logging.info("Pinecone index initialized.")

    def run_consumer():
        try:
            listen_for_updates(model, index)
        except Exception as e:
            logging.error(f"RabbitMQ consumer crashed: {e}")
    
    threading.Thread(target=run_consumer, daemon=True).start()

if __name__ == "__main__":
    uvicorn.run(app, host="localhost", port=8000)