package http

import (
	"context"
	native "net/http"
	"net/url"

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
}

type request struct {
	err    error
	client Client
	domain string
	path   string
	method MethodType
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

func NewRequest(client Client) Request {
	r := &request{
		method: MethodGet,
	}

	if client == nil {
		r.setErrStr("native client is nil")
	}

	return r
}

func (r *request) Go(ctx context.Context, v interface{}) error {
	if r.err != nil {
		return r.err
	}

	req := &native.Request{
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

	return r.client.Do(req, v)
}

func (r *request) setErrStr(s string) {
	if r.err != nil {
		r.err = errors.New(s)
	}
}
