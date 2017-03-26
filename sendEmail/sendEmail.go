package sendEmail

import (
	"fmt"
	"net/smtp"

	"github.com/hhrayr/samantha/configs"
)

func SedEmail(from string, to []string, subject string, body string) error {

	smtpConfigs := configs.GetSmtpConfigs()
	auth := smtp.PlainAuth("",
		smtpConfigs.Username,
		smtpConfigs.Password,
		smtpConfigs.EmailServer)

	msgMime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msgSubject := "Subject: " + subject + "!\n"
	msg := []byte(msgSubject + msgMime + "\n" + body)
	addr := fmt.Sprintf("%s:%v", smtpConfigs.EmailServer, smtpConfigs.Port)

	if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
		return err
	}
	return nil
}
