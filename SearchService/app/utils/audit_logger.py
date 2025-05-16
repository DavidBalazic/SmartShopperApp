from confluent_kafka import Producer
import json
import os
from datetime import datetime, timezone
from app.core.config import Config

KAFKA_BROKER = Config.KAFKA_BROKER
KAFKA_TOPIC = Config.KAFKA_TOPIC

producer = Producer({'bootstrap.servers': KAFKA_BROKER})

def send_audit_log(actor_id, action, resource, service, ip, user_agent, details):
    log_entry = {
        "timestamp": datetime.now(timezone.utc).isoformat(),
        "actor": {
            "id": actor_id,
            "ip": ip,
            "userAgent": user_agent
        },
        "action": action,
        "resource": resource,
        "service": service,
        "details": details
    }

    producer.produce(KAFKA_TOPIC, key=actor_id, value=json.dumps(log_entry))
    producer.flush()