package main

import (
	"log"

	"github.com/AlvinSanudharma/project-management/config"
	"github.com/AlvinSanudharma/project-management/controllers"
	"github.com/AlvinSanudharma/project-management/database/seed"
	"github.com/AlvinSanudharma/project-management/repositories"
	"github.com/AlvinSanudharma/project-management/routes"
	"github.com/AlvinSanudharma/project-management/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":" + port))
}
