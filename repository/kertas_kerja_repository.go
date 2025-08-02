package repository

import (
	"kertas_kerja/contract"
	"kertas_kerja/entity"
	"sort"
	"time"

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
	provinsi string,
) ([]entity.Lelang, error) {
	var hasil []entity.Lelang
	var existingKode []string

	tahunMin := tahunPembuatan - 5
	tahunMax := tahunPembuatan + 5
	tahunLelang := time.Now().Year()
	tahunLelangMin := tahunLelang - 3
	tahunLelangMax := tahunLelang

	// Helper untuk menambah hasil dan kode unik
	appendResult := func(data []entity.Lelang) {
		for _, h := range data {
			if !contains(existingKode, h.Kode) {
				hasil = append(hasil, h)
				existingKode = append(existingKode, h.Kode)
			}
		}
	}

	// Tahap 1
	var tahap1 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl = ?",
		merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi,
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&tahap1)
	appendResult(tahap1)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 2
	var tahap2 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi = ?",
		merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi, provinsi,
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap2)
	appendResult(tahap2)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 3
	var tahap3 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl = ?",
		merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi,
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap3)
	appendResult(tahap3)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 4
	var tahap4 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi = ?",
		merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi, provinsi,
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap4)
	appendResult(tahap4)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 5
	var tahap5 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi ILIKE ?",
		merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi, "%Jawa%",
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap5)
	appendResult(tahap5)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 6
	var tahap6 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi ILIKE ?",
		merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi, "%Jawa%",
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap6)
	appendResult(tahap6)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 7
	var tahap7 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ?",
		merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi,
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap7)
	appendResult(tahap7)
	if len(hasil) >= 3 {
		return limitHasil(hasil), nil
	}

	// Tahap 8
	var tahap8 []entity.Lelang
	r.db.Table("data_lelang").Where(
		"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ?",
		merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi,
	).Order("tahun_lelang DESC, harga_laku ASC").Limit(7 - len(hasil)).Find(&tahap8)
	appendResult(tahap8)

	return limitHasil(hasil), nil
}

// Helper untuk membatasi hasil sesuai aturan
func limitHasil(data []entity.Lelang) []entity.Lelang {
	if len(data) > 7 {
		// Ambil 7 data dengan harga_laku terendah dari tahap terakhir
		sort.Slice(data, func(i, j int) bool {
			return data[i].HargaLaku < data[j].HargaLaku
		})
		return data[:7]
	}
	return data
}

// Helper untuk cek kode unik
func contains(list []string, kode string) bool {
	for _, k := range list {
		if k == kode {
			return true
		}
	}
	return false
}

func (r *kertasKerjaRepo) FindDataLelangByKode(kode string) (*entity.Lelang, error) {
	var lelang entity.Lelang
	err := r.db.Table("data_lelang").Where("kode = ?", kode).First(&lelang).Error
	if err != nil {
		return nil, err
	}
	return &lelang, nil
}
