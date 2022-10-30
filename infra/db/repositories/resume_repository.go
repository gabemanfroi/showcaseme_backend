package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/DTO/project_category"
	"showcaseme/domain/DTO/resume"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/DTO/user_website"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
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
		User: &resume.ReadUserDTOResume{
			ReadUserDTO: &user.ReadUserDTO{
				ID:                u.ID,
				Email:             u.Email,
				FirstName:         u.FirstName,
				Username:          u.Username,
				LastName:          u.LastName,
				Role:              u.Role,
				ProfilePictureUrl: u.ProfilePictureURL,
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
	repository.sqlClient.Where(&models.Project{UserId: u.ID}).Find(&projects)

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

	if len(websites) == 0 {
		return make([]*user_website.ReadUserWebsiteDTO, 0)
	}

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

	if len(carouselItems) == 0 {
		return make([]*carousel_item.ReadCarouselItemDTO, 0)
	}

	for _, c := range carouselItems {
		carouselItemsDTOs = append(carouselItemsDTOs, &carousel_item.ReadCarouselItemDTO{
			ID:       c.ID,
			Content:  c.Content,
			Position: c.Position,
		})
	}
	return carouselItemsDTOs
}

func (repository ResumeRepository) getUserSkills(u *models.User) []*resume.ReadSkillDTOResume {
	var skills []*models.Skill
	repository.sqlClient.Joins("SkillCategory").Where(&models.Skill{UserId: u.ID}).Find(&skills)
	var skillDTOs []*resume.ReadSkillDTOResume

	if len(skills) == 0 {
		return make([]*resume.ReadSkillDTOResume, 0)
	}

	for _, s := range skills {
		skillDTOs = append(skillDTOs, &resume.ReadSkillDTOResume{
			ID:           s.ID,
			Name:         s.Name,
			Proficiency:  s.Proficiency,
			CategoryName: s.SkillCategory.Name,
		})
	}

	return skillDTOs
}

func (repository ResumeRepository) Update(username string, dto *resume.UpdateResumeDTO) (*resume.ReadResumeDTO, error) {
	var u *models.User

	repository.sqlClient.Where(&models.User{Username: username}).First(&u)
	utils.UpdateModelValuesFromDTO(u, dto.User)

	repository.createNewSkills(dto, u)

	repository.sqlClient.Save(&u)
	return repository.GetByUsername(username)
}

func (repository ResumeRepository) createNewSkills(dto *resume.UpdateResumeDTO, u *models.User) {
	for _, s := range *dto.Skills {
		if s.ID == nil {
			skill := &models.Skill{
				Name:        *s.Name,
				Proficiency: *s.Proficiency,
				UserId:      u.ID,
				Show:        true,
			}
			if s.CategoryName == "Languages" {
				var sc *models.SkillCategory
				repository.sqlClient.Where(&models.SkillCategory{Name: "Languages"}).First(&sc)
				skill.SkillCategoryId = sc.ID
			}
			if s.CategoryName == "Programming Languages" {
				var sc *models.SkillCategory
				repository.sqlClient.Where(&models.SkillCategory{Name: "Programming Languages"}).First(&sc)
				skill.SkillCategoryId = sc.ID
			}
			repository.sqlClient.Create(&skill)
		}
	}
}
