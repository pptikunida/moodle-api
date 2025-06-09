package web

type MoodleUserCustomField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MoodleUserPreferenceField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MoodleUserCreateRequest struct {
	CreatePassword    *int                        `json:"createpassword,omitempty"`
	Username          string                      `json:"username"`
	Auth              string                      `json:"auth,omitempty"` // default: manual
	Password          string                      `json:"password,omitempty"`
	Firstname         string                      `json:"firstname"`
	Lastname          string                      `json:"lastname"`
	Email             string                      `json:"email"`
	MailDisplay       *int                        `json:"maildisplay,omitempty"`
	City              string                      `json:"city,omitempty"`
	Country           string                      `json:"country,omitempty"`
	Timezone          string                      `json:"timezone,omitempty"`
	Description       string                      `json:"description,omitempty"`
	FirstnamePhonetic string                      `json:"firstnamephonetic,omitempty"`
	LastnamePhonetic  string                      `json:"lastnamephonetic,omitempty"`
	Middlename        string                      `json:"middlename,omitempty"`
	Alternatename     string                      `json:"alternatename,omitempty"`
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
	CustomFields      []MoodleUserCustomField     `json:"customfields,omitempty"`
	Preferences       []MoodleUserPreferenceField `json:"preferences,omitempty"`
}
