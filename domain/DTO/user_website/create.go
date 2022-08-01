package user_website

type CreateUserWebsiteDTO struct {
	Url    string `json:"url"`
	UserId uint   `json:"userId"`
	Type   string `json:"type"`
}
