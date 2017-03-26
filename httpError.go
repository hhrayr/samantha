package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"fmt"

	"github.com/hhrayr/samantha/api"
	"github.com/hhrayr/samantha/configs"
	"github.com/hhrayr/samantha/utils"
	"github.com/lib/pq"
)

type httpError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewHttpError(err error) httpError {
	status := http.StatusInternalServerError
	message := err.Error()

	switch err := err.(type) {
	case *pq.Error:
		message = err.Message
		if message == "db.error.user_token_expired" ||
			message == "db.error.user_token_invalid" {
			status = http.StatusUnauthorized
		} else if strings.HasPrefix(message, "db.error") {
			status = http.StatusConflict
		} else {
			if configs.GetEnv() != "local" {
				message = "db.error.generic_error"
			}
			status = http.StatusInternalServerError
		}
		break
	case *api.Error:
		message = err.Message
		status = http.StatusConflict
		break
	}

	return httpError{
		Message: message,
		Status:  status,
	}
}

func (he httpError) SetRequestParameters(params map[string]string) {
	requestParams := bytes.NewBufferString("")
	for name, value := range params {
		if requestParams.Len() > 0 {
			requestParams.WriteString(", ")
		}
		requestParams.WriteString(fmt.Sprintf("%s=%s", name, value))
	}
	he.Message = fmt.Sprintf("%s - request params [%s]", he.Message, requestParams.String())
}

func (he httpError) SetRequestUrl(url string) {
	he.Message = fmt.Sprintf("%s - request url [%s]", url)
}

func (he httpError) WriteToResponse(w http.ResponseWriter) {
	errorJSON, err := json.Marshal(he)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(he.Status)
		w.Write(errorJSON)
	}
	utils.LogError("http", he.Message)
}
