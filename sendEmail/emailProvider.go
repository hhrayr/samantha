package sendEmail

type EmailProvider interface {
	GetEmailBody() (string, error)
	GetEmailSubject() string
}

func GetActivationCodeEmailProvider(fullName, activationKey, language string) EmailProvider {
	return NewActivationCodeEmailProvider(fullName, activationKey, language)
}

func GetResetPasswordEmailProvider(fullName, newPassword, language string) EmailProvider {
	return NewResetPasswordEmailProvider(fullName, newPassword, language)
}

func GetInvitationCodeEmailProvider(
	userFullName,
	inviterFullName,
	inviterCompanyName,
	invitationKey,
	language string) EmailProvider {
	return NewInvitationCodeEmailProvider(
		userFullName,
		inviterFullName,
		inviterCompanyName,
		invitationKey,
		language)
}
