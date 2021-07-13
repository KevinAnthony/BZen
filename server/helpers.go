package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func UnmarshalFromRequest(res *http.Request, out interface{}) error{
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read from message body")
	}

	if err := json.Unmarshal(bts, out); err != nil {
		return errors.Wrap(err, "failed to unmarshal body")
	}

	return nil
}

func WriteError(w http.ResponseWriter, err error) {

}

func WriteSuccess(w http.ResponseWriter, out interface{}) {
	bts, err := json.Marshal(out)
	if err != nil {
		WriteError(w, errors.Wrap(err, "could not marshal output"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(bts); err != nil {
		WriteError(w, errors.Wrap(err, "could not write to response"))
	}
}
