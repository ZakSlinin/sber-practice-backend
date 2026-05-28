package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID
	WorkspaceID  uuid.UUID
	Email        string
	PasswordHash string
	Name         string
	Role         string
	TotalPoints  int
}

type RegisterRequest struct {
	Action        string `json:"action"` // "create" | "join"
	WorkspaceName string `json:"workspace_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Name          string `json:"name"`
}

type LoginRequest struct {
	WorkspaceName string `json:"workspace_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
}
