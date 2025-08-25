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
	user, err := a.userRepo.GetById(userID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}

	err = a.userRepo.ValidateUser(userID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.ValidateUser fail: %w", err)
	}

	a.mailService.Enqueue(entity.Mail{
		To:      user.Email,
		Subject: "Validasi User",
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

func (a *userServ) InValidateUser(userID uint64) (*dto.ValidateUserResponse, error) {
	user, err := a.userRepo.GetById(userID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}

	err = a.userRepo.InValidateUser(userID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.InValidateUser fail: %w", err)
	}

	a.mailService.Enqueue(entity.Mail{
		To:      user.Email,
		Subject: "Invalidasi User",
		Body:    fmt.Sprintf("User anda ditolak untuk divalidasi oleh Super Admin. Silahkan hubungi Super Admin untuk informasi lebih lanjut. Terima kasih"),
	})

	response := &dto.ValidateUserResponse{
		StatusCode: http.StatusOK,
		Message:    "User successfully Invalidated",
	}

	return response, nil
}

func (a *userServ) DeleteUser(userID uint64) (*dto.DeleteUserResponse, error) {
	user, err := a.userRepo.GetById(userID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}

	if user.Role == "superadmin" {
		return nil, fmt.Errorf("cannot delete superadmin user")
	}

	err = a.userRepo.DeleteUser(userID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.DeleteUser fail: %w", err)
	}

	response := &dto.DeleteUserResponse{
		StatusCode: http.StatusOK,
		Message:    "User successfully deleted",
	}

	return response, nil
}

func (a *userServ) ChangeUserRole(payload *dto.ChangeUserRoleRequest, userID uint64) (*dto.ChangeUserRoleResponse, error) {
	user, err := a.userRepo.GetById(userID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}

	if payload.Role == "superadmin" {
		return &dto.ChangeUserRoleResponse{
			StatusCode: http.StatusForbidden,
			Message:    "Tidak bisa mengubah role user ke superadmin",
		}, err
	}

	if user.Role == payload.Role {
		return &dto.ChangeUserRoleResponse{
			StatusCode: http.StatusForbidden,
			Message:    "User sudah memiliki role yang sama",
		}, err
	}

	if payload.Role != "satker" && payload.Role != "admin" {
		return &dto.ChangeUserRoleResponse{
			StatusCode: http.StatusForbidden,
			Message:    "Role tidak valid, hanya bisa diubah ke 'satker' atau 'admin'",
		}, err
	}

	err = a.userRepo.ChangeUserRole(userID, payload.Role)
	if err != nil {
		return nil, fmt.Errorf("userRepo.ChangeUserRole fail: %w", err)
	}

	response := &dto.ChangeUserRoleResponse{
		StatusCode: http.StatusOK,
		Message:    "User role successfully changed",
	}

	return response, nil
}
