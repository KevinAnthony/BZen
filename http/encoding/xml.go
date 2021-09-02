package encoding

import "encoding/xml"

type xmlDecoder struct {
}

func NewXMLDecoder() Encoder {
	return xmlDecoder{}
}

func (x xmlDecoder) Encode(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

func (x xmlDecoder) Decode(data []byte, dst interface{}) error {
	return xml.Unmarshal(data, dst)
}
