package article

type UpdateArticleDTO struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
