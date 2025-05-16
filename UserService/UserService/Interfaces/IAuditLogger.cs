using UserService.DTOs.Audit;

namespace UserService.Interfaces
{
    public interface IAuditLogger
    {
        Task LogAsync(AuditLog log);
    }
}
