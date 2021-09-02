package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kevinanthony/bzen/http/encoding"
)

type Request interface {
	Go(ctx context.Context, client http.Client, v interface{}) error
}

type request struct {
	domain string
	path   string
	method MethodType
}

func NewRequest(method MethodType, domain, path string) Request {
	return &request{
		method: method,
		domain: domain,
		path:   path,
	}
}

// TODO http client needs to be abstracted and mocked.
func (r request) Go(ctx context.Context, client http.Client, v interface{}) error {
	req := &http.Request{
		Method: string(r.method),
		URL: &url.URL{
			Scheme: "https",
			Host:   r.domain,
			Path:   r.path,
		},
		Proto:  "https",
		Header: nil,
		Body:   nil,
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return encoding.NewEncoding(resp).Decode(bts, v)
}
