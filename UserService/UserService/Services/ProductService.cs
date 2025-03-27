using Grpc.Core;
using UserService.DTOs;
using UserService.Interfaces;
using UserService.Protos;

namespace UserService.Services
{
    public class ProductService : IProductService
    {
        private readonly UserService.Protos.ProductService.ProductServiceClient _productClient;
        private readonly ILogger<ProductService> _logger;

        public ProductService(UserService.Protos.ProductService.ProductServiceClient productClient, ILogger<ProductService> logger)
        {
            _productClient = productClient;
            _logger = logger;
        }

        public async Task<ProductDTO> GetProductByIdAsync(string productId)
        {
            var request = new ProductIdRequest { Id = productId };
            _logger.LogInformation("Sending request to gRPC service with ProductId: {ProductId}", productId);

            try
            {
                var response = await _productClient.GetProductByIdAsync(request);
                _logger.LogInformation("Received response from gRPC: {@Response}", response);
                var product = response.Product;

                return new ProductDTO
                {
                    Id = product.Id,
                    Name = product.Name,
                    Description = product.Description,
                    Price = product.Price,
                    Quantity = product.Quantity,
                    Unit = product.Unit,
                    Store = product.Store,
                    PricePerUnit = product.PricePerUnit
                };
            }
            catch (RpcException ex)
            {
                _logger.LogError(ex, "gRPC Error: {Status}", ex.Status);
                return null;
            }
        }
    }
}
