package work_experience

import "time"

type ReadWorkExperienceDTO struct {
	ID          uint       `json:"id"`
	Role        string     `json:"role"`
	CompanyName string     `json:"companyName"`
	StartDate   time.Time  `json:"startDate"`
	Description string     `json:"description"`
	EndDate     *time.Time `json:"endDate"`
}
