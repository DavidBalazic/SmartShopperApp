from fastapi.testclient import TestClient
from unittest.mock import patch, MagicMock
from app.main import app 
from app.models.product import Product

client = TestClient(app)

mock_pinecone_results = [
    MagicMock(
        id="prod-1",
        score=0.8,
        metadata={"pricePerUnit": "2.5", "store": "Store A"}
    ),
    MagicMock(
        id="prod-2",
        score=0.6,
        metadata={"pricePerUnit": "3.1", "store": "Store B"}
    )
]

mock_product_service_result = Product(
    name="Test Product",
    description="A cheap product",
    price=2.5,
    quantity=1,
    unit="kg",
    store="Store A",
    pricePerUnit=2.5
)

@patch("app.routes.product.get_product_by_id")
@patch("app.routes.product.query_from_pinecone")
def test_get_cheapest_product(mock_query_from_pinecone, mock_get_product_by_id):
    mock_query_from_pinecone.return_value = mock_pinecone_results
    mock_get_product_by_id.return_value = mock_product_service_result
    
    response = client.get("/cheapest-product/", params={"q": "milk"})
    assert response.status_code == 200

    data = response.json()
    assert data["name"] == "Test Product"
    assert data["pricePerUnit"] == 2.5