package project

type CreateProjectValidator struct {
	ProjectCategoryId  uint `validate:"required,gte=1"`
	UserId  uint `validate:"required,gte=1"`
	Title  string `validate:"required,min=10,max=1000"`
	url  string `validate:"required,min=20"`
}
