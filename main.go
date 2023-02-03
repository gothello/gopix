package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	token = "mytoken"
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

func GetMethodsPayments() error {
	h := map[string]string{

		"accept":        "application/json",
		"content-type":  "application/json",
		"Authorization": `"Bearer ` + token + `"`,
	}

	opt := NewOptions("GET", "https://api.mercadopago.com/v1/payment_methods", "", 3000, h)

	data, err := opt.Request()
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil

}

func CreatePix(amount int, desc string) error {
	h := map[string]string{
		"accept":        "application/json",
		"content-type":  "application/json",
		"Authorization": `"Bearer ` + token + `"`,
	}

	exp := time.Now().Add(time.Hour / 2).Format("2006-01-02T15:04:05.000-04:00")

	body := fmt.Sprintf(`"transaction_amount":"%v","description": "%s","payment_method_id": "pix", "date_of_expiration": "%s",`, amount, desc, exp)

	opt := NewOptions("POST", "https://api.mercadopago.com/v1/payments", body, 0, h)

	data, err := opt.Request()
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

func main() {

	fmt.Println(time.Now().Add(time.Second * 1800).Format("2006-01-02T15:04:05.000-04:00"))

}
