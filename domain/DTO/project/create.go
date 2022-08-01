package project

type CreateProjectDTO struct {
	UserId            uint   `json:"userId"`
	ProjectCategoryId uint   `json:"projectCategoryId"`
	Title             string `json:"title"`
	Url               string `json:"url"`
}
