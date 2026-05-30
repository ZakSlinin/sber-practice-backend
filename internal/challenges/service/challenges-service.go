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

func (s *ChallengesService) CreateChallenge(ctx context.Context, req *models.CreateChallengeRequest) (*models.Challenge, error) {
	challenge := &models.Challenge{
		ID:          uuid.New(),
		WorkspaceID: req.WorkspaceID,
		Description: req.Description,
		Level:       req.Level,
		IsActive:    true,
		CreatedBy:   req.CreatedBy,
	}

	return s.repo.Create(ctx, challenge)
}
