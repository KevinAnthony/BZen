package http_test

import (
	"reflect"
	"testing"

	"github.com/kevinanthony/bzen/http"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRequest(t *testing.T) {
	t.Parallel()

	Convey("NewRequest", t, func() {
		Convey("should return type jsonEncoding", func() {
			actual := http.NewRequest(http.MethodPost, "google.com", "search")

			So(reflect.TypeOf(actual).String(), ShouldEqual, "*http.request")
		})
	})
}

func TestRequest_Go(t *testing.T) {
	t.Parallel()

	Convey("Go", t, func() {
		// TODO no unit test yet
	})
}
