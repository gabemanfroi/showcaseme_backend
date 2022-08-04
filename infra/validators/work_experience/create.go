package work_experience

import "time"

type CreateWorkExperienceValidator struct {
	UserId      uint      `validate:"required,gte=1"`
	CompanyName string    `validate:"required,min=10"`
	Role        string    `validate:"required,min=5"`
	StartDate   time.Time `validate:"required"`
	Description string    `validate:"omitempty,min=10"`
	EndDate     time.Time `validate:"omitempty"`
}
