using UserService.DTOs;

namespace UserService.Interfaces
{
    public interface IUserService
    {
        Task<TokenResponseDTO> RegisterUserAsync(RegisterDTO registerDto);
        Task<TokenResponseDTO> LoginUserAsync(LoginDTO loginDto);
        bool IsTokenValid(string token);
    }
}
