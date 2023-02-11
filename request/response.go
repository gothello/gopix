package request

import "net/http"

type Response struct {
	Body     []byte
	Err      error
	Response *http.Response
}
