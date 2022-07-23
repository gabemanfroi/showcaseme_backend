package skill_category

type UpdateSkillCategoryValidator struct {
	Name string `validate:"omitempty,min=3"`
}
