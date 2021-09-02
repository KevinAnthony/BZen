package encoder

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

func New(resp *http.Response) Encoder {
	mediaType, _, _ := mime.ParseMediaType(resp.Header.Get("content-type"))
	switch mediaType {
	case AcceptXML:
		return NewXML()
	case AcceptJSON:
		return NewJSON()
	default:
		return NewJSON()
	}
}
