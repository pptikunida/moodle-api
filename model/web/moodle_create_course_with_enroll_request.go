package web

type MoodleCreateCourseWithEnrollUserRequest struct {
	CourseData MoodleCourseData `json:"course_data"`
	UserID     int              `json:"user_id"`
	RoleID     int              `json:"role_id"`
}
