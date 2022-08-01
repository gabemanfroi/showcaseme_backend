package user_website

type UpdateUserWebsiteValidator struct {
	Url  string `validate:"omitempty,min=3"`
	Type string `validate:"omitempty,min=3"`
}
