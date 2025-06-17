package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/moodle-api/controllers"
	"github.com/rizkycahyono97/moodle-api/middleware"
)

func SetupRouter(
	r *gin.Engine,
	moodleController *controllers.MoodleController) {

	// protected routes
	protected := r.Group("/api")
	protected.Use(middleware.ApiKeyAuthMiddleware())
	{
		protected.GET("/moodle/site-info", moodleController.CoreWebserviceGetSiteInfo)
		protected.POST("/moodle/users", moodleController.CoreUserCreateUsers)
		protected.POST("/moodle/users/lookup-by-field", moodleController.CoreUserGetUsersByField)
		protected.POST("/moodle/users/update", moodleController.CoreUserUpdateUsers)
		protected.POST("/moodle/users/sync", moodleController.UserSync)
		protected.POST("/moodle/roles/assign", moodleController.CoreRoleAssignRoles)
		protected.POST("/moodle/courses/course", moodleController.CoreCourseCreateCourses)
		protected.POST("/moodle/courses/enrol/manual", moodleController.EnrolManualEnrolUsers)
		protected.POST("/moodle/courses/create-with-enrolment", moodleController.CreateCourseWithEnrollUser)
	}
}
