package market

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ApiClient[T any] struct {
	httpClient http.Client
	baseUrl    string
	authToken  string
}

type ApiConfig struct {
	ApiUrl    string
	AuthToken string
}

func NewApiClient[T any](config ApiConfig) ApiClient[T] {
	return ApiClient[T]{
		httpClient: http.Client{},
		baseUrl:    config.ApiUrl,
		authToken:  config.AuthToken,
	}
}

func (c *ApiClient[T]) BuildUrl(path string, params map[string]string) string {
	url := fmt.Sprintf("%s%s", c.baseUrl, path)
	var paramsStrings []string

	for pname, pvalue := range params {
		paramsStrings = append(paramsStrings, fmt.Sprintf("%s=%s", pname, pvalue))
	}

	return fmt.Sprintf("%s?%s", url, strings.Join(paramsStrings, "&"))
}

func parseResult[T any](data []byte) (*T, error) {
	var resultData T
	err := json.Unmarshal(data, &resultData)

	if err != nil {
		return nil, err
	}

	return &resultData, nil
}

func (client *ApiClient[T]) Get(path string, params map[string]string, headers map[string]string) (*T, error) {
	request, err := client.buildRequest("GET", path, params, headers)
	if err != nil {
		return nil, err
	}
	response, err := client.httpClient.Do(request)
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) Post(path string, params map[string]string, headers map[string]string) (*T, error) {
	request, err := client.buildRequest("POST", path, params, headers)
	if err != nil {
		return nil, err
	}
	response, err := client.httpClient.Do(request)
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) parseResponse(response *http.Response, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	parsed, err := parseResult[T](bodyData)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (client *ApiClient[T]) buildRequest(requestMethod string, path string, params map[string]string, headers map[string]string) (*http.Request, error) {
	url := client.BuildUrl(path, params)
	paramsBuffer := bytes.NewBuffer(make([]byte, 0))

	if len(params) > 0 {
		bodyParams, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		paramsBuffer = bytes.NewBuffer(bodyParams)
	}

	request, err := http.NewRequest(requestMethod, url, paramsBuffer)
	if err != nil {
		return nil, err
	}

	for headerKey, headerValue := range headers {
		request.Header.Add(headerKey, headerValue)
	}

	return request, nil
}
