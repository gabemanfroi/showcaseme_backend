package article

type CreateArticleValidator struct {
	UserId  uint   `validate:"required,gte=1"`
	Title   string `validate:"required,min=10"`
	Content string `validate:"required,min=10"`
}
