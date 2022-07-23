package skill_category

type CreateSkillCategoryValidator struct {
	Name   string `validate:"required,min=3"`
	UserId uint   `validate:"omitempty,gte=0"`
}
