package services

import "github.com/rizkycahyono97/moodle-api/model/web"

type MoodleService interface {
	CheckStatus() (*web.MoodleStatusResponse, error)
	CreateUser()
}
