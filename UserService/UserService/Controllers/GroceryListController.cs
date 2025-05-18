using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;
using UserService.DTOs.GroceryList;
using UserService.Interfaces;

namespace UserService.Controllers
{
    [Authorize]
    [ApiController]
    [Route("api/[controller]")]
    public class GroceryListController : ControllerBase
    {
        private readonly IGroceryListService _groceryListService;
        public GroceryListController(IGroceryListService groceryListService)
        {
            _groceryListService = groceryListService;
        }

        [HttpGet]
        public async Task<IActionResult> GetLists()
        {
            var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
            var groceryLists = await _groceryListService.GetUserGroceryListsAsync(userId);
            if (groceryLists == null || !groceryLists.Any())
                return NotFound("No grocery lists found for the user.");

            return Ok(groceryLists);
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> GetList(int id)
        {
            var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
            var list = await _groceryListService.GetGroceryListByIdAsync(id, userId);
            if (list == null) return NotFound();
            return Ok(list);
        }

        [HttpPost]
        public async Task<IActionResult> CreateList([FromBody] CreateGroceryListDTO dto)
        {
            var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
            var list = await _groceryListService.CreateGroceryListAsync(userId, dto.Name, dto.ProductIds);
            return CreatedAtAction(nameof(GetList), new { id = list.Id }, list);
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteList(int id)
        {
            var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
            var success = await _groceryListService.DeleteGroceryListAsync(id, userId);
            return success ? NoContent() : NotFound();
        }

        [HttpPost("{id}/items")]
        public async Task<IActionResult> AddItems(int id, [FromBody] List<string> productIds)
        {
            var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
            var success = await _groceryListService.AddItemsToListAsync(id, userId, productIds);
            return success ? Ok() : NotFound();
        }

        [HttpDelete("items/{itemId}")]
        public async Task<IActionResult> RemoveItem(int itemId)
        {
            var userId = User.FindFirstValue(ClaimTypes.NameIdentifier);
            var success = await _groceryListService.RemoveItemAsync(itemId, userId);
            return success ? NoContent() : NotFound();
        }
    }
}
