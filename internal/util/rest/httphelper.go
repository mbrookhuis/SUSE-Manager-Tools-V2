// Package rest - rest api helper
package rest

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// HTTPHelperStruct - HTTP call
type HTTPHelperStruct struct {
	Body       []byte
	StatusCode int
	Cookies    []*http.Cookie
}

// HTTPHelper - rest api helper
//
// param: log
// param: retrycount
// param: requsetBody
// param: method
// param: url
// param: insecure
// param: header
// return:
func HTTPHelper(log *zap.Logger, retrycount int, requsetBody []byte, method string, url string, insecure bool, header ...map[string]string) (*HTTPHelperStruct, error) {
	reqBody := bytes.NewBuffer(requsetBody)

	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	for k := range header {
		for i, j := range header[k] {
			req.Header.Add(i, j)
		}
	}
	var res *http.Response
	retry := 0
	log.Info("Connecting to: ", zap.Any("url", url))
	for ok := true; ok; ok = (err != nil || retry == retrycount) {
		res, err = client.Do(req)
		if err != nil {
			log.Info("Retrying connect to: ", zap.Any("retry count", retry))
			time.Sleep(time.Duration(retry*2) * time.Second)
			if retry == retrycount {
				log.Error("Failed connect to: ", zap.Any("url", url))
				return nil, err
			}
		}
		retry++
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	httpHelperStruct := HTTPHelperStruct{
		Body:       body,
		StatusCode: res.StatusCode,
		Cookies:    res.Cookies(),
	}
	return &httpHelperStruct, nil
}
