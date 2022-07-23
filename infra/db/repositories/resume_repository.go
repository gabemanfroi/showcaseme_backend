package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/DTO/resume"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/DTO/skill_category"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type ResumeRepository struct {
	sqlClient *gorm.DB
}

func CreateResumeRepository() *ResumeRepository {
	return &ResumeRepository{sqlClient: db.GetSqlInstance()}
}

func (repository ResumeRepository) GetByUsername(username string) (*resume.ReadResumeDTO, error) {
	var u *models.User
	var skills []*models.Skill
	var carouselItems []*models.CarouselItem

	repository.sqlClient.Where(&models.User{Username: username}).First(&u)

	if u.ID == 0 {
		return nil, errors.New("user not found")
	}

	repository.sqlClient.Joins("SkillCategory").Where(&models.Skill{UserId: u.ID}).Find(&skills)
	var skillDTOs []*skill.ReadSkillDTO

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

	repository.sqlClient.Order("position asc").Where(&models.CarouselItem{UserId: u.ID}).Find(&carouselItems)
	var carouselItemsDTOs []*carousel_item.ReadCarouselItemDTO

	for _, c := range carouselItems {
		carouselItemsDTOs = append(carouselItemsDTOs, &carousel_item.ReadCarouselItemDTO{
			ID:       c.ID,
			Content:  c.Content,
			Position: c.Position,
		})
	}
	return &resume.ReadResumeDTO{
		User: &user.ResumeUserDTO{
			ReadUserDTO: &user.ReadUserDTO{
				ID:        u.ID,
				Email:     u.Email,
				FirstName: u.FirstName,
				Username:  u.Username,
				LastName:  u.LastName,
				Role:      u.Role,
			},
			City:     u.City,
			Country:  u.Country,
			Pronouns: u.Pronouns,
		},
		Skills:        skillDTOs,
		CarouselItems: carouselItemsDTOs,
	}, nil

}
