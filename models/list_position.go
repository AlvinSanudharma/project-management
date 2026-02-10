package models

import "github.com/google/uuid"

type ListPosition struct {
	InternalId int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicId   uuid.UUID `json:"public_id" db:"public_id" gorm:"public_id"`
	BoardId    int64     `json:"board_internal_id" db:"board_internal_id" gorm:"board_internal_id"`
	// ListOrder
}
