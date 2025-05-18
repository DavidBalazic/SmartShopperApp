using Microsoft.AspNetCore.Identity;

namespace UserService.Models
{
    public class GroceryListModel
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
        public List<GroceryItem> Items { get; set; } = new();

        public string UserId { get; set; }
        public IdentityUser User { get; set; }
    }
    public class GroceryItem
    {
        public int Id { get; set; }
        public string ProductId { get; set; }
        public int GroceryListId { get; set; }
        public GroceryListModel GroceryList { get; set; }
    }
}
