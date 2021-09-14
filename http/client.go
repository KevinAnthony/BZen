package http

import (
	"io/ioutil"
	native "net/http"
	"strings"

	"github.com/kevinanthony/bzen/http/encoder"

	"github.com/pkg/errors"
)

var errBadRequest = errors.New("request")

type Client interface {
	Do(req *native.Request, v interface{}) error
}

type Native interface {
	Do(req *native.Request) (*native.Response, error)
}

type client struct {
	encFactory encoder.Factory
	client     Native
}

func NewClient(nativeClient Native, enc encoder.Factory) Client {
	if nativeClient == nil {
		panic("http client is required")
	}

	if enc == nil {
		panic("encoding factory is required")
	}

	return &client{
		encFactory: enc,
		client:     nativeClient,
	}
}

func (c client) Do(req *native.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= native.StatusBadRequest {
		return errors.Wrapf(errBadRequest, "status: %d msg: %s",
			resp.StatusCode, strings.Trim(string(bts), "\""))
	}

	if len(bts) == 0 {
		return nil
	}

	return c.encFactory.Create(resp).Decode(bts, v)
}
