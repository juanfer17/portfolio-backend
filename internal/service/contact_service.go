package service

import (
	"context"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repository"
	"time"
)

type ContactService interface {
	ProcessContactMessage(ctx context.Context, msg models.ContactMessage) error
}

type contactService struct {
	repo repository.PortfolioRepository
}

func NewContactService(repo repository.PortfolioRepository) ContactService {
	return &contactService{repo: repo}
}

func (s *contactService) ProcessContactMessage(ctx context.Context, msg models.ContactMessage) error {
	// Business Logic: Set creation timestamp
	msg.CreatedAt = time.Now()

	// Future enhancement: Send email notification here (e.g., SendGrid, AWS SES)
	// if err := emailProvider.Send(...); err != nil { ... }

	// Persist to database
	return s.repo.SaveContactMessage(ctx, msg)
}
