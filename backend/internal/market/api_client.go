package market

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ApiClient[T any] struct {
	baseUrl   string
	authToken string
}

type ApiConfig struct {
	ApiUrl    string
	AuthToken string
}

func NewApiClient[T any](config ApiConfig) ApiClient[T] {
	return ApiClient[T]{
		baseUrl:   config.ApiUrl,
		authToken: config.AuthToken,
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
	fmt.Println(string(data))
	err := json.Unmarshal(data, &resultData)

	fmt.Println(err.Error())

	if err != nil {
		return nil, err
	}

	return &resultData, nil
}

func (client *ApiClient[T]) Get(path string, params map[string]string) (*T, error) {
	url := client.BuildUrl(path, params)
	fmt.Println(url)
	response, err := http.Get(url)

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
