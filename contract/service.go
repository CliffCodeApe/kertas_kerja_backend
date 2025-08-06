package contract

import (
	"kertas_kerja/dto"
	"kertas_kerja/entity"
)

type Service struct {
	KertasKerja KertasKerjaService
	Auth        AuthService
	User        UserService
	Mail        MailService
}

type KertasKerjaService interface {
	GetDataPembanding(req *dto.KertasKerjaRequest, tahap int) (*dto.KertasKerjaResponse, error)
	GetDataLelangByKode(kode string) (*dto.DataPembandingResponse, error)
	SaveKertasKerjaToExcel(input *dto.KertasKerjaRequest, pembandingList *[]dto.DataPembanding) error
}

type AuthService interface {
	Register(payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
	Login(payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	RefreshToken(token string) (*dto.RefreshTokenResponse, error)
}

type UserService interface {
	ValidateUser(userID uint64) (*dto.ValidateUserResponse, error)
	GetAllUsers() (*dto.GetUserResponse, error)
}

type MailService interface {
	Enqueue(email entity.Mail)
	Dequeue()
}
