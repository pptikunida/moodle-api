package web

type MoodleCreateCourseWithEnrollUserRequest struct {
	CourseData MoodleCourseData `json:"course_data"`
	UserID     int              `json:"teacher_user_id"`
	RoleID     int              `json:"teacher_role_id"`
}
