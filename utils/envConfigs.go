package utils

import "os"

func GetEnvVariable(name, defaultValue string) string {
	res := os.Getenv(name)
	if res == "" {
		res = defaultValue
	}
	return res
}
