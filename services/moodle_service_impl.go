package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/rizkycahyono97/moodle-api/model/web"
	"github.com/rizkycahyono97/moodle-api/utils/helpers"
	"github.com/rizkycahyono97/moodle-api/utils/validation"
)

type MoodleServiceImpl struct {
	client *http.Client
}

func NewMoodleService(client *http.Client) MoodleService {
	return &MoodleServiceImpl{
		client: client,
	}
}

func (s *MoodleServiceImpl) CoreWebserviceGetSiteInfo() (*web.MoodleStatusResponse, error) {
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

func (s *MoodleServiceImpl) CoreUserCreateUsers(req web.MoodleUserCreateRequest) ([]web.MoodleUserCreateResponse, error) {
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, err
	}

	// VALIDASI
	err = validation.CheckMoodleDuplicateFields(s, map[string]string{
		"username": req.Users[0].Username,
		"email":    req.Users[0].Email,
		"idnumber": req.Users[0].IdNumber,
	})
	if err != nil {
		return nil, err
	}

	form := helpers.NewMoodleForm(token, "core_user_create_users")
	for i, user := range req.Users {
		form.Set(fmt.Sprintf("users[%d][username]", i), user.Username)
		form.Set(fmt.Sprintf("users[%d][password]", i), user.Password)
		form.Set(fmt.Sprintf("users[%d][firstname]", i), user.Firstname)
		form.Set(fmt.Sprintf("users[%d][lastname]", i), user.Lastname)
		form.Set(fmt.Sprintf("users[%d][email]", i), user.Email)
		form.Set(fmt.Sprintf("users[%d][idnumber]", i), user.IdNumber)
		form.Set(fmt.Sprintf("users[%d][auth]", i), user.Auth)
		form.Set(fmt.Sprintf("users[%d][city]", i), user.City)
		form.Set(fmt.Sprintf("users[%d][country]", i), user.Country)
		form.Set(fmt.Sprintf("users[%d][timezone]", i), user.Timezone)
		form.Set(fmt.Sprintf("users[%d][description]", i), user.Description)
		form.Set(fmt.Sprintf("users[%d][lang]", i), user.Lang)
		form.Set(fmt.Sprintf("users[%d][calendartype]", i), user.CalendarType)

		for j, cf := range user.CustomFields {
			form.Set(fmt.Sprintf("users[%d][customfields][%d][type]", i, j), cf.Type)
			form.Set(fmt.Sprintf("users[%d][customfields][%d][value]", i, j), cf.Value)
		}
		for j, pref := range user.Preferences {
			form.Set(fmt.Sprintf("users[%d][preferences][%d][type]", i, j), pref.Type)
			form.Set(fmt.Sprintf("users[%d][preferences][%d][value]", i, j), pref.Value)
		}
	}

	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, fmt.Errorf("gagal memanggil Moodle untuk CoreUserCreateUsers: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca body respons CoreUserCreateUsers: %w", err)
	}
	log.Printf("[CoreUserCreateUsers] Raw Response: %s", string(body))

	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return nil, &moodleErr
	}

	var result []web.MoodleUserCreateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("gagal unmarshal respons sukses CoreUserCreateUsers: %w", err)
	}

	return result, nil
}

func (s *MoodleServiceImpl) CoreUserGetUsersByField(req web.MoodleUserGetByFieldRequest) ([]web.MoodleUserGetByFieldResponse, error) {
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
		return nil, validation.ErrNotFound
	}

	if len(result) == 0 {
		return nil, validation.ErrNotFound
	}

	return result, nil
}

func (s *MoodleServiceImpl) CoreUserUpdateUsers(req []web.MoodleUserUpdateRequest) error {
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

// function userSync untuk login siakad, jika tidak ada buatkan user di moodlenya
func (s *MoodleServiceImpl) UserSync(req web.MoodleUserSyncRequest) error {
	// Validasi awal
	if req.Username == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" || req.Email == "" || req.NIM == "" {
		log.Println("[ERROR] UserSync: Permintaan tidak valid, ada field yang kosong")
		return fmt.Errorf("semua field wajib diisi: username, password, firstname, lastname, email, nim")
	}

	// Log permintaan awal
	log.Printf("[INFO] UserSync: Sinkronisasi pengguna dengan NIM '%s' dan username '%s'", req.NIM, req.Username)

	// validasi jika duplikat
	err := validation.CheckMoodleDuplicateFields(s, map[string]string{
		"idnumber": req.NIM,
		"email":    req.Email,
	})
	if err != nil {
		log.Printf("[INFO] UserSync: Duplikasi ditemukan. Error: %v", err)
		return err
	}

	userData := web.MoodleUserCreateData{
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

	//merubah struktur request
	createUserReq := web.MoodleUserCreateRequest{
		Users: []web.MoodleUserCreateData{userData},
	}

	createdUsers, createErr := s.CoreUserCreateUsers(createUserReq)
	if createErr != nil {
		log.Printf("[ERROR] UserSync: Gagal membuat pengguna '%s' di Moodle. Error: %v", req.Username, createErr)
		return fmt.Errorf("gagal membuat user di moodle: %w", createErr)
	}

	// ambil user id yang baru dibuat
	newUserID := createdUsers[0].ID

	// assign Role
	assignReq := web.MoodleRoleAssignRequest{
		Assignments: []web.MoodleRoleAssigment{
			{
				RoleID: req.RoleID,
				UserID: newUserID,
				//ContextLevel: "system",
				ContextID: 1,
			},
		},
	}

	if err := s.CoreRoleAssignRoles(assignReq); err != nil {
		log.Printf("[WARN] UserSync: Gagal assign role ke user '%s'. Error: %v", req.Username, err)
	} else {
		log.Printf("[INFO] UserSync: Berhasil assign role ID %d ke user '%s'", req.RoleID, req.Username)
	}

	log.Printf("[INFO] UserSync: Pengguna '%s' berhasil dibuat di Moodle.", req.Username)
	return nil
}

func (s *MoodleServiceImpl) CoreRoleAssignRoles(req web.MoodleRoleAssignRequest) error {
	moodleUrl, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil
	}

	form := helpers.NewMoodleForm(token, "core_role_assign_roles")

	// loop melalui assignments
	for i, a := range req.Assignments {
		form.Set(fmt.Sprintf("assignments[%d][roleid]", i), strconv.Itoa(a.RoleID))
		form.Set(fmt.Sprintf("assignments[%d][userid]", i), strconv.Itoa(a.UserID))
		if a.ContextID != 0 {
			form.Set(fmt.Sprintf("assignments[%d][contextid]", i), strconv.Itoa(a.ContextID))
		}
		if a.ContextLevel != "" {
			form.Set(fmt.Sprintf("assignments[%d][contextlevel]", i), a.ContextLevel)
		}
		if a.InstanceID != 0 {
			form.Set(fmt.Sprintf("assignments[%d][instanceid]", i), strconv.Itoa(a.InstanceID))
		}
	}

	log.Printf("[DEBUG] CoreRoleAssignRoles Service: Form data yang akan dikirim:\n%s", form.Encode())

	// kirim req POST ke moodle
	resp, err := s.client.PostForm(moodleUrl, form)
	if err != nil {
		return fmt.Errorf("error send request to Moodle: %w", err)
	}

	// baca body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	log.Printf("[DEBUG] CoreRoleAssignRoles Service: Raw response dari Moodle: %s", string(body))

	// cek response
	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return &moodleErr
	}

	return nil

}

func (s *MoodleServiceImpl) CoreCourseCreateCourses(req web.MoodleCoreCourseCreateCoursesRequest) ([]web.MoodleCoreCourseCreateCoursesResponse, error) {
	// load moodle conf
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		log.Printf("[ERROR] CoreCourseCreateCourses: gagal mendapatkan konfigurasi Moodle: %v", err)
		return nil, err
	}

	// req ke moodle
	form := helpers.NewMoodleForm(token, "core_course_create_courses")

	// Encode ke format moodle
	for i, course := range req.Courses {
		prefix := fmt.Sprintf("courses[%d]", i)
		form.Set(prefix+"[fullname]", course.FullName)
		form.Set(prefix+"[shortname]", course.ShortName)
		form.Set(prefix+"[categoryid]", fmt.Sprintf("%d", course.CategoryID))

		if course.IDNumber != "" {
			form.Set(fmt.Sprintf("courses[%d][idnumber]", i), course.IDNumber)
		}
		if course.Summary != "" {
			form.Set(prefix+"[summary]", course.Summary)
		}
		if course.SummaryFormat != 0 {
			form.Set(prefix+"[summaryformat]", fmt.Sprintf("%d", course.SummaryFormat))
		}
		if course.Format != "" {
			form.Set(prefix+"[format]", course.Format)
		}
		if course.Lang != "" {
			form.Set(prefix+"[lang]", course.Lang)
		}
		if course.Visible != 0 {
			form.Set(prefix+"[visible]", fmt.Sprintf("%d", course.Visible))
		}
		//custom field
		for j, field := range course.CustomFields {
			fieldPrefix := fmt.Sprintf("customfields[%d]", j)
			form.Set(fieldPrefix+"[shortname]", field.ShortName)
			form.Set(fieldPrefix+"[value]", field.Value)
		}
	}

	// kirim post ke moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		log.Printf("[ERROR] CoreCourseCreateCourses: gagal melakukan request ke Moodle: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] CoreCourseCreateCourses: gagal membaca body: %v", err)
		return nil, err
	}
	log.Printf("[DEBUG] CoreCourseCreateCourses Raw Response: %s", string(body))

	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return nil, &moodleErr
	}

	// jika tidak error
	var moodleResp []web.MoodleCoreCourseCreateCoursesResponse
	if err := json.Unmarshal(body, &moodleResp); err != nil {
		log.Printf("[ERROR] CoreCourseCreateCourses: gagal decode response sukses: %v", err)
		return nil, fmt.Errorf("gagal unmarshal respons sukses CoreCourseCreateCourses: %w", err)
	}

	// Enroll users

	return moodleResp, nil
}

func (s *MoodleServiceImpl) EnrolManualEnrolUsers(req web.MoodleManualEnrollRequest) error {
	// load config moodle
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan konfigurasi Moodle: %w", err)
	}

	// siapkan form
	form := helpers.NewMoodleForm(token, "enrol_manual_enrol_users")
	for i, enrol := range req.Enrolments {
		form.Set(fmt.Sprintf("enrolments[%d][roleid]", i), strconv.Itoa(enrol.RoleID))
		form.Set(fmt.Sprintf("enrolments[%d][userid]", i), strconv.Itoa(enrol.UserID))
		form.Set(fmt.Sprintf("enrolments[%d][courseid]", i), strconv.Itoa(enrol.CourseID))

		// Jika timestart, timeend, suspend ada, baru set
		if enrol.TimeStart != 0 {
			form.Set(fmt.Sprintf("enrolments[%d][timestart]", i), strconv.Itoa(enrol.TimeStart))
		}
		if enrol.TimeEnd != 0 {
			form.Set(fmt.Sprintf("enrolments[%d][timeend]", i), strconv.Itoa(enrol.TimeEnd))
		}
		if enrol.Suspend != 0 {
			form.Set(fmt.Sprintf("enrolments[%d][suspend]", i), strconv.Itoa(enrol.Suspend))
		}
	}

	// kirim post ke moodle
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		log.Printf("[ERROR] CoreCourseCreateCourses: gagal melakukan request ke Moodle: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] CoreCourseCreateCourses: gagal membaca body: %v", err)
		return err
	}
	log.Printf("[DEBUG] CoreCourseCreateCourses Raw Response: %s", string(body))

	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return &moodleErr
	}

	return nil
}

// function untuk CoreCourseCreateCourses dan Enroll
func (s *MoodleServiceImpl) CreateCourseWithEnrollUser(req web.MoodleCreateCourseWithEnrollUserRequest) (*web.MoodleCreateCourseWithEnrollUserResponse, error) {
	// panggil CoreCourseCreateCourses
	log.Printf("[INFO] CoreCourseCreateCoursesAndEnrolUser: Memulai pembuatan kursus...")
	courseReq := web.MoodleCoreCourseCreateCoursesRequest{
		Courses: []web.MoodleCourseData{req.CourseData},
	}
	createdCourses, err := s.CoreCourseCreateCourses(courseReq)
	if err != nil {
		log.Printf("[ERROR] Gagal pada langkah pembuatan kursus: %v", err)
		return nil, err
	}
	if len(createdCourses) == 0 {
		return nil, errors.New("pembuatan kursus berhasil namun tidak ada data yang dikembalikan Moodle")
	}
	newCourse := createdCourses[0]
	log.Printf("[INFO] Kursus berhasil dibuat dengan ID: %d", newCourse.ID)

	// panggil enrollManualEnrollUsers
	log.Printf("[INFO] Memulai pendaftaran pengguna (UserID: %d) ke kursus (CourseID: %d)", req.UserID, newCourse.ID)
	enrollReq := web.MoodleManualEnrollRequest{
		Enrolments: []web.MoodleManualEnroll{
			{
				RoleID:   req.RoleID,
				UserID:   req.UserID,
				CourseID: newCourse.ID,
			},
		},
	}
	err = s.EnrolManualEnrolUsers(enrollReq)
	if err != nil {
		log.Printf("[ERROR] Gagal pada langkah pendaftaran pengguna: %v", err)
		return nil, fmt.Errorf("kursus berhasil dibuat, tetapi gagal mendaftarkan pengguna: %w", err)
	}
	log.Printf("[INFO] Pendaftaran pengguna berhasil.")

	// response
	finalResponse := &web.MoodleCreateCourseWithEnrollUserResponse{
		CourseID:        newCourse.ID,
		CourseFullName:  newCourse.Fullname,
		CourseShortName: newCourse.ShortName,
		EnrolledUserID:  req.UserID,
		AssignedRoleID:  req.RoleID,
	}

	return finalResponse, nil
}

func (s *MoodleServiceImpl) CoreCourseCreateCategories(req web.MoodleCreateCategoriesRequest) ([]web.MoodleCreateCategoriesResponse, error) {
	//load konfig moodle
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan konfigurasi Moodle: %w", err)
	}

	//form
	form := helpers.NewMoodleForm(token, "core_course_create_categories")

	//loop untuk banyak categories
	for i, category := range req.Categories {
		prefix := fmt.Sprintf("categories[%d]", i)

		form.Set(prefix+"[name]", category.Name)
		form.Set(prefix+"[parent]", strconv.Itoa(category.Parent))

		if category.IDNumber != "" {
			form.Set(prefix+"[idnumber]", category.IDNumber)
		}
		if category.Description != "" {
			form.Set(prefix+"[description]", category.Description)
		}
		if category.DescriptionFormat != 0 {
			form.Set(prefix+"[descriptionformat]", strconv.Itoa(category.DescriptionFormat))
		}
		if category.Theme != "" {
			form.Set(prefix+"[theme]", category.Theme)
		}
	}

	//kirim post ke moodl
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		return nil, fmt.Errorf("gagal memanggil Moodle untuk CreateCategories: %w", err)
	}
	defer resp.Body.Close()

	//baca resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca body respons CreateCategories: %w", err)
	}
	log.Printf("[CreateCategories] Raw Response: %s", string(body))

	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		return nil, &moodleErr
	}

	//jika tidak error
	var result []web.MoodleCreateCategoriesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("gagal unmarshal respons sukses CreateCategories: %w", err)
	}

	return result, nil
}

func (s *MoodleServiceImpl) CoreCourseUpdateCategories(req web.MoodleUpdateCategoriesRequest) error {
	//load moodle
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		log.Printf("[ERROR] UpdateCategories: Gagal mendapatkan konfigurasi Moodle: %v", err)
		return fmt.Errorf("gagal mendapatkan konfigurasi Moodle: %w", err)
	}

	//form
	form := helpers.NewMoodleForm(token, "core_course_update_categories")
	log.Printf("[INFO] UpdateCategories: Memulai proses untuk %d kategori.", len(req.Categories))

	//loop untuk categories
	for i, category := range req.Categories {
		prefix := fmt.Sprintf("categories[%d]", i)
		form.Set(prefix+"[id]", strconv.Itoa(category.ID))
		if category.Name != "" {
			form.Set(prefix+"[name]", category.Name)
		}
		if category.IDNumber != "" {
			form.Set(prefix+"[idnumber]", category.IDNumber)
		}
		// Untuk pointer, cek jika nilainya tidak nil
		if category.Parent != nil {
			form.Set(prefix+"[parent]", strconv.Itoa(*category.Parent))
		}
		if category.Description != "" {
			form.Set(prefix+"[description]", category.Description)
		}
		if category.DescriptionFormat != nil {
			form.Set(prefix+"[descriptionformat]", strconv.Itoa(*category.DescriptionFormat))
		}
		if category.Theme != "" {
			form.Set(prefix+"[theme]", category.Theme)
		}
	}

	// DEBUG LOG: Tampilkan data form yang akan dikirim
	log.Printf("[DEBUG] UpdateCategories: Form data yang akan dikirim:\n%s", form.Encode())

	//kirim post
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		log.Printf("[ERROR] UpdateCategories: Gagal melakukan request ke Moodle: %v", err)
		return fmt.Errorf("gagal memanggil Moodle untuk UpdateCategories: %w", err)
	}
	defer resp.Body.Close()

	//read resp body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] UpdateCategories: Gagal membaca body respons: %v", err)
		return fmt.Errorf("gagal membaca body respons UpdateCategories: %w", err)
	}
	log.Printf("[DEBUG] UpdateCategories Raw Response: %s", string(body))

	//cek jika moodle mengembalikan error
	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		log.Printf("[ERROR] UpdateCategories: Moodle mengembalikan error: %v", &moodleErr)
		return &moodleErr
	}

	log.Printf("[INFO] UpdateCategories: Operasi update kategori berhasil.")
	return nil
}

func (s *MoodleServiceImpl) CoreCourseDeleteCategories(req web.MoodleDeleteCategoriesRequest) error {
	//load konfig moodle
	moodleURL, token, err := helpers.GetMoodleConfig()
	if err != nil {
		log.Printf("[ERROR] DeleteCategories: Gagal mendapatkan konfigurasi Moodle: %v", err)
		return fmt.Errorf("gagal mendapatkan konfigurasi Moodle: %w", err)
	}

	//form
	form := helpers.NewMoodleForm(token, "core_course_delete_categories")
	log.Printf("[INFO] DeleteCategories: Memulai proses untuk menghapus %d kategori.", len(req.Categories))

	for i, category := range req.Categories {
		prefix := fmt.Sprintf("categories[%d]", i)
		form.Set(prefix+"[id]", strconv.Itoa(category.ID))

		if category.NewParent != nil {
			form.Set(prefix+"[newparent]", strconv.Itoa(*category.NewParent))
		}
		if category.Recursive != nil {
			form.Set(prefix+"[recursive]", strconv.Itoa(*category.Recursive))
		}
	}
	log.Printf("[DEBUG] DeleteCategories: Form data yang akan dikirim:\n%s", form.Encode())

	//kirim request
	resp, err := s.client.PostForm(moodleURL, form)
	if err != nil {
		log.Printf("[ERROR] DeleteCategories: Gagal melakukan request ke Moodle: %v", err)
		return fmt.Errorf("gagal memanggil Moodle untuk DeleteCategories: %w", err)
	}
	defer resp.Body.Close()

	//baca response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] DeleteCategories: Gagal membaca body respons: %v", err)
		return fmt.Errorf("gagal membaca body respons DeleteCategories: %w", err)
	}
	log.Printf("[DEBUG] DeleteCategories Raw Response: %s", string(body))

	var moodleErr web.MoodleException
	if json.Unmarshal(body, &moodleErr) == nil && moodleErr.Exception != "" {
		log.Printf("[ERROR] DeleteCategories: Moodle mengembalikan error: %v", &moodleErr)
		return &moodleErr
	}

	//periksa body
	if bodyStr := string(body); bodyStr != "" {
		log.Printf("[WARN] DeleteCategories: Moodle mengembalikan respons tak terduga: %s", bodyStr)
	}

	log.Printf("[INFO] DeleteCategories: Operasi hapus kategori berhasil.")
	return nil
}
