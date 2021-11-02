package encoder

import (
	"github.com/stretchr/testify/mock"
)

var _ Encoder = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func (e *Mock) Encode(data interface{}) ([]byte, error) {
	args := e.Called(data)

	var bts []byte
	if item := args.Get(0); item != nil {
		bts = item.([]byte)
	}

	return bts, args.Error(1)
}

func (e *Mock) Decode(data []byte, dst interface{}) error {
	return e.Called(data, dst).Error(0)
}
