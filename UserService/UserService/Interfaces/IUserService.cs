using System.Security.Claims;
using UserService.DTOs.User;

namespace UserService.Interfaces
{
    public interface IUserService
    {
        Task<RegisterResponse> RegisterUserAsync(RegisterDTO registerDto);
        Task<TokenResponseDTO> LoginUserAsync(LoginDTO loginDto);
        ClaimsPrincipal? IsTokenValid(string token);
    }
}
