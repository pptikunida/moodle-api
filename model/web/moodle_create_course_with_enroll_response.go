package web

type MoodleCreateCourseWithEnrollResponse struct {
	CourseID        int    `json:"course_id"`
	CourseFullName  string `json:"course_fullname"`
	CourseShortName string `json:"course_shortname"`
	EnrolledUserID  int    `json:"enrolled_user_id"`
	AssignRoleID    int    `json:"assign_role_id"`
}
