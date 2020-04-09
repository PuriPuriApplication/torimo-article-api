package response

type ResponseUser struct {
	ID                  uint64 `json:"id"`
	Name                string `json:"name"`
	ExternalServiceID   string `json:"externalServiceID"`
	ExternalServiceType string `json:"externalServiceType"`
	IsDeleted           bool   `json:"isDeleted"`
}
