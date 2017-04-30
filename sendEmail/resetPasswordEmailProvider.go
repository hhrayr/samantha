package sendEmail

import (
	"bytes"
	"html/template"
)

type ResetPasswordEmailProvider struct {
	FullName    string
	NewPassword string
	language    string
}

func NewResetPasswordEmailProvider(fullName, newPassword, language string) *ResetPasswordEmailProvider {
	return &ResetPasswordEmailProvider{
		FullName:    fullName,
		NewPassword: newPassword,
		language:    language,
	}
}

func (e *ResetPasswordEmailProvider) GetEmailBody() (string, error) {
	t, err := template.New("reset_password.tmpl").ParseFiles("./i18n/en/emailTemplates/reset_password.tmpl")
	if err != nil {
		return "", err
	}

	var doc bytes.Buffer
	err = t.Execute(&doc, e)
	if err != nil {
		return "", err
	}

	return doc.String(), nil
}

func (e *ResetPasswordEmailProvider) GetEmailSubject() string {
	return "Reset pasword in autorialto"
}
