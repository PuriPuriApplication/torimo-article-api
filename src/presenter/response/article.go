package response

type ResponseArticle struct {
	ID         uint64              `json:"id"`
	Title      string              `json:"title"`
	Body       string              `json:"body"`
	Status     string              `json:"status"`
	User       *ResponseUser       `json:"user"`
	Shop       *ResponseShop       `json:"shop"`
	Categories []*ResponseCategory `json:"categories"`
}
