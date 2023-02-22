package utils

import (
	"net/http"
	"regexp"
)

var (
	SECRET_AUTH_KEY = ""
)

func IsAdmin(handler func(w http.ResponseWriter, h *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile("Baerer " + SECRET_AUTH_KEY)
		auth := r.Header.Get("Authorization")

		if !re.MatchString(auth) {
			ToErro(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		handler(w, r)
		return
	}
}
