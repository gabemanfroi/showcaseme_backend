package user_website

type ReadUserWebsiteDTO struct {
	ID   uint   `json:"id"`
	Url  string `json:"url"`
	Type string `json:"type"`
}
