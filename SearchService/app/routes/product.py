from fastapi import APIRouter, Query
from typing import Optional
from app.helpers.pinecone_helpers import query_from_pinecone
from app.models.product import Product
from app.dependencies.search_service import model, index
import logging

router = APIRouter()

@router.get("/cheapest-product/", response_model=Optional[Product])
def get_cheapest_product(q: str = Query(..., description="Search query text")):
    logging.info(f"Received query: {q}")
    results = query_from_pinecone(
        query=q,
        index=index,
        model=model,
        namespace="products",
        top_k=10,
        include_metadata=True
    )
    logging.info(f"Query results: {results}")
    filtered = [
        match for match in results if match.score >= 0.3
    ]
    logging.info(f"Filtered results: {filtered}")

    if not filtered:
        return None

    cheapest = min(
        filtered,
        key=lambda x: float(x.metadata.get("pricePerUnit", float("inf")))
    )
    logging.info(f"Cheapest product found: {cheapest}")

    return Product(
        id=cheapest.id,
        score=cheapest.score,
        store=cheapest.metadata.get("store"),
        pricePerUnit=cheapest.metadata.get("pricePerUnit")
    )