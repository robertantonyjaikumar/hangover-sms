package utils

import (
	"fmt"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"net/smtp"
	"sms/config"
)

var emailConfig = config.NewEmail()

func SendEmail(to string, subject string, htmlBody string) error {
	logger.Info("Sending email to " + emailConfig.Sender)
	auth := smtp.PlainAuth("", emailConfig.Sender, emailConfig.Password, emailConfig.Host)

	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n",
		emailConfig.Sender, to, subject)

	message := []byte(headers + htmlBody)

	err := smtp.SendMail(emailConfig.Host+":"+emailConfig.Port, auth, emailConfig.Sender, []string{to}, message)
	return err
}
