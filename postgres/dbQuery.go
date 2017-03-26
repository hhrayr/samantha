package postgres

type DbQuery interface {
	GetQueryText() string
	GetParams() []interface{}
	GetRowsRecord() (interface{}, []interface{})
	NoResult() bool
}
