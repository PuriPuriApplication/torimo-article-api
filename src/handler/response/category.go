package response

type ResponseCategory struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	CreateUser uint64 `json:"createUser"`
	IsDeleted  bool   `json:"isDeleted"`
}
