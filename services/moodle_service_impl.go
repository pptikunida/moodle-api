package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rizkycahyono97/moodle-api/model/web"
	"github.com/rizkycahyono97/moodle-api/utils/helpers"
	"io"
	"log"
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

func (s *MoodleServiceImpl) CreateUser(req web.MoodleUserCreateRequest) (*web.MoodleUserCreateResponse, error) {

	// Load Env & Moodle Request
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}
	form := helpers.NewMoodleForm(token, "core_user_create_users")

	// Set array-style form fields sesuai format Moodle
	form.Set("users[0][username]", req.Username)
	form.Set("users[0][auth]", req.Auth)
	form.Set("users[0][password]", req.Password)
	form.Set("users[0][firstname]", req.Firstname)
	form.Set("users[0][lastname]", req.Lastname)
	form.Set("users[0][email]", req.Email)
	if req.City != "" {
		form.Set("users[0][city]", req.City)
	}
	if req.Country != "" {
		form.Set("users[0][country]", req.Country)
	}
	if req.Timezone != "" {
		form.Set("users[0][timezone]", req.Timezone)
	}
	if req.Description != "" {
		form.Set("users[0][description]", req.Description)
	}
	if req.IdNumber != "" {
		form.Set("users[0][idnumber]", req.IdNumber)
	}
	if req.Lang != "" {
		form.Set("users[0][lang]", req.Lang)
	}
	if req.CalendarType != "" {
		form.Set("users[0][calendartype]", req.CalendarType)
	}
	for i, field := range req.CustomFields {
		form.Set(fmt.Sprintf("users[0][customfields][%d][type]", i), field.Type)
		form.Set(fmt.Sprintf("users[0][customfields][%d][value]", i), field.Value)
	}
	for i, pref := range req.Preferences {
		form.Set(fmt.Sprintf("users[0][preferences][%d][type]", i), pref.Type)
		form.Set(fmt.Sprintf("users[0][preferences][%d][value]", i), pref.Value)
	}

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

	log.Println("[CreateUser] Raw Response:", string(body)) // log

	// Check for moodle error
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to create User: " + string(body))
	}

	// Parse Moodle success
	var result []web.MoodleUserCreateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 || result[0].ID == 0 {
		log.Println("[CreateUser] Warning: Moodle returned empty or invalid user")
		return nil, errors.New("No user returned or user invalid")
	}

	return &result[0], nil
}
