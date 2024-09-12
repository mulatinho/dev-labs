package core

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	ContentType              = "application/json"
	UserAgent                = "golabs/restcall"
	HTTP_METHODS_ALLOWED int = 5
)

type RestCallData struct {
	RestRequest  *http.Request
	RestResponse *http.Response
	RestError    error
	RestBody     []byte
	RestStatus   int
}

func RestCall(method string, url string, body []byte) *RestCallData {
	bodyString := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, url, bodyString)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	if err != nil {
		return &RestCallData{
			RestRequest:  req,
			RestResponse: nil,
			RestError:    err,
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RestCallData{
			RestRequest:  req,
			RestResponse: res,
			RestError:    err,
		}
	}

	responseBody, _ := ioutil.ReadAll(res.Body)
	return &RestCallData{
		RestRequest:  req,
		RestResponse: res,
		RestError:    err,
		RestStatus:   res.StatusCode,
		RestBody:     responseBody,
	}
}
