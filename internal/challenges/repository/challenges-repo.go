package repository

import (
	"context"
	"github.com/ZakSlinin/sber-practice-backend/internal/challenges/models"
	"gorm.io/gorm"
)

type ChallengesRepo interface {
	Create(ctx context.Context, challenge *models.Challenge) (*models.Challenge, error)
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
