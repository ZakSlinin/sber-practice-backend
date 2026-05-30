package handler

import (
	"github.com/ZakSlinin/sber-practice-backend/internal/challenges/models"
	"github.com/ZakSlinin/sber-practice-backend/internal/challenges/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ChallengesHandler struct {
	service *service.ChallengesService
}

func NewChallengesHandler(service *service.ChallengesService) *ChallengesHandler {
	return &ChallengesHandler{service: service}
}

func (s *ChallengesHandler) CreateChallenge(g *gin.Context) {
	var req models.CreateChallengeRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workspaceIDStr := g.MustGet("workspace_id").(string)
	workspaceID, err := uuid.Parse(workspaceIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid workspace_id"})
		return
	}

	userIDStr := g.MustGet("user_id").(string)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	result, err := s.service.CreateChallenge(g.Request.Context(), &req, workspaceID, userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, result)
}
