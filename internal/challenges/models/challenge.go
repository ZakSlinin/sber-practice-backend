package models

import "github.com/google/uuid"

type Challenge struct {
	ID          uuid.UUID `gorm:"column:id"`
	WorkspaceID uuid.UUID `gorm:"column:workspace_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Level       string    `gorm:"column:level"` // "light" | "medium" | "hard"
	IsActive    bool      `gorm:"column:is_active"`
	CreatedBy   uuid.UUID `gorm:"column:created_by"`
}

type CreateChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       string `json:"level"` // "light" | "medium" | "hard"
}

type GetChallengesResponse struct {
	Challenges []*Challenge `json:"challenges"`
}
