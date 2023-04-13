package brapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiClient[T any] struct {
	baseUrl string
}

type ApiResult[T any] struct {
	Results []T `json:"results"`
}

func NewApiClient[T any]() ApiClient[T] {
	return ApiClient[T]{baseUrl: "https://brapi.dev/api"}
}

func (c *ApiClient[T]) makeUrl(path string) string {
	return fmt.Sprintf("%s%s", c.baseUrl, path)
}

func parseResult[T any](data []byte) ApiResult[T] {
	var resultData ApiResult[T]
	// fmt.Println(string(data))
	err := json.Unmarshal(data, &resultData)
	// fmt.Println(resultData.Results)
	if err != nil {
		fmt.Println(err)
	}
	return resultData
}

func (c *ApiClient[T]) Get(path string) ApiResult[T] {
	url := c.makeUrl(path)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return parseResult[T](bodyData)
}
