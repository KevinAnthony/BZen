package http

import (
	"io"
	"net/http"

	"github.com/stretchr/testify/mock"
)

var _ Client = (*ClientMock)(nil)
var _ Native = (*NativeMock)(nil)
var _ io.ReadCloser = (*BodyMock)(nil)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Do(req *http.Request, v interface{}) error {
	return c.Called(req, v).Error(0)
}

type NativeMock struct {
	mock.Mock
}

func (n *NativeMock) Do(req *http.Request) (*http.Response, error) {
	args := n.Called(req)

	var resp *http.Response
	if item := args.Get(0); item != nil {
		resp = item.(*http.Response)
	}

	return resp, args.Error(1)
}

type BodyMock struct {
	mock.Mock
}

func (b *BodyMock) Read(p []byte) (n int, err error) {
	args := b.Called(p)

	return args.Int(0), args.Error(1)
}

func (b *BodyMock) Close() error {
	return b.Called().Error(0)
}
