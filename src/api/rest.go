package main

import (
	"io"
	"net/http"
)

type httpAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

type restAPI interface {
	Do(req *restRequest) (*http.Response, error)
}

type restClient struct {
	httpAPI httpAPI
}

type restRequest struct {
	URL     string
	Method  string
	Payload map[string]interface{}
	Headers map[string]string
}

type restResponse struct {
	StatusCode int
	Body       io.Reader
}

func newRESTClient() *restClient {
	return &restClient{httpAPI: new(http.Client)}
}
