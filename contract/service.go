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
	DataLelang  DataLelangService
}

type KertasKerjaService interface {
	GetDataPembanding(req *dto.IdentitasKendaraan, tahap int) (*dto.KertasKerjaResponse, error)
	GetDataLelangByKode(kode string) (*dto.DataPembandingResponse, error)
	GetAllRiwayatKertasKerja() (*dto.GetAllRiwayatKertasKerjaResponse, error)
	GetRiwayatKertasKerjaByUserID(userID uint64) (*dto.GetAllRiwayatKertasKerjaResponse, error)
	GetRiwayatKertasKerjaByID(id uint64) (*dto.RiwayatKertasKerjaData, error)
	SaveKertasKerja(payload *dto.IsiKertasKerjaRequest, userID uint64) (*dto.RiwayatKertasKerjaResponse, error)
	DeleteRiwayatKertasKerja(id uint64) (*dto.DeleteRiwayatKertasKerjaResponse, error)
	ValidasiKertasKerja(id uint64, pdfPath string) (*dto.ValidasiKertasKerjaResponse, error)
}

type AuthService interface {
	Register(payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
	Login(payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	RefreshToken(token string) (*dto.RefreshTokenResponse, error)
}

type UserService interface {
	ValidateUser(userID uint64) (*dto.ValidateUserResponse, error)
	GetAllUsers() (*dto.GetUserResponse, error)
	InValidateUser(userID uint64) (*dto.ValidateUserResponse, error)
	DeleteUser(userID uint64) (*dto.DeleteUserResponse, error)
	ChangeUserRole(payload *dto.ChangeUserRoleRequest, userID uint64) (*dto.ChangeUserRoleResponse, error)
}

type DataLelangService interface {
	InsertDataLelang(payload *dto.DataLelang) (*dto.InsertDataLelangResponse, error)
}

type MailService interface {
	Enqueue(email entity.Mail)
	Dequeue()
}
