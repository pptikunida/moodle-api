package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rizkycahyono97/moodle-api/config"
	"github.com/rizkycahyono97/moodle-api/controllers"
	"github.com/rizkycahyono97/moodle-api/routes"
	"github.com/rizkycahyono97/moodle-api/services"
	"log"
	"net/http"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpClient := &http.Client{}

	// MoodleInject
	moodleService := services.NewMoodleService(httpClient)
	moodleController := controllers.NewMoodleController(moodleService)

	// Initialize Route
	r := gin.Default()

	// Router
	routes.SetupRouter(
		r,
		moodleController)

	// Port
	appPort := config.GetEnv("APP_PORT", "8080")

	// run application
	if err := r.Run(":" + appPort); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
