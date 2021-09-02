package encoding_test

import (
	"net/http"
	"testing"

	"github.com/kevinanthony/bzen/http/encoding"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEncoding(t *testing.T) {
	t.Parallel()

	Convey("NewEncoding", t, func() {
		resp := &http.Response{
			Header: http.Header{},
		}

		Convey("should return json encoding", func() {
			Convey("when accept is empty", func() {
				actual := encoding.NewEncoding(resp)

				So(actual, ShouldHaveSameTypeAs, encoding.NewJSONDecoder())
			})
			Convey("when accept is application/json", func() {
				resp.Header.Add("content-type", encoding.AcceptJSON)
				actual := encoding.NewEncoding(resp)

				So(actual, ShouldHaveSameTypeAs, encoding.NewJSONDecoder())
			})
		})
		Convey("should return xml encoding", func() {
			Convey("when accept is application/xml", func() {
				resp.Header.Add("content-type", encoding.AcceptXML)
				actual := encoding.NewEncoding(resp)

				So(actual, ShouldHaveSameTypeAs, encoding.NewXMLDecoder())
			})
		})
	})
}
