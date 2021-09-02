package encoder

import "encoding/xml"

type xmlEncoder struct {
}

func NewXML() Encoder {
	return xmlEncoder{}
}

func (x xmlEncoder) Encode(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

func (x xmlEncoder) Decode(data []byte, dst interface{}) error {
	return xml.Unmarshal(data, dst)
}
