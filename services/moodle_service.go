package services

import "github.com/rizkycahyono97/moodle-api/model/web"

type MoodleService interface {
	CheckStatus() (*web.MoodleStatusResponse, error)
	CreateUser(req web.MoodleUserCreateRequest) (*web.MoodleUserCreateResponse, error)
	GetUserByField(req web.MoodleUserGetByFieldRequest) ([]web.MoodleUserGetByFieldResponse, error)
	UpdateUsers(req []web.MoodleUserUpdateRequest) error
}
