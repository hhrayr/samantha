package postgresQueries

type UserUnbind struct {
	Token    string
	Username string
}

func NewUserUnbind(params map[string]string) *UserUnbind {
	return &UserUnbind{
		Token:    params["token"],
		Username: params["username"],
	}
}

func (q *UserUnbind) GetQueryText() string {
	return "SELECT * FROM system.user_unbind($1, $2)"
}

func (q *UserUnbind) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Username,
	}
}

func (q *UserUnbind) NoResult() bool {
	return true
}

func (q *UserUnbind) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
