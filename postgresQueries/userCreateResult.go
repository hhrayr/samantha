package postgresQueries

type UserCreateResult struct {
	ActivationKey string
}

func NewUserCreateResultRecord() (interface{}, []interface{}) {
	record := &UserCreateResult{}
	recordFields := []interface{}{
		&record.ActivationKey,
	}
	return record, recordFields
}
