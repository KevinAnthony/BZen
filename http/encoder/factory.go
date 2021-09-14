package encoder

import (
	"mime"
	"net/http"
)

type Factory interface {
	Create(resp *http.Response) Encoder
}

type factory struct{}

func NewFactory() Factory {
	return factory{}
}

func (f factory) Create(resp *http.Response) Encoder {
	switch mediaType, _, _ := mime.ParseMediaType(resp.Header.Get("content-type")); mediaType {
	case ApplicationXML:
		return NewXML()
	case ApplicationJSON:
		return NewJSON()
	default:
		return NewJSON()
	}
}
