package web

type MoodleUserUpdateResponse struct {
	Warnings []MoodleUserUpdateWarning `json:"warnings,omitempty"`
}

type MoodleUserUpdateWarning struct {
	Item        string `json:"item"`
	ItemID      int    `json:"itemid,omitempty"`
	WarningCode string `json:"warningcode"`
	Message     string `json:"message"`
}
