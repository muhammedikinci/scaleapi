package dtos

import "golang.org/x/crypto/bcrypt"

const ErrUsernameAndPasswordLengthError = "Username and password can be minimum 4 characters"

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
		return ErrUsernameAndPasswordLengthError, false
	}

	return "", true
}

func (rr *LoginRegisterRequest) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(rr.Password), 14)

	rr.Password = string(bytes)
}

func (rr LoginRegisterRequest) CheckPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rr.Password))
	return err == nil
}
