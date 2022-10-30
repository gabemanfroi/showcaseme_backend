package services

import "showcaseme/domain/DTO/resume"

type IResumeService interface {
	GetByUsername(username string) (*resume.ReadResumeDTO, error)
	Update(username string, dto *resume.UpdateResumeDTO) (*resume.ReadResumeDTO, error)
}
