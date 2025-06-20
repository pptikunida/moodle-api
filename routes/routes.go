package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/moodle-api/controllers"
	"github.com/rizkycahyono97/moodle-api/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	r *gin.Engine,
	moodleController *controllers.MoodleController) {

	//swagger routes
	swaggerURL := ginSwagger.URL("/apispec.json")

	// protected routes
	protected := r.Group("/api/v1")
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
		protected.POST("/moodle/categories", moodleController.CoreCourseCreateCategories)
		protected.PUT("/moodle/categories", moodleController.CoreCourseUpdateCategories)
		protected.DELETE("/moodle/categories", moodleController.CoreCourseDeleteCategories)

		//swagger
		protected.GET("/apispec.json", moodleController.ServeSwaggerSpec)
		protected.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
	}
}
