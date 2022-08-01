package user_website

type CreateUserWebsiteValidator struct {
	Url    string `validate:"required,min=3"`
	Type   string `validate:"required,min=3"`
	UserId uint   `validate:"required"`
}
