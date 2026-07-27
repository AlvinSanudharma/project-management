package services

import "github.com/AlvinSanudharma/project-management/repositories"

type listService struct {
	listRepo  repositories.ListRepository
	boardRepo repositories.BoardRepository
}
