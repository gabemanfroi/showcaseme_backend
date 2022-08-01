package project

import "showcaseme/domain/models"

type ReadProjectDTO struct {
	ID              uint                   `json:"id"`
	ProjectCategory models.ProjectCategory `json:"projectCategory"`
	ImageUrl        string                 `json:"imageUrl"`
	Url             string                 `json:"url"`
	Title           string                 `json:"title"`
}
