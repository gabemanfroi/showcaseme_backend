package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/DTO/skill_category"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

type SkillRepository struct {
	sqlClient *gorm.DB
}

func CreateSkillRepository() *SkillRepository {
	return &SkillRepository{sqlClient: db.GetSqlInstance()}
}

func (repository SkillRepository) Create(dto *skill.CreateSkillDTO) (*skill.ReadSkillDTO, error) {
	s := models.Skill{
		UserId:          dto.UserId,
		Name:            dto.Name,
		Proficiency:     dto.Proficiency,
		SkillCategoryId: dto.SkillCategoryId,
	}
	repository.sqlClient.Create(&s)

	if s.ID == 0 {
		return nil, errors.New("an error has occured when creating your skill, verify")
	}

	createdSkill, _ := repository.GetById(s.ID)

	return createdSkill, nil
}

func (repository SkillRepository) GetAll() ([]*skill.ReadSkillDTO, error) {
	var skills []*models.Skill
	var skillDTOs []*skill.ReadSkillDTO
	repository.sqlClient.Joins("SkillCategory").Find(&skills)

	for _, s := range skills {
		skillDTOs = append(skillDTOs, &skill.ReadSkillDTO{
			ID:          s.ID,
			Name:        s.Name,
			Proficiency: s.Proficiency,
			Category: &skill_category.ReadSkillCategoryDTO{
				ID:   s.SkillCategoryId,
				Name: s.SkillCategory.Name,
			},
		})
	}

	return skillDTOs, nil
}

func (repository SkillRepository) GetById(id uint) (*skill.ReadSkillDTO, error) {
	var s *models.Skill

	repository.sqlClient.Joins("SkillCategory").Find(&s, id)

	if s.ID == 0 {
		return nil, errors.New("skill not found")
	}

	return &skill.ReadSkillDTO{
		ID:          s.ID,
		Name:        s.Name,
		Proficiency: s.Proficiency,
		Category: &skill_category.ReadSkillCategoryDTO{
			ID:   s.SkillCategoryId,
			Name: s.SkillCategory.Name,
		},
	}, nil
}

func (repository SkillRepository) Delete(id uint) error {
	var s models.Skill
	repository.sqlClient.Find(&s, id)
	if s.ID == 0 {
		return errors.New("skill not found")
	}
	repository.sqlClient.Delete(&s)
	return nil
}

func (repository SkillRepository) Update(id uint, dto *skill.UpdateSkillDTO) (*skill.ReadSkillDTO, error) {
	var s models.Skill

	repository.sqlClient.Find(&s, id)

	if s.ID == 0 {
		return nil, errors.New("skill not found")
	}

	utils.UpdateModelValuesFromDTO(&s, dto)
	repository.sqlClient.Save(&s)

	updatedSkill, _ := repository.GetById(s.ID)

	return updatedSkill, nil
}
