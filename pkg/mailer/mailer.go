package mailer

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type UserInfo struct {
	Name         string
	Email        string
	EmailPasword string
	SmtpServer   string
	SmtpPort     string
	Hash         string
}

func Send(user UserInfo) error {
	// Отправляет ссылку на подтверждение регистрации http://localhost:8081/verify/{hash}

	userName := user.Name                 //"Yury"
	userEmail := user.Email               //"yury.s.nazarov@mail.ru"
	userEmailPasword := user.EmailPasword // "31pHOmYwf1HhV7o7kcI7"
	userSmtpServer := user.SmtpServer     // "smtp.mail.ru"
	userSmtpPort := user.SmtpPort         // "587"
	userConfirmationLink := fmt.Sprintf("<a href='http://localhost:8081/verify/%s'>registration confirm</a>", user.Hash)

	// Конфигурация Email клиента
	e := &email.Email{
		To:      []string{userEmail},
		From:    fmt.Sprintf("%s <%s>", userName, userEmail),
		Subject: fmt.Sprintf("Hello, %s!", userName),
		HTML:    []byte(userConfirmationLink),
	}
	err := e.Send(
		fmt.Sprintf("%s:%s", userSmtpServer, userSmtpPort),
		smtp.PlainAuth("", userEmail, userEmailPasword, userSmtpServer),
	)
	if err != nil {
		fmt.Printf("ERRO: Send email problem: %x", err.Error())
		return err
	}
	fmt.Println("INFO: Email was send")
	return nil
}
