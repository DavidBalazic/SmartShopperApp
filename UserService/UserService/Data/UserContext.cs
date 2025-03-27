using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using UserService.Models;

namespace UserService.Data
{
    public class UserContext : IdentityDbContext<IdentityUser>
    {
        public UserContext(DbContextOptions<UserContext> options) : base(options) {}
        public DbSet<GroceryListModel> GroceryLists { get; set; }
        public DbSet<GroceryItem> GroceryItems { get; set; }
        protected override void OnModelCreating(ModelBuilder builder)
        {
            base.OnModelCreating(builder);

            builder.Entity<GroceryListModel>()
                .HasOne(gl => gl.User)
                .WithMany()
                .HasForeignKey(gl => gl.UserId);

            builder.Entity<GroceryItem>()
                .HasOne(gi => gi.GroceryList)
                .WithMany(gl => gl.Items)
                .HasForeignKey(gi => gi.GroceryListId);
        }
    }
}
