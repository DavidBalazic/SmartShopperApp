package repository

import "AuditService/internal/models"

type AuditRepository interface {
	SaveLog(log models.AuditLog) error
}