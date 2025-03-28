using Microsoft.AspNetCore.Identity;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using UserService.DTOs;
using UserService.Interfaces;

namespace UserService.Services
{
    //
    public class UsersService : IUserService
    {
        private readonly UserManager<IdentityUser> _userManager;
        private readonly SignInManager<IdentityUser> _signInManager;
        private readonly IConfiguration _configuration;

        public UsersService(UserManager<IdentityUser> userManager, SignInManager<IdentityUser> signInManager, IConfiguration configuration)
        {
            _userManager = userManager;
            _signInManager = signInManager;
            _configuration = configuration;
        }

        public async Task<TokenResponseDTO> RegisterUserAsync(RegisterDTO registerDto)
        {
            var user = new IdentityUser
            {
                UserName = registerDto.UserName,
                Email = registerDto.Email
            };

            var result = await _userManager.CreateAsync(user, registerDto.Password);

            if (!result.Succeeded)
            {
                throw new Exception("User registration failed. Please try again.");
            }

            var token = GenerateJwtToken(user);

            return new TokenResponseDTO
            {
                Token = token,
                Expiration = DateTime.Now.AddMinutes(int.Parse(_configuration["JwtSettings:ExpiryMinutes"]))
            };
        }

        public async Task<TokenResponseDTO> LoginUserAsync(LoginDTO loginDto)
        {
            var user = await _userManager.FindByEmailAsync(loginDto.Email);
            if (user == null)
            {
                throw new Exception("Invalid login attempt.");
            }

            var signInResult = await _signInManager.PasswordSignInAsync(user, loginDto.Password, false, false);
            if (signInResult.Succeeded)
            {
                var token = GenerateJwtToken(user);
                return new TokenResponseDTO
                {
                    Token = token,
                    Expiration = DateTime.Now.AddMinutes(int.Parse(_configuration["JwtSettings:ExpiryMinutes"]))
                };
            }

            throw new Exception("Invalid login attempt.");
        }

        private string GenerateJwtToken(IdentityUser user)
        {
            var claims = new List<Claim>
            {
                new Claim(ClaimTypes.Name, user.UserName),
                new Claim(ClaimTypes.Email, user.Email),
                new Claim(ClaimTypes.NameIdentifier, user.Id)
            };

            var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_configuration["JwtSettings:SecretKey"]));
            var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

            var token = new JwtSecurityToken(
                _configuration["JwtSettings:Issuer"],
                _configuration["JwtSettings:Audience"],
                claims,
                expires: DateTime.Now.AddMinutes(int.Parse(_configuration["JwtSettings:ExpiryMinutes"])),
                signingCredentials: creds);

            return new JwtSecurityTokenHandler().WriteToken(token);
        }

        public bool IsTokenValid(string token)
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var key = Encoding.UTF8.GetBytes(_configuration["JwtSettings:SecretKey"]);

            try
            {
                tokenHandler.ValidateToken(token, new TokenValidationParameters
                {
                    ValidateIssuerSigningKey = true,
                    IssuerSigningKey = new SymmetricSecurityKey(key),
                    ValidateIssuer = true,
                    ValidIssuer = _configuration["JwtSettings:Issuer"],
                    ValidateAudience = true,
                    ValidAudience = _configuration["JwtSettings:Audience"],
                    ValidateLifetime = true,
                    ClockSkew = TimeSpan.Zero
                }, out SecurityToken validatedToken);

                return true;
            }
            catch
            {
                return false;
            }
        }
    }
}
