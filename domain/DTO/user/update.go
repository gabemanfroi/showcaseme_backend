package user

type UpdateUserDTO struct {
	Age       *int    `json:"age"`
	City      *string `json:"city" `
	Country   *string `json:"country" `
	Email     *string `json:"email"  `
	FirstName *string `json:"firstName" `
	LastName  *string `json:"lastName"  `
	Pronouns  *string `json:"pronouns"`
}
