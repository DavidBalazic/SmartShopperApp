package dtos

type AuditActor struct {
	ID        string `json:"id"`
	IP        string `json:"ip"`
	UserAgent string `json:"userAgent"`
}

type AuditLog struct {
	Timestamp string                 `json:"timestamp"`
	Actor     AuditActor             `json:"actor"`
	Action    string                 `json:"action"`
	Resource  string                 `json:"resource"`
	Service   string                 `json:"service"`
	Details   map[string]interface{} `json:"details"`
}