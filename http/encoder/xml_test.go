package encoder_test

import (
	"reflect"
	"testing"

	"github.com/kevinanthony/bzen/http/encoder"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewXML(t *testing.T) {
	t.Parallel()

	Convey("NewXML", t, func() {
		Convey("should return type xmlEncoding", func() {
			actual := encoder.NewXML()

			So(reflect.TypeOf(actual).String(), ShouldEqual, "encoder.xmlEncoder")
		})
	})
}

func TestXML_Encode(t *testing.T) {
	t.Parallel()

	Convey("Encode", t, func() {
		enc := encoder.NewXML()

		Convey("should return json string for structure", func() {
			actual, err := enc.Encode(newTestStruct())

			So(err, ShouldBeNil)
			So(string(actual), ShouldEqual, xmlString())
		})
	})
}

func TestXML_Decode(t *testing.T) {
	t.Parallel()

	Convey("Decode", t, func() {
		enc := encoder.NewXML()

		expected := newTestStruct()
		expected.Map = nil

		Convey("should return json string for structure", func() {
			var actual testStruct

			err := enc.Decode([]byte(xmlString()), &actual)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
	})
}

// nolint: lll // i don't want to format this
func xmlString() string {
	return `<testStruct><slice>one</slice><slice>two</slice><slice>ah-ha-ha</slice><string>something</string><int>42</int><flat>3.1415</flat><time>1989-11-09T18:01:00Z</time><struct><string>else</string><int>99</int><flat>2.7182</flat></struct></testStruct>`
}
