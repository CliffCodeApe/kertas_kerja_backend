package repository

import (
	"kertas_kerja/contract"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		KertasKerja: implKertasKerjaRepository(db),
	}
}
