package work_experience

import "time"

type CreateWorkExperienceDTO struct {
	UserId      uint      `json:"userId"`
	Role        string    `json:"role"`
	CompanyName string    `json:"companyName"`
	StartDate   time.Time `json:"startDate"`
}
