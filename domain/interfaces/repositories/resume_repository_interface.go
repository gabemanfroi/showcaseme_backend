package repositories

import "showcaseme/domain/DTO/resume"

type IResumeRepository interface {
	GetByUsername(username string) (*resume.ReadResumeDTO, error)
	Update(username string, dto *resume.UpdateResumeDTO) (*resume.ReadResumeDTO, error)
}
