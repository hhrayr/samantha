package postgresQueries

type UserUpdate struct {
	Token    string
	Password string
	FullName string
}

func NewUserUpdate(params map[string]string) *UserUpdate {
	return &UserUpdate{
		Token:    params["token"],
		Password: params["password"],
		FullName: params["fullname"],
	}
}

func (q *UserUpdate) GetQueryText() string {
	return "SELECT * FROM system.user_update($1, $2, $3)"
}

func (q *UserUpdate) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Password,
		q.FullName,
	}
}

func (q *UserUpdate) NoResult() bool {
	return true
}

func (q *UserUpdate) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
