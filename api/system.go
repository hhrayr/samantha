package api

import (
	"github.com/hhrayr/samantha/postgres"
	"github.com/hhrayr/samantha/postgresQueries"
	"github.com/hhrayr/samantha/sendEmail"
	"github.com/hhrayr/samantha/utils"
)

type System struct {
	params map[string]string
	method string
}

func newSystem(params map[string]string, method string) *System {
	return &System{
		params: params,
		method: method,
	}
}

func (s *System) invoke() (interface{}, error) {
	switch s.method {
	case "post_login":
		return s.login()
	case "post_logout":
		return nil, s.logout()
	case "post_userupdate":
		return nil, s.userUpdate()
	case "post_usercreate":
		return s.userCreate()
	case "post_useractivate":
		return nil, s.userActivate()
	case "post_userresetpassword":
		return s.userResetPassword()
	case "post_companyusersresetlimitedstatus":
		return nil, s.companyUsersResetLimitedStatus()
	case "post_companyuserssetlimitedstatus":
		return nil, s.companyUsersSetLimitedStatus()
	case "post_companysave":
		return nil, s.companySave()
	case "post_usercreateinvitation":
		return nil, s.userCreateInvitation()
	case "post_userbind":
		return nil, s.userBind()
	case "post_userunbind":
		return nil, s.userUnbind()
	}

	return nil, newError("api.error.method_not_found")
}

func (s *System) login() (*postgresQueries.UserLoginResult, error) {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserLogin(s.params))
	result, err := dbCommand.Run()
	if err != nil {
		return nil, err
	}

	return result[0].(*postgresQueries.UserLoginResult), nil
}

func (s *System) logout() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserLogot(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) userUpdate() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserUpdate(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) userCreate() (*postgresQueries.UserCreateResult, error) {
	user := postgresQueries.NewUserCreate(s.params)
	dbCommand := postgres.NewDbCommand(user)
	result, err := dbCommand.Run()
	if err != nil {
		return nil, err
	}

	createdUser, ok := result[0].(*postgresQueries.UserCreateResult)
	if !ok {
		return nil, newError("Created user data not provided")
	}

	emailProvider := sendEmail.GetActivationCodeEmailProvider(
		user.FullName, createdUser.ActivationKey, s.params["language"])
	emailBody, err := emailProvider.GetEmailBody()
	if err != nil {
		utils.LogError("email", err.Error())
	} else {
		sendEmail.SedEmail("", []string{user.Email}, emailProvider.GetEmailSubject(), emailBody)
	}

	return nil, nil
}

func (s *System) userActivate() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserActivate(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) userResetPassword() (*postgresQueries.UserResetPasswordResult, error) {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserResetPassword(s.params))
	result, err := dbCommand.Run()
	if err != nil {
		return nil, err
	}

	resetPassword, ok := result[0].(*postgresQueries.UserResetPasswordResult)
	if !ok {
		return nil, newError("Reset password user data not provided")
	}

	emailProvider := sendEmail.GetResetPasswordEmailProvider(
		resetPassword.UserFullName, resetPassword.NewPassword, s.params["language"])
	emailBody, err := emailProvider.GetEmailBody()
	if err != nil {
		utils.LogError("email", err.Error())
	} else {
		sendEmail.SedEmail("", []string{resetPassword.UserEmail}, emailProvider.GetEmailSubject(), emailBody)
	}

	return nil, nil
}

func (s *System) companyUsersResetLimitedStatus() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewCompanyUsersResetLimitedStatus(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) companyUsersSetLimitedStatus() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewCompanyUsersSetLimitedStatus(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) companySave() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewCompanySave(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) userCreateInvitation() error {
	userInvitation := postgresQueries.NewUserCreateInvitation(s.params)
	dbCommand := postgres.NewDbCommand(userInvitation)
	result, err := dbCommand.Run()
	if err != nil {
		return err
	}

	invitedUser, ok := result[0].(*postgresQueries.UserCreateInvitationResult)
	if !ok {
		return newError("Invoted user data not provided")
	}

	emailProvider := sendEmail.GetInvitationCodeEmailProvider(
		invitedUser.UserFullName,
		invitedUser.InviterFullName,
		invitedUser.InviterCompanyName,
		invitedUser.InvitationKey,
		s.params["language"])
	emailBody, err := emailProvider.GetEmailBody()
	if err != nil {
		utils.LogError("email", err.Error())
	} else {
		sendEmail.SedEmail("", []string{invitedUser.UserEmail}, emailProvider.GetEmailSubject(), emailBody)
	}

	return nil
}

func (s *System) userBind() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserBind(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func (s *System) userUnbind() error {
	dbCommand := postgres.NewDbCommand(postgresQueries.NewUserUnbind(s.params))
	_, err := dbCommand.Run()
	if err != nil {
		return err
	}

	return nil
}
