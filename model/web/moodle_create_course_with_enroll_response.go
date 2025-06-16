package web

type MoodleCreateCourseWithEnrollUserResponse struct {
	CourseID        int    `json:"course_id"`
	CourseFullName  string `json:"course_fullname"`
	CourseShortName string `json:"course_shortname"`
	EnrolledUserID  int    `json:"enrolled_user_id"`
	AssignedRoleID  int    `json:"assign_role_id"`
}
