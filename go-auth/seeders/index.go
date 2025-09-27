package seeders

import "log"

func RunAllSeeders() {
	log.Println("Running seeders...")

	SeedUsers()
	// Add more seeders here
	// SeedPosts()
	// SeedCategories()

	log.Println("Seeders completed!")
}
