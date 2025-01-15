package services

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/sssamuelll/portfolio_backend/config"
	dkim "github.com/toorop/go-dkim"
)

func SendEmail(to, subject, body string) error {
	from := config.AppConfig.EmailSender
	password := config.AppConfig.EmailPassword
	smtpServer := config.AppConfig.SMTPServer
	smtpPort := config.AppConfig.SMTPPort
	privateKey := config.AppConfig.DKIMPrivateKey
	domain := config.AppConfig.DKIMDomain
	selector := config.AppConfig.DKIMSelector

	// Email headers
	headers := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n", from, to, subject)
	message := headers + body

	// DKIM settings
	dkimOptions := dkim.NewSigOptions()
	dkimOptions.PrivateKey = []byte(privateKey)
	dkimOptions.Domain = domain
	dkimOptions.Selector = selector
	dkimOptions.Headers = []string{"from", "to", "subject"}
	dkimOptions.BodyLength = 0

	// Sign the email with DKIM
	var signedMessage bytes.Buffer
	messageBytes := []byte(message)
	err := dkim.Sign(&messageBytes, dkimOptions)
	if err != nil {
		return fmt.Errorf("failed to sign email with DKIM: %v", err)
	}
	signedMessage.Write(messageBytes)

	// SMTP authentication
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Send the email
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, signedMessage.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
