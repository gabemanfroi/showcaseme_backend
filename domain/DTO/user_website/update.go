package user_website

type UpdateUserWebsiteDTO struct {
	Url  *string `json:"url,omitempty"`
	Type *string `json:"type,omitempty"`
}
