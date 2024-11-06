package util

import (
	"net/http"
)

func RedirectToNewURL(w http.ResponseWriter, r *http.Request, newURL string) {
	http.Redirect(w, r, newURL, http.StatusFound)
}
