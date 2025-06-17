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
		api.GET("/moodle/site-info", moodleController.CoreWebserviceGetSiteInfo)
		api.POST("/moodle/users", moodleController.CoreUserCreateUsers)
		api.POST("/moodle/users/lookup-by-field", moodleController.CoreUserGetUsersByField)
		api.POST("/moodle/users/update", moodleController.CoreUserUpdateUsers)
		api.POST("/moodle/users/sync", moodleController.UserSync)
		api.POST("/moodle/roles/assign", moodleController.CoreRoleAssignRoles)
		api.POST("/moodle/courses/course", moodleController.CoreCourseCreateCourses)
		api.POST("/moodle/courses/enrol/manual", moodleController.EnrolManualEnrolUsers)
		api.POST("/moodle/courses/create-with-enrolment", moodleController.CreateCourseWithEnrollUser)
	}
}
