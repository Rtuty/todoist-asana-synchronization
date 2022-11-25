package main

const apiToken = ""

const (
	apiBaseUrl string = "https://api.todoist.com/rest"
)

type Client struct {
	token   string
	restAPI restAPI
}

func New(token string) *Client {
	return &Client{
		token:   token,
		restAPI: newRESTClient(),
	}
}
