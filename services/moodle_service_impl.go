package services

import (
	"encoding/json"
	"errors"
	"github.com/rizkycahyono97/moodle-api/model/web"
	"github.com/rizkycahyono97/moodle-api/utils/helpers"
	"io"
	"net/http"
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
	// Load Env & Moodle Request
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}
	form := helpers.NewMoodleForm(token, "core_webservice_get_site_info")

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

func (s *MoodleServiceImpl) CreateUser(req web.MoodleUserCreateRequest) ([]web.MoodleUserCreateResponse, error) {
	// CreateUser bisa lebih dari satu
	users := []web.MoodleUserCreateRequest{req}

	// JSON-Encoded
	usersJson, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	// Load Env & Moodle Request
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}
	form := helpers.NewMoodleForm(token, "core_user_create_users")
	form.Set("users", string(usersJson))

	// POST req to moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check for moodle error
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to create User: " + string(body))
	}

	// Parse Moodle success
	var result []web.MoodleUserCreateResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
