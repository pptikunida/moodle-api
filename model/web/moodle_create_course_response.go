package web

type MoodleCoreCourseCreateCoursesResponse struct {
	ID         int    `json:"id"`
	ShortName  string `json:"shortname"`
	Fullname   string `json:"fullname"`
	CategoryID int    `json:"categoryid"`
	IDNumber   int    `json:"idnumber"`
	Summary    string `json:"summary"`
}
