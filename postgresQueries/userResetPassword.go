package postgresQueries

type UserResetPassword struct {
	Token    string
	Username string
}

func NewUserResetPassword(params map[string]string) *UserResetPassword {
	return &UserResetPassword{
		Token:    params["token"],
		Username: params["username"],
	}
}

func (q *UserResetPassword) GetQueryText() string {
	return "SELECT * FROM system.user_reset_password($1, $2)"
}

func (q *UserResetPassword) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Username,
	}
}

func (q *UserResetPassword) NoResult() bool {
	return false
}

func (q *UserResetPassword) GetRowsRecord() (interface{}, []interface{}) {
	return NewUserResetPasswordResultRecord()
}
