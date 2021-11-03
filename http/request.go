package http

import (
	"context"
	"fmt"
	native "net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type Request interface {
	Go(ctx context.Context, v interface{}) error

	Post() Request
	Get() Request
	Put() Request
	Delete() Request

	Domain(string) Request
	Path(string) Request
	Parameter(string, string) Request

	Header(string, string) Request
}

type request struct {
	err    error
	client Client
	domain string
	path   string
	method MethodType

	parameters map[string]string
	headers    map[string]string
}

func (r *request) Post() Request {
	r.method = MethodPost

	return r
}

func (r *request) Get() Request {
	r.method = MethodGet

	return r
}

func (r *request) Put() Request {
	r.method = MethodPut

	return r
}

func (r *request) Delete() Request {
	r.method = MethodDelete

	return r
}

func (r *request) Domain(s string) Request {
	r.domain = s

	return r
}

func (r *request) Path(s string) Request {
	r.path = s

	return r
}

func (r *request) Parameter(pattern, value string) Request {
	r.parameters[pattern] = value

	return r
}

func (r *request) Header(header, value string) Request {
	r.headers[header] = value

	return r
}

func NewRequest(client Client) Request {
	r := &request{
		method:     MethodGet,
		parameters: map[string]string{},
		headers:    map[string]string{},
	}

	if client == nil {
		r.setErrStr("native client is nil")
	}

	r.client = client

	return r
}

func (r *request) Go(ctx context.Context, out interface{}) error {
	if r.err != nil {
		return r.err
	}

	for k, v := range r.parameters {
		if !strings.Contains(r.path, k) {
			return fmt.Errorf("missing parameter %s in path", k)
		}

		r.path = strings.ReplaceAll(r.path, k, v)
	}

	req := &native.Request{
		Method: string(r.method),
		URL: &url.URL{
			Scheme: "https",
			Host:   r.domain,
			Path:   r.path,
		},
		Proto:  "https",
		Header: native.Header{},
		Body:   nil,
	}

	for k, v := range r.headers {
		req.Header.Add(k, v)
	}

	req = req.WithContext(ctx)

	return r.client.Do(req, out)
}

func (r *request) setErrStr(s string) {
	if r.err != nil {
		r.err = errors.New(s)
	}
}
