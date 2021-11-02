package encoder

import (
	jit "github.com/json-iterator/go"
)

type jsonEncoder struct {
	jit jit.API
}

func NewJSON() Encoder {
	return jsonEncoder{
		jit: jit.ConfigCompatibleWithStandardLibrary,
	}
}

func (j jsonEncoder) Encode(data interface{}) ([]byte, error) {
	return j.jit.Marshal(data)
}

func (j jsonEncoder) Decode(data []byte, dst interface{}) error {
	return j.jit.Unmarshal(data, dst)
}
