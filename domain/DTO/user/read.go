package user

type ReadUserDTO struct {
	ID                uint   `json:"id"`
	Email             string `json:"email" `
	FirstName         string `json:"firstName"`
	Username          string `json:"username"`
	LastName          string `json:"lastName"`
	Role              string `json:"role"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
}
