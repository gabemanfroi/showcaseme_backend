package services

import "showcaseme/domain/DTO/resume"

type IResumeService interface {
	GetByUsername(username string) (*resume.ReadResumeDTO, error)
}
