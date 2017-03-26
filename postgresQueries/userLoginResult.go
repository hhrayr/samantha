package postgresQueries

type UserLoginResult struct {
	Token       string      `json:"token"`
	FullName    string      `json:"fullName"`
	Roles       interface{} `json:"roles"`
	Status      string      `json:"-"`
	CompanyName interface{} `json:"companyName"`
}

func NewUserLoginResultRecord() (interface{}, []interface{}) {
	record := &UserLoginResult{}
	recordFields := []interface{}{
		&record.Token,
		&record.FullName,
		&record.Roles,
		&record.Status,
		&record.CompanyName,
	}
	return record, recordFields
}
