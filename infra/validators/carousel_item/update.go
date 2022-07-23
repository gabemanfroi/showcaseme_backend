package carousel_item

type UpdateCarouselItemValidator struct {
	Content  string `validate:"omitempty,min=3"`
	Position uint   `validate:"omitempty,gte=0"`
	UserId   uint8  `validate:"omitempty,gte=0"`
}
