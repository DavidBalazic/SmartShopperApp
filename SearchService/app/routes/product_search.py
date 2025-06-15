from fastapi import APIRouter, Query, Depends, Request
from typing import Optional
from app.dtos.product import Product
from app.services.product_search_service import query_products
import logging
from app.services.product_service import get_product_by_id
from app.dependencies.deps import get_model, get_index, get_reranker
from app.utils.audit_logger import send_audit_log

router = APIRouter()

@router.get("/cheapest-product", response_model=Optional[Product])
def get_cheapest_product(
    request: Request,
    q: str = Query(..., description="Search query text"),
    store: Optional[str] = Query(None, description="Filter results by store name"),
    model=Depends(get_model),
    index=Depends(get_index),
    reranker=Depends(get_reranker)
    ):
    # Extracting request metadata
    user_agent = request.headers.get("user-agent")
    ip = request.client.host
    
    logging.info(f"Received query: {q}, store: {store}")
    matches = query_products(query=q, store=store, model=model, index=index, reranker=reranker)
    
    # If no matches found, log and return None
    if not matches:
        logging.warning(f"No products found for query: {q}, store: {store}")
        
        send_audit_log(
            actor_id="anonymous",  
            action="cheapest-product",
            resource="product",
            service="SearchService",
            ip=ip,
            user_agent=user_agent,
            details={
                "query": q,
                "store": store,
                "productDetail": None,
                "message": "No products found"
            }
        )
        
        return None
    
    # Find the cheapest product
    cheapest = min(
        matches,
        key=lambda x: float(x.metadata.get("pricePerUnit", float("inf")))
    )
    logging.info(f"Cheapest product found: {cheapest}")
    
    product_details = get_product_by_id(cheapest.id)
    logging.info(f"Product details fetched: {product_details}")
    
    product = Product(
        name=product_details.name,
        description=product_details.description,
        price=product_details.price,
        quantity=product_details.quantity,
        unit=product_details.unit,
        store=product_details.store,
        pricePerUnit=product_details.pricePerUnit,
        imageUrl=product_details.imageUrl
    )

    # Log the audit event
    send_audit_log(
        actor_id="anonymous",  # TODO: Extract from JWT
        action="cheapest-product",
        resource="product",
        service="SearchService",
        ip=ip,
        user_agent=user_agent,

        details={
            "query": q,
            "store": store,
            "productDetail": product.model_dump_json(),
        }
    )

    return product
    
@router.get("/search-products", response_model=list[Product])
def get_all_matching_products(
    request: Request,
    q: str = Query(..., description="Search query text"),
    store: Optional[str] = Query(None, description="Optional store filter"),
    model=Depends(get_model),
    index=Depends(get_index),
    reranker=Depends(get_reranker)
):
    # Extracting request metadata
    user_agent = request.headers.get("user-agent")
    ip = request.client.host
    
    logging.info(f"Received query: {q}, store: {store}")
    matches = query_products(query=q, store=store, model=model, index=index, reranker=reranker)

    # TODO: replace with get_products_by_id
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
                pricePerUnit=product_details.pricePerUnit,
                imageUrl=product_details.imageUrl
            )
            products.append(product)
        except Exception as e:
            logging.warning(f"Could not fetch product with ID {match.id}: {e}")
            
    # Log the audit event
    send_audit_log(
        actor_id="anonymous",  # TODO: Extract from JWT
        action="search-products",
        resource="product",
        service="SearchService",
        ip=ip,
        user_agent=user_agent,
        details={
            "query": q,
            "store": store,
            "matches": len(matches),
            "productDetail": [p.model_dump_json() for p in products],
        }
    )
    
    return products