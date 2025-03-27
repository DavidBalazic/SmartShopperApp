using UserService.DTOs;

namespace UserService.Interfaces
{
    public interface IProductService
    {
        public Task<ProductDTO> GetProductByIdAsync(string id);
    }
}
