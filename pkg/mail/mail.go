package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"

	"kertas_kerja/contract"

	"kertas_kerja/entity"
)

type MailService struct {
	From     string
	Password string
	Host     string
	Port     int
	AppUrl   string
	Queue    chan entity.Mail
	Wg       sync.WaitGroup
}

func ImplMailService(from, password, smtpHost, appUrl string, smtpPort int) contract.MailService {
	mailService := &MailService{
		From:     from,
		Password: password,
		Host:     smtpHost,
		Port:     smtpPort,
		AppUrl:   appUrl,
		Queue:    make(chan entity.Mail, 100),
	}

	// Start processing emails in the background
	go mailService.Dequeue()
	return mailService
}

func (m *MailService) Enqueue(email entity.Mail) {
	log.Println("Enqueuing email to:", email.To)
	m.Queue <- email
}

func (m *MailService) Dequeue() {
	for email := range m.Queue {
		m.Wg.Add(1)
		if err := m.SendEmail(email.To, email.Subject, email.Body); err != nil {
			log.Println("Error sending email:", err)
		}
		m.Wg.Done()
	}
}

func (m *MailService) SendEmail(to, subject, body string) error {

	from := m.From
	password := m.Password
	smtpHost := m.Host
	smtpPort := fmt.Sprintf("%d", m.Port)

	// Message format
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
