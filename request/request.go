package request

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (o *Options) Request() *Response {

	req, err := http.NewRequest(o.Method, o.Url, strings.NewReader(o.Body))
	if err != nil {
		return &Response{
			Body:     nil,
			Err:      err,
			Response: nil,
		}
	}

	for p, v := range o.Headers {
		req.Header.Add(p, v)
	}

	if _, ok := o.Headers["Host"]; ok {
		req.Host = o.Headers["Host"]
	}

	c := &http.Client{}
	if o.Timeout != 0 {
		c.Timeout = time.Millisecond * time.Duration(o.Timeout)
	}

	resp, err := c.Do(req)
	if err != nil {
		return &Response{
			Body:     nil,
			Err:      err,
			Response: nil,
		}
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Body:     nil,
			Err:      err,
			Response: nil,
		}
	}

	return &Response{
		Body:     data,
		Err:      nil,
		Response: resp.Request,
	}

}
