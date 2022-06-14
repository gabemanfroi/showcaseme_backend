package skill

type CreateSkillDTO struct {
	Name        string `json:"name"`
	Proficiency uint8  `json:"proficiency"`
	UserId      uint8  `json:"userId"`
}
