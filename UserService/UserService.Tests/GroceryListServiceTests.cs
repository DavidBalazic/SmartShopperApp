using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using Moq;
using UserService.Data;
using UserService.DTOs.Product;
using UserService.Interfaces;
using UserService.Models;
using UserService.Services;

namespace UserService.Tests;

public class GroceryListServiceTests
{
    private readonly UserContext _context;
    private readonly Mock<IProductService> _productServiceMock;
    private readonly Mock<ILogger<GroceryListService>> _loggerMock;
    private readonly GroceryListService _groceryListService;

    public GroceryListServiceTests()
    {
        var options = new DbContextOptionsBuilder<UserContext>()
            .UseInMemoryDatabase(databaseName: Guid.NewGuid().ToString())
            .Options;

        _context = new UserContext(options);
        _productServiceMock = new Mock<IProductService>();
        _loggerMock = new Mock<ILogger<GroceryListService>>();

        _groceryListService = new GroceryListService(_context, _productServiceMock.Object, _loggerMock.Object);
    }

    [Fact]
    public async Task CreateListAsync_ShouldCreateListAndReturnId()
    {
        string userId = "user123";
        string listName = "Weekly Shopping";

        int listId = await _groceryListService.CreateListAsync(userId, listName);

        var list = await _context.GroceryLists.FindAsync(listId);
        Assert.NotNull(list);
        Assert.Equal(userId, list.UserId);
        Assert.Equal(listName, list.Name);
    }
    [Fact]
    public async Task AddItemAsync_ShouldAddItemToList()
    {
        var groceryList = new GroceryListModel { UserId = "user123", Name = "My List" };
        _context.GroceryLists.Add(groceryList);
        await _context.SaveChangesAsync();

        string productId = "product123";

        await _groceryListService.AddItemAsync(groceryList.Id, productId);

        var item = await _context.GroceryItems.FirstOrDefaultAsync(i => i.ProductId == productId);
        Assert.NotNull(item);
        Assert.Equal(groceryList.Id, item.GroceryListId);
    }
    [Fact]
    public async Task GetUserListsAsync_ShouldReturnUserListsWithProducts()
    {
        string userId = "user123";

        var groceryList = new GroceryListModel { UserId = userId, Name = "My List", Items = new List<GroceryItem>() };
        _context.GroceryLists.Add(groceryList);
        await _context.SaveChangesAsync();

        var product = new ProductDTO
        {
            Id = "product123",
            Name = "Milk",
            Description = "1L Milk",
            Price = 1.99,
            Quantity = 1,
            Unit = "L",
            Store = "Supermarket",
            PricePerUnit = 1.99
        };

        _productServiceMock
            .Setup(service => service.GetProductByIdAsync("product123"))
            .ReturnsAsync(product);

        groceryList.Items.Add(new GroceryItem { ProductId = "product123", GroceryListId = groceryList.Id });
        await _context.SaveChangesAsync();

        var result = await _groceryListService.GetUserListsAsync(userId);

        Assert.Single(result);
        Assert.Equal("My List", result[0].Name);
        Assert.Single(result[0].Items);
        Assert.Equal("Milk", result[0].Items[0].Name);
    }
}