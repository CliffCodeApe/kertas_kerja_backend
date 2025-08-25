package service

import (
	"fmt"
	"kertas_kerja/contract"
	"kertas_kerja/dto"
	"kertas_kerja/entity"
	"kertas_kerja/pkg/bcrypt"
	"kertas_kerja/pkg/errs"
	token2 "kertas_kerja/pkg/token"
	"net/http"
)

type authServ struct {
	userRepo contract.UserRepository
}

func implAuthService(repo *contract.Repository) contract.AuthService {
	return &authServ{
		userRepo: repo.User,
	}
}

func (a *authServ) Register(payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	// Check if user already exists
	existingUser, err := a.userRepo.GetByNamaSatker(payload.NamaSatker)
	if err == nil && existingUser != nil {
		return nil, errs.ErrUserAlreadyExists
	}

	// Hash the password
	hashedPassword, err := bcrypt.Generate(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("bcrypt.Hash fail: %w", err)
	}

	// Create user DTO
	newUser := entity.User{
		Email:      payload.Email,
		NamaSatker: payload.NamaSatker,
		KodeKL:     payload.KodeKL,
		Password:   hashedPassword,
	}

	// Save user
	err = a.userRepo.InsertUser(&newUser)
	if err != nil {
		return nil, fmt.Errorf("userRepo.Create fail: %w", err)
	}

	response := &dto.AuthRegisterResponse{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
		Data: dto.RegisterResponse{
			Email:      newUser.Email,
			NamaSatker: newUser.NamaSatker,
			KodeKL:     newUser.KodeKL,
		},
	}

	return response, nil
}

func (a *authServ) Login(payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var response *dto.AuthLoginResponse
	// var loginErr error

	user, err := a.userRepo.GetByNamaSatker(payload.NamaSatker)
	if err != nil {
		return nil, errs.ErrLoginFailed
	}

	if user.IsVerified == "menunggu" {
		return nil, errs.ErrUserNotVerified
	} else if user.IsVerified == "ditolak" {
		return nil, errs.ErrUserNotVerified
	}

	verify := bcrypt.Verify(user.Password, payload.Password)

	if !verify {
		return nil, errs.ErrDecryptFailed
	}

	accessToken, err := token2.GenerateAccessToken(&token2.UserAuthToken{
		ID:          user.ID,
		Email:       user.Email,
		Nama_Satker: user.NamaSatker,
		KodeKL:      user.KodeKL,
		Role:        user.Role,
	})

	if err != nil {
		return nil, fmt.Errorf("token2.GenerateAccessToken fail: %w", err)
	}

	refreshToken, err := token2.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("token2.GenerateRefreshToken fail: %w", err)
	}

	response = &dto.AuthLoginResponse{
		StatusCode: http.StatusOK,
		Message:    "successfully logged in",
		Data: dto.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}

func (a *authServ) RefreshToken(token string) (*dto.RefreshTokenResponse, error) {
	id, err := token2.ValidateRefreshToken(token)
	if err != nil {
		return nil, fmt.Errorf("token2.ValidateRefreshToken fail: %w", err)
	}

	user, err := a.userRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("userRepo.GetById fail: %w", err)
	}

	accessToken, err := token2.GenerateAccessToken(&token2.UserAuthToken{
		ID:          user.ID,
		Email:       user.Email,
		Nama_Satker: user.NamaSatker,
		KodeKL:      user.KodeKL,
		Role:        user.Role,
	})

	if err != nil {
		return nil, fmt.Errorf("token2.GenerateAccessToken fail: %w", err)
	}

	response := &dto.RefreshTokenResponse{
		StatusCode:  http.StatusOK,
		Message:     "Refresh token berhasil",
		AccessToken: accessToken,
	}

	return response, nil
}
