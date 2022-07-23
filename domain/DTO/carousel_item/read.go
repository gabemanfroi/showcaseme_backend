package carousel_item

type ReadCarouselItemDTO struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Position uint   `json:"position"`
}
