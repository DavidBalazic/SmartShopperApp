using System.Text.Json.Serialization;

namespace UserService.DTOs.Audit
{
    public class AuditLog
    {
        [JsonPropertyName("timestamp")]
        public string Timestamp { get; set; } = DateTime.UtcNow.ToString("o"); // ISO 8601 format

        [JsonPropertyName("actor")]
        public AuditActor Actor { get; set; }

        [JsonPropertyName("action")]
        public string Action { get; set; }

        [JsonPropertyName("resource")]
        public string Resource { get; set; }

        [JsonPropertyName("service")]
        public string Service { get; set; }

        [JsonPropertyName("details")]
        public Dictionary<string, object> Details { get; set; }
    }
}
