package repositories

import (
	"showcaseme/domain/DTO/article"
)

type IArticleRepository interface {
	Create(dto *article.CreateArticleDTO) (*article.ReadArticleDTO, error)
	GetAll() ([]*article.ReadArticleDTO, error)
    GetById(id uint) (*article.ReadArticleDTO, error)
    Delete(id uint) error
    Update(id uint, dto *article.UpdateArticleDTO) (*article.ReadArticleDTO, error)
}
