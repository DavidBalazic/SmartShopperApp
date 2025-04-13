import pika
import json
import logging
from app.services.pinecone_service import PineconeService
from app.services.embedding_service import EmbeddingService
from app.helpers.pinecone_helpers import get_embedding
from app.core.config import Config

def callback(ch, method, properties, body, model, index):
    try:
        decoded_body = body.decode('utf-8')
        print(f"message received: {decoded_body}")
        message = json.loads(decoded_body)
        product_id = message["id"]
        name = message.get("name", "")
        store = message.get("store", "")
        pricePerUnit = message.get("pricePerUnit", "")

        embedding = get_embedding(name, model)

        index.upsert(
            vectors=[
                {
                    "id": product_id,
                    "values": embedding,
                    "metadata": {
                        "store": store,
                        "pricePerUnit": pricePerUnit
                    }
                }
            ],
            namespace="products"
        )
        logging.info(f"Upserted product {product_id}, name: {name}, store: {store}, price per unit: {pricePerUnit} to Pinecone.")
        ch.basic_ack(delivery_tag=method.delivery_tag)

    except Exception as e:
        logging.error(f"Error processing message: {e}")
        ch.basic_nack(delivery_tag=method.delivery_tag, requeue=True)


def listen_for_updates():
    model = EmbeddingService.get_model()
    index = PineconeService.get_index()
    
    connection = pika.BlockingConnection(
        pika.ConnectionParameters(Config.RABBITMQ_HOST)
    )
    channel = connection.channel()
    channel.queue_declare(queue=Config.RABBITMQ_QUEUE, durable=True)
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(
        queue=Config.RABBITMQ_QUEUE,
        on_message_callback=lambda ch, method, properties, body: callback(ch, method, properties, body, model, index),
        auto_ack=False
    )
    logging.info("Waiting for messages. Press CTRL+C to exit.")
    channel.start_consuming()