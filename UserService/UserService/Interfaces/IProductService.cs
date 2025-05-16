using UserService.DTOs.Product;

namespace UserService.Interfaces
{
    public interface IProductService
    {
        public Task<ProductDTO> GetProductByIdAsync(string id);
    }
}
