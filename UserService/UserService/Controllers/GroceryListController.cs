using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using UserService.DTOs.GroceryList;
using UserService.Interfaces;

namespace UserService.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class GroceryListController : ControllerBase
    {
        private readonly IGroceryListService _groceryListService;
        public GroceryListController(IGroceryListService groceryListService)
        {
            _groceryListService = groceryListService;
        }

        [HttpPost("create-list")]
        public async Task<IActionResult> CreateList([FromBody] CreateGroceryListDTO request)
        {
            if (string.IsNullOrEmpty(request.UserId) || string.IsNullOrEmpty(request.Name))
                return BadRequest("UserId and ListName are required.");

            var listId = await _groceryListService.CreateListAsync(request.UserId, request.Name);
            return Ok(new { ListId = listId });
        }

        [HttpPost("add-item")]
        public async Task<IActionResult> AddItem([FromBody] AddItemDTO request)
        {
            if (request.ListId <= 0 || string.IsNullOrEmpty(request.ProductId))
                return BadRequest("ListId and ProductId are required.");

            await _groceryListService.AddItemAsync(request.ListId, request.ProductId);
            return Ok("Item added successfully.");
        }

        [HttpGet("user/{userId}")]
        public async Task<IActionResult> GetUserGroceryLists(string userId)
        {
            var groceryLists = await _groceryListService.GetUserListsAsync(userId);
            if (groceryLists == null || !groceryLists.Any())
                return NotFound("No grocery lists found for the user.");

            return Ok(groceryLists);
        }
    }
}
