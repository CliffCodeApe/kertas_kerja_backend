package contract

import "kertas_kerja/entity"

type Repository struct {
	KertasKerja KertasKerjaRepository
}

type KertasKerjaRepository interface {
	FindDataPembanding(
		merek string,
		tipe string,
		tahunPembuatan int,
		kpknl string,
		tahunLelang int,
	) ([]entity.Lelang, error)
	FindDataLelangByKode(kode string) (*entity.Lelang, error)
}
