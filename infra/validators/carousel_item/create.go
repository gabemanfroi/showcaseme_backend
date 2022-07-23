package carousel_item

type CreateCarouselItemValidator struct {
	Content  string `validate:"required,min=3"`
	Position uint   `validate:"gte=0"`
	UserId   uint8  `validate:"required,gte=0"`
}
