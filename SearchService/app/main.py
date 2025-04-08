from app.rabbitmq.consumer import listen_for_updates
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s"
)

if __name__ == "__main__":
    listen_for_updates()