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

func (lr LoginRegisterRequest) Validate() (string, bool) {
	if len(lr.Username) < 3 || len(lr.Password) < 3 {
		return "Username and password can be minimum 4 characters", false
	}

	return "", true
}
