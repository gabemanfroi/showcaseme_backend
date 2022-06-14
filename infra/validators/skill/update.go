package skill

type UpdateSkillValidator struct {
	Name        string `validate:"omitempty,min=3"`
	Proficiency uint8  `validate:"omitempty,gte=0,lte=100"`
}
