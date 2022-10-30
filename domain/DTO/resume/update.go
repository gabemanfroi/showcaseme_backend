package resume

import (
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/DTO/user"
)

type UpdateSkillDTOResume struct {
	*skill.UpdateSkillDTO
	ID           *uint  `json:"id"`
	CategoryName string `json:"categoryName"`
}

type UpdateResumeDTO struct {
	User   *user.UpdateUserDTO     `json:"user"`
	Skills *[]UpdateSkillDTOResume `json:"skills"`
}
