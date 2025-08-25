package entity

import "time"

type DataLelang struct {
	Kode                 string    `gorm:"column:kode" json:"kode"`
	KategoriObjek        string    `gorm:"column:kategori_objek" json:"kategori_objek"`
	TahunLelang          int       `gorm:"column:tahun_lelang" json:"tahun_lelang"`
	Provinsi             string    `gorm:"column:provinsi" json:"provinsi"`
	Kota                 string    `gorm:"column:kota" json:"kota"`
	Kpknl                string    `gorm:"column:kpknl" json:"kpknl"`
	KategoriLokasiJanJun int       `gorm:"column:kategori_lokasi_jan_jun" json:"kategori_lokasi_jan_jun"`
	KategoriLokasiJulDes int       `gorm:"column:kategori_lokasi_jul_des" json:"kategori_lokasi_jul_des"`
	Merek                string    `gorm:"column:merek" json:"merek"`
	Tipe                 string    `gorm:"column:tipe" json:"tipe"`
	TahunPembuatan       int       `gorm:"column:tahun_pembuatan" json:"tahun_pembuatan"`
	Warna                string    `gorm:"column:warna" json:"warna"`
	HargaLaku            float64   `gorm:"column:harga_laku" json:"harga_laku"`
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
}
