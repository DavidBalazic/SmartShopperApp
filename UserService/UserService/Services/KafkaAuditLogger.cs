using Confluent.Kafka;
using Microsoft.Extensions.Options;
using System.Text.Json;
using UserService.DTOs.Audit;
using UserService.Interfaces;
using UserService.Models.Configs;

namespace UserService.Services
{
    public class KafkaAuditLogger : IAuditLogger
    {
        private readonly IProducer<string, string> _producer;
        private readonly string _topic;
        private readonly ILogger<KafkaAuditLogger> _logger;
        private readonly IHttpContextAccessor _httpContextAccessor;

        public KafkaAuditLogger(IOptions<KafkaSettings> options, ILogger<KafkaAuditLogger> logger, IHttpContextAccessor httpContextAccessor)
        {
            var settings = options.Value;

            var config = new ProducerConfig
            {
                BootstrapServers = settings.KafkaBroker,
                MessageTimeoutMs = 1000
            };

            _producer = new ProducerBuilder<string, string>(config).Build();
            _topic = settings.KafkaTopic;
            _logger = logger;
            _httpContextAccessor = httpContextAccessor;

            _logger.LogInformation(settings.KafkaBroker);
            _logger.LogInformation(settings.KafkaTopic);
        }

        public async Task LogAsync(AuditLog log)
        {
            try
            {
                var context = _httpContextAccessor.HttpContext;
                if (context != null)
                {
                    var ip = context.Connection.RemoteIpAddress?.ToString() ?? "unknown";
                    var userAgent = context.Request.Headers["User-Agent"].ToString() ?? "unknown";

                    // Override actor's IP/UserAgent if missing or empty
                    if (log.Actor == null)
                        log.Actor = new AuditActor();

                    if (string.IsNullOrEmpty(log.Actor.IP))
                        log.Actor.IP = ip;

                    if (string.IsNullOrEmpty(log.Actor.UserAgent))
                        log.Actor.UserAgent = userAgent;
                }

                var json = JsonSerializer.Serialize(log);
                var key = log.Actor.ID;
                await _producer.ProduceAsync(_topic, new Message<string, string>
                {
                    Key = key,
                    Value = json
                });
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Failed to send audit log to Kafka");
            }
        }
    }
}
