package dtos

type ResponseMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
