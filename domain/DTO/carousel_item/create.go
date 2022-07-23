package carousel_item

type CreateCarouselItemDTO struct {
	Content  string `json:"content"`
	UserId   uint   `json:"userId"`
	Position uint   `json:"position"`
}
