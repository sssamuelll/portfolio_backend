package services

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

const emailCodeLength = 6

func GenerateEmailCode() string {
	rand.NewSource(time.Now().UnixNano())
	code := ""
	for i := 0; i < emailCodeLength; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func SendEmail(to, subject, body string) error {
	from := "your_email@gmail.com" // Replace with your email
	password := "your_email_password"

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, password, smtpServer)

	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, []byte(msg))
	return err
}
