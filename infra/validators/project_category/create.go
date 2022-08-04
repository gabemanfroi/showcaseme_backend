package project_category

type CreateProjectCategoryValidator struct {
	UserId uint   `validate:"required,gte=1"`
	Name   string `validate:"required,min=10"`
}
