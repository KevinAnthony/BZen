package encoder_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/kevinanthony/bzen/http/encoder"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	t.Parallel()

	Convey("New", t, func() {
		resp := &http.Response{
			Header: http.Header{},
		}

		Convey("should return json encoder", func() {
			Convey("when content-type is empty", func() {
				actual := encoder.New(resp)

				So(actual, ShouldHaveSameTypeAs, encoder.NewJSON())
			})
			Convey("when content-type is application/json", func() {
				resp.Header.Add("content-type", encoder.AcceptJSON)

				actual := encoder.New(resp)

				So(actual, ShouldHaveSameTypeAs, encoder.NewJSON())
			})
		})
		Convey("should return xml encoder", func() {
			Convey("when content-type is application/xml", func() {
				resp.Header.Add("content-type", encoder.AcceptXML)

				actual := encoder.New(resp)

				So(actual, ShouldHaveSameTypeAs, encoder.NewXML())
			})
		})
	})
}

type testStruct struct {
	Map    map[string]int `json:"map" xml:"-"`
	Slice  []string       `json:"slice" xml:"slice"`
	String string         `json:"string" xml:"string"`
	Int    int            `json:"int" xml:"int"`
	Float  float32        `json:"float" xml:"flat"`
	Time   time.Time      `json:"time" xml:"time"`
	Struct subStrict      `json:"sub" xml:"struct"`
}

type subStrict struct {
	String string  `json:"string" xml:"string"`
	Int    int     `json:"int" xml:"int"`
	Float  float32 `json:"float" xml:"flat"`
}

func newTestStruct() testStruct {
	return testStruct{
		Map:    map[string]int{"1": 1, "2": 2, "3": 3},
		Slice:  []string{"one", "two", "ah-ha-ha"},
		String: "something",
		Int:    42,
		Float:  3.1415,
		Time:   time.Date(1989, 11, 9, 18, 01, 00, 00, time.UTC),
		Struct: subStrict{
			String: "else",
			Int:    99,
			Float:  2.7182,
		},
	}
}
