package project

type UpdateProjectDTO struct {
	ProjectCategoryId *uint   `json:"projectCategoryId"`
	Title             *string `json:"title"`
	ImageUrl          *string `json:"imageUrl"`
	Url               *string `json:"url"`
}
