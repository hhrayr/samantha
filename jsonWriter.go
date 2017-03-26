package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hhrayr/samantha/api"
)

type JsonWriter struct {
	requestParams  map[string]string
	requestMethod  string
	responseWriter http.ResponseWriter
}

func newJsonWriter(requestParams map[string]string, w http.ResponseWriter) *JsonWriter {
	return &JsonWriter{
		requestParams:  requestParams,
		responseWriter: w,
	}
}

func (jw *JsonWriter) setRequestMethod(requestMethod string) {
	jw.requestMethod = requestMethod
}

func (jw *JsonWriter) writeApiMethodInvokeResult() {
	var err error
	var apiProvider *api.ApiProvider
	var apiResult interface{}
	if apiProvider, err = api.NewApiProvider(jw.requestParams, jw.requestMethod); err == nil {
		apiResult, err = apiProvider.InvokeMethod()
	}
	if err != nil {
		NewHttpError(err).WriteToResponse(jw.responseWriter)
	} else {
		jw.write(apiResult)
	}
}

func (jw *JsonWriter) write(data interface{}) {
	jsonData, err := encodeJSON(data)
	if err != nil {
		NewHttpError(err).WriteToResponse(jw.responseWriter)
	} else {
		jw.responseWriter.Header().Set("Content-Type", "application/json")
		jw.responseWriter.WriteHeader(http.StatusOK)
		jw.responseWriter.Write(jsonData)
	}
}

func encodeJSON(data interface{}) ([]byte, error) {
	jsonBuffer := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), nil
}
