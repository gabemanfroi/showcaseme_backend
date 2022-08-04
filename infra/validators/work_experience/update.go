package work_experience

import "time"

type UpdateWorkExperienceValidator struct {
	Role        string    `validate:"omitempty,min=25"`
	CompanyName string    `validate:"omitempty,min=25"`
	Description string    `validate:"omitempty,min=20"`
	EndDate     time.Time `validate:"omitempty"`
}
