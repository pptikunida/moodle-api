package web

type MoodleUserUpdateRequest struct {
	ID                int                               `json:"id"`
	Username          string                            `json:"username,omitempty"`
	Auth              string                            `json:"auth,omitempty"`
	Suspended         int                               `json:"suspended,omitempty"`
	Password          string                            `json:"password,omitempty"`
	Firstname         string                            `json:"firstname,omitempty"`
	Lastname          string                            `json:"lastname,omitempty"`
	Email             string                            `json:"email,omitempty"`
	MailDisplay       int                               `json:"maildisplay,omitempty"`
	City              string                            `json:"city,omitempty"`
	Country           string                            `json:"country,omitempty"`
	Timezone          string                            `json:"timezone,omitempty"`
	Description       string                            `json:"description,omitempty"`
	UserPicture       int                               `json:"userpicture,omitempty"`
	FirstnamePhonetic string                            `json:"firstnamephonetic,omitempty"`
	LastnamePhonetic  string                            `json:"lastnamephonetic,omitempty"`
	Middlename        string                            `json:"middlename,omitempty"`
	Alternatename     string                            `json:"alternatename,omitempty"`
	Interests         string                            `json:"interests,omitempty"`
	IdNumber          string                            `json:"idnumber,omitempty"`
	Institution       string                            `json:"institution,omitempty"`
	Department        string                            `json:"department,omitempty"`
	Phone1            string                            `json:"phone1,omitempty"`
	Phone2            string                            `json:"phone2,omitempty"`
	Address           string                            `json:"address,omitempty"`
	Lang              string                            `json:"lang,omitempty"`
	CalendarType      string                            `json:"calendartype,omitempty"`
	Theme             string                            `json:"theme,omitempty"`
	MailFormat        int                               `json:"mailformat,omitempty"`
	CustomFields      []MoodleUserUpdateField           `json:"customFields,omitempty"`
	Preferences       []MoodleUserPreferenceUpdateField `json:"preferences,omitempty"`
}

type MoodleUserUpdateField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MoodleUserPreferenceUpdateField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
