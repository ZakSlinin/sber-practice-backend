package models

import "github.com/google/uuid"

type Workspace struct {
	ID   uuid.UUID
	Name string
}
