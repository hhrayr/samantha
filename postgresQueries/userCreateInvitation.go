package postgresQueries

type UserCreateInvitation struct {
	Token    string
	Username string
	Roles    string
}

func NewUserCreateInvitation(params map[string]string) *UserCreateInvitation {
	return &UserCreateInvitation{
		Token:    params["token"],
		Username: params["username"],
		Roles:    params["roles"],
	}
}

func (q *UserCreateInvitation) GetQueryText() string {
	return "SELECT * FROM system.user_create_invitation($1, $2, $3)"
}

func (q *UserCreateInvitation) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Username,
		q.Roles,
	}
}

func (q *UserCreateInvitation) NoResult() bool {
	return false
}

func (q *UserCreateInvitation) GetRowsRecord() (interface{}, []interface{}) {
	return NewUserCreateInvitationResultRecord()
}
