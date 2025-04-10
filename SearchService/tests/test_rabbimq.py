import json
import pytest
from unittest.mock import MagicMock, patch
from app.rabbitmq.consumer import callback

@patch("app.rabbitmq.consumer.index.upsert")
@patch("app.rabbitmq.consumer.get_embedding")
def test_callback_success(mock_get_embedding, mock_upsert):
    mock_channel = MagicMock()
    mock_method = MagicMock(delivery_tag=123)
    mock_properties = MagicMock()
    
    mock_get_embedding.return_value = [0.1, 0.2, 0.3]

    message_dict = {
        "id": "prod-123",
        "name": "Milk",
        "description": "1L of fresh milk",
        "store": "Store A",
        "pricePerUnit": "2.50"
    }
    body = json.dumps(message_dict).encode("utf-8")

    callback(mock_channel, mock_method, mock_properties, body)

    mock_get_embedding.assert_called_once_with("Milk 1L of fresh milk", mock_get_embedding.call_args[0][1])
    mock_upsert.assert_called_once()
    mock_channel.basic_ack.assert_called_once_with(delivery_tag=123)
    
@patch("app.rabbitmq.consumer.index.upsert", side_effect=Exception("Upsert failed"))
@patch("app.rabbitmq.consumer.get_embedding")
def test_callback_failure(mock_get_embedding, mock_upsert):
    mock_channel = MagicMock()
    mock_method = MagicMock(delivery_tag=999)
    mock_properties = MagicMock()

    mock_get_embedding.return_value = [0.1, 0.2, 0.3]

    message_dict = {
        "id": "prod-999",
        "name": "Bread",
        "description": "Whole wheat bread",
        "store": "Store B",
        "pricePerUnit": "1.20"
    }
    body = json.dumps(message_dict).encode("utf-8")

    callback(mock_channel, mock_method, mock_properties, body)

    mock_channel.basic_nack.assert_called_once_with(delivery_tag=999, requeue=True)