package postgresQueries

type CompanyUsersResetLimitedStatus struct {
	Token       string
	CompanyName string
}

func NewCompanyUsersResetLimitedStatus(params map[string]string) *CompanyUsersResetLimitedStatus {
	return &CompanyUsersResetLimitedStatus{
		Token:       params["token"],
		CompanyName: params["companyname"],
	}
}

func (q *CompanyUsersResetLimitedStatus) GetQueryText() string {
	return "SELECT * FROM system.company_users_reset_limited_status($1, $2)"
}

func (q *CompanyUsersResetLimitedStatus) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.CompanyName,
	}
}

func (q *CompanyUsersResetLimitedStatus) NoResult() bool {
	return true
}

func (q *CompanyUsersResetLimitedStatus) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
