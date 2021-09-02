package http

import (
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

var _ Request = (*RequestMock)(nil)

type RequestMock struct {
	mock.Mock
}

func (r *RequestMock) Go(ctx context.Context, client http.Client, v interface{}) error {
	return r.Called(ctx, client, v).Error(0)
}
