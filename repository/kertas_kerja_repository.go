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
	tahunLelang int,
) ([]entity.Lelang, error) {
	var hasil []entity.Lelang

	// Range tahun pembuatan ±5 tahun
	tahunMin := tahunPembuatan - 5
	tahunMax := tahunPembuatan + 5

	// Ambil tahun lelang terbaru dari database untuk filter range tahun lelang
	// var tahunLelangMax *int
	// r.db.Table("data_lelang").
	// 	Select("MAX(tahun_lelang)").
	// 	Where("merek = ? AND tipe = ? AND tahun_pembuatan BETWEEN ? AND ? AND kpknl = ?", merek, tipe, tahunMin, tahunMax, lokasi).
	// 	Scan(&tahunLelangMax)

	// if tahunLelangMax == nil {
	// 	return hasil, fmt.Errorf("tahun Lelang tidak ditemukan untuk merek %s, tipe %s, tahun pembuatan %d, lokasi %s", merek, tipe, tahunPembuatan, lokasi)
	// }

	fmt.Printf("tahun lelang: %d\n", tahunLelang)

	tahunLelangMin := tahunLelang - 3
	tahunLelangMax := tahunLelang + 3
	// Tahap 1: Query lokasi sama, simpan ke hasil
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe = ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl = ?",
		merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi,
	).Order("tahun_lelang DESC, harga_laku ASC").
		Limit(7).Find(&hasil)

	// Jika hasil kurang dari 7, lanjut tahap 2
	if len(hasil) < 7 {
		var tambahan []entity.Lelang
		var existingKode []string
		for _, h := range hasil {
			existingKode = append(existingKode, h.Kode)
		}

		query := r.db.Table("data_lelang").Where(
			"merek = ? AND tipe = ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ?",
			merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi,
		)
		if len(existingKode) > 0 {
			query = query.Where("kode NOT IN ?", existingKode)
		}
		query.Order("tahun_lelang DESC, harga_laku ASC").
			Limit(7 - len(hasil)).Find(&tambahan)

		// Gabungkan hasil tahap 1 dan tahap 2
		hasil = append(hasil, tambahan...)
	}

	// Batasi maksimal 7 data
	if len(hasil) > 7 {
		hasil = hasil[:7]
	}

	fmt.Printf("Data pembanding ditemukan: %d\n", len(hasil))

	return hasil, nil
}

func (r *kertasKerjaRepo) FindDataLelangByKode(kode string) (*entity.Lelang, error) {
	var lelang entity.Lelang
	err := r.db.Table("data_lelang").Where("kode = ?", kode).First(&lelang).Error
	if err != nil {
		return nil, err
	}
	return &lelang, nil
}
