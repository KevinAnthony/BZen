package encoder_test

import (
	"reflect"
	"testing"

	"github.com/kevinanthony/bzen/http/encoder"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewJSON(t *testing.T) {
	t.Parallel()

	Convey("NewJSon", t, func() {
		Convey("should return type jsonEncoding", func() {
			actual := encoder.NewJSON()

			So(reflect.TypeOf(actual).String(), ShouldEqual, "encoder.jsonEncoder")
		})
	})
}

func TestJSON_Encode(t *testing.T) {
	t.Parallel()

	Convey("Encode", t, func() {
		enc := encoder.NewJSON()

		Convey("should return json string for structure", func() {
			actual, err := enc.Encode(newTestStruct())

			So(err, ShouldBeNil)
			So(string(actual), ShouldEqual, jsonString())
		})
		Convey("should return error when", func() {
			Convey("struct cannot be marshalled", func() {
				actual, err := enc.Encode(badJSONStruct{F: func() {}})

				So(actual, ShouldBeNil)
				So(err, ShouldBeError, "encoder_test.badJSONStruct.F:  Ffunc() is unsupported type")
			})
		})
	})
}

func TestJSON_Decode(t *testing.T) {
	t.Parallel()

	Convey("Decode", t, func() {
		enc := encoder.NewJSON()

		var actual testStruct

		Convey("should return json string for structure", func() {
			err := enc.Decode([]byte(jsonString()), &actual)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, newTestStruct())
		})
		Convey("should return error", func() {
			Convey("when data is empty", func() {
				err := enc.Decode(nil, &actual)

				So(err, ShouldBeError, "readObjectStart: expect { or n, but found \u0000, error found in #0 byte of ...||..., bigger context ...||...")
			})
			Convey("when out is not pointer", func() {
				err := enc.Decode(nil, actual)

				So(err, ShouldBeError, "ReadVal: can only unmarshal into pointer, error found in #0 byte of ...||..., bigger context ...||...")
			})
		})
	})
}

func jsonString() string {
	return `{"map":{"1":1,"2":2,"3":3},"slice":["one","two","ah-ha-ha"],"string":"something","int":42,"float":3.1415,"time":"1989-11-09T18:01:00Z","sub":{"string":"else","int":99,"float":2.7182}}`
}

type badJSONStruct struct {
	F func()
}
