package server

import (
	"net/http"

	goji "goji.io"
)

type HTTPServer struct {
	mux *goji.Mux
}

func NewServer() HTTPServer {
	srv := HTTPServer{
		mux: goji.NewMux(),
	}

	return srv
}

func (s HTTPServer) Run() {
	svr := &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}

	svr.SetKeepAlivesEnabled(false)

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}
