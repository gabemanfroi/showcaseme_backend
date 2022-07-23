package carousel_item

type UpdateCarouselItemDTO struct {
	Content  *string `json:"content,omitempty"`
	Position *uint   `json:"position,omitempty"`
}
