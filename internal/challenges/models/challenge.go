package models

import (
	"github.com/google/uuid"
)

type Challenge struct {
	ID          uuid.UUID
	WorkspaceID uuid.UUID
	Title       string
	Description string
	Level       int
	IsActive    bool
	CreatedBy   uuid.UUID
}

type CreateChallengeRequest struct {
	Title       string
	Description string
	Level       int
}
