package resume

import (
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/DTO/user_website"
)

type ReadUserDTOResume struct {
	*user.ReadUserDTO
	City     string `json:"city"`
	Country  string `json:"country"`
	Pronouns string `json:"pronouns"`
}

type ReadSkillDTOResume struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Proficiency  uint8  `json:"proficiency"`
	CategoryName string `json:"categoryName"`
}

type ReadResumeDTO struct {
	User          *ReadUserDTOResume                   `json:"user"`
	Skills        []*ReadSkillDTOResume                `json:"skills"`
	CarouselItems []*carousel_item.ReadCarouselItemDTO `json:"carouselItems"`
	Websites      []*user_website.ReadUserWebsiteDTO   `json:"websites"'`
	Projects      []*project.ReadProjectDTO            `json:"projects"`
}
