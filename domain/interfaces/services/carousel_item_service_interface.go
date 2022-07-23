package services

import "showcaseme/domain/DTO/carousel_item"

type ICarouselItemService interface {
	Create(dto *carousel_item.CreateCarouselItemDTO) (*carousel_item.ReadCarouselItemDTO, error)
	GetAll() ([]*carousel_item.ReadCarouselItemDTO, error)
	GetById(id uint) (*carousel_item.ReadCarouselItemDTO, error)
	Delete(id uint) error
	Update(id uint, dto *carousel_item.UpdateCarouselItemDTO) (*carousel_item.ReadCarouselItemDTO, error)
}
