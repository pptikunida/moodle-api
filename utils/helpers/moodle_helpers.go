package helpers

import (
	"errors"
	"github.com/rizkycahyono97/moodle-api/config"
	"net/url"
)

func GetMoodleConfig() (string, string, error) {
	moodleURL := config.GetEnv("MOODLE_URL", "")
	token := config.GetEnv("MOODLE_TOKEN", "")

	if moodleURL == "" && token == "" {
		return "", "", errors.New("Moodle URL or token cannot be empty.")
	}

	return moodleURL, token, nil
}

func NewMoodleForm(token, function string) url.Values {
	form := url.Values{}
	form.Set("wstoken", token)
	form.Set("wsfunction", function)
	form.Set("moodlewsrestformat", "json")
	return form
}
