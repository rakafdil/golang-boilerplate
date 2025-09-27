package seeders

import (
	"go-auth/models"
	"go-auth/utils"
	"log"

	"github.com/google/uuid"
)

func SeedUsers() {
	// Check if users already exist
	var count int64
	models.DB.Model(&models.User{}).Count(&count)

	if count > 0 {
		log.Println("Users already exist, skipping seeder")
		return
	}

	// Create sample users
	users := []models.User{
		{
			ID:       uuid.New(),
			Name:     "Admin User",
			Email:    "admin@example.com",
			Password: "admin123",
			Role:     "admin",
		},
		{
			ID:       uuid.New(),
			Name:     "Regular User",
			Email:    "user@example.com",
			Password: "user123",
			Role:     "user",
		},
	}

	// Hash passwords and create users
	for i := range users {
		hashedPassword, err := utils.GenerateHashPassword(users[i].Password)
		if err != nil {
			log.Printf("Error hashing password for %s: %v", users[i].Email, err)
			continue
		}
		users[i].Password = hashedPassword

		if err := models.DB.Create(&users[i]).Error; err != nil {
			log.Printf("Error creating user %s: %v", users[i].Email, err)
		} else {
			log.Printf("Created user: %s", users[i].Email)
		}
	}
}
