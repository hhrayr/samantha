package api

import (
	"fmt"
	"strings"
)

type ApiProvider struct {
	apiDomain string
	apiMethod string
	params    map[string]string
}

func NewApiProvider(params map[string]string, requestMethod string) (*ApiProvider, error) {
	apiDomain := strings.ToLower(params["apidomain"])
	apiMethod := strings.ToLower(fmt.Sprintf("%s_%s", requestMethod, params["apimethod"]))
	token := params["token"]
	if apiDomain != "" && apiMethod != "" && token != "" {
		return &ApiProvider{
			apiDomain: apiDomain,
			apiMethod: apiMethod,
			params:    params,
		}, nil
	}
	return nil, newError("api.error.no_enough_parameters")
}

func (ap *ApiProvider) InvokeMethod() (interface{}, error) {
	switch ap.apiDomain {
	case "system":
		return newSystem(ap.params, ap.apiMethod).invoke()
	}
	return nil, newError("api domain not found")
}
