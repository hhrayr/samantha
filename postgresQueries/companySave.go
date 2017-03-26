package postgresQueries

type CompanySave struct {
	Token        string
	Name         string
	Country      string
	City         string
	AddressLine1 string
	AddressLine2 string
	PostalCode   string
	PhoneNumber  string
}

func NewCompanySave(params map[string]string) *CompanySave {
	res := &CompanySave{
		Token:        params["token"],
		Name:         params["name"],
		Country:      params["country"],
		City:         params["city"],
		AddressLine1: params["addressline1"],
		AddressLine2: params["addressline2"],
		PostalCode:   params["postalcode"],
		PhoneNumber:  params["phonenumber"],
	}

	return res
}

func (q *CompanySave) GetQueryText() string {
	return "SELECT * FROM system.company_save($1, $2, $3, $4, $5, $6, $7, $8)"
}

func (q *CompanySave) GetParams() []interface{} {
	return []interface{}{
		q.Token,
		q.Name,
		q.Country,
		q.City,
		q.AddressLine1,
		q.AddressLine2,
		q.PostalCode,
		q.PhoneNumber,
	}
}

func (q *CompanySave) NoResult() bool {
	return true
}

func (q *CompanySave) GetRowsRecord() (interface{}, []interface{}) {
	return nil, nil
}
