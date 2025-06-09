package web

type MoodleUserCustomField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MoodleUserPreferenceField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MoodleUserCreateResponse struct {
	CreatePassword    int                         `json:"create_password"`
	Username          string                      `json:"username"`
	Auth              string                      `json:"auth"`
	Password          string                      `json:"password"`
	FirstName         string                      `json:"first_name"`
	LastName          string                      `json:"last_name"`
	Email             string                      `json:"email"`
	MailDisplay       int                         `json:"mail_display"`
	City              string                      `json:"city"`
	Country           string                      `json:"country"`
	Timezone          string                      `json:"timezone"`
	Description       string                      `json:"description"`
	FirstnamePhonetic string                      `json:"firstname_phonetic"`
	LastnamePhonetic  string                      `json:"lastname_phonetic"`
	Middlename        string                      `json:"middlename"`
	Alternatename     string                      `json:"alternatename"`
	Interests         string                      `json:"interests,omitempty"`
	IdNumber          string                      `json:"idnumber,omitempty"` // default: ""
	Institution       string                      `json:"institution,omitempty"`
	Department        string                      `json:"department,omitempty"`
	Phone1            string                      `json:"phone1,omitempty"`
	Phone2            string                      `json:"phone2,omitempty"`
	Address           string                      `json:"address,omitempty"`
	Lang              string                      `json:"lang,omitempty"`         // default: "en"
	CalendarType      string                      `json:"calendartype,omitempty"` // default: "gregorian"
	Theme             string                      `json:"theme,omitempty"`
	MailFormat        *int                        `json:"mailformat,omitempty"`
	CustomFields      []MoodleUserCustomField     `json:"custom_fields"`
	Preferences       []MoodleUserPreferenceField `json:"preferences"`
}
