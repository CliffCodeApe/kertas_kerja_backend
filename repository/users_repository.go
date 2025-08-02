package repository

import (
	"context"
	"kertas_kerja/contract"
	"kertas_kerja/entity"
	"time"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func implUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetByNamaSatker(namaSatker string) (*entity.User, error) {
	var user entity.User
	err := u.db.Table("users").First(&user, "nama_satker = ?", namaSatker).Error

	return &user, err
}

func (u *userRepo) GetById(id uint64) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user entity.User
	err := u.db.WithContext(ctx).First(&user, "id = ?", id).Error

	return &user, err
}
