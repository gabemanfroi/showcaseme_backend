package project

import (
	"showcaseme/domain/DTO/project_category"
)

type ReadProjectDTO struct {
	ID              uint                                     `json:"id"`
	ProjectCategory *project_category.ReadProjectCategoryDTO `json:"projectCategory"`
	ImageUrl        string                                   `json:"imageUrl"`
	Url             string                                   `json:"url"`
	Title           string                                   `json:"title"`
}
