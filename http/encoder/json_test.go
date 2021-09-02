package encoder_test

import (
	"reflect"
	"testing"

	"github.com/kevinanthony/bzen/http/encoder"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewJson(t *testing.T) {
	t.Parallel()

	Convey("NewJson", t, func() {
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
	})
}

func TestJSON_Decode(t *testing.T) {
	t.Parallel()

	Convey("Decode", t, func() {
		enc := encoder.NewJSON()

		Convey("should return json string for structure", func() {
			var actual testStruct

			err := enc.Decode([]byte(jsonString()), &actual)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, newTestStruct())
		})
	})
}

// nolint: lll // i don't want to format this
func jsonString() string {
	return `{"map":{"1":1,"2":2,"3":3},"slice":["one","two","ah-ha-ha"],"string":"something","int":42,"float":3.1415,"time":"1989-11-09T18:01:00Z","sub":{"string":"else","int":99,"float":2.7182}}`
}
