from pydantic import BaseModel
from typing import Optional

class Product(BaseModel):
    name: str
    description: Optional[str]
    price: float
    quantity: float
    unit: str
    store: str
    pricePerUnit: Optional[float]
    imageUrl: str