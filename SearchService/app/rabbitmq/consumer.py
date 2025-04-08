import pika
import json
import logging
from sentence_transformers import SentenceTransformer
from app.helpers.pinecone_helpers import initialize_pinecone, get_embedding
from app.core.config import Config


pc, index = initialize_pinecone(api_key=Config.PINECONE_API_KEY, index_name=Config.PINECONE_INDEX_NAME, dimension=384)

model = SentenceTransformer("sentence-transformers/all-MiniLM-L6-v2")
print("Model loaded successfully.")

def callback(ch, method, properties, body):
    try:
        decoded_body = body.decode('utf-8')
        print(f"message received: {decoded_body}")
        message = json.loads(decoded_body)
        product_id = message["id"]
        name = message.get("name", "")
        description = message.get("description", "")
        store = message.get("store", "")
        pricePerUnit = message.get("pricePerUnit", "")


        combined_text = f"{name} {description}"
        embedding = get_embedding(combined_text, model)

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
    connection = pika.BlockingConnection(
        pika.URLParameters(Config.RABBITMQ_HOST)
    )
    channel = connection.channel()
    channel.queue_declare(queue=Config.RABBITMQ_QUEUE, durable=True)
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(
        queue=Config.RABBITMQ_QUEUE,
        on_message_callback=callback,
        auto_ack=False
    )
    logging.info("Waiting for messages. Press CTRL+C to exit.")
    channel.start_consuming()