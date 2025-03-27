using UserService.DTOs;

namespace UserService.Interfaces
{
    public interface IGroceryListService
    {
        Task<int> CreateListAsync(string userId, string listName);
        Task AddItemAsync(int listId, string item);
        Task<List<GroceryListDTO>> GetUserListsAsync(string userId);
    }
}
