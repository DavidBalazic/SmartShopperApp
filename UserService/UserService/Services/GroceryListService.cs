using Microsoft.EntityFrameworkCore;
using System.Net.Http;
using UserService.Data;
using UserService.DTOs;
using UserService.Interfaces;
using UserService.Models;

namespace UserService.Services
{
    public class GroceryListService : IGroceryListService
    {
        private readonly UserContext _context;
        private readonly IProductService _productService;

        public GroceryListService(UserContext context, IProductService productService)
        {
            _context = context;
            _productService = productService;
        }

        public async Task<int> CreateListAsync(string userId, string listName)
        {
            var groceryList = new GroceryListModel
            {
                UserId = userId,
                Name = listName
            };

            _context.GroceryLists.Add(groceryList);
            await _context.SaveChangesAsync();

            return groceryList.Id;
        }
        public async Task AddItemAsync(int listId, string productId)
        {
            var groceryItem = new GroceryItem
            {
                ProductId = productId,
                GroceryListId = listId
            };

            _context.GroceryItems.Add(groceryItem);
            await _context.SaveChangesAsync();
        }
        public async Task<List<GroceryListDTO>> GetUserListsAsync(string userId)
        {
            var groceryLists = await _context.GroceryLists
       .Where(gl => gl.UserId == userId)
       .Include(gl => gl.Items) // Ensure related items are loaded
       .ToListAsync();

            var groceryListDTOs = new List<GroceryListDTO>();

            foreach (var groceryList in groceryLists)
            {
                var groceryListDTO = new GroceryListDTO
                {
                    Id = groceryList.Id,
                    Name = groceryList.Name,
                    Items = new List<ProductDTO>()
                };

                var productTasks = groceryList.Items
                    .Select(item => _productService.GetProductByIdAsync(item.ProductId))
                    .ToList();

                var products = await Task.WhenAll(productTasks);

                foreach (var product in products)
                {
                    if (product != null)
                    {
                        groceryListDTO.Items.Add(new ProductDTO
                        {
                            Id = product.Id,
                            Name = product.Name,
                            Description = product.Description,
                            Price = product.Price,
                            Quantity = product.Quantity,
                            Unit = product.Unit,
                            Store = product.Store,
                            PricePerUnit = product.PricePerUnit
                        });
                    }
                }

                groceryListDTOs.Add(groceryListDTO);
            }

            return groceryListDTOs;
        }
    }
}
