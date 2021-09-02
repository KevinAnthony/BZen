package http

import "net/http"

// this is to hide all net/http types

type (
	MethodType string
)

const (
	MethodPost   = http.MethodPost
	MethodGet    = http.MethodGet
	MethodPut    = http.MethodPut
	MethodDelete = http.MethodDelete
)
