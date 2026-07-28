package services

import (
	"github.com/AlvinSanudharma/project-management/models"
	"github.com/AlvinSanudharma/project-management/repositories"
	"github.com/google/uuid"
)

type listService struct {
	listRepo         repositories.ListRepository
	boardRepo        repositories.BoardRepository
	listPositionRepo repositories.ListPositionRepository
}

type ListWithOrder struct {
	Positions []uuid.UUID
	Lists     []models.List
}

type ListService interface {
	GetByBoardID(boardPublicID string) (*ListWithOrder, error)
	GetByID(id uint) (*models.List, error)
	GetByPublicID(publicID string) (*models.List, error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(list *models.List) error
	UpdatePositions(boardPublicID string, position []uuid.UUID) error
}
