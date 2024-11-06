package util

import "net/url"

func ParseQueryParams(query url.Values) map[string]string {
	params := make(map[string]string)
	for key, values := range query {
		params[key] = values[0]
	}
	return params
}
