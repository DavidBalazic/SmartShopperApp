namespace UserService.DTOs.GroceryList
{
    public class CreateGroceryListDTO
    {
        public string Name { get; set; }
        public List<string> ProductIds { get; set; }
    }
}
