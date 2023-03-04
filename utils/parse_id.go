package utils

import (
	"errors"
	"net/http"
	"strconv"
)

func ParseID(r *http.Request) (int64, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return 0, errors.New("send number id invalid")
	}

	parse, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, errors.New("send number id invalid")
	}

	return parse, nil
}
