import json
from unittest.mock import MagicMock, patch
from app.rabbitmq.consumer import callback
from app.services.pinecone_service import PineconeService
from app.services.embedding_service import EmbeddingService

@patch.object(EmbeddingService, "get_model", return_value=MagicMock())
@patch.object(PineconeService, "get_index", return_value=MagicMock())
@patch("app.rabbitmq.consumer.get_embedding")
def test_callback_success(mock_get_embedding, mock_get_index, mock_get_model):
    mock_channel = MagicMock()
    mock_method = MagicMock(delivery_tag=123)
    mock_properties = MagicMock()
    
    mock_get_embedding.return_value = [0.1, 0.2, 0.3]

    message_dict = {
        "id": "prod-123",
        "name": "Milk",
        "store": "Store A",
        "pricePerUnit": "2.50"
    }
    body = json.dumps(message_dict).encode("utf-8")

    model = MagicMock()
    index = MagicMock()

    callback(mock_channel, mock_method, mock_properties, body, model, index)

    mock_get_embedding.assert_called_once_with("Milk", model)

    index.upsert.assert_called_once()

    mock_channel.basic_ack.assert_called_once_with(delivery_tag=123)



@patch.object(EmbeddingService, "get_model", return_value=MagicMock())
@patch.object(PineconeService, "get_index", return_value=MagicMock())
@patch("app.rabbitmq.consumer.get_embedding")
def test_callback_failure(mock_get_embedding, mock_get_index, mock_get_model):
    mock_get_embedding.return_value = [0.1, 0.2, 0.3]
    
    mock_index = MagicMock()
    mock_index.upsert.side_effect = Exception("Upsert failed")
    
    message_dict = {
        "id": "prod-999",
        "name": "Bread",
        "store": "Store B",
        "pricePerUnit": "1.20"
    }
    body = json.dumps(message_dict).encode("utf-8")
    
    mock_channel = MagicMock()
    mock_method = MagicMock(delivery_tag=999)
    mock_properties = MagicMock()

    callback(mock_channel, mock_method, mock_properties, body, mock_get_model(), mock_index)
    
    mock_channel.basic_nack.assert_called_once_with(delivery_tag=999, requeue=True)