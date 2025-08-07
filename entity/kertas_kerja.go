package entity

import (
	"time"
)

// KertasKerja represents the structure of the kertas kerja entity.
type KertasKerja struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	UserID    uint64    `gorm:"user_id"`
	NamaObjek string    `gorm:"nama_objek"`
	PdfPath   string    `gorm:"pdf_path"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type Lelang struct {
	Kode           string
	Tipe           string
	Merek          string
	TahunPembuatan int
	Kpknl          string
	KategoriLokasi int
	HargaLaku      int64
	TahunLelang    int
}
