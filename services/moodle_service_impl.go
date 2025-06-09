package services

import (
	"encoding/json"
	"errors"
	"github.com/rizkycahyono97/moodle-api/config"
	"github.com/rizkycahyono97/moodle-api/model/web"
	"io"
	"net/http"
	"net/url"
)

type MoodleServiceImpl struct {
	client *http.Client
}

func NewMoodleService(client *http.Client) MoodleService {
	return &MoodleServiceImpl{
		client: client,
	}
}

func (s *MoodleServiceImpl) CheckStatus() (*web.MoodleStatusResponse, error) {
	// Load Env
	moodleURL := config.GetEnv("MOODLE_URL", "")
	token := config.GetEnv("MOODLE_TOKEN", "")

	if moodleURL == "" && token == "" {
		return nil, errors.New("Moodle URL or Token is Empty")
	}

	// Data Request
	form := url.Values{}
	form.Set("wstoken", token)
	form.Set("wsfunction", "core_webservice_get_site_info")
	form.Set("moodlewsrestformat", "json")

	// Kirim POST request ke Moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body)) // debug body asli moodle

	// Status Code Check
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to check Status: " + string(body))
	}

	// Unmarshal JSON ke struct MoodleStatusResponse
	var status web.MoodleStatusResponse
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}
	return &status, nil
}
