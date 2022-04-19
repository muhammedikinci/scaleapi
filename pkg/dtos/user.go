package dtos

type LoginResponse struct {
	Status   bool   `json:"status"`
	Message  string `json:"message"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
