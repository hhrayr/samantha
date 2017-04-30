package httpCache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/hhrayr/samantha/configs"
)

type HttpCachedData struct {
	Header    http.Header
	Timestamp time.Time
	Data      []byte
}

func NewHttpCachedData(header http.Header) *HttpCachedData {
	return &HttpCachedData{
		Header: header,
	}
}

func NewHttpCachedDataFromResponseRecorder(rec *httptest.ResponseRecorder) *HttpCachedData {
	return &HttpCachedData{
		Timestamp: time.Now(),
		Header:    rec.HeaderMap,
		Data:      rec.Body.Bytes(),
	}
}

func (cd *HttpCachedData) SetData(data []byte) {
	cd.Timestamp = time.Now()
	cd.Data = data
}

type HttpCache struct {
	hash string
}

var cachSetupFuncOnce sync.Once

func NewHttpCache(key string) *HttpCache {
	return &HttpCache{
		hash: url.QueryEscape(key),
	}
}

func (ch *HttpCache) getFileName() string {
	return configs.GetEnvConfigs().ResolveHttpCachePath(fmt.Sprintf("%s.cache", ch.hash))
}

func (ch *HttpCache) SetValue(value *HttpCachedData) error {
	file, err := os.OpenFile(ch.getFileName(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	var data []byte
	data, err = json.Marshal(value)
	if err != nil {
		return err
	}

	n, err := file.Write(data)

	if err != nil {
		return err
	} else if n == 0 {
		return errors.New(fmt.Sprintf("no content cashed for %s", ch.hash))
	}

	return nil
}

func (ch *HttpCache) GetValue() (*HttpCachedData, error) {
	data, err := ioutil.ReadFile(ch.getFileName())
	if err != nil {
		return nil, err
	}

	var res *HttpCachedData
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
