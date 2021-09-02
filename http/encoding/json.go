package encoding

import jit "github.com/json-iterator/go"

type jsonDecoder struct {
	jit jit.API
}

func NewJSONDecoder() Encoder {
	return jsonDecoder{
		jit: jit.ConfigCompatibleWithStandardLibrary,
	}
}

func (j jsonDecoder) Encode(data interface{}) ([]byte, error) {
	return j.jit.Marshal(data)
}

func (j jsonDecoder) Decode(data []byte, dst interface{}) error {
	return j.jit.Unmarshal(data, dst)
}
