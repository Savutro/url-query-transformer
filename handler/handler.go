package handler

import (
	"fmt"
	"net/http"

	"github.com/Shared-Info-Platform/url-query-transformer/proxy"
	"github.com/Shared-Info-Platform/url-query-transformer/util"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	queryParams := util.ParseQueryParams(r.URL.Query())

	if queryParams["redirect"] == "true" {
		newURL := util.TransformParams(queryParams)

		fmt.Printf("Redirecting to: %s\n", newURL)

		util.RedirectToNewURL(w, r, newURL)
	} else {
		newURL := util.TransformParams(queryParams)

		fmt.Printf("Fetching and serving content from: %s\n", newURL)

		err := proxy.FetchAndServeContent(w, newURL)
		if err != nil {
			http.Error(w, "Error fetching content", http.StatusInternalServerError)
		}
	}
}
