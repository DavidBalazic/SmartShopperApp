﻿namespace UserService.Models
{
    public class UserModel
    {
        public string? Id { get; set; }

        public string? Email { get; set; }

        public string? PasswordHash { get; set; } 

        public string? Role { get; set; }
    }
}
