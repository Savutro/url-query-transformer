package proxy

import (
	"io"
	"net/http"
)

func FetchAndServeContent(w http.ResponseWriter, newURL string) error {

	resp, err := http.Get(newURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	return err
}
