package web

type MoodleCreateCourseWithEnrollRequest struct {
	CourseData    MoodleCourseData `json:"course_data"`
	TeacherUserID int              `json:"teacher_user_id"`
	TeacherRoleID int              `json:"teacher_role_id"`
}
