package sendEmail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"

	"github.com/hhrayr/samantha/configs"
)

type InvitationCodeEmailProvider struct {
	InvitationUrl      string
	UserFullName       string
	InviterFullName    string
	InviterCompanyName string
	language           string
}

func NewInvitationCodeEmailProvider(
	userFullName,
	inviterFullName,
	inviterCompanyName,
	invitationKey,
	language string,
) *InvitationCodeEmailProvider {
	invitationUrl := configs.GetPenelopeConfigs().ResolveRelativePath(
		fmt.Sprintf("email-client/accept-invitation?code=%s&user_name=%s&company_name=%s",
			invitationKey,
			url.QueryEscape(userFullName),
			url.QueryEscape(inviterCompanyName)))
	return &InvitationCodeEmailProvider{
		InvitationUrl:      invitationUrl,
		UserFullName:       userFullName,
		InviterFullName:    inviterFullName,
		InviterCompanyName: inviterCompanyName,
		language:           language,
	}
}

func (e *InvitationCodeEmailProvider) GetEmailBody() (string, error) {
	t, err := template.New("user_invite.tmpl").ParseFiles("./i18n/en/emailTemplates/user_invite.tmpl")
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

func (e *InvitationCodeEmailProvider) GetEmailSubject() string {
	return "Invitation from autorialto"
}
