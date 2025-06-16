package web

type MoodleManualEnroll struct {
	RoleID    int `json:"roleid" binding:"required"`
	UserID    int `json:"userid" binding:"required"`
	CourseID  int `json:"courseid" binding:"required"`
	TimeStart int `json:"timestart" binding:"required"`
	TimeEnd   int `json:"timeend" binding:"required"`
	Suspend   int `json:"suspend" binding:"required"`
}

type MoodleManualEnrollRequest struct {
	Enrolments []MoodleManualEnroll `json:"enrolments"`
}
