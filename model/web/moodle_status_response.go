package web

// api moodle core_webservice_get_site_info
type MoodleStatusFunction struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type MoodleAdvancedFeature struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MoodleStatusResponse struct {
	SiteName              string                  `json:"sitename"`
	UserName              string                  `json:"username"`
	FirstName             string                  `json:"firstname"`
	LastName              string                  `json:"lastname"`
	FullName              string                  `json:"fullname"`
	Lang                  string                  `json:"lang"`
	UserID                int                     `json:"userid"`
	SiteURL               string                  `json:"siteurl"`
	UserPictureURL        string                  `json:"userpictureurl"`
	Functions             []MoodleStatusFunction  `json:"functions"`
	DownloadFiles         int                     `json:"downloadFiles"`
	UploadFiles           int                     `json:"uploadFiles"`
	Release               string                  `json:"release"`
	Version               string                  `json:"version"`
	MobileCSSURL          string                  `json:"mobilecssurl"`
	AdvancedFeatures      []MoodleAdvancedFeature `json:"advancedfeatures"`
	UserCanManageOwnFiles bool                    `json:"usercanmanageownfiles"`
	UserQuota             int                     `json:"userquota"`
	UserMaxUploadFileSize int                     `json:"usermaxuploadfilesize"`
	UserHomePage          int                     `json:"userhomepage"`
	UserHomePageURL       string                  `json:"userhomepageurl"`
	UserPrivateAccessKey  string                  `json:"userprivateaccesskey"`
	SiteID                int                     `json:"siteid"`
	SiteCalenderType      string                  `json:"sitecalendertype"`
	UserCalenderType      string                  `json:"usercalendertype"`
	UserIsSiteAdmin       bool                    `json:"userissiteadmin"`
	Theme                 string                  `json:"theme"`
	LimitConcurrentLogins int                     `json:"limitconcurrentlogins"`
	UserSessionCount      int                     `json:"usersessioncount"`
	PolicyAgreed          int                     `json:"policyagreed"`
}
