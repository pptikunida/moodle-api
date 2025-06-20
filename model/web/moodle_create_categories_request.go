package web

type CategoryData struct {
	Name              string `json:"name" validate:"required"`
	Parent            int    `json:"parent,omitempty"`
	IDNumber          string `json:"idnumber,omitempty"`
	Description       string `json:"description,omitempty"`
	DescriptionFormat int    `json:"descriptionformat,omitempty"`
	Theme             string `json:"theme,omitempty"`
}

type MoodleCreateCategoriesRequest struct {
	Categories []CategoryData `json:"categories" validate:"required"`
}
