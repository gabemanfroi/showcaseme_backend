package article

type UpdateArticleValidator struct {
	Title  string `validate:"omitempty,min=10,max=1000"`
	Content  string `validate:"omitempty,min=10"`
}
