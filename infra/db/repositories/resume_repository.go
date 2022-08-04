package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/DTO/project_category"
	"showcaseme/domain/DTO/resume"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/DTO/skill_category"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/DTO/user_website"
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

	repository.sqlClient.Where(&models.User{Username: username}).First(&u)

	if u.ID == 0 {
		return nil, errors.New("user not found")
	}

	skillDTOs := repository.getUserSkills(u)
	carouselItemsDTOs := repository.getUserCarouselItems(u)
	websitesDTOs := repository.getUserWebsites(u)
	projectsDTOs := repository.getUserProjects(u)

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
		Websites:      websitesDTOs,
		Projects:      projectsDTOs,
	}, nil

}

func (repository ResumeRepository) getUserProjects(u *models.User) []*project.ReadProjectDTO {
	var projects []*models.Project
	var projectsDTOs []*project.ReadProjectDTO
	repository.sqlClient.Joins("Project").Where(&models.Project{UserId: u.ID}).Find(&projects)

	for _, p := range projects {
		projectsDTOs = append(projectsDTOs, &project.ReadProjectDTO{
			ID: p.ID,
			ProjectCategory: &project_category.ReadProjectCategoryDTO{
				ID:   p.ProjectCategoryId,
				Name: p.ProjectCategory.Name,
			},
			ImageUrl: p.ImageUrl,
			Url:      p.Url,
			Title:    p.Title,
		})
	}
	return projectsDTOs
}

func (repository ResumeRepository) getUserWebsites(u *models.User) []*user_website.ReadUserWebsiteDTO {
	var websites []*models.UserWebsite
	repository.sqlClient.Where(&models.UserWebsite{UserId: u.ID}).Find(&websites)
	var websitesDTOs []*user_website.ReadUserWebsiteDTO

	for _, w := range websites {
		websitesDTOs = append(websitesDTOs, &user_website.ReadUserWebsiteDTO{
			ID:   w.ID,
			Url:  w.Url,
			Type: w.Type,
		})
	}
	return websitesDTOs
}

func (repository ResumeRepository) getUserCarouselItems(u *models.User) []*carousel_item.ReadCarouselItemDTO {
	var carouselItems []*models.CarouselItem
	repository.sqlClient.Order("position asc").Where(&models.CarouselItem{UserId: u.ID}).Find(&carouselItems)
	var carouselItemsDTOs []*carousel_item.ReadCarouselItemDTO

	for _, c := range carouselItems {
		carouselItemsDTOs = append(carouselItemsDTOs, &carousel_item.ReadCarouselItemDTO{
			ID:       c.ID,
			Content:  c.Content,
			Position: c.Position,
		})
	}
	return carouselItemsDTOs
}

func (repository ResumeRepository) getUserSkills(u *models.User) []*skill.ReadSkillDTO {
	var skills []*models.Skill
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
	return skillDTOs
}
