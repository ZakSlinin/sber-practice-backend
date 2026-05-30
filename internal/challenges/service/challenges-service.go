package service

import (
	"context"
	"fmt"
	"github.com/ZakSlinin/sber-practice-backend/internal/challenges/models"
	"github.com/ZakSlinin/sber-practice-backend/internal/challenges/repository"
	"github.com/google/uuid"
)

type ChallengesService struct {
	repo repository.ChallengesRepo
}

func NewChallengesService(repo repository.ChallengesRepo) *ChallengesService {
	return &ChallengesService{repo: repo}
}

func (s *ChallengesService) CreateChallenge(ctx context.Context, req *models.CreateChallengeRequest, workspaceID, userID uuid.UUID) (*models.Challenge, error) {
	if req.Level != "light" && req.Level != "medium" && req.Level != "hard" {
		return nil, fmt.Errorf("invalid level: must be light, medium or hard")
	}

	challenge := &models.Challenge{
		ID:          uuid.New(),
		WorkspaceID: workspaceID,
		Title:       req.Title,
		Description: req.Description,
		Level:       req.Level,
		IsActive:    true,
		CreatedBy:   userID,
	}

	return s.repo.Create(ctx, challenge)
}

func (s *ChallengesService) GetByWorkspace(ctx context.Context, workspaceID uuid.UUID) ([]*models.Challenge, error) {
	return s.repo.GetByWorkspace(ctx, workspaceID)
}

func (s *ChallengesService) GetByID(ctx context.Context, id uuid.UUID) (*models.Challenge, error) {
	return s.repo.GetByID(ctx, id)
}
