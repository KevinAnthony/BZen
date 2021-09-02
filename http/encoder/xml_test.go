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
		Convey("should return error when", func() {
			Convey("struct cannot be marshalled", func() {
				actual, err := enc.Encode(badXMLStruct{M: map[int]int{1: 1}})

				So(actual, ShouldBeNil)
				So(err, ShouldBeError, "xml: unsupported type: map[int]int")
			})
		})
	})
}

func TestXML_Decode(t *testing.T) {
	t.Parallel()

	Convey("Decode", t, func() {
		enc := encoder.NewXML()

		expected := newTestStruct()
		expected.Map = nil
		var actual testStruct

		Convey("should return json string for structure", func() {
			err := enc.Decode([]byte(xmlString()), &actual)

			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should return error", func() {
			Convey("when data is empty", func() {
				err := enc.Decode(nil, &actual)

				So(err, ShouldBeError, "EOF")
			})
			Convey("when out is not pointer", func() {
				err := enc.Decode(nil, actual)

				So(err, ShouldBeError, "non-pointer passed to Unmarshal")
			})
		})
	})
}

func xmlString() string {
	return `<testStruct><slice>one</slice><slice>two</slice><slice>ah-ha-ha</slice><string>something</string><int>42</int><flat>3.1415</flat><time>1989-11-09T18:01:00Z</time><struct><string>else</string><int>99</int><flat>2.7182</flat></struct></testStruct>`
}

type badXMLStruct struct {
	M map[int]int
}
