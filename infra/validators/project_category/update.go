package project_category

type UpdateProjectCategoryValidator struct {
	Name  string `validate:"omitempty,min=10"`
}
