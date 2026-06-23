package controllers

import (
	"github.com/AlvinSanudharma/project-management/models"
	"github.com/AlvinSanudharma/project-management/services"
	"github.com/AlvinSanudharma/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type BoardController struct {
	services services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{
		services: s,
	}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var err error

	board := new(models.Board)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}

	userID, err = uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}
	board.OwnerPublicID = userID

	if err := c.services.Create(board); err != nil {
		return utils.BadRequest(ctx, "Gagal menyimpan data", err.Error())
	}

	return utils.Success(ctx, "Board berhasil dibuat", board)
}

func (c *BoardController) UpdateBoard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	board := new(models.Board)

	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "ID tidak valid", err.Error())
	}

	existingBoard, err := c.services.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Board tidak ditemukan", err.Error())
	}

	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID
	board.OwnerID = existingBoard.OwnerID
	board.OwnerPublicID = existingBoard.OwnerPublicID
	board.CreatedAt = existingBoard.CreatedAt

	if err := c.services.Update(board); err != nil {
		return utils.BadRequest(ctx, "Gagal update board", err.Error())
	}

	return utils.Success(ctx, "Board berhasil diperbarui", board)
}
