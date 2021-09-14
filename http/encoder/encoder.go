package encoder

import (
	"mime"
	"net/http"
)

type (
	AcceptType = string
)

const (
	ApplicationJSON AcceptType = "application/json"
	ApplicationXML  AcceptType = "application/xml"
)

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, dst interface{}) error
}

func New(resp *http.Response) Encoder {
	switch mediaType, _, _ := mime.ParseMediaType(resp.Header.Get("content-type")); mediaType {
	case ApplicationXML:
		return NewXML()
	case ApplicationJSON:
		return NewJSON()
	default:
		return NewJSON()
	}
}
