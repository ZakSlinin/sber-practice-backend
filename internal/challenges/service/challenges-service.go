package service

import (
	"context"
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
	challenge := &models.Challenge{
		ID:          uuid.New(),
		WorkspaceID: workspaceID,
		Description: req.Description,
		Level:       req.Level,
		IsActive:    true,
		CreatedBy:   userID,
	}

	return s.repo.Create(ctx, challenge)
}
