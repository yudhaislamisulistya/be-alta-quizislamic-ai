package util

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"text/template"
	"time"
)

func SendMail(to string, subject, message string) error {
	body := "From: " + os.Getenv("SMTP_USER") + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		message

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))
	smtpAddr := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))

	err := smtp.SendMail(smtpAddr, auth, os.Getenv("SMTP_USER"), []string{to}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func GenerateVerificationCode() string {
	var verificationCode string
	for i := 0; i < 6; i++ {
		digit := strconv.Itoa(rand.Intn(10))
		verificationCode += digit
	}

	return verificationCode
}

func ReadEmailTemplateVericationCode() (*template.Template, error) {
	return template.ParseFiles("lib/templates/email_verification_code.html")
}

func ReadEmailTemplateUserActivation() (*template.Template, error) {
	return template.ParseFiles("lib/templates/email_user_activation.html")
}

func GetExpiredTime(minutes int) int {
	exp := time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	return int(exp)
}
