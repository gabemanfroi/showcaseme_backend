package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
)

type CarouselItemService struct {
	repository repositories.ICarouselItemRepository
}

func CreateCarouselItemService() *CarouselItemService {
	return &CarouselItemService{repository: getCarouselItemRepository()}
}

func (service CarouselItemService) Create(dto *carousel_item.CreateCarouselItemDTO) (*carousel_item.ReadCarouselItemDTO, error) {
	return service.repository.Create(dto)
}

func (service CarouselItemService) GetAll() ([]*carousel_item.ReadCarouselItemDTO, error) {
	return service.repository.GetAll()
}

func (service CarouselItemService) GetById(id uint) (*carousel_item.ReadCarouselItemDTO, error) {
	return service.repository.GetById(id)
}

func (service CarouselItemService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service CarouselItemService) Update(id uint, dto *carousel_item.UpdateCarouselItemDTO) (*carousel_item.ReadCarouselItemDTO, error) {
	return service.repository.Update(id, dto)
}

func getCarouselItemRepository() repositories.ICarouselItemRepository {
	var injector repositories.ICarouselItemRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}
