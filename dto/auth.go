package dto

import "time"

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
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Nim       string    `json:"nim"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuthRegisterResponse struct {
	StatusCode int              `json:"status"`
	Message    string           `json:"message"`
	Data       RegisterResponse `json:"data"`
}

type RegisterResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Nim       string    `json:"nim"`
	CreatedAt time.Time `json:"createdAt"`
}

type EmailRequest struct {
	Email string `json:"email"`
}

type ResetPasswordResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type NewPassword struct {
	NewPassword string `json:"new_password"`
}

type EmailVerificationResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
