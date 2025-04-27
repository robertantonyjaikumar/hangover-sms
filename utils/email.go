package utils

import (
	"fmt"
	"net/smtp"
)

var (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	smtpEmail    = "your-email@gmail.com"
	smtpPassword = "your-app-password"
)

func SendEmail(to string, subject string, htmlBody string) error {
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n",
		smtpEmail, to, subject)

	message := []byte(headers + htmlBody)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, []string{to}, message)
	return err
}
