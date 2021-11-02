package encoder

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

var _ Factory = (*FactoryMock)(nil)

type FactoryMock struct {
	mock.Mock
}

func (f *FactoryMock) Create(resp *http.Response) Encoder {
	args := f.Called(resp)

	var enc Encoder
	if item := args.Get(0); item != nil {
		enc = item.(Encoder)
	}

	return enc
}
