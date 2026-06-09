package seed

import (
	"log"

	"github.com/AlvinSanudharma/project-management/config"
	"github.com/AlvinSanudharma/project-management/models"
	"github.com/AlvinSanudharma/project-management/utils"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("password")

	admin := models.User{
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("Admin seeded successfully")
	}
}
