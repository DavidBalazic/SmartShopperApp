package services

import (
	"AuditService/internal/models"
	"AuditService/internal/repository"
)

type auditService struct {
	repo repository.AuditRepository
}

func NewAuditService(repo repository.AuditRepository) AuditService {
	return &auditService{repo}
}

func (s *auditService) HandleAuditLog(log models.AuditLog) error {
	return s.repo.SaveLog(log)
}
