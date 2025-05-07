import pika
import json
import logging
from app.helpers.pinecone_helpers import get_document_embedding
from app.core.config import Config

def callback(ch, method, properties, body, model, index):
    try:
        decoded_body = body.decode('utf-8')
        print(f"message received: {decoded_body}")
        messages = json.loads(decoded_body)
        
        if isinstance(messages, dict):
            messages = [messages]
        
        for message in messages:
            product_id = message.get("id", "")
            name = message.get("name", "")
            store = message.get("store", "")
            pricePerUnit = message.get("pricePerUnit", "")

            embedding = get_document_embedding(name, model)

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
        ch.basic_nack(delivery_tag=method.delivery_tag, requeue=False)


def listen_for_updates(model, index):
    credentials = pika.PlainCredentials(Config.RABBITMQ_USER, Config.RABBITMQ_PASSWORD)
    parameters = pika.ConnectionParameters(
        host=Config.RABBITMQ_HOST,
        port=Config.RABBITMQ_PORT,
        virtual_host='/',
        credentials=credentials
    )

    connection = pika.BlockingConnection(parameters)
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