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
    private readonly GroceryListService _service;
    private readonly UserContext _context;
    private readonly Mock<IProductService> _mockProductService;
    private readonly Mock<ILogger<GroceryListService>> _mockLogger;

    public GroceryListServiceTests()
    {
        var options = new DbContextOptionsBuilder<UserContext>()
            .UseInMemoryDatabase(databaseName: Guid.NewGuid().ToString())
            .Options;
        _context = new UserContext(options);

        _mockProductService = new Mock<IProductService>();
        _mockLogger = new Mock<ILogger<GroceryListService>>();
        _service = new GroceryListService(_context, _mockProductService.Object, _mockLogger.Object);
    }

    [Fact]
    public async Task CreateGroceryListAsync_ShouldCreateListWithDefaults()
    {
        // Arrange
        var userId = "user123";
        var listName = "";
        var productIds = new List<string> { "p1", "p2" };

        _mockProductService.Setup(p => p.GetProductsByIdsAsync(productIds))
            .ReturnsAsync(productIds.Select(id => new ProductDTO { Id = id }).ToList());

        // Act
        var result = await _service.CreateGroceryListAsync(userId, listName, productIds);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Items.Count);
        Assert.Contains("Shopping list for", result.Name);
        Assert.True(_context.GroceryLists.Any(gl => gl.UserId == userId));
    }

    [Fact]
    public async Task GetUserGroceryListsAsync_ShouldReturnUserLists()
    {
        // Arrange
        var userId = "user1";
        _context.GroceryLists.Add(new GroceryListModel { UserId = userId, Name = "Weekly", CreatedAt = DateTime.UtcNow });
        _context.SaveChanges();

        // Act
        var lists = await _service.GetUserGroceryListsAsync(userId);

        // Assert
        Assert.Single(lists);
        Assert.Equal("Weekly", lists.First().Name);
    }

    [Fact]
    public async Task GetGroceryListByIdAsync_ReturnsListWithProducts()
    {
        // Arrange
        var userId = "user1";
        var list = new GroceryListModel
        {
            UserId = userId,
            Name = "TestList",
            Items = new List<GroceryItem> { new GroceryItem { ProductId = "p1" } }
        };
        _context.GroceryLists.Add(list);
        _context.SaveChanges();

        _mockProductService.Setup(p => p.GetProductsByIdsAsync(It.IsAny<List<string>>()))
            .ReturnsAsync(new List<ProductDTO> { new ProductDTO { Id = "p1" } });

        // Act
        var result = await _service.GetGroceryListByIdAsync(list.Id, userId);

        // Assert
        Assert.NotNull(result);
        Assert.Equal("TestList", result.Name);
        Assert.Single(result.Items);
    }

    [Fact]
    public async Task DeleteGroceryListAsync_DeletesList()
    {
        // Arrange
        var userId = "user2";
        var list = new GroceryListModel { UserId = userId, Name = "ToDelete" };
        _context.GroceryLists.Add(list);
        _context.SaveChanges();

        // Act
        var success = await _service.DeleteGroceryListAsync(list.Id, userId);

        // Assert
        Assert.True(success);
        Assert.Empty(_context.GroceryLists.Where(l => l.Id == list.Id));
    }

    [Fact]
    public async Task AddItemsToListAsync_AddsItemsCorrectly()
    {
        // Arrange
        var userId = "user3";
        var list = new GroceryListModel
        {
            UserId = userId,
            Name = "ToAddTo",
            Items = new List<GroceryItem>()
        };
        _context.GroceryLists.Add(list);
        _context.SaveChanges();

        var newProductIds = new List<string> { "p10", "p20" };

        // Act
        var result = await _service.AddItemsToListAsync(list.Id, userId, newProductIds);

        // Assert
        Assert.True(result);
        var updatedList = _context.GroceryLists.Include(gl => gl.Items).First(gl => gl.Id == list.Id);
        Assert.Equal(2, updatedList.Items.Count);
    }

    [Fact]
    public async Task RemoveItemAsync_RemovesCorrectItem()
    {
        // Arrange
        var userId = "user4";
        var item = new GroceryItem { ProductId = "x" };
        var list = new GroceryListModel
        {
            UserId = userId,
            Name = "WithItem",
            Items = new List<GroceryItem> { item }
        };
        _context.GroceryLists.Add(list);
        _context.SaveChanges();

        var itemId = list.Items.First().Id;

        // Act
        var result = await _service.RemoveItemAsync(itemId, userId);

        // Assert
        Assert.True(result);
        Assert.Empty(_context.GroceryItems);
    }
}