package response

type CreatedArticle struct {
	Id uint64 `json:"id"`
}

// NewCreatedArticle recieves ArticleID and
// returns ArticleID mapped JSON response body.
func NewCreatedArticle(id uint64) *CreatedArticle {
	return &CreatedArticle{Id:id}
}
