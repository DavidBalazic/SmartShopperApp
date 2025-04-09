from pydantic import BaseModel
from typing import Optional

class Product(BaseModel):
    id: str
    score: float
    store: Optional[str]
    pricePerUnit: Optional[float]