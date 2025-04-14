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
    
    # TODO: handle empty
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
    
@router.get("/search-products/", response_model=list[Product])
def get_all_matching_products(
    q: str = Query(..., description="Search query text"),
    store: Optional[str] = Query(None, description="Optional store filter")
):
    logging.info(f"Received query: {q}, store: {store}")
    matches = query_products(query=q, store=store)

    products = []
    for match in matches:
        try:
            product_details = get_product_by_id(match.id)
            logging.info(f"Product details fetched from search-products method: {product_details}")
            product = Product(
                name=product_details.name,
                description=product_details.description,
                price=product_details.price,
                quantity=product_details.quantity,
                unit=product_details.unit,
                store=product_details.store,
                pricePerUnit=product_details.pricePerUnit
            )
            products.append(product)
        except Exception as e:
            logging.warning(f"Could not fetch product with ID {match.id}: {e}")
    
    return products