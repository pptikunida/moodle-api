package web

type MoodleUserGetByFieldResponse struct {
	ID                int                                  `json:"id"`
	Username          string                               `json:"username,omitempty"`
	Firstname         string                               `json:"firstname,omitempty"`
	Lastname          string                               `json:"lastname,omitempty"`
	Fullname          string                               `json:"fullname,omitempty"`
	Email             string                               `json:"email,omitempty"`
	Address           string                               `json:"address,omitempty"`
	Phone1            string                               `json:"phone1,omitempty"`
	Phone2            string                               `json:"phone2,omitempty"`
	Department        string                               `json:"department,omitempty"`
	Institution       string                               `json:"institution,omitempty"`
	IdNumber          string                               `json:"idnumber,omitempty"`
	Interests         string                               `json:"interests,omitempty"`
	FirstAccess       int                                  `json:"firstaccess,omitempty"`
	LastAccess        int                                  `json:"lastaccess,omitempty"`
	Auth              string                               `json:"auth,omitempty"`
	Suspended         bool                                 `json:"suspended,omitempty"`
	Confirmed         bool                                 `json:"confirmed,omitempty"`
	Lang              string                               `json:"lang,omitempty"`
	CalendarType      string                               `json:"calendartype,omitempty"`
	Theme             string                               `json:"theme,omitempty"`
	Timezone          string                               `json:"timezone,omitempty"`
	MailFormat        int                                  `json:"mailformat,omitempty"`
	TrackForums       int                                  `json:"trackforums,omitempty"`
	Description       string                               `json:"description,omitempty"`
	DescriptionFormat int                                  `json:"descriptionformat,omitempty"`
	City              string                               `json:"city,omitempty"`
	Country           string                               `json:"country,omitempty"`
	ProfileImageSmall string                               `json:"profileimageurlsmall"`
	ProfileImage      string                               `json:"profileimageurl"`
	CustomFields      []MoodleUserCustomGetUserByField     `json:"customfields,omitempty"`
	Preferences       []MoodleUserPreferenceGetUserByField `json:"preferences,omitempty"`
}

type MoodleUserCustomGetUserByField struct {
	Type         string `json:"type"`
	Value        string `json:"value"`
	DisplayValue string `json:"displayvalue,omitempty"`
	Name         string `json:"name"`
	ShortName    string `json:"shortname"`
}

type MoodleUserPreferenceGetUserByField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
