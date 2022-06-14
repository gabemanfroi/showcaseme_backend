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

func (repository SkillRepository) Create(dto *skill.CreateSkillDTO) models.Skill {
	skillToBeCreated := models.Skill{
		UserId:      dto.UserId,
		Name:        dto.Name,
		Proficiency: dto.Proficiency,
	}
	repository.sqlClient.Create(&skillToBeCreated)

	return skillToBeCreated
}

func (repository SkillRepository) GetAll() ([]models.Skill, error) {
	var retrievedSkills []models.Skill
	repository.sqlClient.Find(&retrievedSkills)
	return retrievedSkills, nil
}

func (repository SkillRepository) GetById(id string) (models.Skill, error) {
	var retrievedSkill models.Skill
	repository.sqlClient.Find(&retrievedSkill, id)
	return retrievedSkill, nil
}

func (repository SkillRepository) Delete(id string) error {
	var skillToDelete models.Skill
	repository.sqlClient.Find(&skillToDelete, id)
	if skillToDelete.ID == 0 {
		return errors.New("skill not found")
	}
	skillToDelete.Active = false
	repository.sqlClient.Save(&skillToDelete)
	return nil
}

func (repository SkillRepository) Update(id string, dto *skill.UpdateSkillDTO) (models.Skill, error) {
	var skillToBeUpdated models.Skill
	repository.sqlClient.Find(&skillToBeUpdated, id)
	if skillToBeUpdated.ID == 0 {
		return skillToBeUpdated, errors.New("skill not found")
	}
	updateSkillValuesFromDTO(&skillToBeUpdated, dto)
	repository.sqlClient.Save(&skillToBeUpdated)
	return skillToBeUpdated, nil
}

func updateSkillValuesFromDTO(model *models.Skill, dto *skill.UpdateSkillDTO) {
	if dto.Name != nil {
		model.Name = *dto.Name
	}
	if dto.Proficiency != nil {
		model.Proficiency = *dto.Proficiency
	}
}
