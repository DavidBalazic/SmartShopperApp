import grpc
import app.proto.product_pb2 as product_pb2
import app.proto.product_pb2_grpc as product_pb2_grpc
from app.core.config import Config

def get_product_by_id(product_id: str):
    channel = grpc.insecure_channel(Config.GRPC_SERVER_HOST)
    stub = product_pb2_grpc.ProductServiceStub(channel)
    request = product_pb2.ProductIdRequest(id=product_id)
    try:
        response = stub.GetProductById(request)
        return response.product
    except grpc.RpcError as e:
        raise Exception(f"Error calling gRPC service: {e}")