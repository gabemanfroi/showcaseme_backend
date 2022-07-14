package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type SkillRepository struct {
	sqlClient *gorm.DB
}

func CreateSkillRepository() *SkillRepository {
	return &SkillRepository{sqlClient: db.GetSqlInstance()}
}

func (repository SkillRepository) Create(dto *skill.CreateSkillDTO) *models.Skill {
	s := models.Skill{
		UserId:      dto.UserId,
		Name:        dto.Name,
		Proficiency: dto.Proficiency,
	}
	repository.sqlClient.Create(&s)

	return &s
}

func (repository SkillRepository) GetAll() ([]*models.Skill, error) {
	var s []*models.Skill
	repository.sqlClient.Find(&s)
	return s, nil
}

func (repository SkillRepository) GetById(id uint) (*models.Skill, error) {
	var s *models.Skill
	repository.sqlClient.Find(&s, id)
	return s, nil
}

func (repository SkillRepository) Delete(id uint) error {
	var s models.Skill
	repository.sqlClient.Find(&s, id)
	if s.ID == 0 {
		return errors.New("skill not found")
	}
	s.Active = false
	repository.sqlClient.Save(&s)
	return nil
}

func (repository SkillRepository) Update(id uint, dto *skill.UpdateSkillDTO) (*models.Skill, error) {
	var s models.Skill
	repository.sqlClient.Find(&s, id)
	if s.ID == 0 {
		return &s, errors.New("skill not found")
	}
	updateSkillValuesFromDTO(&s, dto)
	repository.sqlClient.Save(&s)
	return &s, nil
}

func updateSkillValuesFromDTO(model *models.Skill, dto *skill.UpdateSkillDTO) {
	if dto.Name != nil {
		model.Name = *dto.Name
	}
	if dto.Proficiency != nil {
		model.Proficiency = *dto.Proficiency
	}
}
