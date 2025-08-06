package contract

import "kertas_kerja/dto"

type Service struct {
	KertasKerja KertasKerjaService
	Auth        AuthService
}

type KertasKerjaService interface {
	GetDataPembanding(req *dto.KertasKerjaRequest, tahap int) (*dto.KertasKerjaResponse, error)
	GetDataLelangByKode(kode string) (*dto.DataPembandingResponse, error)
	SaveKertasKerjaToExcel(input *dto.KertasKerjaRequest, pembandingList *[]dto.DataPembanding) error
}

type AuthService interface {
	Login(payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	RefreshToken(token string) (*dto.RefreshTokenResponse, error)
}
