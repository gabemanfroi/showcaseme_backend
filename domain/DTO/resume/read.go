package resume

import (
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/DTO/user"
)

type ReadResumeDTO struct {
	User          *user.ResumeUserDTO                  `json:"user"`
	Skills        []*skill.ReadSkillDTO                `json:"skills"`
	CarouselItems []*carousel_item.ReadCarouselItemDTO `json:"carouselItems"`
}
