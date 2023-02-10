package utils

import "net/http"

func ToErro(w http.ResponseWriter, s string, code int){
	e := map[string]string{
		"error":s,
	}

	ToJson(w, e, code)
}
