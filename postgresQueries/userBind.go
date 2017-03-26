package postgresQueries

type UserBind struct {
	Token         string
	InvitationKey string
}

func NewUserBind(params map[string]string) *UserBind {
	return &UserBind{
		Token:         params["token"],
		InvitationKey: params["invitationkey"],
	}
}

func (q *UserBind) GetQueryText() string {
	return "SELECT * FROM system.user_bind($1, $2)"
}

func (q *UserBind) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.InvitationKey,
	}
}

func (q *UserBind) NoResult() bool {
	return true
}

func (q *UserBind) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
