package utils

import (
	"fmt"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"net/smtp"
)

var (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
	//smtpEmail    = os.Getenv("EMAIL_SENDER") // todo: have to fix -  not able to read the .env values
	//smtpPassword = os.Getenv("EMAIL_PASSWORD")
	smtpEmail    = "antonyrajjacobg@gmail.com"
	smtpPassword = "omui kyoc qvcl eweb"
)

func SendEmail(to string, subject string, htmlBody string) error {
	logger.Info("Sending email to " + smtpEmail)
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n",
		smtpEmail, to, subject)

	message := []byte(headers + htmlBody)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, []string{to}, message)
	return err
}
