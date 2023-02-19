package request

type Options struct {
	Method  string
	Url     string
	Timeout int
	Body    string
	Headers map[string]string
}

func NewOptions(method, url, body string, timeout int, headers map[string]string) *Options {
	return &Options{
		method,
		url,
		timeout,
		body,
		headers,
	}
}
