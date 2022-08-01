package article

type CreateArticleDTO struct {
    UserId  uint `json:"userId"`
    Title  string `json:"title"`
    Content  string `json:"content"`
}
