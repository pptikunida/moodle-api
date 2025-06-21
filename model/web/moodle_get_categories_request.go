package web

type CategoryGetData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MoodleGetCategoriesRequest struct {
	Criteria         []CategoryData `json:"criteria,omitempty"`
	AddSubcategories *int           `json:"addsubcategories,omitempty"`
}
