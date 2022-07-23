package skill

type CreateSkillValidator struct {
	Name            string `validate:"required,min=3"`
	Proficiency     uint8  `validate:"required,gte=0,lte=100"`
	UserId          uint8  `validate:"required,gte=0"`
	SkillCategoryId uint8  `validate:"required,gte=0"`
}
