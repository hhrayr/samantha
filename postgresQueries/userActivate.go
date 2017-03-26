package postgresQueries

type UserActivate struct {
	Token         string
	ActivationKey string
}

func NewUserActivate(params map[string]string) *UserActivate {
	return &UserActivate{
		Token:         params["token"],
		ActivationKey: params["activationkey"],
	}
}

func (q *UserActivate) GetQueryText() string {
	return "SELECT * FROM system.user_activate($1, $2)"
}

func (q *UserActivate) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.ActivationKey,
	}
}

func (q *UserActivate) NoResult() bool {
	return true
}

func (q *UserActivate) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
