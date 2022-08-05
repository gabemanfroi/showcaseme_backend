package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

type CarouselItemRepository struct {
	sqlClient *gorm.DB
}

func CreateCarouselItemRepository() *CarouselItemRepository {
	return &CarouselItemRepository{sqlClient: db.GetSqlInstance()}
}

func (repository CarouselItemRepository) Create(dto *carousel_item.CreateCarouselItemDTO) (*carousel_item.ReadCarouselItemDTO, error) {
	c := models.CarouselItem{
		UserId:   dto.UserId,
		Content:  dto.Content,
		Position: dto.Position,
	}
	repository.sqlClient.Create(&c)

	if c.ID == 0 {
		return nil, errors.New("an error has occured when creating your carousel_item, verify")
	}

	return &carousel_item.ReadCarouselItemDTO{
		ID:       c.ID,
		Content:  c.Content,
		Position: c.Position,
	}, nil
}

func (repository CarouselItemRepository) GetAll() ([]*carousel_item.ReadCarouselItemDTO, error) {
	var skillCategories []*models.CarouselItem
	var skillCategoriesDto []*carousel_item.ReadCarouselItemDTO

	repository.sqlClient.Find(&skillCategories)

	for _, c := range skillCategories {
		skillCategoriesDto = append(skillCategoriesDto, &carousel_item.ReadCarouselItemDTO{
			ID:       c.ID,
			Content:  c.Content,
			Position: c.Position,
		})
	}

	return skillCategoriesDto, nil
}

func (repository CarouselItemRepository) GetById(id uint) (*carousel_item.ReadCarouselItemDTO, error) {
	var c *models.CarouselItem
	repository.sqlClient.Find(&c, id)

	if c.ID == 0 {
		return nil, errors.New("carousel_item not found")
	}

	return &carousel_item.ReadCarouselItemDTO{ID: c.ID, Content: c.Content}, nil
}

func (repository CarouselItemRepository) Delete(id uint) error {
	var c models.CarouselItem

	repository.sqlClient.Find(&c, id)

	if c.ID == 0 {
		return errors.New("carousel_item not found")
	}

	repository.sqlClient.Delete(&c)
	return nil
}

func (repository CarouselItemRepository) Update(id uint, dto *carousel_item.UpdateCarouselItemDTO) (*carousel_item.ReadCarouselItemDTO, error) {
	var c models.CarouselItem

	repository.sqlClient.Find(&c, id)

	if c.ID == 0 {
		return nil, errors.New("carousel_item not found")
	}

	utils.UpdateModelValuesFromDTO(&c, dto)
	repository.sqlClient.Save(&c)

	return &carousel_item.ReadCarouselItemDTO{
		ID:      c.ID,
		Content: c.Content,
	}, nil
}
