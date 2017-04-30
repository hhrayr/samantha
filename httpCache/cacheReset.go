package httpCache

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/hhrayr/samantha/configs"
	"github.com/hhrayr/samantha/utils"
)

type CacheReset struct {
	root    string
	pattern string
	params  map[string]string
}

type CacheResetStatus struct {
	RawUrl string
	Status string
}

func NewCacheReset(root, pattern string, params map[string]string) *CacheReset {
	return &CacheReset{
		root:    root,
		pattern: pattern,
		params:  params,
	}
}

func (cr *CacheReset) Reset() []*CacheResetStatus {
	var res []*CacheResetStatus

	urls := cr.getUrls()
	for _, url := range urls {
		status := "done"
		err := cacheServerResponseInternaly(url.String())
		if err != nil {
			status = err.Error()
		}
		res = append(res, &CacheResetStatus{
			RawUrl: url.String(),
			Status: status,
		})
	}

	notifyCacheResetStatus(res)
	return res
}

func (cr *CacheReset) getUrls() []*url.URL {
	var res []*url.URL
	filepath.Walk(
		configs.GetEnvConfigs().CachePath,
		func(path string, info os.FileInfo, err error) error {
			filename := filepath.Base(path)
			filename = strings.TrimSuffix(filename, filepath.Ext(filename))
			if strings.HasPrefix(filename, url.QueryEscape(fmt.Sprintf("/%s/%s", cr.root, cr.pattern))) {
				url := parseUrlFromCacheFilename(filename)
				if url != nil && cr.matchUrl(url) {
					res = append(res, url)
				}
			}
			return nil
		})
	return res
}

func (cr *CacheReset) matchUrl(url *url.URL) bool {
	if cr.params["country"] == "intl" {
		return strings.HasSuffix(url.Path, "internationalhome")
	}

	urlQuery := url.Query()
	for paramName, paramValue := range cr.params {
		if urlQuery.Get(paramName) != paramValue {
			return false
		}
	}

	return true
}

func parseUrlFromCacheFilename(filename string) *url.URL {
	rawUrl, err := url.QueryUnescape(filename)
	if err != nil {
		return nil
	}

	res, err := url.Parse(rawUrl)
	if err != nil {
		return nil
	}

	return res
}

func cacheServerResponseInternaly(url string) error {
	envConfigs := configs.GetEnvConfigs()
	req, err := http.NewRequest("GET", envConfigs.ResolveSamanthaAbsoluteUrl(url), nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Reset-Cache", "true")
	_, err = (&http.Client{}).Do(req)
	if err != nil {
		return err
	}

	return nil
}

func notifyCacheResetStatus(resetStatus []*CacheResetStatus) {
	for _, crStatus := range resetStatus {
		if crStatus.Status != "done" {
			utils.LogError(crStatus.Status, "cache-reset-error")
		}
		log.Println(crStatus.RawUrl, crStatus.Status)
	}
}
