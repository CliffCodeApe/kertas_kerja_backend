package contract

import "kertas_kerja/dto"

type Service struct {
	KertasKerja KertasKerjaService
}

type KertasKerjaService interface {
	GetDataPembanding(req *dto.KertasKerjaRequest) (*dto.KertasKerjaResponse, error)
}
