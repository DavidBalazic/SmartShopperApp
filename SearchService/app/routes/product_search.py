from fastapi import APIRouter, Query
from typing import Optional
from app.models.product import Product
from app.services.product_search_service import query_products
import logging
from app.services.product_service import get_product_by_id

router = APIRouter()

@router.get("/cheapest-product/", response_model=Optional[Product])
def get_cheapest_product(
    q: str = Query(..., description="Search query text"),
    store: Optional[str] = Query(None, description="Filter results by store name")
    ):
    logging.info(f"Received query: {q}, store: {store}")
    matches = query_products(query=q, store=store)
    
    cheapest = min(
        matches,
        key=lambda x: float(x.metadata.get("pricePerUnit", float("inf")))
    )
    logging.info(f"Cheapest product found: {cheapest}")
    
    product_details = get_product_by_id(cheapest.id)
    logging.info(f"Product details fetched: {product_details}")

    return Product(
        name=product_details.name,
        description=product_details.description,
        price=product_details.price,
        quantity=product_details.quantity,
        unit=product_details.unit,
        store=product_details.store,
        pricePerUnit=product_details.pricePerUnit
    )