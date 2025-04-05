using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;
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
                return BadRequest(ModelState);

            var registerResponse = await _userService.RegisterUserAsync(registerDto);

            if (!registerResponse.Success)
                return BadRequest(new { status = "error", registerResponse.Message });

            return Ok(new { status = "success", registerResponse.Message });
        }
        [HttpPost("login")]
        public async Task<IActionResult> Login([FromBody] LoginDTO loginDto)
        {
            if (!ModelState.IsValid)
                return BadRequest(ModelState);

            try
            {
                var tokenResponse = await _userService.LoginUserAsync(loginDto);
                return Ok(tokenResponse);
            }
            catch (UnauthorizedAccessException)
            {
                return Unauthorized("Invalid credentials");
            }
        }

        [HttpPost("validate-token")]
        public IActionResult ValidateToken([FromQuery] string token)
        {
            var principal = _userService.IsTokenValid(token);
            if (principal == null)
                return Unauthorized(new { message = "Invalid or expired token." });

            return Ok(new
            {
                username = principal.Identity?.Name,
                roles = principal.Claims.Where(c => c.Type == ClaimTypes.Role).Select(c => c.Value).ToList()
            });
        }
    }
}
