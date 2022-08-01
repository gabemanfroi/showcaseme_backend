package project

type UpdateProjectValidator struct {
	ProjectCategoryId  uint `validate:"omitempty,min=1"`
	Title  string `validate:"omitempty,min=10"`
	ImageUrl  string `validate:"omitempty,min=20"`
	url  string `validate:"omitempty,min=20"`
}
