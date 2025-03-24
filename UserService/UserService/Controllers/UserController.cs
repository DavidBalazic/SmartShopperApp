using Microsoft.AspNetCore.Mvc;
using UserService.DTOs;
using UserService.Interfaces;

namespace UserService.Controllers
{

    [Route("api/[controller]")]
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly IUserService _userService;
        public UserController(IUserService userService)
        {
            _userService = userService;
        }
        [HttpPost("register")]
        public async Task<IActionResult> Register([FromBody] RegisterDTO registerDto)
        {
            if (!ModelState.IsValid)
                return BadRequest("Invalid data.");
            try
            {
                var tokenResponse = await _userService.RegisterUserAsync(registerDto);
                return Ok(tokenResponse);
            }
            catch (Exception ex)
            {
                return BadRequest(ex.Message);
            }
        }
        [HttpPost("login")]
        public async Task<IActionResult> Login([FromBody] LoginDTO loginDto)
        {
            if (!ModelState.IsValid)
                return BadRequest("Invalid data.");
            try
            {
                var tokenResponse = await _userService.LoginUserAsync(loginDto);
                return Ok(tokenResponse);
            }
            catch (Exception ex)
            {
                return BadRequest(ex.Message);
            }
        }

        [HttpPost("validate-token")]
        public IActionResult ValidateToken([FromBody] string token)
        {
            if (string.IsNullOrEmpty(token))
                return BadRequest("Token is required.");

            bool isValid = _userService.IsTokenValid(token);

            if (isValid)
                return Ok(new { message = "Token is valid." });
            else
                return Unauthorized(new { message = "Invalid or expired token." });
        }
    }
}
