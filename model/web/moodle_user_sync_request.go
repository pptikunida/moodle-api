package web

type MoodleUserSyncRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	NIM       string `json:"NIM" validate:"required"`

	// tambahan
	Auth         string `json:"auth,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Description  string `json:"description,omitempty"`
	Lang         string `json:"lang,omitempty"`
	Calendartype string `json:"calendartype,omitempty"`

	CustomFields []MoodleUserSyncCustomField `json:"custom_fields,omitempty"`
	Preferences  []MoodleUserSyncPreference  `json:"preferences,omitempty"`
}

type MoodleUserSyncCustomField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MoodleUserSyncPreference struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
