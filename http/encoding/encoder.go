package encoding

import (
	"mime"
	"net/http"
)

type (
	AcceptType = string
)

const (
	AcceptJSON AcceptType = "application/json"
	AcceptXML  AcceptType = "application/xml"
)

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, dst interface{}) error
}

func NewEncoding(resp *http.Response) Encoder {
	mediaType, _, _ := mime.ParseMediaType(resp.Header.Get("content-type"))
	switch mediaType {
	case AcceptXML:
		return NewXMLDecoder()
	case AcceptJSON:
		return NewJSONDecoder()
	default:
		return NewJSONDecoder()
	}
}
