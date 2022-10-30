package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/resume"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
)

type ResumeService struct {
	repository repositories.IResumeRepository
}

func CreateResumeService() *ResumeService { return &ResumeService{repository: getResumeRepository()} }

func getResumeRepository() repositories.IResumeRepository {
	var injector repositories.IResumeRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}

func (service ResumeService) GetByUsername(username string) (*resume.ReadResumeDTO, error) {
	return service.repository.GetByUsername(username)
}

func (service ResumeService) Update(username string, dto *resume.UpdateResumeDTO) (*resume.ReadResumeDTO, error) {
	return service.repository.Update(username, dto)
}
