package service

import (
	"context"
	"github.com/ZakSlinin/sber-practice-backend/internal/workspace/models"
	"github.com/ZakSlinin/sber-practice-backend/internal/workspace/repository"
	"github.com/google/uuid"
)

type WorkspaceService struct {
	repo repository.WorkspaceRepository
}

func NewWorkspaceService(repo repository.WorkspaceRepository) *WorkspaceService {
	return &WorkspaceService{repo: repo}
}

func (s *WorkspaceService) Create(ctx context.Context, name string) (*models.Workspace, error) {
	workspace := &models.Workspace{
		ID:   uuid.New(),
		Name: name,
	}
	return s.repo.Create(ctx, workspace)
}

func (s *WorkspaceService) GetByName(ctx context.Context, name string) (*models.Workspace, error) {
	return s.repo.GetByName(ctx, name)
}
