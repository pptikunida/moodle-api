package validation

import (
	"errors"
	"fmt"

	"github.com/rizkycahyono97/moodle-api/contracts" // hanya interface!
	"github.com/rizkycahyono97/moodle-api/model/web"
)

var ErrNotFound = errors.New("data dengan kriteria yang diberikan tidak ditemukan") //

func CheckMoodleDuplicateField(svc contracts.MoodleUserGetter, field, value string) error {
	users, err := svc.CoreUserGetUsersByField(web.MoodleUserGetByFieldRequest{
		Field:  field,
		Values: []string{value},
	})
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		return fmt.Errorf("gagal memeriksa field '%s' di Moodle: %w", field, err)
	}
	if len(users) > 0 {
		return fmt.Errorf("field '%s' dengan nilai '%s' sudah digunakan di Moodle", field, value)
	}
	return nil
}

func CheckMoodleDuplicateFields(svc contracts.MoodleUserGetter, fields map[string]string) error {
	for field, value := range fields {
		if err := CheckMoodleDuplicateField(svc, field, value); err != nil {
			return err
		}
	}
	return nil
}
