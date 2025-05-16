package services

import "AuditService/internal/models"

type AuditService interface {
	HandleAuditLog(log models.AuditLog) error
}