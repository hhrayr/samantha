package configs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/hhrayr/samantha/utils"
)

type EnvConfigs struct {
	Env       string
	CachePath string
}

func NewEnvConfigs() *EnvConfigs {
	return &EnvConfigs{
		Env:       utils.GetEnvVariable("SAMANTHA_ENV", "local"),
		CachePath: "cache",
	}
}

func (env *EnvConfigs) GetSamanthaHostUrl() string {
	switch env.Env {
	case "local":
		return "http://localhost:4040/"
	}
	return "https://google.com"
}

func (env *EnvConfigs) ResolveHttpCachePath(filename string) string {
	return fmt.Sprintf("%s/%s", env.CachePath, filename)
}

func (env *EnvConfigs) ResolveSamanthaAbsoluteUrl(url string) string {
	pageUrl := url
	if strings.HasPrefix(url, "/") {
		pageUrl = strings.TrimPrefix(url, "/")
	}
	return fmt.Sprintf("%s%s", env.GetSamanthaHostUrl(), pageUrl)
}

var envConfigInstance *EnvConfigs
var envConfigInstanceOnce sync.Once

func GetEnvConfigs() *EnvConfigs {
	envConfigInstanceOnce.Do(func() {
		envConfigInstance = NewEnvConfigs()
	})
	return envConfigInstance
}
