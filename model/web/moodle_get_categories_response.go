package web

type MoodleGetCategoriesResponse struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	IDNumber          string `json:"idnumber,omitempty"`
	Description       string `json:"description"`
	DescriptionFormat string `json:"descriptionformat"`
	Parent            int    `json:"parent"`
	SortOrder         int    `json:"sortorder"`
	CourseCount       int    `json:"coursecount"`
	Visible           *int   `json:"visible,omitempty"`
	VisibleOld        *int   `json:"visibleold,omitempty"`
	TimeModified      *int   `json:"timemodified,omitempty"`
	Depth             int    `json:"depth"` // Tingkat kedalaman kategori
	Path              string `json:"path"`  // Path ID, contoh: /1/7
	Theme             string `json:"theme,omitempty"`
}
