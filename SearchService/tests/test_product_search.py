from fastapi.testclient import TestClient
from unittest.mock import patch, MagicMock, ANY
from app.main import app 
from app.dtos.product import Product
from app.dependencies.deps import get_model, get_index

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

def fake_get_model():
    return MagicMock(name="FakeModel")

def fake_get_index():
    return MagicMock(name="FakeIndex")

@patch("app.routes.product_search.send_audit_log")
@patch("app.routes.product_search.get_product_by_id")
@patch("app.routes.product_search.query_products")
def test_get_cheapest_product(mock_query_from_pinecone, mock_get_product_by_id, mock_send_audit_log):
    mock_query_from_pinecone.return_value = mock_pinecone_results
    mock_get_product_by_id.return_value = mock_product_service_result
    mock_send_audit_log.return_value = None
    app.dependency_overrides[get_model] = fake_get_model
    app.dependency_overrides[get_index] = fake_get_index
    
    response = client.get("/cheapest-product/", params={"q": "milk"})
    assert response.status_code == 200

    data = response.json()
    assert data["name"] == "Test Product"
    assert data["pricePerUnit"] == 2.5
    
    app.dependency_overrides = {}

@patch("app.routes.product_search.send_audit_log")
@patch("app.routes.product_search.get_product_by_id")
@patch("app.routes.product_search.query_products")
def test_get_cheapest_product_with_store_filter(mock_query_from_pinecone, mock_get_product_by_id, mock_send_audit_log):
    mock_query_from_pinecone.return_value = mock_pinecone_results
    mock_get_product_by_id.return_value = mock_product_service_result
    mock_send_audit_log.return_value = None
    app.dependency_overrides[get_model] = fake_get_model
    app.dependency_overrides[get_index] = fake_get_index

    response = client.get("/cheapest-product/", params={"q": "milk", "store": "store a"})
    assert response.status_code == 200

    data = response.json()
    assert data["name"] == "Test Product"
    assert data["store"] == "Store A"
    assert data["pricePerUnit"] == 2.5

    mock_query_from_pinecone.assert_called_once_with(
        query="milk",
        store="store a",
        model=ANY,
        index=ANY
    )
    
    app.dependency_overrides = {}