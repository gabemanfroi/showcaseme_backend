package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/skill_category"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type SkillCategoryRepository struct {
	sqlClient *gorm.DB
}

func CreateSkillCategoryRepository() *SkillCategoryRepository {
	return &SkillCategoryRepository{sqlClient: db.GetSqlInstance()}
}

func (repository SkillCategoryRepository) Create(dto *skill_category.CreateSkillCategoryDTO) (*skill_category.ReadSkillCategoryDTO, error) {
	s := models.SkillCategory{
		UserId: dto.UserId,
		Name:   dto.Name,
	}
	repository.sqlClient.Create(&s)

	if s.ID == 0 {
		return nil, errors.New("an error has occured when creating your skill_category, verify")
	}

	return &skill_category.ReadSkillCategoryDTO{
		ID:   s.ID,
		Name: s.Name,
	}, nil
}

func (repository SkillCategoryRepository) GetAll() ([]*skill_category.ReadSkillCategoryDTO, error) {
	var skillCategories []*models.SkillCategory
	var skillCategoriesDto []*skill_category.ReadSkillCategoryDTO

	repository.sqlClient.Find(&skillCategories)

	for _, s := range skillCategories {
		skillCategoriesDto = append(skillCategoriesDto, &skill_category.ReadSkillCategoryDTO{
			ID:   s.ID,
			Name: s.Name,
		})
	}

	return skillCategoriesDto, nil
}

func (repository SkillCategoryRepository) GetById(id uint) (*skill_category.ReadSkillCategoryDTO, error) {
	var s *models.SkillCategory
	repository.sqlClient.Find(&s, id)

	if s.ID == 0 {
		return nil, errors.New("skill_category not found")
	}

	return &skill_category.ReadSkillCategoryDTO{ID: s.ID, Name: s.Name}, nil
}

func (repository SkillCategoryRepository) Delete(id uint) error {
	var s models.SkillCategory

	repository.sqlClient.Find(&s, id)

	if s.ID == 0 {
		return errors.New("skill_category not found")
	}

	repository.sqlClient.Where(&models.Skill{SkillCategoryId: s.ID}).Delete(&models.Skill{})
	repository.sqlClient.Delete(&s)
	return nil
}

func (repository SkillCategoryRepository) Update(id uint, dto *skill_category.UpdateSkillCategoryDTO) (*skill_category.ReadSkillCategoryDTO, error) {
	var s models.SkillCategory

	repository.sqlClient.Find(&s, id)

	if s.ID == 0 {
		return nil, errors.New("skill_category not found")
	}

	updateSkillCategoryValuesFromDTO(&s, dto)
	repository.sqlClient.Save(&s)

	return &skill_category.ReadSkillCategoryDTO{
		ID:   s.ID,
		Name: s.Name,
	}, nil
}

func updateSkillCategoryValuesFromDTO(model *models.SkillCategory, dto *skill_category.UpdateSkillCategoryDTO) {
	if dto.Name != nil {
		model.Name = *dto.Name
	}
}
