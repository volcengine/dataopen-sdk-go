/*
 * Copyright 2023 DataOpen SDK Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
)

type ParamsValueType interface{}

type ClientStruct struct {
	app_id, app_secret   string
	url, expiration string
	_access_token        string
	_ttl, _token_time    int64
}

func Client(app_id, app_secret, url, expiration string) *ClientStruct {
	if url == "" {
		url = "https://analytics.volcengineapi.com"
	}

	if expiration == "" {
		expiration = "1800"
	}

	return &ClientStruct{
		app_id: app_id, app_secret: app_secret, url: url, expiration: expiration,
		_ttl: 0, _access_token: "", _token_time: 0,
	}
}

type ApiResponse struct {
	Code    int
	Message string
	Data    *struct {
		Access_token string
		Ttl          int64
	}
}

func (c *ClientStruct) Request(serviceUrl, method string, headers map[string]string, params map[string]ParamsValueType, body map[string]interface{}) (map[string]interface{}, error) {
	upperCaseMethod := strings.ToUpper(method)

	if c._access_token == "" || !c._valid_token() {
		if err := c.GetToken(); err != nil {
			return nil, err
		}
	}

	newHeaders := map[string]string{
		"Authorization": c._access_token,
		"Content-Type":  "application/json",
		"x-sdk-source":  "go",
	}

	for key, value := range headers {
		newHeaders[key] = value
	}

	completedUrl := c.url + serviceUrl
	queryUrl := c._joint_query(completedUrl, params)

	var resp map[string]interface{}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(upperCaseMethod, queryUrl, bytes.NewBuffer(jsonBody))

	for k, v := range newHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	json.NewDecoder(response.Body).Decode(&resp)

	return resp, nil
}

func (c *ClientStruct) GetToken() error {
	authorizationUrl := "/dataopen/open-apis/v1/authorization"
	completedUrl := c.url + authorizationUrl

	mapBody := map[string]string{
		"app_id":     c.app_id,
		"app_secret": c.app_secret,
	}

	body, _ := json.Marshal(mapBody)
	req, _ := http.NewRequest("POST", completedUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var apiResponse ApiResponse
	body, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &apiResponse)

	tokenTime := time.Now().Unix()
	if apiResponse.Code == 200 && apiResponse.Data != nil {
		c._ttl = apiResponse.Data.Ttl
		c._token_time = tokenTime
		c._access_token = apiResponse.Data.Access_token
	}

	return nil
}

func (c *ClientStruct) IsAuthenticated() bool {
	return c._access_token != ""
}

func (c *ClientStruct) _joint_query(url string, params map[string]ParamsValueType) string {
	var paramStr []string
	for key, val := range params {
		paramStr = append(paramStr, key+"="+val.(string))
	}
	return url + "?" + strings.Join(paramStr, "&")
}

func (c *ClientStruct) _valid_token() bool {
	currentTime := time.Now().Unix()
	return currentTime-c._token_time < c._ttl
}
