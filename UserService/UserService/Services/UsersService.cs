using Microsoft.AspNetCore.Identity;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using UserService.DTOs.Audit;
using UserService.DTOs.User;
using UserService.Interfaces;
using UserService.Models;

namespace UserService.Services
{
    public class UsersService : IUserService
    {
        private readonly UserManager<IdentityUser> _userManager;
        private readonly RoleManager<IdentityRole> _roleManager;
        private readonly IConfiguration _configuration;
        private readonly IAuditLogger _auditLogger;

        public UsersService(UserManager<IdentityUser> userManager, RoleManager<IdentityRole> roleManager, IConfiguration configuration, IAuditLogger auditLogger)
        {
            _userManager = userManager;
            _roleManager = roleManager;
            _configuration = configuration;
            _auditLogger = auditLogger;
        }

        public async Task<RegisterResponse> RegisterUserAsync(RegisterDTO registerDto)
        {
            var userExists = await _userManager.FindByNameAsync(registerDto.Username);
            if (userExists != null)
            {
                return new RegisterResponse { Success = false, Message = "User already exists" };
            }

            var user = new IdentityUser
            {
                Email = registerDto.Email,
                SecurityStamp = Guid.NewGuid().ToString(),
                UserName = registerDto.Username
            };

            var result = await _userManager.CreateAsync(user, registerDto.Password);

            var auditDetails = new Dictionary<string, object> 
            {
                { "username", registerDto.Username },
                { "email", registerDto.Email },
                { "status", result.Succeeded ? "success" : "fail" }
            };

            await _auditLogger.LogAsync(new AuditLog
            {
                Actor = new AuditActor { ID = registerDto.Username },
                Action = "register",
                Resource = "user",
                Service = "UserService",
                Details = auditDetails
            });

            if (!result.Succeeded)
            {
                return new RegisterResponse { Success = false, Message = "User creation failed. Please check the details and try again." };
            }

            return new RegisterResponse { Success = true, Message = "User created successfully" };
        }

        public async Task<TokenResponseDTO> LoginUserAsync(LoginDTO loginDto)
        {
            var user = await _userManager.FindByNameAsync(loginDto.Username);
            var loginSuccess = user != null && await _userManager.CheckPasswordAsync(user, loginDto.Password);

            await _auditLogger.LogAsync(new AuditLog
            {
                Actor = new AuditActor { ID = loginDto.Username },
                Action = "login",
                Resource = "user",
                Service = "UserService",
                Details = new Dictionary<string, object>
                {
                    { "username", loginDto.Username },
                    { "status", loginSuccess ? "success" : "fail" }
                }
            });

            if (!loginSuccess)
                throw new UnauthorizedAccessException("Invalid credentials");

            var userRoles = await _userManager.GetRolesAsync(user);

            var authClaims = new List<Claim>
            {
                new Claim(ClaimTypes.Name, user.UserName),
                new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
            };

            foreach (var role in userRoles)
            {
                authClaims.Add(new Claim(ClaimTypes.Role, role));
            }

            var authSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_configuration["JwtSettings:SecretKey"]));

            var token = new JwtSecurityToken(
                issuer: _configuration["JwtSettings:Issuer"],
                audience: _configuration["JwtSettings:Audience"],
                expires: DateTime.Now.AddHours(3),
                claims: authClaims,
                signingCredentials: new SigningCredentials(authSigningKey, SecurityAlgorithms.HmacSha256)
            );

            var tokenString = new JwtSecurityTokenHandler().WriteToken(token);
            return new TokenResponseDTO
            {
                Token = new JwtSecurityTokenHandler().WriteToken(token),
                Expiration = token.ValidTo
            };
        }

        public ClaimsPrincipal? IsTokenValid(string token)
        {
            try
            {
                var tokenHandler = new JwtSecurityTokenHandler();
                var key = Encoding.UTF8.GetBytes(_configuration["JwtSettings:SecretKey"]);

                var validationParameters = new TokenValidationParameters
                {
                    ValidateIssuer = true,
                    ValidIssuer = _configuration["JwtSettings:Issuer"],

                    ValidateAudience = true,
                    ValidAudience = _configuration["JwtSettings:Audience"],

                    ValidateLifetime = true,
                    ClockSkew = TimeSpan.Zero, 

                    ValidateIssuerSigningKey = true,
                    IssuerSigningKey = new SymmetricSecurityKey(key),

                    RoleClaimType = ClaimTypes.Role
                };

                SecurityToken validatedToken;
                var principal = tokenHandler.ValidateToken(token, validationParameters, out validatedToken);
                return principal;
            }
            catch (SecurityTokenExpiredException)
            {
                return null;
            }
            catch (Exception)
            {
                return null;
            }
        }
    }
}
