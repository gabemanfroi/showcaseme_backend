package skill

type CreateSkillDTO struct {
	Name            string `json:"name"`
	Proficiency     uint8  `json:"proficiency"`
	UserId          uint   `json:"userId"`
	SkillCategoryId uint   `json:"skillCategoryId"`
}
