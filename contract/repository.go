package contract

import (
	"kertas_kerja/entity"
)

type Repository struct {
	KertasKerja KertasKerjaRepository
	User        UserRepository
	DataLelang  DataLelangRepository
}

type UserRepository interface {
	GetByNamaSatker(namaSatker string) (*entity.User, error)
	GetById(id uint64) (*entity.User, error)
	GetUsers() ([]*entity.User, error)
	ValidateUser(userID uint64) error
	InsertUser(user *entity.User) error
	InValidateUser(userID uint64) error
	DeleteUser(userID uint64) error
	ChangeUserRole(userID uint64, role string) error
}

type KertasKerjaRepository interface {
	FindDataPembanding(
		merek string,
		tipe string,
		tahunPembuatan int,
		kpknl string,
		provinsi string,
		tahap int,
	) ([]entity.Lelang, error)
	FindDataLelangByKode(kode string) (*entity.Lelang, error)
	GetAllRiwayatKertasKerja() ([]*entity.KertasKerja, error)
	GetRiwayatKertasKerjaByID(id uint64) (entity.KertasKerja, error)
	GetRiwayatKertasKerjaByUserID(userID uint64) ([]*entity.KertasKerja, error)
	InsertRiwayatKertasKerja(kk *entity.KertasKerja) error
	DeleteRiwayatKertasKerja(id uint64) error
	ValidasiKertasKerja(id uint64, pdfPath string) error
}

type DataLelangRepository interface {
	InsertDataLelang(dataLelang *entity.DataLelang) error
}
