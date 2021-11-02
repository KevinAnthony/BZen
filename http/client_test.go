package http_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	native "net/http"
	"testing"

	"github.com/kevinanthony/bzen/http"
	"github.com/kevinanthony/bzen/http/encoder"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	Convey("NewClient", t, func() {
		clnt := &native.Client{}
		fact := &encoder.FactoryMock{}

		Convey("should return new client", func() {
			f := func() { http.NewClient(clnt, fact) }

			So(f, ShouldNotPanic)
		})
		Convey("should panic when", func() {
			Convey("client transport is nil", func() {
				f := func() { http.NewClient(nil, fact) }

				So(f, ShouldPanicWith, "http client is required")
			})
			Convey("encoder factory is nil", func() {
				f := func() { http.NewClient(clnt, nil) }

				So(f, ShouldPanicWith, "encoding factory is required")
			})
		})
	})
}

func TestClient_Do(t *testing.T) {
	t.Parallel()

	Convey("Do", t, func() {
		req, err := native.NewRequest(http.MethodGet, "https://test.com/test", nil)
		So(err, ShouldBeNil)

		factoryMock := &encoder.FactoryMock{}
		clientMock := &http.NativeMock{}
		encoderMock := &encoder.Mock{}
		bodyMock := &http.BodyMock{}
		mocks := []interface{}{factoryMock, clientMock, encoderMock, bodyMock}

		client := http.NewClient(clientMock, factoryMock)

		doCall := clientMock.On("Do", req).Once()
		blank := struct{}{}

		Convey("should return no error when", func() {
			Convey("http response returns 200", func() {
				Convey("and body is empty", func() {
					doCall.Return(newResponse(native.StatusOK, nil), nil)

					err := client.Do(req, &blank)

					So(err, ShouldBeNil)
					mock.AssertExpectationsForObjects(t, mocks...)
				})
				Convey("and body is valid", func() {
					type T struct {
						Int int `json:"int"`
					}
					var actual T
					expected := T{Int: 1}

					resp := newResponse(native.StatusOK, actual)
					doCall.Return(resp, nil)

					factoryMock.On("Create", resp).Return(encoderMock).Once()
					encoderMock.On("Decode", mock.Anything, mock.Anything).
						Once().
						Return(nil).Run(func(args mock.Arguments) {
						tmp := args.Get(1).(*T)
						*tmp = expected
					})

					err := client.Do(req, &actual)

					So(err, ShouldBeNil)
					So(actual, ShouldResemble, expected)
					mock.AssertExpectationsForObjects(t, mocks...)
				})
			})
		})
		Convey("should return error when", func() {
			Convey("http do returns error", func() {
				doCall.Return(nil, errors.New("this is my boomstick"))
				err := client.Do(req, nil)

				So(err, ShouldBeError, "this is my boomstick")
				mock.AssertExpectationsForObjects(t, mocks...)
			})
			Convey("http response is >= 400", func() {
				doCall.Return(newResponse(native.StatusTeapot, "you are a teapot"), nil)

				err := client.Do(req, &blank)

				So(err, ShouldBeError, "418: you are a teapot: bad request")
				mock.AssertExpectationsForObjects(t, mocks...)
			})
			Convey("read body returns an error", func() {
				resp := newResponse(native.StatusOK, nil)
				resp.Body = bodyMock
				bodyMock.On("Close").Once().Return(nil)
				bodyMock.On("Read", mock.Anything).Once().Return(0, errors.New("everybody bock mock"))

				doCall.Return(resp, nil)

				err := client.Do(req, &blank)

				So(err, ShouldBeError, "everybody bock mock")
				mock.AssertExpectationsForObjects(t, mocks...)
			})
			Convey("decode returns an error", func() {
				resp := newResponse(native.StatusOK, "junk")
				doCall.Return(resp, nil)

				factoryMock.On("Create", resp).Return(encoderMock).Once()
				encoderMock.On("Decode", mock.Anything, mock.Anything).Return(errors.New("error decoding"))

				err := client.Do(req, &blank)

				So(err, ShouldBeError, "error decoding")
				mock.AssertExpectationsForObjects(t, mocks...)
			})
		})
	})
}

func newResponse(status int, data interface{}) *native.Response {
	var body []byte

	if data != nil {
		bts, err := encoder.NewJSON().Encode(data)
		if err != nil {
			panic(err)
		}

		body = bts
	}

	return &native.Response{
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
		StatusCode: status,
	}
}
