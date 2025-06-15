using Grpc.Core;
using UserService.DTOs.Product;
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
                    PricePerUnit = product.PricePerUnit,
                    ImageUrl = product.ImageUrl
                };
            }
            catch (RpcException ex)
            {
                _logger.LogError(ex, "gRPC Error: {Status}", ex.Status);
                return null;
            }
        }

        public async Task<List<ProductDTO>> GetProductsByIdsAsync(List<string> productIds)
        {
            var request = new ProductsIdsRequest();
            request.Ids.AddRange(productIds);

            _logger.LogInformation("Sending request to gRPC service to fetch products by IDs: {Ids}", string.Join(", ", productIds));

            try
            {
                var response = await _productClient.GetProductsByIdsAsync(request);
                _logger.LogInformation("Received {Count} products from gRPC", response.Products.Count);

                return response.Products.Select(p => new ProductDTO
                {
                    Id = p.Id,
                    Name = p.Name,
                    Description = p.Description,
                    Price = p.Price,
                    Quantity = p.Quantity,
                    Unit = p.Unit,
                    Store = p.Store,
                    PricePerUnit = p.PricePerUnit,
                    ImageUrl = p.ImageUrl
                }).ToList();
            }
            catch (RpcException ex)
            {
                _logger.LogError(ex, "gRPC error during GetProductsByIdsAsync: {Status}", ex.Status);
                return new List<ProductDTO>();
            }
        }
    }
}
