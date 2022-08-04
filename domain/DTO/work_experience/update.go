package work_experience

import "time"

type UpdateWorkExperienceDTO struct {
	Role        *string    `json:"role"`
	CompanyName *string    `json:"companyName"`
	EndDate     *time.Time `json:"endDate"`
	Description *string    `json:"description"`
	StartDate   *time.Time `json:"startDate"`
}
