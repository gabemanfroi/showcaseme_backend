package user

type ReadUserDTO struct {
	ID        uint   `json:"id"`
	Email     string `json:"email" `
	FirstName string `json:"firstName"`
	Username  string `json:"username"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

type ResumeUserDTO struct {
	*ReadUserDTO
	City     string `json:"city"`
	Country  string `json:"country"`
	Pronouns string `json:"pronouns"`
}
