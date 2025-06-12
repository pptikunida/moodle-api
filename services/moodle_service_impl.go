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

var ErrNotFound = errors.New("data dengan kriteria yang diberikan tidak ditemukan")

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

// file: services/moodle_service_impl.go

func (s *MoodleServiceImpl) CreateUser(req web.MoodleUserCreateRequest) ([]web.MoodleUserCreateResponse, error) {
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}

	// VALIDASI
	err = CheckMoodleDuplicateFields(s, map[string]string{
		"username": req.Username,
		"email":    req.Email,
		"idnumber": req.IdNumber,
	})
	if err != nil {
		return nil, err
	}

	form := helpers.NewMoodleForm(token, "core_user_create_users")
	// Mengisi form dengan semua data dari request
	form.Set("users[0][username]", req.Username)
	form.Set("users[0][password]", req.Password)
	form.Set("users[0][firstname]", req.Firstname)
	form.Set("users[0][lastname]", req.Lastname)
	form.Set("users[0][email]", req.Email)
	form.Set("users[0][idnumber]", req.IdNumber)
	form.Set("users[0][auth]", req.Auth)
	form.Set("users[0][city]", req.City)
	form.Set("users[0][country]", req.Country)
	form.Set("users[0][timezone]", req.Timezone)
	form.Set("users[0][description]", req.Description)
	form.Set("users[0][lang]", req.Lang)
	form.Set("users[0][calendartype]", req.CalendarType)

	for i, cf := range req.CustomFields {
		form.Set(fmt.Sprintf("users[0][customfields][%d][type]", i), cf.Type)
		form.Set(fmt.Sprintf("users[0][customfields][%d][value]", i), cf.Value)
	}
	for i, pref := range req.Preferences {
		form.Set(fmt.Sprintf("users[0][preferences][%d][type]", i), pref.Type)
		form.Set(fmt.Sprintf("users[0][preferences][%d][value]", i), pref.Value)
	}

	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, fmt.Errorf("gagal memanggil Moodle untuk CreateUser: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca body respons CreateUser: %w", err)
	}
	log.Printf("[CreateUser] Raw Response: %s", string(body))

	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return nil, &moodleErr
	}

	var result []web.MoodleUserCreateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("gagal unmarshal respons sukses CreateUser: %w", err)
	}

	return result, nil
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

	// moodle exeception
	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return nil, &moodleErr
	}

	// Jika tidak ada error, unmarshal ke slice of users
	var result []web.MoodleUserGetByFieldResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ErrNotFound
	}

	if len(result) == 0 {
		return nil, ErrNotFound
	}

	return result, nil
}

func (s *MoodleServiceImpl) UpdateUsers(req []web.MoodleUserUpdateRequest) error {
	// load moodle conf
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return err
	}
	form := helpers.NewMoodleForm(token, "core_user_update_users")

	for i, user := range req {
		if user.ID == 0 {
			return fmt.Errorf("Invalid user id: %d", user.ID)
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
		if user.Password != "" {
			form.Set(fmt.Sprintf("users[%d][password]", i), user.Password)
		}
		// ... add other optional fields in the same way (auth, city, etc.)

		// Handle custom fields
		for j, cf := range user.CustomFields {
			form.Set(fmt.Sprintf("users[%d][customfields][%d][type]", i, j), cf.Type)
			form.Set(fmt.Sprintf("users[%d][customfields][%d][value]", i, j), cf.Value)
		}

		// Handle preferences
		for j, pref := range user.Preferences {
			form.Set(fmt.Sprintf("users[%d][preferences][%d][type]", i, j), pref.Type)
			form.Set(fmt.Sprintf("users[%d][preferences][%d][value]", i, j), pref.Value)
		}
	}

	// send POST to moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return fmt.Errorf("error send request to Moodle: %w", err)
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	log.Printf("[DEBUG] Moodle update response (status %d): %s", resp.StatusCode, string(body))

	// check for moodle API if error
	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Message != "" {
		return &moodleErr
	}

	// check http method
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Moodle update returned invalid response (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}

func (s *MoodleServiceImpl) UserSync(req web.MoodleUserSyncRequest) error {
	// Validasi awal
	if req.Username == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" || req.Email == "" || req.NIM == "" {
		log.Println("[ERROR] UserSync: Permintaan tidak valid, ada field yang kosong")
		return fmt.Errorf("semua field wajib diisi: username, password, firstname, lastname, email, nim")
	}

	// Log permintaan awal
	log.Printf("[INFO] UserSync: Sinkronisasi pengguna dengan NIM '%s' dan username '%s'", req.NIM, req.Username)

	// Cek user di Moodle berdasarkan idnumber (NIM)
	existingByNIM, err := s.GetUserByField(web.MoodleUserGetByFieldRequest{
		Field:  "idnumber",
		Values: []string{req.NIM},
	})
	if err == nil && len(existingByNIM) > 0 {
		log.Printf("[INFO] UserSync: Field 'idnumber' dengan nilai '%s' sudah digunakan di Moodle. Sinkronisasi dilewati.", req.NIM)
		return fmt.Errorf("idnumber (NIM) '%s' sudah digunakan di Moodle", req.NIM)
	}

	existingByEmail, err := s.GetUserByField(web.MoodleUserGetByFieldRequest{
		Field:  "email",
		Values: []string{req.Email},
	})
	if err == nil && len(existingByEmail) > 0 {
		log.Printf("[INFO] UserSync: Field 'email' dengan nilai '%s' sudah digunakan di Moodle. Sinkronisasi dilewati.", req.Email)
		return fmt.Errorf("email '%s' sudah digunakan di Moodle", req.Email)
	}

	createUserReq := web.MoodleUserCreateRequest{
		Username:  req.Username,
		Password:  req.Password,
		Firstname: req.FirstName,
		Lastname:  req.LastName,
		Email:     req.Email,
		IdNumber:  req.NIM,
		Auth:      "manual",
		// Optional: bisa diisi default kalau tidak dikirim
		City:         "Indonesia",
		Country:      "ID",
		Timezone:     "Asia/Jakarta",
		Lang:         "en",
		CalendarType: "gregorian",
	}

	_, createErr := s.CreateUser(createUserReq)
	if createErr != nil {
		log.Printf("[ERROR] UserSync: Gagal membuat pengguna '%s' di Moodle. Error: %v", req.Username, createErr)
		return fmt.Errorf("gagal membuat user di moodle: %w", createErr)
	}

	log.Printf("[INFO] UserSync: Pengguna '%s' berhasil dibuat di Moodle.", req.Username)
	return nil
}

// function cek validasi jika ada duplikat
func CheckMoodleDuplicateField(svc MoodleService, field, value string) error {
	users, err := svc.GetUserByField(web.MoodleUserGetByFieldRequest{
		Field:  field,
		Values: []string{value},
	})
	// Tangani kasus tidak ditemukan → ini bukan error untuk duplikasi
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil // Tidak ditemukan berarti tidak duplikat
		}
		// Error selain ErrNotFound dianggap error asli
		return fmt.Errorf("gagal memeriksa field '%s' di Moodle: %w", field, err)
	}
	if len(users) > 0 {
		return fmt.Errorf("field '%s' dengan nilai '%s' sudah digunakan di Moodle", field, value)
	}
	return nil
}

// function cek validasi jika ada duplikat banyak
func CheckMoodleDuplicateFields(svc MoodleService, fields map[string]string) error {
	for field, value := range fields {
		if err := CheckMoodleDuplicateField(svc, field, value); err != nil {
			return err
		}
	}
	return nil
}
