package service

import (
	"context"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/models"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/repository"
	models2 "github.com/ZakSlinin/sber-practice-backend/internal/workspace/models"
	repository2 "github.com/ZakSlinin/sber-practice-backend/internal/workspace/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
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

type ErrorMessage struct {
	Error     string                 `json:"error"`
	Message   string                 `json:"message"`
	Timestamp *timestamppb.Timestamp `json:"timestamp"`
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

func (s *AuthService) GetByEmailAndWorkspace(ctx context.Context, email string, workspaceID uuid.UUID) (*models.User, *ErrorMessage) {
	user, err := s.repo.GetByEmailAndWorkspace(ctx, email, workspaceID)

	if err != nil {
		if err.Error() == "user not found" {
			return nil, &ErrorMessage{Error: "USER_NOT_FOUND", Message: "User not found", Timestamp: timestamppb.New(time.Now())}
		}
		return nil, &ErrorMessage{
			Error:     "INTERNAL_ERROR",
			Message:   err.Error(),
			Timestamp: timestamppb.Now(),
		}
	}

	return user, nil
}
