package service

import (
	"kertas_kerja/contract"
	"kertas_kerja/pkg/mail"
	"log"
	"os"
	"strconv"
)

func New(repo *contract.Repository) *contract.Service {
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	mailService := mail.ImplMailService(
		os.Getenv("SMTP_FROM"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("BASE_URL"),
		smtpPort,
	)

	return &contract.Service{
		KertasKerja: implKertasKerjaService(repo),
		Auth:        implAuthService(repo),
		User:        implUserService(repo),
		Mail:        mailService,
	}
}
