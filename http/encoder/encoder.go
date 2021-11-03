package encoder

type (
	AcceptType = string
)

const (
	ApplicationJSON AcceptType = "application/json"
	ApplicationXML  AcceptType = "application/xml"
	ApplicaitonYAML AcceptType = "application/yaml"
)

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, dst interface{}) error
}
