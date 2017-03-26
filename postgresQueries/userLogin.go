package postgresQueries

type UserLogin struct {
	Token    string
	Username string
	Password string
}

func NewUserLogin(params map[string]string) *UserLogin {
	return &UserLogin{
		Token:    params["token"],
		Username: params["username"],
		Password: params["password"],
	}
}

func (q *UserLogin) GetQueryText() string {
	return "SELECT * FROM system.user_login($1, $2, $3)"
}

func (q *UserLogin) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Username,
		q.Password,
	}
}

func (q *UserLogin) NoResult() bool {
	return false
}

func (q *UserLogin) GetRowsRecord() (interface{}, []interface{}) {
	return NewUserLoginResultRecord()
}
