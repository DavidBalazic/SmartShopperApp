package models

import "time"

type AuditActor struct {
	ID        string `json:"id" bson:"id"`
	IP        string `json:"ip" bson:"ip"`
	UserAgent string `json:"userAgent" bson:"userAgent"`
}

type AuditLog struct {
	Timestamp time.Time              `json:"timestamp" bson:"timestamp"`
	Actor     AuditActor             `json:"actor" bson:"actor"`
	Action    string                 `json:"action" bson:"action"`
	Resource  string                 `json:"resource" bson:"resource"`
	Service   string                 `json:"service" bson:"service"`
	Details   map[string]interface{} `json:"details" bson:"details"`
}