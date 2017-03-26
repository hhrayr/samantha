package postgresQueries

type UserResetPasswordResult struct {
	NewPassword  string
	UserEmail    string
	UserFullName string
}

func NewUserResetPasswordResultRecord() (interface{}, []interface{}) {
	record := &UserResetPasswordResult{}
	recordFields := []interface{}{
		&record.NewPassword,
		&record.UserEmail,
		&record.UserFullName,
	}
	return record, recordFields
}
