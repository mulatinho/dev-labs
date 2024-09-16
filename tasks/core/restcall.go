/*
 * The golabs/tasks application.
 *
 * Copyright (C) 2024 Alexandre Mulatinho <alex@mulatinho.net>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either  version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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
