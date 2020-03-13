package request

type CreateArticleRequest struct {
	Title       string   `json:"title"  validate:"required,max=100"`
	Body        string   `json:"body"   validate:"required"`
	Status      string   `json:"status" validate:"required,max=16"`
	UserID      uint64   `json:"userId" validate:"required"`
	ShopID      uint64   `json:"shopId"`
	CategoryIDs []uint64 `json:"categoryIds"`
}
