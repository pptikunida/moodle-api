package contracts

import "github.com/rizkycahyono97/moodle-api/model/web"

type MoodleUserGetter interface {
	CoreUserGetUsersByField(req web.MoodleUserGetByFieldRequest) ([]web.MoodleUserGetByFieldResponse, error)
}
