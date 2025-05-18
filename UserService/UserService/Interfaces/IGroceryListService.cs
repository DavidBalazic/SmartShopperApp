using UserService.DTOs.GroceryList;
using UserService.Models;

namespace UserService.Interfaces
{
    public interface IGroceryListService
    {
        Task<GroceryListDTO> CreateGroceryListAsync(string userId, string listName, List<string> productIds);
        Task<IEnumerable<GroceryListSummaryDTO>> GetUserGroceryListsAsync(string userId);
        Task<GroceryListDTO> GetGroceryListByIdAsync(int listId, string userId);
        Task<bool> DeleteGroceryListAsync(int listId, string userId);
        Task<bool> AddItemsToListAsync(int listId, string userId, List<string> productIds);
        Task<bool> RemoveItemAsync(int itemId, string userId);
    }
}
