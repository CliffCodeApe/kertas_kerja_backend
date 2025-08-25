package dto

import "time"

type DataLelang struct {
	Kode                          string    `json:"kode" binding:"required"`
	KategoriObjek                 string    `json:"kategori_objek" binding:"required"`
	TahunLelang                   int       `json:"tahun_lelang" binding:"required"`
	Provinsi                      string    `json:"provinsi" binding:"required"`
	Kota                          string    `json:"kota" binding:"required"`
	Kpknl                         string    `json:"kpknl" binding:"required"`
	KategoriLokasiJanuariDanJuni  int       `json:"kategori_lokasi_jan_jun" binding:"required"`
	KategoriLokasiJuliDanDesember int       `json:"kategori_lokasi_jul_des" binding:"required"`
	Merek                         string    `json:"merek" binding:"required"`
	Tipe                          string    `json:"tipe" binding:"required"`
	TahunPembuatan                int       `json:"tahun_pembuatan" binding:"required"`
	Warna                         string    `json:"warna" binding:"required"`
	HargaLaku                     float64   `json:"harga_laku" binding:"required"`
	CreatedAt                     time.Time `json:"created_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
}

type InsertDataLelangResponse struct {
	Status  string     `json:"status_code"`
	Message string     `json:"message"`
	Data    DataLelang `json:"data"`
}
