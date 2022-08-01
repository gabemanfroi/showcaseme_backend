package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/article"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type ArticleRepository struct {
	sqlClient *gorm.DB
}

func CreateArticleRepository() *ArticleRepository {
	return &ArticleRepository{sqlClient: db.GetSqlInstance()}
}

func (repository ArticleRepository) Create(dto *article.CreateArticleDTO) (*article.ReadArticleDTO, error) {
	a := models.Article{
		UserId:  dto.UserId,
		Title:   dto.Title,
		Content: dto.Content,
	}
	repository.sqlClient.Create(&a)

	if a.ID == 0 {
		return nil, errors.New("an error has occured when creating your article, verify")
	}

	createdarticle, _ := repository.GetById(a.ID)

	return createdarticle, nil
}

func (repository ArticleRepository) GetAll() ([]*article.ReadArticleDTO, error) {
	var articles []*models.Article
	var articleDTOs []*article.ReadArticleDTO

	repository.sqlClient.Find(&articles)

	for _, a := range articles {
		articleDTOs = append(articleDTOs, &article.ReadArticleDTO{
			ID:      a.ID,
			Title:   a.Title,
			Content: a.Content,
		})
	}

	return articleDTOs, nil
}

func (repository ArticleRepository) GetById(id uint) (*article.ReadArticleDTO, error) {
	var a *models.Article

	repository.sqlClient.Find(&a, id)

	if a.ID == 0 {
		return nil, errors.New("article not found")
	}

	return &article.ReadArticleDTO{
		ID:      a.ID,
		Title:   a.Title,
		Content: a.Content,
	}, nil
}

func (repository ArticleRepository) Delete(id uint) error {
	var a models.Article
	repository.sqlClient.Find(&a, id)
	if a.ID == 0 {
		return errors.New("article not found")
	}
	repository.sqlClient.Delete(&a)
	return nil
}

func (repository ArticleRepository) Update(id uint, dto *article.UpdateArticleDTO) (*article.ReadArticleDTO, error) {
	var a models.Article

	repository.sqlClient.Find(&a, id)

	if a.ID == 0 {
		return nil, errors.New("article not found")
	}

	updateArticleValuesFromDTO(&a, dto)
	repository.sqlClient.Save(&a)

	updatedArticle, _ := repository.GetById(a.ID)

	return updatedArticle, nil
}

func updateArticleValuesFromDTO(model *models.Article, dto *article.UpdateArticleDTO) {
	if dto.Title != nil {
		model.Title = *dto.Title
	}
	if dto.Content != nil {
		model.Content = *dto.Content
	}
}
