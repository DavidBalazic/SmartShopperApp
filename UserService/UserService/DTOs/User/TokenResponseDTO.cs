namespace UserService.DTOs.User
{
    public class TokenResponseDTO
    {
        public string Token { get; set; }
        public DateTime Expiration { get; set; }
    }
}
