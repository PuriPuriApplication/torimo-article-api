package response

type ResponseShop struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	StationID  uint64 `json:"stationId"`
	CreateUser uint64 `json:"createUser"`
	IsDeleted  bool   `json:"isDeleted"`
}
