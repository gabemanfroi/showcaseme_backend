package skill

type UpdateSkillDTO struct {
	Name            *string `json:"name,omitempty"`
	Proficiency     *uint8  `json:"proficiency,omitempty"`
	SkillCategoryId *uint   `json:"skillCategoryId"`
}
