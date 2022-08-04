package resume

import (
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/DTO/user_website"
)

type ReadResumeDTO struct {
	User          *user.ResumeUserDTO                  `json:"user"`
	Skills        []*skill.ReadSkillDTO                `json:"skills"`
	CarouselItems []*carousel_item.ReadCarouselItemDTO `json:"carouselItems"`
	Websites      []*user_website.ReadUserWebsiteDTO   `json:"websites"'`
	Projects      []*project.ReadProjectDTO            `json:"projects"`
}
