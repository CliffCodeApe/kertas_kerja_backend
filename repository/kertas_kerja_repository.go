package repository

import (
	"fmt"
	"kertas_kerja/contract"
	"kertas_kerja/entity"

	"gorm.io/gorm"
)

type kertasKerjaRepo struct {
	db *gorm.DB
}

func implKertasKerjaRepository(db *gorm.DB) contract.KertasKerjaRepository {
	return &kertasKerjaRepo{
		db: db,
	}
}

func (r *kertasKerjaRepo) FindDataPembanding(
	merek string,
	tipe string,
	tahunPembuatan int,
	lokasi string,
) ([]entity.Lelang, error) {
	var hasil []entity.Lelang

	// Range tahun pembuatan ±5 tahun
	tahunMin := tahunPembuatan - 5
	tahunMax := tahunPembuatan + 5

	// Ambil tahun lelang terbaru dari database untuk filter range tahun lelang
	var tahunLelangMax int
	r.db.Table("data_lelang").
		Select("MAX(tahun_lelang)").
		Where("merek = ? AND tipe = ? AND tahun_pembuatan BETWEEN ? AND ? AND kpknl = ?", merek, tipe, tahunMin, tahunMax, lokasi).
		Scan(&tahunLelangMax)

	// Jika tidak ada data, return hasil kosong
	if tahunLelangMax == 0 {
		return hasil, nil
	}

	tahunLelangMin := tahunLelangMax - 3

	// Tahap 1: Lokasi sama
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe = ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl = ?",
		merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi,
	).Order("tahun_lelang DESC, harga_laku ASC").
		Limit(7).Find(&hasil)

	// if len(hasil) < 3 {
	// 	// Tahap 2: Lokasi berbeda
	// 	var tambahan []entity.Lelang
	// 	r.db.Where(
	// 		"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND kpknl != ?",
	// 		merek, tipe, tahunPembuatan, lokasi,
	// 	).Order("tahun_lelang DESC, harga_laku ASC").
	// 		Limit(10).Find(&tambahan)
	// 	hasil = append(hasil, tambahan...)
	// }

	// Batasi maksimal 7 data
	if len(hasil) > 7 {
		hasil = hasil[:7]
	}

	fmt.Printf("PARAMS: merek=%s, tipe=%s, tahun_pembuatan=%d, kpknl=%s\n", merek, tipe, tahunPembuatan, lokasi)
	fmt.Printf("%+v\n", hasil)

	return hasil, nil
}
