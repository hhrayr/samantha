package configs

import (
	"fmt"
	"strings"
	"sync"
)

type PenelopeConfigs struct {
	host string
}

func (c *PenelopeConfigs) ResolveRelativePath(path string) string {
	normalisedPath := strings.TrimPrefix(path, "/")
	return fmt.Sprintf("%s%s", c.host, normalisedPath)
}

func NewPenelopeConfigs(env string) *PenelopeConfigs {
	return &PenelopeConfigs{
		host: getPenelopeHost(env),
	}
}

func getPenelopeHost(env string) string {
	if env == "local" {
		return "http://localhost:8080/"
	}
	return "https://google.com/"
}

var penelopeConfigInstance *PenelopeConfigs
var penelopeConfigInstanceOnce sync.Once

func GetPenelopeConfigs() *PenelopeConfigs {
	penelopeConfigInstanceOnce.Do(func() {
		penelopeConfigInstance = NewPenelopeConfigs(GetEnvConfigs().Env)
	})
	return penelopeConfigInstance
}
