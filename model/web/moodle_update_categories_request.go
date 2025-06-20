package web

type UpdateCategoryData struct {
	ID                int    `json:"id" validate:"required"`
	Name              string `json:"name,omitempty"`
	IDNumber          string `json:"idnumber,omitempty"`
	Parent            *int   `json:"parent,omitempty"`
	Description       string `json:"description,omitempty"`
	DescriptionFormat *int   `json:"descriptionformat,omitempty"`
	Theme             string `json:"theme,omitempty"`
}

type MoodleUpdateCategoriesRequest struct {
	Categories []UpdateCategoryData `json:"categories" validate:"required"`
}
