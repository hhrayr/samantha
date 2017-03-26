package postgresQueries

type UserCreateInvitationResult struct {
	InvitationKey      string `json:"invitationKey"`
	UserEmail          string `json:"userEmail"`
	UserFullName       string `json:"userFullName"`
	InviterFullName    string `json:"inviterFullName"`
	InviterCompanyName string `json:"inviterCompanyName"`
}

func NewUserCreateInvitationResultRecord() (interface{}, []interface{}) {
	record := &UserCreateInvitationResult{}
	recordFields := []interface{}{
		&record.InvitationKey,
		&record.UserEmail,
		&record.UserFullName,
		&record.InviterFullName,
		&record.InviterCompanyName,
	}
	return record, recordFields
}
