package http

import (
	"context"

	"github.com/stretchr/testify/mock"
)

var _ Request = (*RequestMock)(nil)

type RequestMock struct {
	mock.Mock
}

func (r *RequestMock) Parameter(key, value string) Request {
	r.Called(key, value)

	return r
}

func (r *RequestMock) Header(key, value string) Request {
	r.Called(key, value)

	return r
}

func (r *RequestMock) Go(ctx context.Context, v interface{}) error {
	return r.Called(ctx, v).Error(0)
}

func (r *RequestMock) Post() Request {
	r.Called()

	return r
}

func (r *RequestMock) Get() Request {
	r.Called()

	return r
}

func (r *RequestMock) Put() Request {
	r.Called()

	return r
}

func (r *RequestMock) Delete() Request {
	r.Called()

	return r
}

func (r *RequestMock) Domain(s string) Request {
	r.Called(s)

	return r
}

func (r *RequestMock) Path(s string) Request {
	r.Called(s)

	return r
}
