package repository

import (
	"context"
	"fmt"
	"github.com/ZakSlinin/sber-practice-backend/internal/workspace/models"
	"gorm.io/gorm"
)

type WorkspaceRepository interface {
	Create(ctx context.Context, workspace models.Workspace) (*models.Workspace, error)
	GetByName(ctx context.Context, name string) (*models.Workspace, error)
}

type PostgresWorkspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) *PostgresWorkspaceRepository {
	return &PostgresWorkspaceRepository{db: db}
}

func (r *PostgresWorkspaceRepository) Create(ctx context.Context, workspace *models.Workspace) (*models.Workspace, error) {
	result := r.db.WithContext(ctx).Create(workspace)

	if result.Error != nil {
		return nil, result.Error
	}
	return workspace, nil
}

func (r *PostgresWorkspaceRepository) GetByName(ctx context.Context, name string) (*models.Workspace, error) {
	var workspace models.Workspace
	result := r.db.WithContext(ctx).Where(&workspace, "name = ?", name)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("workspace %s not found", name)
		}
		return nil, result.Error
	}

	return &workspace, nil
}
