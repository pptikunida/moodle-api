package web

type DeleteCategoryData struct {
	ID        int  `json:"id" validate:"required"`
	NewParent *int `json:"newparent,omitempty"`
	Recursive *int `json:"recursive,omitempty"` // Default to "0" //1: recursively delete all contents inside this category, 0 (default): move contents to newparent or current parent category (except if parent is root)
}

type MoodleDeleteCategoriesRequest struct {
	Categories []DeleteCategoryData `json:"categories" validate:"required"`
}
