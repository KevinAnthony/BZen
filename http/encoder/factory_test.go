package encoder_test

import (
	"net/http"
	"testing"

	"github.com/kevinanthony/bzen/http/encoder"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFactory(t *testing.T) {
	t.Parallel()

	Convey("NewFactory", t, func() {
		Convey("should return new factory", func() {
			f := encoder.NewFactory()

			So(f, ShouldImplement, (*encoder.Factory)(nil))
		})
	})
}

func TestFactoryMock_Create(t *testing.T) {
	t.Parallel()

	Convey("Create", t, func() {
		resp := &http.Response{
			Header: http.Header{},
		}
		factory := encoder.NewFactory()

		Convey("should return json encoder", func() {
			Convey("when content-type is empty", func() {
				actual := factory.Create(resp)

				So(actual, ShouldHaveSameTypeAs, encoder.NewJSON())
			})
			Convey("when content-type is application/json", func() {
				resp.Header.Add("content-type", encoder.ApplicationJSON)

				actual := factory.Create(resp)

				So(actual, ShouldHaveSameTypeAs, encoder.NewJSON())
			})
		})
		Convey("should return xml encoder", func() {
			Convey("when content-type is application/xml", func() {
				resp.Header.Add("content-type", encoder.ApplicationXML)

				actual := factory.Create(resp)

				So(actual, ShouldHaveSameTypeAs, encoder.NewXML())
			})
		})
	})
}
