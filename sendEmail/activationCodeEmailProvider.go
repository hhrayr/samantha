package sendEmail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/hhrayr/samantha/configs"
)

type ActivationCodeEmailProvider struct {
	FullName      string
	ActivationUrl string
	language      string
}

func NewActivationCodeEmailProvider(fullName, activationKey, language string) *ActivationCodeEmailProvider {
	activationUrl := configs.GetPenelopeConfigs().ResolveRelativePath(fmt.Sprintf("email-client/activate?code=%s", activationKey))
	return &ActivationCodeEmailProvider{
		FullName:      fullName,
		ActivationUrl: activationUrl,
		language:      language,
	}
}

func (e *ActivationCodeEmailProvider) GetEmailBody() (string, error) {
	t, err := template.New("activation.tmpl").ParseFiles("./i18n/en/emailTemplates/activation.tmpl")
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

func (e *ActivationCodeEmailProvider) GetEmailSubject() string {
	return "Registration in autorialto"
}
