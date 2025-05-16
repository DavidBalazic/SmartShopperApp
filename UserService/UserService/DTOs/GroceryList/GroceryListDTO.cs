using UserService.DTOs.Product;

namespace UserService.DTOs.GroceryList
{
    public class GroceryListDTO
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public DateTime CreatedAt { get; set; }
        public List<ProductDTO> Items { get; set; }
    }
}
