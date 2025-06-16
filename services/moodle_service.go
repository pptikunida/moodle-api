package services

import (
	"github.com/rizkycahyono97/moodle-api/contracts"
	"github.com/rizkycahyono97/moodle-api/model/web"
)

type MoodleService interface {
	contracts.MoodleUserGetter // men-embed interface baru menghindari cyclic import
	CheckStatus() (*web.MoodleStatusResponse, error)
	CreateUser(req web.MoodleUserCreateRequest) ([]web.MoodleUserCreateResponse, error)
	GetUserByField(req web.MoodleUserGetByFieldRequest) ([]web.MoodleUserGetByFieldResponse, error)
	UpdateUsers(req []web.MoodleUserUpdateRequest) error
	UserSync(req web.MoodleUserSyncRequest) error
	AssignRole(req web.MoodleRoleAssignRequest) error
	CreateCourse(req web.MoodleCreateCourseRequest) ([]web.MoodleCreateCourseResponse, error)
	EnrollManualEnrolUsers(req web.MoodleManualEnrollRequest) error
}
