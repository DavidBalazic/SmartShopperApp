using System.Text.Json.Serialization;

namespace UserService.DTOs.Audit
{
    public class AuditActor
    {
        [JsonPropertyName("id")]
        public string ID { get; set; }

        [JsonPropertyName("ip")]
        public string IP { get; set; }

        [JsonPropertyName("userAgent")]
        public string UserAgent { get; set; }
    }
}
