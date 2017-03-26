package postgresQueries

type CompanyUsersSetLimitedStatus struct {
	Token       string
	CompanyName string
}

func NewCompanyUsersSetLimitedStatus(params map[string]string) *CompanyUsersSetLimitedStatus {
	return &CompanyUsersSetLimitedStatus{
		Token:       params["token"],
		CompanyName: params["companyname"],
	}
}

func (q *CompanyUsersSetLimitedStatus) GetQueryText() string {
	return "SELECT * FROM system.company_users_set_limited_status($1, $2)"
}

func (q *CompanyUsersSetLimitedStatus) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.CompanyName,
	}
}

func (q *CompanyUsersSetLimitedStatus) NoResult() bool {
	return true
}

func (q *CompanyUsersSetLimitedStatus) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
