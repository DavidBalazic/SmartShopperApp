using Microsoft.EntityFrameworkCore;
using System.Net.Http;
using System.Xml.Linq;
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

        public async Task<GroceryListDTO> CreateGroceryListAsync(string userId, string listName, List<string> productIds)
        {
            if (string.IsNullOrWhiteSpace(listName))
            {
                listName = GenerateDefaultListName();
                _logger.LogInformation("List name not provided. Generated name: {ListName}", listName);
            }

            _logger.LogInformation("Creating grocery list for user {UserId} with name {ListName}", userId, listName);

            var groceryList = new GroceryListModel
            {
                Name = listName,
                UserId = userId,
                Items = productIds.Select(pid => new GroceryItem { ProductId = pid }).ToList()
            };

            _context.GroceryLists.Add(groceryList);
            await _context.SaveChangesAsync();

            var products = await _productService.GetProductsByIdsAsync(productIds);

            return new GroceryListDTO
            {
                Id = groceryList.Id,
                Name = groceryList.Name,
                CreatedAt = groceryList.CreatedAt,
                Items = products
            };
        }

        public async Task<IEnumerable<GroceryListSummaryDTO>> GetUserGroceryListsAsync(string userId)
        {
            _logger.LogInformation("Retrieving user grocery lists for {UserId}", userId);
            return await _context.GroceryLists
                .Where(gl => gl.UserId == userId)
                .Select(gl => new GroceryListSummaryDTO
                {
                    Id = gl.Id,
                    Name = gl.Name,
                    CreatedAt = gl.CreatedAt
                })
                .ToListAsync();

        }

        public async Task<GroceryListDTO> GetGroceryListByIdAsync(int listId, string userId)
        {
            _logger.LogInformation("Fetching grocery list ID {ListId} for user {UserId}", listId, userId);

            var groceryList = await _context.GroceryLists
                .Include(gl => gl.Items)
                .FirstOrDefaultAsync(gl => gl.Id == listId && gl.UserId == userId);

            if (groceryList == null)
            {
                _logger.LogWarning("Grocery list with ID {ListId} not found for user {UserId}", listId, userId);
                return null;
            }

            var productIds = groceryList.Items.Select(item => item.ProductId).ToList();

            _logger.LogInformation("Fetching {Count} products for grocery list ID {ListId}", productIds.Count, listId);

            var products = await _productService.GetProductsByIdsAsync(productIds);

            var dto = new GroceryListDTO
            {
                Id = groceryList.Id,
                Name = groceryList.Name,
                CreatedAt = groceryList.CreatedAt,
                Items = products
            };

            _logger.LogInformation("Successfully returned grocery list ID {ListId} for user {UserId}", listId, userId);

            return dto;
        }

        public async Task<bool> DeleteGroceryListAsync(int listId, string userId)
        {
            _logger.LogInformation("Deleting grocery list ID {ListId} for user {UserId}", listId, userId);
            var list = await _context.GroceryLists.FirstOrDefaultAsync(gl => gl.Id == listId && gl.UserId == userId);
            if (list == null) return false;

            _context.GroceryLists.Remove(list);
            await _context.SaveChangesAsync();
            return true;
        }

        public async Task<bool> AddItemsToListAsync(int listId, string userId, List<string> productIds)
        {
            _logger.LogInformation("Adding items to grocery list ID {ListId} for user {UserId}", listId, userId);
            var list = await _context.GroceryLists.Include(gl => gl.Items)
                        .FirstOrDefaultAsync(gl => gl.Id == listId && gl.UserId == userId);
            if (list == null) return false;

            var newItems = productIds.Select(pid => new GroceryItem { ProductId = pid }).ToList();
            list.Items.AddRange(newItems);
            await _context.SaveChangesAsync();
            return true;
        }

        public async Task<bool> RemoveItemAsync(int itemId, string userId)
        {
            _logger.LogInformation("Removing item ID {ItemId} from grocery list for user {UserId}", itemId, userId);
            var item = await _context.GroceryItems
                .Include(i => i.GroceryList)
                .FirstOrDefaultAsync(i => i.Id == itemId && i.GroceryList.UserId == userId);

            if (item == null) return false;

            _context.GroceryItems.Remove(item);
            await _context.SaveChangesAsync();
            return true;
        }

        private string GenerateDefaultListName()
        {
            var now = DateTime.Now;
            var dayOfWeek = now.DayOfWeek.ToString().ToLower();
            var hour = now.Hour;

            string timeOfDay = hour switch
            {
                >= 5 and < 12 => "morning",
                >= 12 and < 17 => "afternoon",
                >= 17 and < 21 => "evening",
                _ => "night"
            };

            return $"Shopping list for {dayOfWeek} {timeOfDay}";
        }
    }
}
