using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;

namespace UserService.Data
{
    public class UserContext : IdentityDbContext<IdentityUser>
    {
        public UserContext(DbContextOptions<UserContext> options) : base(options)
        {
            
        }
    }
}
