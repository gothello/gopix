package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

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

func (o *Options) Request() ([]byte, error) {
	var data = []byte{}

	req, err := http.NewRequest(o.Method, o.Url, strings.NewReader(o.Body))
	if err != nil {
		return data, err
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
		return data, err
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	return data, nil

}

func main() {

	h := map[string]string{}

	opt := NewOptions("GET", "https://www.google.com.br", "", 1000, h)

	d, err := opt.Request()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(d))

}
