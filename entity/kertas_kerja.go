package entity

import (
	"time"
)

// KertasKerja represents the structure of the kertas kerja entity.
type KertasKerja struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	UserID     uint64    `gorm:"user_id"`
	NamaObjek  string    `gorm:"nama_objek"`
	PdfPath    string    `gorm:"pdf_path"`
	ExcelPath  string    `gorm:"excel_path"`
	IsVerified bool      `gorm:"is_verified"`
	KodeKL     string    `gorm:"-"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
}

type Lelang struct {
	Kode           string `json:"kode"`
	Tipe           string `json:"tipe_lelang"`
	Merek          string `json:"merek"`
	TahunPembuatan int    `json:"tahun_pembuatan"`
	Kpknl          string `json:"kpknl"`
	KategoriLokasi int    `json:"kategori_lokasi"`
	HargaLaku      int64  `json:"harga_laku"`
	TahunLelang    int    `json:"tahun_lelang"`
}
