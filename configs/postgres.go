package configs

import "sync"

type PostgresConfigs struct {
	SmanthaSessionToken string
	ConnectionString    string
}

func newPostgresConfigs() *PostgresConfigs {
	return &PostgresConfigs{
		SmanthaSessionToken: "b320a271-ef64-4100-928e-405187aa5443",
		ConnectionString:    "postgres://postgres:__123456@localhost/samantha?sslmode=disable",
	}
}

var postgresConfigInstance *PostgresConfigs
var postgresConfigInstanceOnce sync.Once

func GetPgConfigs() *PostgresConfigs {
	postgresConfigInstanceOnce.Do(func() {
		postgresConfigInstance = newPostgresConfigs()
	})
	return postgresConfigInstance
}
