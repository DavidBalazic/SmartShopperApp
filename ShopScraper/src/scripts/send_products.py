import pandas as pd
import grpc
import argparse
from typing import List
from src.grpc_clients import product_pb2, product_pb2_grpc
from core.config import Config

BATCH_SIZE = 100 

def read_products_from_csv(store_name: str) -> List[product_pb2.AddProductRequest]:
    file_path = f"data/preprocessed/{store_name}/{store_name}_data.csv"
    df = pd.read_csv(file_path)
    products = []

    for _, row in df.iterrows():
        try:
            product = product_pb2.AddProductRequest(
                name=str(row["name"]),
                #description="",
                price=float(row["price"]),
                #quantity=0,
                unit=str(row["unit"]),
                store=str(row["store"]),
                pricePerUnit=float(row["price_per_unit"]),
            )
            products.append(product)
        except Exception as e:
            print(f"[ERROR] Skipping row due to error: {e}")
    
    return products

def send_products_to_grpc(products: List[product_pb2.AddProductRequest]):
    channel = grpc.insecure_channel(Config.GRPC_SERVER_HOST)
    stub = product_pb2_grpc.ProductServiceStub(channel)

    for i in range(0, len(products), BATCH_SIZE):
        batch = products[i:i + BATCH_SIZE]
        try:
            request = product_pb2.AddProductsRequest(products=batch)
            response = stub.AddProducts(request)
            print(f"[INFO] Sent batch {i//BATCH_SIZE + 1}, {len(response.products)} products added.")
        except grpc.RpcError as e:
            print(f"[GRPC ERROR] Batch {i//BATCH_SIZE + 1} failed: {e.details()}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Send preprocessed product data to product service")
    parser.add_argument("--store", required=True, help="Name of the store (e.g. 'spar', 'hofer')")

    args = parser.parse_args()
    products = read_products_from_csv(args.store.lower())
    send_products_to_grpc(products)
