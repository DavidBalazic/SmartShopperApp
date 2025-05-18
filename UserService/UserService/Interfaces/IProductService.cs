using UserService.DTOs.Product;

namespace UserService.Interfaces
{
    public interface IProductService
    {
        public Task<ProductDTO> GetProductByIdAsync(string id);
        public Task<List<ProductDTO>> GetProductsByIdsAsync(List<string> ids);
    }
}
