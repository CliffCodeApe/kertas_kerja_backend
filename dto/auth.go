package dto

type AuthLoginRequest struct {
	NamaSatker string `json:"nama_satker"`
	Password   string `json:"password"`
}

type AuthLoginResponse struct {
	StatusCode int           `json:"status"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	StatusCode  int    `json:"status"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

type AuthRegisterRequest struct {
	NamaSatker string `json:"nama_satker"`
	KodeKL     string `json:"kode_kl"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type AuthRegisterResponse struct {
	StatusCode int              `json:"status"`
	Message    string           `json:"message"`
	Data       RegisterResponse `json:"data"`
}

type RegisterResponse struct {
	NamaSatker string `json:"nama_satker"`
	KodeKL     string `json:"kode_kl"`
	Email      string `json:"email"`
}

type EmailRequest struct {
	Email string `json:"email"`
}

type ValidateUserResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type DeleteUserResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type ChangeUserRoleRequest struct {
	Role string `json:"role"`
}

type ChangeUserRoleResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type GetUserResponse struct {
	StatusCode int        `json:"status"`
	Message    string     `json:"message"`
	Data       []UserData `json:"data"`
}

type UserData struct {
	ID         uint64 `json:"id"`
	NamaSatker string `json:"nama_satker"`
	Email      string `json:"email"`
	KodeKL     string `json:"kode_kl"`
	Role       string `json:"role"`
	IsVerified string `json:"is_verified"`
}
