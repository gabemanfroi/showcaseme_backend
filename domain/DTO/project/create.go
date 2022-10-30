package project

import "mime/multipart"

type CreateProjectDTO struct {
	UserId            uint                  `json:"userId" form:"userId"`
	ProjectCategoryId uint                  `json:"projectCategoryId" form:"projectCategoryId"`
	Title             string                `json:"title" form:"title"`
	Url               string                `json:"url" form:"url"`
	BackgroundImage   *multipart.FileHeader `form:"backgroundImage"`
}
