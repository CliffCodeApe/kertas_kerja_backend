package service

import "kertas_kerja/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		KertasKerja: implKertasKerjaService(repo),
		Auth:        implAuthService(repo),
	}
}
