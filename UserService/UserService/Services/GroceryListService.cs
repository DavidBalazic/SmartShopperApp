using Microsoft.EntityFrameworkCore;
using System.Net.Http;
using UserService.Data;
using UserService.DTOs.GroceryList;
using UserService.DTOs.Product;
using UserService.Interfaces;
using UserService.Models;

namespace UserService.Services
{
    public class GroceryListService : IGroceryListService
    {
        private readonly UserContext _context;
        private readonly IProductService _productService;
        private readonly ILogger<GroceryListService> _logger;

        public GroceryListService(UserContext context, IProductService productService, ILogger<GroceryListService> logger)
        {
            _context = context;
            _productService = productService;
            _logger = logger;
        }

        public async Task<int> CreateListAsync(string userId, string listName)
        {
            _logger.LogInformation("Creating grocery list for user {UserId} with name {ListName}", userId, listName);

            var groceryList = new GroceryListModel
            {
                UserId = userId,
                Name = listName
            };

            _context.GroceryLists.Add(groceryList);
            await _context.SaveChangesAsync();

            _logger.LogInformation("Grocery list created with ID {ListId}", groceryList.Id);

            return groceryList.Id;
        }
        public async Task AddItemAsync(int listId, string productId)
        {
            _logger.LogInformation("Adding product {ProductId} to grocery list {ListId}", productId, listId);

            var groceryItem = new GroceryItem
            {
                ProductId = productId,
                GroceryListId = listId
            };

            _context.GroceryItems.Add(groceryItem);
            await _context.SaveChangesAsync();

            _logger.LogInformation("Product {ProductId} added to grocery list {ListId}", productId, listId);
        }
        public async Task<List<GroceryListDTO>> GetUserListsAsync(string userId)
        {
            _logger.LogInformation("Processing grocery list for user {UserId}", userId);

            var groceryLists = await _context.GroceryLists
                .Where(gl => gl.UserId == userId)
                .Include(gl => gl.Items) 
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
                    else
                    {
                        _logger.LogWarning("Product with ID {ProductId} not found", product?.Id);
                    }
                }

                groceryListDTOs.Add(groceryListDTO);
            }
            _logger.LogInformation("Fetched {ListCount} grocery lists for user {UserId}", groceryListDTOs.Count, userId);

            return groceryListDTOs;
        }
    }
}
