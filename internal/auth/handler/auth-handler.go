package handler

import (
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/models"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(g *gin.Context) {
	var req models.RegisterRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.authService.Register(g.Request.Context(), &req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, models.RegisterResponse{
		ID:          result.ID,
		WorkspaceID: result.WorkspaceID,
		Email:       result.Email,
		Name:        result.Name,
		Role:        result.Role,
	})

}

func (h *AuthHandler) Login(g *gin.Context) {
	var req models.LoginRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(g.Request.Context(), &req)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"token": token})
}
