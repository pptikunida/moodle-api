package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/moodle-api/controllers"
)

func SetupRouter(
	r *gin.Engine,
	moodleController *controllers.MoodleController) {

	// routes
	api := r.Group("/api")
	{
		api.GET("/users/status", moodleController.CheckStatus)
		api.POST("/users", moodleController.CreateUser)
		api.POST("/users/lookup", moodleController.GetUserByField)
		api.POST("/users/update", moodleController.UpdateUser)
		api.POST("/users/sync", moodleController.UserSync)
		api.POST("/users/assign-role", moodleController.AssignRole)
	}
}
