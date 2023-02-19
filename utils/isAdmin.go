package utils

import (
	"net/http"
	"regexp"
)

var (
	SecureKey = "ok"
)

func IsAdmin(handler func(w http.ResponseWriter, h *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile("Baerer " + SecureKey)
		auth := r.Header.Get("Authorization")

		if !re.Match([]byte(auth)) {
			ToErro(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		handler(w, r)
		return
	}
}
