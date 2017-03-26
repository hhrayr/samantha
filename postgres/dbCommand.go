package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/hhrayr/samantha/configs"
)

type DbCommand struct {
	query DbQuery
}

func NewDbCommand(query DbQuery) *DbCommand {
	return &DbCommand{
		query: query,
	}
}

func (dbCommand *DbCommand) Run() ([]interface{}, error) {
	dbConnection, err := getDBConnection()
	if err != nil {
		return nil, err
	}

	rows, err := dbConnection.Query(dbCommand.query.GetQueryText(), dbCommand.getDBParams()...)
	if err != nil {
		return nil, err
	}

	if dbCommand.query.NoResult() {
		rows.Close()
		return nil, nil
	}

	return dbCommand.fetchResults(rows)
}

func (dbCommand *DbCommand) getDBParams() []interface{} {
	var res []interface{}
	if queryParams := dbCommand.query.GetParams(); queryParams != nil && len(queryParams) > 0 {
		res = make([]interface{}, len(queryParams))
		for index, param := range queryParams {
			switch paramVal := param.(type) {
			case string:
				if paramVal != "" {
					res[index] = paramVal
				}
				break
			default:
				res[index] = paramVal
				break
			}
		}
	}
	return res
}

func (dbCommand *DbCommand) fetchResults(rows *sql.Rows) ([]interface{}, error) {
	defer rows.Close()

	var result []interface{}
	for rows.Next() {
		record, recordFields := dbCommand.query.GetRowsRecord()
		err := rows.Scan(recordFields...)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getDBConnection() (*sql.DB, error) {
	return sql.Open("postgres", configs.GetPgConfigs().ConnectionString)
}
