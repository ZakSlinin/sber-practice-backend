package repository

import (
	"context"
	"github.com/ZakSlinin/sber-practice-backend/internal/challenges/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChallengesRepo interface {
	Create(ctx context.Context, challenge *models.Challenge) (*models.Challenge, error)
	GetByWorkspace(ctx context.Context, workspaceID uuid.UUID) ([]*models.Challenge, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Challenge, error)
}

type PostgresChallengesRepo struct {
	db *gorm.DB
}

func NewPostgresChallengesRepo(db *gorm.DB) *PostgresChallengesRepo {
	return &PostgresChallengesRepo{db: db}
}

func (r *PostgresChallengesRepo) Create(ctx context.Context, challenge *models.Challenge) (*models.Challenge, error) {
	result := r.db.WithContext(ctx).Create(challenge)

	if result.Error != nil {
		return nil, result.Error
	}

	return challenge, nil
}

func (r *PostgresChallengesRepo) GetByWorkspace(ctx context.Context, workspaceID uuid.UUID) ([]*models.Challenge, error) {
	var challenges []*models.Challenge
	result := r.db.WithContext(ctx).
		Where("workspace_id = ?", workspaceID).
		Find(&challenges)
	if result.Error != nil {
		return nil, result.Error
	}
	return challenges, nil
}

func (r *PostgresChallengesRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Challenge, error) {
	var challenge models.Challenge
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&challenge)
	if result.Error != nil {
		return nil, result.Error
	}
	return &challenge, nil
}
