package services

import (
	"showcaseme/domain/DTO/article"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
	"github.com/golobby/container/v3"
)

type ArticleService struct {
	repository repositories.IArticleRepository
}

func CreateArticleService() *ArticleService { return &ArticleService{repository: getArticleService()} }

func (service ArticleService) Create(dto *article.CreateArticleDTO) (*article.ReadArticleDTO, error) {
	return service.repository.Create(dto)
}

func (service ArticleService) GetAll() ([]*article.ReadArticleDTO, error) {
	return service.repository.GetAll()
}

func (service ArticleService) GetById(id uint) (*article.ReadArticleDTO, error) {
	return service.repository.GetById(id)
}

func (service ArticleService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service ArticleService) Update(id uint, dto *article.UpdateArticleDTO) (*article.ReadArticleDTO, error) {
	return service.repository.Update(id, dto)
}

func getArticleService() repositories.IArticleRepository {
	var injector repositories.IArticleRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving ArticleRepository instance")
	return injector
}
