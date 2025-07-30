package entity

import (
	"time"
)

// KertasKerja represents the structure of the kertas kerja entity.
type KertasKerja struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Deskripsi string    `json:"deskripsi"`
	NamaObjek string    `json:"nama_objek"`
	FilePath  string    `json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
