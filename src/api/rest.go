package main

import "net/http"

type httpAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

type restAPI interface {
	Do(req *restRequest) (*http.Response, error)
}

type restRequest struct {
	httpAPI httpAPI
}
