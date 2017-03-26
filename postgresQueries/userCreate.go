package postgresQueries

type UserCreate struct {
	Token    string
	Username string
	Email    string
	Password string
	FullName string
}

func NewUserCreate(params map[string]string) *UserCreate {
	res := &UserCreate{
		Token:    params["token"],
		Username: params["username"],
		Email:    params["email"],
		Password: params["password"],
		FullName: params["fullname"],
	}

	return res
}

func (q *UserCreate) GetQueryText() string {
	return "SELECT * FROM system.user_create($1, $2, $3, $4, $5)"
}

func (q *UserCreate) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Username,
		q.Email,
		q.Password,
		q.FullName,
	}
}

func (q *UserCreate) NoResult() bool {
	return false
}

func (q *UserCreate) GetRowsRecord() (interface{}, []interface{}) {
	return NewUserCreateResultRecord()
}
