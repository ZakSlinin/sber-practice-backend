package service

import (
	"context"
	"fmt"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/models"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/repository"
	models2 "github.com/ZakSlinin/sber-practice-backend/internal/workspace/models"
	repository2 "github.com/ZakSlinin/sber-practice-backend/internal/workspace/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService struct {
	repo          repository.AuthRepository
	workspaceRepo repository2.WorkspaceRepository
}

func NewAuthService(repo *repository.PostgresAuthRepository, workspace *repository2.PostgresWorkspaceRepository) *AuthService {
	return &AuthService{
		repo:          repo,
		workspaceRepo: workspace,
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {

	var workspace *models2.Workspace
	var err error

	if req.Action == "create" {
		workspace, err = s.workspaceRepo.Create(ctx, &models2.Workspace{
			ID:   uuid.New(),
			Name: req.WorkspaceName,
		})
	} else {
		workspace, err = s.workspaceRepo.GetByName(ctx, req.WorkspaceName)
	}
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role := "pm"
	if req.Action == "create" {
		role = "admin"
	}

	user := &models.User{
		ID:           uuid.New(),
		WorkspaceID:  workspace.ID,
		Email:        req.Email,
		Name:         req.Name,
		PasswordHash: string(hash),
		Role:         role,
	}

	return s.repo.Create(ctx, user)
}

func (s *AuthService) GetByEmailAndWorkspace(ctx context.Context, email string, workspaceID uuid.UUID) (*models.User, error) {
	user, err := s.repo.GetByEmailAndWorkspace(ctx, email, workspaceID)

	if err != nil {
		if err.Error() == "user not found" {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (string, error) {
	workspace, err := s.workspaceRepo.GetByName(ctx, req.WorkspaceName)
	if err != nil {
		return "", fmt.Errorf("workspace not found")
	}

	user, err := s.repo.GetByEmailAndWorkspace(ctx, req.Email, workspace.ID)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", fmt.Errorf("could not generate token")
	}

	return token, nil
}

func generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":      user.ID,
		"workspace_id": user.WorkspaceID,
		"role":         user.Role,
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
