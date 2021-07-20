package server

import (
	"net/http"

	goji "goji.io"
	"goji.io/pat"
)

type HTTPServer struct {
	mux *goji.Mux
}

func NewServer() HTTPServer {
	srv := HTTPServer{
		mux: goji.NewMux(),
	}

	srv.mux.HandleFunc(pat.Post("/calc/add/two"), AddTwoNumbers)
	srv.mux.HandleFunc(pat.Post("/calc/multiply/two"), Multiply)
	srv.mux.HandleFunc(pat.Post("/calc/sum"), Sum)

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
