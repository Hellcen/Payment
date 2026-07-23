package response

type RegisterResponseDTO struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	Status   string `json:"status"`
}
