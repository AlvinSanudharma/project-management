package models

import "github.com/google/uuid"

type User struct {
	InternalID int64
	PublicID   uuid.UUID
	Name       string
	Email      string
	Password   string
	Role       string
}
