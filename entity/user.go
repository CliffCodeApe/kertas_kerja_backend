package entity

import "time"

type User struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	NamaSatker string    `gorm:"nama_satker"`
	KodeKL     string    `gorm:"kode_kl"`
	Email      string    `gorm:"email"`
	Role       string    `gorm:"column:role;default:'satker'"`
	Password   string    `gorm:"password"`
	IsVerified string    `gorm:"is_verified;default:menunggu"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
}
