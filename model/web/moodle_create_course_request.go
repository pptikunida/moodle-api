package web

type MoodleCreateCourseRequest struct {
	Courses []MoodleCourseData `json:"courses" binding:"required"`
}

type MoodleCourseData struct {
	FullName            string               `json:"fullname" binding:"required"`
	ShortName           string               `json:"shortname" binding:"required"`
	CategoryID          int                  `json:"categoryid" binding:"required"`
	IDNumber            string               `json:"idnumber,omitempty"`
	Summary             string               `json:"summary,omitempty"`
	SummaryFormat       int                  `json:"summaryformat,omitempty"` // Default: 1
	Format              string               `json:"format,omitempty"`        // Default: "topics"
	ShowGrades          int                  `json:"showgrades,omitempty"`    // Default: 1
	NewsItems           int                  `json:"newsitems,omitempty"`     // Default: 5
	StartDate           int64                `json:"startdate,omitempty"`     // Unix timestamp
	EndDate             int64                `json:"enddate,omitempty"`
	NumSections         int                  `json:"numsections,omitempty"` // Deprecated
	MaxBytes            int                  `json:"maxbytes,omitempty"`    // Default: 0
	ShowReports         int                  `json:"showreports,omitempty"` // Default: 0
	Visible             int                  `json:"visible,omitempty"`     // 1 = show, 0 = hidden
	HiddenSections      int                  `json:"hiddensections,omitempty"`
	GroupMode           int                  `json:"groupmode,omitempty"`      // Default: 0
	GroupModeForce      int                  `json:"groupmodeforce,omitempty"` // Default: 0
	DefaultGroupingID   int                  `json:"defaultgroupingid,omitempty"`
	EnableCompletion    int                  `json:"enablecompletion,omitempty"`
	CompletionNotify    int                  `json:"completionnotify,omitempty"`
	Lang                string               `json:"lang,omitempty"`
	ForceTheme          string               `json:"forcetheme,omitempty"`
	CourseFormatOptions []CourseFormatOption `json:"courseformatoptions,omitempty"`
	CustomFields        []MoodleCustomField  `json:"customfields,omitempty"`
}

type CourseFormatOption struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type MoodleCustomField struct {
	ShortName string `json:"shortname" binding:"required"`
	Value     string `json:"value" binding:"required"`
}
