package controllers

import (
	"github.com/AlvinSanudharma/project-management/models"
	"github.com/AlvinSanudharma/project-management/services"
	"github.com/AlvinSanudharma/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registrasi Gagal", err.Error())
	}

	var userResponse models.UserResponse
	_ = copier.Copy(&userResponse, user)

	return utils.Success(ctx, "Registrasi Success", userResponse)
}
