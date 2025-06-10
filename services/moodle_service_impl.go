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

func (s *MoodleServiceImpl) GetUserByField(req web.MoodleUserGetByFieldRequest) ([]web.MoodleUserGetByFieldResponse, error) {
	// validasi field yang diperbolehkan
	allowedFields := map[string]bool{
		"id":       true,
		"idnumber": true,
		"username": true,
		"email":    true,
	}
	if !allowedFields[req.Field] {
		return nil, fmt.Errorf("Invalid field: %s", req.Field)
	}

	// Load env & moodle req
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}
	form := helpers.NewMoodleForm(token, "core_user_get_users_by_field")
	form.Set("field", req.Field)

	// Set parameter values[0], values[1]
	for i, v := range req.Values {
		form.Set(fmt.Sprintf("values[%d]", i), v)
	}

	// kirim POST ke moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, fmt.Errorf("error calling Moodle: %w", err)
	}
	defer resp.Body.Close()

	// Baca response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	//debug
	log.Printf("[DEBUG] Request Field: %s", req.Field)
	log.Printf("[DEBUG] Request Values: %v", req.Values)
	log.Printf("[DEBUG] Encoded Form: %v", form.Encode())
	log.Printf("[DEBUG] Raw Moodle Response: %s", string(body))

	// Cek apakah response mengandung error Moodle
	var maybeError map[string]interface{}
	if err := json.Unmarshal(body, &maybeError); err == nil {
		if _, exists := maybeError["exception"]; exists {
			return nil, fmt.Errorf("error from Moodle: %s", string(body))
		}
	}

	// Jika tidak ada error, unmarshal ke slice of users
	var result []web.MoodleUserGetByFieldResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Moodle user response: %w", err)
	}

	// validasi jika tidak ada user yang ditemukan
	if len(result) == 0 {
		return nil, errors.New("No user returned or invalid field")
	}

	return result, nil
}

func (s *MoodleServiceImpl) UpdateUsers(req []web.MoodleUserUpdateRequest) ([]web.MoodleUserUpdateResponse, error) {
	// Load config
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}
	form := helpers.NewMoodleForm(token, "core_user_update_users")

	// konversi array struct ke map string
	for i, user := range req {
		if user.ID == 0 {
			return nil, fmt.Errorf("user ID is required for update")
		}
		form.Set(fmt.Sprintf("users[%d][id]", i), fmt.Sprintf("%d", user.ID))
		if user.Username != "" {
			form.Set(fmt.Sprintf("users[%d][username]", i), user.Username)
		}
		if user.Email != "" {
			form.Set(fmt.Sprintf("users[%d][email]", i), user.Email)
		}
		if user.Firstname != "" {
			form.Set(fmt.Sprintf("users[%d][firstname]", i), user.Firstname)
		}
		if user.Lastname != "" {
			form.Set(fmt.Sprintf("users[%d][lastname]", i), user.Lastname)
		}
		if user.Auth != "" {
			form.Set(fmt.Sprintf("users[%d][auth]", i), user.Auth)
		}
		//if user.Suspended != nil {
		//	form.Set(fmt.Sprintf("users[%d][suspended]", i), fmt.Sprintf("%d", *user.Suspended))
		//}
		if user.Password != "" {
			form.Set(fmt.Sprintf("users[%d][password]", i), user.Password)
		}
		if user.City != "" {
			form.Set(fmt.Sprintf("users[%d][city]", i), user.City)
		}
		if user.Country != "" {
			form.Set(fmt.Sprintf("users[%d][country]", i), user.Country)
		}
		if user.Description != "" {
			form.Set(fmt.Sprintf("users[%d][description]", i), user.Description)
		}
		//if user.MailDisplay != nil {
		//	form.Set(fmt.Sprintf("users[%d][maildisplay]", i), fmt.Sprintf("%d", *user.MailDisplay))
		//}
		if user.Timezone != "" {
			form.Set(fmt.Sprintf("users[%d][timezone]", i), user.Timezone)
		}
		if user.Lang != "" {
			form.Set(fmt.Sprintf("users[%d][lang]", i), user.Lang)
		}
		//custom field lain
		for j, cf := range user.CustomFields {
			form.Set(fmt.Sprintf("users[%d][customfields][%d][type]", i, j), cf.Type)
			form.Set(fmt.Sprintf("users[%d][customfields][%d][value]", i, j), cf.Value)
		}
		// Preferences
		for j, pref := range user.Preferences {
			form.Set(fmt.Sprintf("users[%d][preferences][%d][type]", i, j), pref.Type)
			form.Set(fmt.Sprintf("users[%d][preferences][%d][value]", i, j), pref.Value)
		}
	}

	// kirim req ke moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// baca resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	// debug
	log.Printf("[DEBUG] Moodle update response: %s", string(body))

	// cek kemungkinan error dari moodle
	var maybeError map[string]interface{}
	if er := json.Unmarshal(body, &maybeError); er != nil {
		if _, exists := maybeError["exception"]; exists {
			return nil, fmt.Errorf("error from Moodle: %s", string(body))
		}
	}

	// decode response
	var result web.MoodleUserUpdateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Moodle user response: %w", err)
	}
	return []web.MoodleUserUpdateResponse{result}, nil
}
