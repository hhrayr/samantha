package postgresQueries

type UserLogout struct {
	Token string
}

func NewUserLogot(params map[string]string) *UserLogout {
	return &UserLogout{
		Token: params["token"],
	}
}

func (q *UserLogout) GetQueryText() string {
	return "SELECT * FROM system.user_logout($1)"
}

func (q *UserLogout) GetParams() []interface{} {
	return []interface{}{
		q.Token,
	}
}

func (q *UserLogout) NoResult() bool {
	return true
}

func (q *UserLogout) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
