package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hhrayr/samantha/api"
	"github.com/hhrayr/samantha/configs"
	"github.com/hhrayr/samantha/utils"
	"github.com/lib/pq"
)

type httpError struct {
	Details string `json:"details,omitempty"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewHttpError(err error, requestUrl string) httpError {
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
			if configs.GetEnvConfigs().Env != "local" {
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

	res := httpError{
		Message: message,
		Status:  status,
	}

	if configs.GetEnvConfigs().Env == "local" {
		res.Details = fmt.Sprintf("%s | %s", requestUrl, message)
	}

	return res
}

func (he httpError) WriteToResponse(w http.ResponseWriter) {
	errorJSON, err := json.Marshal(he)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(he.Status)
		w.Write(errorJSON)
	}
	utils.LogError(he.Details, "http")
}
