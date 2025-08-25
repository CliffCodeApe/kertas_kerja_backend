package repository

import (
	"fmt"
	"kertas_kerja/contract"
	"kertas_kerja/entity"
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
	tahap int,
) ([]entity.Lelang, error) {
	var hasil []entity.Lelang
	tahunMin := tahunPembuatan - 5
	tahunMax := tahunPembuatan + 5
	tahunLelang := time.Now().Year()
	tahunLelangMin := tahunLelang - 3
	tahunLelangMax := tahunLelang

	switch tahap {
	case 1:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl ILIKE ?",
			merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, "%"+lokasi+"%",
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 2:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi = ?",
			merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi, provinsi,
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 3:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe = ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl ILIKE ?",
			merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, "%"+lokasi+"%",
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 4:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi = ?",
			merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi, provinsi,
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 5:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi ILIKE ?",
			merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi, "%"+provinsi+"%",
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 6:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ? AND provinsi ILIKE ?",
			merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi, "%"+provinsi+"%",
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 7:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe = ? AND tahun_pembuatan = ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ?",
			merek, tipe, tahunPembuatan, tahunLelangMin, tahunLelangMax, lokasi,
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	case 8:
		r.db.Table("data_lelang").Where(
			"merek = ? AND tipe != ? AND tahun_pembuatan BETWEEN ? AND ? AND tahun_lelang BETWEEN ? AND ? AND kpknl != ?",
			merek, tipe, tahunMin, tahunMax, tahunLelangMin, tahunLelangMax, lokasi,
		).Order("tahun_lelang DESC, harga_laku ASC").Limit(7).Find(&hasil)
	}

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

func (r *kertasKerjaRepo) GetRiwayatKertasKerjaByUserID(userID uint64) ([]*entity.KertasKerja, error) {

	var riwayat []entity.KertasKerja
	err := r.db.Table("kertas_kerja").Where("user_id = ?", userID).Find(&riwayat).Error
	if err != nil {
		return nil, err
	}

	result := make([]*entity.KertasKerja, len(riwayat))
	for i, item := range riwayat {
		result[i] = &item
	}

	return result, nil
}

func (r *kertasKerjaRepo) GetRiwayatKertasKerjaByID(id uint64) (entity.KertasKerja, error) {

	var riwayat entity.KertasKerja
	err := r.db.Table("kertas_kerja").Where("id = ?", id).First(&riwayat).Error
	if err != nil {
		return entity.KertasKerja{}, err
	}

	return riwayat, nil
}

func (r *kertasKerjaRepo) GetAllRiwayatKertasKerja() ([]*entity.KertasKerja, error) {
	var result []*entity.KertasKerja
	err := r.db.Table("kertas_kerja AS k").
		Where("u.role = ?", "satker").
		Select("k.id, k.nama_objek, k.nup, k.kode_satker, k.hasil_nilai_taksiran, k.excel_path, k.pdf_path, k.is_verified, k.created_at, k.updated_at").
		Joins("JOIN users u ON u.id = k.user_id").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *kertasKerjaRepo) InsertRiwayatKertasKerja(kk *entity.KertasKerja) error {
	return r.db.Table("kertas_kerja").Create(kk).Error
}

func (r *kertasKerjaRepo) DeleteRiwayatKertasKerja(id uint64) error {
	return r.db.Table("kertas_kerja").Where("id = ?", id).Delete(&entity.KertasKerja{}).Error
}

func (r *kertasKerjaRepo) ValidasiKertasKerja(id uint64, pdfPath string) error {
	var kertasKerja entity.KertasKerja
	err := r.db.Table("kertas_kerja").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_verified": true,
			"pdf_path":    pdfPath,
		}).Error
	if err != nil {
		return fmt.Errorf("data kertas kerja tidak ditemukan: %w", err)
	}

	// Lakukan validasi terhadap data kertas kerja
	if kertasKerja.IsVerified {
		return fmt.Errorf("kertas kerja sudah diverifikasi")
	}

	return nil
}
