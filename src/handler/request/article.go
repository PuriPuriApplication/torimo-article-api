package request

type RequestArticle struct {
	Title string           `json:"title"  validate:"required,max=100"`
	Body string            `json:"body"   validate:"required"`
	Status string          `json:"status" validate:"required,max=16"`
	UserID int64           `json:"userId" validate:"required"`
	ShopID string          `json:"shopId"`
	CategoryIDs []int      `json:"categoryIds"`
}