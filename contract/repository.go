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
		kategoriLokasi int,
		tahunPenilaian int,
	) ([]entity.Lelang, error)
}
