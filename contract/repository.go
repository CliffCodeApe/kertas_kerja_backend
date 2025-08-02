package contract

import (
	"kertas_kerja/entity"
)

type Repository struct {
	KertasKerja KertasKerjaRepository
	User        UserRepository
}

type UserRepository interface {
	GetByNamaSatker(namaSatker string) (*entity.User, error)
	GetById(id uint64) (*entity.User, error)
}

type KertasKerjaRepository interface {
	FindDataPembanding(
		merek string,
		tipe string,
		tahunPembuatan int,
		kpknl string,
		provinsi string,
	) ([]entity.Lelang, error)
	FindDataLelangByKode(kode string) (*entity.Lelang, error)
}
