package skill

import "showcaseme/domain/DTO/skill_category"

type ReadSkillDTO struct {
	ID          uint                                 `json:"id"`
	Name        string                               `json:"name"`
	Proficiency uint8                                `json:"proficiency"`
	Category    *skill_category.ReadSkillCategoryDTO `json:"category"`
}
