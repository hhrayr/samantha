package configs

import "os"

func GetEnv() string {
	env := os.Getenv("SAMANTHA_ENV")
	if env == "" {
		env = "local"
	}
	return env
}
