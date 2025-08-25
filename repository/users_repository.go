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

func (u *userRepo) GetUsers() ([]*entity.User, error) {
	var users []*entity.User
	err := u.db.Table("users").Select("id, nama_satker, email, kode_kl, role, is_verified, created_at, updated_at").Where("role IN ?", []string{"satker", "admin"}).Find(&users).Error
	return users, err
}

func (u *userRepo) ValidateUser(userID uint64) error {
	return u.db.Table("users").Where("id = ?", userID).Update("is_verified", "tervalidasi").Error
}

func (u *userRepo) InsertUser(user *entity.User) error {
	err := u.db.Table("users").Create(&user).Error
	return err
}

func (u *userRepo) InValidateUser(userID uint64) error {
	return u.db.Table("users").Where("id = ?", userID).Update("is_verified", "ditolak").Error
}

func (u *userRepo) DeleteUser(userID uint64) error {
	return u.db.Table("users").Where("id = ?", userID).Delete(&entity.User{}).Error
}

func (u *userRepo) ChangeUserRole(userID uint64, role string) error {
	return u.db.Table("users").Where("id = ?", userID).Update("role", role).Error

}
