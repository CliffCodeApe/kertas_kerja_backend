package service

import (
	"fmt"
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/entity"
	"kertas_kerja/pkg/errs"
	"kertas_kerja/pkg/mail"
	"log"
	"net/http"
	"os"
	"strconv"
)

type userServ struct {
	userRepo    contract.UserRepository
	mailService contract.MailService
}

func implUserService(repo *contract.Repository) contract.UserService {
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}
	return &userServ{
		userRepo: repo.User,
		mailService: mail.ImplMailService(
			os.Getenv("SMTP_FROM"),
			os.Getenv("SMTP_PASSWORD"),
			os.Getenv("SMTP_HOST"),
			os.Getenv("BASE_URL"),
			smtpPort,
		),
	}
}

func (a *userServ) ValidateUser(userID uint64) (*dto.ValidateUserResponse, error) {
	err := a.userRepo.ValidateUser(userID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.ValidateUser fail: %w", err)
	}

	user, err := a.userRepo.GetById(userID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}

	a.mailService.Enqueue(entity.Mail{
		To:      user.Email,
		Subject: "Validasi Satker",
		Body:    fmt.Sprintf("User anda sudah diverifikasi oleh Super Admin, silahkan login. Terima Kasih"),
	})

	response := &dto.ValidateUserResponse{
		StatusCode: http.StatusOK,
		Message:    "User successfully validated",
	}

	return response, nil
}

func (a *userServ) GetAllUsers() (*dto.GetUserResponse, error) {
	users, err := a.userRepo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("userRepo.GetUsers fail: %w", err)
	}

	var userData []dto.UserData
	for _, user := range users {
		userData = append(userData, dto.UserData{
			ID:         user.ID,
			NamaSatker: user.NamaSatker,
			Email:      user.Email,
			KodeKL:     user.KodeKL,
			Role:       user.Role,
			IsVerified: user.IsVerified,
		})
	}

	response := &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "Users retrieved successfully",
		Data:       userData,
	}

	return response, nil
}
