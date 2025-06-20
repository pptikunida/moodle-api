package services

import (
	"github.com/rizkycahyono97/moodle-api/contracts"
	"github.com/rizkycahyono97/moodle-api/model/web"
)

type MoodleService interface {
	contracts.MoodleUserGetter // men-embed interface baru menghindari cyclic import
	CoreWebserviceGetSiteInfo() (*web.MoodleStatusResponse, error)
	CoreUserCreateUsers(req web.MoodleUserCreateRequest) ([]web.MoodleUserCreateResponse, error)
	CoreUserGetUsersByField(req web.MoodleUserGetByFieldRequest) ([]web.MoodleUserGetByFieldResponse, error)
	CoreUserUpdateUsers(req []web.MoodleUserUpdateRequest) error
	UserSync(req web.MoodleUserSyncRequest) error
	CoreRoleAssignRoles(req web.MoodleRoleAssignRequest) error
	CoreCourseCreateCourses(req web.MoodleCoreCourseCreateCoursesRequest) ([]web.MoodleCoreCourseCreateCoursesResponse, error)
	EnrolManualEnrolUsers(req web.MoodleManualEnrollRequest) error
	CreateCourseWithEnrollUser(req web.MoodleCreateCourseWithEnrollUserRequest) (*web.MoodleCreateCourseWithEnrollUserResponse, error)
	CoreCourseCreateCategories(req web.MoodleCreateCategoriesRequest) ([]web.MoodleCreateCategoriesResponse, error)
	CoreCourseUpdateCategories(req web.MoodleUpdateCategoriesRequest) error
}
