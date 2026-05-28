package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByEmailAndWorkspace(ctx context.Context, email string, workspaceID uuid.UUID) (*models.User, error)
}

type PostgresAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{db: db}
}

func (r *PostgresAuthRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result := r.db.WithContext(ctx).Create(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *PostgresAuthRepository) GetByEmailAndWorkspace(ctx context.Context, email string, workspaceID uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("email = ? AND workspace_id = ?", email, workspaceID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}
