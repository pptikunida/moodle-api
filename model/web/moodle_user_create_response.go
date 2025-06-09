package web

type MoodleUserCreateResponse struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
}
