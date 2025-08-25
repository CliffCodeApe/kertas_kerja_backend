package repository

import (
	"kertas_kerja/contract"
	"kertas_kerja/entity"

	"gorm.io/gorm"
)

type dataLelangRepo struct {
	db *gorm.DB
}

func implDataLelangRepository(db *gorm.DB) contract.DataLelangRepository {
	return &dataLelangRepo{
		db: db,
	}
}

func (r *dataLelangRepo) InsertDataLelang(dataLelang *entity.DataLelang) error {
	return r.db.Table("data_lelang").Create(dataLelang).Error
}
