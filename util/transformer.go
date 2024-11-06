package util

import "net/url"

func TransformParams(params map[string]string) string {

	baseURL := "https://example.com/updated?"

	newParams := url.Values{}
	for key, value := range params {

		newKey := "transformed_" + key
		newValue := "transformed_" + value

		newParams.Set(newKey, newValue)
	}
	return baseURL + newParams.Encode()
}
